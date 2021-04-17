// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/go-swagger/dockerctl/models"

	"github.com/spf13/cobra"
)

// Schema cli for PushImageInfo

// register flags to command
func registerModelPushImageInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPushImageInfoError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPushImageInfoProgress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPushImageInfoProgressDetail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPushImageInfoStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPushImageInfoError(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPushImageInfoProgress(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerPushImageInfoProgressDetail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var progressDetailFlagName string
	if cmdPrefix == "" {
		progressDetailFlagName = "progressDetail"
	} else {
		progressDetailFlagName = fmt.Sprintf("%v.progressDetail", cmdPrefix)
	}

	if err := registerModelProgressDetailFlags(depth+1, progressDetailFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerPushImageInfoStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPushImageInfoFlags(depth int, m *models.PushImageInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorAdded := retrievePushImageInfoErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, progressAdded := retrievePushImageInfoProgressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || progressAdded

	err, progressDetailAdded := retrievePushImageInfoProgressDetailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || progressDetailAdded

	err, statusAdded := retrievePushImageInfoStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrievePushImageInfoErrorFlags(depth int, m *models.PushImageInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePushImageInfoProgressFlags(depth int, m *models.PushImageInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePushImageInfoProgressDetailFlags(depth int, m *models.PushImageInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	progressDetailFlagName := fmt.Sprintf("%v.progressDetail", cmdPrefix)
	if cmd.Flags().Changed(progressDetailFlagName) {
		// info: complex object progressDetail ProgressDetail is retrieved outside this Changed() block
	}
	progressDetailFlagValue := m.ProgressDetail
	if swag.IsZero(progressDetailFlagValue) {
		progressDetailFlagValue = &models.ProgressDetail{}
	}

	err, progressDetailAdded := retrieveModelProgressDetailFlags(depth+1, progressDetailFlagValue, progressDetailFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || progressDetailAdded
	if progressDetailAdded {
		m.ProgressDetail = progressDetailFlagValue
	}

	return nil, retAdded
}

func retrievePushImageInfoStatusFlags(depth int, m *models.PushImageInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
