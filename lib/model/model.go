package model

type Config struct {
	Commands []Command
}

type Command struct {
	BaseCommand  string
	Services     []string
	Environments []Environment
}

type Environment struct {
	Name               string
	PreServiceCommand  string
	PostServiceCommand string
	Confirm            bool
}

func (c *Config) GetCommandByService(serviceName string) *Command {
	for _, command := range c.Commands {
		for _, service := range command.Services {
			if service == serviceName {
				return &command
			}
		}
	}
	return nil
}

func (c *Command) GetEnvironmentsString() string {
	environments := ""
	for _, environment := range c.Environments {
		environments += environment.Name + ", "
	}
	// Trim last ", "
	return environments[:len(environments)-2]
}
