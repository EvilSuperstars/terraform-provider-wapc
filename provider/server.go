package provider

import (
	"context"

	"github.com/EvilSuperstars/terraform-provider-wapc/tfplugin5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// grpcProviderServer implements the Protobuf ProviderServer interface.
type grpcProviderServer struct {}

func NewGRPCProviderServer() tfplugin5.ProviderServer {
	return &grpcProviderServer{}
}

//////// Information about what a provider supports/expects

func (s *grpcProviderServer) GetSchema(ctx context.Context, req *tfplugin5.GetProviderSchema_Request) (*tfplugin5.GetProviderSchema_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::GetSchema not implemented")
}

func (s *grpcProviderServer) PrepareProviderConfig(ctx context.Context, req *tfplugin5.PrepareProviderConfig_Request) (*tfplugin5.PrepareProviderConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::PrepareProviderConfig not implemented")
}

func (s *grpcProviderServer) ValidateResourceTypeConfig(ctx context.Context, req *tfplugin5.ValidateResourceTypeConfig_Request) (*tfplugin5.ValidateResourceTypeConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ValidateResourceTypeConfig not implemented")
}

func (s *grpcProviderServer) ValidateDataSourceConfig(ctx context.Context, req *tfplugin5.ValidateDataSourceConfig_Request) (*tfplugin5.ValidateDataSourceConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ValidateDataSourceConfig not implemented")
}

func (s *grpcProviderServer) UpgradeResourceState(ctx context.Context, req *tfplugin5.UpgradeResourceState_Request) (*tfplugin5.UpgradeResourceState_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::UpgradeResourceState not implemented")
}

//////// One-time initialization, called before other functions below

func (s *grpcProviderServer) Configure(ctx context.Context, req *tfplugin5.Configure_Request) (*tfplugin5.Configure_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::Configure not implemented")
}

//////// Managed Resource Lifecycle

func (s *grpcProviderServer) ReadResource(ctx context.Context, req *tfplugin5.ReadResource_Request) (*tfplugin5.ReadResource_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ReadResource not implemented")
}

func (s *grpcProviderServer) PlanResourceChange(ctx context.Context, req *tfplugin5.PlanResourceChange_Request) (*tfplugin5.PlanResourceChange_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::PlanResourceChange not implemented")
}

func (s *grpcProviderServer) ApplyResourceChange(ctx context.Context, req *tfplugin5.ApplyResourceChange_Request) (*tfplugin5.ApplyResourceChange_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ApplyResourceChange not implemented")
}

func (s *grpcProviderServer) ImportResourceState(ctx context.Context, req *tfplugin5.ImportResourceState_Request) (*tfplugin5.ImportResourceState_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ImportResourceState not implemented")
}

func (s *grpcProviderServer) ReadDataSource(ctx context.Context, req *tfplugin5.ReadDataSource_Request) (*tfplugin5.ReadDataSource_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::ReadDataSource not implemented")
}

//////// Graceful Shutdown

func (s *grpcProviderServer) Stop(ctx context.Context, req *tfplugin5.Stop_Request) (*tfplugin5.Stop_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "ProviderServer::Stop not implemented")
}
