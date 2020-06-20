package cmd

import (
	"bufio"
	"fmt"
	"io"
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

// MarkFlagsRequired ...
func MarkFlagsRequired(cmd *cobra.Command, names ...string) {
	// cobra does not merge its local flagset with the persistent flagset
	// when Flags() is called. By calling InheretedFlags(), we force a merge.

	cmd.InheritedFlags()
	for _, name := range names {
		if !viper.IsSet(name) {
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

/* func bindAllFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		fmt.Println(f.Name)
		pp.Println(cmd.PersistentFlags().Lookup(f.Name))
		if err := v.BindPFlag(f.Name, cmd.PersistentFlags().Lookup(f.Name)); err != nil {
			panic(err) // Should never happen
		}
	})
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if err := v.BindPFlag(f.Name, cmd.Flags().Lookup(f.Name)); err != nil {
			panic(err) // Should never happen
		}
	})
} */

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
