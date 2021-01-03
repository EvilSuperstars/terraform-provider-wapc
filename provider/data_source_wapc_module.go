package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

type dataSourceWapcModule struct{}

//
// tfprotov5.DataSourceServer interface methods.
//

// ValidateDataSourceConfig is called when Terraform is checking that a data source's configuration is valid.
// It is guaranteed to have types conforming to your schema, but it is not guaranteed that all values will be known.
// This is your opportunity to do custom or advanced validation prior to a plan being generated.
func (d dataSourceWapcModule) ValidateDataSourceConfig(ctx context.Context, req *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	log.Println("[DEBUG] Enter DataSourceWapcModule::ValidateDataSourceConfig")
	defer log.Println("[DEBUG] Exit DataSourceWapcModule::ValidateDataSourceConfig")

	return &tfprotov5.ValidateDataSourceConfigResponse{}, nil
}

// ReadDataSource is called when Terraform is refreshing a data source's state.
func (d dataSourceWapcModule) ReadDataSource(ctx context.Context, req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	log.Println("[DEBUG] Enter DataSourceWapcModule::ReadDataSource")
	defer log.Println("[DEBUG] Exit DataSourceWapcModule::ReadDataSource")

	_, err := req.Config.Unmarshal(tftypes.Object{})
	if err != nil {
		return &tfprotov5.ReadDataSourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error unmarshaling config",
					Detail:   fmt.Sprintf("Error unmarshaling config: %s", err.Error()),
				},
			},
		}, nil
	}

	return nil, nil
}
