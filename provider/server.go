package provider

import (
	"context"
	"log"

	"github.com/EvilSuperstars/terraform-provider-wapc/tfplugin5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpcProviderServer implements the Protobuf ProviderServer interface.
type grpcProviderServer struct {
	stopCh chan struct{}
}

func NewGRPCProviderServer() tfplugin5.ProviderServer {
	return &grpcProviderServer{
		stopCh: make(chan struct{}),
	}
}

//
// ProviderServer interface methods.
//

//////// Information about what a provider supports/expects

func (s *grpcProviderServer) GetSchema(ctx context.Context, req *tfplugin5.GetProviderSchema_Request) (*tfplugin5.GetProviderSchema_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::GetSchema")

	configSchema, err := GetConfigSchema()
	if err != nil {
		return nil, err
	}

	dsSchemas, err := GetDataSourceSchemas()
	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.GetProviderSchema_Response{
		DataSourceSchemas: dsSchemas,
		Provider: configSchema,
	}

	log.Println("[DEBUG] Exit ProviderServer::GetSchema")
	return resp, nil
}

func (s *grpcProviderServer) PrepareProviderConfig(ctx context.Context, req *tfplugin5.PrepareProviderConfig_Request) (*tfplugin5.PrepareProviderConfig_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::PrepareProviderConfig")

	log.Println("[DEBUG] Exit ProviderServer::PrepareProviderConfig")
	return &tfplugin5.PrepareProviderConfig_Response{}, nil
}

func (s *grpcProviderServer) ValidateResourceTypeConfig(ctx context.Context, req *tfplugin5.ValidateResourceTypeConfig_Request) (*tfplugin5.ValidateResourceTypeConfig_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::ValidateResourceTypeConfig")

	log.Println("[DEBUG] Exit ProviderServer::ValidateResourceTypeConfig")
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ValidateResourceTypeConfig not implemented")
}

func (s *grpcProviderServer) ValidateDataSourceConfig(ctx context.Context, req *tfplugin5.ValidateDataSourceConfig_Request) (*tfplugin5.ValidateDataSourceConfig_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::ValidateDataSourceConfig")

	log.Println("[DEBUG] Exit ProviderServer::ValidateDataSourceConfig")
	return &tfplugin5.ValidateDataSourceConfig_Response{}, nil
}

func (s *grpcProviderServer) UpgradeResourceState(ctx context.Context, req *tfplugin5.UpgradeResourceState_Request) (*tfplugin5.UpgradeResourceState_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::UpgradeResourceState")

	log.Println("[DEBUG] Exit ProviderServer::UpgradeResourceState")
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::UpgradeResourceState not implemented")
}

//////// One-time initialization, called before other functions below

func (s *grpcProviderServer) Configure(ctx context.Context, req *tfplugin5.Configure_Request) (*tfplugin5.Configure_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::Configure")

	log.Println("[DEBUG] Exit ProviderServer::Configure")
	return &tfplugin5.Configure_Response{}, nil
}

//////// Managed Resource Lifecycle

func (s *grpcProviderServer) ReadResource(ctx context.Context, req *tfplugin5.ReadResource_Request) (*tfplugin5.ReadResource_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::ReadResource")

	log.Println("[DEBUG] Exit ProviderServer::ReadResource")
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ReadResource not implemented")
}

func (s *grpcProviderServer) PlanResourceChange(ctx context.Context, req *tfplugin5.PlanResourceChange_Request) (*tfplugin5.PlanResourceChange_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::PlanResourceChange")

	log.Println("[DEBUG] Exit ProviderServer::PlanResourceChange")
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::PlanResourceChange not implemented")
}

func (s *grpcProviderServer) ApplyResourceChange(ctx context.Context, req *tfplugin5.ApplyResourceChange_Request) (*tfplugin5.ApplyResourceChange_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::ApplyResourceChange")

	log.Println("[DEBUG] Exit ProviderServer::ApplyResourceChange")
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ApplyResourceChange not implemented")
}

func (s *grpcProviderServer) ImportResourceState(ctx context.Context, req *tfplugin5.ImportResourceState_Request) (*tfplugin5.ImportResourceState_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::ImportResourceState")

	log.Println("[DEBUG] Exit ProviderServer::ImportResourceState")
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ImportResourceState not implemented")
}

func (s *grpcProviderServer) ReadDataSource(ctx context.Context, req *tfplugin5.ReadDataSource_Request) (*tfplugin5.ReadDataSource_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::ReadDataSource")

	log.Println("[DEBUG] Exit ProviderServer::ReadDataSource")
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ReadDataSource not implemented")
}

//////// Graceful Shutdown

func (s *grpcProviderServer) Stop(_ context.Context, _ *tfplugin5.Stop_Request) (*tfplugin5.Stop_Response, error) {
	log.Println("[DEBUG] Enter ProviderServer::Stop")

	close(s.stopCh)
	s.stopCh = make(chan struct{})

	log.Println("[DEBUG] Exit ProviderServer::Stop")
	return &tfplugin5.Stop_Response{}, nil
}
