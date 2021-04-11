// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/service"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationServiceServiceListCmd returns a cmd to handle operation serviceList
func makeOperationServiceServiceListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ServiceList",
		Short: ``,
		RunE:  runOperationServiceServiceList,
	}

	if err := registerOperationServiceServiceListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationServiceServiceList uses cmd flags to call endpoint api
func runOperationServiceServiceList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := service.NewServiceListParams()
	if err, _ := retrieveOperationServiceServiceListFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationServiceServiceListStatusFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationServiceServiceListResult(appCli.Service.ServiceList(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationServiceServiceListParamFlags registers all flags needed to fill params
func registerOperationServiceServiceListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationServiceServiceListFiltersParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationServiceServiceListStatusParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationServiceServiceListFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `A JSON encoded value of the filters (a ` + "`" + `map[string][]string` + "`" + `) to process on the services list. Available filters:

- ` + "`" + `id=<service id>` + "`" + `
- ` + "`" + `label=<service label>` + "`" + `
- ` + "`" + `mode=["replicated"|"global"]` + "`" + `
- ` + "`" + `name=<service name>` + "`" + `
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
func registerOperationServiceServiceListStatusParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	statusDescription := `Include service status, with count of running and desired tasks`

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault bool

	_ = cmd.PersistentFlags().Bool(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

func retrieveOperationServiceServiceListFiltersFlag(m *service.ServiceListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationServiceServiceListStatusFlag(m *service.ServiceListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("status") {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetBool(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = &statusFlagValue

	}
	return nil, retAdded
}

// parseOperationServiceServiceListResult parses request result and return the string content
func parseOperationServiceServiceListResult(resp0 *service.ServiceListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*service.ServiceListOK)
		if ok {
			if !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*service.ServiceListInternalServerError)
		if ok {
			if !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*service.ServiceListServiceUnavailable)
		if ok {
			if !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
