// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/system"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSystemSystemPingHeadCmd returns a cmd to handle operation systemPingHead
func makeOperationSystemSystemPingHeadCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SystemPingHead",
		Short: `This is a dummy endpoint you can use to test if the server is accessible.`,
		RunE:  runOperationSystemSystemPingHead,
	}

	if err := registerOperationSystemSystemPingHeadParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSystemSystemPingHead uses cmd flags to call endpoint api
func runOperationSystemSystemPingHead(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := system.NewSystemPingHeadParams()
	// make request and then print result
	if err := printOperationSystemSystemPingHeadResult(appCli.System.SystemPingHead(params)); err != nil {
		return err
	}
	return nil
}

// printOperationSystemSystemPingHeadResult prints output to stdout
func printOperationSystemSystemPingHeadResult(resp0 *system.SystemPingHeadOK, respErr error) error {
	if respErr != nil {
		return respErr
	}

	if !swag.IsZero(resp0.Payload) {
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationSystemSystemPingHeadParamFlags registers all flags needed to fill params
func registerOperationSystemSystemPingHeadParamFlags(cmd *cobra.Command) error {
	return nil
}