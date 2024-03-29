// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/network"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNetworkNetworkDisconnectCmd returns a cmd to handle operation networkDisconnect
func makeOperationNetworkNetworkDisconnectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "NetworkDisconnect",
		Short: ``,
		RunE:  runOperationNetworkNetworkDisconnect,
	}

	if err := registerOperationNetworkNetworkDisconnectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNetworkNetworkDisconnect uses cmd flags to call endpoint api
func runOperationNetworkNetworkDisconnect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := network.NewNetworkDisconnectParams()
	if err, _ := retrieveOperationNetworkNetworkDisconnectContainerFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationNetworkNetworkDisconnectIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationNetworkNetworkDisconnectResult(appCli.Network.NetworkDisconnect(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationNetworkNetworkDisconnectParamFlags registers all flags needed to fill params
func registerOperationNetworkNetworkDisconnectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNetworkNetworkDisconnectContainerParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationNetworkNetworkDisconnectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNetworkNetworkDisconnectContainerParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var containerFlagName string
	if cmdPrefix == "" {
		containerFlagName = "container"
	} else {
		containerFlagName = fmt.Sprintf("%v.container", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(containerFlagName, "", "Optional json string for [container]. ")

	// add flags for body
	if err := registerModelNetworkDisconnectBodyFlags(0, "networkDisconnectBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationNetworkNetworkDisconnectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Network ID or name`

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

func retrieveOperationNetworkNetworkDisconnectContainerFlag(m *network.NetworkDisconnectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("container") {
		// Read container string from cmd and unmarshal
		containerValueStr, err := cmd.Flags().GetString("container")
		if err != nil {
			return err, false
		}

		containerValue := network.NetworkDisconnectBody{}
		if err := json.Unmarshal([]byte(containerValueStr), &containerValue); err != nil {
			return fmt.Errorf("cannot unmarshal container string in NetworkDisconnectBody: %v", err), false
		}
		m.Container = containerValue
	}
	containerValueModel := m.Container
	if swag.IsZero(containerValueModel) {
		containerValueModel = network.NetworkDisconnectBody{}
	}
	err, added := retrieveModelNetworkDisconnectBodyFlags(0, &containerValueModel, "networkDisconnectBody", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Container = containerValueModel
	}
	if dryRun && debug {

		containerValueDebugBytes, err := json.Marshal(m.Container)
		if err != nil {
			return err, false
		}
		logDebugf("Container dry-run payload: %v", string(containerValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationNetworkNetworkDisconnectIDFlag(m *network.NetworkDisconnectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationNetworkNetworkDisconnectResult parses request result and return the string content
func parseOperationNetworkNetworkDisconnectResult(resp0 *network.NetworkDisconnectOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning networkDisconnectOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*network.NetworkDisconnectForbidden)
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
		resp2, ok := iResp2.(*network.NetworkDisconnectNotFound)
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
		resp3, ok := iResp3.(*network.NetworkDisconnectInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response networkDisconnectOK is not supported by go-swagger cli yet.

	return "", nil
}

// register flags to command
func registerModelNetworkDisconnectBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNetworkDisconnectBodyContainer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkDisconnectBodyForce(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkDisconnectBodyContainer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	containerDescription := `The ID or name of the container to disconnect from the network.`

	var containerFlagName string
	if cmdPrefix == "" {
		containerFlagName = "Container"
	} else {
		containerFlagName = fmt.Sprintf("%v.Container", cmdPrefix)
	}

	var containerFlagDefault string

	_ = cmd.PersistentFlags().String(containerFlagName, containerFlagDefault, containerDescription)

	return nil
}

func registerNetworkDisconnectBodyForce(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	forceDescription := `Force the container to disconnect from the network.`

	var forceFlagName string
	if cmdPrefix == "" {
		forceFlagName = "Force"
	} else {
		forceFlagName = fmt.Sprintf("%v.Force", cmdPrefix)
	}

	var forceFlagDefault bool

	_ = cmd.PersistentFlags().Bool(forceFlagName, forceFlagDefault, forceDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNetworkDisconnectBodyFlags(depth int, m *network.NetworkDisconnectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, containerAdded := retrieveNetworkDisconnectBodyContainerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containerAdded

	err, forceAdded := retrieveNetworkDisconnectBodyForceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || forceAdded

	return nil, retAdded
}

func retrieveNetworkDisconnectBodyContainerFlags(depth int, m *network.NetworkDisconnectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	containerFlagName := fmt.Sprintf("%v.Container", cmdPrefix)
	if cmd.Flags().Changed(containerFlagName) {

		var containerFlagName string
		if cmdPrefix == "" {
			containerFlagName = "Container"
		} else {
			containerFlagName = fmt.Sprintf("%v.Container", cmdPrefix)
		}

		containerFlagValue, err := cmd.Flags().GetString(containerFlagName)
		if err != nil {
			return err, false
		}
		m.Container = containerFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNetworkDisconnectBodyForceFlags(depth int, m *network.NetworkDisconnectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	forceFlagName := fmt.Sprintf("%v.Force", cmdPrefix)
	if cmd.Flags().Changed(forceFlagName) {

		var forceFlagName string
		if cmdPrefix == "" {
			forceFlagName = "Force"
		} else {
			forceFlagName = fmt.Sprintf("%v.Force", cmdPrefix)
		}

		forceFlagValue, err := cmd.Flags().GetBool(forceFlagName)
		if err != nil {
			return err, false
		}
		m.Force = forceFlagValue

		retAdded = true
	}

	return nil, retAdded
}
