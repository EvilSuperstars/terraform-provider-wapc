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

	configType := schemaType(dataSourceWapcModuleSchema())
	config, err := req.Config.Unmarshal(configType)
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

	filename, _, err := tftypes.WalkAttributePath(config, tftypes.AttributePath{
		Steps: []tftypes.AttributePathStep{tftypes.AttributeName("filename")},
	})
	if err != nil {
		return &tfprotov5.ReadDataSourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error getting filename",
					Detail:   fmt.Sprintf("Error getting filename: %s", err.Error()),
				},
			},
		}, nil
	}

	_, err = NewWapcModule(filename.(string))
	if err != nil {
		return &tfprotov5.ReadDataSourceResponse{
			Diagnostics: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "Error creating waPC module",
					Detail:   fmt.Sprintf("Error creating waPC module: %s", err.Error()),
				},
			},
		}, nil
	}

	return nil, nil
}

func dataSourceWapcModuleSchema() *tfprotov5.Schema {
	return &tfprotov5.Schema{
		Version: 1,
		Block: &tfprotov5.SchemaBlock{
			Version: 1,
			Attributes: []*tfprotov5.SchemaAttribute{
				{
					Name:            "filename",
					Type:            tftypes.String,
					Description:     "The filename of the waPC-compliant WebAssembly module.",
					DescriptionKind: tfprotov5.StringKindPlain,
					Required:        true,
				},
				{
					Name:            "operation",
					Type:            tftypes.String,
					Description:     "The name of the operation to invoke.",
					DescriptionKind: tfprotov5.StringKindPlain,
					Required:        true,
				},
				{
					Name:            "input",
					Type:            tftypes.DynamicPseudoType,
					Description:     "The operation's input.",
					DescriptionKind: tfprotov5.StringKindPlain,
					Required:        true,
				},
				{
					Name:            "result",
					Type:            tftypes.DynamicPseudoType,
					Description:     "The operation's result.",
					DescriptionKind: tfprotov5.StringKindPlain,
					Computed:        true,
				},
			},
		},
	}
}

// Returns a schema's type suitable for unmarshalling.
func schemaType(schema *tfprotov5.Schema) tftypes.Type {
	attributeTypes := map[string]tftypes.Type{}

	for _, attribute := range schema.Block.Attributes {
		attributeTypes[attribute.Name] = attribute.Type
	}

	return tftypes.Object{AttributeTypes: attributeTypes}
}
