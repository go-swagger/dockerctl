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

// makeOperationNetworkNetworkInspectCmd returns a cmd to handle operation networkInspect
func makeOperationNetworkNetworkInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "NetworkInspect",
		Short: ``,
		RunE:  runOperationNetworkNetworkInspect,
	}

	if err := registerOperationNetworkNetworkInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNetworkNetworkInspect uses cmd flags to call endpoint api
func runOperationNetworkNetworkInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := network.NewNetworkInspectParams()
	if err, _ := retrieveOperationNetworkNetworkInspectIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationNetworkNetworkInspectScopeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationNetworkNetworkInspectVerboseFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationNetworkNetworkInspectResult(appCli.Network.NetworkInspect(params)); err != nil {
		return err
	}
	return nil
}

// registerOperationNetworkNetworkInspectParamFlags registers all flags needed to fill params
func registerOperationNetworkNetworkInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNetworkNetworkInspectIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationNetworkNetworkInspectScopeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationNetworkNetworkInspectVerboseParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNetworkNetworkInspectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationNetworkNetworkInspectScopeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	scopeDescription := `Filter the network by scope (swarm, global, or local)`

	var scopeFlagName string
	if cmdPrefix == "" {
		scopeFlagName = "scope"
	} else {
		scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
	}

	var scopeFlagDefault string

	_ = cmd.PersistentFlags().String(scopeFlagName, scopeFlagDefault, scopeDescription)

	return nil
}
func registerOperationNetworkNetworkInspectVerboseParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	verboseDescription := `Detailed inspect output for troubleshooting`

	var verboseFlagName string
	if cmdPrefix == "" {
		verboseFlagName = "verbose"
	} else {
		verboseFlagName = fmt.Sprintf("%v.verbose", cmdPrefix)
	}

	var verboseFlagDefault bool

	_ = cmd.PersistentFlags().Bool(verboseFlagName, verboseFlagDefault, verboseDescription)

	return nil
}

func retrieveOperationNetworkNetworkInspectIDFlag(m *network.NetworkInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationNetworkNetworkInspectScopeFlag(m *network.NetworkInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("scope") {

		var scopeFlagName string
		if cmdPrefix == "" {
			scopeFlagName = "scope"
		} else {
			scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
		}

		scopeFlagValue, err := cmd.Flags().GetString(scopeFlagName)
		if err != nil {
			return err, false
		}
		m.Scope = &scopeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationNetworkNetworkInspectVerboseFlag(m *network.NetworkInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("verbose") {

		var verboseFlagName string
		if cmdPrefix == "" {
			verboseFlagName = "verbose"
		} else {
			verboseFlagName = fmt.Sprintf("%v.verbose", cmdPrefix)
		}

		verboseFlagValue, err := cmd.Flags().GetBool(verboseFlagName)
		if err != nil {
			return err, false
		}
		m.Verbose = &verboseFlagValue

	}
	return nil, retAdded
}

// printOperationNetworkNetworkInspectResult prints output to stdout
func printOperationNetworkNetworkInspectResult(resp0 *network.NetworkInspectOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*network.NetworkInspectOK)
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
		resp1, ok := iResp1.(*network.NetworkInspectNotFound)
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
		resp2, ok := iResp2.(*network.NetworkInspectInternalServerError)
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
