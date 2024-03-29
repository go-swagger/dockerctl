// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/image"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationImageImageGetCmd returns a cmd to handle operation imageGet
func makeOperationImageImageGetCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "ImageGet",
		Short: `Get a tarball containing all images and metadata for a repository.

If ` + "`" + `name` + "`" + ` is a specific name and tag (e.g. ` + "`" + `ubuntu:latest` + "`" + `), then only that image (and its parents) are returned. If ` + "`" + `name` + "`" + ` is an image ID, similarly only that image (and its parents) are returned, but with the exclusion of the ` + "`" + `repositories` + "`" + ` file in the tarball, as there were no image names referenced.

### Image tarball format

An image tarball contains one directory per image layer (named using its long ID), each containing these files:

- ` + "`" + `VERSION` + "`" + `: currently ` + "`" + `1.0` + "`" + ` - the file format version
- ` + "`" + `json` + "`" + `: detailed layer information, similar to ` + "`" + `docker inspect layer_id` + "`" + `
- ` + "`" + `layer.tar` + "`" + `: A tarfile containing the filesystem changes in this layer

The ` + "`" + `layer.tar` + "`" + ` file contains ` + "`" + `aufs` + "`" + ` style ` + "`" + `.wh..wh.aufs` + "`" + ` files and directories for storing attribute changes and deletions.

If the tarball defines a repository, the tarball should also include a ` + "`" + `repositories` + "`" + ` file at the root that contains a list of repository and tag names mapped to layer IDs.

` + "`" + `` + "`" + `` + "`" + `json
{
  "hello-world": {
    "latest": "565a9d68a73f6706862bfe8409a7f659776d4d60a8d096eb4a3cbce6999cc2a1"
  }
}
` + "`" + `` + "`" + `` + "`" + `
`,
		RunE: runOperationImageImageGet,
	}

	if err := registerOperationImageImageGetParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageGet uses cmd flags to call endpoint api
func runOperationImageImageGet(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageGetParams()
	if err, _ := retrieveOperationImageImageGetNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImageImageGetResult(appCli.Image.ImageGet(params, &bytes.Buffer{}))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImageImageGetParamFlags registers all flags needed to fill params
func registerOperationImageImageGetParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageGetNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageGetNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Image name or ID`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func retrieveOperationImageImageGetNameFlag(m *image.ImageGetParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

	}
	return nil, retAdded
}

// parseOperationImageImageGetResult parses request result and return the string content
func parseOperationImageImageGetResult(resp0 *image.ImageGetOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*image.ImageGetOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*image.ImageGetInternalServerError)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}
