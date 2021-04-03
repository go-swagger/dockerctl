// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/image"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationImageImageHistoryCmd returns a cmd to handle operation imageHistory
func makeOperationImageImageHistoryCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "ImageHistory",
		Short: `Return parent layers of an image.`,
		RunE:  runOperationImageImageHistory,
	}

	if err := registerOperationImageImageHistoryParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationImageImageHistory uses cmd flags to call endpoint api
func runOperationImageImageHistory(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := image.NewImageHistoryParams()
	if err, _ := retrieveOperationImageImageHistoryNameFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationImageImageHistoryResult(appCli.Image.ImageHistory(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationImageImageHistoryNameFlag(m *image.ImageHistoryParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// printOperationImageImageHistoryResult prints output to stdout
func printOperationImageImageHistoryResult(resp0 *image.ImageHistoryOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*image.ImageHistoryOK)
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
		resp1, ok := iResp1.(*image.ImageHistoryNotFound)
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
		resp2, ok := iResp2.(*image.ImageHistoryInternalServerError)
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

// registerOperationImageImageHistoryParamFlags registers all flags needed to fill params
func registerOperationImageImageHistoryParamFlags(cmd *cobra.Command) error {
	if err := registerOperationImageImageHistoryNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationImageImageHistoryNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	nameDescription := `Required. Image name or ID`

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

// register flags to command
func registerModelHistoryResponseItemFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerHistoryResponseItemComment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerHistoryResponseItemCreated(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerHistoryResponseItemCreatedBy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerHistoryResponseItemID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerHistoryResponseItemSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerHistoryResponseItemTags(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerHistoryResponseItemComment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	commentDescription := `Required. `

	var commentFlagName string
	if cmdPrefix == "" {
		commentFlagName = "Comment"
	} else {
		commentFlagName = fmt.Sprintf("%v.Comment", cmdPrefix)
	}

	var commentFlagDefault string

	_ = cmd.PersistentFlags().String(commentFlagName, commentFlagDefault, commentDescription)

	return nil
}

func registerHistoryResponseItemCreated(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdDescription := `Required. `

	var createdFlagName string
	if cmdPrefix == "" {
		createdFlagName = "Created"
	} else {
		createdFlagName = fmt.Sprintf("%v.Created", cmdPrefix)
	}

	var createdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(createdFlagName, createdFlagDefault, createdDescription)

	return nil
}

func registerHistoryResponseItemCreatedBy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	createdByDescription := `Required. `

	var createdByFlagName string
	if cmdPrefix == "" {
		createdByFlagName = "CreatedBy"
	} else {
		createdByFlagName = fmt.Sprintf("%v.CreatedBy", cmdPrefix)
	}

	var createdByFlagDefault string

	_ = cmd.PersistentFlags().String(createdByFlagName, createdByFlagDefault, createdByDescription)

	return nil
}

func registerHistoryResponseItemID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `Required. `

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

func registerHistoryResponseItemSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	sizeDescription := `Required. `

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "Size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
	}

	var sizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}

func registerHistoryResponseItemTags(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Tags []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelHistoryResponseItemFlags(depth int, m *image.HistoryResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, commentAdded := retrieveHistoryResponseItemCommentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || commentAdded

	err, createdAdded := retrieveHistoryResponseItemCreatedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdAdded

	err, createdByAdded := retrieveHistoryResponseItemCreatedByFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || createdByAdded

	err, idAdded := retrieveHistoryResponseItemIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, sizeAdded := retrieveHistoryResponseItemSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sizeAdded

	err, tagsAdded := retrieveHistoryResponseItemTagsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tagsAdded

	return nil, retAdded
}

func retrieveHistoryResponseItemCommentFlags(depth int, m *image.HistoryResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	commentFlagName := fmt.Sprintf("%v.Comment", cmdPrefix)
	if cmd.Flags().Changed(commentFlagName) {

		var commentFlagName string
		if cmdPrefix == "" {
			commentFlagName = "Comment"
		} else {
			commentFlagName = fmt.Sprintf("%v.Comment", cmdPrefix)
		}

		commentFlagValue, err := cmd.Flags().GetString(commentFlagName)
		if err != nil {
			return err, false
		}
		m.Comment = commentFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveHistoryResponseItemCreatedFlags(depth int, m *image.HistoryResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

		createdFlagValue, err := cmd.Flags().GetInt64(createdFlagName)
		if err != nil {
			return err, false
		}
		m.Created = createdFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveHistoryResponseItemCreatedByFlags(depth int, m *image.HistoryResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	createdByFlagName := fmt.Sprintf("%v.CreatedBy", cmdPrefix)
	if cmd.Flags().Changed(createdByFlagName) {

		var createdByFlagName string
		if cmdPrefix == "" {
			createdByFlagName = "CreatedBy"
		} else {
			createdByFlagName = fmt.Sprintf("%v.CreatedBy", cmdPrefix)
		}

		createdByFlagValue, err := cmd.Flags().GetString(createdByFlagName)
		if err != nil {
			return err, false
		}
		m.CreatedBy = createdByFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveHistoryResponseItemIDFlags(depth int, m *image.HistoryResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveHistoryResponseItemSizeFlags(depth int, m *image.HistoryResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	sizeFlagName := fmt.Sprintf("%v.Size", cmdPrefix)
	if cmd.Flags().Changed(sizeFlagName) {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "Size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.Size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetInt64(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = sizeFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveHistoryResponseItemTagsFlags(depth int, m *image.HistoryResponseItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	tagsFlagName := fmt.Sprintf("%v.Tags", cmdPrefix)
	if cmd.Flags().Changed(tagsFlagName) {
		// warning: Tags array type []string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}
