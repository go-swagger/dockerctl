// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/network"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNetworkNetworkListCmd returns a cmd to handle operation networkList
func makeOperationNetworkNetworkListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "NetworkList",
		Short: `Returns a list of networks. For details on the format, see [the network inspect endpoint](#operation/NetworkInspect).

Note that it uses a different, smaller representation of a network than inspecting a single network. For example,
the list of containers attached to the network is not propagated in API versions 1.28 and up.
`,
		RunE: runOperationNetworkNetworkList,
	}

	if err := registerOperationNetworkNetworkListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNetworkNetworkList uses cmd flags to call endpoint api
func runOperationNetworkNetworkList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := network.NewNetworkListParams()
	if err, _ := retrieveOperationNetworkNetworkListFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationNetworkNetworkListResult(appCli.Network.NetworkList(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationNetworkNetworkListFiltersFlag(m *network.NetworkListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationNetworkNetworkListResult prints output to stdout
func printOperationNetworkNetworkListResult(resp0 *network.NetworkListOK, respErr error) error {
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

// registerOperationNetworkNetworkListParamFlags registers all flags needed to fill params
func registerOperationNetworkNetworkListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNetworkNetworkListFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNetworkNetworkListFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `JSON encoded value of the filters (a ` + "`" + `map[string][]string` + "`" + `) to process on the networks list. Available filters:

- ` + "`" + `dangling=<boolean>` + "`" + ` When set to ` + "`" + `true` + "`" + ` (or ` + "`" + `1` + "`" + `), returns all
   networks that are not in use by a container. When set to ` + "`" + `false` + "`" + `
   (or ` + "`" + `0` + "`" + `), only networks that are in use by one or more
   containers are returned.
- ` + "`" + `driver=<driver-name>` + "`" + ` Matches a network's driver.
- ` + "`" + `id=<network-id>` + "`" + ` Matches all or part of a network ID.
- ` + "`" + `label=<key>` + "`" + ` or ` + "`" + `label=<key>=<value>` + "`" + ` of a network label.
- ` + "`" + `name=<network-name>` + "`" + ` Matches all or part of a network name.
- ` + "`" + `scope=["swarm"|"global"|"local"]` + "`" + ` Filters networks by scope (` + "`" + `swarm` + "`" + `, ` + "`" + `global` + "`" + `, or ` + "`" + `local` + "`" + `).
- ` + "`" + `type=["custom"|"builtin"]` + "`" + ` Filters networks by type. The ` + "`" + `custom` + "`" + ` keyword returns all user-defined networks.
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
