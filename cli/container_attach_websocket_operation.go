// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/container"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationContainerContainerAttachWebsocketCmd returns a cmd to handle operation containerAttachWebsocket
func makeOperationContainerContainerAttachWebsocketCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerAttachWebsocket",
		Short: ``,
		RunE:  runOperationContainerContainerAttachWebsocket,
	}

	if err := registerOperationContainerContainerAttachWebsocketParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerAttachWebsocket uses cmd flags to call endpoint api
func runOperationContainerContainerAttachWebsocket(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerAttachWebsocketParams()
	if err, _ := retrieveOperationContainerContainerAttachWebsocketDetachKeysFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachWebsocketIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachWebsocketLogsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachWebsocketStderrFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachWebsocketStdinFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachWebsocketStdoutFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachWebsocketStreamFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerAttachWebsocketResult(appCli.Container.ContainerAttachWebsocket(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerAttachWebsocketDetachKeysFlag(m *container.ContainerAttachWebsocketParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("detachKeys") {

		var detachKeysFlagName string
		if cmdPrefix == "" {
			detachKeysFlagName = "detachKeys"
		} else {
			detachKeysFlagName = fmt.Sprintf("%v.detachKeys", cmdPrefix)
		}

		detachKeysFlagValue, err := cmd.Flags().GetString(detachKeysFlagName)
		if err != nil {
			return err, false
		}
		m.DetachKeys = &detachKeysFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerContainerAttachWebsocketIDFlag(m *container.ContainerAttachWebsocketParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerAttachWebsocketLogsFlag(m *container.ContainerAttachWebsocketParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("logs") {

		var logsFlagName string
		if cmdPrefix == "" {
			logsFlagName = "logs"
		} else {
			logsFlagName = fmt.Sprintf("%v.logs", cmdPrefix)
		}

		logsFlagValue, err := cmd.Flags().GetBool(logsFlagName)
		if err != nil {
			return err, false
		}
		m.Logs = &logsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerContainerAttachWebsocketStderrFlag(m *container.ContainerAttachWebsocketParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("stderr") {

		var stderrFlagName string
		if cmdPrefix == "" {
			stderrFlagName = "stderr"
		} else {
			stderrFlagName = fmt.Sprintf("%v.stderr", cmdPrefix)
		}

		stderrFlagValue, err := cmd.Flags().GetBool(stderrFlagName)
		if err != nil {
			return err, false
		}
		m.Stderr = &stderrFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerContainerAttachWebsocketStdinFlag(m *container.ContainerAttachWebsocketParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("stdin") {

		var stdinFlagName string
		if cmdPrefix == "" {
			stdinFlagName = "stdin"
		} else {
			stdinFlagName = fmt.Sprintf("%v.stdin", cmdPrefix)
		}

		stdinFlagValue, err := cmd.Flags().GetBool(stdinFlagName)
		if err != nil {
			return err, false
		}
		m.Stdin = &stdinFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerContainerAttachWebsocketStdoutFlag(m *container.ContainerAttachWebsocketParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("stdout") {

		var stdoutFlagName string
		if cmdPrefix == "" {
			stdoutFlagName = "stdout"
		} else {
			stdoutFlagName = fmt.Sprintf("%v.stdout", cmdPrefix)
		}

		stdoutFlagValue, err := cmd.Flags().GetBool(stdoutFlagName)
		if err != nil {
			return err, false
		}
		m.Stdout = &stdoutFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerContainerAttachWebsocketStreamFlag(m *container.ContainerAttachWebsocketParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("stream") {

		var streamFlagName string
		if cmdPrefix == "" {
			streamFlagName = "stream"
		} else {
			streamFlagName = fmt.Sprintf("%v.stream", cmdPrefix)
		}

		streamFlagValue, err := cmd.Flags().GetBool(streamFlagName)
		if err != nil {
			return err, false
		}
		m.Stream = &streamFlagValue

	}
	return nil, retAdded
}

// printOperationContainerContainerAttachWebsocketResult prints output to stdout
func printOperationContainerContainerAttachWebsocketResult(resp0 *container.ContainerAttachWebsocketOK, respErr error) error {
	if respErr != nil {

		// Non schema case: warning containerAttachWebsocketSwitchingProtocols is not supported

		// Non schema case: warning containerAttachWebsocketOK is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*container.ContainerAttachWebsocketBadRequest)
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
		resp3, ok := iResp3.(*container.ContainerAttachWebsocketNotFound)
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

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*container.ContainerAttachWebsocketInternalServerError)
		if ok {
			if !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		return respErr
	}

	// warning: non schema response containerAttachWebsocketOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationContainerContainerAttachWebsocketParamFlags registers all flags needed to fill params
func registerOperationContainerContainerAttachWebsocketParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerAttachWebsocketDetachKeysParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachWebsocketIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachWebsocketLogsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachWebsocketStderrParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachWebsocketStdinParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachWebsocketStdoutParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachWebsocketStreamParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerAttachWebsocketDetachKeysParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	detachKeysDescription := `Override the key sequence for detaching a container.Format is a single character ` + "`" + `[a-Z]` + "`" + ` or ` + "`" + `ctrl-<value>` + "`" + ` where ` + "`" + `<value>` + "`" + ` is one of: ` + "`" + `a-z` + "`" + `, ` + "`" + `@` + "`" + `, ` + "`" + `^` + "`" + `, ` + "`" + `[` + "`" + `, ` + "`" + `,` + "`" + `, or ` + "`" + `_` + "`" + `.`

	var detachKeysFlagName string
	if cmdPrefix == "" {
		detachKeysFlagName = "detachKeys"
	} else {
		detachKeysFlagName = fmt.Sprintf("%v.detachKeys", cmdPrefix)
	}

	var detachKeysFlagDefault string

	_ = cmd.PersistentFlags().String(detachKeysFlagName, detachKeysFlagDefault, detachKeysDescription)

	return nil
}
func registerOperationContainerContainerAttachWebsocketIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID or name of the container`

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
func registerOperationContainerContainerAttachWebsocketLogsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	logsDescription := `Return logs`

	var logsFlagName string
	if cmdPrefix == "" {
		logsFlagName = "logs"
	} else {
		logsFlagName = fmt.Sprintf("%v.logs", cmdPrefix)
	}

	var logsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(logsFlagName, logsFlagDefault, logsDescription)

	return nil
}
func registerOperationContainerContainerAttachWebsocketStderrParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	stderrDescription := `Attach to ` + "`" + `stderr` + "`" + ``

	var stderrFlagName string
	if cmdPrefix == "" {
		stderrFlagName = "stderr"
	} else {
		stderrFlagName = fmt.Sprintf("%v.stderr", cmdPrefix)
	}

	var stderrFlagDefault bool

	_ = cmd.PersistentFlags().Bool(stderrFlagName, stderrFlagDefault, stderrDescription)

	return nil
}
func registerOperationContainerContainerAttachWebsocketStdinParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	stdinDescription := `Attach to ` + "`" + `stdin` + "`" + ``

	var stdinFlagName string
	if cmdPrefix == "" {
		stdinFlagName = "stdin"
	} else {
		stdinFlagName = fmt.Sprintf("%v.stdin", cmdPrefix)
	}

	var stdinFlagDefault bool

	_ = cmd.PersistentFlags().Bool(stdinFlagName, stdinFlagDefault, stdinDescription)

	return nil
}
func registerOperationContainerContainerAttachWebsocketStdoutParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	stdoutDescription := `Attach to ` + "`" + `stdout` + "`" + ``

	var stdoutFlagName string
	if cmdPrefix == "" {
		stdoutFlagName = "stdout"
	} else {
		stdoutFlagName = fmt.Sprintf("%v.stdout", cmdPrefix)
	}

	var stdoutFlagDefault bool

	_ = cmd.PersistentFlags().Bool(stdoutFlagName, stdoutFlagDefault, stdoutDescription)

	return nil
}
func registerOperationContainerContainerAttachWebsocketStreamParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	streamDescription := `Return stream`

	var streamFlagName string
	if cmdPrefix == "" {
		streamFlagName = "stream"
	} else {
		streamFlagName = fmt.Sprintf("%v.stream", cmdPrefix)
	}

	var streamFlagDefault bool

	_ = cmd.PersistentFlags().Bool(streamFlagName, streamFlagDefault, streamDescription)

	return nil
}
