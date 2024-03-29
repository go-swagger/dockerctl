// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/config"
	"github.com/go-swagger/dockerctl/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationConfigConfigUpdateCmd returns a cmd to handle operation configUpdate
func makeOperationConfigConfigUpdateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ConfigUpdate",
		Short: ``,
		RunE:  runOperationConfigConfigUpdate,
	}

	if err := registerOperationConfigConfigUpdateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationConfigConfigUpdate uses cmd flags to call endpoint api
func runOperationConfigConfigUpdate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := config.NewConfigUpdateParams()
	if err, _ := retrieveOperationConfigConfigUpdateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationConfigConfigUpdateIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationConfigConfigUpdateVersionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationConfigConfigUpdateResult(appCli.Config.ConfigUpdate(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationConfigConfigUpdateParamFlags registers all flags needed to fill params
func registerOperationConfigConfigUpdateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationConfigConfigUpdateBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationConfigConfigUpdateIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationConfigConfigUpdateVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationConfigConfigUpdateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. The spec of the config to update. Currently, only the Labels field can be updated. All other fields must remain unchanged from the [ConfigInspect endpoint](#operation/ConfigInspect) response values.")

	// add flags for body
	if err := registerModelConfigSpecFlags(0, "configSpec", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationConfigConfigUpdateIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. The ID or name of the config`

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
func registerOperationConfigConfigUpdateVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	versionDescription := `Required. The version number of the config object being updated. This is required to avoid conflicting writes.`

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

func retrieveOperationConfigConfigUpdateBodyFlag(m *config.ConfigUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.ConfigSpec{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.ConfigSpec: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.ConfigSpec{}
	}
	err, added := retrieveModelConfigSpecFlags(0, bodyValueModel, "configSpec", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationConfigConfigUpdateIDFlag(m *config.ConfigUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationConfigConfigUpdateVersionFlag(m *config.ConfigUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationConfigConfigUpdateResult parses request result and return the string content
func parseOperationConfigConfigUpdateResult(resp0 *config.ConfigUpdateOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning configUpdateOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*config.ConfigUpdateBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*config.ConfigUpdateNotFound)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*config.ConfigUpdateInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*config.ConfigUpdateServiceUnavailable)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response configUpdateOK is not supported by go-swagger cli yet.

	return "", nil
}
