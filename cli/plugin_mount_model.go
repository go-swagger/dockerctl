// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for PluginMount

// register flags to command
func registerModelPluginMountFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPluginMountDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginMountDestination(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginMountName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginMountOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginMountSettable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginMountSource(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginMountType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPluginMountDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `Required. `

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "Description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.Description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerPluginMountDestination(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	destinationDescription := `Required. `

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

func registerPluginMountName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. `

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

func registerPluginMountOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Options []string array type is not supported by go-swagger cli yet

	return nil
}

func registerPluginMountSettable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Settable []string array type is not supported by go-swagger cli yet

	return nil
}

func registerPluginMountSource(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sourceDescription := `Required. `

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

func registerPluginMountType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Required. `

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
func retrieveModelPluginMountFlags(depth int, m *models.PluginMount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrievePluginMountDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, destinationAdded := retrievePluginMountDestinationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || destinationAdded

	err, nameAdded := retrievePluginMountNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, optionsAdded := retrievePluginMountOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || optionsAdded

	err, settableAdded := retrievePluginMountSettableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || settableAdded

	err, sourceAdded := retrievePluginMountSourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceAdded

	err, typeAdded := retrievePluginMountTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrievePluginMountDescriptionFlags(depth int, m *models.PluginMount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	descriptionFlagName := fmt.Sprintf("%v.Description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "Description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.Description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrievePluginMountDestinationFlags(depth int, m *models.PluginMount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePluginMountNameFlags(depth int, m *models.PluginMount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePluginMountOptionsFlags(depth int, m *models.PluginMount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	optionsFlagName := fmt.Sprintf("%v.Options", cmdPrefix)
	if cmd.Flags().Changed(optionsFlagName) {
		// warning: Options array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrievePluginMountSettableFlags(depth int, m *models.PluginMount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	settableFlagName := fmt.Sprintf("%v.Settable", cmdPrefix)
	if cmd.Flags().Changed(settableFlagName) {
		// warning: Settable array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrievePluginMountSourceFlags(depth int, m *models.PluginMount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Source = &sourceFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrievePluginMountTypeFlags(depth int, m *models.PluginMount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
