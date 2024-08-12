package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var cmdShip = &cobra.Command{
	Use:   "print [string to print]",
	Short: "Deploy a service.",
	Long:  `Deploy a service based on your config.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("No service provided")
		}
		serviceName := args[0]
		command := config.GetCommandByService(serviceName)
		if command == nil {
			log.Fatalf("No command found for service: %s", serviceName)
		}
		if len(args) == 1 {

		}
	},
}
