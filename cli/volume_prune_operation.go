// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/volume"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVolumeVolumePruneCmd returns a cmd to handle operation volumePrune
func makeOperationVolumeVolumePruneCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "VolumePrune",
		Short: ``,
		RunE:  runOperationVolumeVolumePrune,
	}

	if err := registerOperationVolumeVolumePruneParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVolumeVolumePrune uses cmd flags to call endpoint api
func runOperationVolumeVolumePrune(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := volume.NewVolumePruneParams()
	if err, _ := retrieveOperationVolumeVolumePruneFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVolumeVolumePruneResult(appCli.Volume.VolumePrune(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationVolumeVolumePruneParamFlags registers all flags needed to fill params
func registerOperationVolumeVolumePruneParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVolumeVolumePruneFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVolumeVolumePruneFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `Filters to process on the prune list, encoded as JSON (a ` + "`" + `map[string][]string` + "`" + `).

Available filters:
- ` + "`" + `label` + "`" + ` (` + "`" + `label=<key>` + "`" + `, ` + "`" + `label=<key>=<value>` + "`" + `, ` + "`" + `label!=<key>` + "`" + `, or ` + "`" + `label!=<key>=<value>` + "`" + `) Prune volumes with (or without, in case ` + "`" + `label!=...` + "`" + ` is used) the specified labels.
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

func retrieveOperationVolumeVolumePruneFiltersFlag(m *volume.VolumePruneParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationVolumeVolumePruneResult parses request result and return the string content
func parseOperationVolumeVolumePruneResult(resp0 *volume.VolumePruneOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*volume.VolumePruneOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*volume.VolumePruneInternalServerError)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}

// register flags to command
func registerModelVolumePruneOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVolumePruneOKBodySpaceReclaimed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVolumePruneOKBodyVolumesDeleted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVolumePruneOKBodySpaceReclaimed(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	spaceReclaimedDescription := `Disk space reclaimed in bytes`

	var spaceReclaimedFlagName string
	if cmdPrefix == "" {
		spaceReclaimedFlagName = "SpaceReclaimed"
	} else {
		spaceReclaimedFlagName = fmt.Sprintf("%v.SpaceReclaimed", cmdPrefix)
	}

	var spaceReclaimedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(spaceReclaimedFlagName, spaceReclaimedFlagDefault, spaceReclaimedDescription)

	return nil
}

func registerVolumePruneOKBodyVolumesDeleted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: VolumesDeleted []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVolumePruneOKBodyFlags(depth int, m *volume.VolumePruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, spaceReclaimedAdded := retrieveVolumePruneOKBodySpaceReclaimedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || spaceReclaimedAdded

	err, volumesDeletedAdded := retrieveVolumePruneOKBodyVolumesDeletedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || volumesDeletedAdded

	return nil, retAdded
}

func retrieveVolumePruneOKBodySpaceReclaimedFlags(depth int, m *volume.VolumePruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	spaceReclaimedFlagName := fmt.Sprintf("%v.SpaceReclaimed", cmdPrefix)
	if cmd.Flags().Changed(spaceReclaimedFlagName) {

		var spaceReclaimedFlagName string
		if cmdPrefix == "" {
			spaceReclaimedFlagName = "SpaceReclaimed"
		} else {
			spaceReclaimedFlagName = fmt.Sprintf("%v.SpaceReclaimed", cmdPrefix)
		}

		spaceReclaimedFlagValue, err := cmd.Flags().GetInt64(spaceReclaimedFlagName)
		if err != nil {
			return err, false
		}
		m.SpaceReclaimed = spaceReclaimedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVolumePruneOKBodyVolumesDeletedFlags(depth int, m *volume.VolumePruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	volumesDeletedFlagName := fmt.Sprintf("%v.VolumesDeleted", cmdPrefix)
	if cmd.Flags().Changed(volumesDeletedFlagName) {
		// warning: VolumesDeleted array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
