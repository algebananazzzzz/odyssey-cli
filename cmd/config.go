/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	constants "github.com/algebananazzzzz/odysseycli/constants"
	"github.com/algebananazzzzz/odysseycli/models"
	"github.com/algebananazzzzz/odysseycli/operations"
	"github.com/algebananazzzzz/odysseycli/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure settings used for Odyssey projects.",
	Long:  `Configure the Terraform backend state configuration and the AWS region for your new Odyssey projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("odyssey-config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		// Add the .odyssey directory in the home directory as a search path
		home, err := os.UserHomeDir()

		if err == nil {
			viper.AddConfigPath(filepath.Join(home, ".odyssey"))
		}

		if err := viper.ReadInConfig(); err != nil {
			var configFileNotFound viper.ConfigFileNotFoundError
			if !errors.As(err, &configFileNotFound) {
				fmt.Printf(`%sError: Configuration file not accessible.

It looks like the configuration file is missing or not accessible: %s%s

`, constants.Red, constants.Reset, err.Error())
				return
			} else {
				viper.SetConfigFile(filepath.Join(home, ".odyssey", "odyssey-config.yml"))
				fmt.Printf(`%sWarning: Configuration file not found. Creating a new configuration.%s
`, constants.Yellow, constants.Reset)

				// Set default values
				viper.SetDefault("region", "ap-southeast-1")
				viper.SetDefault("backend.region", "ap-southeast-1")
				viper.SetDefault("backend.workspace_key_prefix", "")
				viper.SetDefault("backend.bucket", "")
			}
		}

		var (
			region                    = viper.GetString("region")
			backendRegion             = viper.GetString("backend.region")
			backendWorkspaceKeyPrefix = viper.GetString("backend.workspace_key_prefix")
			backendBucket             = viper.GetString("backend.bucket")
			confirm                   bool
		)

		if err := huh.NewNote().
			Title("OdysseyCli - Config").
			Description(fmt.Sprintf(`This command sets up the Terraform backend state and AWS region for your new Odyssey projects.
The configuration will be saved in %s%s.odyssey/odyssey-config.yml%s in your home directory.

`, constants.Bold, constants.Blue, constants.Reset)).Next(true).Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if err := huh.NewNote().
			Title("To create an S3 bucket for storing Terraform state:").
			Description(`1. Log in to your AWS Management Console.
2. Navigate to the S3 service.
3. Click "Create bucket" and follow the prompts.

`).Next(true).Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if err := huh.NewInput().
			Title("Please enter the AWS region where you would like to deploy your infrastructure (e.g., ap-southeast-1).").
			Value(&region).
			Placeholder("ap-southeast-1").
			Validate(utils.ValidateAlphanumeric).
			Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if err := huh.NewInput().
			Title("Please provide the name of the S3 bucket where your Terraform state will be stored").
			Value(&backendBucket).
			Validate(utils.ValidateAlphanumeric).
			Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if err := huh.NewInput().
			Title("Please provide the region of the S3 bucket where your Terraform state will be stored").
			Value(&backendRegion).
			Placeholder("ap-southeast-1").
			Validate(utils.ValidateAlphanumeric).
			Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if err := huh.NewInput().
			Title("Enter the key prefix for your Terraform backend workspace used to organize your state files (e.g., tfstate)").
			Value(&backendWorkspaceKeyPrefix).
			Placeholder("tfstate").
			Validate(utils.ValidateAlphanumeric).
			Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		globalConfig := constants.GlobalConfig{
			Region: region,
			BackendConfig: constants.BackendConfig{
				Bucket:             backendBucket,
				WorkspaceKeyPrefix: backendWorkspaceKeyPrefix,
				Region:             backendRegion,
			},
		}

		if err := huh.NewConfirm().
			Title("Please confirm that the configurations are correct before proceeding:").
			Description(
				utils.FormatNewGlobalConfig(globalConfig)).
			Value(&confirm).Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if confirm {
			tasks := []models.Task{
				{
					Name: "Writing Configuration",
					Run:  operations.WriteConfigFile(globalConfig, viper.ConfigFileUsed()),
				},
			}

			p := tea.NewProgram(models.NewModel(tasks))

			if _, err := p.Run(); err != nil {
				fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			}

		} else {
			fmt.Printf("%sExiting: Confirmation not received.%s\n", constants.Yellow, constants.Reset)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
