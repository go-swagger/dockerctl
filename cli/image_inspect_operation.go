// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/image"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationImageImageInspectCmd returns a cmd to handle operation imageInspect
func makeOperationImageImageInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ImageInspect",
		Short: `Return low-level information about an image.`,
		RunE:  runOperationImageImageInspect,
	}

	if err := registerOperationImageImageInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageInspect uses cmd flags to call endpoint api
func runOperationImageImageInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageInspectParams()
	if err, _ := retrieveOperationImageImageInspectNameFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationImageImageInspectResult(appCli.Image.ImageInspect(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationImageImageInspectNameFlag(m *image.ImageInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationImageImageInspectResult prints output to stdout
func printOperationImageImageInspectResult(resp0 *image.ImageInspectOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*image.ImageInspectOK)
		if ok {
			if !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*image.ImageInspectNotFound)
		if ok {
			if !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*image.ImageInspectInternalServerError)
		if ok {
			if !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		return respErr
	}

	if !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return err
		}
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationImageImageInspectParamFlags registers all flags needed to fill params
func registerOperationImageImageInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageInspectNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageInspectNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Image name or id`

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
