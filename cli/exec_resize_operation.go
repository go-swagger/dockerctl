// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/exec"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationExecExecResizeCmd returns a cmd to handle operation execResize
func makeOperationExecExecResizeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ExecResize",
		Short: `Resize the TTY session used by an exec instance. This endpoint only works if ` + "`" + `tty` + "`" + ` was specified as part of creating and starting the exec instance.`,
		RunE:  runOperationExecExecResize,
	}

	if err := registerOperationExecExecResizeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationExecExecResize uses cmd flags to call endpoint api
func runOperationExecExecResize(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := exec.NewExecResizeParams()
	if err, _ := retrieveOperationExecExecResizeHFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationExecExecResizeIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationExecExecResizeWFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationExecExecResizeResult(appCli.Exec.ExecResize(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationExecExecResizeParamFlags registers all flags needed to fill params
func registerOperationExecExecResizeParamFlags(cmd *cobra.Command) error {
	if err := registerOperationExecExecResizeHParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationExecExecResizeIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationExecExecResizeWParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationExecExecResizeHParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	hDescription := `Height of the TTY session in characters`

	var hFlagName string
	if cmdPrefix == "" {
		hFlagName = "h"
	} else {
		hFlagName = fmt.Sprintf("%v.h", cmdPrefix)
	}

	var hFlagDefault int64

	_ = cmd.PersistentFlags().Int64(hFlagName, hFlagDefault, hDescription)

	return nil
}
func registerOperationExecExecResizeIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Exec instance ID`

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
func registerOperationExecExecResizeWParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	wDescription := `Width of the TTY session in characters`

	var wFlagName string
	if cmdPrefix == "" {
		wFlagName = "w"
	} else {
		wFlagName = fmt.Sprintf("%v.w", cmdPrefix)
	}

	var wFlagDefault int64

	_ = cmd.PersistentFlags().Int64(wFlagName, wFlagDefault, wDescription)

	return nil
}

func retrieveOperationExecExecResizeHFlag(m *exec.ExecResizeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("h") {

		var hFlagName string
		if cmdPrefix == "" {
			hFlagName = "h"
		} else {
			hFlagName = fmt.Sprintf("%v.h", cmdPrefix)
		}

		hFlagValue, err := cmd.Flags().GetInt64(hFlagName)
		if err != nil {
			return err, false
		}
		m.H = &hFlagValue

	}
	return nil, retAdded
}
func retrieveOperationExecExecResizeIDFlag(m *exec.ExecResizeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationExecExecResizeWFlag(m *exec.ExecResizeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("w") {

		var wFlagName string
		if cmdPrefix == "" {
			wFlagName = "w"
		} else {
			wFlagName = fmt.Sprintf("%v.w", cmdPrefix)
		}

		wFlagValue, err := cmd.Flags().GetInt64(wFlagName)
		if err != nil {
			return err, false
		}
		m.W = &wFlagValue

	}
	return nil, retAdded
}

// parseOperationExecExecResizeResult parses request result and return the string content
func parseOperationExecExecResizeResult(resp0 *exec.ExecResizeCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning execResizeCreated is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*exec.ExecResizeNotFound)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response execResizeCreated is not supported by go-swagger cli yet.

	return "", nil
}
