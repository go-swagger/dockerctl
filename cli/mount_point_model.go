// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for MountPoint

// register flags to command
func registerModelMountPointFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMountPointDestination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointPropagation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointRW(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointSource(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountPointType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountPointDestination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	destinationDescription := ``

	var destinationFlagName string
	if cmdPrefix == "" {
		destinationFlagName = "Destination"
	} else {
		destinationFlagName = fmt.Sprintf("%v.Destination", cmdPrefix)
	}

	var destinationFlagDefault string

	_ = cmd.PersistentFlags().String(destinationFlagName, destinationFlagDefault, destinationDescription)

	return nil
}

func registerMountPointDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerMountPointMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modeDescription := ``

	var modeFlagName string
	if cmdPrefix == "" {
		modeFlagName = "Mode"
	} else {
		modeFlagName = fmt.Sprintf("%v.Mode", cmdPrefix)
	}

	var modeFlagDefault string

	_ = cmd.PersistentFlags().String(modeFlagName, modeFlagDefault, modeDescription)

	return nil
}

func registerMountPointName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "Name"
	} else {
		nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerMountPointPropagation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	propagationDescription := ``

	var propagationFlagName string
	if cmdPrefix == "" {
		propagationFlagName = "Propagation"
	} else {
		propagationFlagName = fmt.Sprintf("%v.Propagation", cmdPrefix)
	}

	var propagationFlagDefault string

	_ = cmd.PersistentFlags().String(propagationFlagName, propagationFlagDefault, propagationDescription)

	return nil
}

func registerMountPointRW(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rWDescription := ``

	var rWFlagName string
	if cmdPrefix == "" {
		rWFlagName = "RW"
	} else {
		rWFlagName = fmt.Sprintf("%v.RW", cmdPrefix)
	}

	var rWFlagDefault bool

	_ = cmd.PersistentFlags().Bool(rWFlagName, rWFlagDefault, rWDescription)

	return nil
}

func registerMountPointSource(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sourceDescription := ``

	var sourceFlagName string
	if cmdPrefix == "" {
		sourceFlagName = "Source"
	} else {
		sourceFlagName = fmt.Sprintf("%v.Source", cmdPrefix)
	}

	var sourceFlagDefault string

	_ = cmd.PersistentFlags().String(sourceFlagName, sourceFlagDefault, sourceDescription)

	return nil
}

func registerMountPointType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := ``

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "Type"
	} else {
		typeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMountPointFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, destinationAdded := retrieveMountPointDestinationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || destinationAdded

	err, driverAdded := retrieveMountPointDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverAdded

	err, modeAdded := retrieveMountPointModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	err, nameAdded := retrieveMountPointNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, propagationAdded := retrieveMountPointPropagationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || propagationAdded

	err, rWAdded := retrieveMountPointRWFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rWAdded

	err, sourceAdded := retrieveMountPointSourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceAdded

	err, typeAdded := retrieveMountPointTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveMountPointDestinationFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	destinationFlagName := fmt.Sprintf("%v.Destination", cmdPrefix)
	if cmd.Flags().Changed(destinationFlagName) {

		var destinationFlagName string
		if cmdPrefix == "" {
			destinationFlagName = "Destination"
		} else {
			destinationFlagName = fmt.Sprintf("%v.Destination", cmdPrefix)
		}

		destinationFlagValue, err := cmd.Flags().GetString(destinationFlagName)
		if err != nil {
			return err, false
		}
		m.Destination = destinationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointDriverFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveMountPointModeFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	modeFlagName := fmt.Sprintf("%v.Mode", cmdPrefix)
	if cmd.Flags().Changed(modeFlagName) {

		var modeFlagName string
		if cmdPrefix == "" {
			modeFlagName = "Mode"
		} else {
			modeFlagName = fmt.Sprintf("%v.Mode", cmdPrefix)
		}

		modeFlagValue, err := cmd.Flags().GetString(modeFlagName)
		if err != nil {
			return err, false
		}
		m.Mode = modeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointNameFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.Name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "Name"
		} else {
			nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointPropagationFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	propagationFlagName := fmt.Sprintf("%v.Propagation", cmdPrefix)
	if cmd.Flags().Changed(propagationFlagName) {

		var propagationFlagName string
		if cmdPrefix == "" {
			propagationFlagName = "Propagation"
		} else {
			propagationFlagName = fmt.Sprintf("%v.Propagation", cmdPrefix)
		}

		propagationFlagValue, err := cmd.Flags().GetString(propagationFlagName)
		if err != nil {
			return err, false
		}
		m.Propagation = propagationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointRWFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rWFlagName := fmt.Sprintf("%v.RW", cmdPrefix)
	if cmd.Flags().Changed(rWFlagName) {

		var rWFlagName string
		if cmdPrefix == "" {
			rWFlagName = "RW"
		} else {
			rWFlagName = fmt.Sprintf("%v.RW", cmdPrefix)
		}

		rWFlagValue, err := cmd.Flags().GetBool(rWFlagName)
		if err != nil {
			return err, false
		}
		m.RW = rWFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointSourceFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sourceFlagName := fmt.Sprintf("%v.Source", cmdPrefix)
	if cmd.Flags().Changed(sourceFlagName) {

		var sourceFlagName string
		if cmdPrefix == "" {
			sourceFlagName = "Source"
		} else {
			sourceFlagName = fmt.Sprintf("%v.Source", cmdPrefix)
		}

		sourceFlagValue, err := cmd.Flags().GetString(sourceFlagName)
		if err != nil {
			return err, false
		}
		m.Source = sourceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountPointTypeFlags(depth int, m *models.MountPoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.Type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "Type"
		} else {
			typeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
