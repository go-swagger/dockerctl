// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/config"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationConfigConfigDeleteCmd returns a cmd to handle operation configDelete
func makeOperationConfigConfigDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ConfigDelete",
		Short: ``,
		RunE:  runOperationConfigConfigDelete,
	}

	if err := registerOperationConfigConfigDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationConfigConfigDelete uses cmd flags to call endpoint api
func runOperationConfigConfigDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := config.NewConfigDeleteParams()
	if err, _ := retrieveOperationConfigConfigDeleteIDFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationConfigConfigDeleteResult(appCli.Config.ConfigDelete(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationConfigConfigDeleteIDFlag(m *config.ConfigDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// printOperationConfigConfigDeleteResult prints output to stdout
func printOperationConfigConfigDeleteResult(resp0 *config.ConfigDeleteNoContent, respErr error) error {
	if respErr != nil {

		// Non schema case: warning configDeleteNoContent is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*config.ConfigDeleteNotFound)
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
		resp2, ok := iResp2.(*config.ConfigDeleteInternalServerError)
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

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*config.ConfigDeleteServiceUnavailable)
		if ok {
			if !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		return respErr
	}

	// warning: non schema response configDeleteNoContent is not supported by go-swagger cli yet.

	return nil
}

// registerOperationConfigConfigDeleteParamFlags registers all flags needed to fill params
func registerOperationConfigConfigDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationConfigConfigDeleteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationConfigConfigDeleteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID of the config`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}
