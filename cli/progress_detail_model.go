// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for ProgressDetail

// register flags to command
func registerModelProgressDetailFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerProgressDetailCurrent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerProgressDetailTotal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerProgressDetailCurrent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	currentDescription := ``

	var currentFlagName string
	if cmdPrefix == "" {
		currentFlagName = "current"
	} else {
		currentFlagName = fmt.Sprintf("%v.current", cmdPrefix)
	}

	var currentFlagDefault int64

	_ = cmd.PersistentFlags().Int64(currentFlagName, currentFlagDefault, currentDescription)

	return nil
}

func registerProgressDetailTotal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalDescription := ``

	var totalFlagName string
	if cmdPrefix == "" {
		totalFlagName = "total"
	} else {
		totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
	}

	var totalFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalFlagName, totalFlagDefault, totalDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelProgressDetailFlags(depth int, m *models.ProgressDetail, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, currentAdded := retrieveProgressDetailCurrentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currentAdded

	err, totalAdded := retrieveProgressDetailTotalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalAdded

	return nil, retAdded
}

func retrieveProgressDetailCurrentFlags(depth int, m *models.ProgressDetail, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	currentFlagName := fmt.Sprintf("%v.current", cmdPrefix)
	if cmd.Flags().Changed(currentFlagName) {

		var currentFlagName string
		if cmdPrefix == "" {
			currentFlagName = "current"
		} else {
			currentFlagName = fmt.Sprintf("%v.current", cmdPrefix)
		}

		currentFlagValue, err := cmd.Flags().GetInt64(currentFlagName)
		if err != nil {
			return err, false
		}
		m.Current = currentFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveProgressDetailTotalFlags(depth int, m *models.ProgressDetail, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	totalFlagName := fmt.Sprintf("%v.total", cmdPrefix)
	if cmd.Flags().Changed(totalFlagName) {

		var totalFlagName string
		if cmdPrefix == "" {
			totalFlagName = "total"
		} else {
			totalFlagName = fmt.Sprintf("%v.total", cmdPrefix)
		}

		totalFlagValue, err := cmd.Flags().GetInt64(totalFlagName)
		if err != nil {
			return err, false
		}
		m.Total = totalFlagValue

		retAdded = true
	}
	return nil, retAdded
}
