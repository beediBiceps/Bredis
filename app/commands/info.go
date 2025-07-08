package commands

import (
	"fmt"
)

type InfoCommand struct{}

func (i *InfoCommand) Name() string {
	return "INFO"
}

func (i *InfoCommand) Execute(args []string) (string, error) {
	response := fmt.Sprintf("$%d\r\n%s\r\n", len(info), info)
	return response, nil
}
