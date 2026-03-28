// Package internal implements the workflow-plugin-deployment external plugin,
// providing deployment pipeline step types.
package internal

import (
	"context"
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// deploymentPlugin implements sdk.PluginProvider.
type deploymentPlugin struct{}

// NewDeploymentPlugin returns a new deploymentPlugin instance.
func NewDeploymentPlugin() sdk.PluginProvider {
	return &deploymentPlugin{}
}

// Manifest returns plugin metadata.
func (p *deploymentPlugin) Manifest() sdk.PluginManifest {
	return sdk.PluginManifest{
		Name:        "workflow-plugin-deployment",
		Version:     "0.1.0",
		Author:      "GoCodeAlone",
		Description: "Deployment pipeline steps: rolling, blue-green, canary, verify, rollback, container_build, deploy",
	}
}


// StepTypes returns the step type names this plugin provides.
func (p *deploymentPlugin) StepTypes() []string {
	return []string{
		"step.deploy",
		"step.deploy_rolling",
		"step.deploy_blue_green",
		"step.deploy_canary",
		"step.deploy_verify",
		"step.deploy_rollback",
		"step.container_build",
	}
}

// CreateStep creates a step instance of the given type.
func (p *deploymentPlugin) CreateStep(typeName, name string, config map[string]any) (sdk.StepInstance, error) {
	switch typeName {
	case "step.deploy",
		"step.deploy_rolling",
		"step.deploy_blue_green",
		"step.deploy_canary",
		"step.deploy_verify",
		"step.deploy_rollback",
		"step.container_build":
		return &deployStep{name: name, stepType: typeName, config: config}, nil
	default:
		return nil, fmt.Errorf("deployment plugin: unknown step type %q", typeName)
	}
}

// deployStep is a stub StepInstance for deployment step types.
// TODO: Implement actual deployment strategies.
type deployStep struct {
	name     string
	stepType string
	config   map[string]any
}

// Execute runs the deployment step.
// TODO: Implement rolling, blue-green, canary strategies with health checking and rollback.
func (s *deployStep) Execute(
	_ context.Context,
	_ map[string]any,
	_ map[string]map[string]any,
	_ map[string]any,
	_ map[string]any,
	_ map[string]any,
) (*sdk.StepResult, error) {
	service, _ := s.config["service"].(string)
	image, _ := s.config["image"].(string)
	strategy, _ := s.config["strategy"].(string)
	if strategy == "" {
		strategy = s.stepType
	}

	return &sdk.StepResult{
		Output: map[string]any{
			"status":   "deployed",
			"service":  service,
			"image":    image,
			"strategy": strategy,
			"message":  fmt.Sprintf("TODO: %s not yet implemented in external plugin", s.stepType),
		},
	}, nil
}
