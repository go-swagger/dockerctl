// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/go-swagger/dockerctl/models"

	"github.com/spf13/cobra"
)

// Schema cli for Mount

// register flags to command
func registerModelMountFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMountBindOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountConsistency(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountReadOnly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountSource(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountTarget(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountTmpfsOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountVolumeOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountBindOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var bindOptionsFlagName string
	if cmdPrefix == "" {
		bindOptionsFlagName = "BindOptions"
	} else {
		bindOptionsFlagName = fmt.Sprintf("%v.BindOptions", cmdPrefix)
	}

	if err := registerModelMountBindOptionsFlags(depth+1, bindOptionsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountConsistency(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	consistencyDescription := `The consistency requirement for the mount: ` + "`" + `default` + "`" + `, ` + "`" + `consistent` + "`" + `, ` + "`" + `cached` + "`" + `, or ` + "`" + `delegated` + "`" + `.`

	var consistencyFlagName string
	if cmdPrefix == "" {
		consistencyFlagName = "Consistency"
	} else {
		consistencyFlagName = fmt.Sprintf("%v.Consistency", cmdPrefix)
	}

	var consistencyFlagDefault string

	_ = cmd.PersistentFlags().String(consistencyFlagName, consistencyFlagDefault, consistencyDescription)

	return nil
}

func registerMountReadOnly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	readOnlyDescription := `Whether the mount should be read-only.`

	var readOnlyFlagName string
	if cmdPrefix == "" {
		readOnlyFlagName = "ReadOnly"
	} else {
		readOnlyFlagName = fmt.Sprintf("%v.ReadOnly", cmdPrefix)
	}

	var readOnlyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(readOnlyFlagName, readOnlyFlagDefault, readOnlyDescription)

	return nil
}

func registerMountSource(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sourceDescription := `Mount source (e.g. a volume name, a host path).`

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

func registerMountTarget(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	targetDescription := `Container path.`

	var targetFlagName string
	if cmdPrefix == "" {
		targetFlagName = "Target"
	} else {
		targetFlagName = fmt.Sprintf("%v.Target", cmdPrefix)
	}

	var targetFlagDefault string

	_ = cmd.PersistentFlags().String(targetFlagName, targetFlagDefault, targetDescription)

	return nil
}

func registerMountTmpfsOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var tmpfsOptionsFlagName string
	if cmdPrefix == "" {
		tmpfsOptionsFlagName = "TmpfsOptions"
	} else {
		tmpfsOptionsFlagName = fmt.Sprintf("%v.TmpfsOptions", cmdPrefix)
	}

	if err := registerModelMountTmpfsOptionsFlags(depth+1, tmpfsOptionsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["bind","volume","tmpfs","npipe"]. The mount type. Available types:

- ` + "`" + `bind` + "`" + ` Mounts a file or directory from the host into the container. Must exist prior to creating the container.
- ` + "`" + `volume` + "`" + ` Creates a volume with the given name and options (or uses a pre-existing volume with the same name and options). These are **not** removed when the container is removed.
- ` + "`" + `tmpfs` + "`" + ` Create a tmpfs with the given options. The mount source cannot be specified for tmpfs.
- ` + "`" + `npipe` + "`" + ` Mounts a named pipe from the host into the container. Must exist prior to creating the container.
`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "Type"
	} else {
		typeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["bind","volume","tmpfs","npipe"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerMountVolumeOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var volumeOptionsFlagName string
	if cmdPrefix == "" {
		volumeOptionsFlagName = "VolumeOptions"
	} else {
		volumeOptionsFlagName = fmt.Sprintf("%v.VolumeOptions", cmdPrefix)
	}

	if err := registerModelMountVolumeOptionsFlags(depth+1, volumeOptionsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMountFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, bindOptionsAdded := retrieveMountBindOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bindOptionsAdded

	err, consistencyAdded := retrieveMountConsistencyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || consistencyAdded

	err, readOnlyAdded := retrieveMountReadOnlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || readOnlyAdded

	err, sourceAdded := retrieveMountSourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceAdded

	err, targetAdded := retrieveMountTargetFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetAdded

	err, tmpfsOptionsAdded := retrieveMountTmpfsOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tmpfsOptionsAdded

	err, typeAdded := retrieveMountTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, volumeOptionsAdded := retrieveMountVolumeOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || volumeOptionsAdded

	return nil, retAdded
}

func retrieveMountBindOptionsFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bindOptionsFlagName := fmt.Sprintf("%v.BindOptions", cmdPrefix)
	if cmd.Flags().Changed(bindOptionsFlagName) {
		// info: complex object BindOptions MountBindOptions is retrieved outside this Changed() block
	}
	bindOptionsFlagValue := m.BindOptions
	if swag.IsZero(bindOptionsFlagValue) {
		bindOptionsFlagValue = &models.MountBindOptions{}
	}

	err, bindOptionsAdded := retrieveModelMountBindOptionsFlags(depth+1, bindOptionsFlagValue, bindOptionsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bindOptionsAdded
	if bindOptionsAdded {
		m.BindOptions = bindOptionsFlagValue
	}

	return nil, retAdded
}

func retrieveMountConsistencyFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	consistencyFlagName := fmt.Sprintf("%v.Consistency", cmdPrefix)
	if cmd.Flags().Changed(consistencyFlagName) {

		var consistencyFlagName string
		if cmdPrefix == "" {
			consistencyFlagName = "Consistency"
		} else {
			consistencyFlagName = fmt.Sprintf("%v.Consistency", cmdPrefix)
		}

		consistencyFlagValue, err := cmd.Flags().GetString(consistencyFlagName)
		if err != nil {
			return err, false
		}
		m.Consistency = consistencyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountReadOnlyFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	readOnlyFlagName := fmt.Sprintf("%v.ReadOnly", cmdPrefix)
	if cmd.Flags().Changed(readOnlyFlagName) {

		var readOnlyFlagName string
		if cmdPrefix == "" {
			readOnlyFlagName = "ReadOnly"
		} else {
			readOnlyFlagName = fmt.Sprintf("%v.ReadOnly", cmdPrefix)
		}

		readOnlyFlagValue, err := cmd.Flags().GetBool(readOnlyFlagName)
		if err != nil {
			return err, false
		}
		m.ReadOnly = readOnlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountSourceFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Source = sourceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountTargetFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	targetFlagName := fmt.Sprintf("%v.Target", cmdPrefix)
	if cmd.Flags().Changed(targetFlagName) {

		var targetFlagName string
		if cmdPrefix == "" {
			targetFlagName = "Target"
		} else {
			targetFlagName = fmt.Sprintf("%v.Target", cmdPrefix)
		}

		targetFlagValue, err := cmd.Flags().GetString(targetFlagName)
		if err != nil {
			return err, false
		}
		m.Target = targetFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountTmpfsOptionsFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tmpfsOptionsFlagName := fmt.Sprintf("%v.TmpfsOptions", cmdPrefix)
	if cmd.Flags().Changed(tmpfsOptionsFlagName) {
		// info: complex object TmpfsOptions MountTmpfsOptions is retrieved outside this Changed() block
	}
	tmpfsOptionsFlagValue := m.TmpfsOptions
	if swag.IsZero(tmpfsOptionsFlagValue) {
		tmpfsOptionsFlagValue = &models.MountTmpfsOptions{}
	}

	err, tmpfsOptionsAdded := retrieveModelMountTmpfsOptionsFlags(depth+1, tmpfsOptionsFlagValue, tmpfsOptionsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tmpfsOptionsAdded
	if tmpfsOptionsAdded {
		m.TmpfsOptions = tmpfsOptionsFlagValue
	}

	return nil, retAdded
}

func retrieveMountTypeFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveMountVolumeOptionsFlags(depth int, m *models.Mount, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	volumeOptionsFlagName := fmt.Sprintf("%v.VolumeOptions", cmdPrefix)
	if cmd.Flags().Changed(volumeOptionsFlagName) {
		// info: complex object VolumeOptions MountVolumeOptions is retrieved outside this Changed() block
	}
	volumeOptionsFlagValue := m.VolumeOptions
	if swag.IsZero(volumeOptionsFlagValue) {
		volumeOptionsFlagValue = &models.MountVolumeOptions{}
	}

	err, volumeOptionsAdded := retrieveModelMountVolumeOptionsFlags(depth+1, volumeOptionsFlagValue, volumeOptionsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || volumeOptionsAdded
	if volumeOptionsAdded {
		m.VolumeOptions = volumeOptionsFlagValue
	}

	return nil, retAdded
}

// Extra schema cli for MountBindOptions

// register flags to command
func registerModelMountBindOptionsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMountBindOptionsNonRecursive(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountBindOptionsPropagation(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountBindOptionsNonRecursive(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nonRecursiveDescription := `Disable recursive bind mount.`

	var nonRecursiveFlagName string
	if cmdPrefix == "" {
		nonRecursiveFlagName = "NonRecursive"
	} else {
		nonRecursiveFlagName = fmt.Sprintf("%v.NonRecursive", cmdPrefix)
	}

	var nonRecursiveFlagDefault bool

	_ = cmd.PersistentFlags().Bool(nonRecursiveFlagName, nonRecursiveFlagDefault, nonRecursiveDescription)

	return nil
}

func registerMountBindOptionsPropagation(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	propagationDescription := `Enum: ["private","rprivate","shared","rshared","slave","rslave"]. A propagation mode with the value ` + "`" + `[r]private` + "`" + `, ` + "`" + `[r]shared` + "`" + `, or ` + "`" + `[r]slave` + "`" + `.`

	var propagationFlagName string
	if cmdPrefix == "" {
		propagationFlagName = "Propagation"
	} else {
		propagationFlagName = fmt.Sprintf("%v.Propagation", cmdPrefix)
	}

	var propagationFlagDefault string

	_ = cmd.PersistentFlags().String(propagationFlagName, propagationFlagDefault, propagationDescription)

	if err := cmd.RegisterFlagCompletionFunc(propagationFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["private","rprivate","shared","rshared","slave","rslave"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMountBindOptionsFlags(depth int, m *models.MountBindOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nonRecursiveAdded := retrieveMountBindOptionsNonRecursiveFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nonRecursiveAdded

	err, propagationAdded := retrieveMountBindOptionsPropagationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || propagationAdded

	return nil, retAdded
}

func retrieveMountBindOptionsNonRecursiveFlags(depth int, m *models.MountBindOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nonRecursiveFlagName := fmt.Sprintf("%v.NonRecursive", cmdPrefix)
	if cmd.Flags().Changed(nonRecursiveFlagName) {

		var nonRecursiveFlagName string
		if cmdPrefix == "" {
			nonRecursiveFlagName = "NonRecursive"
		} else {
			nonRecursiveFlagName = fmt.Sprintf("%v.NonRecursive", cmdPrefix)
		}

		nonRecursiveFlagValue, err := cmd.Flags().GetBool(nonRecursiveFlagName)
		if err != nil {
			return err, false
		}
		m.NonRecursive = &nonRecursiveFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountBindOptionsPropagationFlags(depth int, m *models.MountBindOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	propagationFlagName := fmt.Sprintf("%v.Propagation", cmdPrefix)
	if cmd.Flags().Changed(propagationFlagName) {

		var propagationFlagName string
		if cmdPrefix == "" {
			propagationFlagName = "Propagation"
		} else {
			propagationFlagName = fmt.Sprintf("%v.Propagation", cmdPrefix)
		}

		propagationFlagValue, err := cmd.Flags().GetString(propagationFlagName)
		if err != nil {
			return err, false
		}
		m.Propagation = propagationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for MountTmpfsOptions

// register flags to command
func registerModelMountTmpfsOptionsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMountTmpfsOptionsMode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountTmpfsOptionsSizeBytes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountTmpfsOptionsMode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	modeDescription := `The permission mode for the tmpfs mount in an integer.`

	var modeFlagName string
	if cmdPrefix == "" {
		modeFlagName = "Mode"
	} else {
		modeFlagName = fmt.Sprintf("%v.Mode", cmdPrefix)
	}

	var modeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(modeFlagName, modeFlagDefault, modeDescription)

	return nil
}

func registerMountTmpfsOptionsSizeBytes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeBytesDescription := `The size for the tmpfs mount in bytes.`

	var sizeBytesFlagName string
	if cmdPrefix == "" {
		sizeBytesFlagName = "SizeBytes"
	} else {
		sizeBytesFlagName = fmt.Sprintf("%v.SizeBytes", cmdPrefix)
	}

	var sizeBytesFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeBytesFlagName, sizeBytesFlagDefault, sizeBytesDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMountTmpfsOptionsFlags(depth int, m *models.MountTmpfsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, modeAdded := retrieveMountTmpfsOptionsModeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || modeAdded

	err, sizeBytesAdded := retrieveMountTmpfsOptionsSizeBytesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeBytesAdded

	return nil, retAdded
}

func retrieveMountTmpfsOptionsModeFlags(depth int, m *models.MountTmpfsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	modeFlagName := fmt.Sprintf("%v.Mode", cmdPrefix)
	if cmd.Flags().Changed(modeFlagName) {

		var modeFlagName string
		if cmdPrefix == "" {
			modeFlagName = "Mode"
		} else {
			modeFlagName = fmt.Sprintf("%v.Mode", cmdPrefix)
		}

		modeFlagValue, err := cmd.Flags().GetInt64(modeFlagName)
		if err != nil {
			return err, false
		}
		m.Mode = modeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMountTmpfsOptionsSizeBytesFlags(depth int, m *models.MountTmpfsOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sizeBytesFlagName := fmt.Sprintf("%v.SizeBytes", cmdPrefix)
	if cmd.Flags().Changed(sizeBytesFlagName) {

		var sizeBytesFlagName string
		if cmdPrefix == "" {
			sizeBytesFlagName = "SizeBytes"
		} else {
			sizeBytesFlagName = fmt.Sprintf("%v.SizeBytes", cmdPrefix)
		}

		sizeBytesFlagValue, err := cmd.Flags().GetInt64(sizeBytesFlagName)
		if err != nil {
			return err, false
		}
		m.SizeBytes = sizeBytesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for MountVolumeOptions

// register flags to command
func registerModelMountVolumeOptionsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMountVolumeOptionsDriverConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountVolumeOptionsLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountVolumeOptionsNoCopy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountVolumeOptionsDriverConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var driverConfigFlagName string
	if cmdPrefix == "" {
		driverConfigFlagName = "DriverConfig"
	} else {
		driverConfigFlagName = fmt.Sprintf("%v.DriverConfig", cmdPrefix)
	}

	if err := registerModelMountVolumeOptionsDriverConfigFlags(depth+1, driverConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountVolumeOptionsLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerMountVolumeOptionsNoCopy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	noCopyDescription := `Populate volume with data from the target.`

	var noCopyFlagName string
	if cmdPrefix == "" {
		noCopyFlagName = "NoCopy"
	} else {
		noCopyFlagName = fmt.Sprintf("%v.NoCopy", cmdPrefix)
	}

	var noCopyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(noCopyFlagName, noCopyFlagDefault, noCopyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMountVolumeOptionsFlags(depth int, m *models.MountVolumeOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, driverConfigAdded := retrieveMountVolumeOptionsDriverConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverConfigAdded

	err, labelsAdded := retrieveMountVolumeOptionsLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, noCopyAdded := retrieveMountVolumeOptionsNoCopyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || noCopyAdded

	return nil, retAdded
}

func retrieveMountVolumeOptionsDriverConfigFlags(depth int, m *models.MountVolumeOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	driverConfigFlagName := fmt.Sprintf("%v.DriverConfig", cmdPrefix)
	if cmd.Flags().Changed(driverConfigFlagName) {
		// info: complex object DriverConfig MountVolumeOptionsDriverConfig is retrieved outside this Changed() block
	}
	driverConfigFlagValue := m.DriverConfig
	if swag.IsZero(driverConfigFlagValue) {
		driverConfigFlagValue = &models.MountVolumeOptionsDriverConfig{}
	}

	err, driverConfigAdded := retrieveModelMountVolumeOptionsDriverConfigFlags(depth+1, driverConfigFlagValue, driverConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverConfigAdded
	if driverConfigAdded {
		m.DriverConfig = driverConfigFlagValue
	}

	return nil, retAdded
}

func retrieveMountVolumeOptionsLabelsFlags(depth int, m *models.MountVolumeOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveMountVolumeOptionsNoCopyFlags(depth int, m *models.MountVolumeOptions, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	noCopyFlagName := fmt.Sprintf("%v.NoCopy", cmdPrefix)
	if cmd.Flags().Changed(noCopyFlagName) {

		var noCopyFlagName string
		if cmdPrefix == "" {
			noCopyFlagName = "NoCopy"
		} else {
			noCopyFlagName = fmt.Sprintf("%v.NoCopy", cmdPrefix)
		}

		noCopyFlagValue, err := cmd.Flags().GetBool(noCopyFlagName)
		if err != nil {
			return err, false
		}
		m.NoCopy = &noCopyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// Extra schema cli for MountVolumeOptionsDriverConfig

// register flags to command
func registerModelMountVolumeOptionsDriverConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMountVolumeOptionsDriverConfigName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMountVolumeOptionsDriverConfigOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMountVolumeOptionsDriverConfigName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of the driver to use to create the volume.`

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

func registerMountVolumeOptionsDriverConfigOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Options map[string]string map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMountVolumeOptionsDriverConfigFlags(depth int, m *models.MountVolumeOptionsDriverConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrieveMountVolumeOptionsDriverConfigNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, optionsAdded := retrieveMountVolumeOptionsDriverConfigOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || optionsAdded

	return nil, retAdded
}

func retrieveMountVolumeOptionsDriverConfigNameFlags(depth int, m *models.MountVolumeOptionsDriverConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveMountVolumeOptionsDriverConfigOptionsFlags(depth int, m *models.MountVolumeOptionsDriverConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
