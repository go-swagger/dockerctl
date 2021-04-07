// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for Image

// register flags to command
func registerModelImageFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageArchitecture(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageAuthor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageComment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageContainer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageContainerConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageCreated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageDockerVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageGraphDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageMetadata(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageOs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageOsVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRepoDigests(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRepoTags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRootFS(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageVirtualSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageArchitecture(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	architectureDescription := `Required. `

	var architectureFlagName string
	if cmdPrefix == "" {
		architectureFlagName = "Architecture"
	} else {
		architectureFlagName = fmt.Sprintf("%v.Architecture", cmdPrefix)
	}

	var architectureFlagDefault string

	_ = cmd.PersistentFlags().String(architectureFlagName, architectureFlagDefault, architectureDescription)

	return nil
}

func registerImageAuthor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	authorDescription := `Required. `

	var authorFlagName string
	if cmdPrefix == "" {
		authorFlagName = "Author"
	} else {
		authorFlagName = fmt.Sprintf("%v.Author", cmdPrefix)
	}

	var authorFlagDefault string

	_ = cmd.PersistentFlags().String(authorFlagName, authorFlagDefault, authorDescription)

	return nil
}

func registerImageComment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	commentDescription := `Required. `

	var commentFlagName string
	if cmdPrefix == "" {
		commentFlagName = "Comment"
	} else {
		commentFlagName = fmt.Sprintf("%v.Comment", cmdPrefix)
	}

	var commentFlagDefault string

	_ = cmd.PersistentFlags().String(commentFlagName, commentFlagDefault, commentDescription)

	return nil
}

func registerImageConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var configFlagName string
	if cmdPrefix == "" {
		configFlagName = "Config"
	} else {
		configFlagName = fmt.Sprintf("%v.Config", cmdPrefix)
	}

	if err := registerModelContainerConfigFlags(depth+1, configFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageContainer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	containerDescription := `Required. `

	var containerFlagName string
	if cmdPrefix == "" {
		containerFlagName = "Container"
	} else {
		containerFlagName = fmt.Sprintf("%v.Container", cmdPrefix)
	}

	var containerFlagDefault string

	_ = cmd.PersistentFlags().String(containerFlagName, containerFlagDefault, containerDescription)

	return nil
}

func registerImageContainerConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var containerConfigFlagName string
	if cmdPrefix == "" {
		containerConfigFlagName = "ContainerConfig"
	} else {
		containerConfigFlagName = fmt.Sprintf("%v.ContainerConfig", cmdPrefix)
	}

	if err := registerModelContainerConfigFlags(depth+1, containerConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageCreated(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

	var createdFlagDefault string

	_ = cmd.PersistentFlags().String(createdFlagName, createdFlagDefault, createdDescription)

	return nil
}

func registerImageDockerVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dockerVersionDescription := `Required. `

	var dockerVersionFlagName string
	if cmdPrefix == "" {
		dockerVersionFlagName = "DockerVersion"
	} else {
		dockerVersionFlagName = fmt.Sprintf("%v.DockerVersion", cmdPrefix)
	}

	var dockerVersionFlagDefault string

	_ = cmd.PersistentFlags().String(dockerVersionFlagName, dockerVersionFlagDefault, dockerVersionDescription)

	return nil
}

func registerImageGraphDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var graphDriverFlagName string
	if cmdPrefix == "" {
		graphDriverFlagName = "GraphDriver"
	} else {
		graphDriverFlagName = fmt.Sprintf("%v.GraphDriver", cmdPrefix)
	}

	if err := registerModelGraphDriverDataFlags(depth+1, graphDriverFlagName, cmd); err != nil {
		return err
	}

	if err := cmd.MarkPersistentFlagRequired(graphDriverFlagName); err != nil {
		return err
	}

	return nil
}

func registerImageID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerImageMetadata(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var metadataFlagName string
	if cmdPrefix == "" {
		metadataFlagName = "Metadata"
	} else {
		metadataFlagName = fmt.Sprintf("%v.Metadata", cmdPrefix)
	}

	if err := registerModelImageMetadataFlags(depth+1, metadataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageOs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	osDescription := `Required. `

	var osFlagName string
	if cmdPrefix == "" {
		osFlagName = "Os"
	} else {
		osFlagName = fmt.Sprintf("%v.Os", cmdPrefix)
	}

	var osFlagDefault string

	_ = cmd.PersistentFlags().String(osFlagName, osFlagDefault, osDescription)

	return nil
}

func registerImageOsVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	osVersionDescription := ``

	var osVersionFlagName string
	if cmdPrefix == "" {
		osVersionFlagName = "OsVersion"
	} else {
		osVersionFlagName = fmt.Sprintf("%v.OsVersion", cmdPrefix)
	}

	var osVersionFlagDefault string

	_ = cmd.PersistentFlags().String(osVersionFlagName, osVersionFlagDefault, osVersionDescription)

	return nil
}

func registerImageParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	parentDescription := `Required. `

	var parentFlagName string
	if cmdPrefix == "" {
		parentFlagName = "Parent"
	} else {
		parentFlagName = fmt.Sprintf("%v.Parent", cmdPrefix)
	}

	var parentFlagDefault string

	_ = cmd.PersistentFlags().String(parentFlagName, parentFlagDefault, parentDescription)

	return nil
}

func registerImageRepoDigests(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: RepoDigests []string array type is not supported by go-swagger cli yet

	return nil
}

func registerImageRepoTags(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: RepoTags []string array type is not supported by go-swagger cli yet

	return nil
}

func registerImageRootFS(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var rootFSFlagName string
	if cmdPrefix == "" {
		rootFSFlagName = "RootFS"
	} else {
		rootFSFlagName = fmt.Sprintf("%v.RootFS", cmdPrefix)
	}

	if err := registerModelImageRootFSFlags(depth+1, rootFSFlagName, cmd); err != nil {
		return err
	}

	if err := cmd.MarkPersistentFlagRequired(rootFSFlagName); err != nil {
		return err
	}

	return nil
}

func registerImageSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerImageVirtualSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelImageFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, architectureAdded := retrieveImageArchitectureFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || architectureAdded

	err, authorAdded := retrieveImageAuthorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || authorAdded

	err, commentAdded := retrieveImageCommentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || commentAdded

	err, configAdded := retrieveImageConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || configAdded

	err, containerAdded := retrieveImageContainerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containerAdded

	err, containerConfigAdded := retrieveImageContainerConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containerConfigAdded

	err, createdAdded := retrieveImageCreatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAdded

	err, dockerVersionAdded := retrieveImageDockerVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dockerVersionAdded

	err, graphDriverAdded := retrieveImageGraphDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || graphDriverAdded

	err, idAdded := retrieveImageIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, metadataAdded := retrieveImageMetadataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metadataAdded

	err, osAdded := retrieveImageOsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || osAdded

	err, osVersionAdded := retrieveImageOsVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || osVersionAdded

	err, parentAdded := retrieveImageParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	err, repoDigestsAdded := retrieveImageRepoDigestsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || repoDigestsAdded

	err, repoTagsAdded := retrieveImageRepoTagsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || repoTagsAdded

	err, rootFSAdded := retrieveImageRootFSFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rootFSAdded

	err, sizeAdded := retrieveImageSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, virtualSizeAdded := retrieveImageVirtualSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || virtualSizeAdded

	return nil, retAdded
}

func retrieveImageArchitectureFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	architectureFlagName := fmt.Sprintf("%v.Architecture", cmdPrefix)
	if cmd.Flags().Changed(architectureFlagName) {

		var architectureFlagName string
		if cmdPrefix == "" {
			architectureFlagName = "Architecture"
		} else {
			architectureFlagName = fmt.Sprintf("%v.Architecture", cmdPrefix)
		}

		architectureFlagValue, err := cmd.Flags().GetString(architectureFlagName)
		if err != nil {
			return err, false
		}
		m.Architecture = architectureFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageAuthorFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	authorFlagName := fmt.Sprintf("%v.Author", cmdPrefix)
	if cmd.Flags().Changed(authorFlagName) {

		var authorFlagName string
		if cmdPrefix == "" {
			authorFlagName = "Author"
		} else {
			authorFlagName = fmt.Sprintf("%v.Author", cmdPrefix)
		}

		authorFlagValue, err := cmd.Flags().GetString(authorFlagName)
		if err != nil {
			return err, false
		}
		m.Author = authorFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageCommentFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	commentFlagName := fmt.Sprintf("%v.Comment", cmdPrefix)
	if cmd.Flags().Changed(commentFlagName) {

		var commentFlagName string
		if cmdPrefix == "" {
			commentFlagName = "Comment"
		} else {
			commentFlagName = fmt.Sprintf("%v.Comment", cmdPrefix)
		}

		commentFlagValue, err := cmd.Flags().GetString(commentFlagName)
		if err != nil {
			return err, false
		}
		m.Comment = commentFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageConfigFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	configFlagName := fmt.Sprintf("%v.Config", cmdPrefix)
	if cmd.Flags().Changed(configFlagName) {

		configFlagValue := &models.ContainerConfig{}
		err, added := retrieveModelContainerConfigFlags(depth+1, configFlagValue, configFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.Config = configFlagValue
		}
	}
	return nil, retAdded
}

func retrieveImageContainerFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	containerFlagName := fmt.Sprintf("%v.Container", cmdPrefix)
	if cmd.Flags().Changed(containerFlagName) {

		var containerFlagName string
		if cmdPrefix == "" {
			containerFlagName = "Container"
		} else {
			containerFlagName = fmt.Sprintf("%v.Container", cmdPrefix)
		}

		containerFlagValue, err := cmd.Flags().GetString(containerFlagName)
		if err != nil {
			return err, false
		}
		m.Container = containerFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageContainerConfigFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	containerConfigFlagName := fmt.Sprintf("%v.ContainerConfig", cmdPrefix)
	if cmd.Flags().Changed(containerConfigFlagName) {

		containerConfigFlagValue := &models.ContainerConfig{}
		err, added := retrieveModelContainerConfigFlags(depth+1, containerConfigFlagValue, containerConfigFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.ContainerConfig = containerConfigFlagValue
		}
	}
	return nil, retAdded
}

func retrieveImageCreatedFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		createdFlagValue, err := cmd.Flags().GetString(createdFlagName)
		if err != nil {
			return err, false
		}
		m.Created = createdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageDockerVersionFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	dockerVersionFlagName := fmt.Sprintf("%v.DockerVersion", cmdPrefix)
	if cmd.Flags().Changed(dockerVersionFlagName) {

		var dockerVersionFlagName string
		if cmdPrefix == "" {
			dockerVersionFlagName = "DockerVersion"
		} else {
			dockerVersionFlagName = fmt.Sprintf("%v.DockerVersion", cmdPrefix)
		}

		dockerVersionFlagValue, err := cmd.Flags().GetString(dockerVersionFlagName)
		if err != nil {
			return err, false
		}
		m.DockerVersion = dockerVersionFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageGraphDriverFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	graphDriverFlagName := fmt.Sprintf("%v.GraphDriver", cmdPrefix)
	if cmd.Flags().Changed(graphDriverFlagName) {

		graphDriverFlagValue := &models.GraphDriverData{}
		err, added := retrieveModelGraphDriverDataFlags(depth+1, graphDriverFlagValue, graphDriverFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.GraphDriver = graphDriverFlagValue
		}
	}
	return nil, retAdded
}

func retrieveImageIDFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveImageMetadataFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	metadataFlagName := fmt.Sprintf("%v.Metadata", cmdPrefix)
	if cmd.Flags().Changed(metadataFlagName) {

		metadataFlagValue := &models.ImageMetadata{}
		err, added := retrieveModelImageMetadataFlags(depth+1, metadataFlagValue, metadataFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.Metadata = metadataFlagValue
		}
	}
	return nil, retAdded
}

func retrieveImageOsFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	osFlagName := fmt.Sprintf("%v.Os", cmdPrefix)
	if cmd.Flags().Changed(osFlagName) {

		var osFlagName string
		if cmdPrefix == "" {
			osFlagName = "Os"
		} else {
			osFlagName = fmt.Sprintf("%v.Os", cmdPrefix)
		}

		osFlagValue, err := cmd.Flags().GetString(osFlagName)
		if err != nil {
			return err, false
		}
		m.Os = osFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageOsVersionFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	osVersionFlagName := fmt.Sprintf("%v.OsVersion", cmdPrefix)
	if cmd.Flags().Changed(osVersionFlagName) {

		var osVersionFlagName string
		if cmdPrefix == "" {
			osVersionFlagName = "OsVersion"
		} else {
			osVersionFlagName = fmt.Sprintf("%v.OsVersion", cmdPrefix)
		}

		osVersionFlagValue, err := cmd.Flags().GetString(osVersionFlagName)
		if err != nil {
			return err, false
		}
		m.OsVersion = osVersionFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageParentFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	parentFlagName := fmt.Sprintf("%v.Parent", cmdPrefix)
	if cmd.Flags().Changed(parentFlagName) {

		var parentFlagName string
		if cmdPrefix == "" {
			parentFlagName = "Parent"
		} else {
			parentFlagName = fmt.Sprintf("%v.Parent", cmdPrefix)
		}

		parentFlagValue, err := cmd.Flags().GetString(parentFlagName)
		if err != nil {
			return err, false
		}
		m.Parent = parentFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageRepoDigestsFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveImageRepoTagsFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveImageRootFSFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	rootFSFlagName := fmt.Sprintf("%v.RootFS", cmdPrefix)
	if cmd.Flags().Changed(rootFSFlagName) {

		rootFSFlagValue := &models.ImageRootFS{}
		err, added := retrieveModelImageRootFSFlags(depth+1, rootFSFlagValue, rootFSFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.RootFS = rootFSFlagValue
		}
	}
	return nil, retAdded
}

func retrieveImageSizeFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveImageVirtualSizeFlags(depth int, m *models.Image, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// Extra schema cli for ImageMetadata

// register flags to command
func registerModelImageMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageMetadataLastTagTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageMetadataLastTagTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastTagTimeDescription := ``

	var lastTagTimeFlagName string
	if cmdPrefix == "" {
		lastTagTimeFlagName = "LastTagTime"
	} else {
		lastTagTimeFlagName = fmt.Sprintf("%v.LastTagTime", cmdPrefix)
	}

	var lastTagTimeFlagDefault string

	_ = cmd.PersistentFlags().String(lastTagTimeFlagName, lastTagTimeFlagDefault, lastTagTimeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImageMetadataFlags(depth int, m *models.ImageMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, lastTagTimeAdded := retrieveImageMetadataLastTagTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastTagTimeAdded

	return nil, retAdded
}

func retrieveImageMetadataLastTagTimeFlags(depth int, m *models.ImageMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	lastTagTimeFlagName := fmt.Sprintf("%v.LastTagTime", cmdPrefix)
	if cmd.Flags().Changed(lastTagTimeFlagName) {

		var lastTagTimeFlagName string
		if cmdPrefix == "" {
			lastTagTimeFlagName = "LastTagTime"
		} else {
			lastTagTimeFlagName = fmt.Sprintf("%v.LastTagTime", cmdPrefix)
		}

		lastTagTimeFlagValue, err := cmd.Flags().GetString(lastTagTimeFlagName)
		if err != nil {
			return err, false
		}
		m.LastTagTime = lastTagTimeFlagValue

		retAdded = true
	}
	return nil, retAdded
}

// Extra schema cli for ImageRootFS

// register flags to command
func registerModelImageRootFSFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImageRootFSBaseLayer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRootFSLayers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImageRootFSType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImageRootFSBaseLayer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	baseLayerDescription := ``

	var baseLayerFlagName string
	if cmdPrefix == "" {
		baseLayerFlagName = "BaseLayer"
	} else {
		baseLayerFlagName = fmt.Sprintf("%v.BaseLayer", cmdPrefix)
	}

	var baseLayerFlagDefault string

	_ = cmd.PersistentFlags().String(baseLayerFlagName, baseLayerFlagDefault, baseLayerDescription)

	return nil
}

func registerImageRootFSLayers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Layers []string array type is not supported by go-swagger cli yet

	return nil
}

func registerImageRootFSType(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelImageRootFSFlags(depth int, m *models.ImageRootFS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, baseLayerAdded := retrieveImageRootFSBaseLayerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || baseLayerAdded

	err, layersAdded := retrieveImageRootFSLayersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || layersAdded

	err, typeAdded := retrieveImageRootFSTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveImageRootFSBaseLayerFlags(depth int, m *models.ImageRootFS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	baseLayerFlagName := fmt.Sprintf("%v.BaseLayer", cmdPrefix)
	if cmd.Flags().Changed(baseLayerFlagName) {

		var baseLayerFlagName string
		if cmdPrefix == "" {
			baseLayerFlagName = "BaseLayer"
		} else {
			baseLayerFlagName = fmt.Sprintf("%v.BaseLayer", cmdPrefix)
		}

		baseLayerFlagValue, err := cmd.Flags().GetString(baseLayerFlagName)
		if err != nil {
			return err, false
		}
		m.BaseLayer = baseLayerFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveImageRootFSLayersFlags(depth int, m *models.ImageRootFS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	layersFlagName := fmt.Sprintf("%v.Layers", cmdPrefix)
	if cmd.Flags().Changed(layersFlagName) {
		// warning: Layers array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveImageRootFSTypeFlags(depth int, m *models.ImageRootFS, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
