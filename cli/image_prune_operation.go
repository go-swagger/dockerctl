// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/image"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationImageImagePruneCmd returns a cmd to handle operation imagePrune
func makeOperationImageImagePruneCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ImagePrune",
		Short: ``,
		RunE:  runOperationImageImagePrune,
	}

	if err := registerOperationImageImagePruneParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImagePrune uses cmd flags to call endpoint api
func runOperationImageImagePrune(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImagePruneParams()
	if err, _ := retrieveOperationImageImagePruneFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImageImagePruneResult(appCli.Image.ImagePrune(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImageImagePruneParamFlags registers all flags needed to fill params
func registerOperationImageImagePruneParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImagePruneFiltersParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImagePruneFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `Filters to process on the prune list, encoded as JSON (a ` + "`" + `map[string][]string` + "`" + `). Available filters:

- ` + "`" + `dangling=<boolean>` + "`" + ` When set to ` + "`" + `true` + "`" + ` (or ` + "`" + `1` + "`" + `), prune only
   unused *and* untagged images. When set to ` + "`" + `false` + "`" + `
   (or ` + "`" + `0` + "`" + `), all unused images are pruned.
- ` + "`" + `until=<string>` + "`" + ` Prune images created before this timestamp. The ` + "`" + `<timestamp>` + "`" + ` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. ` + "`" + `10m` + "`" + `, ` + "`" + `1h30m` + "`" + `) computed relative to the daemon machine’s time.
- ` + "`" + `label` + "`" + ` (` + "`" + `label=<key>` + "`" + `, ` + "`" + `label=<key>=<value>` + "`" + `, ` + "`" + `label!=<key>` + "`" + `, or ` + "`" + `label!=<key>=<value>` + "`" + `) Prune images with (or without, in case ` + "`" + `label!=...` + "`" + ` is used) the specified labels.
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

func retrieveOperationImageImagePruneFiltersFlag(m *image.ImagePruneParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationImageImagePruneResult parses request result and return the string content
func parseOperationImageImagePruneResult(resp0 *image.ImagePruneOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*image.ImagePruneOK)
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
		resp1, ok := iResp1.(*image.ImagePruneInternalServerError)
		if ok {
			if !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
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

// register flags to command
func registerModelImagePruneOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerImagePruneOKBodyImagesDeleted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerImagePruneOKBodySpaceReclaimed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerImagePruneOKBodyImagesDeleted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: ImagesDeleted []*models.ImageDeleteResponseItem array type is not supported by go-swagger cli yet

	return nil
}

func registerImagePruneOKBodySpaceReclaimed(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelImagePruneOKBodyFlags(depth int, m *image.ImagePruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, imagesDeletedAdded := retrieveImagePruneOKBodyImagesDeletedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imagesDeletedAdded

	err, spaceReclaimedAdded := retrieveImagePruneOKBodySpaceReclaimedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || spaceReclaimedAdded

	return nil, retAdded
}

func retrieveImagePruneOKBodyImagesDeletedFlags(depth int, m *image.ImagePruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imagesDeletedFlagName := fmt.Sprintf("%v.ImagesDeleted", cmdPrefix)
	if cmd.Flags().Changed(imagesDeletedFlagName) {
		// warning: ImagesDeleted array type []*models.ImageDeleteResponseItem is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveImagePruneOKBodySpaceReclaimedFlags(depth int, m *image.ImagePruneOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
