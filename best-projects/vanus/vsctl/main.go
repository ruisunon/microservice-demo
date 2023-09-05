// Copyright 2023 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// vsctl is a command line application that controls vanus.
package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/vanus-labs/vanus/vsctl/command"
)

const (
	cliName        = "vsctl"
	cliDescription = "the command-line application for vanus"
)

var (
	globalFlags = command.GlobalFlags{}
	rootCmd     = &cobra.Command{
		Use:        cliName,
		Short:      cliDescription,
		SuggestFor: []string{"vsctl"},
	}
)

func init() {
	cobra.EnablePrefixMatching = true
	cobra.EnableCommandSorting = false
	rootCmd.PersistentFlags().StringVar(&globalFlags.Endpoint, "endpoint",
		"127.0.0.1:8080", "the endpoints of vanus controller")
	rootCmd.PersistentFlags().StringVar(&globalFlags.OperatorEndpoint, "operator-endpoint",
		"127.0.0.1:8080", "the endpoints of vanus operator")
	rootCmd.PersistentFlags().StringVarP(&globalFlags.ConfigFile, "config", "C",
		"~/.vanus/vanus.yml", "the config file of vsctl")
	rootCmd.PersistentFlags().BoolVarP(&globalFlags.Debug, "debug", "D", false,
		"is debug mode enable")
	rootCmd.PersistentFlags().StringVar(&globalFlags.Format, "format", "table",
		"the output format of vsctl, json or table")
	rootCmd.PersistentFlags().StringVar(&globalFlags.Token, "token", "admin",
		"the user token")

	if os.Getenv("VANUS_TOKEN") != "" {
		globalFlags.Token = os.Getenv("VANUS_TOKEN")
	}

	if os.Getenv("VANUS_GATEWAY") != "" {
		globalFlags.Endpoint = os.Getenv("VANUS_GATEWAY")
	}

	if os.Getenv("VANUS_OPERATOR") != "" {
		globalFlags.OperatorEndpoint = os.Getenv("VANUS_OPERATOR")
	}

	rootCmd.AddCommand(
		command.NewEventCommand(),
		command.NewEventbusCommand(),
		command.NewSubscriptionCommand(),
		command.NewClusterCommand(),
		command.NewConnectorCommand(),
		command.NewNamespaceCommand(),
		command.NewUserCommand(),
		command.NewPermissionCommand(),
		newVersionCommand(),
	)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	MustStart()
}

func Start() error {
	return rootCmd.Execute()
}

func MustStart() {
	if err := Start(); err != nil {
		color.Red("vsctl run error: %s", err)
		os.Exit(-1)
	}
}
