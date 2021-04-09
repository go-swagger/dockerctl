// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for ClusterInfo

// register flags to command
func registerModelClusterInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerClusterInfoCreatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoDataPathPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoDefaultAddrPool(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoRootRotationInProgress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoSpec(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoSubnetSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoTLSInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoUpdatedAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterInfoVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerClusterInfoCreatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdAtDescription := `Date and time at which the swarm was initialised in
[RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
`

	var createdAtFlagName string
	if cmdPrefix == "" {
		createdAtFlagName = "CreatedAt"
	} else {
		createdAtFlagName = fmt.Sprintf("%v.CreatedAt", cmdPrefix)
	}

	var createdAtFlagDefault string

	_ = cmd.PersistentFlags().String(createdAtFlagName, createdAtFlagDefault, createdAtDescription)

	return nil
}

func registerClusterInfoDataPathPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive DataPathPort uint32 is not supported by go-swagger cli yet

	return nil
}

func registerClusterInfoDefaultAddrPool(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: DefaultAddrPool []string array type is not supported by go-swagger cli yet

	return nil
}

func registerClusterInfoID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of the swarm.`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "ID"
	} else {
		idFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerClusterInfoRootRotationInProgress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rootRotationInProgressDescription := `Whether there is currently a root CA rotation in progress for the swarm`

	var rootRotationInProgressFlagName string
	if cmdPrefix == "" {
		rootRotationInProgressFlagName = "RootRotationInProgress"
	} else {
		rootRotationInProgressFlagName = fmt.Sprintf("%v.RootRotationInProgress", cmdPrefix)
	}

	var rootRotationInProgressFlagDefault bool

	_ = cmd.PersistentFlags().Bool(rootRotationInProgressFlagName, rootRotationInProgressFlagDefault, rootRotationInProgressDescription)

	return nil
}

func registerClusterInfoSpec(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var specFlagName string
	if cmdPrefix == "" {
		specFlagName = "Spec"
	} else {
		specFlagName = fmt.Sprintf("%v.Spec", cmdPrefix)
	}

	if err := registerModelSwarmSpecFlags(depth+1, specFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClusterInfoSubnetSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive SubnetSize uint32 is not supported by go-swagger cli yet

	return nil
}

func registerClusterInfoTLSInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var tlsInfoFlagName string
	if cmdPrefix == "" {
		tlsInfoFlagName = "TLSInfo"
	} else {
		tlsInfoFlagName = fmt.Sprintf("%v.TLSInfo", cmdPrefix)
	}

	if err := registerModelTLSInfoFlags(depth+1, tlsInfoFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClusterInfoUpdatedAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updatedAtDescription := `Date and time at which the swarm was last updated in
[RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
`

	var updatedAtFlagName string
	if cmdPrefix == "" {
		updatedAtFlagName = "UpdatedAt"
	} else {
		updatedAtFlagName = fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
	}

	var updatedAtFlagDefault string

	_ = cmd.PersistentFlags().String(updatedAtFlagName, updatedAtFlagDefault, updatedAtDescription)

	return nil
}

func registerClusterInfoVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "Version"
	} else {
		versionFlagName = fmt.Sprintf("%v.Version", cmdPrefix)
	}

	if err := registerModelObjectVersionFlags(depth+1, versionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelClusterInfoFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, createdAtAdded := retrieveClusterInfoCreatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAtAdded

	err, dataPathPortAdded := retrieveClusterInfoDataPathPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataPathPortAdded

	err, defaultAddrPoolAdded := retrieveClusterInfoDefaultAddrPoolFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultAddrPoolAdded

	err, idAdded := retrieveClusterInfoIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, rootRotationInProgressAdded := retrieveClusterInfoRootRotationInProgressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rootRotationInProgressAdded

	err, specAdded := retrieveClusterInfoSpecFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || specAdded

	err, subnetSizeAdded := retrieveClusterInfoSubnetSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || subnetSizeAdded

	err, tlsInfoAdded := retrieveClusterInfoTLSInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tlsInfoAdded

	err, updatedAtAdded := retrieveClusterInfoUpdatedAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updatedAtAdded

	err, versionAdded := retrieveClusterInfoVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveClusterInfoCreatedAtFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	createdAtFlagName := fmt.Sprintf("%v.CreatedAt", cmdPrefix)
	if cmd.Flags().Changed(createdAtFlagName) {

		var createdAtFlagName string
		if cmdPrefix == "" {
			createdAtFlagName = "CreatedAt"
		} else {
			createdAtFlagName = fmt.Sprintf("%v.CreatedAt", cmdPrefix)
		}

		createdAtFlagValue, err := cmd.Flags().GetString(createdAtFlagName)
		if err != nil {
			return err, false
		}
		m.CreatedAt = createdAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterInfoDataPathPortFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataPathPortFlagName := fmt.Sprintf("%v.DataPathPort", cmdPrefix)
	if cmd.Flags().Changed(dataPathPortFlagName) {

		// warning: primitive DataPathPort uint32 is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterInfoDefaultAddrPoolFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultAddrPoolFlagName := fmt.Sprintf("%v.DefaultAddrPool", cmdPrefix)
	if cmd.Flags().Changed(defaultAddrPoolFlagName) {
		// warning: DefaultAddrPool array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveClusterInfoIDFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.ID", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "ID"
		} else {
			idFlagName = fmt.Sprintf("%v.ID", cmdPrefix)
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

func retrieveClusterInfoRootRotationInProgressFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rootRotationInProgressFlagName := fmt.Sprintf("%v.RootRotationInProgress", cmdPrefix)
	if cmd.Flags().Changed(rootRotationInProgressFlagName) {

		var rootRotationInProgressFlagName string
		if cmdPrefix == "" {
			rootRotationInProgressFlagName = "RootRotationInProgress"
		} else {
			rootRotationInProgressFlagName = fmt.Sprintf("%v.RootRotationInProgress", cmdPrefix)
		}

		rootRotationInProgressFlagValue, err := cmd.Flags().GetBool(rootRotationInProgressFlagName)
		if err != nil {
			return err, false
		}
		m.RootRotationInProgress = rootRotationInProgressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterInfoSpecFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	specFlagName := fmt.Sprintf("%v.Spec", cmdPrefix)
	if cmd.Flags().Changed(specFlagName) {

		specFlagValue := models.SwarmSpec{}
		err, added := retrieveModelSwarmSpecFlags(depth+1, &specFlagValue, specFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.Spec = &specFlagValue
		}
	}

	return nil, retAdded
}

func retrieveClusterInfoSubnetSizeFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	subnetSizeFlagName := fmt.Sprintf("%v.SubnetSize", cmdPrefix)
	if cmd.Flags().Changed(subnetSizeFlagName) {

		// warning: primitive SubnetSize uint32 is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterInfoTLSInfoFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tlsInfoFlagName := fmt.Sprintf("%v.TLSInfo", cmdPrefix)
	if cmd.Flags().Changed(tlsInfoFlagName) {

		tlsInfoFlagValue := models.TLSInfo{}
		err, added := retrieveModelTLSInfoFlags(depth+1, &tlsInfoFlagValue, tlsInfoFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.TLSInfo = &tlsInfoFlagValue
		}
	}

	return nil, retAdded
}

func retrieveClusterInfoUpdatedAtFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updatedAtFlagName := fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
	if cmd.Flags().Changed(updatedAtFlagName) {

		var updatedAtFlagName string
		if cmdPrefix == "" {
			updatedAtFlagName = "UpdatedAt"
		} else {
			updatedAtFlagName = fmt.Sprintf("%v.UpdatedAt", cmdPrefix)
		}

		updatedAtFlagValue, err := cmd.Flags().GetString(updatedAtFlagName)
		if err != nil {
			return err, false
		}
		m.UpdatedAt = updatedAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterInfoVersionFlags(depth int, m *models.ClusterInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.Version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		versionFlagValue := models.ObjectVersion{}
		err, added := retrieveModelObjectVersionFlags(depth+1, &versionFlagValue, versionFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.Version = &versionFlagValue
		}
	}

	return nil, retAdded
}
