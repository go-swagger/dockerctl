// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/plugin"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationPluginPluginDeleteCmd returns a cmd to handle operation pluginDelete
func makeOperationPluginPluginDeleteCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PluginDelete",
		Short: ``,
		RunE:  runOperationPluginPluginDelete,
	}

	if err := registerOperationPluginPluginDeleteParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginDelete uses cmd flags to call endpoint api
func runOperationPluginPluginDelete(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginDeleteParams()
	if err, _ := retrieveOperationPluginPluginDeleteForceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPluginPluginDeleteNameFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationPluginPluginDeleteResult(appCli.Plugin.PluginDelete(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationPluginPluginDeleteForceFlag(m *plugin.PluginDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("force") {

		var forceFlagName string
		if cmdPrefix == "" {
			forceFlagName = "force"
		} else {
			forceFlagName = fmt.Sprintf("%v.force", cmdPrefix)
		}

		forceFlagValue, err := cmd.Flags().GetBool(forceFlagName)
		if err != nil {
			return err, false
		}
		m.Force = &forceFlagValue

	}
	return nil, retAdded
}
func retrieveOperationPluginPluginDeleteNameFlag(m *plugin.PluginDeleteParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationPluginPluginDeleteResult prints output to stdout
func printOperationPluginPluginDeleteResult(resp0 *plugin.PluginDeleteOK, respErr error) error {
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

// registerOperationPluginPluginDeleteParamFlags registers all flags needed to fill params
func registerOperationPluginPluginDeleteParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginDeleteForceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPluginPluginDeleteNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginDeleteForceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	forceDescription := `Disable the plugin before removing. This may result in issues if the plugin is in use by a container.`

	var forceFlagName string
	if cmdPrefix == "" {
		forceFlagName = "force"
	} else {
		forceFlagName = fmt.Sprintf("%v.force", cmdPrefix)
	}

	var forceFlagDefault bool

	_ = cmd.PersistentFlags().Bool(forceFlagName, forceFlagDefault, forceDescription)

	return nil
}
func registerOperationPluginPluginDeleteNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
