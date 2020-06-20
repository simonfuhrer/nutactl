package cmd

import (
	"fmt"

	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/simonfuhrer/nutactl/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newVMListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List all VM",
		Aliases:               []string{"l", "li"},
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMList),
	}
	addOutputFormatFlags(cmd.Flags(), "table")
	flags := cmd.Flags()
	flags.StringP("filter", "f", "", "FIQL filter (e.g. vm_name==x3012.*)")
	flags.StringP("cluster", "c", "", "filter vms by cluster")
	return cmd
}

func runVMList(cli *CLI, cmd *cobra.Command, args []string) error {
	filter := viper.GetString("filter")
	cluster := viper.GetString("cluster")
	var list *schema.VMListIntent
	var err error

	var filtercluster string

	if cluster != "" {
		clustername, err := cli.Client().Cluster.Get(cli.Context, cluster)
		if err != nil {
			return err
		}
		filtercluster = fmt.Sprintf("cluster=in=%s", clustername.Metadata.UUID)
	}

	if filter != "" || cluster != "" {
		if filter != "" {
			filter = fmt.Sprintf("%s;%s", filter, filtercluster)
		} else {
			filter = filtercluster
		}
		list, err = cli.Client().VM.List(
			cli.Context,
			&schema.DSMetadata{Length: utils.Int64Ptr(500), Filter: filter},
		)
		if err != nil {
			return err
		}
	} else {
		list, err = cli.Client().VM.All(cli.Context)
		if err != nil {
			return err
		}
	}

	return outputResponse(displayers.VMs{VMListIntent: *list})
}
