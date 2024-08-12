/*
Copyright © 2024 Daniel Zhou algebananazzzzz.devops@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
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

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Odyssey project.",
	Long:  `Create a new Odyssey project in your current directory with the pre-configured Terraform backend state and AWS region settings. Users can choose the project code, specify the type of project to create, and determine the number of environments needed.`,
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
			fmt.Printf(`%sError: Configuration file not found.

It looks like the configuration file is missing or not accessible. Please ensure that the configuration file is present in the expected directory.
%s
To set up your configuration, please run the following command:
			
	%s%sodyssey-cli config

`, constants.Red, constants.Reset, constants.Bold, constants.Purple)
			return
		}

		var (
			projectCode               string
			projectDir                string
			projectType               string
			projectDeploymentStrategy int
			confirm                   bool
		)

		if err := huh.NewNote().
			Title("OdysseyCli - New").
			Description("This command creates a new Odyssey project in your current directory with the pre-configured Terraform backend state and AWS region settings.\n\n").Next(true).Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if err := huh.NewInput().
			Title("Think of a Project Code for your new project!").
			Value(&projectCode).
			Placeholder("My-Awesome-Project").
			Validate(utils.ValidateAlphanumeric).
			Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		projectDir = utils.ConvertNameToFilePath(projectCode)

		if err := huh.NewInput().
			Title("Which directory would you like to store your project files?").
			Value(&projectDir).
			Validate(utils.ValidateFilePath).
			Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if err := huh.NewSelect[string]().
			Title("What type of project would you like to create?").
			Options(
				models.ProjectTypeHuhOptions()...,
			).
			Value(&projectType).
			Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if err := huh.NewSelect[int]().
			Title("Choose how many environments to deploy to. This will impact your CI/CD workflow.").
			Options(
				models.DeploymentStrategyHuhOptions()...,
			).
			Value(&projectDeploymentStrategy).Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		projectConfig := constants.ProjectConfig{
			Code:               projectCode,
			Dir:                projectDir,
			Type:               projectType,
			DeploymentStrategy: projectDeploymentStrategy,
			GlobalConfig: constants.GlobalConfig{
				Region: viper.GetString("region"),
				BackendConfig: constants.BackendConfig{
					Bucket:             viper.GetString("backend.bucket"),
					Region:             viper.GetString("backend.region"),
					WorkspaceKeyPrefix: viper.GetString("backend.workspace_key_prefix"),
				},
			},
		}

		if err := huh.NewConfirm().
			Title("Shall we create the project? Here are the selected configuration:").
			Description(
				utils.FormatNewProjectConfig(projectConfig)).
			Value(&confirm).Run(); err != nil {
			fmt.Printf("%sError: %s%s\n", constants.Red, err, constants.Reset)
			os.Exit(1)
		}

		if confirm {
			tasks := []models.Task{
				{
					Name: "Sourcing Project Files",
					Run:  operations.CloneProjectFiles(projectConfig),
				},
				{
					Name: "Initialize Git",
					Run:  operations.InitGit(projectDir, nil),
				},
				{
					Name: "Setup Terraform Submodules",
					Run:  operations.AddSubmodule(projectDir, "infra/modules", constants.TERRAFORM_LIBRARY_GIT_URL),
				},
				{
					Name: "Customizing Project Files",
					Run:  operations.CustomizeContentFiles(projectConfig),
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
	rootCmd.AddCommand(newCmd)
}
