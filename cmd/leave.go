// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var env string
var node string

// leaveCmd represents the leave command
var leaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "Sends a request for node to LEAVE the cluster",
	Long:  `Sends a request for node to LEAVE the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Leave command called ...")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		sections := ParseINI(cfgFile)
		if sections.Section(env).HasKey("nodes") {
			nodes := sections.Section(env).Key("nodes").String()
			nodeandport, err := nodeExists(nodes)
			if err == nil {
				runLeave(nodeandport)
			} else {
				log.Panic(err)
			}
		} else {
			//TODO: This will be shown, even if the Environment is not found.
			log.Panic("Cannot find 'nodes' value from the config file.")
		}
	},
}

func init() {
	RootCmd.AddCommand(leaveCmd)
	leaveCmd.PersistentFlags().StringVar(&env, "env", "", "Environment to contact")
	leaveCmd.PersistentFlags().StringVar(&node, "node", "", "Node name to be used")
}

func nodeExists(nodes string) (string, error) {
	nodesandports := strings.Split(nodes, ",")
	var onlynodenames []string
	for _, nodename := range nodesandports {
		onlynodenames = append(onlynodenames, strings.Split(nodename, ":")[0])
	}
	for i, nodename := range onlynodenames {
		//Check if the passed node exists in the config file.
		if node == nodename {
			return nodesandports[i], nil
		}
	}
	return "", errors.New("No such node in the config file. Exiting!")
}

func runLeave(nodeandport string) {
	// Get the node URL
	split := strings.Split(nodeandport, ":")
	node, port := split[0], split[1]
	nodeURL := GetNodeURL(node, port)
	// TODO: Get the other node URL? Is this needed?
	cmdArgs := []string{node, port, "leave", nodeURL}
	out, err := exec.Command("akka-cluster", cmdArgs...).Output()
	if err == nil {
		fmt.Printf("%s", out)
	}
}
