// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/service"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServiceServiceCreateCmd returns a cmd to handle operation serviceCreate
func makeOperationServiceServiceCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ServiceCreate",
		Short: ``,
		RunE:  runOperationServiceServiceCreate,
	}

	if err := registerOperationServiceServiceCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServiceServiceCreate uses cmd flags to call endpoint api
func runOperationServiceServiceCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := service.NewServiceCreateParams()
	if err, _ := retrieveOperationServiceServiceCreateXRegistryAuthFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServiceServiceCreateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationServiceServiceCreateResult(appCli.Service.ServiceCreate(params)); err != nil {
		return err
	}
	return nil
}

// registerOperationServiceServiceCreateParamFlags registers all flags needed to fill params
func registerOperationServiceServiceCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServiceServiceCreateXRegistryAuthParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServiceServiceCreateBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServiceServiceCreateXRegistryAuthParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xRegistryAuthDescription := `A base64-encoded auth configuration for pulling from private registries. [See the authentication section for details.](#section/Authentication)`

	var xRegistryAuthFlagName string
	if cmdPrefix == "" {
		xRegistryAuthFlagName = "X-Registry-Auth"
	} else {
		xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
	}

	var xRegistryAuthFlagDefault string

	_ = cmd.PersistentFlags().String(xRegistryAuthFlagName, xRegistryAuthFlagDefault, xRegistryAuthDescription)

	return nil
}
func registerOperationServiceServiceCreateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	exampleBodyStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(bodyFlagName, "", fmt.Sprintf("Optional json string for [body] of form %v.", string(exampleBodyStr)))

	// add flags for body
	if err := registerModelServiceCreateBodyFlags(0, "serviceCreateBody", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationServiceServiceCreateXRegistryAuthFlag(m *service.ServiceCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Registry-Auth") {

		var xRegistryAuthFlagName string
		if cmdPrefix == "" {
			xRegistryAuthFlagName = "X-Registry-Auth"
		} else {
			xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
		}

		xRegistryAuthFlagValue, err := cmd.Flags().GetString(xRegistryAuthFlagName)
		if err != nil {
			return err, false
		}
		m.XRegistryAuth = &xRegistryAuthFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServiceServiceCreateBodyFlag(m *service.ServiceCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := service.ServiceCreateBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in ServiceCreateBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = service.ServiceCreateBody{}
	}
	err, added := retrieveModelServiceCreateBodyFlags(0, &bodyValueModel, "serviceCreateBody", cmd)
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

// printOperationServiceServiceCreateResult prints output to stdout
func printOperationServiceServiceCreateResult(resp0 *service.ServiceCreateCreated, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*service.ServiceCreateCreated)
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
		resp1, ok := iResp1.(*service.ServiceCreateBadRequest)
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
		resp2, ok := iResp2.(*service.ServiceCreateForbidden)
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
		resp3, ok := iResp3.(*service.ServiceCreateConflict)
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

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*service.ServiceCreateInternalServerError)
		if ok {
			if !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		var iResp5 interface{} = respErr
		resp5, ok := iResp5.(*service.ServiceCreateServiceUnavailable)
		if ok {
			if !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
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
func registerModelServiceCreateBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.ServiceSpec flags

	if err := registerModelServiceSpecFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// ServiceCreateParamsBodyAO1 ServiceCreateParamsBodyAllOf1 register is skipped

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceCreateBodyFlags(depth int, m *service.ServiceCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.ServiceSpec
	err, serviceCreateParamsBodyAO0Added := retrieveModelServiceSpecFlags(depth, &m.ServiceSpec, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serviceCreateParamsBodyAO0Added

	return nil, retAdded
}

// register flags to command
func registerModelServiceCreateCreatedBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceCreateCreatedBodyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceCreateCreatedBodyWarning(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceCreateCreatedBodyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of the created service.`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "ID"
	} else {
		idFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerServiceCreateCreatedBodyWarning(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	warningDescription := `Optional warning message`

	var warningFlagName string
	if cmdPrefix == "" {
		warningFlagName = "Warning"
	} else {
		warningFlagName = fmt.Sprintf("%v.Warning", cmdPrefix)
	}

	var warningFlagDefault string

	_ = cmd.PersistentFlags().String(warningFlagName, warningFlagDefault, warningDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceCreateCreatedBodyFlags(depth int, m *service.ServiceCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveServiceCreateCreatedBodyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, warningAdded := retrieveServiceCreateCreatedBodyWarningFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || warningAdded

	return nil, retAdded
}

func retrieveServiceCreateCreatedBodyIDFlags(depth int, m *service.ServiceCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.ID", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "ID"
		} else {
			idFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServiceCreateCreatedBodyWarningFlags(depth int, m *service.ServiceCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	warningFlagName := fmt.Sprintf("%v.Warning", cmdPrefix)
	if cmd.Flags().Changed(warningFlagName) {

		var warningFlagName string
		if cmdPrefix == "" {
			warningFlagName = "Warning"
		} else {
			warningFlagName = fmt.Sprintf("%v.Warning", cmdPrefix)
		}

		warningFlagValue, err := cmd.Flags().GetString(warningFlagName)
		if err != nil {
			return err, false
		}
		m.Warning = warningFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// interface{} register and retrieve functions are not rendered by go-swagger cli
