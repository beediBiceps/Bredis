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
    response := ""
	if args[0] == "replication"{
		role := cfg.GetRole()
        response = fmt.Sprintf("role:%s\r\n", role)
	}
	return "+" + response + "\r\n", nil
}