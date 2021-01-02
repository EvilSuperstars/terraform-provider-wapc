package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

type dataSourceWapc struct{}

//
// tfprotov5.DataSourceServer interface methods.
//

// ValidateDataSourceConfig is called when Terraform is checking that a data source's configuration is valid.
// It is guaranteed to have types conforming to your schema, but it is not guaranteed that all values will be known.
// This is your opportunity to do custom or advanced validation prior to a plan being generated.
func (d dataSourceWapc) ValidateDataSourceConfig(ctx context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	log.Println("[DEBUG] Enter DataSourceWapc::ValidateDataSourceConfig")
	defer log.Println("[DEBUG] Exit DataSourceWapc::ValidateDataSourceConfig")

	return &tfprotov5.ValidateDataSourceConfigResponse{}, nil
}

// ReadDataSource is called when Terraform is refreshing a data source's state.
func (d dataSourceWapc) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	log.Println("[DEBUG] Enter DataSourceWapc::ReadDataSource")
	defer log.Println("[DEBUG] Exit DataSourceWapc::ReadDataSource")

	return nil, nil
}
