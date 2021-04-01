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
func registerModelRuntimeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRuntimePath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRuntimeRuntimeArgs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRuntimePath(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pathDescription := `Name and, optional, path, of the OCI executable binary.

If the path is omitted, the daemon searches the host's ` + "`" + `$PATH` + "`" + ` for the
binary and uses the first result.
`

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

func registerRuntimeRuntimeArgs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: runtimeArgs []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRuntimeFlags(depth int, m *models.Runtime, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, pathAdded := retrieveRuntimePathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pathAdded

	err, runtimeArgsAdded := retrieveRuntimeRuntimeArgsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || runtimeArgsAdded

	return nil, retAdded
}

func retrieveRuntimePathFlags(depth int, m *models.Runtime, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	pathFlagName := fmt.Sprintf("%v.path", cmdPrefix)
	if cmd.Flags().Changed(pathFlagName) {

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

		retAdded = true
	}
	return nil, retAdded
}

func retrieveRuntimeRuntimeArgsFlags(depth int, m *models.Runtime, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	runtimeArgsFlagName := fmt.Sprintf("%v.runtimeArgs", cmdPrefix)
	if cmd.Flags().Changed(runtimeArgsFlagName) {
		// warning: runtimeArgs array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
