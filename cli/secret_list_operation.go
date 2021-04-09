// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/secret"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSecretSecretListCmd returns a cmd to handle operation secretList
func makeOperationSecretSecretListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SecretList",
		Short: ``,
		RunE:  runOperationSecretSecretList,
	}

	if err := registerOperationSecretSecretListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSecretSecretList uses cmd flags to call endpoint api
func runOperationSecretSecretList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := secret.NewSecretListParams()
	if err, _ := retrieveOperationSecretSecretListFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationSecretSecretListResult(appCli.Secret.SecretList(params)); err != nil {
		return err
	}
	return nil
}

// registerOperationSecretSecretListParamFlags registers all flags needed to fill params
func registerOperationSecretSecretListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSecretSecretListFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSecretSecretListFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `A JSON encoded value of the filters (a ` + "`" + `map[string][]string` + "`" + `) to process on the secrets list. Available filters:

- ` + "`" + `id=<secret id>` + "`" + `
- ` + "`" + `label=<key> or label=<key>=value` + "`" + `
- ` + "`" + `name=<secret name>` + "`" + `
- ` + "`" + `names=<secret name>` + "`" + `
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

func retrieveOperationSecretSecretListFiltersFlag(m *secret.SecretListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationSecretSecretListResult prints output to stdout
func printOperationSecretSecretListResult(resp0 *secret.SecretListOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*secret.SecretListOK)
		if ok {
			if !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*secret.SecretListInternalServerError)
		if ok {
			if !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*secret.SecretListServiceUnavailable)
		if ok {
			if !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return err
				}
				fmt.Println(string(msgStr))
				return nil
			}
		}

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
