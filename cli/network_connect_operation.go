// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/network"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNetworkNetworkConnectCmd returns a cmd to handle operation networkConnect
func makeOperationNetworkNetworkConnectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "NetworkConnect",
		Short: ``,
		RunE:  runOperationNetworkNetworkConnect,
	}

	if err := registerOperationNetworkNetworkConnectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNetworkNetworkConnect uses cmd flags to call endpoint api
func runOperationNetworkNetworkConnect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := network.NewNetworkConnectParams()
	if err, _ := retrieveOperationNetworkNetworkConnectContainerFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationNetworkNetworkConnectIDFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationNetworkNetworkConnectResult(appCli.Network.NetworkConnect(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationNetworkNetworkConnectContainerFlag(m *network.NetworkConnectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("container") {
		// Read container string from cmd and unmarshal
		containerValueStr, err := cmd.Flags().GetString("container")
		if err != nil {
			return err, false
		}

		containerValue := network.NetworkConnectBody{}
		if err := json.Unmarshal([]byte(containerValueStr), &containerValue); err != nil {
			return fmt.Errorf("cannot unmarshal container string in NetworkConnectBody: %v", err), false
		}
		m.Container = containerValue
	}
	containerValueModel := m.Container
	if swag.IsZero(containerValueModel) {
		containerValueModel = network.NetworkConnectBody{}
	}
	err, added := retrieveModelNetworkConnectBodyFlags(0, &containerValueModel, "networkConnectBody", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Container = containerValueModel
	}
	containerValueDebugBytes, err := json.Marshal(m.Container)
	if err != nil {
		return err, false
	}
	logDebugf("Container payload: %v", string(containerValueDebugBytes))
	return nil, retAdded
}
func retrieveOperationNetworkNetworkConnectIDFlag(m *network.NetworkConnectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationNetworkNetworkConnectResult prints output to stdout
func printOperationNetworkNetworkConnectResult(resp0 *network.NetworkConnectOK, respErr error) error {
	if respErr != nil {

		// Non schema case: warning networkConnectOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*network.NetworkConnectForbidden)
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
		resp2, ok := iResp2.(*network.NetworkConnectNotFound)
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
		resp3, ok := iResp3.(*network.NetworkConnectInternalServerError)
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

	// warning: non schema response networkConnectOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationNetworkNetworkConnectParamFlags registers all flags needed to fill params
func registerOperationNetworkNetworkConnectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNetworkNetworkConnectContainerParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationNetworkNetworkConnectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNetworkNetworkConnectContainerParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var containerFlagName string
	if cmdPrefix == "" {
		containerFlagName = "container"
	} else {
		containerFlagName = fmt.Sprintf("%v.container", cmdPrefix)
	}

	exampleContainerStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(containerFlagName, "", fmt.Sprintf("Optional json string for [container] of form %v.", string(exampleContainerStr)))

	// add flags for body
	if err := registerModelNetworkConnectBodyFlags(0, "networkConnectBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationNetworkNetworkConnectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

// register flags to command
func registerModelNetworkConnectBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNetworkConnectBodyContainer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkConnectBodyEndpointConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkConnectBodyContainer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	containerDescription := `The ID or name of the container to connect to the network.`

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

func registerNetworkConnectBodyEndpointConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var endpointConfigFlagName string
	if cmdPrefix == "" {
		endpointConfigFlagName = "EndpointConfig"
	} else {
		endpointConfigFlagName = fmt.Sprintf("%v.EndpointConfig", cmdPrefix)
	}

	registerModelNetworkConnectBodyFlags(depth+1, endpointConfigFlagName, cmd)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNetworkConnectBodyFlags(depth int, m *network.NetworkConnectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, containerAdded := retrieveNetworkConnectBodyContainerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containerAdded

	err, endpointConfigAdded := retrieveNetworkConnectBodyEndpointConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointConfigAdded

	return nil, retAdded
}

func retrieveNetworkConnectBodyContainerFlags(depth int, m *network.NetworkConnectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNetworkConnectBodyEndpointConfigFlags(depth int, m *network.NetworkConnectBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	endpointConfigFlagName := fmt.Sprintf("%v.EndpointConfig", cmdPrefix)
	if cmd.Flags().Changed(endpointConfigFlagName) {

		endpointConfigFlagValue := &network.NetworkConnectBody{}
		err, added := retrieveModelNetworkConnectBodyFlags(depth+1, endpointConfigFlagValue, endpointConfigFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}
