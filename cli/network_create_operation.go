// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/network"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationNetworkNetworkCreateCmd returns a cmd to handle operation networkCreate
func makeOperationNetworkNetworkCreateCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "NetworkCreate",
		Short: ``,
		RunE:  runOperationNetworkNetworkCreate,
	}

	if err := registerOperationNetworkNetworkCreateParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationNetworkNetworkCreate uses cmd flags to call endpoint api
func runOperationNetworkNetworkCreate(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := network.NewNetworkCreateParams()
	if err, _ := retrieveOperationNetworkNetworkCreateNetworkConfigFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationNetworkNetworkCreateResult(appCli.Network.NetworkCreate(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationNetworkNetworkCreateNetworkConfigFlag(m *network.NetworkCreateParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("networkConfig") {
		// Read networkConfig string from cmd and unmarshal
		networkConfigValueStr, err := cmd.Flags().GetString("networkConfig")
		if err != nil {
			return err, false
		}

		networkConfigValue := network.NetworkCreateBody{}
		if err := json.Unmarshal([]byte(networkConfigValueStr), &networkConfigValue); err != nil {
			return fmt.Errorf("cannot unmarshal networkConfig string in NetworkCreateBody: %v", err), false
		}
		m.NetworkConfig = networkConfigValue
	}
	networkConfigValueModel := m.NetworkConfig
	if swag.IsZero(networkConfigValueModel) {
		networkConfigValueModel = network.NetworkCreateBody{}
	}
	err, added := retrieveModelNetworkCreateBodyFlags(0, &networkConfigValueModel, "networkCreateBody", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.NetworkConfig = networkConfigValueModel
	}
	networkConfigValueDebugBytes, err := json.Marshal(m.NetworkConfig)
	if err != nil {
		return err, false
	}
	logDebugf("NetworkConfig payload: %v", string(networkConfigValueDebugBytes))
	return nil, retAdded
}

// printOperationNetworkNetworkCreateResult prints output to stdout
func printOperationNetworkNetworkCreateResult(resp0 *network.NetworkCreateCreated, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*network.NetworkCreateCreated)
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
		resp1, ok := iResp1.(*network.NetworkCreateForbidden)
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
		resp2, ok := iResp2.(*network.NetworkCreateNotFound)
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

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*network.NetworkCreateInternalServerError)
		if ok {
			if !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
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

// registerOperationNetworkNetworkCreateParamFlags registers all flags needed to fill params
func registerOperationNetworkNetworkCreateParamFlags(cmd *cobra.Command) error {
	if err := registerOperationNetworkNetworkCreateNetworkConfigParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationNetworkNetworkCreateNetworkConfigParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var networkConfigFlagName string
	if cmdPrefix == "" {
		networkConfigFlagName = "networkConfig"
	} else {
		networkConfigFlagName = fmt.Sprintf("%v.networkConfig", cmdPrefix)
	}

	exampleNetworkConfigStr := "go-swagger TODO"
	_ = cmd.PersistentFlags().String(networkConfigFlagName, "", fmt.Sprintf("Optional json string for [networkConfig] of form %v.Network configuration", string(exampleNetworkConfigStr)))

	// add flags for body
	if err := registerModelNetworkCreateBodyFlags(0, "networkCreateBody", cmd); err != nil {
		return err
	}

	return nil
}

// register flags to command
func registerModelNetworkCreateBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNetworkCreateBodyAttachable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyCheckDuplicate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyDriver(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyEnableIPV6(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyIPAM(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyIngress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyInternal(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyLabels(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateBodyOptions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkCreateBodyAttachable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	attachableDescription := `Globally scoped network is manually attachable by regular containers from workers in swarm mode.`

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

func registerNetworkCreateBodyCheckDuplicate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	checkDuplicateDescription := `Check for networks with duplicate names. Since Network is primarily keyed based on a random ID and not on the name, and network name is strictly a user-friendly alias to the network which is uniquely identified using ID, there is no guaranteed way to check for duplicates. CheckDuplicate is there to provide a best effort checking of any networks which has the same name but it is not guaranteed to catch all name collisions.`

	var checkDuplicateFlagName string
	if cmdPrefix == "" {
		checkDuplicateFlagName = "CheckDuplicate"
	} else {
		checkDuplicateFlagName = fmt.Sprintf("%v.CheckDuplicate", cmdPrefix)
	}

	var checkDuplicateFlagDefault bool

	_ = cmd.PersistentFlags().Bool(checkDuplicateFlagName, checkDuplicateFlagDefault, checkDuplicateDescription)

	return nil
}

func registerNetworkCreateBodyDriver(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	driverDescription := `Name of the network driver plugin to use.`

	var driverFlagName string
	if cmdPrefix == "" {
		driverFlagName = "Driver"
	} else {
		driverFlagName = fmt.Sprintf("%v.Driver", cmdPrefix)
	}

	var driverFlagDefault string = "bridge"

	_ = cmd.PersistentFlags().String(driverFlagName, driverFlagDefault, driverDescription)

	return nil
}

func registerNetworkCreateBodyEnableIPV6(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enableIpv6Description := `Enable IPv6 on the network.`

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

func registerNetworkCreateBodyIPAM(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ipAMFlagName string
	if cmdPrefix == "" {
		ipAMFlagName = "IPAM"
	} else {
		ipAMFlagName = fmt.Sprintf("%v.IPAM", cmdPrefix)
	}

	registerModelNetworkCreateBodyFlags(depth+1, ipAMFlagName, cmd)

	return nil
}

func registerNetworkCreateBodyIngress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ingressDescription := `Ingress network is the network which provides the routing-mesh in swarm mode.`

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

func registerNetworkCreateBodyInternal(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	internalDescription := `Restrict external access to the network.`

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

func registerNetworkCreateBodyLabels(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Labels map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerNetworkCreateBodyName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. The network's name.`

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

func registerNetworkCreateBodyOptions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Options map[string]string map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNetworkCreateBodyFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attachableAdded := retrieveNetworkCreateBodyAttachableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attachableAdded

	err, checkDuplicateAdded := retrieveNetworkCreateBodyCheckDuplicateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || checkDuplicateAdded

	err, driverAdded := retrieveNetworkCreateBodyDriverFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverAdded

	err, enableIpv6Added := retrieveNetworkCreateBodyEnableIPV6Flags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enableIpv6Added

	err, ipAMAdded := retrieveNetworkCreateBodyIPAMFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAMAdded

	err, ingressAdded := retrieveNetworkCreateBodyIngressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ingressAdded

	err, internalAdded := retrieveNetworkCreateBodyInternalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || internalAdded

	err, labelsAdded := retrieveNetworkCreateBodyLabelsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || labelsAdded

	err, nameAdded := retrieveNetworkCreateBodyNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, optionsAdded := retrieveNetworkCreateBodyOptionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || optionsAdded

	return nil, retAdded
}

func retrieveNetworkCreateBodyAttachableFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNetworkCreateBodyCheckDuplicateFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	checkDuplicateFlagName := fmt.Sprintf("%v.CheckDuplicate", cmdPrefix)
	if cmd.Flags().Changed(checkDuplicateFlagName) {

		var checkDuplicateFlagName string
		if cmdPrefix == "" {
			checkDuplicateFlagName = "CheckDuplicate"
		} else {
			checkDuplicateFlagName = fmt.Sprintf("%v.CheckDuplicate", cmdPrefix)
		}

		checkDuplicateFlagValue, err := cmd.Flags().GetBool(checkDuplicateFlagName)
		if err != nil {
			return err, false
		}
		m.CheckDuplicate = checkDuplicateFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkCreateBodyDriverFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Driver = &driverFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkCreateBodyEnableIPV6Flags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNetworkCreateBodyIPAMFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	ipAMFlagName := fmt.Sprintf("%v.IPAM", cmdPrefix)
	if cmd.Flags().Changed(ipAMFlagName) {

		ipAMFlagValue := &network.NetworkCreateBody{}
		err, added := retrieveModelNetworkCreateBodyFlags(depth+1, ipAMFlagValue, ipAMFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveNetworkCreateBodyIngressFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNetworkCreateBodyInternalFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNetworkCreateBodyLabelsFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNetworkCreateBodyNameFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveNetworkCreateBodyOptionsFlags(depth int, m *network.NetworkCreateBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// register flags to command
func registerModelNetworkCreateCreatedBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNetworkCreateCreatedBodyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNetworkCreateCreatedBodyWarning(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNetworkCreateCreatedBodyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of the created network.`

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

func registerNetworkCreateCreatedBodyWarning(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	warningDescription := ``

	var warningFlagName string
	if cmdPrefix == "" {
		warningFlagName = "Warning"
	} else {
		warningFlagName = fmt.Sprintf("%v.Warning", cmdPrefix)
	}

	var warningFlagDefault string

	_ = cmd.PersistentFlags().String(warningFlagName, warningFlagDefault, warningDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNetworkCreateCreatedBodyFlags(depth int, m *network.NetworkCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, idAdded := retrieveNetworkCreateCreatedBodyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, warningAdded := retrieveNetworkCreateCreatedBodyWarningFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || warningAdded

	return nil, retAdded
}

func retrieveNetworkCreateCreatedBodyIDFlags(depth int, m *network.NetworkCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveNetworkCreateCreatedBodyWarningFlags(depth int, m *network.NetworkCreateCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	warningFlagName := fmt.Sprintf("%v.Warning", cmdPrefix)
	if cmd.Flags().Changed(warningFlagName) {

		var warningFlagName string
		if cmdPrefix == "" {
			warningFlagName = "Warning"
		} else {
			warningFlagName = fmt.Sprintf("%v.Warning", cmdPrefix)
		}

		warningFlagValue, err := cmd.Flags().GetString(warningFlagName)
		if err != nil {
			return err, false
		}
		m.Warning = warningFlagValue

		retAdded = true
	}
	return nil, retAdded
}
