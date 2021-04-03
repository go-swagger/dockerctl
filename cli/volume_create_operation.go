// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/volume"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVolumeVolumeCreateCmd returns a cmd to handle operation volumeCreate
func makeOperationVolumeVolumeCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "VolumeCreate",
		Short: ``,
		RunE:  runOperationVolumeVolumeCreate,
	}

	if err := registerOperationVolumeVolumeCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVolumeVolumeCreate uses cmd flags to call endpoint api
func runOperationVolumeVolumeCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := volume.NewVolumeCreateParams()
	if err, _ := retrieveOperationVolumeVolumeCreateVolumeConfigFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationVolumeVolumeCreateResult(appCli.Volume.VolumeCreate(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationVolumeVolumeCreateVolumeConfigFlag(m *volume.VolumeCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("volumeConfig") {
		// Read volumeConfig string from cmd and unmarshal
		volumeConfigValueStr, err := cmd.Flags().GetString("volumeConfig")
		if err != nil {
			return err, false
		}

		volumeConfigValue := volume.VolumeCreateBody{}
		if err := json.Unmarshal([]byte(volumeConfigValueStr), &volumeConfigValue); err != nil {
			return fmt.Errorf("cannot unmarshal volumeConfig string in VolumeCreateBody: %v", err), false
		}
		m.VolumeConfig = volumeConfigValue
	}
	volumeConfigValueModel := m.VolumeConfig
	if swag.IsZero(volumeConfigValueModel) {
		volumeConfigValueModel = volume.VolumeCreateBody{}
	}
	err, added := retrieveModelVolumeCreateBodyFlags(0, &volumeConfigValueModel, "volumeCreateBody", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.VolumeConfig = volumeConfigValueModel
	}
	volumeConfigValueDebugBytes, err := json.Marshal(m.VolumeConfig)
	if err != nil {
		return err, false
	}
	logDebugf("VolumeConfig payload: %v", string(volumeConfigValueDebugBytes))
	return nil, retAdded
}

// printOperationVolumeVolumeCreateResult prints output to stdout
func printOperationVolumeVolumeCreateResult(resp0 *volume.VolumeCreateCreated, respErr error) error {
	if respErr != nil {
		return respErr
	}

	if !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return err
		}
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationVolumeVolumeCreateParamFlags registers all flags needed to fill params
func registerOperationVolumeVolumeCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVolumeVolumeCreateVolumeConfigParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVolumeVolumeCreateVolumeConfigParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var volumeConfigFlagName string
	if cmdPrefix == "" {
		volumeConfigFlagName = "volumeConfig"
	} else {
		volumeConfigFlagName = fmt.Sprintf("%v.volumeConfig", cmdPrefix)
	}

	exampleVolumeConfigStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(volumeConfigFlagName, "", fmt.Sprintf("Optional json string for [volumeConfig] of form %v.Volume configuration", string(exampleVolumeConfigStr)))

	// add flags for body
	if err := registerModelVolumeCreateBodyFlags(0, "volumeCreateBody", cmd); err != nil {
		return err
	}

	return nil
}

// register flags to command
func registerModelVolumeCreateBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVolumeCreateBodyDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeCreateBodyDriverOpts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeCreateBodyLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumeCreateBodyName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVolumeCreateBodyDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	driverDescription := `Name of the volume driver to use.`

	var driverFlagName string
	if cmdPrefix == "" {
		driverFlagName = "Driver"
	} else {
		driverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
	}

	var driverFlagDefault string = "local"

	_ = cmd.PersistentFlags().String(driverFlagName, driverFlagDefault, driverDescription)

	return nil
}

func registerVolumeCreateBodyDriverOpts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: DriverOpts map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerVolumeCreateBodyLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerVolumeCreateBodyName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `The new volume's name. If not specified, Docker generates a name.`

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVolumeCreateBodyFlags(depth int, m *volume.VolumeCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, driverAdded := retrieveVolumeCreateBodyDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverAdded

	err, driverOptsAdded := retrieveVolumeCreateBodyDriverOptsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverOptsAdded

	err, labelsAdded := retrieveVolumeCreateBodyLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, nameAdded := retrieveVolumeCreateBodyNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveVolumeCreateBodyDriverFlags(depth int, m *volume.VolumeCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVolumeCreateBodyDriverOptsFlags(depth int, m *volume.VolumeCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	driverOptsFlagName := fmt.Sprintf("%v.DriverOpts", cmdPrefix)
	if cmd.Flags().Changed(driverOptsFlagName) {
		// warning: DriverOpts map type map[string]string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveVolumeCreateBodyLabelsFlags(depth int, m *volume.VolumeCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveVolumeCreateBodyNameFlags(depth int, m *volume.VolumeCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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