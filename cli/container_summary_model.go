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

// Schema cli for []*ContainerSummaryItems0

// Name: [ContainerSummary], Type:[[]*ContainerSummaryItems0], register and retrieve functions are not rendered by go-swagger cli

// Extra schema cli for ContainerSummaryItems0

// register flags to command
func registerModelContainerSummaryItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerSummaryItems0Command(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0Created(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0HostConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0ID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0Image(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0ImageID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0Labels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0Mounts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0Names(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0NetworkSettings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0Ports(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0SizeRootFs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0SizeRw(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0State(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerSummaryItems0Status(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerSummaryItems0Command(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	commandDescription := `Command to run when starting the container`

	var commandFlagName string
	if cmdPrefix == "" {
		commandFlagName = "Command"
	} else {
		commandFlagName = fmt.Sprintf("%v.Command", cmdPrefix)
	}

	var commandFlagDefault string

	_ = cmd.PersistentFlags().String(commandFlagName, commandFlagDefault, commandDescription)

	return nil
}

func registerContainerSummaryItems0Created(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdDescription := `When the container was created`

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

func registerContainerSummaryItems0HostConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var hostConfigFlagName string
	if cmdPrefix == "" {
		hostConfigFlagName = "HostConfig"
	} else {
		hostConfigFlagName = fmt.Sprintf("%v.HostConfig", cmdPrefix)
	}

	if err := registerModelContainerSummaryItems0HostConfigFlags(depth+1, hostConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerSummaryItems0ID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of this container`

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

func registerContainerSummaryItems0Image(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	imageDescription := `The name of the image used when creating this container`

	var imageFlagName string
	if cmdPrefix == "" {
		imageFlagName = "Image"
	} else {
		imageFlagName = fmt.Sprintf("%v.Image", cmdPrefix)
	}

	var imageFlagDefault string

	_ = cmd.PersistentFlags().String(imageFlagName, imageFlagDefault, imageDescription)

	return nil
}

func registerContainerSummaryItems0ImageID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	imageIdDescription := `The ID of the image that this container was created from`

	var imageIdFlagName string
	if cmdPrefix == "" {
		imageIdFlagName = "ImageID"
	} else {
		imageIdFlagName = fmt.Sprintf("%v.ImageID", cmdPrefix)
	}

	var imageIdFlagDefault string

	_ = cmd.PersistentFlags().String(imageIdFlagName, imageIdFlagDefault, imageIdDescription)

	return nil
}

func registerContainerSummaryItems0Labels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerContainerSummaryItems0Mounts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Mounts []*Mount array type is not supported by go-swagger cli yet

	return nil
}

func registerContainerSummaryItems0Names(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Names []string array type is not supported by go-swagger cli yet

	return nil
}

func registerContainerSummaryItems0NetworkSettings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var networkSettingsFlagName string
	if cmdPrefix == "" {
		networkSettingsFlagName = "NetworkSettings"
	} else {
		networkSettingsFlagName = fmt.Sprintf("%v.NetworkSettings", cmdPrefix)
	}

	if err := registerModelContainerSummaryItems0NetworkSettingsFlags(depth+1, networkSettingsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerSummaryItems0Ports(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Ports []*Port array type is not supported by go-swagger cli yet

	return nil
}

func registerContainerSummaryItems0SizeRootFs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeRootFsDescription := `The total size of all the files in this container`

	var sizeRootFsFlagName string
	if cmdPrefix == "" {
		sizeRootFsFlagName = "SizeRootFs"
	} else {
		sizeRootFsFlagName = fmt.Sprintf("%v.SizeRootFs", cmdPrefix)
	}

	var sizeRootFsFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeRootFsFlagName, sizeRootFsFlagDefault, sizeRootFsDescription)

	return nil
}

func registerContainerSummaryItems0SizeRw(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeRwDescription := `The size of files that have been created or changed by this container`

	var sizeRwFlagName string
	if cmdPrefix == "" {
		sizeRwFlagName = "SizeRw"
	} else {
		sizeRwFlagName = fmt.Sprintf("%v.SizeRw", cmdPrefix)
	}

	var sizeRwFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeRwFlagName, sizeRwFlagDefault, sizeRwDescription)

	return nil
}

func registerContainerSummaryItems0State(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	stateDescription := `The state of this container (e.g. ` + "`" + `Exited` + "`" + `)`

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "State"
	} else {
		stateFlagName = fmt.Sprintf("%v.State", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	return nil
}

func registerContainerSummaryItems0Status(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := `Additional human-readable status of this container (e.g. ` + "`" + `Exit 0` + "`" + `)`

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "Status"
	} else {
		statusFlagName = fmt.Sprintf("%v.Status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerSummaryItems0Flags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, commandAdded := retrieveContainerSummaryItems0CommandFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || commandAdded

	err, createdAdded := retrieveContainerSummaryItems0CreatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAdded

	err, hostConfigAdded := retrieveContainerSummaryItems0HostConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostConfigAdded

	err, idAdded := retrieveContainerSummaryItems0IDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, imageAdded := retrieveContainerSummaryItems0ImageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imageAdded

	err, imageIdAdded := retrieveContainerSummaryItems0ImageIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imageIdAdded

	err, labelsAdded := retrieveContainerSummaryItems0LabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, mountsAdded := retrieveContainerSummaryItems0MountsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mountsAdded

	err, namesAdded := retrieveContainerSummaryItems0NamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || namesAdded

	err, networkSettingsAdded := retrieveContainerSummaryItems0NetworkSettingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkSettingsAdded

	err, portsAdded := retrieveContainerSummaryItems0PortsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || portsAdded

	err, sizeRootFsAdded := retrieveContainerSummaryItems0SizeRootFsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeRootFsAdded

	err, sizeRwAdded := retrieveContainerSummaryItems0SizeRwFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeRwAdded

	err, stateAdded := retrieveContainerSummaryItems0StateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stateAdded

	err, statusAdded := retrieveContainerSummaryItems0StatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrieveContainerSummaryItems0CommandFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	commandFlagName := fmt.Sprintf("%v.Command", cmdPrefix)
	if cmd.Flags().Changed(commandFlagName) {

		var commandFlagName string
		if cmdPrefix == "" {
			commandFlagName = "Command"
		} else {
			commandFlagName = fmt.Sprintf("%v.Command", cmdPrefix)
		}

		commandFlagValue, err := cmd.Flags().GetString(commandFlagName)
		if err != nil {
			return err, false
		}
		m.Command = commandFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0CreatedFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveContainerSummaryItems0HostConfigFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hostConfigFlagName := fmt.Sprintf("%v.HostConfig", cmdPrefix)
	if cmd.Flags().Changed(hostConfigFlagName) {
		// info: complex object HostConfig ContainerSummaryItems0HostConfig is retrieved outside this Changed() block
	}
	hostConfigFlagValue := m.HostConfig
	if swag.IsZero(hostConfigFlagValue) {
		hostConfigFlagValue = &models.ContainerSummaryItems0HostConfig{}
	}

	err, hostConfigAdded := retrieveModelContainerSummaryItems0HostConfigFlags(depth+1, hostConfigFlagValue, hostConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostConfigAdded
	if hostConfigAdded {
		m.HostConfig = hostConfigFlagValue
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0IDFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveContainerSummaryItems0ImageFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imageFlagName := fmt.Sprintf("%v.Image", cmdPrefix)
	if cmd.Flags().Changed(imageFlagName) {

		var imageFlagName string
		if cmdPrefix == "" {
			imageFlagName = "Image"
		} else {
			imageFlagName = fmt.Sprintf("%v.Image", cmdPrefix)
		}

		imageFlagValue, err := cmd.Flags().GetString(imageFlagName)
		if err != nil {
			return err, false
		}
		m.Image = imageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0ImageIDFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imageIdFlagName := fmt.Sprintf("%v.ImageID", cmdPrefix)
	if cmd.Flags().Changed(imageIdFlagName) {

		var imageIdFlagName string
		if cmdPrefix == "" {
			imageIdFlagName = "ImageID"
		} else {
			imageIdFlagName = fmt.Sprintf("%v.ImageID", cmdPrefix)
		}

		imageIdFlagValue, err := cmd.Flags().GetString(imageIdFlagName)
		if err != nil {
			return err, false
		}
		m.ImageID = imageIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0LabelsFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveContainerSummaryItems0MountsFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mountsFlagName := fmt.Sprintf("%v.Mounts", cmdPrefix)
	if cmd.Flags().Changed(mountsFlagName) {
		// warning: Mounts array type []*Mount is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0NamesFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	namesFlagName := fmt.Sprintf("%v.Names", cmdPrefix)
	if cmd.Flags().Changed(namesFlagName) {
		// warning: Names array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0NetworkSettingsFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkSettingsFlagName := fmt.Sprintf("%v.NetworkSettings", cmdPrefix)
	if cmd.Flags().Changed(networkSettingsFlagName) {
		// info: complex object NetworkSettings ContainerSummaryItems0NetworkSettings is retrieved outside this Changed() block
	}
	networkSettingsFlagValue := m.NetworkSettings
	if swag.IsZero(networkSettingsFlagValue) {
		networkSettingsFlagValue = &models.ContainerSummaryItems0NetworkSettings{}
	}

	err, networkSettingsAdded := retrieveModelContainerSummaryItems0NetworkSettingsFlags(depth+1, networkSettingsFlagValue, networkSettingsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkSettingsAdded
	if networkSettingsAdded {
		m.NetworkSettings = networkSettingsFlagValue
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0PortsFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	portsFlagName := fmt.Sprintf("%v.Ports", cmdPrefix)
	if cmd.Flags().Changed(portsFlagName) {
		// warning: Ports array type []*Port is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0SizeRootFsFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeRootFsFlagName := fmt.Sprintf("%v.SizeRootFs", cmdPrefix)
	if cmd.Flags().Changed(sizeRootFsFlagName) {

		var sizeRootFsFlagName string
		if cmdPrefix == "" {
			sizeRootFsFlagName = "SizeRootFs"
		} else {
			sizeRootFsFlagName = fmt.Sprintf("%v.SizeRootFs", cmdPrefix)
		}

		sizeRootFsFlagValue, err := cmd.Flags().GetInt64(sizeRootFsFlagName)
		if err != nil {
			return err, false
		}
		m.SizeRootFs = sizeRootFsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0SizeRwFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeRwFlagName := fmt.Sprintf("%v.SizeRw", cmdPrefix)
	if cmd.Flags().Changed(sizeRwFlagName) {

		var sizeRwFlagName string
		if cmdPrefix == "" {
			sizeRwFlagName = "SizeRw"
		} else {
			sizeRwFlagName = fmt.Sprintf("%v.SizeRw", cmdPrefix)
		}

		sizeRwFlagValue, err := cmd.Flags().GetInt64(sizeRwFlagName)
		if err != nil {
			return err, false
		}
		m.SizeRw = sizeRwFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0StateFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stateFlagName := fmt.Sprintf("%v.State", cmdPrefix)
	if cmd.Flags().Changed(stateFlagName) {

		var stateFlagName string
		if cmdPrefix == "" {
			stateFlagName = "State"
		} else {
			stateFlagName = fmt.Sprintf("%v.State", cmdPrefix)
		}

		stateFlagValue, err := cmd.Flags().GetString(stateFlagName)
		if err != nil {
			return err, false
		}
		m.State = stateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveContainerSummaryItems0StatusFlags(depth int, m *models.ContainerSummaryItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.Status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "Status"
		} else {
			statusFlagName = fmt.Sprintf("%v.Status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ContainerSummaryItems0HostConfig

// register flags to command
func registerModelContainerSummaryItems0HostConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerSummaryItems0HostConfigNetworkMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerSummaryItems0HostConfigNetworkMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	networkModeDescription := ``

	var networkModeFlagName string
	if cmdPrefix == "" {
		networkModeFlagName = "NetworkMode"
	} else {
		networkModeFlagName = fmt.Sprintf("%v.NetworkMode", cmdPrefix)
	}

	var networkModeFlagDefault string

	_ = cmd.PersistentFlags().String(networkModeFlagName, networkModeFlagDefault, networkModeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerSummaryItems0HostConfigFlags(depth int, m *models.ContainerSummaryItems0HostConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, networkModeAdded := retrieveContainerSummaryItems0HostConfigNetworkModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkModeAdded

	return nil, retAdded
}

func retrieveContainerSummaryItems0HostConfigNetworkModeFlags(depth int, m *models.ContainerSummaryItems0HostConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkModeFlagName := fmt.Sprintf("%v.NetworkMode", cmdPrefix)
	if cmd.Flags().Changed(networkModeFlagName) {

		var networkModeFlagName string
		if cmdPrefix == "" {
			networkModeFlagName = "NetworkMode"
		} else {
			networkModeFlagName = fmt.Sprintf("%v.NetworkMode", cmdPrefix)
		}

		networkModeFlagValue, err := cmd.Flags().GetString(networkModeFlagName)
		if err != nil {
			return err, false
		}
		m.NetworkMode = networkModeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for ContainerSummaryItems0NetworkSettings

// register flags to command
func registerModelContainerSummaryItems0NetworkSettingsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerSummaryItems0NetworkSettingsNetworks(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerSummaryItems0NetworkSettingsNetworks(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Networks map[string]EndpointSettings map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerSummaryItems0NetworkSettingsFlags(depth int, m *models.ContainerSummaryItems0NetworkSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, networksAdded := retrieveContainerSummaryItems0NetworkSettingsNetworksFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networksAdded

	return nil, retAdded
}

func retrieveContainerSummaryItems0NetworkSettingsNetworksFlags(depth int, m *models.ContainerSummaryItems0NetworkSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networksFlagName := fmt.Sprintf("%v.Networks", cmdPrefix)
	if cmd.Flags().Changed(networksFlagName) {
		// warning: Networks map type map[string]EndpointSettings is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
