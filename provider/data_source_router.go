package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

type dataSourceRouter map[string]tfprotov5.DataSourceServer

//
// tfprotov5.DataSourceServer interface methods.
//

// ValidateDataSourceConfig is called when Terraform is checking that a data source's configuration is valid.
// It is guaranteed to have types conforming to your schema, but it is not guaranteed that all values will be known.
// This is your opportunity to do custom or advanced validation prior to a plan being generated.
func (r dataSourceRouter) ValidateDataSourceConfig(ctx context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	log.Println("[DEBUG] Enter DataSourceRouter::ValidateDataSourceConfig")
	defer log.Println("[DEBUG] Exit DataSourceRouter::ValidateDataSourceConfig")

	ds, ok := r[req.TypeName]
	if !ok {
		return nil, fmt.Errorf("Unsupported data source: %s", req.TypeName)
	}

	return ds.ValidateDataSourceConfig(ctx, req)
}

// ReadDataSource is called when Terraform is refreshing a data source's state.
func (r dataSourceRouter) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	log.Println("[DEBUG] Enter DataSourceRouter::ReadDataSource")
	defer log.Println("[DEBUG] Exit DataSourceRouter::ReadDataSource")

	ds, ok := r[req.TypeName]
	if !ok {
		return nil, fmt.Errorf("Unsupported data source: %s", req.TypeName)
	}

	return ds.ReadDataSource(ctx, req)
}
