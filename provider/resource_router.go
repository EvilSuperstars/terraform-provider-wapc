package provider

import (
	"context"
	"errors"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

type resourceRouter map[string]tfprotov5.ResourceServer

//
// tfprotov5.ResourceServer interface methods.
//

// ValidateResourceTypeConfig is called when Terraform is checking that a resource's configuration is valid.
// It is guaranteed to have types conforming to your schema.
// This is your opportunity to do custom or advanced validation prior to a plan being generated.
func (r resourceRouter) ValidateResourceTypeConfig(ctx context.Context, req *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::ValidateResourceTypeConfig")

	log.Println("[DEBUG] Exit ProviderServer::ValidateResourceTypeConfig")
	return nil, errors.New("ProviderServer::ValidateResourceTypeConfig not implemented")
}

// UpgradeResourceState is called when Terraform has encountered a esource with a state in a schema that doesn't match the schema's current version.
// It is the provider's responsibility to modify the state to upgrade it to the latest state schema.
func (r resourceRouter) UpgradeResourceState(ctx context.Context, req *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::UpgradeResourceState")

	log.Println("[DEBUG] Exit ProviderServer::UpgradeResourceState")
	return nil, errors.New("ProviderServer::UpgradeResourceState not implemented")
}

// ReadResource is called when Terraform is refreshing a resource's state.
func (r resourceRouter) ReadResource(ctx context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::ReadResource")

	log.Println("[DEBUG] Exit ProviderServer::ReadResource")
	return nil, errors.New("ProviderServer::ReadResource not implemented")
}

// PlanResourceChange is called when Terraform is attempting to calculate a plan for a resource.
// Terraform will suggest a proposed new state, which the provider can modify or return unmodified to influence Terraform's plan.
func (r resourceRouter) PlanResourceChange(ctx context.Context, req *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::PlanResourceChange")

	log.Println("[DEBUG] Exit ProviderServer::PlanResourceChange")
	return nil, errors.New("ProviderServer::PlanResourceChange not implemented")
}

// ApplyResourceChange is called when Terraform has detected a diff between the resource's state and the user's config, and the user has approved a planned change.
// The provider is to apply the changes contained in the plan, and return the resulting state.
func (r resourceRouter) ApplyResourceChange(ctx context.Context, req *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::ApplyResourceChange")

	log.Println("[DEBUG] Exit ProviderServer::ApplyResourceChange")
	return nil, errors.New("ProviderServer::ApplyResourceChange not implemented")
}

// ImportResourceState is called when a user has requested Terraform import a resource.
// The provider should fetch the information  specified by the passed ID and return it as one or more resource states for Terraform to assume control of.
func (r resourceRouter) ImportResourceState(ctx context.Context, req *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::ImportResourceState")

	log.Println("[DEBUG] Exit ProviderServer::ImportResourceState")
	return nil, errors.New("ProviderServer::ImportResourceState not implemented")
}
