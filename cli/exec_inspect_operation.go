// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/exec"
	"github.com/go-swagger/dockerctl/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationExecExecInspectCmd returns a cmd to handle operation execInspect
func makeOperationExecExecInspectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ExecInspect",
		Short: `Return low-level information about an exec instance.`,
		RunE:  runOperationExecExecInspect,
	}

	if err := registerOperationExecExecInspectParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationExecExecInspect uses cmd flags to call endpoint api
func runOperationExecExecInspect(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := exec.NewExecInspectParams()
	if err, _ := retrieveOperationExecExecInspectIDFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationExecExecInspectResult(appCli.Exec.ExecInspect(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationExecExecInspectIDFlag(m *exec.ExecInspectParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationExecExecInspectResult prints output to stdout
func printOperationExecExecInspectResult(resp0 *exec.ExecInspectOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*exec.ExecInspectOK)
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
		resp1, ok := iResp1.(*exec.ExecInspectNotFound)
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
		resp2, ok := iResp2.(*exec.ExecInspectInternalServerError)
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

// registerOperationExecExecInspectParamFlags registers all flags needed to fill params
func registerOperationExecExecInspectParamFlags(cmd *cobra.Command) error {
	if err := registerOperationExecExecInspectIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationExecExecInspectIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

// register flags to command
func registerModelExecInspectOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerExecInspectOKBodyCanRemove(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyContainerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyDetachKeys(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyExitCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyOpenStderr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyOpenStdin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyOpenStdout(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyPid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyProcessConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecInspectOKBodyRunning(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerExecInspectOKBodyCanRemove(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	canRemoveDescription := ``

	var canRemoveFlagName string
	if cmdPrefix == "" {
		canRemoveFlagName = "CanRemove"
	} else {
		canRemoveFlagName = fmt.Sprintf("%v.CanRemove", cmdPrefix)
	}

	var canRemoveFlagDefault bool

	_ = cmd.PersistentFlags().Bool(canRemoveFlagName, canRemoveFlagDefault, canRemoveDescription)

	return nil
}

func registerExecInspectOKBodyContainerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	containerIdDescription := ``

	var containerIdFlagName string
	if cmdPrefix == "" {
		containerIdFlagName = "ContainerID"
	} else {
		containerIdFlagName = fmt.Sprintf("%v.ContainerID", cmdPrefix)
	}

	var containerIdFlagDefault string

	_ = cmd.PersistentFlags().String(containerIdFlagName, containerIdFlagDefault, containerIdDescription)

	return nil
}

func registerExecInspectOKBodyDetachKeys(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	detachKeysDescription := ``

	var detachKeysFlagName string
	if cmdPrefix == "" {
		detachKeysFlagName = "DetachKeys"
	} else {
		detachKeysFlagName = fmt.Sprintf("%v.DetachKeys", cmdPrefix)
	}

	var detachKeysFlagDefault string

	_ = cmd.PersistentFlags().String(detachKeysFlagName, detachKeysFlagDefault, detachKeysDescription)

	return nil
}

func registerExecInspectOKBodyExitCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	exitCodeDescription := ``

	var exitCodeFlagName string
	if cmdPrefix == "" {
		exitCodeFlagName = "ExitCode"
	} else {
		exitCodeFlagName = fmt.Sprintf("%v.ExitCode", cmdPrefix)
	}

	var exitCodeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(exitCodeFlagName, exitCodeFlagDefault, exitCodeDescription)

	return nil
}

func registerExecInspectOKBodyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerExecInspectOKBodyOpenStderr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	openStderrDescription := ``

	var openStderrFlagName string
	if cmdPrefix == "" {
		openStderrFlagName = "OpenStderr"
	} else {
		openStderrFlagName = fmt.Sprintf("%v.OpenStderr", cmdPrefix)
	}

	var openStderrFlagDefault bool

	_ = cmd.PersistentFlags().Bool(openStderrFlagName, openStderrFlagDefault, openStderrDescription)

	return nil
}

func registerExecInspectOKBodyOpenStdin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	openStdinDescription := ``

	var openStdinFlagName string
	if cmdPrefix == "" {
		openStdinFlagName = "OpenStdin"
	} else {
		openStdinFlagName = fmt.Sprintf("%v.OpenStdin", cmdPrefix)
	}

	var openStdinFlagDefault bool

	_ = cmd.PersistentFlags().Bool(openStdinFlagName, openStdinFlagDefault, openStdinDescription)

	return nil
}

func registerExecInspectOKBodyOpenStdout(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	openStdoutDescription := ``

	var openStdoutFlagName string
	if cmdPrefix == "" {
		openStdoutFlagName = "OpenStdout"
	} else {
		openStdoutFlagName = fmt.Sprintf("%v.OpenStdout", cmdPrefix)
	}

	var openStdoutFlagDefault bool

	_ = cmd.PersistentFlags().Bool(openStdoutFlagName, openStdoutFlagDefault, openStdoutDescription)

	return nil
}

func registerExecInspectOKBodyPid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pidDescription := `The system process ID for the exec process.`

	var pidFlagName string
	if cmdPrefix == "" {
		pidFlagName = "Pid"
	} else {
		pidFlagName = fmt.Sprintf("%v.Pid", cmdPrefix)
	}

	var pidFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pidFlagName, pidFlagDefault, pidDescription)

	return nil
}

func registerExecInspectOKBodyProcessConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var processConfigFlagName string
	if cmdPrefix == "" {
		processConfigFlagName = "ProcessConfig"
	} else {
		processConfigFlagName = fmt.Sprintf("%v.ProcessConfig", cmdPrefix)
	}

	if err := registerModelProcessConfigFlags(depth+1, processConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerExecInspectOKBodyRunning(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	runningDescription := ``

	var runningFlagName string
	if cmdPrefix == "" {
		runningFlagName = "Running"
	} else {
		runningFlagName = fmt.Sprintf("%v.Running", cmdPrefix)
	}

	var runningFlagDefault bool

	_ = cmd.PersistentFlags().Bool(runningFlagName, runningFlagDefault, runningDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelExecInspectOKBodyFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, canRemoveAdded := retrieveExecInspectOKBodyCanRemoveFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || canRemoveAdded

	err, containerIdAdded := retrieveExecInspectOKBodyContainerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containerIdAdded

	err, detachKeysAdded := retrieveExecInspectOKBodyDetachKeysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || detachKeysAdded

	err, exitCodeAdded := retrieveExecInspectOKBodyExitCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || exitCodeAdded

	err, idAdded := retrieveExecInspectOKBodyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, openStderrAdded := retrieveExecInspectOKBodyOpenStderrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || openStderrAdded

	err, openStdinAdded := retrieveExecInspectOKBodyOpenStdinFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || openStdinAdded

	err, openStdoutAdded := retrieveExecInspectOKBodyOpenStdoutFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || openStdoutAdded

	err, pidAdded := retrieveExecInspectOKBodyPidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pidAdded

	err, processConfigAdded := retrieveExecInspectOKBodyProcessConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || processConfigAdded

	err, runningAdded := retrieveExecInspectOKBodyRunningFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || runningAdded

	return nil, retAdded
}

func retrieveExecInspectOKBodyCanRemoveFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	canRemoveFlagName := fmt.Sprintf("%v.CanRemove", cmdPrefix)
	if cmd.Flags().Changed(canRemoveFlagName) {

		var canRemoveFlagName string
		if cmdPrefix == "" {
			canRemoveFlagName = "CanRemove"
		} else {
			canRemoveFlagName = fmt.Sprintf("%v.CanRemove", cmdPrefix)
		}

		canRemoveFlagValue, err := cmd.Flags().GetBool(canRemoveFlagName)
		if err != nil {
			return err, false
		}
		m.CanRemove = canRemoveFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyContainerIDFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	containerIdFlagName := fmt.Sprintf("%v.ContainerID", cmdPrefix)
	if cmd.Flags().Changed(containerIdFlagName) {

		var containerIdFlagName string
		if cmdPrefix == "" {
			containerIdFlagName = "ContainerID"
		} else {
			containerIdFlagName = fmt.Sprintf("%v.ContainerID", cmdPrefix)
		}

		containerIdFlagValue, err := cmd.Flags().GetString(containerIdFlagName)
		if err != nil {
			return err, false
		}
		m.ContainerID = containerIdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyDetachKeysFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	detachKeysFlagName := fmt.Sprintf("%v.DetachKeys", cmdPrefix)
	if cmd.Flags().Changed(detachKeysFlagName) {

		var detachKeysFlagName string
		if cmdPrefix == "" {
			detachKeysFlagName = "DetachKeys"
		} else {
			detachKeysFlagName = fmt.Sprintf("%v.DetachKeys", cmdPrefix)
		}

		detachKeysFlagValue, err := cmd.Flags().GetString(detachKeysFlagName)
		if err != nil {
			return err, false
		}
		m.DetachKeys = detachKeysFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyExitCodeFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	exitCodeFlagName := fmt.Sprintf("%v.ExitCode", cmdPrefix)
	if cmd.Flags().Changed(exitCodeFlagName) {

		var exitCodeFlagName string
		if cmdPrefix == "" {
			exitCodeFlagName = "ExitCode"
		} else {
			exitCodeFlagName = fmt.Sprintf("%v.ExitCode", cmdPrefix)
		}

		exitCodeFlagValue, err := cmd.Flags().GetInt64(exitCodeFlagName)
		if err != nil {
			return err, false
		}
		m.ExitCode = exitCodeFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyIDFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveExecInspectOKBodyOpenStderrFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	openStderrFlagName := fmt.Sprintf("%v.OpenStderr", cmdPrefix)
	if cmd.Flags().Changed(openStderrFlagName) {

		var openStderrFlagName string
		if cmdPrefix == "" {
			openStderrFlagName = "OpenStderr"
		} else {
			openStderrFlagName = fmt.Sprintf("%v.OpenStderr", cmdPrefix)
		}

		openStderrFlagValue, err := cmd.Flags().GetBool(openStderrFlagName)
		if err != nil {
			return err, false
		}
		m.OpenStderr = openStderrFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyOpenStdinFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	openStdinFlagName := fmt.Sprintf("%v.OpenStdin", cmdPrefix)
	if cmd.Flags().Changed(openStdinFlagName) {

		var openStdinFlagName string
		if cmdPrefix == "" {
			openStdinFlagName = "OpenStdin"
		} else {
			openStdinFlagName = fmt.Sprintf("%v.OpenStdin", cmdPrefix)
		}

		openStdinFlagValue, err := cmd.Flags().GetBool(openStdinFlagName)
		if err != nil {
			return err, false
		}
		m.OpenStdin = openStdinFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyOpenStdoutFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	openStdoutFlagName := fmt.Sprintf("%v.OpenStdout", cmdPrefix)
	if cmd.Flags().Changed(openStdoutFlagName) {

		var openStdoutFlagName string
		if cmdPrefix == "" {
			openStdoutFlagName = "OpenStdout"
		} else {
			openStdoutFlagName = fmt.Sprintf("%v.OpenStdout", cmdPrefix)
		}

		openStdoutFlagValue, err := cmd.Flags().GetBool(openStdoutFlagName)
		if err != nil {
			return err, false
		}
		m.OpenStdout = openStdoutFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyPidFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	pidFlagName := fmt.Sprintf("%v.Pid", cmdPrefix)
	if cmd.Flags().Changed(pidFlagName) {

		var pidFlagName string
		if cmdPrefix == "" {
			pidFlagName = "Pid"
		} else {
			pidFlagName = fmt.Sprintf("%v.Pid", cmdPrefix)
		}

		pidFlagValue, err := cmd.Flags().GetInt64(pidFlagName)
		if err != nil {
			return err, false
		}
		m.Pid = pidFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyProcessConfigFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	processConfigFlagName := fmt.Sprintf("%v.ProcessConfig", cmdPrefix)
	if cmd.Flags().Changed(processConfigFlagName) {

		processConfigFlagValue := &models.ProcessConfig{}
		err, added := retrieveModelProcessConfigFlags(depth+1, processConfigFlagValue, processConfigFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.ProcessConfig = processConfigFlagValue
		}
	}
	return nil, retAdded
}

func retrieveExecInspectOKBodyRunningFlags(depth int, m *exec.ExecInspectOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	runningFlagName := fmt.Sprintf("%v.Running", cmdPrefix)
	if cmd.Flags().Changed(runningFlagName) {

		var runningFlagName string
		if cmdPrefix == "" {
			runningFlagName = "Running"
		} else {
			runningFlagName = fmt.Sprintf("%v.Running", cmdPrefix)
		}

		runningFlagValue, err := cmd.Flags().GetBool(runningFlagName)
		if err != nil {
			return err, false
		}
		m.Running = runningFlagValue

		retAdded = true
	}
	return nil, retAdded
}
