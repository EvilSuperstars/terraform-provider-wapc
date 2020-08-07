package provider

import (
	"fmt"

	"github.com/EvilSuperstars/terraform-provider-wapc/tfplugin5"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/msgpack"
)

// GetDataSourceSchemas returns this provider's configuration schema.
func GetConfigSchema() (*tfplugin5.Schema, error) {
	return &tfplugin5.Schema{
		Version: 1,
		Block: &tfplugin5.Schema_Block{
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
						Optional: true,
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

// GetObjectTypeFromSchema returns a cty.Type that can wholly represent the schema input.
func GetObjectTypeFromSchema(schema *tfplugin5.Schema) (cty.Type, error) {
	bm := make(map[string]cty.Type)

	for _, att := range schema.Block.Attributes {
		var t cty.Type

		err := t.UnmarshalJSON(att.Type)
		if err != nil {
			return cty.NilType, fmt.Errorf("failed to unmarshall type %s: %w", string(att.Type), err)
		}

		bm[att.Name] = t
	}

	return cty.Object(bm), nil
}

// MarshalDataSource marshals a cty.Value into a msgpack-ed data source.
func MarshalDataSource(ds string, data *cty.Value) ([]byte, error) {
	s, err := GetDataSourceSchemas()
	if err != nil {
		return nil, err
	}

	t, err := GetObjectTypeFromSchema(s[ds])
	if err != nil {
		return nil, err
	}

	return msgpack.Marshal(*data, t)
}

// UnmarshalDataSource umarshals a msgpack-ed data source into its corresponding cty.Value.
func UnmarshalDataSource(ds string, data []byte) (cty.Value, error) {
	s, err := GetDataSourceSchemas()
	if err != nil {
		return cty.NilVal, err
	}

	t, err := GetObjectTypeFromSchema(s[ds])
	if err != nil {
		return cty.NilVal, err
	}

	return msgpack.Unmarshal(data, t)
}
