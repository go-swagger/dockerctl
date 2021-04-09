// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for Volume

// register flags to command
func registerModelVolumeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVolumeCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeMountpoint(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeScope(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeUsageData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVolumeCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdAtDescription := `Date/Time the volume was created.`

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

func registerVolumeDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	driverDescription := `Required. Name of the volume driver used by the volume.`

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

func registerVolumeLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerVolumeMountpoint(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	mountpointDescription := `Required. Mount path of the volume on the host.`

	var mountpointFlagName string
	if cmdPrefix == "" {
		mountpointFlagName = "Mountpoint"
	} else {
		mountpointFlagName = fmt.Sprintf("%v.Mountpoint", cmdPrefix)
	}

	var mountpointFlagDefault string

	_ = cmd.PersistentFlags().String(mountpointFlagName, mountpointFlagDefault, mountpointDescription)

	return nil
}

func registerVolumeName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. Name of the volume.`

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

func registerVolumeOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Options map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerVolumeScope(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	scopeDescription := `Required. The level at which the volume exists. Either ` + "`" + `global` + "`" + ` for cluster-wide, or ` + "`" + `local` + "`" + ` for machine level.`

	var scopeFlagName string
	if cmdPrefix == "" {
		scopeFlagName = "Scope"
	} else {
		scopeFlagName = fmt.Sprintf("%v.Scope", cmdPrefix)
	}

	var scopeFlagDefault string = "local"

	_ = cmd.PersistentFlags().String(scopeFlagName, scopeFlagDefault, scopeDescription)

	return nil
}

func registerVolumeStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Status map[string]interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerVolumeUsageData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var usageDataFlagName string
	if cmdPrefix == "" {
		usageDataFlagName = "UsageData"
	} else {
		usageDataFlagName = fmt.Sprintf("%v.UsageData", cmdPrefix)
	}

	if err := registerModelVolumeUsageDataFlags(depth+1, usageDataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVolumeFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, createdAtAdded := retrieveVolumeCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, driverAdded := retrieveVolumeDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverAdded

	err, labelsAdded := retrieveVolumeLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, mountpointAdded := retrieveVolumeMountpointFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mountpointAdded

	err, nameAdded := retrieveVolumeNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, optionsAdded := retrieveVolumeOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || optionsAdded

	err, scopeAdded := retrieveVolumeScopeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeAdded

	err, statusAdded := retrieveVolumeStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, usageDataAdded := retrieveVolumeUsageDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usageDataAdded

	return nil, retAdded
}

func retrieveVolumeCreatedAtFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVolumeDriverFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVolumeLabelsFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVolumeMountpointFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mountpointFlagName := fmt.Sprintf("%v.Mountpoint", cmdPrefix)
	if cmd.Flags().Changed(mountpointFlagName) {

		var mountpointFlagName string
		if cmdPrefix == "" {
			mountpointFlagName = "Mountpoint"
		} else {
			mountpointFlagName = fmt.Sprintf("%v.Mountpoint", cmdPrefix)
		}

		mountpointFlagValue, err := cmd.Flags().GetString(mountpointFlagName)
		if err != nil {
			return err, false
		}
		m.Mountpoint = mountpointFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVolumeNameFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVolumeOptionsFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	optionsFlagName := fmt.Sprintf("%v.Options", cmdPrefix)
	if cmd.Flags().Changed(optionsFlagName) {
		// warning: Options map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveVolumeScopeFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scopeFlagName := fmt.Sprintf("%v.Scope", cmdPrefix)
	if cmd.Flags().Changed(scopeFlagName) {

		var scopeFlagName string
		if cmdPrefix == "" {
			scopeFlagName = "Scope"
		} else {
			scopeFlagName = fmt.Sprintf("%v.Scope", cmdPrefix)
		}

		scopeFlagValue, err := cmd.Flags().GetString(scopeFlagName)
		if err != nil {
			return err, false
		}
		m.Scope = scopeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVolumeStatusFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.Status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {
		// warning: Status map type map[string]interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveVolumeUsageDataFlags(depth int, m *models.Volume, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usageDataFlagName := fmt.Sprintf("%v.UsageData", cmdPrefix)
	if cmd.Flags().Changed(usageDataFlagName) {

		usageDataFlagValue := models.VolumeUsageData{}
		err, added := retrieveModelVolumeUsageDataFlags(depth+1, &usageDataFlagValue, usageDataFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.UsageData = &usageDataFlagValue
		}
	}

	return nil, retAdded
}

// Extra schema cli for VolumeUsageData

// register flags to command
func registerModelVolumeUsageDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVolumeUsageDataRefCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeUsageDataSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVolumeUsageDataRefCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	refCountDescription := `Required. The number of containers referencing this volume. This field
is set to ` + "`" + `-1` + "`" + ` if the reference-count is not available.
`

	var refCountFlagName string
	if cmdPrefix == "" {
		refCountFlagName = "RefCount"
	} else {
		refCountFlagName = fmt.Sprintf("%v.RefCount", cmdPrefix)
	}

	var refCountFlagDefault int64 = -1

	_ = cmd.PersistentFlags().Int64(refCountFlagName, refCountFlagDefault, refCountDescription)

	return nil
}

func registerVolumeUsageDataSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := `Required. Amount of disk space used by the volume (in bytes). This information
is only available for volumes created with the ` + "`" + `"local"` + "`" + ` volume
driver. For volumes created with other volume drivers, this field
is set to ` + "`" + `-1` + "`" + ` ("not available")
`

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "Size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
	}

	var sizeFlagDefault int64 = -1

	_ = cmd.PersistentFlags().Int64(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVolumeUsageDataFlags(depth int, m *models.VolumeUsageData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, refCountAdded := retrieveVolumeUsageDataRefCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || refCountAdded

	err, sizeAdded := retrieveVolumeUsageDataSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	return nil, retAdded
}

func retrieveVolumeUsageDataRefCountFlags(depth int, m *models.VolumeUsageData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	refCountFlagName := fmt.Sprintf("%v.RefCount", cmdPrefix)
	if cmd.Flags().Changed(refCountFlagName) {

		var refCountFlagName string
		if cmdPrefix == "" {
			refCountFlagName = "RefCount"
		} else {
			refCountFlagName = fmt.Sprintf("%v.RefCount", cmdPrefix)
		}

		refCountFlagValue, err := cmd.Flags().GetInt64(refCountFlagName)
		if err != nil {
			return err, false
		}
		m.RefCount = refCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVolumeUsageDataSizeFlags(depth int, m *models.VolumeUsageData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
