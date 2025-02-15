// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"encoding/json"

	mfclients "github.com/mainflux/mainflux/pkg/clients"
	mfxsdk "github.com/mainflux/mainflux/pkg/sdk/go"
	"github.com/spf13/cobra"
)

var cmdUsers = []cobra.Command{
	{
		Use:   "create <name> <username> <password> <user_auth_token>",
		Short: "Create user",
		Long: "Create user with provided name, username and password. Token in optional\n" +
			"For example:\n" +
			"\tmainflux-cli users create user user@example.com 12345678 $USER_AUTH_TOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 3 || len(args) > 4 {
				logUsage(cmd.Use)
				return
			}
			if len(args) == 3 {
				args = append(args, "")
			}

			user := mfxsdk.User{
				Name: args[0],
				Credentials: mfxsdk.Credentials{
					Identity: args[1],
					Secret:   args[2],
				},
				Status: mfclients.EnabledStatus.String(),
			}
			user, err := sdk.CreateUser(user, args[3])
			if err != nil {
				logError(err)
				return
			}

			logJSON(user)
		},
	},
	{
		Use:   "get [all | <user_id> ] <user_auth_token>",
		Short: "Get users",
		Long: "Get all users or get user by id. Users can be filtered by name or metadata or status\n" +
			"Usage:\n" +
			"\tmainflux-cli users get all <user_auth_token> - lists all users\n" +
			"\tmainflux-cli users get all <user_auth_token> --offset <offset> --limit <limit> - lists all users with provided offset and limit\n" +
			"\tmainflux-cli users get <user_id> <user_auth_token> - shows user with provided <user_id>\n",
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
				Email:    "",
				Offset:   uint64(Offset),
				Limit:    uint64(Limit),
				Metadata: metadata,
				Status:   Status,
			}
			if args[0] == all {
				l, err := sdk.Users(pageMetadata, args[1])
				if err != nil {
					logError(err)
					return
				}
				logJSON(l)
				return
			}
			u, err := sdk.User(args[0], args[1])
			if err != nil {
				logError(err)
				return
			}

			logJSON(u)
		},
	},
	{
		Use:   "token <username> <password>",
		Short: "Get token",
		Long: "Generate new token from username and password\n" +
			"For example:\n" +
			"\tmainflux-cli users token user@example.com 12345678\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				logUsage(cmd.Use)
				return
			}

			user := mfxsdk.User{
				Credentials: mfxsdk.Credentials{
					Identity: args[0],
					Secret:   args[1],
				},
			}
			token, err := sdk.CreateToken(user)
			if err != nil {
				logError(err)
				return
			}

			logJSON(token)

		},
	},
	{
		Use:   "refreshtoken <token>",
		Short: "Get token",
		Long: "Generate new token from refresh token\n" +
			"For example:\n" +
			"\tmainflux-cli users refreshtoken <refresh_token>\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				logUsage(cmd.Use)
				return
			}

			token, err := sdk.RefreshToken(args[0])
			if err != nil {
				logError(err)
				return
			}

			logJSON(token)

		},
	},
	{
		Use:   "update [<user_id> <JSON_string> | tags <user_id> <tags> | identity <user_id> <identity> | owner <user_id> <owner>] <user_auth_token>",
		Short: "Update user",
		Long: "Updates either user name and metadata or user tags or user identity or user owner\n" +
			"Usage:\n" +
			"\tmainflux-cli users update <user_id> '{\"name\":\"new name\", \"metadata\":{\"key\": \"value\"}}' $USERTOKEN - updates user name and metadata\n" +
			"\tmainflux-cli users update tags <user_id> '[\"tag1\", \"tag2\"]' $USERTOKEN - updates user tags\n" +
			"\tmainflux-cli users update identity <user_id> newidentity@example.com $USERTOKEN - updates user identity\n" +
			"\tmainflux-cli users update owner <user_id> <owner_id> $USERTOKEN - updates user owner\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 4 && len(args) != 3 {
				logUsage(cmd.Use)
				return
			}

			var user mfxsdk.User
			if args[0] == "tags" {
				if err := json.Unmarshal([]byte(args[2]), &user.Tags); err != nil {
					logError(err)
					return
				}
				user.ID = args[1]
				user, err := sdk.UpdateUserTags(user, args[3])
				if err != nil {
					logError(err)
					return
				}

				logJSON(user)
				return
			}

			if args[0] == "identity" {
				user.ID = args[1]
				user.Credentials.Identity = args[2]
				user, err := sdk.UpdateUserIdentity(user, args[3])
				if err != nil {
					logError(err)
					return
				}

				logJSON(user)
				return

			}

			if args[0] == "owner" {
				user.ID = args[1]
				user.Owner = args[2]
				user, err := sdk.UpdateUserOwner(user, args[3])
				if err != nil {
					logError(err)
					return
				}

				logJSON(user)
				return

			}

			if err := json.Unmarshal([]byte(args[1]), &user); err != nil {
				logError(err)
				return
			}
			user.ID = args[0]
			user, err := sdk.UpdateUser(user, args[2])
			if err != nil {
				logError(err)
				return
			}

			logJSON(user)
		},
	},
	{
		Use:   "profile <user_auth_token>",
		Short: "Get user profile",
		Long: "Get user profile\n" +
			"Usage:\n" +
			"\tmainflux-cli users profile $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				logUsage(cmd.Use)
				return
			}

			user, err := sdk.UserProfile(args[0])
			if err != nil {
				logError(err)
				return
			}

			logJSON(user)
		},
	},
	{
		Use:   "password <old_password> <password> <user_auth_token>",
		Short: "Update password",
		Long: "Update password\n" +
			"Usage:\n" +
			"\tmainflux-cli users password old_password new_password $USERTOKEN\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 3 {
				logUsage(cmd.Use)
				return
			}

			user, err := sdk.UpdatePassword(args[0], args[1], args[2])
			if err != nil {
				logError(err)
				return
			}

			logJSON(user)
		},
	},
	{
		Use:   "enable <user_id> <user_auth_token>",
		Short: "Change user status to enabled",
		Long: "Change user status to enabled\n" +
			"Usage:\n" +
			"\tmainflux-cli users enable <user_id> <user_auth_token>\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				logUsage(cmd.Use)
				return
			}

			user, err := sdk.EnableUser(args[0], args[1])
			if err != nil {
				logError(err)
				return
			}

			logJSON(user)
		},
	},
	{
		Use:   "disable <user_id> <user_auth_token>",
		Short: "Change user status to disabled",
		Long: "Change user status to disabled\n" +
			"Usage:\n" +
			"\tmainflux-cli users disable <user_id> <user_auth_token>\n",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				logUsage(cmd.Use)
				return
			}

			user, err := sdk.DisableUser(args[0], args[1])
			if err != nil {
				logError(err)
				return
			}

			logJSON(user)
		},
	},
}

// NewUsersCmd returns users command.
func NewUsersCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "users [create | get | update | token | password | enable | disable]",
		Short: "Users management",
		Long:  `Users management: create accounts and tokens"`,
	}

	for i := range cmdUsers {
		cmd.AddCommand(&cmdUsers[i])
	}

	return &cmd
}
