// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/plugin"

	"github.com/spf13/cobra"
)

// makeOperationPluginPluginUpgradeCmd returns a cmd to handle operation pluginUpgrade
func makeOperationPluginPluginUpgradeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PluginUpgrade",
		Short: ``,
		RunE:  runOperationPluginPluginUpgrade,
	}

	if err := registerOperationPluginPluginUpgradeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginUpgrade uses cmd flags to call endpoint api
func runOperationPluginPluginUpgrade(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginUpgradeParams()
	if err, _ := retrieveOperationPluginPluginUpgradeXRegistryAuthFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPluginPluginUpgradeBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPluginPluginUpgradeNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPluginPluginUpgradeRemoteFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationPluginPluginUpgradeResult(appCli.Plugin.PluginUpgrade(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationPluginPluginUpgradeXRegistryAuthFlag(m *plugin.PluginUpgradeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Registry-Auth") {

		var xRegistryAuthFlagName string
		if cmdPrefix == "" {
			xRegistryAuthFlagName = "X-Registry-Auth"
		} else {
			xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
		}

		xRegistryAuthFlagValue, err := cmd.Flags().GetString(xRegistryAuthFlagName)
		if err != nil {
			return err, false
		}
		m.XRegistryAuth = &xRegistryAuthFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPluginPluginUpgradeBodyFlag(m *plugin.PluginUpgradeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// warning: body array type []*PluginUpgradeParamsBodyItems0 is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
func retrieveOperationPluginPluginUpgradeNameFlag(m *plugin.PluginUpgradeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPluginPluginUpgradeRemoteFlag(m *plugin.PluginUpgradeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("remote") {

		var remoteFlagName string
		if cmdPrefix == "" {
			remoteFlagName = "remote"
		} else {
			remoteFlagName = fmt.Sprintf("%v.remote", cmdPrefix)
		}

		remoteFlagValue, err := cmd.Flags().GetString(remoteFlagName)
		if err != nil {
			return err, false
		}
		m.Remote = remoteFlagValue

	}
	return nil, retAdded
}

// printOperationPluginPluginUpgradeResult prints output to stdout
func printOperationPluginPluginUpgradeResult(resp0 *plugin.PluginUpgradeNoContent, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response pluginUpgradeNoContent is not supported by go-swagger cli yet.

	return nil
}

// registerOperationPluginPluginUpgradeParamFlags registers all flags needed to fill params
func registerOperationPluginPluginUpgradeParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginUpgradeXRegistryAuthParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPluginPluginUpgradeBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPluginPluginUpgradeNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPluginPluginUpgradeRemoteParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginUpgradeXRegistryAuthParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xRegistryAuthDescription := `A base64-encoded auth configuration to use when pulling a plugin from a registry. [See the authentication section for details.](#section/Authentication)`

	var xRegistryAuthFlagName string
	if cmdPrefix == "" {
		xRegistryAuthFlagName = "X-Registry-Auth"
	} else {
		xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
	}

	var xRegistryAuthFlagDefault string

	_ = cmd.PersistentFlags().String(xRegistryAuthFlagName, xRegistryAuthFlagDefault, xRegistryAuthDescription)

	return nil
}
func registerOperationPluginPluginUpgradeBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {
	// warning: go type []*PluginUpgradeParamsBodyItems0 is not supported by go-swagger cli yet.
	return nil
}
func registerOperationPluginPluginUpgradeNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. The name of the plugin. The ` + "`" + `:latest` + "`" + ` tag is optional, and is the default if omitted.`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}
func registerOperationPluginPluginUpgradeRemoteParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	remoteDescription := `Required. Remote reference to upgrade to.

The ` + "`" + `:latest` + "`" + ` tag is optional, and is used as the default if omitted.
`

	var remoteFlagName string
	if cmdPrefix == "" {
		remoteFlagName = "remote"
	} else {
		remoteFlagName = fmt.Sprintf("%v.remote", cmdPrefix)
	}

	var remoteFlagDefault string

	_ = cmd.PersistentFlags().String(remoteFlagName, remoteFlagDefault, remoteDescription)

	return nil
}

// register flags to command
func registerModelPluginUpgradeParamsBodyItems0Flags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPluginUpgradeParamsBodyItems0Description(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginUpgradeParamsBodyItems0Name(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPluginUpgradeParamsBodyItems0Value(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPluginUpgradeParamsBodyItems0Description(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := ``

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "Description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.Description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerPluginUpgradeParamsBodyItems0Name(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

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

func registerPluginUpgradeParamsBodyItems0Value(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Value []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPluginUpgradeParamsBodyItems0Flags(depth int, m *plugin.PluginUpgradeParamsBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrievePluginUpgradeParamsBodyItems0DescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, nameAdded := retrievePluginUpgradeParamsBodyItems0NameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, valueAdded := retrievePluginUpgradeParamsBodyItems0ValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrievePluginUpgradeParamsBodyItems0DescriptionFlags(depth int, m *plugin.PluginUpgradeParamsBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	descriptionFlagName := fmt.Sprintf("%v.Description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "Description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.Description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrievePluginUpgradeParamsBodyItems0NameFlags(depth int, m *plugin.PluginUpgradeParamsBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrievePluginUpgradeParamsBodyItems0ValueFlags(depth int, m *plugin.PluginUpgradeParamsBodyItems0, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	valueFlagName := fmt.Sprintf("%v.Value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {
		// warning: Value array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
