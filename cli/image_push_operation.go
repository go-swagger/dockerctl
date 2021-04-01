// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/image"

	"github.com/spf13/cobra"
)

// makeOperationImageImagePushCmd returns a cmd to handle operation imagePush
func makeOperationImageImagePushCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "ImagePush",
		Short: `Push an image to a registry.

If you wish to push an image on to a private registry, that image must already have a tag which references the registry. For example, ` + "`" + `registry.example.com/myimage:latest` + "`" + `.

The push is cancelled if the HTTP connection is closed.
`,
		RunE: runOperationImageImagePush,
	}

	if err := registerOperationImageImagePushParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImagePush uses cmd flags to call endpoint api
func runOperationImageImagePush(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImagePushParams()
	if err, _ := retrieveOperationImageImagePushXRegistryAuthFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImagePushNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImagePushTagFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationImageImagePushResult(appCli.Image.ImagePush(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationImageImagePushXRegistryAuthFlag(m *image.ImagePushParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Registry-Auth") {

		var xRegistryAuthFlagName string
		if cmdPrefix == "" {
			xRegistryAuthFlagName = "X-Registry-Auth"
		} else {
			xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
		}

		xRegistryAuthFlagValue, err := cmd.Flags().GetString(xRegistryAuthFlagName)
		if err != nil {
			return err, false
		}
		m.XRegistryAuth = xRegistryAuthFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImagePushNameFlag(m *image.ImagePushParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("name") {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImagePushTagFlag(m *image.ImagePushParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("tag") {

		var tagFlagName string
		if cmdPrefix == "" {
			tagFlagName = "tag"
		} else {
			tagFlagName = fmt.Sprintf("%v.tag", cmdPrefix)
		}

		tagFlagValue, err := cmd.Flags().GetString(tagFlagName)
		if err != nil {
			return err, false
		}
		m.Tag = &tagFlagValue

	}
	return nil, retAdded
}

// printOperationImageImagePushResult prints output to stdout
func printOperationImageImagePushResult(resp0 *image.ImagePushOK, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response imagePushOK is not supported by go-swagger cli yet.

	return nil
}

// registerOperationImageImagePushParamFlags registers all flags needed to fill params
func registerOperationImageImagePushParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImagePushXRegistryAuthParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImagePushNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImagePushTagParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImagePushXRegistryAuthParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xRegistryAuthDescription := `Required. A base64-encoded auth configuration. [See the authentication section for details.](#section/Authentication)`

	var xRegistryAuthFlagName string
	if cmdPrefix == "" {
		xRegistryAuthFlagName = "X-Registry-Auth"
	} else {
		xRegistryAuthFlagName = fmt.Sprintf("%v.X-Registry-Auth", cmdPrefix)
	}

	var xRegistryAuthFlagDefault string

	_ = cmd.PersistentFlags().String(xRegistryAuthFlagName, xRegistryAuthFlagDefault, xRegistryAuthDescription)

	return nil
}
func registerOperationImageImagePushNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Image name or ID.`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}
func registerOperationImageImagePushTagParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tagDescription := `The tag to associate with the image on the registry.`

	var tagFlagName string
	if cmdPrefix == "" {
		tagFlagName = "tag"
	} else {
		tagFlagName = fmt.Sprintf("%v.tag", cmdPrefix)
	}

	var tagFlagDefault string

	_ = cmd.PersistentFlags().String(tagFlagName, tagFlagDefault, tagDescription)

	return nil
}
