package builder

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/jlineaweaver/send-it/lib/model"
)

func Build(cfg model.Config, args []string) {
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

	fmt.Printf("%s %s %s", command.BaseCommand, service, environment.Name)
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
