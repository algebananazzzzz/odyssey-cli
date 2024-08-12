package operations

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/algebananazzzzz/odysseycli/constants"
)

func CloneProjectFiles(config constants.ProjectConfig) func() error {
	return func() error {
		// Create a temporary directory for cloning
		tempDir, err := os.MkdirTemp("", "git-clone-")
		if err != nil {
			return fmt.Errorf("failed to create temp directory: %v", err)
		}

		defer os.RemoveAll(tempDir) // Clean up temp directory when done

		// Clone the repository
		cmd := exec.Command("git", "clone", "--depth", "1", constants.ODYSSEY_PROJECT_GIT_URL, tempDir)

		if _, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to clone repository: %v", err)
		}

		// Ensure the local path exists
		if err := os.MkdirAll(config.Dir, 0755); err != nil {
			return fmt.Errorf("failed to create local directory: %v", err)
		}

		// Copy the project folder from the cloned repo to the local path
		sourcePath := filepath.Join(tempDir, config.Type) + string(os.PathSeparator) + "."

		cmd = exec.Command("cp", "-R", sourcePath, config.Dir)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to copy folder: %v, %s", err, output)
		}

		// Copy the CI/CD configuration file from the cloned repo to the local path
		sourcePath = filepath.Join(tempDir, constants.CICD_TEMPLATE_DIR, "gitlab", fmt.Sprintf("%s.%d.yml", config.Type, config.DeploymentStrategy))
		localPath := filepath.Join(config.Dir, ".gitlab-ci.yml")

		cmd = exec.Command("cp", sourcePath, localPath)
		if _, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to copy file: %v", err)
		}

		return nil
	}
}

func InitGit(localPath string, originURL *string) func() error {
	return func() error {
		// Initialize a new repository
		cmd := exec.Command("git", "init")
		cmd.Dir = localPath
		if _, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to initialize repository: %v", err)
		}

		// If originURL is provided, add it as a remote
		if originURL != nil {
			cmd = exec.Command("git", "remote", "add", "origin", *originURL)
			cmd.Dir = localPath
			if _, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("failed to add remote origin: %v", err)
			}
		}

		return nil
	}
}

func AddSubmodule(repoPath, submodulePath, submoduleURL string) func() error {
	return func() error { // Construct the git submodule add command
		cmd := exec.Command("git", "submodule", "add", submoduleURL, submodulePath)
		cmd.Dir = repoPath

		// Run the command
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to add submodule: %v, output: %s", err, string(output))
		}

		return nil
	}
}
