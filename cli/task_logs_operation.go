// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/task"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTaskTaskLogsCmd returns a cmd to handle operation taskLogs
func makeOperationTaskTaskLogsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "TaskLogs",
		Short: `Get ` + "`" + `stdout` + "`" + ` and ` + "`" + `stderr` + "`" + ` logs from a task. See also [` + "`" + `/containers/{id}/logs` + "`" + `](#operation/ContainerLogs).

**Note**: This endpoint works only for services with the ` + "`" + `local` + "`" + `, ` + "`" + `json-file` + "`" + ` or ` + "`" + `journald` + "`" + ` logging drivers.
`,
		RunE: runOperationTaskTaskLogs,
	}

	if err := registerOperationTaskTaskLogsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTaskTaskLogs uses cmd flags to call endpoint api
func runOperationTaskTaskLogs(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := task.NewTaskLogsParams()
	if err, _ := retrieveOperationTaskTaskLogsDetailsFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskTaskLogsFollowFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskTaskLogsIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskTaskLogsSinceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskTaskLogsStderrFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskTaskLogsStdoutFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskTaskLogsTailFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTaskTaskLogsTimestampsFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationTaskTaskLogsResult(appCli.Task.TaskLogs(params, &bytes.Buffer{})); err != nil {
		return err
	}
	return nil
}

func retrieveOperationTaskTaskLogsDetailsFlag(m *task.TaskLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("details") {

		var detailsFlagName string
		if cmdPrefix == "" {
			detailsFlagName = "details"
		} else {
			detailsFlagName = fmt.Sprintf("%v.details", cmdPrefix)
		}

		detailsFlagValue, err := cmd.Flags().GetBool(detailsFlagName)
		if err != nil {
			return err, false
		}
		m.Details = &detailsFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTaskTaskLogsFollowFlag(m *task.TaskLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("follow") {

		var followFlagName string
		if cmdPrefix == "" {
			followFlagName = "follow"
		} else {
			followFlagName = fmt.Sprintf("%v.follow", cmdPrefix)
		}

		followFlagValue, err := cmd.Flags().GetBool(followFlagName)
		if err != nil {
			return err, false
		}
		m.Follow = &followFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTaskTaskLogsIDFlag(m *task.TaskLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTaskTaskLogsSinceFlag(m *task.TaskLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("since") {

		var sinceFlagName string
		if cmdPrefix == "" {
			sinceFlagName = "since"
		} else {
			sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
		}

		sinceFlagValue, err := cmd.Flags().GetInt64(sinceFlagName)
		if err != nil {
			return err, false
		}
		m.Since = &sinceFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTaskTaskLogsStderrFlag(m *task.TaskLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTaskTaskLogsStdoutFlag(m *task.TaskLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationTaskTaskLogsTailFlag(m *task.TaskLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tail") {

		var tailFlagName string
		if cmdPrefix == "" {
			tailFlagName = "tail"
		} else {
			tailFlagName = fmt.Sprintf("%v.tail", cmdPrefix)
		}

		tailFlagValue, err := cmd.Flags().GetString(tailFlagName)
		if err != nil {
			return err, false
		}
		m.Tail = &tailFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTaskTaskLogsTimestampsFlag(m *task.TaskLogsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("timestamps") {

		var timestampsFlagName string
		if cmdPrefix == "" {
			timestampsFlagName = "timestamps"
		} else {
			timestampsFlagName = fmt.Sprintf("%v.timestamps", cmdPrefix)
		}

		timestampsFlagValue, err := cmd.Flags().GetBool(timestampsFlagName)
		if err != nil {
			return err, false
		}
		m.Timestamps = &timestampsFlagValue

	}
	return nil, retAdded
}

// printOperationTaskTaskLogsResult prints output to stdout
func printOperationTaskTaskLogsResult(resp0 *task.TaskLogsOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*task.TaskLogsOK)
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
		resp1, ok := iResp1.(*task.TaskLogsNotFound)
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
		resp2, ok := iResp2.(*task.TaskLogsInternalServerError)
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
		resp3, ok := iResp3.(*task.TaskLogsServiceUnavailable)
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
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationTaskTaskLogsParamFlags registers all flags needed to fill params
func registerOperationTaskTaskLogsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTaskTaskLogsDetailsParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskTaskLogsFollowParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskTaskLogsIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskTaskLogsSinceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskTaskLogsStderrParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskTaskLogsStdoutParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskTaskLogsTailParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTaskTaskLogsTimestampsParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTaskTaskLogsDetailsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	detailsDescription := `Show task context and extra details provided to logs.`

	var detailsFlagName string
	if cmdPrefix == "" {
		detailsFlagName = "details"
	} else {
		detailsFlagName = fmt.Sprintf("%v.details", cmdPrefix)
	}

	var detailsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(detailsFlagName, detailsFlagDefault, detailsDescription)

	return nil
}
func registerOperationTaskTaskLogsFollowParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	followDescription := `Keep connection after returning logs.`

	var followFlagName string
	if cmdPrefix == "" {
		followFlagName = "follow"
	} else {
		followFlagName = fmt.Sprintf("%v.follow", cmdPrefix)
	}

	var followFlagDefault bool

	_ = cmd.PersistentFlags().Bool(followFlagName, followFlagDefault, followDescription)

	return nil
}
func registerOperationTaskTaskLogsIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationTaskTaskLogsSinceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sinceDescription := `Only return logs since this time, as a UNIX timestamp`

	var sinceFlagName string
	if cmdPrefix == "" {
		sinceFlagName = "since"
	} else {
		sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
	}

	var sinceFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sinceFlagName, sinceFlagDefault, sinceDescription)

	return nil
}
func registerOperationTaskTaskLogsStderrParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	stderrDescription := `Return logs from ` + "`" + `stderr` + "`" + ``

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
func registerOperationTaskTaskLogsStdoutParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	stdoutDescription := `Return logs from ` + "`" + `stdout` + "`" + ``

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
func registerOperationTaskTaskLogsTailParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tailDescription := `Only return this number of log lines from the end of the logs. Specify as an integer or ` + "`" + `all` + "`" + ` to output all log lines.`

	var tailFlagName string
	if cmdPrefix == "" {
		tailFlagName = "tail"
	} else {
		tailFlagName = fmt.Sprintf("%v.tail", cmdPrefix)
	}

	var tailFlagDefault string = "all"

	_ = cmd.PersistentFlags().String(tailFlagName, tailFlagDefault, tailDescription)

	return nil
}
func registerOperationTaskTaskLogsTimestampsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	timestampsDescription := `Add timestamps to every log line`

	var timestampsFlagName string
	if cmdPrefix == "" {
		timestampsFlagName = "timestamps"
	} else {
		timestampsFlagName = fmt.Sprintf("%v.timestamps", cmdPrefix)
	}

	var timestampsFlagDefault bool

	_ = cmd.PersistentFlags().Bool(timestampsFlagName, timestampsFlagDefault, timestampsDescription)

	return nil
}
