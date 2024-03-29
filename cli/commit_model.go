// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for Commit

// register flags to command
func registerModelCommitFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCommitExpected(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCommitID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCommitExpected(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	expectedDescription := `Commit ID of external tool expected by dockerd as set at build time.
`

	var expectedFlagName string
	if cmdPrefix == "" {
		expectedFlagName = "Expected"
	} else {
		expectedFlagName = fmt.Sprintf("%v.Expected", cmdPrefix)
	}

	var expectedFlagDefault string

	_ = cmd.PersistentFlags().String(expectedFlagName, expectedFlagDefault, expectedDescription)

	return nil
}

func registerCommitID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Actual commit ID of external tool.`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "ID"
	} else {
		idFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCommitFlags(depth int, m *models.Commit, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, expectedAdded := retrieveCommitExpectedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || expectedAdded

	err, idAdded := retrieveCommitIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveCommitExpectedFlags(depth int, m *models.Commit, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	expectedFlagName := fmt.Sprintf("%v.Expected", cmdPrefix)
	if cmd.Flags().Changed(expectedFlagName) {

		var expectedFlagName string
		if cmdPrefix == "" {
			expectedFlagName = "Expected"
		} else {
			expectedFlagName = fmt.Sprintf("%v.Expected", cmdPrefix)
		}

		expectedFlagValue, err := cmd.Flags().GetString(expectedFlagName)
		if err != nil {
			return err, false
		}
		m.Expected = expectedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCommitIDFlags(depth int, m *models.Commit, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.ID", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "ID"
		} else {
			idFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
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
