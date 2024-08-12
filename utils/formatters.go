package utils

import (
	"fmt"

	"github.com/algebananazzzzz/odysseycli/constants"
)

func FormatNewProjectConfig(config constants.ProjectConfig) string {
	return fmt.Sprintf(`Project Code: %s
Directory: %s
Project Type: %v
Deployment Strategy: %v
		`, config.Code, config.Dir,
		GetProjectType(config.Type),
		GetGitflowStrategy(config.DeploymentStrategy))
}

func FormatNewGlobalConfig(config constants.GlobalConfig) string {
	return fmt.Sprintf(`Region: %s

Terraform State
Bucket: %s
Region: %s
Workspace key prefix: %s
			`, config.Region, config.BackendConfig.Bucket, config.BackendConfig.Region, config.BackendConfig.WorkspaceKeyPrefix)
}
