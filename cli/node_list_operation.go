// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/node"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNodeNodeListCmd returns a cmd to handle operation nodeList
func makeOperationNodeNodeListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "NodeList",
		Short: ``,
		RunE:  runOperationNodeNodeList,
	}

	if err := registerOperationNodeNodeListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNodeNodeList uses cmd flags to call endpoint api
func runOperationNodeNodeList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := node.NewNodeListParams()
	if err, _ := retrieveOperationNodeNodeListFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationNodeNodeListResult(appCli.Node.NodeList(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationNodeNodeListParamFlags registers all flags needed to fill params
func registerOperationNodeNodeListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNodeNodeListFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNodeNodeListFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `Filters to process on the nodes list, encoded as JSON (a ` + "`" + `map[string][]string` + "`" + `).

Available filters:
- ` + "`" + `id=<node id>` + "`" + `
- ` + "`" + `label=<engine label>` + "`" + `
- ` + "`" + `membership=` + "`" + `(` + "`" + `accepted` + "`" + `|` + "`" + `pending` + "`" + `)` + "`" + `
- ` + "`" + `name=<node name>` + "`" + `
- ` + "`" + `node.label=<node label>` + "`" + `
- ` + "`" + `role=` + "`" + `(` + "`" + `manager` + "`" + `|` + "`" + `worker` + "`" + `)` + "`" + `
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

func retrieveOperationNodeNodeListFiltersFlag(m *node.NodeListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationNodeNodeListResult parses request result and return the string content
func parseOperationNodeNodeListResult(resp0 *node.NodeListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*node.NodeListOK)
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
		resp1, ok := iResp1.(*node.NodeListInternalServerError)
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
		resp2, ok := iResp2.(*node.NodeListServiceUnavailable)
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
