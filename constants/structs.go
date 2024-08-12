package constants

type BackendConfig struct {
	Bucket             string `yaml:"bucket"`
	WorkspaceKeyPrefix string `yaml:"workspace_key_prefix"`
	Region             string `yaml:"region"`
}

type GlobalConfig struct {
	Region        string        `yaml:"region"`
	BackendConfig BackendConfig `yaml:"backend"`
}

type ProjectConfig struct {
	Code               string
	Dir                string
	Type               string
	DeploymentStrategy int
	GlobalConfig       GlobalConfig
	Env                string
}
