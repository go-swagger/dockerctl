// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/models"
	"github.com/spf13/cobra"
)

// register flags to command
func registerModelIndexInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIndexInfoMirrors(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIndexInfoName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIndexInfoOfficial(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIndexInfoSecure(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIndexInfoMirrors(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Mirrors []string array type is not supported by go-swagger cli yet

	return nil
}

func registerIndexInfoName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of the registry, such as "docker.io".
`

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

func registerIndexInfoOfficial(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	officialDescription := `Indicates whether this is an official registry (i.e., Docker Hub / docker.io)
`

	var officialFlagName string
	if cmdPrefix == "" {
		officialFlagName = "Official"
	} else {
		officialFlagName = fmt.Sprintf("%v.Official", cmdPrefix)
	}

	var officialFlagDefault bool

	_ = cmd.PersistentFlags().Bool(officialFlagName, officialFlagDefault, officialDescription)

	return nil
}

func registerIndexInfoSecure(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	secureDescription := `Indicates if the registry is part of the list of insecure
registries.

If ` + "`" + `false` + "`" + `, the registry is insecure. Insecure registries accept
un-encrypted (HTTP) and/or untrusted (HTTPS with certificates from
unknown CAs) communication.

> **Warning**: Insecure registries can be useful when running a local
> registry. However, because its use creates security vulnerabilities
> it should ONLY be enabled for testing purposes. For increased
> security, users should add their CA to their system's list of
> trusted CAs instead of enabling this option.
`

	var secureFlagName string
	if cmdPrefix == "" {
		secureFlagName = "Secure"
	} else {
		secureFlagName = fmt.Sprintf("%v.Secure", cmdPrefix)
	}

	var secureFlagDefault bool

	_ = cmd.PersistentFlags().Bool(secureFlagName, secureFlagDefault, secureDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIndexInfoFlags(depth int, m *models.IndexInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, mirrorsAdded := retrieveIndexInfoMirrorsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mirrorsAdded

	err, nameAdded := retrieveIndexInfoNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, officialAdded := retrieveIndexInfoOfficialFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || officialAdded

	err, secureAdded := retrieveIndexInfoSecureFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secureAdded

	return nil, retAdded
}

func retrieveIndexInfoMirrorsFlags(depth int, m *models.IndexInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveIndexInfoNameFlags(depth int, m *models.IndexInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveIndexInfoOfficialFlags(depth int, m *models.IndexInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	officialFlagName := fmt.Sprintf("%v.Official", cmdPrefix)
	if cmd.Flags().Changed(officialFlagName) {

		var officialFlagName string
		if cmdPrefix == "" {
			officialFlagName = "Official"
		} else {
			officialFlagName = fmt.Sprintf("%v.Official", cmdPrefix)
		}

		officialFlagValue, err := cmd.Flags().GetBool(officialFlagName)
		if err != nil {
			return err, false
		}
		m.Official = officialFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveIndexInfoSecureFlags(depth int, m *models.IndexInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	secureFlagName := fmt.Sprintf("%v.Secure", cmdPrefix)
	if cmd.Flags().Changed(secureFlagName) {

		var secureFlagName string
		if cmdPrefix == "" {
			secureFlagName = "Secure"
		} else {
			secureFlagName = fmt.Sprintf("%v.Secure", cmdPrefix)
		}

		secureFlagValue, err := cmd.Flags().GetBool(secureFlagName)
		if err != nil {
			return err, false
		}
		m.Secure = secureFlagValue

		retAdded = true
	}
	return nil, retAdded
}