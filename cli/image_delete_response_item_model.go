// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for ImageDeleteResponseItem

// register flags to command
func registerModelImageDeleteResponseItemFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageDeleteResponseItemDeleted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageDeleteResponseItemUntagged(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageDeleteResponseItemDeleted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deletedDescription := `The image ID of an image that was deleted`

	var deletedFlagName string
	if cmdPrefix == "" {
		deletedFlagName = "Deleted"
	} else {
		deletedFlagName = fmt.Sprintf("%v.Deleted", cmdPrefix)
	}

	var deletedFlagDefault string

	_ = cmd.PersistentFlags().String(deletedFlagName, deletedFlagDefault, deletedDescription)

	return nil
}

func registerImageDeleteResponseItemUntagged(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	untaggedDescription := `The image ID of an image that was untagged`

	var untaggedFlagName string
	if cmdPrefix == "" {
		untaggedFlagName = "Untagged"
	} else {
		untaggedFlagName = fmt.Sprintf("%v.Untagged", cmdPrefix)
	}

	var untaggedFlagDefault string

	_ = cmd.PersistentFlags().String(untaggedFlagName, untaggedFlagDefault, untaggedDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageDeleteResponseItemFlags(depth int, m *models.ImageDeleteResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, deletedAdded := retrieveImageDeleteResponseItemDeletedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deletedAdded

	err, untaggedAdded := retrieveImageDeleteResponseItemUntaggedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || untaggedAdded

	return nil, retAdded
}

func retrieveImageDeleteResponseItemDeletedFlags(depth int, m *models.ImageDeleteResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deletedFlagName := fmt.Sprintf("%v.Deleted", cmdPrefix)
	if cmd.Flags().Changed(deletedFlagName) {

		var deletedFlagName string
		if cmdPrefix == "" {
			deletedFlagName = "Deleted"
		} else {
			deletedFlagName = fmt.Sprintf("%v.Deleted", cmdPrefix)
		}

		deletedFlagValue, err := cmd.Flags().GetString(deletedFlagName)
		if err != nil {
			return err, false
		}
		m.Deleted = deletedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveImageDeleteResponseItemUntaggedFlags(depth int, m *models.ImageDeleteResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	untaggedFlagName := fmt.Sprintf("%v.Untagged", cmdPrefix)
	if cmd.Flags().Changed(untaggedFlagName) {

		var untaggedFlagName string
		if cmdPrefix == "" {
			untaggedFlagName = "Untagged"
		} else {
			untaggedFlagName = fmt.Sprintf("%v.Untagged", cmdPrefix)
		}

		untaggedFlagValue, err := cmd.Flags().GetString(untaggedFlagName)
		if err != nil {
			return err, false
		}
		m.Untagged = untaggedFlagValue

		retAdded = true
	}

	return nil, retAdded
}
