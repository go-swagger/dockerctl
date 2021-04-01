// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/models"
	"github.com/spf13/cobra"
)

// register flags to command
func registerModelBuildInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerBuildInfoAux(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBuildInfoError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBuildInfoErrorDetail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBuildInfoID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBuildInfoProgress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBuildInfoProgressDetail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBuildInfoStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerBuildInfoStream(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerBuildInfoAux(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var auxFlagName string
	if cmdPrefix == "" {
		auxFlagName = "aux"
	} else {
		auxFlagName = fmt.Sprintf("%v.aux", cmdPrefix)
	}

	registerModelBuildInfoFlags(depth+1, auxFlagName, cmd)

	return nil
}

func registerBuildInfoError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorDescription := ``

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "error"
	} else {
		errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
	}

	var errorFlagDefault string

	_ = cmd.PersistentFlags().String(errorFlagName, errorFlagDefault, errorDescription)

	return nil
}

func registerBuildInfoErrorDetail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var errorDetailFlagName string
	if cmdPrefix == "" {
		errorDetailFlagName = "errorDetail"
	} else {
		errorDetailFlagName = fmt.Sprintf("%v.errorDetail", cmdPrefix)
	}

	registerModelBuildInfoFlags(depth+1, errorDetailFlagName, cmd)

	return nil
}

func registerBuildInfoID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerBuildInfoProgress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	progressDescription := ``

	var progressFlagName string
	if cmdPrefix == "" {
		progressFlagName = "progress"
	} else {
		progressFlagName = fmt.Sprintf("%v.progress", cmdPrefix)
	}

	var progressFlagDefault string

	_ = cmd.PersistentFlags().String(progressFlagName, progressFlagDefault, progressDescription)

	return nil
}

func registerBuildInfoProgressDetail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var progressDetailFlagName string
	if cmdPrefix == "" {
		progressDetailFlagName = "progressDetail"
	} else {
		progressDetailFlagName = fmt.Sprintf("%v.progressDetail", cmdPrefix)
	}

	registerModelBuildInfoFlags(depth+1, progressDetailFlagName, cmd)

	return nil
}

func registerBuildInfoStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := ``

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

func registerBuildInfoStream(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	streamDescription := ``

	var streamFlagName string
	if cmdPrefix == "" {
		streamFlagName = "stream"
	} else {
		streamFlagName = fmt.Sprintf("%v.stream", cmdPrefix)
	}

	var streamFlagDefault string

	_ = cmd.PersistentFlags().String(streamFlagName, streamFlagDefault, streamDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelBuildInfoFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, auxAdded := retrieveBuildInfoAuxFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || auxAdded

	err, errorAdded := retrieveBuildInfoErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, errorDetailAdded := retrieveBuildInfoErrorDetailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorDetailAdded

	err, idAdded := retrieveBuildInfoIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, progressAdded := retrieveBuildInfoProgressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || progressAdded

	err, progressDetailAdded := retrieveBuildInfoProgressDetailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || progressDetailAdded

	err, statusAdded := retrieveBuildInfoStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, streamAdded := retrieveBuildInfoStreamFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamAdded

	return nil, retAdded
}

func retrieveBuildInfoAuxFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	auxFlagName := fmt.Sprintf("%v.aux", cmdPrefix)
	if cmd.Flags().Changed(auxFlagName) {

		auxFlagValue := &models.BuildInfo{}
		err, added := retrieveModelBuildInfoFlags(depth+1, auxFlagValue, auxFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveBuildInfoErrorFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	errorFlagName := fmt.Sprintf("%v.error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {

		var errorFlagName string
		if cmdPrefix == "" {
			errorFlagName = "error"
		} else {
			errorFlagName = fmt.Sprintf("%v.error", cmdPrefix)
		}

		errorFlagValue, err := cmd.Flags().GetString(errorFlagName)
		if err != nil {
			return err, false
		}
		m.Error = errorFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveBuildInfoErrorDetailFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	errorDetailFlagName := fmt.Sprintf("%v.errorDetail", cmdPrefix)
	if cmd.Flags().Changed(errorDetailFlagName) {

		errorDetailFlagValue := &models.BuildInfo{}
		err, added := retrieveModelBuildInfoFlags(depth+1, errorDetailFlagValue, errorDetailFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveBuildInfoIDFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

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

		retAdded = true
	}
	return nil, retAdded
}

func retrieveBuildInfoProgressFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	progressFlagName := fmt.Sprintf("%v.progress", cmdPrefix)
	if cmd.Flags().Changed(progressFlagName) {

		var progressFlagName string
		if cmdPrefix == "" {
			progressFlagName = "progress"
		} else {
			progressFlagName = fmt.Sprintf("%v.progress", cmdPrefix)
		}

		progressFlagValue, err := cmd.Flags().GetString(progressFlagName)
		if err != nil {
			return err, false
		}
		m.Progress = progressFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveBuildInfoProgressDetailFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	progressDetailFlagName := fmt.Sprintf("%v.progressDetail", cmdPrefix)
	if cmd.Flags().Changed(progressDetailFlagName) {

		progressDetailFlagValue := &models.BuildInfo{}
		err, added := retrieveModelBuildInfoFlags(depth+1, progressDetailFlagValue, progressDetailFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveBuildInfoStatusFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveBuildInfoStreamFlags(depth int, m *models.BuildInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	streamFlagName := fmt.Sprintf("%v.stream", cmdPrefix)
	if cmd.Flags().Changed(streamFlagName) {

		var streamFlagName string
		if cmdPrefix == "" {
			streamFlagName = "stream"
		} else {
			streamFlagName = fmt.Sprintf("%v.stream", cmdPrefix)
		}

		streamFlagValue, err := cmd.Flags().GetString(streamFlagName)
		if err != nil {
			return err, false
		}
		m.Stream = streamFlagValue

		retAdded = true
	}
	return nil, retAdded
}
