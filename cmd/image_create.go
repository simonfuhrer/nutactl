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
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newImageCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] imagename",
		Short:                 "Create an image",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runImageCreate),
	}
	flags := cmd.Flags()
	flags.StringP("source-uri", "s", "", "source image URI")
	flags.StringP("source-path", "f", "", "source image path")
	_ = cmd.MarkFlagFilename("source-path")
	flags.StringP("type", "t", "disk", "image type (iso or disk")
	flags.BoolP("wait", "w", true, "wait to be completed")
	addImageFlags(flags)

	return cmd
}

func runImageCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	sourceURI, _ := cmd.Flags().GetString("source-uri")
	sourcePath, _ := cmd.Flags().GetString("source-path")
	imagetype, _ := cmd.Flags().GetString("type")
	waitfor, _ := cmd.Flags().GetBool("wait")

	if len(sourceURI) > 0 && len(sourcePath) > 0 {
		return fmt.Errorf("both source-uri and source-path provided")
	}

	if len(sourceURI) == 0 && len(sourcePath) == 0 {
		return fmt.Errorf("both source-uri and source-path are missing")
	}

	if len(sourcePath) != 0 {
		_, err := utils.Exists(sourcePath)
		waitfor = true
		if err != nil {
			return err
		}
	}

	req := &schema.ImageIntent{
		Spec: &schema.Image{
			Resources: &schema.ImageResources{
				ImageType: "DISK_IMAGE",
			},
		},
		Metadata: &schema.Metadata{
			Kind: "image",
		},
	}

	if imagetype == "iso" {
		req.Spec.Resources.ImageType = "ISO_IMAGE"
	}

	if len(sourceURI) != 0 {
		req.Spec.Resources.SourceURI = sourceURI
	}

	err := createUpdateImageHelper(name, req)
	if err != nil {
		return err
	}

	result, err := cli.Client().Image.Create(cli.Context, req)
	if err != nil {
		return err
	}

	if waitfor {
		taskUUID := result.Status.ExecutionContext.TaskUUID.(string)
		err := cli.WaitTask(cli.Context, taskUUID, 180)
		if err != nil {
			errdelete := cli.Client().Image.Delete(cli.Context, result.Metadata.UUID)
			if errdelete != nil {
				return errors.Wrap(err, errdelete.Error())
			}
			return err
		}
	}

	if len(sourcePath) != 0 {
		file, err := os.Open(sourcePath)
		if err != nil {
			return err
		}
		defer file.Close()

		fileContents, err := ioutil.ReadAll(file)
		if err != nil {
			return fmt.Errorf("cannot read file %s", err)
		}

		_, err = cli.Client().Image.Upload(cli.Context, result.Metadata.UUID, fileContents)
		if err != nil {
			err := cli.Client().Image.Delete(cli.Context, result.Metadata.UUID)
			return err
		}
	}

	fmt.Printf("Image %s with uuid %s created\n", result.Spec.Name, result.Metadata.UUID)

	return nil
}
