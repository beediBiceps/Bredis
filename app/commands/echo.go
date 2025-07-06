package commands

import "strings"

type EchoCommand struct{}

func (c *EchoCommand) Name() string {
	return "echo"
}

func (c *EchoCommand) Execute(args []string) string {
	return strings.Join(args, " ")
}
