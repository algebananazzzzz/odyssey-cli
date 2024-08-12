package constants

const ODYSSEY_PROJECT_GIT_URL = "https://github.com/algebananazzzzz/OdysseyFramework.git"
const TERRAFORM_LIBRARY_GIT_URL = "https://github.com/algebananazzzzz/terraform-modules"
const CICD_TEMPLATE_DIR = "cicd-templates"

var PROJECT_TYPES = map[string]string{
	"simple-static-site": "Simple Frontend Page",
	"simple-api":         "Simple Backend API",
}

var GITFLOW_STRATEGIES = map[int]string{
	1: "Production only, for quick POCs",
	2: "Staging + Production, for stable releases",
}

var MODIFICATION_TYPES = map[int]string{
	0: "Add deployment strategy",
	1: "Modify project name",
}
