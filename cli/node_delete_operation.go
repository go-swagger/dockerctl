// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/node"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNodeNodeDeleteCmd returns a cmd to handle operation nodeDelete
func makeOperationNodeNodeDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "NodeDelete",
		Short: ``,
		RunE:  runOperationNodeNodeDelete,
	}

	if err := registerOperationNodeNodeDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNodeNodeDelete uses cmd flags to call endpoint api
func runOperationNodeNodeDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := node.NewNodeDeleteParams()
	if err, _ := retrieveOperationNodeNodeDeleteForceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationNodeNodeDeleteIDFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationNodeNodeDeleteResult(appCli.Node.NodeDelete(params)); err != nil {
		return err
	}
	return nil
}

// registerOperationNodeNodeDeleteParamFlags registers all flags needed to fill params
func registerOperationNodeNodeDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNodeNodeDeleteForceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationNodeNodeDeleteIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNodeNodeDeleteForceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	forceDescription := `Force remove a node from the swarm`

	var forceFlagName string
	if cmdPrefix == "" {
		forceFlagName = "force"
	} else {
		forceFlagName = fmt.Sprintf("%v.force", cmdPrefix)
	}

	var forceFlagDefault bool

	_ = cmd.PersistentFlags().Bool(forceFlagName, forceFlagDefault, forceDescription)

	return nil
}
func registerOperationNodeNodeDeleteIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. The ID or name of the node`

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

func retrieveOperationNodeNodeDeleteForceFlag(m *node.NodeDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("force") {

		var forceFlagName string
		if cmdPrefix == "" {
			forceFlagName = "force"
		} else {
			forceFlagName = fmt.Sprintf("%v.force", cmdPrefix)
		}

		forceFlagValue, err := cmd.Flags().GetBool(forceFlagName)
		if err != nil {
			return err, false
		}
		m.Force = &forceFlagValue

	}
	return nil, retAdded
}
func retrieveOperationNodeNodeDeleteIDFlag(m *node.NodeDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationNodeNodeDeleteResult prints output to stdout
func printOperationNodeNodeDeleteResult(resp0 *node.NodeDeleteOK, respErr error) error {
	if respErr != nil {

		// Non schema case: warning nodeDeleteOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*node.NodeDeleteNotFound)
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
		resp2, ok := iResp2.(*node.NodeDeleteInternalServerError)
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
		resp3, ok := iResp3.(*node.NodeDeleteServiceUnavailable)
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

	// warning: non schema response nodeDeleteOK is not supported by go-swagger cli yet.

	return nil
}
