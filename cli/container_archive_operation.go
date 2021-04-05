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

// makeOperationContainerContainerArchiveCmd returns a cmd to handle operation containerArchive
func makeOperationContainerContainerArchiveCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerArchive",
		Short: `Get a tar archive of a resource in the filesystem of container id.`,
		RunE:  runOperationContainerContainerArchive,
	}

	if err := registerOperationContainerContainerArchiveParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerArchive uses cmd flags to call endpoint api
func runOperationContainerContainerArchive(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerArchiveParams()
	if err, _ := retrieveOperationContainerContainerArchiveIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerArchivePathFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerArchiveResult(appCli.Container.ContainerArchive(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerArchiveIDFlag(m *container.ContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerArchivePathFlag(m *container.ContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("path") {

		var pathFlagName string
		if cmdPrefix == "" {
			pathFlagName = "path"
		} else {
			pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
		}

		pathFlagValue, err := cmd.Flags().GetString(pathFlagName)
		if err != nil {
			return err, false
		}
		m.Path = pathFlagValue

	}
	return nil, retAdded
}

// printOperationContainerContainerArchiveResult prints output to stdout
func printOperationContainerContainerArchiveResult(resp0 *container.ContainerArchiveOK, respErr error) error {
	if respErr != nil {

		// Non schema case: warning containerArchiveOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*container.ContainerArchiveBadRequest)
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
		resp2, ok := iResp2.(*container.ContainerArchiveNotFound)
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
		resp3, ok := iResp3.(*container.ContainerArchiveInternalServerError)
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

	// warning: non schema response containerArchiveOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationContainerContainerArchiveParamFlags registers all flags needed to fill params
func registerOperationContainerContainerArchiveParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerArchiveIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerArchivePathParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerArchiveIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerArchivePathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pathDescription := `Required. Resource in the container’s filesystem to archive.`

	var pathFlagName string
	if cmdPrefix == "" {
		pathFlagName = "path"
	} else {
		pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
	}

	var pathFlagDefault string

	_ = cmd.PersistentFlags().String(pathFlagName, pathFlagDefault, pathDescription)

	return nil
}

// register flags to command
func registerModelContainerArchiveBadRequestBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// allOf containerArchiveBadRequestBodyAO0 is not supported by go-swwagger cli yet

	// allOf containerArchiveBadRequestBodyAO1 is not supported by go-swwagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerArchiveBadRequestBodyFlags(depth int, m *container.ContainerArchiveBadRequestBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// allOf containerArchiveBadRequestBodyAO0 is not supported by go-swwagger cli yet

	// allOf containerArchiveBadRequestBodyAO1 is not supported by go-swwagger cli yet

	return nil, retAdded
}
