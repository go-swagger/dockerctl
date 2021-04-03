// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/plugin"

	"github.com/spf13/cobra"
)

// makeOperationPluginPluginEnableCmd returns a cmd to handle operation pluginEnable
func makeOperationPluginPluginEnableCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PluginEnable",
		Short: ``,
		RunE:  runOperationPluginPluginEnable,
	}

	if err := registerOperationPluginPluginEnableParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginEnable uses cmd flags to call endpoint api
func runOperationPluginPluginEnable(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginEnableParams()
	if err, _ := retrieveOperationPluginPluginEnableNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPluginPluginEnableTimeoutFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationPluginPluginEnableResult(appCli.Plugin.PluginEnable(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationPluginPluginEnableNameFlag(m *plugin.PluginEnableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPluginPluginEnableTimeoutFlag(m *plugin.PluginEnableParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("timeout") {

		var timeoutFlagName string
		if cmdPrefix == "" {
			timeoutFlagName = "timeout"
		} else {
			timeoutFlagName = fmt.Sprintf("%v.timeout", cmdPrefix)
		}

		timeoutFlagValue, err := cmd.Flags().GetInt64(timeoutFlagName)
		if err != nil {
			return err, false
		}
		m.Timeout = &timeoutFlagValue

	}
	return nil, retAdded
}

// printOperationPluginPluginEnableResult prints output to stdout
func printOperationPluginPluginEnableResult(resp0 *plugin.PluginEnableOK, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response pluginEnableOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationPluginPluginEnableParamFlags registers all flags needed to fill params
func registerOperationPluginPluginEnableParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginEnableNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPluginPluginEnableTimeoutParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginEnableNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPluginPluginEnableTimeoutParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	timeoutDescription := `Set the HTTP client timeout (in seconds)`

	var timeoutFlagName string
	if cmdPrefix == "" {
		timeoutFlagName = "timeout"
	} else {
		timeoutFlagName = fmt.Sprintf("%v.timeout", cmdPrefix)
	}

	var timeoutFlagDefault int64

	_ = cmd.PersistentFlags().Int64(timeoutFlagName, timeoutFlagDefault, timeoutDescription)

	return nil
}