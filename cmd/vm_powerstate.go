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

	"github.com/spf13/cobra"
	v2 "github.com/tecbiz-ch/nutanix-go-sdk/schema/v2"
)

func newVMPowerStateOnCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "poweron [FLAGS] VM",
		Short:                 "Poweron a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      false,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMPowerStateOn),
	}

	return cmd
}

func newVMPowerStateOffCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "poweroff [FLAGS] VM",
		Short:                 "Poweroff a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      false,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMPowerStateOff),
	}

	return cmd
}

func newVMPowerStateRebootCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "reboot [FLAGS] VM",
		Short:                 "Reboot a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      false,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMPowerStateReboot),
	}

	return cmd
}

func newVMPowerStateResetCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "powerreset [FLAGS] VM",
		Short:                 "Powerreset a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      false,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMPowerStateReset),
	}

	return cmd
}

func newVMPowerStateShutdownCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "shutdown [FLAGS] VM",
		Short:                 "Shutdown a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      false,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMPowerStateShutdown),
	}

	return cmd
}

func runVMPowerStateShutdown(cli *CLI, cmd *cobra.Command, args []string) error {
	return changeVMPowerState(cli, v2.PowerStateACPISHUTDOWN, args)
}

func runVMPowerStateReboot(cli *CLI, cmd *cobra.Command, args []string) error {
	return changeVMPowerState(cli, v2.PowerStateACPIREBOOT, args)
}

func runVMPowerStateReset(cli *CLI, cmd *cobra.Command, args []string) error {
	return changeVMPowerState(cli, v2.PowerStateRESET, args)
}

func runVMPowerStateOff(cli *CLI, cmd *cobra.Command, args []string) error {
	return changeVMPowerState(cli, v2.PowerStateOFF, args)
}

func runVMPowerStateOn(cli *CLI, cmd *cobra.Command, args []string) error {
	return changeVMPowerState(cli, v2.PowerStateON, args)
}

func changeVMPowerState(cli *CLI, powerState v2.PowerState, args []string) error {
	idOrName := args[0]

	vm, err := cli.Client().VM.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}
	if vm == nil {
		return fmt.Errorf("vm not found: %s", idOrName)
	}
	task, err := cli.Client().VM.SetPowerState(cli.Context, powerState, vm)
	if err != nil {
		return err
	}
	err = cli.WaitTask(cli.Context, task.TaskUUID, 180)
	if err != nil {
		return err
	}
	fmt.Printf("VM State changed to: %v\n", powerState)
	return nil
}
