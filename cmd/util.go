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
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

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
			err = fmt.Errorf("please specify jsonpath using -o jsonpath=<path>")
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

type generator func(*schema.DSMetadata) (interface{}, error)

func paginateResp(gen generator, opts *schema.DSMetadata) (chan interface{}, error) {
	resp, err := gen(opts)
	if err != nil {
		return nil, err
	}
	var buffer int64
	buffer = 1
	v := reflect.ValueOf(resp).Elem().FieldByName("Metadata")
	metadata := v.Interface().(*schema.ListMetadata)

	var wg sync.WaitGroup
	if metadata.Length > 0 && metadata.TotalMatches > 0 {
		buffer = metadata.TotalMatches/metadata.Length + 1
	}
	responsechannel := make(chan interface{}, buffer)
	errorchannel := make(chan error, buffer)
	responsechannel <- resp

	if metadata.Length < metadata.TotalMatches {
		var i int64
		for i = *opts.Length; i < metadata.TotalMatches; i += *opts.Length {
			wg.Add(1)
			go func(i int64) {
				defer wg.Done()
				pagedopts := *opts
				pagedopts.Offset = utils.Int64Ptr(i)
				resp, err := gen(&pagedopts)
				if err != nil {
					errorchannel <- err
				}
				responsechannel <- resp
			}(i)

		}

	}
	go func() {
		wg.Wait()
		close(responsechannel)
		close(errorchannel)
	}()

	for err := range errorchannel {
		if err != nil {
			return nil, err
		}
	}

	return responsechannel, err
}
