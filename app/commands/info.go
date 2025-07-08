// commands/info.go
package commands

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/config"
)

type InfoCommand struct{}

func (i *InfoCommand) Name() string {
	return "INFO"
}

func (i *InfoCommand) Execute(args []string) (string, error) {
	cfg := config.GetConfig()
	
	if cfg == nil {
		return "-ERR Server not properly initialized\r\n", nil
	}

	clusterInfo := cfg.GetClusterInfo()
	role := cfg.GetRole()
	

	response := fmt.Sprintf("# Replication\r\nrole:%s\r\n", role)
	
	if role == "master" {
		response += fmt.Sprintf("connected_slaves:%d\r\n", clusterInfo.connectedSlaves)
		response += fmt.Sprintf("master_replid:%s\r\n", clusterInfo.masterReplID)
		response += fmt.Sprintf("master_repl_offset:%d\r\n", clusterInfo.masterReplOffset)
	}
	
	return "+" + response + "\r\n", nil
}