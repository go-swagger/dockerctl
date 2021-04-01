// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/node"

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
	// make request and then print result
	if err := printOperationNodeNodeListResult(appCli.Node.NodeList(params)); err != nil {
		return err
	}
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

// printOperationNodeNodeListResult prints output to stdout
func printOperationNodeNodeListResult(resp0 *node.NodeListOK, respErr error) error {
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
