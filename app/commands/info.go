package commands

import "fmt"

type InfoCommand struct{}

func (i *InfoCommand) Name() string {
	return "INFO"
}

func (i *InfoCommand) Execute(flag string, args []string) (string, error) {
	return "+OK\r\n", nil
}

