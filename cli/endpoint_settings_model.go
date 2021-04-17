// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/go-swagger/dockerctl/models"

	"github.com/spf13/cobra"
)

// Schema cli for EndpointSettings

// register flags to command
func registerModelEndpointSettingsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEndpointSettingsAliases(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsDriverOpts(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsEndpointID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsGateway(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsGlobalIPV6Address(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsGlobalIPV6PrefixLen(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsIPAMConfig(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsIPAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsIPPrefixLen(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsIPV6Gateway(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsLinks(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsMacAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointSettingsNetworkID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEndpointSettingsAliases(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Aliases []string array type is not supported by go-swagger cli yet

	return nil
}

func registerEndpointSettingsDriverOpts(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: DriverOpts map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerEndpointSettingsEndpointID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endpointIdDescription := `Unique ID for the service endpoint in a Sandbox.
`

	var endpointIdFlagName string
	if cmdPrefix == "" {
		endpointIdFlagName = "EndpointID"
	} else {
		endpointIdFlagName = fmt.Sprintf("%v.EndpointID", cmdPrefix)
	}

	var endpointIdFlagDefault string

	_ = cmd.PersistentFlags().String(endpointIdFlagName, endpointIdFlagDefault, endpointIdDescription)

	return nil
}

func registerEndpointSettingsGateway(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	gatewayDescription := `Gateway address for this network.
`

	var gatewayFlagName string
	if cmdPrefix == "" {
		gatewayFlagName = "Gateway"
	} else {
		gatewayFlagName = fmt.Sprintf("%v.Gateway", cmdPrefix)
	}

	var gatewayFlagDefault string

	_ = cmd.PersistentFlags().String(gatewayFlagName, gatewayFlagDefault, gatewayDescription)

	return nil
}

func registerEndpointSettingsGlobalIPV6Address(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	globalIpv6AddressDescription := `Global IPv6 address.
`

	var globalIpv6AddressFlagName string
	if cmdPrefix == "" {
		globalIpv6AddressFlagName = "GlobalIPv6Address"
	} else {
		globalIpv6AddressFlagName = fmt.Sprintf("%v.GlobalIPv6Address", cmdPrefix)
	}

	var globalIpv6AddressFlagDefault string

	_ = cmd.PersistentFlags().String(globalIpv6AddressFlagName, globalIpv6AddressFlagDefault, globalIpv6AddressDescription)

	return nil
}

func registerEndpointSettingsGlobalIPV6PrefixLen(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	globalIpv6PrefixLenDescription := `Mask length of the global IPv6 address.
`

	var globalIpv6PrefixLenFlagName string
	if cmdPrefix == "" {
		globalIpv6PrefixLenFlagName = "GlobalIPv6PrefixLen"
	} else {
		globalIpv6PrefixLenFlagName = fmt.Sprintf("%v.GlobalIPv6PrefixLen", cmdPrefix)
	}

	var globalIpv6PrefixLenFlagDefault int64

	_ = cmd.PersistentFlags().Int64(globalIpv6PrefixLenFlagName, globalIpv6PrefixLenFlagDefault, globalIpv6PrefixLenDescription)

	return nil
}

func registerEndpointSettingsIPAMConfig(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var ipAMConfigFlagName string
	if cmdPrefix == "" {
		ipAMConfigFlagName = "IPAMConfig"
	} else {
		ipAMConfigFlagName = fmt.Sprintf("%v.IPAMConfig", cmdPrefix)
	}

	if err := registerModelEndpointIPAMConfigFlags(depth+1, ipAMConfigFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerEndpointSettingsIPAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipAddressDescription := `IPv4 address.
`

	var ipAddressFlagName string
	if cmdPrefix == "" {
		ipAddressFlagName = "IPAddress"
	} else {
		ipAddressFlagName = fmt.Sprintf("%v.IPAddress", cmdPrefix)
	}

	var ipAddressFlagDefault string

	_ = cmd.PersistentFlags().String(ipAddressFlagName, ipAddressFlagDefault, ipAddressDescription)

	return nil
}

func registerEndpointSettingsIPPrefixLen(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipPrefixLenDescription := `Mask length of the IPv4 address.
`

	var ipPrefixLenFlagName string
	if cmdPrefix == "" {
		ipPrefixLenFlagName = "IPPrefixLen"
	} else {
		ipPrefixLenFlagName = fmt.Sprintf("%v.IPPrefixLen", cmdPrefix)
	}

	var ipPrefixLenFlagDefault int64

	_ = cmd.PersistentFlags().Int64(ipPrefixLenFlagName, ipPrefixLenFlagDefault, ipPrefixLenDescription)

	return nil
}

func registerEndpointSettingsIPV6Gateway(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipv6GatewayDescription := `IPv6 gateway address.
`

	var ipv6GatewayFlagName string
	if cmdPrefix == "" {
		ipv6GatewayFlagName = "IPv6Gateway"
	} else {
		ipv6GatewayFlagName = fmt.Sprintf("%v.IPv6Gateway", cmdPrefix)
	}

	var ipv6GatewayFlagDefault string

	_ = cmd.PersistentFlags().String(ipv6GatewayFlagName, ipv6GatewayFlagDefault, ipv6GatewayDescription)

	return nil
}

func registerEndpointSettingsLinks(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Links []string array type is not supported by go-swagger cli yet

	return nil
}

func registerEndpointSettingsMacAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	macAddressDescription := `MAC address for the endpoint on this network.
`

	var macAddressFlagName string
	if cmdPrefix == "" {
		macAddressFlagName = "MacAddress"
	} else {
		macAddressFlagName = fmt.Sprintf("%v.MacAddress", cmdPrefix)
	}

	var macAddressFlagDefault string

	_ = cmd.PersistentFlags().String(macAddressFlagName, macAddressFlagDefault, macAddressDescription)

	return nil
}

func registerEndpointSettingsNetworkID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	networkIdDescription := `Unique ID of the network.
`

	var networkIdFlagName string
	if cmdPrefix == "" {
		networkIdFlagName = "NetworkID"
	} else {
		networkIdFlagName = fmt.Sprintf("%v.NetworkID", cmdPrefix)
	}

	var networkIdFlagDefault string

	_ = cmd.PersistentFlags().String(networkIdFlagName, networkIdFlagDefault, networkIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEndpointSettingsFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, aliasesAdded := retrieveEndpointSettingsAliasesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || aliasesAdded

	err, driverOptsAdded := retrieveEndpointSettingsDriverOptsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || driverOptsAdded

	err, endpointIdAdded := retrieveEndpointSettingsEndpointIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointIdAdded

	err, gatewayAdded := retrieveEndpointSettingsGatewayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || gatewayAdded

	err, globalIpv6AddressAdded := retrieveEndpointSettingsGlobalIPV6AddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || globalIpv6AddressAdded

	err, globalIpv6PrefixLenAdded := retrieveEndpointSettingsGlobalIPV6PrefixLenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || globalIpv6PrefixLenAdded

	err, ipAMConfigAdded := retrieveEndpointSettingsIPAMConfigFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAMConfigAdded

	err, ipAddressAdded := retrieveEndpointSettingsIPAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAddressAdded

	err, ipPrefixLenAdded := retrieveEndpointSettingsIPPrefixLenFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipPrefixLenAdded

	err, ipv6GatewayAdded := retrieveEndpointSettingsIPV6GatewayFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipv6GatewayAdded

	err, linksAdded := retrieveEndpointSettingsLinksFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || linksAdded

	err, macAddressAdded := retrieveEndpointSettingsMacAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || macAddressAdded

	err, networkIdAdded := retrieveEndpointSettingsNetworkIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || networkIdAdded

	return nil, retAdded
}

func retrieveEndpointSettingsAliasesFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	aliasesFlagName := fmt.Sprintf("%v.Aliases", cmdPrefix)
	if cmd.Flags().Changed(aliasesFlagName) {
		// warning: Aliases array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveEndpointSettingsDriverOptsFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	driverOptsFlagName := fmt.Sprintf("%v.DriverOpts", cmdPrefix)
	if cmd.Flags().Changed(driverOptsFlagName) {
		// warning: DriverOpts map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveEndpointSettingsEndpointIDFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointIdFlagName := fmt.Sprintf("%v.EndpointID", cmdPrefix)
	if cmd.Flags().Changed(endpointIdFlagName) {

		var endpointIdFlagName string
		if cmdPrefix == "" {
			endpointIdFlagName = "EndpointID"
		} else {
			endpointIdFlagName = fmt.Sprintf("%v.EndpointID", cmdPrefix)
		}

		endpointIdFlagValue, err := cmd.Flags().GetString(endpointIdFlagName)
		if err != nil {
			return err, false
		}
		m.EndpointID = endpointIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointSettingsGatewayFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	gatewayFlagName := fmt.Sprintf("%v.Gateway", cmdPrefix)
	if cmd.Flags().Changed(gatewayFlagName) {

		var gatewayFlagName string
		if cmdPrefix == "" {
			gatewayFlagName = "Gateway"
		} else {
			gatewayFlagName = fmt.Sprintf("%v.Gateway", cmdPrefix)
		}

		gatewayFlagValue, err := cmd.Flags().GetString(gatewayFlagName)
		if err != nil {
			return err, false
		}
		m.Gateway = gatewayFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointSettingsGlobalIPV6AddressFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	globalIpv6AddressFlagName := fmt.Sprintf("%v.GlobalIPv6Address", cmdPrefix)
	if cmd.Flags().Changed(globalIpv6AddressFlagName) {

		var globalIpv6AddressFlagName string
		if cmdPrefix == "" {
			globalIpv6AddressFlagName = "GlobalIPv6Address"
		} else {
			globalIpv6AddressFlagName = fmt.Sprintf("%v.GlobalIPv6Address", cmdPrefix)
		}

		globalIpv6AddressFlagValue, err := cmd.Flags().GetString(globalIpv6AddressFlagName)
		if err != nil {
			return err, false
		}
		m.GlobalIPV6Address = globalIpv6AddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointSettingsGlobalIPV6PrefixLenFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	globalIpv6PrefixLenFlagName := fmt.Sprintf("%v.GlobalIPv6PrefixLen", cmdPrefix)
	if cmd.Flags().Changed(globalIpv6PrefixLenFlagName) {

		var globalIpv6PrefixLenFlagName string
		if cmdPrefix == "" {
			globalIpv6PrefixLenFlagName = "GlobalIPv6PrefixLen"
		} else {
			globalIpv6PrefixLenFlagName = fmt.Sprintf("%v.GlobalIPv6PrefixLen", cmdPrefix)
		}

		globalIpv6PrefixLenFlagValue, err := cmd.Flags().GetInt64(globalIpv6PrefixLenFlagName)
		if err != nil {
			return err, false
		}
		m.GlobalIPV6PrefixLen = globalIpv6PrefixLenFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointSettingsIPAMConfigFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipAMConfigFlagName := fmt.Sprintf("%v.IPAMConfig", cmdPrefix)
	if cmd.Flags().Changed(ipAMConfigFlagName) {
		// info: complex object IPAMConfig EndpointIPAMConfig is retrieved outside this Changed() block
	}
	ipAMConfigFlagValue := m.IPAMConfig
	if swag.IsZero(ipAMConfigFlagValue) {
		ipAMConfigFlagValue = &models.EndpointIPAMConfig{}
	}

	err, ipAMConfigAdded := retrieveModelEndpointIPAMConfigFlags(depth+1, ipAMConfigFlagValue, ipAMConfigFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAMConfigAdded
	if ipAMConfigAdded {
		m.IPAMConfig = ipAMConfigFlagValue
	}

	return nil, retAdded
}

func retrieveEndpointSettingsIPAddressFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipAddressFlagName := fmt.Sprintf("%v.IPAddress", cmdPrefix)
	if cmd.Flags().Changed(ipAddressFlagName) {

		var ipAddressFlagName string
		if cmdPrefix == "" {
			ipAddressFlagName = "IPAddress"
		} else {
			ipAddressFlagName = fmt.Sprintf("%v.IPAddress", cmdPrefix)
		}

		ipAddressFlagValue, err := cmd.Flags().GetString(ipAddressFlagName)
		if err != nil {
			return err, false
		}
		m.IPAddress = ipAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointSettingsIPPrefixLenFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipPrefixLenFlagName := fmt.Sprintf("%v.IPPrefixLen", cmdPrefix)
	if cmd.Flags().Changed(ipPrefixLenFlagName) {

		var ipPrefixLenFlagName string
		if cmdPrefix == "" {
			ipPrefixLenFlagName = "IPPrefixLen"
		} else {
			ipPrefixLenFlagName = fmt.Sprintf("%v.IPPrefixLen", cmdPrefix)
		}

		ipPrefixLenFlagValue, err := cmd.Flags().GetInt64(ipPrefixLenFlagName)
		if err != nil {
			return err, false
		}
		m.IPPrefixLen = ipPrefixLenFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointSettingsIPV6GatewayFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipv6GatewayFlagName := fmt.Sprintf("%v.IPv6Gateway", cmdPrefix)
	if cmd.Flags().Changed(ipv6GatewayFlagName) {

		var ipv6GatewayFlagName string
		if cmdPrefix == "" {
			ipv6GatewayFlagName = "IPv6Gateway"
		} else {
			ipv6GatewayFlagName = fmt.Sprintf("%v.IPv6Gateway", cmdPrefix)
		}

		ipv6GatewayFlagValue, err := cmd.Flags().GetString(ipv6GatewayFlagName)
		if err != nil {
			return err, false
		}
		m.IPV6Gateway = ipv6GatewayFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointSettingsLinksFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	linksFlagName := fmt.Sprintf("%v.Links", cmdPrefix)
	if cmd.Flags().Changed(linksFlagName) {
		// warning: Links array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveEndpointSettingsMacAddressFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	macAddressFlagName := fmt.Sprintf("%v.MacAddress", cmdPrefix)
	if cmd.Flags().Changed(macAddressFlagName) {

		var macAddressFlagName string
		if cmdPrefix == "" {
			macAddressFlagName = "MacAddress"
		} else {
			macAddressFlagName = fmt.Sprintf("%v.MacAddress", cmdPrefix)
		}

		macAddressFlagValue, err := cmd.Flags().GetString(macAddressFlagName)
		if err != nil {
			return err, false
		}
		m.MacAddress = macAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointSettingsNetworkIDFlags(depth int, m *models.EndpointSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	networkIdFlagName := fmt.Sprintf("%v.NetworkID", cmdPrefix)
	if cmd.Flags().Changed(networkIdFlagName) {

		var networkIdFlagName string
		if cmdPrefix == "" {
			networkIdFlagName = "NetworkID"
		} else {
			networkIdFlagName = fmt.Sprintf("%v.NetworkID", cmdPrefix)
		}

		networkIdFlagValue, err := cmd.Flags().GetString(networkIdFlagName)
		if err != nil {
			return err, false
		}
		m.NetworkID = networkIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
