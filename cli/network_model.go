// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for Network

// register flags to command
func registerModelNetworkFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNetworkAttachable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkContainers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkEnableIPV6(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkIPAM(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkIngress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkInternal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkScope(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkAttachable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	attachableDescription := ``

	var attachableFlagName string
	if cmdPrefix == "" {
		attachableFlagName = "Attachable"
	} else {
		attachableFlagName = fmt.Sprintf("%v.Attachable", cmdPrefix)
	}

	var attachableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(attachableFlagName, attachableFlagDefault, attachableDescription)

	return nil
}

func registerNetworkContainers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Containers map[string]NetworkContainer map type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkCreated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdDescription := ``

	var createdFlagName string
	if cmdPrefix == "" {
		createdFlagName = "Created"
	} else {
		createdFlagName = fmt.Sprintf("%v.Created", cmdPrefix)
	}

	var createdFlagDefault string

	_ = cmd.PersistentFlags().String(createdFlagName, createdFlagDefault, createdDescription)

	return nil
}

func registerNetworkDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	driverDescription := ``

	var driverFlagName string
	if cmdPrefix == "" {
		driverFlagName = "Driver"
	} else {
		driverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
	}

	var driverFlagDefault string

	_ = cmd.PersistentFlags().String(driverFlagName, driverFlagDefault, driverDescription)

	return nil
}

func registerNetworkEnableIPV6(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enableIpv6Description := ``

	var enableIpv6FlagName string
	if cmdPrefix == "" {
		enableIpv6FlagName = "EnableIPv6"
	} else {
		enableIpv6FlagName = fmt.Sprintf("%v.EnableIPv6", cmdPrefix)
	}

	var enableIpv6FlagDefault bool

	_ = cmd.PersistentFlags().Bool(enableIpv6FlagName, enableIpv6FlagDefault, enableIpv6Description)

	return nil
}

func registerNetworkIPAM(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ipAMFlagName string
	if cmdPrefix == "" {
		ipAMFlagName = "IPAM"
	} else {
		ipAMFlagName = fmt.Sprintf("%v.IPAM", cmdPrefix)
	}

	if err := registerModelIPAMFlags(depth+1, ipAMFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerNetworkIngress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ingressDescription := ``

	var ingressFlagName string
	if cmdPrefix == "" {
		ingressFlagName = "Ingress"
	} else {
		ingressFlagName = fmt.Sprintf("%v.Ingress", cmdPrefix)
	}

	var ingressFlagDefault bool

	_ = cmd.PersistentFlags().Bool(ingressFlagName, ingressFlagDefault, ingressDescription)

	return nil
}

func registerNetworkInternal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	internalDescription := ``

	var internalFlagName string
	if cmdPrefix == "" {
		internalFlagName = "Internal"
	} else {
		internalFlagName = fmt.Sprintf("%v.Internal", cmdPrefix)
	}

	var internalFlagDefault bool

	_ = cmd.PersistentFlags().Bool(internalFlagName, internalFlagDefault, internalDescription)

	return nil
}

func registerNetworkLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "Name"
	} else {
		nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerNetworkOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Options map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkScope(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	scopeDescription := ``

	var scopeFlagName string
	if cmdPrefix == "" {
		scopeFlagName = "Scope"
	} else {
		scopeFlagName = fmt.Sprintf("%v.Scope", cmdPrefix)
	}

	var scopeFlagDefault string

	_ = cmd.PersistentFlags().String(scopeFlagName, scopeFlagDefault, scopeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNetworkFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attachableAdded := retrieveNetworkAttachableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attachableAdded

	err, containersAdded := retrieveNetworkContainersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containersAdded

	err, createdAdded := retrieveNetworkCreatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAdded

	err, driverAdded := retrieveNetworkDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverAdded

	err, enableIpv6Added := retrieveNetworkEnableIPV6Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enableIpv6Added

	err, ipAMAdded := retrieveNetworkIPAMFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAMAdded

	err, idAdded := retrieveNetworkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, ingressAdded := retrieveNetworkIngressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ingressAdded

	err, internalAdded := retrieveNetworkInternalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || internalAdded

	err, labelsAdded := retrieveNetworkLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, nameAdded := retrieveNetworkNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, optionsAdded := retrieveNetworkOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || optionsAdded

	err, scopeAdded := retrieveNetworkScopeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeAdded

	return nil, retAdded
}

func retrieveNetworkAttachableFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	attachableFlagName := fmt.Sprintf("%v.Attachable", cmdPrefix)
	if cmd.Flags().Changed(attachableFlagName) {

		var attachableFlagName string
		if cmdPrefix == "" {
			attachableFlagName = "Attachable"
		} else {
			attachableFlagName = fmt.Sprintf("%v.Attachable", cmdPrefix)
		}

		attachableFlagValue, err := cmd.Flags().GetBool(attachableFlagName)
		if err != nil {
			return err, false
		}
		m.Attachable = attachableFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkContainersFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	containersFlagName := fmt.Sprintf("%v.Containers", cmdPrefix)
	if cmd.Flags().Changed(containersFlagName) {
		// warning: Containers map type map[string]NetworkContainer is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveNetworkCreatedFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	createdFlagName := fmt.Sprintf("%v.Created", cmdPrefix)
	if cmd.Flags().Changed(createdFlagName) {

		var createdFlagName string
		if cmdPrefix == "" {
			createdFlagName = "Created"
		} else {
			createdFlagName = fmt.Sprintf("%v.Created", cmdPrefix)
		}

		createdFlagValue, err := cmd.Flags().GetString(createdFlagName)
		if err != nil {
			return err, false
		}
		m.Created = createdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkDriverFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	driverFlagName := fmt.Sprintf("%v.Driver", cmdPrefix)
	if cmd.Flags().Changed(driverFlagName) {

		var driverFlagName string
		if cmdPrefix == "" {
			driverFlagName = "Driver"
		} else {
			driverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
		}

		driverFlagValue, err := cmd.Flags().GetString(driverFlagName)
		if err != nil {
			return err, false
		}
		m.Driver = driverFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkEnableIPV6Flags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	enableIpv6FlagName := fmt.Sprintf("%v.EnableIPv6", cmdPrefix)
	if cmd.Flags().Changed(enableIpv6FlagName) {

		var enableIpv6FlagName string
		if cmdPrefix == "" {
			enableIpv6FlagName = "EnableIPv6"
		} else {
			enableIpv6FlagName = fmt.Sprintf("%v.EnableIPv6", cmdPrefix)
		}

		enableIpv6FlagValue, err := cmd.Flags().GetBool(enableIpv6FlagName)
		if err != nil {
			return err, false
		}
		m.EnableIPV6 = enableIpv6FlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkIPAMFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	ipAMFlagName := fmt.Sprintf("%v.IPAM", cmdPrefix)
	if cmd.Flags().Changed(ipAMFlagName) {

		ipAMFlagValue := &models.IPAM{}
		err, added := retrieveModelIPAMFlags(depth+1, ipAMFlagValue, ipAMFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
		if added {
			m.IPAM = ipAMFlagValue
		}
	}
	return nil, retAdded
}

func retrieveNetworkIDFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNetworkIngressFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	ingressFlagName := fmt.Sprintf("%v.Ingress", cmdPrefix)
	if cmd.Flags().Changed(ingressFlagName) {

		var ingressFlagName string
		if cmdPrefix == "" {
			ingressFlagName = "Ingress"
		} else {
			ingressFlagName = fmt.Sprintf("%v.Ingress", cmdPrefix)
		}

		ingressFlagValue, err := cmd.Flags().GetBool(ingressFlagName)
		if err != nil {
			return err, false
		}
		m.Ingress = ingressFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkInternalFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	internalFlagName := fmt.Sprintf("%v.Internal", cmdPrefix)
	if cmd.Flags().Changed(internalFlagName) {

		var internalFlagName string
		if cmdPrefix == "" {
			internalFlagName = "Internal"
		} else {
			internalFlagName = fmt.Sprintf("%v.Internal", cmdPrefix)
		}

		internalFlagValue, err := cmd.Flags().GetBool(internalFlagName)
		if err != nil {
			return err, false
		}
		m.Internal = internalFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkLabelsFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	labelsFlagName := fmt.Sprintf("%v.Labels", cmdPrefix)
	if cmd.Flags().Changed(labelsFlagName) {
		// warning: Labels map type map[string]string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveNetworkNameFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	nameFlagName := fmt.Sprintf("%v.Name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "Name"
		} else {
			nameFlagName = fmt.Sprintf("%v.Name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkOptionsFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	optionsFlagName := fmt.Sprintf("%v.Options", cmdPrefix)
	if cmd.Flags().Changed(optionsFlagName) {
		// warning: Options map type map[string]string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveNetworkScopeFlags(depth int, m *models.Network, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	scopeFlagName := fmt.Sprintf("%v.Scope", cmdPrefix)
	if cmd.Flags().Changed(scopeFlagName) {

		var scopeFlagName string
		if cmdPrefix == "" {
			scopeFlagName = "Scope"
		} else {
			scopeFlagName = fmt.Sprintf("%v.Scope", cmdPrefix)
		}

		scopeFlagValue, err := cmd.Flags().GetString(scopeFlagName)
		if err != nil {
			return err, false
		}
		m.Scope = scopeFlagValue

		retAdded = true
	}
	return nil, retAdded
}
