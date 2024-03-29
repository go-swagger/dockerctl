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

// makeOperationContainerPutContainerArchiveCmd returns a cmd to handle operation putContainerArchive
func makeOperationContainerPutContainerArchiveCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PutContainerArchive",
		Short: `Upload a tar archive to be extracted to a path in the filesystem of container id.`,
		RunE:  runOperationContainerPutContainerArchive,
	}

	if err := registerOperationContainerPutContainerArchiveParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerPutContainerArchive uses cmd flags to call endpoint api
func runOperationContainerPutContainerArchive(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewPutContainerArchiveParams()
	if err, _ := retrieveOperationContainerPutContainerArchiveCopyUIDGIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerPutContainerArchiveIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerPutContainerArchiveInputStreamFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerPutContainerArchiveNoOverwriteDirNonDirFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerPutContainerArchivePathFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerPutContainerArchiveResult(appCli.Container.PutContainerArchive(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerPutContainerArchiveParamFlags registers all flags needed to fill params
func registerOperationContainerPutContainerArchiveParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerPutContainerArchiveCopyUIDGIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerPutContainerArchiveIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerPutContainerArchiveInputStreamParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerPutContainerArchiveNoOverwriteDirNonDirParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerPutContainerArchivePathParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerPutContainerArchiveCopyUIDGIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	copyUidGIdDescription := `If “1”, “true”, then it will copy UID/GID maps to the dest file or dir`

	var copyUidGIdFlagName string
	if cmdPrefix == "" {
		copyUidGIdFlagName = "copyUIDGID"
	} else {
		copyUidGIdFlagName = fmt.Sprintf("%v.copyUIDGID", cmdPrefix)
	}

	var copyUidGIdFlagDefault string

	_ = cmd.PersistentFlags().String(copyUidGIdFlagName, copyUidGIdFlagDefault, copyUidGIdDescription)

	return nil
}
func registerOperationContainerPutContainerArchiveIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerPutContainerArchiveInputStreamParamFlags(cmdPrefix string, cmd *cobra.Command) error {
	// warning: go type io.ReadCloser is not supported by go-swagger cli yet.
	return nil
}
func registerOperationContainerPutContainerArchiveNoOverwriteDirNonDirParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	noOverwriteDirNonDirDescription := `If “1”, “true”, or “True” then it will be an error if unpacking the given content would cause an existing directory to be replaced with a non-directory and vice versa.`

	var noOverwriteDirNonDirFlagName string
	if cmdPrefix == "" {
		noOverwriteDirNonDirFlagName = "noOverwriteDirNonDir"
	} else {
		noOverwriteDirNonDirFlagName = fmt.Sprintf("%v.noOverwriteDirNonDir", cmdPrefix)
	}

	var noOverwriteDirNonDirFlagDefault string

	_ = cmd.PersistentFlags().String(noOverwriteDirNonDirFlagName, noOverwriteDirNonDirFlagDefault, noOverwriteDirNonDirDescription)

	return nil
}
func registerOperationContainerPutContainerArchivePathParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pathDescription := `Required. Path to a directory in the container to extract the archive’s contents into. `

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

func retrieveOperationContainerPutContainerArchiveCopyUIDGIDFlag(m *container.PutContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("copyUIDGID") {

		var copyUidGIdFlagName string
		if cmdPrefix == "" {
			copyUidGIdFlagName = "copyUIDGID"
		} else {
			copyUidGIdFlagName = fmt.Sprintf("%v.copyUIDGID", cmdPrefix)
		}

		copyUidGIdFlagValue, err := cmd.Flags().GetString(copyUidGIdFlagName)
		if err != nil {
			return err, false
		}
		m.CopyUIDGID = &copyUidGIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerPutContainerArchiveIDFlag(m *container.PutContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerPutContainerArchiveInputStreamFlag(m *container.PutContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("inputStream") {
		// warning: io.ReadCloser is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
func retrieveOperationContainerPutContainerArchiveNoOverwriteDirNonDirFlag(m *container.PutContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("noOverwriteDirNonDir") {

		var noOverwriteDirNonDirFlagName string
		if cmdPrefix == "" {
			noOverwriteDirNonDirFlagName = "noOverwriteDirNonDir"
		} else {
			noOverwriteDirNonDirFlagName = fmt.Sprintf("%v.noOverwriteDirNonDir", cmdPrefix)
		}

		noOverwriteDirNonDirFlagValue, err := cmd.Flags().GetString(noOverwriteDirNonDirFlagName)
		if err != nil {
			return err, false
		}
		m.NoOverwriteDirNonDir = &noOverwriteDirNonDirFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerPutContainerArchivePathFlag(m *container.PutContainerArchiveParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationContainerPutContainerArchiveResult parses request result and return the string content
func parseOperationContainerPutContainerArchiveResult(resp0 *container.PutContainerArchiveOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning putContainerArchiveOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*container.PutContainerArchiveBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*container.PutContainerArchiveForbidden)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*container.PutContainerArchiveNotFound)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*container.PutContainerArchiveInternalServerError)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response putContainerArchiveOK is not supported by go-swagger cli yet.

	return "", nil
}
