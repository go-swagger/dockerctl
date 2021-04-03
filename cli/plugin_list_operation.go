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

// makeOperationPluginPluginListCmd returns a cmd to handle operation pluginList
func makeOperationPluginPluginListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "PluginList",
		Short: `Returns information about installed plugins.`,
		RunE:  runOperationPluginPluginList,
	}

	if err := registerOperationPluginPluginListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPluginPluginList uses cmd flags to call endpoint api
func runOperationPluginPluginList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := plugin.NewPluginListParams()
	if err, _ := retrieveOperationPluginPluginListFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationPluginPluginListResult(appCli.Plugin.PluginList(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationPluginPluginListFiltersFlag(m *plugin.PluginListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filters") {

		var filtersFlagName string
		if cmdPrefix == "" {
			filtersFlagName = "filters"
		} else {
			filtersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
		}

		filtersFlagValue, err := cmd.Flags().GetString(filtersFlagName)
		if err != nil {
			return err, false
		}
		m.Filters = &filtersFlagValue

	}
	return nil, retAdded
}

// printOperationPluginPluginListResult prints output to stdout
func printOperationPluginPluginListResult(resp0 *plugin.PluginListOK, respErr error) error {
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

// registerOperationPluginPluginListParamFlags registers all flags needed to fill params
func registerOperationPluginPluginListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPluginPluginListFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPluginPluginListFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `A JSON encoded value of the filters (a ` + "`" + `map[string][]string` + "`" + `) to process on the plugin list. Available filters:

- ` + "`" + `capability=<capability name>` + "`" + `
- ` + "`" + `enable=<true>|<false>` + "`" + `
`

	var filtersFlagName string
	if cmdPrefix == "" {
		filtersFlagName = "filters"
	} else {
		filtersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
	}

	var filtersFlagDefault string

	_ = cmd.PersistentFlags().String(filtersFlagName, filtersFlagDefault, filtersDescription)

	return nil
}