// Command workflow-plugin-deployment is a workflow engine external plugin that
// provides deployment pipeline step types: rolling, blue-green, canary, verify,
// rollback, and container build.
package main

import (
	"github.com/GoCodeAlone/workflow-plugin-deployment/internal"
	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

func main() {
	sdk.Serve(internal.NewDeploymentPlugin())
}
