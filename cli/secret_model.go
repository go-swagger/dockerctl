// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/go-swagger/dockerctl/models"

	"github.com/spf13/cobra"
)

// Schema cli for Secret

// register flags to command
func registerModelSecretFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSecretCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecretID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecretSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecretUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSecretVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSecretCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerSecretID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerSecretSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var specFlagName string
	if cmdPrefix == "" {
		specFlagName = "Spec"
	} else {
		specFlagName = fmt.Sprintf("%v.Spec", cmdPrefix)
	}

	if err := registerModelSecretSpecFlags(depth+1, specFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerSecretUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerSecretVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "Version"
	} else {
		versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
	}

	if err := registerModelObjectVersionFlags(depth+1, versionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSecretFlags(depth int, m *models.Secret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, createdAtAdded := retrieveSecretCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, idAdded := retrieveSecretIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, specAdded := retrieveSecretSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || specAdded

	err, updatedAtAdded := retrieveSecretUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	err, versionAdded := retrieveSecretVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveSecretCreatedAtFlags(depth int, m *models.Secret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveSecretIDFlags(depth int, m *models.Secret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveSecretSpecFlags(depth int, m *models.Secret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	specFlagName := fmt.Sprintf("%v.Spec", cmdPrefix)
	if cmd.Flags().Changed(specFlagName) {
		// info: complex object Spec SecretSpec is retrieved outside this Changed() block
	}
	specFlagValue := m.Spec
	if swag.IsZero(specFlagValue) {
		specFlagValue = &models.SecretSpec{}
	}

	err, specAdded := retrieveModelSecretSpecFlags(depth+1, specFlagValue, specFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || specAdded
	if specAdded {
		m.Spec = specFlagValue
	}

	return nil, retAdded
}

func retrieveSecretUpdatedAtFlags(depth int, m *models.Secret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveSecretVersionFlags(depth int, m *models.Secret, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.Version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {
		// info: complex object Version ObjectVersion is retrieved outside this Changed() block
	}
	versionFlagValue := m.Version
	if swag.IsZero(versionFlagValue) {
		versionFlagValue = &models.ObjectVersion{}
	}

	err, versionAdded := retrieveModelObjectVersionFlags(depth+1, versionFlagValue, versionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded
	if versionAdded {
		m.Version = versionFlagValue
	}

	return nil, retAdded
}
