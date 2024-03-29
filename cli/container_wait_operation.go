// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/container"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationContainerContainerWaitCmd returns a cmd to handle operation containerWait
func makeOperationContainerContainerWaitCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerWait",
		Short: `Block until a container stops, then returns the exit code.`,
		RunE:  runOperationContainerContainerWait,
	}

	if err := registerOperationContainerContainerWaitParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerWait uses cmd flags to call endpoint api
func runOperationContainerContainerWait(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerWaitParams()
	if err, _ := retrieveOperationContainerContainerWaitConditionFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerWaitIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerContainerWaitResult(appCli.Container.ContainerWait(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerContainerWaitParamFlags registers all flags needed to fill params
func registerOperationContainerContainerWaitParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerWaitConditionParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerWaitIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerWaitConditionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	conditionDescription := `Wait until a container state reaches the given condition, either 'not-running' (default), 'next-exit', or 'removed'.`

	var conditionFlagName string
	if cmdPrefix == "" {
		conditionFlagName = "condition"
	} else {
		conditionFlagName = fmt.Sprintf("%v.condition", cmdPrefix)
	}

	var conditionFlagDefault string = "not-running"

	_ = cmd.PersistentFlags().String(conditionFlagName, conditionFlagDefault, conditionDescription)

	return nil
}
func registerOperationContainerContainerWaitIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. ID or name of the container`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationContainerContainerWaitConditionFlag(m *container.ContainerWaitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("condition") {

		var conditionFlagName string
		if cmdPrefix == "" {
			conditionFlagName = "condition"
		} else {
			conditionFlagName = fmt.Sprintf("%v.condition", cmdPrefix)
		}

		conditionFlagValue, err := cmd.Flags().GetString(conditionFlagName)
		if err != nil {
			return err, false
		}
		m.Condition = &conditionFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerContainerWaitIDFlag(m *container.ContainerWaitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationContainerContainerWaitResult parses request result and return the string content
func parseOperationContainerContainerWaitResult(resp0 *container.ContainerWaitOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*container.ContainerWaitOK)
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
		resp1, ok := iResp1.(*container.ContainerWaitNotFound)
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
		resp2, ok := iResp2.(*container.ContainerWaitInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
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
func registerModelContainerWaitOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerWaitOKBodyError(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerWaitOKBodyStatusCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerWaitOKBodyError(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var errorFlagName string
	if cmdPrefix == "" {
		errorFlagName = "Error"
	} else {
		errorFlagName = fmt.Sprintf("%v.Error", cmdPrefix)
	}

	if err := registerModelContainerWaitOKBodyErrorFlags(depth+1, errorFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerWaitOKBodyStatusCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusCodeDescription := `Required. Exit code of the container`

	var statusCodeFlagName string
	if cmdPrefix == "" {
		statusCodeFlagName = "StatusCode"
	} else {
		statusCodeFlagName = fmt.Sprintf("%v.StatusCode", cmdPrefix)
	}

	var statusCodeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(statusCodeFlagName, statusCodeFlagDefault, statusCodeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerWaitOKBodyFlags(depth int, m *container.ContainerWaitOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorAdded := retrieveContainerWaitOKBodyErrorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded

	err, statusCodeAdded := retrieveContainerWaitOKBodyStatusCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusCodeAdded

	return nil, retAdded
}

func retrieveContainerWaitOKBodyErrorFlags(depth int, m *container.ContainerWaitOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorFlagName := fmt.Sprintf("%v.Error", cmdPrefix)
	if cmd.Flags().Changed(errorFlagName) {
		// info: complex object Error ContainerWaitOKBodyError is retrieved outside this Changed() block
	}
	errorFlagValue := m.Error
	if swag.IsZero(errorFlagValue) {
		errorFlagValue = &container.ContainerWaitOKBodyError{}
	}

	err, errorAdded := retrieveModelContainerWaitOKBodyErrorFlags(depth+1, errorFlagValue, errorFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorAdded
	if errorAdded {
		m.Error = errorFlagValue
	}

	return nil, retAdded
}

func retrieveContainerWaitOKBodyStatusCodeFlags(depth int, m *container.ContainerWaitOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusCodeFlagName := fmt.Sprintf("%v.StatusCode", cmdPrefix)
	if cmd.Flags().Changed(statusCodeFlagName) {

		var statusCodeFlagName string
		if cmdPrefix == "" {
			statusCodeFlagName = "StatusCode"
		} else {
			statusCodeFlagName = fmt.Sprintf("%v.StatusCode", cmdPrefix)
		}

		statusCodeFlagValue, err := cmd.Flags().GetInt64(statusCodeFlagName)
		if err != nil {
			return err, false
		}
		m.StatusCode = statusCodeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

// register flags to command
func registerModelContainerWaitOKBodyErrorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerWaitOKBodyErrorMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerWaitOKBodyErrorMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := `Details of an error`

	var messageFlagName string
	if cmdPrefix == "" {
		messageFlagName = "Message"
	} else {
		messageFlagName = fmt.Sprintf("%v.Message", cmdPrefix)
	}

	var messageFlagDefault string

	_ = cmd.PersistentFlags().String(messageFlagName, messageFlagDefault, messageDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerWaitOKBodyErrorFlags(depth int, m *container.ContainerWaitOKBodyError, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, messageAdded := retrieveContainerWaitOKBodyErrorMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	return nil, retAdded
}

func retrieveContainerWaitOKBodyErrorMessageFlags(depth int, m *container.ContainerWaitOKBodyError, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	messageFlagName := fmt.Sprintf("%v.Message", cmdPrefix)
	if cmd.Flags().Changed(messageFlagName) {

		var messageFlagName string
		if cmdPrefix == "" {
			messageFlagName = "Message"
		} else {
			messageFlagName = fmt.Sprintf("%v.Message", cmdPrefix)
		}

		messageFlagValue, err := cmd.Flags().GetString(messageFlagName)
		if err != nil {
			return err, false
		}
		m.Message = messageFlagValue

		retAdded = true
	}

	return nil, retAdded
}
