// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// register flags to command
func registerModelServiceFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServiceCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceEndpoint(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceServiceStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceUpdateStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServiceVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServiceCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdAtDescription := ``

	var createdAtFlagName string
	if cmdPrefix == "" {
		createdAtFlagName = "CreatedAt"
	} else {
		createdAtFlagName = fmt.Sprintf("%v.CreatedAt", cmdPrefix)
	}

	var createdAtFlagDefault string

	_ = cmd.PersistentFlags().String(createdAtFlagName, createdAtFlagDefault, createdAtDescription)

	return nil
}

func registerServiceEndpoint(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var endpointFlagName string
	if cmdPrefix == "" {
		endpointFlagName = "Endpoint"
	} else {
		endpointFlagName = fmt.Sprintf("%v.Endpoint", cmdPrefix)
	}

	registerModelServiceFlags(depth+1, endpointFlagName, cmd)

	return nil
}

func registerServiceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerServiceServiceStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var serviceStatusFlagName string
	if cmdPrefix == "" {
		serviceStatusFlagName = "ServiceStatus"
	} else {
		serviceStatusFlagName = fmt.Sprintf("%v.ServiceStatus", cmdPrefix)
	}

	registerModelServiceFlags(depth+1, serviceStatusFlagName, cmd)

	return nil
}

func registerServiceSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var specFlagName string
	if cmdPrefix == "" {
		specFlagName = "Spec"
	} else {
		specFlagName = fmt.Sprintf("%v.Spec", cmdPrefix)
	}

	registerModelServiceFlags(depth+1, specFlagName, cmd)

	return nil
}

func registerServiceUpdateStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var updateStatusFlagName string
	if cmdPrefix == "" {
		updateStatusFlagName = "UpdateStatus"
	} else {
		updateStatusFlagName = fmt.Sprintf("%v.UpdateStatus", cmdPrefix)
	}

	registerModelServiceFlags(depth+1, updateStatusFlagName, cmd)

	return nil
}

func registerServiceUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updatedAtDescription := ``

	var updatedAtFlagName string
	if cmdPrefix == "" {
		updatedAtFlagName = "UpdatedAt"
	} else {
		updatedAtFlagName = fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
	}

	var updatedAtFlagDefault string

	_ = cmd.PersistentFlags().String(updatedAtFlagName, updatedAtFlagDefault, updatedAtDescription)

	return nil
}

func registerServiceVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "Version"
	} else {
		versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
	}

	registerModelServiceFlags(depth+1, versionFlagName, cmd)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServiceFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, createdAtAdded := retrieveServiceCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, endpointAdded := retrieveServiceEndpointFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointAdded

	err, idAdded := retrieveServiceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, serviceStatusAdded := retrieveServiceServiceStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serviceStatusAdded

	err, specAdded := retrieveServiceSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || specAdded

	err, updateStatusAdded := retrieveServiceUpdateStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updateStatusAdded

	err, updatedAtAdded := retrieveServiceUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	err, versionAdded := retrieveServiceVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveServiceCreatedAtFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	createdAtFlagName := fmt.Sprintf("%v.CreatedAt", cmdPrefix)
	if cmd.Flags().Changed(createdAtFlagName) {

		var createdAtFlagName string
		if cmdPrefix == "" {
			createdAtFlagName = "CreatedAt"
		} else {
			createdAtFlagName = fmt.Sprintf("%v.CreatedAt", cmdPrefix)
		}

		createdAtFlagValue, err := cmd.Flags().GetString(createdAtFlagName)
		if err != nil {
			return err, false
		}
		m.CreatedAt = createdAtFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveServiceEndpointFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	endpointFlagName := fmt.Sprintf("%v.Endpoint", cmdPrefix)
	if cmd.Flags().Changed(endpointFlagName) {

		endpointFlagValue := &models.Service{}
		err, added := retrieveModelServiceFlags(depth+1, endpointFlagValue, endpointFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveServiceIDFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveServiceServiceStatusFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	serviceStatusFlagName := fmt.Sprintf("%v.ServiceStatus", cmdPrefix)
	if cmd.Flags().Changed(serviceStatusFlagName) {

		serviceStatusFlagValue := &models.Service{}
		err, added := retrieveModelServiceFlags(depth+1, serviceStatusFlagValue, serviceStatusFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveServiceSpecFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	specFlagName := fmt.Sprintf("%v.Spec", cmdPrefix)
	if cmd.Flags().Changed(specFlagName) {

		specFlagValue := &models.Service{}
		err, added := retrieveModelServiceFlags(depth+1, specFlagValue, specFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveServiceUpdateStatusFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	updateStatusFlagName := fmt.Sprintf("%v.UpdateStatus", cmdPrefix)
	if cmd.Flags().Changed(updateStatusFlagName) {

		updateStatusFlagValue := &models.Service{}
		err, added := retrieveModelServiceFlags(depth+1, updateStatusFlagValue, updateStatusFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveServiceUpdatedAtFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	updatedAtFlagName := fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
	if cmd.Flags().Changed(updatedAtFlagName) {

		var updatedAtFlagName string
		if cmdPrefix == "" {
			updatedAtFlagName = "UpdatedAt"
		} else {
			updatedAtFlagName = fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
		}

		updatedAtFlagValue, err := cmd.Flags().GetString(updatedAtFlagName)
		if err != nil {
			return err, false
		}
		m.UpdatedAt = updatedAtFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveServiceVersionFlags(depth int, m *models.Service, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	versionFlagName := fmt.Sprintf("%v.Version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		versionFlagValue := &models.Service{}
		err, added := retrieveModelServiceFlags(depth+1, versionFlagValue, versionFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}
