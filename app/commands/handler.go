package commands

import (
    "fmt"
    "strings"
)

type CommandRegistry struct {
    cmds map[string]Command
}

func NewCommandRegistry() *CommandRegistry {
    r := &CommandRegistry{cmds: make(map[string]Command)}
    r.Register(&PingCommand{})
    r.Register(&EchoCommand{})
    return r
}

func (r *CommandRegistry) Register(cmd Command) {
    r.cmds[strings.ToUpper(cmd.Name())] = cmd
}

func (r *CommandRegistry) ExecuteCommand(name string, args []string) (string, error) {
    cmd, exists := r.cmds[strings.ToUpper(name)]
    if !exists {
        return "", fmt.Errorf("unknown command: %s", name)
    }
    return cmd.Execute(args)
}
