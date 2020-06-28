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
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
	v2 "github.com/tecbiz-ch/nutanix-go-sdk/schema/v2"
)

func newVMCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] vmname",
		Short:                 "Create a VM",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMCreate),
	}
	flags := cmd.Flags()
	flags.StringP("cluster", "c", "", "Cluster Name or UUID)")
	flags.String("image", "", "Image Name or UUID")
	flags.String("vm", "", "VM Name or UUID")
	flags.Bool("start-after-create", false, "Start VM right after creation")
	flags.String("user-data", "", "Read user data from specified file")
	_ = cmd.MarkFlagFilename("user-data")
	markFlagsRequired(cmd, "cluster")
	addVMFlags(flags)

	return cmd
}

func runVMCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	vmIdorName := viper.GetString("vm")
	clusterIdorName := viper.GetString("cluster")
	imageIdorName := viper.GetString("image")

	startVM := viper.GetBool("start-after-create")
	userDataFileName := viper.GetString("user-data")

	if len(imageIdorName) > 0 && len(vmIdorName) > 0 {
		return fmt.Errorf("both image and vm provided")
	}

	if len(imageIdorName) == 0 && len(vmIdorName) == 0 {
		return fmt.Errorf("both image and vm are missing")
	}

	vmexists, err := cli.Client().VM.Get(cli.Context, name)
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return err
	}
	if vmexists != nil {
		return fmt.Errorf("vm %s already exists with uuid %s", name, vmexists.Metadata.UUID)
	}

	cluster, err := cli.Client().Cluster.Get(cli.Context, clusterIdorName)
	if err != nil {
		return err
	}

	req := &schema.VMIntent{
		Spec: &schema.VM{
			Resources: &schema.VMResources{
				SerialPortList: []*schema.VMSerialPort{
					{
						Index:       0,
						IsConnected: true,
					},
				},
				DiskList: []*schema.VMDisk{
					{
						DeviceProperties: &schema.VMDiskDeviceProperties{
							DeviceType: "DISK",
							DiskAddress: &schema.DiskAddress{
								//DeviceIndex: 0,
								AdapterType: "SCSI",
							},
						},
					},
				},
			},
			ClusterReference: &schema.Reference{
				Kind: "cluster",
				UUID: cluster.Metadata.UUID,
			},
		},
		Metadata: &schema.Metadata{
			Kind: "vm",
		},
	}

	err = createUpdateVMHelper(name, req, cli)
	if err != nil {
		return err
	}

	powerState := v2.PowerStateOFF
	if startVM {
		powerState = v2.PowerStateON
	}

	req.Spec.Resources.PowerState = string(powerState)

	if len(imageIdorName) != 0 {
		image, err := cli.Client().Image.Get(cli.Context, imageIdorName)
		if err != nil {
			return fmt.Errorf("image not found %s", imageIdorName)
		}
		req.Spec.Resources.DiskList[0].DataSourceReference = &schema.Reference{
			Kind: "image",
			UUID: image.Metadata.UUID,
		}
	}

	if len(vmIdorName) != 0 {
		vm, err := cli.Client().VM.Get(cli.Context, vmIdorName)
		if err != nil {
			return err
		}
		req.Spec.Resources.ParentReference = &schema.Reference{
			Kind: "vm",
			UUID: vm.Metadata.UUID,
		}
	}

	if len(userDataFileName) != 0 {

		userdata, err := ioutil.ReadFile(userDataFileName)
		if err != nil {
			return err
		}
		userdataEncoded := base64.StdEncoding.EncodeToString(userdata)
		metadata := schema.MetaData{
			Hostname: name,
		}
		metadatastr, err := metadata.ToBase64()
		if err != nil {
			return err
		}
		guestCustomization := &schema.GuestCustomization{
			CloudInit: &schema.GuestCustomizationCloudInit{
				UserData: userdataEncoded,
				MetaData: metadatastr,
			},
		}
		req.Spec.Resources.GuestCustomization = guestCustomization
	}

	result, err := cli.Client().VM.Create(cli.Context, req)
	if err != nil {
		return err
	}

	taskUUID := result.Status.ExecutionContext.TaskUUID.(string)

	err = cli.WaitTask(cli.Context, taskUUID, 180)
	if err != nil {
		errdelete := cli.Client().VM.Delete(cli.Context, result.Metadata.UUID)
		if errdelete != nil {
			return errors.Wrap(err, errdelete.Error())
		}
		return err
	}

	fmt.Printf("VM %s with uuid %s created\n", result.Spec.Name, result.Metadata.UUID)
	return nil
}
