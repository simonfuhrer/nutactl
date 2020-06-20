package main

import (
	"github.com/simonfuhrer/nutactl/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	c := cmd.NewCLI()
	cobra.EnableCommandSorting = false
	cobra.EnablePrefixMatching = false
	if err := c.RootCommand.Execute(); err != nil {
		logrus.Errorf("error: %s", err.Error())

	}
}
