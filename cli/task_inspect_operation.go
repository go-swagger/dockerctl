// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/task"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTaskTaskInspectCmd returns a cmd to handle operation taskInspect
func makeOperationTaskTaskInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "TaskInspect",
		Short: ``,
		RunE:  runOperationTaskTaskInspect,
	}

	if err := registerOperationTaskTaskInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskTaskInspect uses cmd flags to call endpoint api
func runOperationTaskTaskInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewTaskInspectParams()
	if err, _ := retrieveOperationTaskTaskInspectIDFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationTaskTaskInspectResult(appCli.Task.TaskInspect(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationTaskTaskInspectIDFlag(m *task.TaskInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationTaskTaskInspectResult prints output to stdout
func printOperationTaskTaskInspectResult(resp0 *task.TaskInspectOK, respErr error) error {
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

// registerOperationTaskTaskInspectParamFlags registers all flags needed to fill params
func registerOperationTaskTaskInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskTaskInspectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskTaskInspectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID of the task`

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
