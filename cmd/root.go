package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/jlineaweaver/send-it/lib/builder"
	"github.com/jlineaweaver/send-it/lib/model"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	config      model.Config

	rootCmd = &cobra.Command{
		Use:   "send-it",
		Short: "A tool to help you remember deploy commands.",
		Long: `send-it is a tool that allows you to deploy services
		without having to remember multiple different deploy commands`,
		Run: sendIt,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
}

func initConfig() {
	if cfgFile == "" {
		// Use config file from the flag.
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		cfgFile = fmt.Sprintf("%s/.send-it/config.json", home)
	}

	file, err := os.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("Failed to find or open config file err:%s", err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file err:%s", err)
	}
}

func sendIt(cmd *cobra.Command, args []string) {
	c := builder.Build(config, args)
	if len(c) == 0 {
		return
	}
	commandSlice := strings.Split(c, " ")

	command := exec.Command(commandSlice[0], commandSlice[1:]...)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	command.Dir = dir
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err = command.Run()
	if err != nil {
		log.Fatal(err)
	}
}
