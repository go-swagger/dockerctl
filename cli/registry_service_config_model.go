// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/dockerctl/models"
	"github.com/spf13/cobra"
)

// Schema cli for RegistryServiceConfig

// register flags to command
func registerModelRegistryServiceConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRegistryServiceConfigAllowNondistributableArtifactsCIDRs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryServiceConfigAllowNondistributableArtifactsHostnames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryServiceConfigIndexConfigs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryServiceConfigInsecureRegistryCIDRs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryServiceConfigMirrors(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegistryServiceConfigAllowNondistributableArtifactsCIDRs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: AllowNondistributableArtifactsCIDRs []string array type is not supported by go-swagger cli yet

	return nil
}

func registerRegistryServiceConfigAllowNondistributableArtifactsHostnames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: AllowNondistributableArtifactsHostnames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerRegistryServiceConfigIndexConfigs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: IndexConfigs map[string]IndexInfo map type is not supported by go-swagger cli yet

	return nil
}

func registerRegistryServiceConfigInsecureRegistryCIDRs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: InsecureRegistryCIDRs []string array type is not supported by go-swagger cli yet

	return nil
}

func registerRegistryServiceConfigMirrors(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: Mirrors []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRegistryServiceConfigFlags(depth int, m *models.RegistryServiceConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, allowNondistributableArtifactsCIdRsAdded := retrieveRegistryServiceConfigAllowNondistributableArtifactsCIDRsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || allowNondistributableArtifactsCIdRsAdded

	err, allowNondistributableArtifactsHostnamesAdded := retrieveRegistryServiceConfigAllowNondistributableArtifactsHostnamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || allowNondistributableArtifactsHostnamesAdded

	err, indexConfigsAdded := retrieveRegistryServiceConfigIndexConfigsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || indexConfigsAdded

	err, insecureRegistryCIdRsAdded := retrieveRegistryServiceConfigInsecureRegistryCIDRsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || insecureRegistryCIdRsAdded

	err, mirrorsAdded := retrieveRegistryServiceConfigMirrorsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mirrorsAdded

	return nil, retAdded
}

func retrieveRegistryServiceConfigAllowNondistributableArtifactsCIDRsFlags(depth int, m *models.RegistryServiceConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	allowNondistributableArtifactsCIdRsFlagName := fmt.Sprintf("%v.AllowNondistributableArtifactsCIDRs", cmdPrefix)
	if cmd.Flags().Changed(allowNondistributableArtifactsCIdRsFlagName) {
		// warning: AllowNondistributableArtifactsCIDRs array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRegistryServiceConfigAllowNondistributableArtifactsHostnamesFlags(depth int, m *models.RegistryServiceConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	allowNondistributableArtifactsHostnamesFlagName := fmt.Sprintf("%v.AllowNondistributableArtifactsHostnames", cmdPrefix)
	if cmd.Flags().Changed(allowNondistributableArtifactsHostnamesFlagName) {
		// warning: AllowNondistributableArtifactsHostnames array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRegistryServiceConfigIndexConfigsFlags(depth int, m *models.RegistryServiceConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	indexConfigsFlagName := fmt.Sprintf("%v.IndexConfigs", cmdPrefix)
	if cmd.Flags().Changed(indexConfigsFlagName) {
		// warning: IndexConfigs map type map[string]IndexInfo is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRegistryServiceConfigInsecureRegistryCIDRsFlags(depth int, m *models.RegistryServiceConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	insecureRegistryCIdRsFlagName := fmt.Sprintf("%v.InsecureRegistryCIDRs", cmdPrefix)
	if cmd.Flags().Changed(insecureRegistryCIdRsFlagName) {
		// warning: InsecureRegistryCIDRs array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRegistryServiceConfigMirrorsFlags(depth int, m *models.RegistryServiceConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mirrorsFlagName := fmt.Sprintf("%v.Mirrors", cmdPrefix)
	if cmd.Flags().Changed(mirrorsFlagName) {
		// warning: Mirrors array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
