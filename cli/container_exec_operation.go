// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/exec"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationExecContainerExecCmd returns a cmd to handle operation containerExec
func makeOperationExecContainerExecCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerExec",
		Short: `Run a command inside a running container.`,
		RunE:  runOperationExecContainerExec,
	}

	if err := registerOperationExecContainerExecParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationExecContainerExec uses cmd flags to call endpoint api
func runOperationExecContainerExec(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := exec.NewContainerExecParams()
	if err, _ := retrieveOperationExecContainerExecExecConfigFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationExecContainerExecIDFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationExecContainerExecResult(appCli.Exec.ContainerExec(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationExecContainerExecExecConfigFlag(m *exec.ContainerExecParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("execConfig") {
		// Read execConfig string from cmd and unmarshal
		execConfigValueStr, err := cmd.Flags().GetString("execConfig")
		if err != nil {
			return err, false
		}

		execConfigValue := exec.ContainerExecBody{}
		if err := json.Unmarshal([]byte(execConfigValueStr), &execConfigValue); err != nil {
			return fmt.Errorf("cannot unmarshal execConfig string in ContainerExecBody: %v", err), false
		}
		m.ExecConfig = execConfigValue
	}
	execConfigValueModel := m.ExecConfig
	if swag.IsZero(execConfigValueModel) {
		execConfigValueModel = exec.ContainerExecBody{}
	}
	err, added := retrieveModelContainerExecBodyFlags(0, &execConfigValueModel, "containerExecBody", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.ExecConfig = execConfigValueModel
	}
	execConfigValueDebugBytes, err := json.Marshal(m.ExecConfig)
	if err != nil {
		return err, false
	}
	logDebugf("ExecConfig payload: %v", string(execConfigValueDebugBytes))
	return nil, retAdded
}
func retrieveOperationExecContainerExecIDFlag(m *exec.ContainerExecParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationExecContainerExecResult prints output to stdout
func printOperationExecContainerExecResult(resp0 *exec.ContainerExecCreated, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*exec.ContainerExecCreated)
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
		resp1, ok := iResp1.(*exec.ContainerExecNotFound)
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
		resp2, ok := iResp2.(*exec.ContainerExecConflict)
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
		resp3, ok := iResp3.(*exec.ContainerExecInternalServerError)
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

	if !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return err
		}
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationExecContainerExecParamFlags registers all flags needed to fill params
func registerOperationExecContainerExecParamFlags(cmd *cobra.Command) error {
	if err := registerOperationExecContainerExecExecConfigParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationExecContainerExecIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationExecContainerExecExecConfigParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var execConfigFlagName string
	if cmdPrefix == "" {
		execConfigFlagName = "execConfig"
	} else {
		execConfigFlagName = fmt.Sprintf("%v.execConfig", cmdPrefix)
	}

	exampleExecConfigStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(execConfigFlagName, "", fmt.Sprintf("Optional json string for [execConfig] of form %v.Exec configuration", string(exampleExecConfigStr)))

	// add flags for body
	if err := registerModelContainerExecBodyFlags(0, "containerExecBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationExecContainerExecIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID or name of container`

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
func registerModelContainerExecBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerExecBodyAttachStderr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyAttachStdin(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyAttachStdout(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyCmd(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyDetachKeys(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyEnv(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyPrivileged(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyTty(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyUser(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerExecBodyWorkingDir(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerExecBodyAttachStderr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	attachStderrDescription := `Attach to ` + "`" + `stderr` + "`" + ` of the exec command.`

	var attachStderrFlagName string
	if cmdPrefix == "" {
		attachStderrFlagName = "AttachStderr"
	} else {
		attachStderrFlagName = fmt.Sprintf("%v.AttachStderr", cmdPrefix)
	}

	var attachStderrFlagDefault bool

	_ = cmd.PersistentFlags().Bool(attachStderrFlagName, attachStderrFlagDefault, attachStderrDescription)

	return nil
}

func registerContainerExecBodyAttachStdin(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	attachStdinDescription := `Attach to ` + "`" + `stdin` + "`" + ` of the exec command.`

	var attachStdinFlagName string
	if cmdPrefix == "" {
		attachStdinFlagName = "AttachStdin"
	} else {
		attachStdinFlagName = fmt.Sprintf("%v.AttachStdin", cmdPrefix)
	}

	var attachStdinFlagDefault bool

	_ = cmd.PersistentFlags().Bool(attachStdinFlagName, attachStdinFlagDefault, attachStdinDescription)

	return nil
}

func registerContainerExecBodyAttachStdout(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	attachStdoutDescription := `Attach to ` + "`" + `stdout` + "`" + ` of the exec command.`

	var attachStdoutFlagName string
	if cmdPrefix == "" {
		attachStdoutFlagName = "AttachStdout"
	} else {
		attachStdoutFlagName = fmt.Sprintf("%v.AttachStdout", cmdPrefix)
	}

	var attachStdoutFlagDefault bool

	_ = cmd.PersistentFlags().Bool(attachStdoutFlagName, attachStdoutFlagDefault, attachStdoutDescription)

	return nil
}

func registerContainerExecBodyCmd(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Cmd []string array type is not supported by go-swagger cli yet

	return nil
}

func registerContainerExecBodyDetachKeys(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	detachKeysDescription := `Override the key sequence for detaching a container. Format is a single character ` + "`" + `[a-Z]` + "`" + ` or ` + "`" + `ctrl-<value>` + "`" + ` where ` + "`" + `<value>` + "`" + ` is one of: ` + "`" + `a-z` + "`" + `, ` + "`" + `@` + "`" + `, ` + "`" + `^` + "`" + `, ` + "`" + `[` + "`" + `, ` + "`" + `,` + "`" + ` or ` + "`" + `_` + "`" + `.`

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

func registerContainerExecBodyEnv(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Env []string array type is not supported by go-swagger cli yet

	return nil
}

func registerContainerExecBodyPrivileged(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	privilegedDescription := `Runs the exec process with extended privileges.`

	var privilegedFlagName string
	if cmdPrefix == "" {
		privilegedFlagName = "Privileged"
	} else {
		privilegedFlagName = fmt.Sprintf("%v.Privileged", cmdPrefix)
	}

	var privilegedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(privilegedFlagName, privilegedFlagDefault, privilegedDescription)

	return nil
}

func registerContainerExecBodyTty(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ttyDescription := `Allocate a pseudo-TTY.`

	var ttyFlagName string
	if cmdPrefix == "" {
		ttyFlagName = "Tty"
	} else {
		ttyFlagName = fmt.Sprintf("%v.Tty", cmdPrefix)
	}

	var ttyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(ttyFlagName, ttyFlagDefault, ttyDescription)

	return nil
}

func registerContainerExecBodyUser(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	userDescription := `The user, and optionally, group to run the exec process inside the container. Format is one of: ` + "`" + `user` + "`" + `, ` + "`" + `user:group` + "`" + `, ` + "`" + `uid` + "`" + `, or ` + "`" + `uid:gid` + "`" + `.`

	var userFlagName string
	if cmdPrefix == "" {
		userFlagName = "User"
	} else {
		userFlagName = fmt.Sprintf("%v.User", cmdPrefix)
	}

	var userFlagDefault string

	_ = cmd.PersistentFlags().String(userFlagName, userFlagDefault, userDescription)

	return nil
}

func registerContainerExecBodyWorkingDir(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	workingDirDescription := `The working directory for the exec process inside the container.`

	var workingDirFlagName string
	if cmdPrefix == "" {
		workingDirFlagName = "WorkingDir"
	} else {
		workingDirFlagName = fmt.Sprintf("%v.WorkingDir", cmdPrefix)
	}

	var workingDirFlagDefault string

	_ = cmd.PersistentFlags().String(workingDirFlagName, workingDirFlagDefault, workingDirDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerExecBodyFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attachStderrAdded := retrieveContainerExecBodyAttachStderrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attachStderrAdded

	err, attachStdinAdded := retrieveContainerExecBodyAttachStdinFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attachStdinAdded

	err, attachStdoutAdded := retrieveContainerExecBodyAttachStdoutFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attachStdoutAdded

	err, cmdAdded := retrieveContainerExecBodyCmdFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cmdAdded

	err, detachKeysAdded := retrieveContainerExecBodyDetachKeysFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || detachKeysAdded

	err, envAdded := retrieveContainerExecBodyEnvFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || envAdded

	err, privilegedAdded := retrieveContainerExecBodyPrivilegedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || privilegedAdded

	err, ttyAdded := retrieveContainerExecBodyTtyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ttyAdded

	err, userAdded := retrieveContainerExecBodyUserFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userAdded

	err, workingDirAdded := retrieveContainerExecBodyWorkingDirFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || workingDirAdded

	return nil, retAdded
}

func retrieveContainerExecBodyAttachStderrFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	attachStderrFlagName := fmt.Sprintf("%v.AttachStderr", cmdPrefix)
	if cmd.Flags().Changed(attachStderrFlagName) {

		var attachStderrFlagName string
		if cmdPrefix == "" {
			attachStderrFlagName = "AttachStderr"
		} else {
			attachStderrFlagName = fmt.Sprintf("%v.AttachStderr", cmdPrefix)
		}

		attachStderrFlagValue, err := cmd.Flags().GetBool(attachStderrFlagName)
		if err != nil {
			return err, false
		}
		m.AttachStderr = attachStderrFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveContainerExecBodyAttachStdinFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	attachStdinFlagName := fmt.Sprintf("%v.AttachStdin", cmdPrefix)
	if cmd.Flags().Changed(attachStdinFlagName) {

		var attachStdinFlagName string
		if cmdPrefix == "" {
			attachStdinFlagName = "AttachStdin"
		} else {
			attachStdinFlagName = fmt.Sprintf("%v.AttachStdin", cmdPrefix)
		}

		attachStdinFlagValue, err := cmd.Flags().GetBool(attachStdinFlagName)
		if err != nil {
			return err, false
		}
		m.AttachStdin = attachStdinFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveContainerExecBodyAttachStdoutFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	attachStdoutFlagName := fmt.Sprintf("%v.AttachStdout", cmdPrefix)
	if cmd.Flags().Changed(attachStdoutFlagName) {

		var attachStdoutFlagName string
		if cmdPrefix == "" {
			attachStdoutFlagName = "AttachStdout"
		} else {
			attachStdoutFlagName = fmt.Sprintf("%v.AttachStdout", cmdPrefix)
		}

		attachStdoutFlagValue, err := cmd.Flags().GetBool(attachStdoutFlagName)
		if err != nil {
			return err, false
		}
		m.AttachStdout = attachStdoutFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveContainerExecBodyCmdFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	cmdFlagName := fmt.Sprintf("%v.Cmd", cmdPrefix)
	if cmd.Flags().Changed(cmdFlagName) {
		// warning: Cmd array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveContainerExecBodyDetachKeysFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveContainerExecBodyEnvFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	envFlagName := fmt.Sprintf("%v.Env", cmdPrefix)
	if cmd.Flags().Changed(envFlagName) {
		// warning: Env array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveContainerExecBodyPrivilegedFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	privilegedFlagName := fmt.Sprintf("%v.Privileged", cmdPrefix)
	if cmd.Flags().Changed(privilegedFlagName) {

		var privilegedFlagName string
		if cmdPrefix == "" {
			privilegedFlagName = "Privileged"
		} else {
			privilegedFlagName = fmt.Sprintf("%v.Privileged", cmdPrefix)
		}

		privilegedFlagValue, err := cmd.Flags().GetBool(privilegedFlagName)
		if err != nil {
			return err, false
		}
		m.Privileged = &privilegedFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveContainerExecBodyTtyFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	ttyFlagName := fmt.Sprintf("%v.Tty", cmdPrefix)
	if cmd.Flags().Changed(ttyFlagName) {

		var ttyFlagName string
		if cmdPrefix == "" {
			ttyFlagName = "Tty"
		} else {
			ttyFlagName = fmt.Sprintf("%v.Tty", cmdPrefix)
		}

		ttyFlagValue, err := cmd.Flags().GetBool(ttyFlagName)
		if err != nil {
			return err, false
		}
		m.Tty = ttyFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveContainerExecBodyUserFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	userFlagName := fmt.Sprintf("%v.User", cmdPrefix)
	if cmd.Flags().Changed(userFlagName) {

		var userFlagName string
		if cmdPrefix == "" {
			userFlagName = "User"
		} else {
			userFlagName = fmt.Sprintf("%v.User", cmdPrefix)
		}

		userFlagValue, err := cmd.Flags().GetString(userFlagName)
		if err != nil {
			return err, false
		}
		m.User = userFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveContainerExecBodyWorkingDirFlags(depth int, m *exec.ContainerExecBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	workingDirFlagName := fmt.Sprintf("%v.WorkingDir", cmdPrefix)
	if cmd.Flags().Changed(workingDirFlagName) {

		var workingDirFlagName string
		if cmdPrefix == "" {
			workingDirFlagName = "WorkingDir"
		} else {
			workingDirFlagName = fmt.Sprintf("%v.WorkingDir", cmdPrefix)
		}

		workingDirFlagValue, err := cmd.Flags().GetString(workingDirFlagName)
		if err != nil {
			return err, false
		}
		m.WorkingDir = workingDirFlagValue

		retAdded = true
	}
	return nil, retAdded
}
