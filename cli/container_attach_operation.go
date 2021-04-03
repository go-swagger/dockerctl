// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/container"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationContainerContainerAttachCmd returns a cmd to handle operation containerAttach
func makeOperationContainerContainerAttachCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "ContainerAttach",
		Short: `Attach to a container to read its output or send it input. You can attach to the same container multiple times and you can reattach to containers that have been detached.

Either the ` + "`" + `stream` + "`" + ` or ` + "`" + `logs` + "`" + ` parameter must be ` + "`" + `true` + "`" + ` for this endpoint to do anything.

See [the documentation for the ` + "`" + `docker attach` + "`" + ` command](https://docs.docker.com/engine/reference/commandline/attach/) for more details.

### Hijacking

This endpoint hijacks the HTTP connection to transport ` + "`" + `stdin` + "`" + `, ` + "`" + `stdout` + "`" + `, and ` + "`" + `stderr` + "`" + ` on the same socket.

This is the response from the daemon for an attach request:

` + "`" + `` + "`" + `` + "`" + `
HTTP/1.1 200 OK
Content-Type: application/vnd.docker.raw-stream

[STREAM]
` + "`" + `` + "`" + `` + "`" + `

After the headers and two new lines, the TCP connection can now be used for raw, bidirectional communication between the client and server.

To hint potential proxies about connection hijacking, the Docker client can also optionally send connection upgrade headers.

For example, the client sends this request to upgrade the connection:

` + "`" + `` + "`" + `` + "`" + `
POST /containers/16253994b7c4/attach?stream=1&stdout=1 HTTP/1.1
Upgrade: tcp
Connection: Upgrade
` + "`" + `` + "`" + `` + "`" + `

The Docker daemon will respond with a ` + "`" + `101 UPGRADED` + "`" + ` response, and will similarly follow with the raw stream:

` + "`" + `` + "`" + `` + "`" + `
HTTP/1.1 101 UPGRADED
Content-Type: application/vnd.docker.raw-stream
Connection: Upgrade
Upgrade: tcp

[STREAM]
` + "`" + `` + "`" + `` + "`" + `

### Stream format

When the TTY setting is disabled in [` + "`" + `POST /containers/create` + "`" + `](#operation/ContainerCreate), the stream over the hijacked connected is multiplexed to separate out ` + "`" + `stdout` + "`" + ` and ` + "`" + `stderr` + "`" + `. The stream consists of a series of frames, each containing a header and a payload.

The header contains the information which the stream writes (` + "`" + `stdout` + "`" + ` or ` + "`" + `stderr` + "`" + `). It also contains the size of the associated frame encoded in the last four bytes (` + "`" + `uint32` + "`" + `).

It is encoded on the first eight bytes like this:

` + "`" + `` + "`" + `` + "`" + `go
header := [8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}
` + "`" + `` + "`" + `` + "`" + `

` + "`" + `STREAM_TYPE` + "`" + ` can be:

- 0: ` + "`" + `stdin` + "`" + ` (is written on ` + "`" + `stdout` + "`" + `)
- 1: ` + "`" + `stdout` + "`" + `
- 2: ` + "`" + `stderr` + "`" + `

` + "`" + `SIZE1, SIZE2, SIZE3, SIZE4` + "`" + ` are the four bytes of the ` + "`" + `uint32` + "`" + ` size encoded as big endian.

Following the header is the payload, which is the specified number of bytes of ` + "`" + `STREAM_TYPE` + "`" + `.

The simplest way to implement this protocol is the following:

1. Read 8 bytes.
2. Choose ` + "`" + `stdout` + "`" + ` or ` + "`" + `stderr` + "`" + ` depending on the first byte.
3. Extract the frame size from the last four bytes.
4. Read the extracted size and output it on the correct output.
5. Goto 1.

### Stream format when using a TTY

When the TTY setting is enabled in [` + "`" + `POST /containers/create` + "`" + `](#operation/ContainerCreate), the stream is not multiplexed. The data exchanged over the hijacked connection is simply the raw data from the process PTY and client's ` + "`" + `stdin` + "`" + `.
`,
		RunE: runOperationContainerContainerAttach,
	}

	if err := registerOperationContainerContainerAttachParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerAttach uses cmd flags to call endpoint api
func runOperationContainerContainerAttach(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerAttachParams()
	if err, _ := retrieveOperationContainerContainerAttachDetachKeysFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachLogsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachStderrFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachStdinFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachStdoutFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerAttachStreamFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerAttachResult(appCli.Container.ContainerAttach(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerAttachDetachKeysFlag(m *container.ContainerAttachParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerAttachIDFlag(m *container.ContainerAttachParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerAttachLogsFlag(m *container.ContainerAttachParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerAttachStderrFlag(m *container.ContainerAttachParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerAttachStdinFlag(m *container.ContainerAttachParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerAttachStdoutFlag(m *container.ContainerAttachParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerAttachStreamFlag(m *container.ContainerAttachParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationContainerContainerAttachResult prints output to stdout
func printOperationContainerContainerAttachResult(resp0 *container.ContainerAttachOK, respErr error) error {
	if respErr != nil {

		// Non schema case: warning containerAttachSwitchingProtocols is not supported

		// Non schema case: warning containerAttachOK is not supported

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*container.ContainerAttachBadRequest)
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
		resp3, ok := iResp3.(*container.ContainerAttachNotFound)
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
		resp4, ok := iResp4.(*container.ContainerAttachInternalServerError)
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

	// warning: non schema response containerAttachOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationContainerContainerAttachParamFlags registers all flags needed to fill params
func registerOperationContainerContainerAttachParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerAttachDetachKeysParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachLogsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachStderrParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachStdinParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachStdoutParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerAttachStreamParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerAttachDetachKeysParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	detachKeysDescription := `Override the key sequence for detaching a container.Format is a single character ` + "`" + `[a-Z]` + "`" + ` or ` + "`" + `ctrl-<value>` + "`" + ` where ` + "`" + `<value>` + "`" + ` is one of: ` + "`" + `a-z` + "`" + `, ` + "`" + `@` + "`" + `, ` + "`" + `^` + "`" + `, ` + "`" + `[` + "`" + `, ` + "`" + `,` + "`" + ` or ` + "`" + `_` + "`" + `.`

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
func registerOperationContainerContainerAttachIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerAttachLogsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	logsDescription := `Replay previous logs from the container.

This is useful for attaching to a container that has started and you want to output everything since the container started.

If ` + "`" + `stream` + "`" + ` is also enabled, once all the previous output has been returned, it will seamlessly transition into streaming current output.
`

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
func registerOperationContainerContainerAttachStderrParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerAttachStdinParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerAttachStdoutParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerAttachStreamParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	streamDescription := `Stream attached streams from the time the request was made onwards`

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
