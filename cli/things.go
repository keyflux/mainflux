// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"encoding/json"

	mfclients "github.com/mainflux/mainflux/pkg/clients"
	mfxsdk "github.com/mainflux/mainflux/pkg/sdk/go"
	"github.com/spf13/cobra"
)

var cmdThings = []cobra.Command{
	{
		Use:   "create <JSON_thing> <user_auth_token>",
		Short: "Create thing",
		Long: "Creates new thing with provided name and metadata\n" +
			"Usage:\n" +
			"\tmainflux-cli things create '{\"name\":\"new thing\", \"metadata\":{\"key\": \"value\"}}' $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				logUsage(cmd.Use)
				return
			}

			var thing mfxsdk.Thing
			if err := json.Unmarshal([]byte(args[0]), &thing); err != nil {
				logError(err)
				return
			}
			thing.Status = mfclients.EnabledStatus.String()
			thing, err := sdk.CreateThing(thing, args[1])
			if err != nil {
				logError(err)
				return
			}

			logJSON(thing)
		},
	},
	{
		Use:   "get [all | <thing_id>] <user_auth_token>",
		Short: "Get things",
		Long: "Get all things or get thing by id. Things can be filtered by name or metadata\n" +
			"Usage:\n" +
			"\tmainflux-cli things get all $USERTOKEN - lists all things\n" +
			"\tmainflux-cli things get all $USERTOKEN --offset=10 --limit=10 - lists all things with offset and limit\n" +
			"\tmainflux-cli things get <thing_id> $USERTOKEN - shows thing with provided <thing_id>\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				logUsage(cmd.Use)
				return
			}
			metadata, err := convertMetadata(Metadata)
			if err != nil {
				logError(err)
				return
			}
			pageMetadata := mfxsdk.PageMetadata{
				Name:     "",
				Offset:   uint64(Offset),
				Limit:    uint64(Limit),
				Metadata: metadata,
			}
			if args[0] == all {
				l, err := sdk.Things(pageMetadata, args[1])
				if err != nil {
					logError(err)
					return
				}
				logJSON(l)
				return
			}
			t, err := sdk.Thing(args[0], args[1])
			if err != nil {
				logError(err)
				return
			}

			logJSON(t)
		},
	},
	{
		Use:   "identify <thing_key>",
		Short: "Identify thing",
		Long: "Validates thing's key and returns its ID\n" +
			"Usage:\n" +
			"\tmainflux-cli things identify <thing_key>\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				logUsage(cmd.Use)
				return
			}

			i, err := sdk.IdentifyThing(args[0])
			if err != nil {
				logError(err)
				return
			}

			logJSON(i)
		},
	},
	{
		Use:   "update [<thing_id> <JSON_string> | tags <thing_id> <tags> | secret <thing_id> <secret> | owner <thing_id> <owner> ] <user_auth_token>",
		Short: "Update thing",
		Long: "Updates thing with provided id, name and metadata, or updates thing tags, secret or owner\n" +
			"Usage:\n" +
			"\tmainflux-cli things update <thing_id> '{\"name\":\"new name\", \"metadata\":{\"key\": \"value\"}}' $USERTOKEN\n" +
			"\tmainflux-cli things update tags <thing_id> '{\"tag1\":\"value1\", \"tag2\":\"value2\"}' $USERTOKEN\n" +
			"\tmainflux-cli things update secret <thing_id> newsecret $USERTOKEN\n" +
			"\tmainflux-cli things update owner <thing_id> <owner_id> $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 4 && len(args) != 3 {
				logUsage(cmd.Use)
				return
			}

			var thing mfxsdk.Thing
			if args[0] == "tags" {
				if err := json.Unmarshal([]byte(args[2]), &thing.Tags); err != nil {
					logError(err)
					return
				}
				thing.ID = args[1]
				thing, err := sdk.UpdateThingTags(thing, args[3])
				if err != nil {
					logError(err)
					return
				}

				logJSON(thing)
				return
			}

			if args[0] == "secret" {
				thing, err := sdk.UpdateThingSecret(args[1], args[2], args[3])
				if err != nil {
					logError(err)
					return
				}

				logJSON(thing)
				return
			}

			if args[0] == "owner" {
				thing.ID = args[1]
				thing.Owner = args[2]
				thing, err := sdk.UpdateThingOwner(thing, args[3])
				if err != nil {
					logError(err)
					return
				}

				logJSON(thing)
				return
			}

			if err := json.Unmarshal([]byte(args[1]), &thing); err != nil {
				logError(err)
				return
			}
			thing.ID = args[0]
			thing, err := sdk.UpdateThing(thing, args[2])
			if err != nil {
				logError(err)
				return
			}

			logJSON(thing)
		},
	},
	{
		Use:   "enable <thing_id> <user_auth_token>",
		Short: "Change thing status to enabled",
		Long: "Change thing status to enabled\n" +
			"Usage:\n" +
			"\tmainflux-cli things enable <thing_id> $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				logUsage(cmd.Use)
				return
			}

			thing, err := sdk.EnableThing(args[0], args[1])
			if err != nil {
				logError(err)
				return
			}

			logJSON(thing)
		},
	},
	{
		Use:   "disable <thing_id> <user_auth_token>",
		Short: "Change thing status to disabled",
		Long: "Change thing status to disabled\n" +
			"Usage:\n" +
			"\tmainflux-cli things disable <thing_id> $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				logUsage(cmd.Use)
				return
			}

			thing, err := sdk.DisableThing(args[0], args[1])
			if err != nil {
				logError(err)
				return
			}

			logJSON(thing)
		},
	},
	{
		Use:   "share <channel_id> <user_id> <allowed_actions> <user_auth_token>",
		Short: "Share thing with a user",
		Long: "Share thing with a user\n" +
			"Usage:\n" +
			"\tmainflux-cli things share <channel_id> <user_id> '[\"c_list\", \"c_delete\"]' $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 4 {
				logUsage(cmd.Use)
				return
			}
			var actions []string
			if err := json.Unmarshal([]byte(args[2]), &actions); err != nil {
				logError(err)
				return
			}

			err := sdk.ShareThing(args[0], args[1], actions, args[3])
			if err != nil {
				logError(err)
				return
			}

			logOK()
		},
	},
	{
		Use:   "connect <thing_id> <channel_id> <user_auth_token>",
		Short: "Connect thing",
		Long: "Connect thing to the channel\n" +
			"Usage:\n" +
			"\tmainflux-cli things connect <thing_id> <channel_id> $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 3 {
				logUsage(cmd.Use)
				return
			}

			connIDs := mfxsdk.ConnectionIDs{
				ChannelIDs: []string{args[1]},
				ThingIDs:   []string{args[0]},
			}
			if err := sdk.Connect(connIDs, args[2]); err != nil {
				logError(err)
				return
			}

			logOK()
		},
	},
	{
		Use:   "disconnect <thing_id> <channel_id> <user_auth_token>",
		Short: "Disconnect thing",
		Long: "Disconnect thing to the channel\n" +
			"Usage:\n" +
			"\tmainflux-cli things disconnect <thing_id> <channel_id> $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 3 {
				logUsage(cmd.Use)
				return
			}

			connIDs := mfxsdk.ConnectionIDs{
				ThingIDs:   []string{args[0]},
				ChannelIDs: []string{args[1]},
			}
			if err := sdk.Disconnect(connIDs, args[2]); err != nil {
				logError(err)
				return
			}

			logOK()
		},
	},
	{
		Use:   "connections <thing_id> <user_auth_token>",
		Short: "Connected list",
		Long: "List of Channels connected to Thing\n" +
			"Usage:\n" +
			"\tmainflux-cli connections <thing_id> $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				logUsage(cmd.Use)
				return
			}
			pm := mfxsdk.PageMetadata{
				Offset: uint64(Offset),
				Limit:  uint64(Limit),
			}
			cl, err := sdk.ChannelsByThing(args[0], pm, args[1])
			if err != nil {
				logError(err)
				return
			}

			logJSON(cl)
		},
	},
}

// NewThingsCmd returns things command.
func NewThingsCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "things [create | get | update | delete | share | connect | disconnect | connections | not-connected]",
		Short: "Things management",
		Long:  `Things management: create, get, update, delete or share Thing, connect or disconnect Thing from Channel and get the list of Channels connected or disconnected from a Thing`,
	}

	for i := range cmdThings {
		cmd.AddCommand(&cmdThings[i])
	}

	return &cmd
}
