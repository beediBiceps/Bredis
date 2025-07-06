package commands

type EchoCommand struct{}

func (e *EchoCommand) Name() string {
	return "ECHO"
}

func (e *EchoCommand) Execute(args []string) (string, error) {
	if len(args) == 0 {
		return "-ERR wrong number of arguments for 'echo' command\r\n", nil
	}
	return "+" + args[0] + "\r\n", nil
}