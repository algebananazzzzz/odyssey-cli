package operations

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/algebananazzzzz/odyssey-cli/constants"
	"gopkg.in/yaml.v3"
)

func WriteConfigFile(config constants.GlobalConfig, filePath string) func() error {
	return func() error {
		// Ensure the directory exists
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}

		content, err := yaml.Marshal(config)
		if err != nil {
			return fmt.Errorf("error marshalling yaml: %w", err)
		}

		if err := os.WriteFile(filePath, content, 0666); err != nil {
			return fmt.Errorf("error writing file: %w", err)
		}

		return nil
	}
}
