// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/container"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationContainerContainerStatsCmd returns a cmd to handle operation containerStats
func makeOperationContainerContainerStatsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "ContainerStats",
		Short: `This endpoint returns a live stream of a container’s resource usage
statistics.

The ` + "`" + `precpu_stats` + "`" + ` is the CPU statistic of the *previous* read, and is
used to calculate the CPU usage percentage. It is not an exact copy
of the ` + "`" + `cpu_stats` + "`" + ` field.

If either ` + "`" + `precpu_stats.online_cpus` + "`" + ` or ` + "`" + `cpu_stats.online_cpus` + "`" + ` is
nil then for compatibility with older daemons the length of the
corresponding ` + "`" + `cpu_usage.percpu_usage` + "`" + ` array should be used.
`,
		RunE: runOperationContainerContainerStats,
	}

	if err := registerOperationContainerContainerStatsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerStats uses cmd flags to call endpoint api
func runOperationContainerContainerStats(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerStatsParams()
	if err, _ := retrieveOperationContainerContainerStatsIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerStatsStreamFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerStatsResult(appCli.Container.ContainerStats(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerStatsIDFlag(m *container.ContainerStatsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerStatsStreamFlag(m *container.ContainerStatsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationContainerContainerStatsResult prints output to stdout
func printOperationContainerContainerStatsResult(resp0 *container.ContainerStatsOK, respErr error) error {
	if respErr != nil {
		return respErr
	}

	if !swag.IsZero(resp0.Payload) {
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationContainerContainerStatsParamFlags registers all flags needed to fill params
func registerOperationContainerContainerStatsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerStatsIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerStatsStreamParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerStatsIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerStatsStreamParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	streamDescription := `Stream the output. If false, the stats will be output once and then it will disconnect.`

	var streamFlagName string
	if cmdPrefix == "" {
		streamFlagName = "stream"
	} else {
		streamFlagName = fmt.Sprintf("%v.stream", cmdPrefix)
	}

	var streamFlagDefault bool = true

	_ = cmd.PersistentFlags().Bool(streamFlagName, streamFlagDefault, streamDescription)

	return nil
}