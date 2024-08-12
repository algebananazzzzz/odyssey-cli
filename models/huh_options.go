package models

import (
	constants "github.com/algebananazzzzz/odysseycli/constants"
	"github.com/charmbracelet/huh"
)

func ProjectTypeHuhOptions() []huh.Option[string] {
	options := []huh.Option[string]{}
	for idx, val := range constants.PROJECT_TYPES {
		options = append(options, huh.NewOption(val, idx))
	}
	return options
}

func DeploymentStrategyHuhOptions() []huh.Option[int] {
	options := []huh.Option[int]{}
	for idx, val := range constants.GITFLOW_STRATEGIES {
		options = append(options, huh.NewOption(val, idx))
	}
	return options
}

func ModificationTypeHuhOptions(projectConfig constants.ProjectConfig) []huh.Option[int] {
	if projectConfig.DeploymentStrategy != 0 {
		return []huh.Option[int]{
			huh.NewOption("Modify deployment strategy", 1),
			huh.NewOption("Modify project name", 2),
		}
	} else {
		return []huh.Option[int]{
			huh.NewOption("Add deployment strategy", 0),
			huh.NewOption("Modify project name", 2),
		}
	}
}
