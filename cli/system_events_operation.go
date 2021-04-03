// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/dockerctl/client/system"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSystemSystemEventsCmd returns a cmd to handle operation systemEvents
func makeOperationSystemSystemEventsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "SystemEvents",
		Short: `Stream real-time events from the server.

Various objects within Docker report events when something happens to them.

Containers report these events: ` + "`" + `attach` + "`" + `, ` + "`" + `commit` + "`" + `, ` + "`" + `copy` + "`" + `, ` + "`" + `create` + "`" + `, ` + "`" + `destroy` + "`" + `, ` + "`" + `detach` + "`" + `, ` + "`" + `die` + "`" + `, ` + "`" + `exec_create` + "`" + `, ` + "`" + `exec_detach` + "`" + `, ` + "`" + `exec_start` + "`" + `, ` + "`" + `exec_die` + "`" + `, ` + "`" + `export` + "`" + `, ` + "`" + `health_status` + "`" + `, ` + "`" + `kill` + "`" + `, ` + "`" + `oom` + "`" + `, ` + "`" + `pause` + "`" + `, ` + "`" + `rename` + "`" + `, ` + "`" + `resize` + "`" + `, ` + "`" + `restart` + "`" + `, ` + "`" + `start` + "`" + `, ` + "`" + `stop` + "`" + `, ` + "`" + `top` + "`" + `, ` + "`" + `unpause` + "`" + `, and ` + "`" + `update` + "`" + `

Images report these events: ` + "`" + `delete` + "`" + `, ` + "`" + `import` + "`" + `, ` + "`" + `load` + "`" + `, ` + "`" + `pull` + "`" + `, ` + "`" + `push` + "`" + `, ` + "`" + `save` + "`" + `, ` + "`" + `tag` + "`" + `, and ` + "`" + `untag` + "`" + `

Volumes report these events: ` + "`" + `create` + "`" + `, ` + "`" + `mount` + "`" + `, ` + "`" + `unmount` + "`" + `, and ` + "`" + `destroy` + "`" + `

Networks report these events: ` + "`" + `create` + "`" + `, ` + "`" + `connect` + "`" + `, ` + "`" + `disconnect` + "`" + `, ` + "`" + `destroy` + "`" + `, ` + "`" + `update` + "`" + `, and ` + "`" + `remove` + "`" + `

The Docker daemon reports these events: ` + "`" + `reload` + "`" + `

Services report these events: ` + "`" + `create` + "`" + `, ` + "`" + `update` + "`" + `, and ` + "`" + `remove` + "`" + `

Nodes report these events: ` + "`" + `create` + "`" + `, ` + "`" + `update` + "`" + `, and ` + "`" + `remove` + "`" + `

Secrets report these events: ` + "`" + `create` + "`" + `, ` + "`" + `update` + "`" + `, and ` + "`" + `remove` + "`" + `

Configs report these events: ` + "`" + `create` + "`" + `, ` + "`" + `update` + "`" + `, and ` + "`" + `remove` + "`" + `
`,
		RunE: runOperationSystemSystemEvents,
	}

	if err := registerOperationSystemSystemEventsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSystemSystemEvents uses cmd flags to call endpoint api
func runOperationSystemSystemEvents(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := system.NewSystemEventsParams()
	if err, _ := retrieveOperationSystemSystemEventsFiltersFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSystemSystemEventsSinceFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSystemSystemEventsUntilFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationSystemSystemEventsResult(appCli.System.SystemEvents(params)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationSystemSystemEventsFiltersFlag(m *system.SystemEventsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("filters") {

		var filtersFlagName string
		if cmdPrefix == "" {
			filtersFlagName = "filters"
		} else {
			filtersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
		}

		filtersFlagValue, err := cmd.Flags().GetString(filtersFlagName)
		if err != nil {
			return err, false
		}
		m.Filters = &filtersFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSystemSystemEventsSinceFlag(m *system.SystemEventsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("since") {

		var sinceFlagName string
		if cmdPrefix == "" {
			sinceFlagName = "since"
		} else {
			sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
		}

		sinceFlagValue, err := cmd.Flags().GetString(sinceFlagName)
		if err != nil {
			return err, false
		}
		m.Since = &sinceFlagValue

	}
	return nil, retAdded
}
func retrieveOperationSystemSystemEventsUntilFlag(m *system.SystemEventsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("until") {

		var untilFlagName string
		if cmdPrefix == "" {
			untilFlagName = "until"
		} else {
			untilFlagName = fmt.Sprintf("%v.until", cmdPrefix)
		}

		untilFlagValue, err := cmd.Flags().GetString(untilFlagName)
		if err != nil {
			return err, false
		}
		m.Until = &untilFlagValue

	}
	return nil, retAdded
}

// printOperationSystemSystemEventsResult prints output to stdout
func printOperationSystemSystemEventsResult(resp0 *system.SystemEventsOK, respErr error) error {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*system.SystemEventsOK)
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
		resp1, ok := iResp1.(*system.SystemEventsBadRequest)
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
		resp2, ok := iResp2.(*system.SystemEventsInternalServerError)
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

// registerOperationSystemSystemEventsParamFlags registers all flags needed to fill params
func registerOperationSystemSystemEventsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSystemSystemEventsFiltersParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSystemSystemEventsSinceParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSystemSystemEventsUntilParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSystemSystemEventsFiltersParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	filtersDescription := `A JSON encoded value of filters (a ` + "`" + `map[string][]string` + "`" + `) to process on the event list. Available filters:

- ` + "`" + `config=<string>` + "`" + ` config name or ID
- ` + "`" + `container=<string>` + "`" + ` container name or ID
- ` + "`" + `daemon=<string>` + "`" + ` daemon name or ID
- ` + "`" + `event=<string>` + "`" + ` event type
- ` + "`" + `image=<string>` + "`" + ` image name or ID
- ` + "`" + `label=<string>` + "`" + ` image or container label
- ` + "`" + `network=<string>` + "`" + ` network name or ID
- ` + "`" + `node=<string>` + "`" + ` node ID
- ` + "`" + `plugin` + "`" + `=<string> plugin name or ID
- ` + "`" + `scope` + "`" + `=<string> local or swarm
- ` + "`" + `secret=<string>` + "`" + ` secret name or ID
- ` + "`" + `service=<string>` + "`" + ` service name or ID
- ` + "`" + `type=<string>` + "`" + ` object to filter by, one of ` + "`" + `container` + "`" + `, ` + "`" + `image` + "`" + `, ` + "`" + `volume` + "`" + `, ` + "`" + `network` + "`" + `, ` + "`" + `daemon` + "`" + `, ` + "`" + `plugin` + "`" + `, ` + "`" + `node` + "`" + `, ` + "`" + `service` + "`" + `, ` + "`" + `secret` + "`" + ` or ` + "`" + `config` + "`" + `
- ` + "`" + `volume=<string>` + "`" + ` volume name
`

	var filtersFlagName string
	if cmdPrefix == "" {
		filtersFlagName = "filters"
	} else {
		filtersFlagName = fmt.Sprintf("%v.filters", cmdPrefix)
	}

	var filtersFlagDefault string

	_ = cmd.PersistentFlags().String(filtersFlagName, filtersFlagDefault, filtersDescription)

	return nil
}
func registerOperationSystemSystemEventsSinceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sinceDescription := `Show events created since this timestamp then stream new events.`

	var sinceFlagName string
	if cmdPrefix == "" {
		sinceFlagName = "since"
	} else {
		sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
	}

	var sinceFlagDefault string

	_ = cmd.PersistentFlags().String(sinceFlagName, sinceFlagDefault, sinceDescription)

	return nil
}
func registerOperationSystemSystemEventsUntilParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	untilDescription := `Show events created until this timestamp then stop streaming.`

	var untilFlagName string
	if cmdPrefix == "" {
		untilFlagName = "until"
	} else {
		untilFlagName = fmt.Sprintf("%v.until", cmdPrefix)
	}

	var untilFlagDefault string

	_ = cmd.PersistentFlags().String(untilFlagName, untilFlagDefault, untilDescription)

	return nil
}

// register flags to command
func registerModelSystemEventsOKBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSystemEventsOKBodyAction(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemEventsOKBodyActor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemEventsOKBodyType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemEventsOKBodyTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemEventsOKBodyTimeNano(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSystemEventsOKBodyAction(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	actionDescription := `The type of event`

	var actionFlagName string
	if cmdPrefix == "" {
		actionFlagName = "Action"
	} else {
		actionFlagName = fmt.Sprintf("%v.Action", cmdPrefix)
	}

	var actionFlagDefault string

	_ = cmd.PersistentFlags().String(actionFlagName, actionFlagDefault, actionDescription)

	return nil
}

func registerSystemEventsOKBodyActor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var actorFlagName string
	if cmdPrefix == "" {
		actorFlagName = "Actor"
	} else {
		actorFlagName = fmt.Sprintf("%v.Actor", cmdPrefix)
	}

	registerModelSystemEventsOKBodyFlags(depth+1, actorFlagName, cmd)

	return nil
}

func registerSystemEventsOKBodyType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `The type of object emitting the event`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "Type"
	} else {
		typeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

func registerSystemEventsOKBodyTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timeDescription := `Timestamp of event`

	var timeFlagName string
	if cmdPrefix == "" {
		timeFlagName = "time"
	} else {
		timeFlagName = fmt.Sprintf("%v.time", cmdPrefix)
	}

	var timeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(timeFlagName, timeFlagDefault, timeDescription)

	return nil
}

func registerSystemEventsOKBodyTimeNano(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timeNanoDescription := `Timestamp of event, with nanosecond accuracy`

	var timeNanoFlagName string
	if cmdPrefix == "" {
		timeNanoFlagName = "timeNano"
	} else {
		timeNanoFlagName = fmt.Sprintf("%v.timeNano", cmdPrefix)
	}

	var timeNanoFlagDefault int64

	_ = cmd.PersistentFlags().Int64(timeNanoFlagName, timeNanoFlagDefault, timeNanoDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSystemEventsOKBodyFlags(depth int, m *system.SystemEventsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, actionAdded := retrieveSystemEventsOKBodyActionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || actionAdded

	err, actorAdded := retrieveSystemEventsOKBodyActorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || actorAdded

	err, typeAdded := retrieveSystemEventsOKBodyTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, timeAdded := retrieveSystemEventsOKBodyTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeAdded

	err, timeNanoAdded := retrieveSystemEventsOKBodyTimeNanoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeNanoAdded

	return nil, retAdded
}

func retrieveSystemEventsOKBodyActionFlags(depth int, m *system.SystemEventsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	actionFlagName := fmt.Sprintf("%v.Action", cmdPrefix)
	if cmd.Flags().Changed(actionFlagName) {

		var actionFlagName string
		if cmdPrefix == "" {
			actionFlagName = "Action"
		} else {
			actionFlagName = fmt.Sprintf("%v.Action", cmdPrefix)
		}

		actionFlagValue, err := cmd.Flags().GetString(actionFlagName)
		if err != nil {
			return err, false
		}
		m.Action = actionFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSystemEventsOKBodyActorFlags(depth int, m *system.SystemEventsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	actorFlagName := fmt.Sprintf("%v.Actor", cmdPrefix)
	if cmd.Flags().Changed(actorFlagName) {

		actorFlagValue := &system.SystemEventsOKBody{}
		err, added := retrieveModelSystemEventsOKBodyFlags(depth+1, actorFlagValue, actorFlagName, cmd)
		if err != nil {
			return err, false
		}
		retAdded = retAdded || added
	}
	return nil, retAdded
}

func retrieveSystemEventsOKBodyTypeFlags(depth int, m *system.SystemEventsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	typeFlagName := fmt.Sprintf("%v.Type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "Type"
		} else {
			typeFlagName = fmt.Sprintf("%v.Type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSystemEventsOKBodyTimeFlags(depth int, m *system.SystemEventsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	timeFlagName := fmt.Sprintf("%v.time", cmdPrefix)
	if cmd.Flags().Changed(timeFlagName) {

		var timeFlagName string
		if cmdPrefix == "" {
			timeFlagName = "time"
		} else {
			timeFlagName = fmt.Sprintf("%v.time", cmdPrefix)
		}

		timeFlagValue, err := cmd.Flags().GetInt64(timeFlagName)
		if err != nil {
			return err, false
		}
		m.Time = timeFlagValue

		retAdded = true
	}
	return nil, retAdded
}

func retrieveSystemEventsOKBodyTimeNanoFlags(depth int, m *system.SystemEventsOKBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	timeNanoFlagName := fmt.Sprintf("%v.timeNano", cmdPrefix)
	if cmd.Flags().Changed(timeNanoFlagName) {

		var timeNanoFlagName string
		if cmdPrefix == "" {
			timeNanoFlagName = "timeNano"
		} else {
			timeNanoFlagName = fmt.Sprintf("%v.timeNano", cmdPrefix)
		}

		timeNanoFlagValue, err := cmd.Flags().GetInt64(timeNanoFlagName)
		if err != nil {
			return err, false
		}
		m.TimeNano = timeNanoFlagValue

		retAdded = true
	}
	return nil, retAdded
}

// register flags to command
func registerModelSystemEventsOKBodyActorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSystemEventsOKBodyActorAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSystemEventsOKBodyActorID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSystemEventsOKBodyActorAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}
	// warning: Attributes map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerSystemEventsOKBodyActorID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of the object emitting the event`

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSystemEventsOKBodyActorFlags(depth int, m *system.SystemEventsOKBodyActor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrieveSystemEventsOKBodyActorAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, idAdded := retrieveSystemEventsOKBodyActorIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveSystemEventsOKBodyActorAttributesFlags(depth int, m *system.SystemEventsOKBodyActor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false
	attributesFlagName := fmt.Sprintf("%v.Attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// warning: Attributes map type map[string]string is not supported by go-swagger cli yet
	}
	return nil, retAdded
}

func retrieveSystemEventsOKBodyActorIDFlags(depth int, m *system.SystemEventsOKBodyActor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
