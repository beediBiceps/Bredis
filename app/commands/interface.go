
package commands

type Command interface {
	Name() string
	Execute(args []string) (string, error)
}

type Store struct{
    store map[string]string
}