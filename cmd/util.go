// Copyright © 2020 Simon Fuhrer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func GenerateMac() net.HardwareAddr {
	buf := make([]byte, 6)
	var mac net.HardwareAddr

	_, err := rand.Read(buf)
	if err != nil {
	}

	// Set the local bit
	buf[0] |= 2

	mac = append(mac, buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])

	return mac
}

// MarkFlagsRequired ...
func MarkFlagsRequired(cmd *cobra.Command, names ...string) {
	// cobra does not merge its local flagset with the persistent flagset
	// when Flags() is called. By calling InheretedFlags(), we force a merge.

	cmd.InheritedFlags()
	for _, name := range names {
		if !viper.IsSet(name) || viper.GetString(name) == "" {
			if err := cobra.MarkFlagRequired(cmd.Flags(), name); err != nil {
				panic(err)
			}
		}
	}
}

func addOutputFormatFlags(flags *pflag.FlagSet, defaultformat string) {
	flags.StringP("output", "o", defaultformat, "json|yaml|table")
}

// BindAllFlags ...
func BindAllFlags(cmd *cobra.Command) {
	_ = viper.BindPFlags(cmd.Flags())
	_ = viper.BindPFlags(cmd.PersistentFlags())
}

func checkErr(err error) error {
	if err == nil {
		return nil
	}
	logrus.Error(err.Error())
	os.Exit(255)
	return nil
}

func outputResponse(d displayers.Displayable) error {
	var err error
	outputFormat := viper.GetString("output")
	switch {
	case outputFormat == "table":
		err = d.TableData(os.Stdout)
	case outputFormat == "json":
		err = d.JSON(os.Stdout)
	case outputFormat == "yaml":
		err = d.YAML(os.Stdout)
	case outputFormat == "pp":
		err = d.PP(os.Stdout)
	case outputFormat == "text":
		err = d.Text(os.Stdout)
	case strings.HasPrefix(outputFormat, "jsonpath"):
		fields := strings.SplitN(outputFormat, "=", 2)
		if len(fields) != 2 {
			err = errors.New("please specify jsonpath using -o jsonpath=<path>")
			break
		}
		template := fields[1]
		err = d.JSONPath(os.Stdout, template)
	default:
		// TODO handle this using cobra itself?
		err = errors.Errorf("output format %s not supported", outputFormat)
	}

	return err
}

func warnConfirm(msg string, args ...interface{}) {
	fmt.Fprintf(color.Output, "%s: %s", colorWarn, fmt.Sprintf(msg, args...))
}

var retrieveUserInput = func(message string) (string, error) {
	return readUserInput(os.Stdin, message)
}

func readUserInput(in io.Reader, message string) (string, error) {
	reader := bufio.NewReader(in)
	warnConfirm("Are you sure you want to " + message + " (y/N) ? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	answer = strings.TrimRight(answer, "\r\n")

	return strings.ToLower(answer), nil
}

func askForConfirm(message string) error {
	answer, err := retrieveUserInput(message)
	if err != nil {
		return fmt.Errorf("unable to parse users input: %s", err)
	}

	if answer != "y" && answer != "ye" && answer != "yes" {
		return fmt.Errorf("invalid user input")
	}

	return nil
}
