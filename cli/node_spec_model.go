// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for NodeSpec

// register flags to command
func registerModelNodeSpecFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNodeSpecAvailability(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNodeSpecLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNodeSpecName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNodeSpecRole(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNodeSpecAvailability(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	availabilityDescription := `Availability of the node.`

	var availabilityFlagName string
	if cmdPrefix == "" {
		availabilityFlagName = "Availability"
	} else {
		availabilityFlagName = fmt.Sprintf("%v.Availability", cmdPrefix)
	}

	var availabilityFlagDefault string

	_ = cmd.PersistentFlags().String(availabilityFlagName, availabilityFlagDefault, availabilityDescription)

	return nil
}

func registerNodeSpecLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerNodeSpecName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name for the node.`

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

func registerNodeSpecRole(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	roleDescription := `Role of the node.`

	var roleFlagName string
	if cmdPrefix == "" {
		roleFlagName = "Role"
	} else {
		roleFlagName = fmt.Sprintf("%v.Role", cmdPrefix)
	}

	var roleFlagDefault string

	_ = cmd.PersistentFlags().String(roleFlagName, roleFlagDefault, roleDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNodeSpecFlags(depth int, m *models.NodeSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, availabilityAdded := retrieveNodeSpecAvailabilityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || availabilityAdded

	err, labelsAdded := retrieveNodeSpecLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, nameAdded := retrieveNodeSpecNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, roleAdded := retrieveNodeSpecRoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || roleAdded

	return nil, retAdded
}

func retrieveNodeSpecAvailabilityFlags(depth int, m *models.NodeSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	availabilityFlagName := fmt.Sprintf("%v.Availability", cmdPrefix)
	if cmd.Flags().Changed(availabilityFlagName) {

		var availabilityFlagName string
		if cmdPrefix == "" {
			availabilityFlagName = "Availability"
		} else {
			availabilityFlagName = fmt.Sprintf("%v.Availability", cmdPrefix)
		}

		availabilityFlagValue, err := cmd.Flags().GetString(availabilityFlagName)
		if err != nil {
			return err, false
		}
		m.Availability = availabilityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNodeSpecLabelsFlags(depth int, m *models.NodeSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	labelsFlagName := fmt.Sprintf("%v.Labels", cmdPrefix)
	if cmd.Flags().Changed(labelsFlagName) {
		// warning: Labels map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNodeSpecNameFlags(depth int, m *models.NodeSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNodeSpecRoleFlags(depth int, m *models.NodeSpec, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	roleFlagName := fmt.Sprintf("%v.Role", cmdPrefix)
	if cmd.Flags().Changed(roleFlagName) {

		var roleFlagName string
		if cmdPrefix == "" {
			roleFlagName = "Role"
		} else {
			roleFlagName = fmt.Sprintf("%v.Role", cmdPrefix)
		}

		roleFlagValue, err := cmd.Flags().GetString(roleFlagName)
		if err != nil {
			return err, false
		}
		m.Role = roleFlagValue

		retAdded = true
	}

	return nil, retAdded
}
