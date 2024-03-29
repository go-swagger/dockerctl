// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/container"
	"github.com/go-swagger/dockerctl/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationContainerContainerCreateCmd returns a cmd to handle operation containerCreate
func makeOperationContainerContainerCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerCreate",
		Short: ``,
		RunE:  runOperationContainerContainerCreate,
	}

	if err := registerOperationContainerContainerCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerCreate uses cmd flags to call endpoint api
func runOperationContainerContainerCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerCreateParams()
	if err, _ := retrieveOperationContainerContainerCreateBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerCreateNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerContainerCreateResult(appCli.Container.ContainerCreate(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerContainerCreateParamFlags registers all flags needed to fill params
func registerOperationContainerContainerCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerCreateBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerCreateNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerCreateBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. Container to create")

	// add flags for body
	if err := registerModelContainerCreateBodyFlags(0, "containerCreateBody", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationContainerContainerCreateNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Assign the specified name to the container. Must match ` + "`" + `/?[a-zA-Z0-9][a-zA-Z0-9_.-]+` + "`" + `.`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func retrieveOperationContainerContainerCreateBodyFlag(m *container.ContainerCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := container.ContainerCreateBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in ContainerCreateBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = container.ContainerCreateBody{}
	}
	err, added := retrieveModelContainerCreateBodyFlags(0, &bodyValueModel, "containerCreateBody", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationContainerContainerCreateNameFlag(m *container.ContainerCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

	}
	return nil, retAdded
}

// parseOperationContainerContainerCreateResult parses request result and return the string content
func parseOperationContainerContainerCreateResult(resp0 *container.ContainerCreateCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*container.ContainerCreateCreated)
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
		resp1, ok := iResp1.(*container.ContainerCreateBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*container.ContainerCreateNotFound)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*container.ContainerCreateConflict)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*container.ContainerCreateInternalServerError)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
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
func registerModelContainerCreateBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	// register embedded models.ContainerConfig flags

	if err := registerModelContainerConfigFlags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	// register anonymous fields for ContainerCreateParamsBodyAO1

	if err := registerContainerCreateBodyAnonContainerCreateParamsBodyAO1HostConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerCreateBodyAnonContainerCreateParamsBodyAO1NetworkingConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

// inline definition name ContainerCreateParamsBodyAO1, type

func registerContainerCreateBodyAnonContainerCreateParamsBodyAO1HostConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var hostConfigFlagName string
	if cmdPrefix == "" {
		hostConfigFlagName = "HostConfig"
	} else {
		hostConfigFlagName = fmt.Sprintf("%v.HostConfig", cmdPrefix)
	}

	if err := registerModelHostConfigFlags(depth+1, hostConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerCreateBodyAnonContainerCreateParamsBodyAO1NetworkingConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var networkingConfigFlagName string
	if cmdPrefix == "" {
		networkingConfigFlagName = "NetworkingConfig"
	} else {
		networkingConfigFlagName = fmt.Sprintf("%v.NetworkingConfig", cmdPrefix)
	}

	if err := registerModelContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigFlags(depth+1, networkingConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerCreateBodyFlags(depth int, m *container.ContainerCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	// retrieve model models.ContainerConfig
	err, containerCreateParamsBodyAO0Added := retrieveModelContainerConfigFlags(depth, &m.ContainerConfig, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containerCreateParamsBodyAO0Added

	// retrieve allOf ContainerCreateParamsBodyAO1 fields

	err, hostConfigAdded := retrieveContainerCreateBodyAnonContainerCreateParamsBodyAO1HostConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostConfigAdded

	err, networkingConfigAdded := retrieveContainerCreateBodyAnonContainerCreateParamsBodyAO1NetworkingConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkingConfigAdded

	return nil, retAdded
}

// define retrieve functions for fields for inline definition name ContainerCreateParamsBodyAO1

func retrieveContainerCreateBodyAnonContainerCreateParamsBodyAO1HostConfigFlags(depth int, m *container.ContainerCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hostConfigFlagName := fmt.Sprintf("%v.HostConfig", cmdPrefix)
	if cmd.Flags().Changed(hostConfigFlagName) {
		// info: complex object HostConfig models.HostConfig is retrieved outside this Changed() block
	}
	hostConfigFlagValue := m.HostConfig
	if swag.IsZero(hostConfigFlagValue) {
		hostConfigFlagValue = &models.HostConfig{}
	}

	err, hostConfigAdded := retrieveModelHostConfigFlags(depth+1, hostConfigFlagValue, hostConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostConfigAdded
	if hostConfigAdded {
		m.HostConfig = hostConfigFlagValue
	}

	return nil, retAdded
}

func retrieveContainerCreateBodyAnonContainerCreateParamsBodyAO1NetworkingConfigFlags(depth int, m *container.ContainerCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkingConfigFlagName := fmt.Sprintf("%v.NetworkingConfig", cmdPrefix)
	if cmd.Flags().Changed(networkingConfigFlagName) {
		// info: complex object NetworkingConfig ContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfig is retrieved outside this Changed() block
	}
	networkingConfigFlagValue := m.NetworkingConfig
	if swag.IsZero(networkingConfigFlagValue) {
		networkingConfigFlagValue = &container.ContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfig{}
	}

	err, networkingConfigAdded := retrieveModelContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigFlags(depth+1, networkingConfigFlagValue, networkingConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkingConfigAdded
	if networkingConfigAdded {
		m.NetworkingConfig = networkingConfigFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelContainerCreateCreatedBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerCreateCreatedBodyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerCreateCreatedBodyWarnings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerCreateCreatedBodyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. The ID of the created container`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "Id"
	} else {
		idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerContainerCreateCreatedBodyWarnings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Warnings []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerCreateCreatedBodyFlags(depth int, m *container.ContainerCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveContainerCreateCreatedBodyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, warningsAdded := retrieveContainerCreateCreatedBodyWarningsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || warningsAdded

	return nil, retAdded
}

func retrieveContainerCreateCreatedBodyIDFlags(depth int, m *container.ContainerCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.Id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "Id"
		} else {
			idFlagName = fmt.Sprintf("%v.Id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveContainerCreateCreatedBodyWarningsFlags(depth int, m *container.ContainerCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	warningsFlagName := fmt.Sprintf("%v.Warnings", cmdPrefix)
	if cmd.Flags().Changed(warningsFlagName) {
		// warning: Warnings array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

// register flags to command
func registerModelContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigEndpointsConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigEndpointsConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: EndpointsConfig map[string]models.EndpointSettings map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigFlags(depth int, m *container.ContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endpointsConfigAdded := retrieveContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigEndpointsConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointsConfigAdded

	return nil, retAdded
}

func retrieveContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfigEndpointsConfigFlags(depth int, m *container.ContainerCreateParamsBodyContainerCreateParamsBodyAO1NetworkingConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointsConfigFlagName := fmt.Sprintf("%v.EndpointsConfig", cmdPrefix)
	if cmd.Flags().Changed(endpointsConfigFlagName) {
		// warning: EndpointsConfig map type map[string]models.EndpointSettings is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
