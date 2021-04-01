// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/container"

	"github.com/spf13/cobra"
)

// makeOperationContainerContainerStopCmd returns a cmd to handle operation containerStop
func makeOperationContainerContainerStopCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerStop",
		Short: ``,
		RunE:  runOperationContainerContainerStop,
	}

	if err := registerOperationContainerContainerStopParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerStop uses cmd flags to call endpoint api
func runOperationContainerContainerStop(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerStopParams()
	if err, _ := retrieveOperationContainerContainerStopIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerStopTFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerStopResult(appCli.Container.ContainerStop(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerStopIDFlag(m *container.ContainerStopParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerStopTFlag(m *container.ContainerStopParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("t") {

		var tFlagName string
		if cmdPrefix == "" {
			tFlagName = "t"
		} else {
			tFlagName = fmt.Sprintf("%v.t", cmdPrefix)
		}

		tFlagValue, err := cmd.Flags().GetInt64(tFlagName)
		if err != nil {
			return err, false
		}
		m.T = &tFlagValue

	}
	return nil, retAdded
}

// printOperationContainerContainerStopResult prints output to stdout
func printOperationContainerContainerStopResult(resp0 *container.ContainerStopNoContent, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response containerStopNoContent is not supported by go-swagger cli yet.

	return nil
}

// registerOperationContainerContainerStopParamFlags registers all flags needed to fill params
func registerOperationContainerContainerStopParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerStopIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerStopTParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerStopIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerStopTParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tDescription := `Number of seconds to wait before killing the container`

	var tFlagName string
	if cmdPrefix == "" {
		tFlagName = "t"
	} else {
		tFlagName = fmt.Sprintf("%v.t", cmdPrefix)
	}

	var tFlagDefault int64

	_ = cmd.PersistentFlags().Int64(tFlagName, tFlagDefault, tDescription)

	return nil
}
