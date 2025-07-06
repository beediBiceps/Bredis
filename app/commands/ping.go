package commands

type PingCommand struct{}

func (p *PingCommand) Name() string {
	return "PING"
}

func (p *PingCommand) Execute(args []string) (string, error) {
	if len(args) == 0 {
		return "+PONG\r\n", nil
	}
	return "+" + args[0] + "\r\n", nil
}