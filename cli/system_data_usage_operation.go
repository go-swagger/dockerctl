// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/system"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSystemSystemDataUsageCmd returns a cmd to handle operation systemDataUsage
func makeOperationSystemSystemDataUsageCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SystemDataUsage",
		Short: ``,
		RunE:  runOperationSystemSystemDataUsage,
	}

	if err := registerOperationSystemSystemDataUsageParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSystemSystemDataUsage uses cmd flags to call endpoint api
func runOperationSystemSystemDataUsage(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := system.NewSystemDataUsageParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSystemSystemDataUsageResult(appCli.System.SystemDataUsage(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSystemSystemDataUsageParamFlags registers all flags needed to fill params
func registerOperationSystemSystemDataUsageParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationSystemSystemDataUsageResult parses request result and return the string content
func parseOperationSystemSystemDataUsageResult(resp0 *system.SystemDataUsageOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*system.SystemDataUsageOK)
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
		resp1, ok := iResp1.(*system.SystemDataUsageInternalServerError)
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
func registerModelSystemDataUsageOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSystemDataUsageOKBodyBuildCache(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemDataUsageOKBodyContainers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemDataUsageOKBodyImages(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemDataUsageOKBodyLayersSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemDataUsageOKBodyVolumes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSystemDataUsageOKBodyBuildCache(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: BuildCache []*models.BuildCache array type is not supported by go-swagger cli yet

	return nil
}

func registerSystemDataUsageOKBodyContainers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Containers []models.ContainerSummary array type is not supported by go-swagger cli yet

	return nil
}

func registerSystemDataUsageOKBodyImages(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Images []*models.ImageSummary array type is not supported by go-swagger cli yet

	return nil
}

func registerSystemDataUsageOKBodyLayersSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	layersSizeDescription := ``

	var layersSizeFlagName string
	if cmdPrefix == "" {
		layersSizeFlagName = "LayersSize"
	} else {
		layersSizeFlagName = fmt.Sprintf("%v.LayersSize", cmdPrefix)
	}

	var layersSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(layersSizeFlagName, layersSizeFlagDefault, layersSizeDescription)

	return nil
}

func registerSystemDataUsageOKBodyVolumes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Volumes []*models.Volume array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSystemDataUsageOKBodyFlags(depth int, m *system.SystemDataUsageOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, buildCacheAdded := retrieveSystemDataUsageOKBodyBuildCacheFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || buildCacheAdded

	err, containersAdded := retrieveSystemDataUsageOKBodyContainersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containersAdded

	err, imagesAdded := retrieveSystemDataUsageOKBodyImagesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imagesAdded

	err, layersSizeAdded := retrieveSystemDataUsageOKBodyLayersSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || layersSizeAdded

	err, volumesAdded := retrieveSystemDataUsageOKBodyVolumesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || volumesAdded

	return nil, retAdded
}

func retrieveSystemDataUsageOKBodyBuildCacheFlags(depth int, m *system.SystemDataUsageOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	buildCacheFlagName := fmt.Sprintf("%v.BuildCache", cmdPrefix)
	if cmd.Flags().Changed(buildCacheFlagName) {
		// warning: BuildCache array type []*models.BuildCache is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSystemDataUsageOKBodyContainersFlags(depth int, m *system.SystemDataUsageOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	containersFlagName := fmt.Sprintf("%v.Containers", cmdPrefix)
	if cmd.Flags().Changed(containersFlagName) {
		// warning: Containers array type []models.ContainerSummary is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSystemDataUsageOKBodyImagesFlags(depth int, m *system.SystemDataUsageOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imagesFlagName := fmt.Sprintf("%v.Images", cmdPrefix)
	if cmd.Flags().Changed(imagesFlagName) {
		// warning: Images array type []*models.ImageSummary is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSystemDataUsageOKBodyLayersSizeFlags(depth int, m *system.SystemDataUsageOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	layersSizeFlagName := fmt.Sprintf("%v.LayersSize", cmdPrefix)
	if cmd.Flags().Changed(layersSizeFlagName) {

		var layersSizeFlagName string
		if cmdPrefix == "" {
			layersSizeFlagName = "LayersSize"
		} else {
			layersSizeFlagName = fmt.Sprintf("%v.LayersSize", cmdPrefix)
		}

		layersSizeFlagValue, err := cmd.Flags().GetInt64(layersSizeFlagName)
		if err != nil {
			return err, false
		}
		m.LayersSize = layersSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSystemDataUsageOKBodyVolumesFlags(depth int, m *system.SystemDataUsageOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	volumesFlagName := fmt.Sprintf("%v.Volumes", cmdPrefix)
	if cmd.Flags().Changed(volumesFlagName) {
		// warning: Volumes array type []*models.Volume is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
