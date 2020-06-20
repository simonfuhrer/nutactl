package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newSubnetCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] subnetname",
		Short:                 "Create an subnet",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runSubnetCreate),
	}
	flags := cmd.Flags()
	flags.String("cluster", "", "Cluster (UUID or name)")
	MarkFlagsRequired(cmd, "cluster")
	addSubnetFlags(flags)

	return cmd
}

func runSubnetCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	clusterIdorName := viper.GetString("cluster")

	subnet, _ := cli.Client().Subnet.Get(cli.Context, name)
	if subnet != nil {
		return fmt.Errorf("subnet %s already exists with uuid %s", name, subnet.Metadata.UUID)
	}
	cluster, err := cli.Client().Cluster.Get(cli.Context, clusterIdorName)
	if err != nil {
		return fmt.Errorf("cluster not found %s", clusterIdorName)
	}

	req := &schema.SubnetIntent{
		Spec: &schema.Subnet{
			Resources: &schema.SubnetResources{
				SubnetType: "VLAN",
			},
			ClusterReference: &schema.Reference{
				Kind: "cluster",
				UUID: cluster.Metadata.UUID,
			},
		},
		Metadata: &schema.Metadata{
			Kind: "subnet",
		},
	}

	err = createUpdateSubnetHelper(name, req)
	if err != nil {
		return err
	}

	result, err := cli.Client().Subnet.Create(cli.Context, req)
	if err != nil {
		return err
	}

	fmt.Printf("Subnet %s with uuid %s created\n", result.Spec.Name, result.Metadata.UUID)
	return nil
}
