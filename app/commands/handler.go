package commands

import (
	"fmt"
	"strings"
)

type CommandHandler struct {
	cmds map[string]Command
}

func NewCommandHandler() *CommandHandler {
	r := &CommandHandler{cmds: make(map[string]Command)}
	store := NewStore()
	r.Register(&PingCommand{})
	r.Register(&EchoCommand{})
	r.Register(&SetCommand{store: store})
	r.Register(&GetCommand{store: store})
	return r
}

func (r *CommandHandler) Register(cmd Command) {
	r.cmds[strings.ToUpper(cmd.Name())] = cmd
}

func (r *CommandHandler) ExecuteCommand(name string, args []string) (string, error) {
	cmd, exists := r.cmds[strings.ToUpper(name)]
	if !exists {
		return "", fmt.Errorf("unknown command: %s", name)
	}
	return cmd.Execute(args)
}