// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for DeviceRequest

// register flags to command
func registerModelDeviceRequestFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDeviceRequestCapabilities(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceRequestCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceRequestDeviceIDs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceRequestDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDeviceRequestOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDeviceRequestCapabilities(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Capabilities [][]string array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceRequestCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	countDescription := ``

	var countFlagName string
	if cmdPrefix == "" {
		countFlagName = "Count"
	} else {
		countFlagName = fmt.Sprintf("%v.Count", cmdPrefix)
	}

	var countFlagDefault int64

	_ = cmd.PersistentFlags().Int64(countFlagName, countFlagDefault, countDescription)

	return nil
}

func registerDeviceRequestDeviceIDs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: DeviceIDs []string array type is not supported by go-swagger cli yet

	return nil
}

func registerDeviceRequestDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	driverDescription := ``

	var driverFlagName string
	if cmdPrefix == "" {
		driverFlagName = "Driver"
	} else {
		driverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
	}

	var driverFlagDefault string

	_ = cmd.PersistentFlags().String(driverFlagName, driverFlagDefault, driverDescription)

	return nil
}

func registerDeviceRequestOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Options map[string]string map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDeviceRequestFlags(depth int, m *models.DeviceRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, capabilitiesAdded := retrieveDeviceRequestCapabilitiesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || capabilitiesAdded

	err, countAdded := retrieveDeviceRequestCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || countAdded

	err, deviceIDsAdded := retrieveDeviceRequestDeviceIDsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deviceIDsAdded

	err, driverAdded := retrieveDeviceRequestDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverAdded

	err, optionsAdded := retrieveDeviceRequestOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || optionsAdded

	return nil, retAdded
}

func retrieveDeviceRequestCapabilitiesFlags(depth int, m *models.DeviceRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	capabilitiesFlagName := fmt.Sprintf("%v.Capabilities", cmdPrefix)
	if cmd.Flags().Changed(capabilitiesFlagName) {
		// warning: Capabilities array type [][]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceRequestCountFlags(depth int, m *models.DeviceRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	countFlagName := fmt.Sprintf("%v.Count", cmdPrefix)
	if cmd.Flags().Changed(countFlagName) {

		var countFlagName string
		if cmdPrefix == "" {
			countFlagName = "Count"
		} else {
			countFlagName = fmt.Sprintf("%v.Count", cmdPrefix)
		}

		countFlagValue, err := cmd.Flags().GetInt64(countFlagName)
		if err != nil {
			return err, false
		}
		m.Count = countFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceRequestDeviceIDsFlags(depth int, m *models.DeviceRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deviceIDsFlagName := fmt.Sprintf("%v.DeviceIDs", cmdPrefix)
	if cmd.Flags().Changed(deviceIDsFlagName) {
		// warning: DeviceIDs array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDeviceRequestDriverFlags(depth int, m *models.DeviceRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Driver = driverFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDeviceRequestOptionsFlags(depth int, m *models.DeviceRequest, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
