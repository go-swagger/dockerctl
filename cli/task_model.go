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
func registerModelTaskFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTaskAssignedGenericResources(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskDesiredState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskNodeID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskServiceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskSlot(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTaskVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTaskAssignedGenericResources(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: AssignedGenericResources GenericResources array type is not supported by go-swagger cli yet

	return nil
}

func registerTaskCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTaskDesiredState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive DesiredState TaskState is not supported by go-swagger cli yet

	return nil
}

func registerTaskID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of the task.`

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

func registerTaskLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerTaskName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of the task.`

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

func registerTaskNodeID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodeIdDescription := `The ID of the node that this task is on.`

	var nodeIdFlagName string
	if cmdPrefix == "" {
		nodeIdFlagName = "NodeID"
	} else {
		nodeIdFlagName = fmt.Sprintf("%v.NodeID", cmdPrefix)
	}

	var nodeIdFlagDefault string

	_ = cmd.PersistentFlags().String(nodeIdFlagName, nodeIdFlagDefault, nodeIdDescription)

	return nil
}

func registerTaskServiceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serviceIdDescription := `The ID of the service this task is part of.`

	var serviceIdFlagName string
	if cmdPrefix == "" {
		serviceIdFlagName = "ServiceID"
	} else {
		serviceIdFlagName = fmt.Sprintf("%v.ServiceID", cmdPrefix)
	}

	var serviceIdFlagDefault string

	_ = cmd.PersistentFlags().String(serviceIdFlagName, serviceIdFlagDefault, serviceIdDescription)

	return nil
}

func registerTaskSlot(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	slotDescription := ``

	var slotFlagName string
	if cmdPrefix == "" {
		slotFlagName = "Slot"
	} else {
		slotFlagName = fmt.Sprintf("%v.Slot", cmdPrefix)
	}

	var slotFlagDefault int64

	_ = cmd.PersistentFlags().Int64(slotFlagName, slotFlagDefault, slotDescription)

	return nil
}

func registerTaskSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var specFlagName string
	if cmdPrefix == "" {
		specFlagName = "Spec"
	} else {
		specFlagName = fmt.Sprintf("%v.Spec", cmdPrefix)
	}

	registerModelTaskFlags(depth+1, specFlagName, cmd)

	return nil
}

func registerTaskStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "Status"
	} else {
		statusFlagName = fmt.Sprintf("%v.Status", cmdPrefix)
	}

	registerModelTaskFlags(depth+1, statusFlagName, cmd)

	return nil
}

func registerTaskUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerTaskVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "Version"
	} else {
		versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
	}

	registerModelTaskFlags(depth+1, versionFlagName, cmd)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTaskFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, assignedGenericResourcesAdded := retrieveTaskAssignedGenericResourcesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || assignedGenericResourcesAdded

	err, createdAtAdded := retrieveTaskCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, desiredStateAdded := retrieveTaskDesiredStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || desiredStateAdded

	err, idAdded := retrieveTaskIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, labelsAdded := retrieveTaskLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, nameAdded := retrieveTaskNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, nodeIdAdded := retrieveTaskNodeIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodeIdAdded

	err, serviceIdAdded := retrieveTaskServiceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serviceIdAdded

	err, slotAdded := retrieveTaskSlotFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || slotAdded

	err, specAdded := retrieveTaskSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || specAdded

	err, statusAdded := retrieveTaskStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, updatedAtAdded := retrieveTaskUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	err, versionAdded := retrieveTaskVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveTaskAssignedGenericResourcesFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	assignedGenericResourcesFlagName := fmt.Sprintf("%v.AssignedGenericResources", cmdPrefix)
	if cmd.Flags().Changed(assignedGenericResourcesFlagName) {
		// warning: AssignedGenericResources array type GenericResources is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveTaskCreatedAtFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTaskDesiredStateFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	desiredStateFlagName := fmt.Sprintf("%v.DesiredState", cmdPrefix)
	if cmd.Flags().Changed(desiredStateFlagName) {

		// warning: primitive DesiredState TaskState is not supported by go-swagger cli yet

		retAdded = true
	}
	return nil, retAdded
}

func retrieveTaskIDFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTaskLabelsFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTaskNameFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTaskNodeIDFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	nodeIdFlagName := fmt.Sprintf("%v.NodeID", cmdPrefix)
	if cmd.Flags().Changed(nodeIdFlagName) {

		var nodeIdFlagName string
		if cmdPrefix == "" {
			nodeIdFlagName = "NodeID"
		} else {
			nodeIdFlagName = fmt.Sprintf("%v.NodeID", cmdPrefix)
		}

		nodeIdFlagValue, err := cmd.Flags().GetString(nodeIdFlagName)
		if err != nil {
			return err, false
		}
		m.NodeID = nodeIdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveTaskServiceIDFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	serviceIdFlagName := fmt.Sprintf("%v.ServiceID", cmdPrefix)
	if cmd.Flags().Changed(serviceIdFlagName) {

		var serviceIdFlagName string
		if cmdPrefix == "" {
			serviceIdFlagName = "ServiceID"
		} else {
			serviceIdFlagName = fmt.Sprintf("%v.ServiceID", cmdPrefix)
		}

		serviceIdFlagValue, err := cmd.Flags().GetString(serviceIdFlagName)
		if err != nil {
			return err, false
		}
		m.ServiceID = serviceIdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveTaskSlotFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	slotFlagName := fmt.Sprintf("%v.Slot", cmdPrefix)
	if cmd.Flags().Changed(slotFlagName) {

		var slotFlagName string
		if cmdPrefix == "" {
			slotFlagName = "Slot"
		} else {
			slotFlagName = fmt.Sprintf("%v.Slot", cmdPrefix)
		}

		slotFlagValue, err := cmd.Flags().GetInt64(slotFlagName)
		if err != nil {
			return err, false
		}
		m.Slot = slotFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveTaskSpecFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	specFlagName := fmt.Sprintf("%v.Spec", cmdPrefix)
	if cmd.Flags().Changed(specFlagName) {

		specFlagValue := &models.Task{}
		err, added := retrieveModelTaskFlags(depth+1, specFlagValue, specFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveTaskStatusFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	statusFlagName := fmt.Sprintf("%v.Status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		statusFlagValue := &models.Task{}
		err, added := retrieveModelTaskFlags(depth+1, statusFlagValue, statusFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveTaskUpdatedAtFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveTaskVersionFlags(depth int, m *models.Task, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	versionFlagName := fmt.Sprintf("%v.Version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		versionFlagValue := &models.Task{}
		err, added := retrieveModelTaskFlags(depth+1, versionFlagValue, versionFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}
