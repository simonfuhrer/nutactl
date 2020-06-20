package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newProjectCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "project",
		Short:                 "Manage projects",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runProject),
	}
	cmd.AddCommand(
		newProjectListCommand(cli),
		newProjectDescribeCommand(cli),
		newProjectUpdateCommand(cli),
		newProjectCreateCommand(cli),
		newProjectDeleteCommand(cli),
	)
	return cmd
}

func runProject(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func createUpdateProjectHelper(name string, req *schema.ProjectIntent) error {
	description := viper.GetString("description")

	if len(name) != 0 {
		req.Spec.Name = name
	}
	if len(description) != 0 {
		req.Spec.Description = description
		req.Metadata.UseCategoriesMapping = true
		m := make(map[string][]string)
		m["BI-ForemanOperatingSystem"] = []string{"Windows_2016"}
		req.Metadata.CategoriesMapping = nil

		/* 		m := make(map[string]string)
		   		m["BI-ForemanOperatingSystem"] = "Windows_20162"
		   		project.Metadata.Categories = m */
	}
	return nil
}

func addProjectFlags(flags *pflag.FlagSet) {
	flags.StringP("description", "d", "", "Description")
}
