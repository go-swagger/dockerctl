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

// makeOperationConfigConfigCreateCmd returns a cmd to handle operation configCreate
func makeOperationConfigConfigCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ConfigCreate",
		Short: ``,
		RunE:  runOperationConfigConfigCreate,
	}

	if err := registerOperationConfigConfigCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationConfigConfigCreate uses cmd flags to call endpoint api
func runOperationConfigConfigCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := config.NewConfigCreateParams()
	if err, _ := retrieveOperationConfigConfigCreateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationConfigConfigCreateResult(appCli.Config.ConfigCreate(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationConfigConfigCreateBodyFlag(m *config.ConfigCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := config.ConfigCreateBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in ConfigCreateBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = config.ConfigCreateBody{}
	}
	err, added := retrieveModelConfigCreateBodyFlags(0, &bodyValueModel, "configCreateBody", cmd)
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

// printOperationConfigConfigCreateResult prints output to stdout
func printOperationConfigConfigCreateResult(resp0 *config.ConfigCreateCreated, respErr error) error {
	if respErr != nil {
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

// registerOperationConfigConfigCreateParamFlags registers all flags needed to fill params
func registerOperationConfigConfigCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationConfigConfigCreateBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationConfigConfigCreateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	exampleBodyStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(bodyFlagName, "", fmt.Sprintf("Optional json string for [body] of form %v.", string(exampleBodyStr)))

	// add flags for body
	if err := registerModelConfigCreateBodyFlags(0, "configCreateBody", cmd); err != nil {
		return err
	}

	return nil
}

// register flags to command
func registerModelConfigCreateBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// allOf ConfigCreateParamsBodyAO0 is not supported by go-swwagger cli yet

	// allOf ConfigCreateParamsBodyAO1 is not supported by go-swwagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelConfigCreateBodyFlags(depth int, m *config.ConfigCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// allOf ConfigCreateParamsBodyAO0 is not supported by go-swwagger cli yet

	// allOf ConfigCreateParamsBodyAO1 is not supported by go-swwagger cli yet

	return nil, retAdded
}

// interface{} register and retrieve functions are not rendered by go-swagger cli