// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/secret"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSecretSecretInspectCmd returns a cmd to handle operation secretInspect
func makeOperationSecretSecretInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SecretInspect",
		Short: ``,
		RunE:  runOperationSecretSecretInspect,
	}

	if err := registerOperationSecretSecretInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSecretSecretInspect uses cmd flags to call endpoint api
func runOperationSecretSecretInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := secret.NewSecretInspectParams()
	if err, _ := retrieveOperationSecretSecretInspectIDFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationSecretSecretInspectResult(appCli.Secret.SecretInspect(params)); err != nil {
		return err
	}
	return nil
}

// registerOperationSecretSecretInspectParamFlags registers all flags needed to fill params
func registerOperationSecretSecretInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSecretSecretInspectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSecretSecretInspectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID of the secret`

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

func retrieveOperationSecretSecretInspectIDFlag(m *secret.SecretInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationSecretSecretInspectResult prints output to stdout
func printOperationSecretSecretInspectResult(resp0 *secret.SecretInspectOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*secret.SecretInspectOK)
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
		resp1, ok := iResp1.(*secret.SecretInspectNotFound)
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
		resp2, ok := iResp2.(*secret.SecretInspectInternalServerError)
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
		resp3, ok := iResp3.(*secret.SecretInspectServiceUnavailable)
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

	if !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return err
		}
		fmt.Println(string(msgStr))
	}

	return nil
}
