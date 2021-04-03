// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/plugin"

	"github.com/spf13/cobra"
)

// makeOperationPluginPluginPushCmd returns a cmd to handle operation pluginPush
func makeOperationPluginPluginPushCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "PluginPush",
		Short: `Push a plugin to the registry.
`,
		RunE: runOperationPluginPluginPush,
	}

	if err := registerOperationPluginPluginPushParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginPush uses cmd flags to call endpoint api
func runOperationPluginPluginPush(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginPushParams()
	if err, _ := retrieveOperationPluginPluginPushNameFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationPluginPluginPushResult(appCli.Plugin.PluginPush(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationPluginPluginPushNameFlag(m *plugin.PluginPushParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationPluginPluginPushResult prints output to stdout
func printOperationPluginPluginPushResult(resp0 *plugin.PluginPushOK, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response pluginPushOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationPluginPluginPushParamFlags registers all flags needed to fill params
func registerOperationPluginPluginPushParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginPushNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginPushNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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