// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/config"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationConfigConfigListCmd returns a cmd to handle operation configList
func makeOperationConfigConfigListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ConfigList",
		Short: ``,
		RunE:  runOperationConfigConfigList,
	}

	if err := registerOperationConfigConfigListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationConfigConfigList uses cmd flags to call endpoint api
func runOperationConfigConfigList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := config.NewConfigListParams()
	if err, _ := retrieveOperationConfigConfigListFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationConfigConfigListResult(appCli.Config.ConfigList(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationConfigConfigListFiltersFlag(m *config.ConfigListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationConfigConfigListResult prints output to stdout
func printOperationConfigConfigListResult(resp0 *config.ConfigListOK, respErr error) error {
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

// registerOperationConfigConfigListParamFlags registers all flags needed to fill params
func registerOperationConfigConfigListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationConfigConfigListFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationConfigConfigListFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `A JSON encoded value of the filters (a ` + "`" + `map[string][]string` + "`" + `) to process on the configs list. Available filters:

- ` + "`" + `id=<config id>` + "`" + `
- ` + "`" + `label=<key> or label=<key>=value` + "`" + `
- ` + "`" + `name=<config name>` + "`" + `
- ` + "`" + `names=<config name>` + "`" + `
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
