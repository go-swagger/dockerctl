// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/service"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServiceServiceUpdateCmd returns a cmd to handle operation serviceUpdate
func makeOperationServiceServiceUpdateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ServiceUpdate",
		Short: ``,
		RunE:  runOperationServiceServiceUpdate,
	}

	if err := registerOperationServiceServiceUpdateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServiceServiceUpdate uses cmd flags to call endpoint api
func runOperationServiceServiceUpdate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := service.NewServiceUpdateParams()
	if err, _ := retrieveOperationServiceServiceUpdateXRegistryAuthFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServiceServiceUpdateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServiceServiceUpdateIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServiceServiceUpdateRegistryAuthFromFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServiceServiceUpdateRollbackFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServiceServiceUpdateVersionFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationServiceServiceUpdateResult(appCli.Service.ServiceUpdate(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationServiceServiceUpdateXRegistryAuthFlag(m *service.ServiceUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServiceServiceUpdateBodyFlag(m *service.ServiceUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := service.ServiceUpdateBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in ServiceUpdateBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = service.ServiceUpdateBody{}
	}
	err, added := retrieveModelServiceUpdateBodyFlags(0, &bodyValueModel, "serviceUpdateBody", cmd)
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
func retrieveOperationServiceServiceUpdateIDFlag(m *service.ServiceUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServiceServiceUpdateRegistryAuthFromFlag(m *service.ServiceUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("registryAuthFrom") {

		var registryAuthFromFlagName string
		if cmdPrefix == "" {
			registryAuthFromFlagName = "registryAuthFrom"
		} else {
			registryAuthFromFlagName = fmt.Sprintf("%v.registryAuthFrom", cmdPrefix)
		}

		registryAuthFromFlagValue, err := cmd.Flags().GetString(registryAuthFromFlagName)
		if err != nil {
			return err, false
		}
		m.RegistryAuthFrom = &registryAuthFromFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServiceServiceUpdateRollbackFlag(m *service.ServiceUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("rollback") {

		var rollbackFlagName string
		if cmdPrefix == "" {
			rollbackFlagName = "rollback"
		} else {
			rollbackFlagName = fmt.Sprintf("%v.rollback", cmdPrefix)
		}

		rollbackFlagValue, err := cmd.Flags().GetString(rollbackFlagName)
		if err != nil {
			return err, false
		}
		m.Rollback = &rollbackFlagValue

	}
	return nil, retAdded
}
func retrieveOperationServiceServiceUpdateVersionFlag(m *service.ServiceUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("version") {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetInt64(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = versionFlagValue

	}
	return nil, retAdded
}

// printOperationServiceServiceUpdateResult prints output to stdout
func printOperationServiceServiceUpdateResult(resp0 *service.ServiceUpdateOK, respErr error) error {
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

// registerOperationServiceServiceUpdateParamFlags registers all flags needed to fill params
func registerOperationServiceServiceUpdateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServiceServiceUpdateXRegistryAuthParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServiceServiceUpdateBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServiceServiceUpdateIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServiceServiceUpdateRegistryAuthFromParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServiceServiceUpdateRollbackParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServiceServiceUpdateVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServiceServiceUpdateXRegistryAuthParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationServiceServiceUpdateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	exampleBodyStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(bodyFlagName, "", fmt.Sprintf("Optional json string for [body] of form %v.", string(exampleBodyStr)))

	// add flags for body
	if err := registerModelServiceUpdateBodyFlags(0, "serviceUpdateBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationServiceServiceUpdateIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID or name of service.`

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
func registerOperationServiceServiceUpdateRegistryAuthFromParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	registryAuthFromDescription := `If the ` + "`" + `X-Registry-Auth` + "`" + ` header is not specified, this parameter
indicates where to find registry authorization credentials.
`

	var registryAuthFromFlagName string
	if cmdPrefix == "" {
		registryAuthFromFlagName = "registryAuthFrom"
	} else {
		registryAuthFromFlagName = fmt.Sprintf("%v.registryAuthFrom", cmdPrefix)
	}

	var registryAuthFromFlagDefault string = "spec"

	_ = cmd.PersistentFlags().String(registryAuthFromFlagName, registryAuthFromFlagDefault, registryAuthFromDescription)

	return nil
}
func registerOperationServiceServiceUpdateRollbackParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	rollbackDescription := `Set to this parameter to ` + "`" + `previous` + "`" + ` to cause a server-side rollback
to the previous service spec. The supplied spec will be ignored in
this case.
`

	var rollbackFlagName string
	if cmdPrefix == "" {
		rollbackFlagName = "rollback"
	} else {
		rollbackFlagName = fmt.Sprintf("%v.rollback", cmdPrefix)
	}

	var rollbackFlagDefault string

	_ = cmd.PersistentFlags().String(rollbackFlagName, rollbackFlagDefault, rollbackDescription)

	return nil
}
func registerOperationServiceServiceUpdateVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	versionDescription := `Required. The version number of the service object being updated. This is required to avoid conflicting writes. This version number should be the value as currently set on the service *before* the update. You can find the current version by calling ` + "`" + `GET /services/{id}` + "`" + ``

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault int64

	_ = cmd.PersistentFlags().Int64(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

// register flags to command
func registerModelServiceUpdateBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// allOf ServiceUpdateParamsBodyAO0 is not supported by go-swwagger cli yet

	// allOf ServiceUpdateParamsBodyAO1 is not supported by go-swwagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceUpdateBodyFlags(depth int, m *service.ServiceUpdateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// allOf ServiceUpdateParamsBodyAO0 is not supported by go-swwagger cli yet

	// allOf ServiceUpdateParamsBodyAO1 is not supported by go-swwagger cli yet

	return nil, retAdded
}

// interface{} register and retrieve functions are not rendered by go-swagger cli