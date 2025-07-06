package commands

import "fmt"

type SetCommand struct{
    store *Store
}

func (s *SetCommand) Name() string {
    return "SET"
}

func (s *SetCommand) Execute(args []string) (string, error) {
    if len(args) != 2 {
        return "", fmt.Errorf("wrong number of arguments for 'set' command")
    }
    fmt.Println(args)
    s.store.Set(args[0], args[1])
    return "+OK\r\n", nil
}