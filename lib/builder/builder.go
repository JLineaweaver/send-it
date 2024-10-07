package builder

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/jlineaweaver/send-it/lib/model"
)

func Build(cfg model.Config, args []string) []string {
	// If no args we're in full interactive mode
	serviceDictonary := cfg.BuildCommandHelpers()
	var service string
	var environment model.Environment
	var envString string
	var command *model.Command

	if len(args) == 1 {
		// One arguement this is your service
		service = args[0]
	} else if len(args) == 2 {
		service = args[0]
		envString = args[1]
	} else if len(args) > 2 {
		log.Fatal("You've added too many arguments")
	}

	if service == "" {
		service = SelectService(serviceDictonary)
	}
	command = cfg.GetCommandByService(service)
	environmentDictonary := cfg.BuildEnvironmentHelpers(command)
	if envString == "" {
		environment = SelectEnvironment(command, environmentDictonary)
	} else {
		environment = command.GetEnvironmentByString(envString)
	}

	// Reset envString variable and build the command
	envCommand := []string{environment.Name}
	if environment.PreServiceCommand != "" {
		envCommand = append([]string{environment.PreServiceCommand}, envCommand...)
	}
	if environment.PostServiceCommand != "" {
		envCommand = append(envCommand, environment.PostServiceCommand)
	}

	c := append([]string{command.BaseCommand, service}, envCommand...)

	if !environment.SkipConfirm {
		var confirm string
		// Make sure they want to run this
		fmt.Printf("Here is the command to run\n")
		fmt.Printf("$ %s\n", strings.Join(c, " "))
		fmt.Println("Do you want to run this command? (y/n)")
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" || confirm == "yes" || confirm == "Yes" {
			return c
		} else {
			fmt.Println("Failed to confirm, exiting")
			return []string{}
		}
	}
	return c
}

func SelectService(serviceDictonary map[int]string) string {
	var service string
	// Build our list of keys in order
	keys := make([]int, 0, len(serviceDictonary))
	for key := range serviceDictonary {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	fmt.Println("Choose a service by typing the name or number:")
	for _, key := range keys {
		fmt.Printf("%d: %s\n", key, serviceDictonary[key])
	}
	fmt.Println()
	fmt.Scanln(&service)

	// Check if int or string
	if num, err := strconv.Atoi(service); err == nil {
		service = serviceDictonary[num]
	}
	return service
}

func SelectEnvironment(command *model.Command, environmentDictonary map[int]model.Environment) model.Environment {
	var envString string
	keys := make([]int, 0, len(environmentDictonary))
	for key := range environmentDictonary {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	fmt.Println("Choose an environment by typing the name or number:")
	// Display list of environments to send
	for _, key := range keys {
		envString = environmentDictonary[key].Alias
		if envString == "" {
			envString = environmentDictonary[key].Name
		}
		fmt.Printf("%d: %s\n", key, envString)
	}
	fmt.Println()
	fmt.Scanln(&envString)
	// Check if int or string
	if num, err := strconv.Atoi(envString); err == nil {
		return environmentDictonary[num]
	} else {
		return command.GetEnvironmentByString(envString)
	}
}
