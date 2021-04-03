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
func registerModelImageSummaryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageSummaryContainers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummaryCreated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummaryID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummaryLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummaryParentID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummaryRepoDigests(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummaryRepoTags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummarySharedSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummarySize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSummaryVirtualSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageSummaryContainers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	containersDescription := `Required. `

	var containersFlagName string
	if cmdPrefix == "" {
		containersFlagName = "Containers"
	} else {
		containersFlagName = fmt.Sprintf("%v.Containers", cmdPrefix)
	}

	var containersFlagDefault int64

	_ = cmd.PersistentFlags().Int64(containersFlagName, containersFlagDefault, containersDescription)

	return nil
}

func registerImageSummaryCreated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdDescription := `Required. `

	var createdFlagName string
	if cmdPrefix == "" {
		createdFlagName = "Created"
	} else {
		createdFlagName = fmt.Sprintf("%v.Created", cmdPrefix)
	}

	var createdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(createdFlagName, createdFlagDefault, createdDescription)

	return nil
}

func registerImageSummaryID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "Id"
	} else {
		idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerImageSummaryLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerImageSummaryParentID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	parentIdDescription := `Required. `

	var parentIdFlagName string
	if cmdPrefix == "" {
		parentIdFlagName = "ParentId"
	} else {
		parentIdFlagName = fmt.Sprintf("%v.ParentId", cmdPrefix)
	}

	var parentIdFlagDefault string

	_ = cmd.PersistentFlags().String(parentIdFlagName, parentIdFlagDefault, parentIdDescription)

	return nil
}

func registerImageSummaryRepoDigests(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: RepoDigests []string array type is not supported by go-swagger cli yet

	return nil
}

func registerImageSummaryRepoTags(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: RepoTags []string array type is not supported by go-swagger cli yet

	return nil
}

func registerImageSummarySharedSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sharedSizeDescription := `Required. `

	var sharedSizeFlagName string
	if cmdPrefix == "" {
		sharedSizeFlagName = "SharedSize"
	} else {
		sharedSizeFlagName = fmt.Sprintf("%v.SharedSize", cmdPrefix)
	}

	var sharedSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sharedSizeFlagName, sharedSizeFlagDefault, sharedSizeDescription)

	return nil
}

func registerImageSummarySize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := `Required. `

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "Size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
	}

	var sizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

func registerImageSummaryVirtualSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	virtualSizeDescription := `Required. `

	var virtualSizeFlagName string
	if cmdPrefix == "" {
		virtualSizeFlagName = "VirtualSize"
	} else {
		virtualSizeFlagName = fmt.Sprintf("%v.VirtualSize", cmdPrefix)
	}

	var virtualSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(virtualSizeFlagName, virtualSizeFlagDefault, virtualSizeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageSummaryFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, containersAdded := retrieveImageSummaryContainersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containersAdded

	err, createdAdded := retrieveImageSummaryCreatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAdded

	err, idAdded := retrieveImageSummaryIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, labelsAdded := retrieveImageSummaryLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, parentIdAdded := retrieveImageSummaryParentIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentIdAdded

	err, repoDigestsAdded := retrieveImageSummaryRepoDigestsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || repoDigestsAdded

	err, repoTagsAdded := retrieveImageSummaryRepoTagsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || repoTagsAdded

	err, sharedSizeAdded := retrieveImageSummarySharedSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sharedSizeAdded

	err, sizeAdded := retrieveImageSummarySizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, virtualSizeAdded := retrieveImageSummaryVirtualSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || virtualSizeAdded

	return nil, retAdded
}

func retrieveImageSummaryContainersFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	containersFlagName := fmt.Sprintf("%v.Containers", cmdPrefix)
	if cmd.Flags().Changed(containersFlagName) {

		var containersFlagName string
		if cmdPrefix == "" {
			containersFlagName = "Containers"
		} else {
			containersFlagName = fmt.Sprintf("%v.Containers", cmdPrefix)
		}

		containersFlagValue, err := cmd.Flags().GetInt64(containersFlagName)
		if err != nil {
			return err, false
		}
		m.Containers = containersFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageSummaryCreatedFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	createdFlagName := fmt.Sprintf("%v.Created", cmdPrefix)
	if cmd.Flags().Changed(createdFlagName) {

		var createdFlagName string
		if cmdPrefix == "" {
			createdFlagName = "Created"
		} else {
			createdFlagName = fmt.Sprintf("%v.Created", cmdPrefix)
		}

		createdFlagValue, err := cmd.Flags().GetInt64(createdFlagName)
		if err != nil {
			return err, false
		}
		m.Created = createdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageSummaryIDFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	idFlagName := fmt.Sprintf("%v.Id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "Id"
		} else {
			idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
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

func retrieveImageSummaryLabelsFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveImageSummaryParentIDFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	parentIdFlagName := fmt.Sprintf("%v.ParentId", cmdPrefix)
	if cmd.Flags().Changed(parentIdFlagName) {

		var parentIdFlagName string
		if cmdPrefix == "" {
			parentIdFlagName = "ParentId"
		} else {
			parentIdFlagName = fmt.Sprintf("%v.ParentId", cmdPrefix)
		}

		parentIdFlagValue, err := cmd.Flags().GetString(parentIdFlagName)
		if err != nil {
			return err, false
		}
		m.ParentID = parentIdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageSummaryRepoDigestsFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	repoDigestsFlagName := fmt.Sprintf("%v.RepoDigests", cmdPrefix)
	if cmd.Flags().Changed(repoDigestsFlagName) {
		// warning: RepoDigests array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveImageSummaryRepoTagsFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	repoTagsFlagName := fmt.Sprintf("%v.RepoTags", cmdPrefix)
	if cmd.Flags().Changed(repoTagsFlagName) {
		// warning: RepoTags array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveImageSummarySharedSizeFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	sharedSizeFlagName := fmt.Sprintf("%v.SharedSize", cmdPrefix)
	if cmd.Flags().Changed(sharedSizeFlagName) {

		var sharedSizeFlagName string
		if cmdPrefix == "" {
			sharedSizeFlagName = "SharedSize"
		} else {
			sharedSizeFlagName = fmt.Sprintf("%v.SharedSize", cmdPrefix)
		}

		sharedSizeFlagValue, err := cmd.Flags().GetInt64(sharedSizeFlagName)
		if err != nil {
			return err, false
		}
		m.SharedSize = sharedSizeFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageSummarySizeFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	sizeFlagName := fmt.Sprintf("%v.Size", cmdPrefix)
	if cmd.Flags().Changed(sizeFlagName) {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "Size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetInt64(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = sizeFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageSummaryVirtualSizeFlags(depth int, m *models.ImageSummary, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	virtualSizeFlagName := fmt.Sprintf("%v.VirtualSize", cmdPrefix)
	if cmd.Flags().Changed(virtualSizeFlagName) {

		var virtualSizeFlagName string
		if cmdPrefix == "" {
			virtualSizeFlagName = "VirtualSize"
		} else {
			virtualSizeFlagName = fmt.Sprintf("%v.VirtualSize", cmdPrefix)
		}

		virtualSizeFlagValue, err := cmd.Flags().GetInt64(virtualSizeFlagName)
		if err != nil {
			return err, false
		}
		m.VirtualSize = virtualSizeFlagValue

		retAdded = true
	}
	return nil, retAdded
}