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

// makeOperationSecretSecretCreateCmd returns a cmd to handle operation secretCreate
func makeOperationSecretSecretCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SecretCreate",
		Short: ``,
		RunE:  runOperationSecretSecretCreate,
	}

	if err := registerOperationSecretSecretCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSecretSecretCreate uses cmd flags to call endpoint api
func runOperationSecretSecretCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := secret.NewSecretCreateParams()
	if err, _ := retrieveOperationSecretSecretCreateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationSecretSecretCreateResult(appCli.Secret.SecretCreate(params)); err != nil {
		return err
	}
	return nil
}

// registerOperationSecretSecretCreateParamFlags registers all flags needed to fill params
func registerOperationSecretSecretCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSecretSecretCreateBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSecretSecretCreateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	exampleBodyStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(bodyFlagName, "", fmt.Sprintf("Optional json string for [body] of form %v.", string(exampleBodyStr)))

	// add flags for body
	if err := registerModelSecretCreateBodyFlags(0, "secretCreateBody", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationSecretSecretCreateBodyFlag(m *secret.SecretCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := secret.SecretCreateBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in SecretCreateBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = secret.SecretCreateBody{}
	}
	err, added := retrieveModelSecretCreateBodyFlags(0, &bodyValueModel, "secretCreateBody", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	bodyValueDebugBytes, err := json.Marshal(m.Body)
	if err != nil {
		return err, false
	}
	logDebugf("Body payload: %v", string(bodyValueDebugBytes))
	return nil, retAdded
}

// printOperationSecretSecretCreateResult prints output to stdout
func printOperationSecretSecretCreateResult(resp0 *secret.SecretCreateCreated, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*secret.SecretCreateCreated)
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
		resp1, ok := iResp1.(*secret.SecretCreateConflict)
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
		resp2, ok := iResp2.(*secret.SecretCreateInternalServerError)
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
		resp3, ok := iResp3.(*secret.SecretCreateServiceUnavailable)
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

// register flags to command
func registerModelSecretCreateBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.SecretSpec flags

	if err := registerModelSecretSpecFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// SecretCreateParamsBodyAO1 SecretCreateParamsBodyAllOf1 register is skipped

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSecretCreateBodyFlags(depth int, m *secret.SecretCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.SecretSpec
	err, secretCreateParamsBodyAO0Added := retrieveModelSecretSpecFlags(depth, &m.SecretSpec, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secretCreateParamsBodyAO0Added

	return nil, retAdded
}

// interface{} register and retrieve functions are not rendered by go-swagger cli
