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
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newSubnetCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "subnet",
		Short:                 "Manage subnets",
		Aliases:               []string{"s", "sub"},
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		RunE:                  cli.wrap(runSubnet),
	}
	cmd.AddCommand(
		newSubnetListCommand(cli),
		newSubnetDescribeCommand(cli),
		newSubnetCreateCommand(cli),
		newSubnetUpdateCommand(cli),
		newSubnetDeleteCommand(cli),
	)
	return cmd
}

func runSubnet(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func createUpdateSubnetHelper(name string, req *schema.SubnetIntent) error {
	description := viper.GetString("description")
	vlanID := viper.GetInt64("vlan-id")
	networkCIDR := viper.GetString("ip-range")
	gateway := viper.GetString("gateway")
	domain := viper.GetString("domain")
	dnsservers := viper.GetStringSlice("dns-servers")
	ippool := viper.GetStringSlice("ip-pool")

	if vlanID > 4094 {
		return fmt.Errorf("--vlanid the maximum number of VLANs is 4094")
	}

	if req.Spec.Resources.IPConfig == nil {
		if len(gateway) != 0 || len(networkCIDR) != 0 || len(ippool) > 0 || len(dnsservers) > 0 || len(domain) != 0 || len(gateway) != 0 {
			req.Spec.Resources.IPConfig = &schema.IPConfig{}
		}
		if len(domain) != 0 || len(dnsservers) > 0 {
			req.Spec.Resources.IPConfig.DHCPOptions = &schema.DHCPOptions{}
		}
	}

	if vlanID > 0 {
		req.Spec.Resources.VlanID = &vlanID
	}

	if len(name) != 0 {
		req.Spec.Name = name

	}

	if len(description) != 0 {
		req.Spec.Description = description
	}

	if len(networkCIDR) != 0 {
		cidr, network, err := net.ParseCIDR(networkCIDR)
		if err != nil {
			return err
		}
		req.Spec.Resources.IPConfig.SubnetIP = cidr.String()
		prefixSize, _ := network.Mask.Size()
		req.Spec.Resources.IPConfig.PrefixLength = int64(prefixSize)

	}
	if len(gateway) != 0 {
		req.Spec.Resources.IPConfig.DefaultGatewayIP = gateway
	}
	if len(domain) != 0 {
		req.Spec.Resources.IPConfig.DHCPOptions.DomainName = domain
		req.Spec.Resources.IPConfig.DHCPOptions.DomainSearchList = []string{domain}
	}
	if len(dnsservers) > 0 {
		req.Spec.Resources.IPConfig.DHCPOptions.DomainNameServerList = dnsservers
	}

	if len(ippool) > 0 {
		if len(ippool) != 2 {
			return fmt.Errorf("for --ip-pool your must provide a start and a end address")
		}
		req.Spec.Resources.IPConfig.PoolList = []*schema.IPPool{
			{
				Range: fmt.Sprintf("%s %s", ippool[0], ippool[1]),
			},
		}
	}
	return nil
}

func addSubnetFlags(flags *pflag.FlagSet) {
	flags.StringP("description", "d", "", "Description")
	flags.Int64("vlan-id", 0, "VlanID")
	flags.String("ip-range", "", "Network CIDR")
	flags.String("gateway", "", "Default Gateway IP")
	flags.String("domain", "", "Default Domainname")
	flags.StringSlice("dns-servers", nil, "Default DNS Servers seperated with a comma")
	flags.StringSlice("ip-pool", nil, "Start address to end address seperated with a comma")
}
