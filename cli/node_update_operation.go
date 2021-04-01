// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/node"
	"github.com/go-openapi/dockerctl/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNodeNodeUpdateCmd returns a cmd to handle operation nodeUpdate
func makeOperationNodeNodeUpdateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "NodeUpdate",
		Short: ``,
		RunE:  runOperationNodeNodeUpdate,
	}

	if err := registerOperationNodeNodeUpdateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNodeNodeUpdate uses cmd flags to call endpoint api
func runOperationNodeNodeUpdate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := node.NewNodeUpdateParams()
	if err, _ := retrieveOperationNodeNodeUpdateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationNodeNodeUpdateIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationNodeNodeUpdateVersionFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationNodeNodeUpdateResult(appCli.Node.NodeUpdate(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationNodeNodeUpdateBodyFlag(m *node.NodeUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.NodeSpec{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.NodeSpec: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.NodeSpec{}
	}
	err, added := retrieveModelNodeSpecFlags(0, bodyValueModel, "nodeSpec", cmd)
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
func retrieveOperationNodeNodeUpdateIDFlag(m *node.NodeUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationNodeNodeUpdateVersionFlag(m *node.NodeUpdateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationNodeNodeUpdateResult prints output to stdout
func printOperationNodeNodeUpdateResult(resp0 *node.NodeUpdateOK, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response nodeUpdateOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationNodeNodeUpdateParamFlags registers all flags needed to fill params
func registerOperationNodeNodeUpdateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNodeNodeUpdateBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationNodeNodeUpdateIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationNodeNodeUpdateVersionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNodeNodeUpdateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	exampleBodyStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(bodyFlagName, "", fmt.Sprintf("Optional json string for [body] of form %v.", string(exampleBodyStr)))

	// add flags for body
	if err := registerModelNodeSpecFlags(0, "nodeSpec", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationNodeNodeUpdateIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. The ID of the node`

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
func registerOperationNodeNodeUpdateVersionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	versionDescription := `Required. The version number of the node object being updated. This is required to avoid conflicting writes.`

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
