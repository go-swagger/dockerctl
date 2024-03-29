// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for PeerNode

// register flags to command
func registerModelPeerNodeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPeerNodeAddr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPeerNodeNodeID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPeerNodeAddr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	addrDescription := `IP address and ports at which this node can be reached.
`

	var addrFlagName string
	if cmdPrefix == "" {
		addrFlagName = "Addr"
	} else {
		addrFlagName = fmt.Sprintf("%v.Addr", cmdPrefix)
	}

	var addrFlagDefault string

	_ = cmd.PersistentFlags().String(addrFlagName, addrFlagDefault, addrDescription)

	return nil
}

func registerPeerNodeNodeID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodeIdDescription := `Unique identifier of for this node in the swarm.`

	var nodeIdFlagName string
	if cmdPrefix == "" {
		nodeIdFlagName = "NodeID"
	} else {
		nodeIdFlagName = fmt.Sprintf("%v.NodeID", cmdPrefix)
	}

	var nodeIdFlagDefault string

	_ = cmd.PersistentFlags().String(nodeIdFlagName, nodeIdFlagDefault, nodeIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPeerNodeFlags(depth int, m *models.PeerNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, addrAdded := retrievePeerNodeAddrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || addrAdded

	err, nodeIdAdded := retrievePeerNodeNodeIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodeIdAdded

	return nil, retAdded
}

func retrievePeerNodeAddrFlags(depth int, m *models.PeerNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	addrFlagName := fmt.Sprintf("%v.Addr", cmdPrefix)
	if cmd.Flags().Changed(addrFlagName) {

		var addrFlagName string
		if cmdPrefix == "" {
			addrFlagName = "Addr"
		} else {
			addrFlagName = fmt.Sprintf("%v.Addr", cmdPrefix)
		}

		addrFlagValue, err := cmd.Flags().GetString(addrFlagName)
		if err != nil {
			return err, false
		}
		m.Addr = addrFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePeerNodeNodeIDFlags(depth int, m *models.PeerNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nodeIdFlagName := fmt.Sprintf("%v.NodeID", cmdPrefix)
	if cmd.Flags().Changed(nodeIdFlagName) {

		var nodeIdFlagName string
		if cmdPrefix == "" {
			nodeIdFlagName = "NodeID"
		} else {
			nodeIdFlagName = fmt.Sprintf("%v.NodeID", cmdPrefix)
		}

		nodeIdFlagValue, err := cmd.Flags().GetString(nodeIdFlagName)
		if err != nil {
			return err, false
		}
		m.NodeID = nodeIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
