package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	//rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")

	//rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(cmdShip)
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
