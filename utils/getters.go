package utils

import (
	"github.com/algebananazzzzz/odyssey-cli/constants"
)

func GetProjectType(projectType string) string {
	return constants.PROJECT_TYPES[projectType]
}

func GetGitflowStrategy(projectGitFlowStrategy int) string {
	return constants.GITFLOW_STRATEGIES[projectGitFlowStrategy]
}

func GetEnvironments(projectGitFlowStrategy int) []string {
	switch projectGitFlowStrategy {
	case 2:
		return []string{"prd", "stg"}
	case 3:
		return []string{"prd", "preprd", "stg"}
	default:
		return []string{"prd"}
	}
}

func GetModificationType(i int) string {
	return constants.MODIFICATION_TYPES[i]
}
