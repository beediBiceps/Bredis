package commands

import "fmt"

type GetCommand struct{
    store *Store
}

func (g *GetCommand) Name() string {
    return "GET"
}
    
func (g *GetCommand) Execute(args []string) (string, error) {
    if len(args) != 1 {
        return "", fmt.Errorf("wrong number of arguments for 'get' command")
    }
    value, err := g.store.Get(args[0])
    if err != nil {
        return "", err
    }
    if value == "" {
        return "$-1\r\n", nil
    }
    return fmt.Sprintf("$%d\r\n%s\r\n", len(value), value), nil
}