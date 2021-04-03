// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/dockerctl/client/image"

	"github.com/spf13/cobra"
)

// makeOperationImageImageTagCmd returns a cmd to handle operation imageTag
func makeOperationImageImageTagCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ImageTag",
		Short: `Tag an image so that it becomes part of a repository.`,
		RunE:  runOperationImageImageTag,
	}

	if err := registerOperationImageImageTagParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageTag uses cmd flags to call endpoint api
func runOperationImageImageTag(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageTagParams()
	if err, _ := retrieveOperationImageImageTagNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageTagRepoFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageTagTagFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationImageImageTagResult(appCli.Image.ImageTag(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationImageImageTagNameFlag(m *image.ImageTagParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationImageImageTagRepoFlag(m *image.ImageTagParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("repo") {

		var repoFlagName string
		if cmdPrefix == "" {
			repoFlagName = "repo"
		} else {
			repoFlagName = fmt.Sprintf("%v.repo", cmdPrefix)
		}

		repoFlagValue, err := cmd.Flags().GetString(repoFlagName)
		if err != nil {
			return err, false
		}
		m.Repo = &repoFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageTagTagFlag(m *image.ImageTagParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationImageImageTagResult prints output to stdout
func printOperationImageImageTagResult(resp0 *image.ImageTagCreated, respErr error) error {
	if respErr != nil {
		return respErr
	}

	// warning: non schema response imageTagCreated is not supported by go-swagger cli yet.

	return nil
}

// registerOperationImageImageTagParamFlags registers all flags needed to fill params
func registerOperationImageImageTagParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageTagNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageTagRepoParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageTagTagParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageTagNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Image name or ID to tag.`

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
func registerOperationImageImageTagRepoParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	repoDescription := `The repository to tag in. For example, ` + "`" + `someuser/someimage` + "`" + `.`

	var repoFlagName string
	if cmdPrefix == "" {
		repoFlagName = "repo"
	} else {
		repoFlagName = fmt.Sprintf("%v.repo", cmdPrefix)
	}

	var repoFlagDefault string

	_ = cmd.PersistentFlags().String(repoFlagName, repoFlagDefault, repoDescription)

	return nil
}
func registerOperationImageImageTagTagParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tagDescription := `The name of the new tag.`

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