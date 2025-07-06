package commands

import (
    "fmt"
    "strconv"
)

type SetCommand struct{
    store *Store
}

func (s *SetCommand) Name() string {
    return "SET"
}

func (s *SetCommand) Execute(args []string) (string, error) {
    fmt.Println(args)

    if len(args)>2{
        expiryMs, err := strconv.Atoi(args[3])
        if err != nil {
            return "", fmt.Errorf("wrong number of arguments for 'set' command")
        }
        s.store.SetWithExpiry(args[0], args[1], int64(expiryMs))
        return "+OK\r\n", nil
    }
    s.store.Set(args[0], args[1])
    return "+OK\r\n", nil
}