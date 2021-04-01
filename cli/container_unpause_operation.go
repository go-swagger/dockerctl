// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/container"

	"github.com/spf13/cobra"
)

// makeOperationContainerContainerUnpauseCmd returns a cmd to handle operation containerUnpause
func makeOperationContainerContainerUnpauseCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerUnpause",
		Short: `Resume a container which has been paused.`,
		RunE:  runOperationContainerContainerUnpause,
	}

	if err := registerOperationContainerContainerUnpauseParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerUnpause uses cmd flags to call endpoint api
func runOperationContainerContainerUnpause(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerUnpauseParams()
	if err, _ := retrieveOperationContainerContainerUnpauseIDFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerUnpauseResult(appCli.Container.ContainerUnpause(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerUnpauseIDFlag(m *container.ContainerUnpauseParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationContainerContainerUnpauseResult prints output to stdout
func printOperationContainerContainerUnpauseResult(resp0 *container.ContainerUnpauseNoContent, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response containerUnpauseNoContent is not supported by go-swagger cli yet.

	return nil
}

// registerOperationContainerContainerUnpauseParamFlags registers all flags needed to fill params
func registerOperationContainerContainerUnpauseParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerUnpauseIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerUnpauseIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
