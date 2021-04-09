// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/volume"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVolumeVolumeInspectCmd returns a cmd to handle operation volumeInspect
func makeOperationVolumeVolumeInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "VolumeInspect",
		Short: ``,
		RunE:  runOperationVolumeVolumeInspect,
	}

	if err := registerOperationVolumeVolumeInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVolumeVolumeInspect uses cmd flags to call endpoint api
func runOperationVolumeVolumeInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := volume.NewVolumeInspectParams()
	if err, _ := retrieveOperationVolumeVolumeInspectNameFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationVolumeVolumeInspectResult(appCli.Volume.VolumeInspect(params)); err != nil {
		return err
	}
	return nil
}

// registerOperationVolumeVolumeInspectParamFlags registers all flags needed to fill params
func registerOperationVolumeVolumeInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVolumeVolumeInspectNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVolumeVolumeInspectNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Volume name or ID`

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

func retrieveOperationVolumeVolumeInspectNameFlag(m *volume.VolumeInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationVolumeVolumeInspectResult prints output to stdout
func printOperationVolumeVolumeInspectResult(resp0 *volume.VolumeInspectOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*volume.VolumeInspectOK)
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
		resp1, ok := iResp1.(*volume.VolumeInspectNotFound)
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
		resp2, ok := iResp2.(*volume.VolumeInspectInternalServerError)
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
