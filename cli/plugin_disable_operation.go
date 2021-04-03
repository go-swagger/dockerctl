// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/plugin"

	"github.com/spf13/cobra"
)

// makeOperationPluginPluginDisableCmd returns a cmd to handle operation pluginDisable
func makeOperationPluginPluginDisableCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PluginDisable",
		Short: ``,
		RunE:  runOperationPluginPluginDisable,
	}

	if err := registerOperationPluginPluginDisableParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginDisable uses cmd flags to call endpoint api
func runOperationPluginPluginDisable(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginDisableParams()
	if err, _ := retrieveOperationPluginPluginDisableNameFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationPluginPluginDisableResult(appCli.Plugin.PluginDisable(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationPluginPluginDisableNameFlag(m *plugin.PluginDisableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationPluginPluginDisableResult prints output to stdout
func printOperationPluginPluginDisableResult(resp0 *plugin.PluginDisableOK, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response pluginDisableOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationPluginPluginDisableParamFlags registers all flags needed to fill params
func registerOperationPluginPluginDisableParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginDisableNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginDisableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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