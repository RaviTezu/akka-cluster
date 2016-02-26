package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/go-ini/ini"
)

//ParseINI - To parse the INI file and return the config.
func ParseINI(filename string) *ini.File {
	conf, err := ini.Load(filename)
	if err != nil {
		log.Panic(err)
	}
	return conf
}

func checkAkkaCLI() bool {
	out, err := exec.Command("which", "akka-cluster").Output()
	if err == nil {
		fmt.Printf("Found the akka-cluster binary here: %s", out)
		return true
	}
	return false
}

var status Status

// Member struct - which will be used to match against cluster-status output.
type Member struct {
	Address string   `json:"address"`
	Status  string   `json:"status"`
	Roles   []string `json:"roles"`
}

// Status struct - which will be used to match againstthe cluster-status output.
type Status struct {
	Selfaddress string   `json:"self-address"`
	Members     []Member `json:"members"`
	Unreachable []string `json:"unreachable"`
}

// GetNodeURL - Takes in a passed node and returns the node URL to be used for akka-cluster
func GetNodeURL(node string, port string) string {
	if checkAkkaCLI() {
		cmdArgs := []string{node, port, "cluster-status"}
		out, err := exec.Command("akka-cluster", cmdArgs...).Output()
		//Parse the output of cluster-status and get the node URL.
		if err == nil {
			json.Unmarshal([]byte(strings.TrimLeft(string(out), "Querying cluster status\n")), &status)
		} else {
			log.Panic(err)
		}
	} else {
		log.Panic("Cannot find akka-cluster binary on the machine, Please add it to $PATH.")
	}
	return status.Selfaddress
}

//GetOtherNodeURL - Takes a node and then returns the other node URL in the cluster.
func GetOtherNodeURL(node string) string {
	return "none for now"
}
