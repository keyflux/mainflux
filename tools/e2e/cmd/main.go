// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Package main contains e2e tool for testing Mainflux.
package main

import (
	"log"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/mainflux/mainflux/tools/e2e"
	"github.com/spf13/cobra"
)

const defNum = uint64(10)

func main() {
	econf := e2e.Config{}

	var rootCmd = &cobra.Command{
		Use:   "e2e",
		Short: "e2e is end-to-end testing tool for Mainflux",
		Long: "Tool for testing end-to-end flow of mainflux by doing a couple of operations namely:\n" +
			"1. Creating, viewing, updating and changing status of users, groups, things and channels.\n" +
			"2. Connecting users and groups to each other and things and channels to each other.\n" +
			"3. Sending messages from things to channels on all 4 protocol adapters (HTTP, WS, CoAP and MQTT).\n" +
			"Complete documentation is available at https://docs.mainflux.io",
		Example: "Here is a simple example of using e2e tool.\n" +
			"Use the following commands from the root mainflux directory:\n\n" +
			"go run tools/e2e/cmd/main.go\n" +
			"go run tools/e2e/cmd/main.go --host 142.93.118.47\n" +
			"go run tools/e2e/cmd/main.go --host localhost --num 10 --num_of_messages 100 --prefix e2e",
		Run: func(_ *cobra.Command, _ []string) {
			e2e.Test(econf)
		},
	}

	cc.Init(&cc.Config{
		RootCmd:       rootCmd,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		CmdShortDescr: cc.Magenta,
		Example:       cc.Italic + cc.Magenta,
		ExecName:      cc.Bold,
		Flags:         cc.HiGreen + cc.Bold,
		FlagsDescr:    cc.Green,
		FlagsDataType: cc.White + cc.Italic,
	})

	// Root Flags
	rootCmd.PersistentFlags().StringVarP(&econf.Host, "host", "H", "localhost", "address for a running mainflux instance")
	rootCmd.PersistentFlags().StringVarP(&econf.Prefix, "prefix", "p", "", "name prefix for users, groups, things and channels")
	rootCmd.PersistentFlags().Uint64VarP(&econf.Num, "num", "n", defNum, "number of users, groups, channels and things to create and connect")
	rootCmd.PersistentFlags().Uint64VarP(&econf.NumOfMsg, "num_of_messages", "N", defNum, "number of messages to send")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
