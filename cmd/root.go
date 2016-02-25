// Copyright © 2016 ravitezu ravi-teja@live.in
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

package cmd

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string
var env string
var node string

//RootCmd - This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "akka-cluster",
	Short: "A wrapper around the actual akka-cluster command.",
	Long: `A wrapper around the actual akka-cluster command.
	Instead of passing the node-hostname, JMX port over the command line.
	This wrapper will read those values from a config(JSON) file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.akka-cluster.ini)")
	RootCmd.PersistentFlags().StringVar(&env, "environment", "stage", "Environment to contact")
	RootCmd.PersistentFlags().StringVar(&node, "node", "stage", "Environment to contact")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// enable ability to specify config file via flag
		cfg, err := ini.Load(cfgFile)
	} else {
		//TODO: Make the default config file path to be accessible easily.
		cfg, err := ini.Load(os.Getenv("HOME") + "/.akka-cluster.ini")
	}
}
