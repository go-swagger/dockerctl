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

// makeOperationImageImageGetAllCmd returns a cmd to handle operation imageGetAll
func makeOperationImageImageGetAllCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "ImageGetAll",
		Short: `Get a tarball containing all images and metadata for several image repositories.

For each value of the ` + "`" + `names` + "`" + ` parameter: if it is a specific name and tag (e.g. ` + "`" + `ubuntu:latest` + "`" + `), then only that image (and its parents) are returned; if it is an image ID, similarly only that image (and its parents) are returned and there would be no names referenced in the 'repositories' file for this image ID.

For details on the format, see [the export image endpoint](#operation/ImageGet).
`,
		RunE: runOperationImageImageGetAll,
	}

	if err := registerOperationImageImageGetAllParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageGetAll uses cmd flags to call endpoint api
func runOperationImageImageGetAll(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageGetAllParams()
	if err, _ := retrieveOperationImageImageGetAllNamesFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImageImageGetAllResult(appCli.Image.ImageGetAll(params, &bytes.Buffer{}))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImageImageGetAllParamFlags registers all flags needed to fill params
func registerOperationImageImageGetAllParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageGetAllNamesParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageGetAllNamesParamFlags(cmdPrefix string, cmd *cobra.Command) error {
	// warning: go type []string is not supported by go-swagger cli yet.
	return nil
}

func retrieveOperationImageImageGetAllNamesFlag(m *image.ImageGetAllParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("names") {
		// warning: names array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

// parseOperationImageImageGetAllResult parses request result and return the string content
func parseOperationImageImageGetAllResult(resp0 *image.ImageGetAllOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*image.ImageGetAllOK)
		if ok {
			if !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*image.ImageGetAllInternalServerError)
		if ok {
			if !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0.Payload) {
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}
