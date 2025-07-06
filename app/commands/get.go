package commands

import "fmt"

type GetCommand struct{
    store Store
}

func (g *GetCommand) Name() string {
    return "GET"
}
    
func (g *GetCommand) Execute(args []string) (string, error) {
    if len(args) != 1 {
        return "", fmt.Errorf("wrong number of arguments for 'get' command")
    }
    return g.store.Get(args[0])
}
