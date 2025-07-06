package commands

type PingCommand struct{}

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Execute(args []string) (string, error) {
	return "+PONG\r\n", nil
}
