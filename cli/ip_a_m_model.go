// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for IPAM

// register flags to command
func registerModelIPAMFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIPAMConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPAMDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPAMOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIPAMConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Config []map[string]string array type is not supported by go-swagger cli yet

	return nil
}

func registerIPAMDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	driverDescription := `Name of the IPAM driver to use.`

	var driverFlagName string
	if cmdPrefix == "" {
		driverFlagName = "Driver"
	} else {
		driverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
	}

	var driverFlagDefault string = "default"

	_ = cmd.PersistentFlags().String(driverFlagName, driverFlagDefault, driverDescription)

	return nil
}

func registerIPAMOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Options map[string]string map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIPAMFlags(depth int, m *models.IPAM, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, configAdded := retrieveIPAMConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || configAdded

	err, driverAdded := retrieveIPAMDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverAdded

	err, optionsAdded := retrieveIPAMOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || optionsAdded

	return nil, retAdded
}

func retrieveIPAMConfigFlags(depth int, m *models.IPAM, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	configFlagName := fmt.Sprintf("%v.Config", cmdPrefix)
	if cmd.Flags().Changed(configFlagName) {
		// warning: Config array type []map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIPAMDriverFlags(depth int, m *models.IPAM, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	driverFlagName := fmt.Sprintf("%v.Driver", cmdPrefix)
	if cmd.Flags().Changed(driverFlagName) {

		var driverFlagName string
		if cmdPrefix == "" {
			driverFlagName = "Driver"
		} else {
			driverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
		}

		driverFlagValue, err := cmd.Flags().GetString(driverFlagName)
		if err != nil {
			return err, false
		}
		m.Driver = &driverFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIPAMOptionsFlags(depth int, m *models.IPAM, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	optionsFlagName := fmt.Sprintf("%v.Options", cmdPrefix)
	if cmd.Flags().Changed(optionsFlagName) {
		// warning: Options map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
