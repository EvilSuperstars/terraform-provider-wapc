package provider

import (
	"github.com/EvilSuperstars/terraform-provider-wapc/tfplugin5"
	"github.com/hashicorp/go-cty/cty"
)

// GetDataSourceSchemas returns this provider's configuration schema.
func GetConfigSchema() (*tfplugin5.Schema, error) {
	return &tfplugin5.Schema{
		Version: 1,
		Block:   &tfplugin5.Schema_Block{
			Attributes: []*tfplugin5.Schema_Attribute{},
		},
	}, nil
}

// GetDataSourceSchemas returns this provider's data source schemas.
func GetDataSourceSchemas() (map[string]*tfplugin5.Schema, error) {
	oType, err := cty.DynamicPseudoType.MarshalJSON()
	if err != nil {
		return nil, err
	}

	sType, err := cty.String.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return map[string]*tfplugin5.Schema{
		"wapc_module": {
			Version: 1,
			Block: &tfplugin5.Schema_Block{
				Attributes: []*tfplugin5.Schema_Attribute{
					{
						Name:     "filename",
						Type:     sType,
						Required: true,
					},
					{
						Name:     "operation",
						Type:     sType,
						Required: true,
					},
					{
						Name:     "input",
						Type:     oType,
						Required: true,
					},
					{
						Name:     "result",
						Type:     oType,
						Computed: true,
					},
				},
			},
		},
	}, nil
}
