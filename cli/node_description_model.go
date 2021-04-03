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
func registerModelNodeDescriptionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNodeDescriptionEngine(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNodeDescriptionHostname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNodeDescriptionPlatform(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNodeDescriptionResources(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNodeDescriptionTLSInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNodeDescriptionEngine(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var engineFlagName string
	if cmdPrefix == "" {
		engineFlagName = "Engine"
	} else {
		engineFlagName = fmt.Sprintf("%v.Engine", cmdPrefix)
	}

	registerModelNodeDescriptionFlags(depth+1, engineFlagName, cmd)

	return nil
}

func registerNodeDescriptionHostname(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hostnameDescription := ``

	var hostnameFlagName string
	if cmdPrefix == "" {
		hostnameFlagName = "Hostname"
	} else {
		hostnameFlagName = fmt.Sprintf("%v.Hostname", cmdPrefix)
	}

	var hostnameFlagDefault string

	_ = cmd.PersistentFlags().String(hostnameFlagName, hostnameFlagDefault, hostnameDescription)

	return nil
}

func registerNodeDescriptionPlatform(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var platformFlagName string
	if cmdPrefix == "" {
		platformFlagName = "Platform"
	} else {
		platformFlagName = fmt.Sprintf("%v.Platform", cmdPrefix)
	}

	registerModelNodeDescriptionFlags(depth+1, platformFlagName, cmd)

	return nil
}

func registerNodeDescriptionResources(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var resourcesFlagName string
	if cmdPrefix == "" {
		resourcesFlagName = "Resources"
	} else {
		resourcesFlagName = fmt.Sprintf("%v.Resources", cmdPrefix)
	}

	registerModelNodeDescriptionFlags(depth+1, resourcesFlagName, cmd)

	return nil
}

func registerNodeDescriptionTLSInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var tlsInfoFlagName string
	if cmdPrefix == "" {
		tlsInfoFlagName = "TLSInfo"
	} else {
		tlsInfoFlagName = fmt.Sprintf("%v.TLSInfo", cmdPrefix)
	}

	registerModelNodeDescriptionFlags(depth+1, tlsInfoFlagName, cmd)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNodeDescriptionFlags(depth int, m *models.NodeDescription, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, engineAdded := retrieveNodeDescriptionEngineFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || engineAdded

	err, hostnameAdded := retrieveNodeDescriptionHostnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostnameAdded

	err, platformAdded := retrieveNodeDescriptionPlatformFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || platformAdded

	err, resourcesAdded := retrieveNodeDescriptionResourcesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || resourcesAdded

	err, tlsInfoAdded := retrieveNodeDescriptionTLSInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tlsInfoAdded

	return nil, retAdded
}

func retrieveNodeDescriptionEngineFlags(depth int, m *models.NodeDescription, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	engineFlagName := fmt.Sprintf("%v.Engine", cmdPrefix)
	if cmd.Flags().Changed(engineFlagName) {

		engineFlagValue := &models.NodeDescription{}
		err, added := retrieveModelNodeDescriptionFlags(depth+1, engineFlagValue, engineFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveNodeDescriptionHostnameFlags(depth int, m *models.NodeDescription, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	hostnameFlagName := fmt.Sprintf("%v.Hostname", cmdPrefix)
	if cmd.Flags().Changed(hostnameFlagName) {

		var hostnameFlagName string
		if cmdPrefix == "" {
			hostnameFlagName = "Hostname"
		} else {
			hostnameFlagName = fmt.Sprintf("%v.Hostname", cmdPrefix)
		}

		hostnameFlagValue, err := cmd.Flags().GetString(hostnameFlagName)
		if err != nil {
			return err, false
		}
		m.Hostname = hostnameFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNodeDescriptionPlatformFlags(depth int, m *models.NodeDescription, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	platformFlagName := fmt.Sprintf("%v.Platform", cmdPrefix)
	if cmd.Flags().Changed(platformFlagName) {

		platformFlagValue := &models.NodeDescription{}
		err, added := retrieveModelNodeDescriptionFlags(depth+1, platformFlagValue, platformFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveNodeDescriptionResourcesFlags(depth int, m *models.NodeDescription, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	resourcesFlagName := fmt.Sprintf("%v.Resources", cmdPrefix)
	if cmd.Flags().Changed(resourcesFlagName) {

		resourcesFlagValue := &models.NodeDescription{}
		err, added := retrieveModelNodeDescriptionFlags(depth+1, resourcesFlagValue, resourcesFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveNodeDescriptionTLSInfoFlags(depth int, m *models.NodeDescription, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	tlsInfoFlagName := fmt.Sprintf("%v.TLSInfo", cmdPrefix)
	if cmd.Flags().Changed(tlsInfoFlagName) {

		tlsInfoFlagValue := &models.NodeDescription{}
		err, added := retrieveModelNodeDescriptionFlags(depth+1, tlsInfoFlagValue, tlsInfoFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}