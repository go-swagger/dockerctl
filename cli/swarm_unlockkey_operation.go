// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/swarm"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSwarmSwarmUnlockkeyCmd returns a cmd to handle operation swarmUnlockkey
func makeOperationSwarmSwarmUnlockkeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "SwarmUnlockkey",
		Short: ``,
		RunE:  runOperationSwarmSwarmUnlockkey,
	}

	if err := registerOperationSwarmSwarmUnlockkeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSwarmSwarmUnlockkey uses cmd flags to call endpoint api
func runOperationSwarmSwarmUnlockkey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := swarm.NewSwarmUnlockkeyParams()
	// make request and then print result
	if err := printOperationSwarmSwarmUnlockkeyResult(appCli.Swarm.SwarmUnlockkey(params)); err != nil {
		return err
	}
	return nil
}

// printOperationSwarmSwarmUnlockkeyResult prints output to stdout
func printOperationSwarmSwarmUnlockkeyResult(resp0 *swarm.SwarmUnlockkeyOK, respErr error) error {
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

// registerOperationSwarmSwarmUnlockkeyParamFlags registers all flags needed to fill params
func registerOperationSwarmSwarmUnlockkeyParamFlags(cmd *cobra.Command) error {
	return nil
}

// register flags to command
func registerModelSwarmUnlockkeyOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSwarmUnlockkeyOKBodyUnlockKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSwarmUnlockkeyOKBodyUnlockKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	unlockKeyDescription := `The swarm's unlock key.`

	var unlockKeyFlagName string
	if cmdPrefix == "" {
		unlockKeyFlagName = "UnlockKey"
	} else {
		unlockKeyFlagName = fmt.Sprintf("%v.UnlockKey", cmdPrefix)
	}

	var unlockKeyFlagDefault string

	_ = cmd.PersistentFlags().String(unlockKeyFlagName, unlockKeyFlagDefault, unlockKeyDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSwarmUnlockkeyOKBodyFlags(depth int, m *swarm.SwarmUnlockkeyOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, unlockKeyAdded := retrieveSwarmUnlockkeyOKBodyUnlockKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unlockKeyAdded

	return nil, retAdded
}

func retrieveSwarmUnlockkeyOKBodyUnlockKeyFlags(depth int, m *swarm.SwarmUnlockkeyOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	unlockKeyFlagName := fmt.Sprintf("%v.UnlockKey", cmdPrefix)
	if cmd.Flags().Changed(unlockKeyFlagName) {

		var unlockKeyFlagName string
		if cmdPrefix == "" {
			unlockKeyFlagName = "UnlockKey"
		} else {
			unlockKeyFlagName = fmt.Sprintf("%v.UnlockKey", cmdPrefix)
		}

		unlockKeyFlagValue, err := cmd.Flags().GetString(unlockKeyFlagName)
		if err != nil {
			return err, false
		}
		m.UnlockKey = unlockKeyFlagValue

		retAdded = true
	}
	return nil, retAdded
}
