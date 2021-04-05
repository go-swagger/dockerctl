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

// makeOperationContainerContainerTopCmd returns a cmd to handle operation containerTop
func makeOperationContainerContainerTopCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ContainerTop",
		Short: `On Unix systems, this is done by running the ` + "`" + `ps` + "`" + ` command. This endpoint is not supported on Windows.`,
		RunE:  runOperationContainerContainerTop,
	}

	if err := registerOperationContainerContainerTopParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerContainerTop uses cmd flags to call endpoint api
func runOperationContainerContainerTop(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewContainerTopParams()
	if err, _ := retrieveOperationContainerContainerTopIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerContainerTopPsArgsFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationContainerContainerTopResult(appCli.Container.ContainerTop(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationContainerContainerTopIDFlag(m *container.ContainerTopParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerContainerTopPsArgsFlag(m *container.ContainerTopParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ps_args") {

		var psArgsFlagName string
		if cmdPrefix == "" {
			psArgsFlagName = "ps_args"
		} else {
			psArgsFlagName = fmt.Sprintf("%v.ps_args", cmdPrefix)
		}

		psArgsFlagValue, err := cmd.Flags().GetString(psArgsFlagName)
		if err != nil {
			return err, false
		}
		m.PsArgs = &psArgsFlagValue

	}
	return nil, retAdded
}

// printOperationContainerContainerTopResult prints output to stdout
func printOperationContainerContainerTopResult(resp0 *container.ContainerTopOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*container.ContainerTopOK)
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
		resp1, ok := iResp1.(*container.ContainerTopNotFound)
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
		resp2, ok := iResp2.(*container.ContainerTopInternalServerError)
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

// registerOperationContainerContainerTopParamFlags registers all flags needed to fill params
func registerOperationContainerContainerTopParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerContainerTopIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerContainerTopPsArgsParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerContainerTopIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerContainerTopPsArgsParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	psArgsDescription := `The arguments to pass to ` + "`" + `ps` + "`" + `. For example, ` + "`" + `aux` + "`" + ``

	var psArgsFlagName string
	if cmdPrefix == "" {
		psArgsFlagName = "ps_args"
	} else {
		psArgsFlagName = fmt.Sprintf("%v.ps_args", cmdPrefix)
	}

	var psArgsFlagDefault string = "-ef"

	_ = cmd.PersistentFlags().String(psArgsFlagName, psArgsFlagDefault, psArgsDescription)

	return nil
}

// register flags to command
func registerModelContainerTopOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerContainerTopOKBodyProcesses(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerContainerTopOKBodyTitles(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerContainerTopOKBodyProcesses(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Processes [][]string array type is not supported by go-swagger cli yet

	return nil
}

func registerContainerTopOKBodyTitles(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Titles []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelContainerTopOKBodyFlags(depth int, m *container.ContainerTopOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, processesAdded := retrieveContainerTopOKBodyProcessesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || processesAdded

	err, titlesAdded := retrieveContainerTopOKBodyTitlesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || titlesAdded

	return nil, retAdded
}

func retrieveContainerTopOKBodyProcessesFlags(depth int, m *container.ContainerTopOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	processesFlagName := fmt.Sprintf("%v.Processes", cmdPrefix)
	if cmd.Flags().Changed(processesFlagName) {
		// warning: Processes array type [][]string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveContainerTopOKBodyTitlesFlags(depth int, m *container.ContainerTopOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	titlesFlagName := fmt.Sprintf("%v.Titles", cmdPrefix)
	if cmd.Flags().Changed(titlesFlagName) {
		// warning: Titles array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
