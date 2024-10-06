package model

import (
	"log"
	"sort"
)

type Config struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	BaseCommand  string        `json:"base_command"`
	Services     []string      `json:"services"`
	Environments []Environment `json:"environments"`
}

type Environment struct {
	Name               string `json:"name"`
	Alias              string `json:"alias"`
	PreServiceCommand  string `json:"pre_service_command"`
	PostServiceCommand string `json:"post_service_command"`
	Confirm            bool   `json:"confirm"`
}

func (c *Config) BuildCommandHelpers() map[int]string {
	d := make(map[int]string)
	services := []string{}
	for _, command := range c.Commands {
		services = append(services, command.Services...)
	}
	sort.Strings(services)

	for i, service := range services {
		d[i+1] = service
	}

	return d
}

func (c *Config) BuildEnvironmentHelpers(command *Command) map[int]Environment {
	d := make(map[int]Environment)
	environments := []string{}
	for _, env := range command.Environments {
		str := env.Alias
		if str == "" {
			str = env.Name
		}
		environments = append(environments, str)
	}
	sort.Strings(environments)

	for i, environment := range environments {
		d[i+1] = command.GetEnvironmentByString(environment)
	}
	return d

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

func (c *Command) GetEnvironmentByString(envName string) Environment {
	for _, env := range c.Environments {
		if env.Name == envName || env.Alias == envName {
			return env
		}
	}
	log.Fatal("Environment not found")
	return Environment{}
}
