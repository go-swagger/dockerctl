// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/dockerctl/client/image"
	"github.com/go-swagger/dockerctl/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationImageImageCommitCmd returns a cmd to handle operation imageCommit
func makeOperationImageImageCommitCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ImageCommit",
		Short: ``,
		RunE:  runOperationImageImageCommit,
	}

	if err := registerOperationImageImageCommitParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageCommit uses cmd flags to call endpoint api
func runOperationImageImageCommit(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageCommitParams()
	if err, _ := retrieveOperationImageImageCommitAuthorFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCommitChangesFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCommitCommentFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCommitContainerFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCommitContainerConfigFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCommitPauseFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCommitRepoFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationImageImageCommitTagFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationImageImageCommitResult(appCli.Image.ImageCommit(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationImageImageCommitParamFlags registers all flags needed to fill params
func registerOperationImageImageCommitParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageCommitAuthorParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCommitChangesParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCommitCommentParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCommitContainerParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCommitContainerConfigParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCommitPauseParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCommitRepoParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationImageImageCommitTagParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageCommitAuthorParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	authorDescription := `Author of the image (e.g., ` + "`" + `John Hannibal Smith <hannibal@a-team.com>` + "`" + `)`

	var authorFlagName string
	if cmdPrefix == "" {
		authorFlagName = "author"
	} else {
		authorFlagName = fmt.Sprintf("%v.author", cmdPrefix)
	}

	var authorFlagDefault string

	_ = cmd.PersistentFlags().String(authorFlagName, authorFlagDefault, authorDescription)

	return nil
}
func registerOperationImageImageCommitChangesParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	changesDescription := `` + "`" + `Dockerfile` + "`" + ` instructions to apply while committing`

	var changesFlagName string
	if cmdPrefix == "" {
		changesFlagName = "changes"
	} else {
		changesFlagName = fmt.Sprintf("%v.changes", cmdPrefix)
	}

	var changesFlagDefault string

	_ = cmd.PersistentFlags().String(changesFlagName, changesFlagDefault, changesDescription)

	return nil
}
func registerOperationImageImageCommitCommentParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	commentDescription := `Commit message`

	var commentFlagName string
	if cmdPrefix == "" {
		commentFlagName = "comment"
	} else {
		commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
	}

	var commentFlagDefault string

	_ = cmd.PersistentFlags().String(commentFlagName, commentFlagDefault, commentDescription)

	return nil
}
func registerOperationImageImageCommitContainerParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	containerDescription := `The ID or name of the container to commit`

	var containerFlagName string
	if cmdPrefix == "" {
		containerFlagName = "container"
	} else {
		containerFlagName = fmt.Sprintf("%v.container", cmdPrefix)
	}

	var containerFlagDefault string

	_ = cmd.PersistentFlags().String(containerFlagName, containerFlagDefault, containerDescription)

	return nil
}
func registerOperationImageImageCommitContainerConfigParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var containerConfigFlagName string
	if cmdPrefix == "" {
		containerConfigFlagName = "containerConfig"
	} else {
		containerConfigFlagName = fmt.Sprintf("%v.containerConfig", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(containerConfigFlagName, "", "Optional json string for [containerConfig]. The container configuration")

	// add flags for body
	if err := registerModelContainerConfigFlags(0, "containerConfig", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationImageImageCommitPauseParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	pauseDescription := `Whether to pause the container before committing`

	var pauseFlagName string
	if cmdPrefix == "" {
		pauseFlagName = "pause"
	} else {
		pauseFlagName = fmt.Sprintf("%v.pause", cmdPrefix)
	}

	var pauseFlagDefault bool = true

	_ = cmd.PersistentFlags().Bool(pauseFlagName, pauseFlagDefault, pauseDescription)

	return nil
}
func registerOperationImageImageCommitRepoParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	repoDescription := `Repository name for the created image`

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
func registerOperationImageImageCommitTagParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	tagDescription := `Tag name for the create image`

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

func retrieveOperationImageImageCommitAuthorFlag(m *image.ImageCommitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("author") {

		var authorFlagName string
		if cmdPrefix == "" {
			authorFlagName = "author"
		} else {
			authorFlagName = fmt.Sprintf("%v.author", cmdPrefix)
		}

		authorFlagValue, err := cmd.Flags().GetString(authorFlagName)
		if err != nil {
			return err, false
		}
		m.Author = &authorFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCommitChangesFlag(m *image.ImageCommitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("changes") {

		var changesFlagName string
		if cmdPrefix == "" {
			changesFlagName = "changes"
		} else {
			changesFlagName = fmt.Sprintf("%v.changes", cmdPrefix)
		}

		changesFlagValue, err := cmd.Flags().GetString(changesFlagName)
		if err != nil {
			return err, false
		}
		m.Changes = &changesFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCommitCommentFlag(m *image.ImageCommitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("comment") {

		var commentFlagName string
		if cmdPrefix == "" {
			commentFlagName = "comment"
		} else {
			commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
		}

		commentFlagValue, err := cmd.Flags().GetString(commentFlagName)
		if err != nil {
			return err, false
		}
		m.Comment = &commentFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCommitContainerFlag(m *image.ImageCommitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("container") {

		var containerFlagName string
		if cmdPrefix == "" {
			containerFlagName = "container"
		} else {
			containerFlagName = fmt.Sprintf("%v.container", cmdPrefix)
		}

		containerFlagValue, err := cmd.Flags().GetString(containerFlagName)
		if err != nil {
			return err, false
		}
		m.Container = &containerFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCommitContainerConfigFlag(m *image.ImageCommitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("containerConfig") {
		// Read containerConfig string from cmd and unmarshal
		containerConfigValueStr, err := cmd.Flags().GetString("containerConfig")
		if err != nil {
			return err, false
		}

		containerConfigValue := models.ContainerConfig{}
		if err := json.Unmarshal([]byte(containerConfigValueStr), &containerConfigValue); err != nil {
			return fmt.Errorf("cannot unmarshal containerConfig string in models.ContainerConfig: %v", err), false
		}
		m.ContainerConfig = &containerConfigValue
	}
	containerConfigValueModel := m.ContainerConfig
	if swag.IsZero(containerConfigValueModel) {
		containerConfigValueModel = &models.ContainerConfig{}
	}
	err, added := retrieveModelContainerConfigFlags(0, containerConfigValueModel, "containerConfig", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.ContainerConfig = containerConfigValueModel
	}
	if dryRun && debug {

		containerConfigValueDebugBytes, err := json.Marshal(m.ContainerConfig)
		if err != nil {
			return err, false
		}
		logDebugf("ContainerConfig dry-run payload: %v", string(containerConfigValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationImageImageCommitPauseFlag(m *image.ImageCommitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("pause") {

		var pauseFlagName string
		if cmdPrefix == "" {
			pauseFlagName = "pause"
		} else {
			pauseFlagName = fmt.Sprintf("%v.pause", cmdPrefix)
		}

		pauseFlagValue, err := cmd.Flags().GetBool(pauseFlagName)
		if err != nil {
			return err, false
		}
		m.Pause = &pauseFlagValue

	}
	return nil, retAdded
}
func retrieveOperationImageImageCommitRepoFlag(m *image.ImageCommitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationImageImageCommitTagFlag(m *image.ImageCommitParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationImageImageCommitResult parses request result and return the string content
func parseOperationImageImageCommitResult(resp0 *image.ImageCommitCreated, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*image.ImageCommitCreated)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*image.ImageCommitNotFound)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*image.ImageCommitInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
