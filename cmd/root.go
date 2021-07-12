/*
Copyright © 2021 Srihari Vishnu srihari.vishnu@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
func NewRootCmd(cli *client.Client) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dockbox",
		Short: "Try out code without creating any side effects!",
		Long: `
Usage: dockbox [OPTIONS] COMMAND

Manage workspaces and dependencies with ease in an isolated, secure environment.
	
To get started with dockbox, try entering:

	dockbox create <url>`,
	}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dockbox.yaml)")

	rootCmd.AddCommand(
		NewCleanCommand(cli),
		NewCreateCommand(cli),
		NewEnterCommand(cli),
		NewListCommand(cli),
		NewTreeCommand(cli),
	)
	return rootCmd
}

func Execute() {
	cli, err := client.NewClientWithOpts()
	CheckError(err)
	rootCmd := NewRootCmd(cli)
	CheckError(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".dockbox" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dockbox")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
