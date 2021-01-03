package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

type server struct {
	providerSchema     *tfprotov5.Schema
	providerMetaSchema *tfprotov5.Schema
	resourceSchemas    map[string]*tfprotov5.Schema
	dataSourceSchemas  map[string]*tfprotov5.Schema

	resourceRouter
	dataSourceRouter
}

//
// tfprotov5.ProviderServer interface methods.
//

// GetProviderSchema is called when Terraform needs to know what the provider's schema is,
// along with the schemas of all its resources and data sources.
func (s *server) GetProviderSchema(ctx context.Context, req *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::GetProviderSchema")
	defer log.Println("[DEBUG] Exit ProviderServer::GetProviderSchema")

	return &tfprotov5.GetProviderSchemaResponse{
		Provider:          s.providerSchema,
		ProviderMeta:      s.providerMetaSchema,
		ResourceSchemas:   s.resourceSchemas,
		DataSourceSchemas: s.dataSourceSchemas,
	}, nil
}

// PrepareProviderConfig is called to give a provider a chance to modify the configuration the user specified before validation.
func (s *server) PrepareProviderConfig(ctx context.Context, req *tfprotov5.PrepareProviderConfigRequest) (*tfprotov5.PrepareProviderConfigResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::PrepareProviderConfig")
	defer log.Println("[DEBUG] Exit ProviderServer::PrepareProviderConfig")

	return &tfprotov5.PrepareProviderConfigResponse{
		PreparedConfig: req.Config,
	}, nil
}

// ConfigureProvider is called to pass the user-specified provider configuration to the provider.
func (s *server) ConfigureProvider(ctx context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::ConfigureProvider")
	defer log.Println("[DEBUG] Exit ProviderServer::ConfigureProvider")

	return &tfprotov5.ConfigureProviderResponse{}, nil
}

// StopProvider is called when Terraform would like providers to shut down as quickly as possible, and usually represents an interrupt.
func (s *server) StopProvider(ctx context.Context, req *tfprotov5.StopProviderRequest) (*tfprotov5.StopProviderResponse, error) {
	log.Println("[DEBUG] Enter ProviderServer::StopProvider")
	defer log.Println("[DEBUG] Exit ProviderServer::StopProvider")

	return &tfprotov5.StopProviderResponse{}, nil
}

// Provider entrypoint.
func Server() tfprotov5.ProviderServer {
	return &server{
		providerSchema: &tfprotov5.Schema{
			Version: 1,
			Block: &tfprotov5.SchemaBlock{
				Version: 1,
			},
		},
		dataSourceSchemas: map[string]*tfprotov5.Schema{
			"wapc_module": dataSourceWapcModuleSchema(),
		},
		dataSourceRouter: dataSourceRouter{
			"wapc_module": dataSourceWapcModule{},
		},
	}
}
