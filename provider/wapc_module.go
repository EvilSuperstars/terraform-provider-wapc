package provider

import (
	"log"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/msgpack"
)

// InvokeWapcModule invokes a waPC module.
func InvokeWapcModule(config *cty.Value) (*cty.Value, error) {
	filename := config.GetAttr("filename").AsString()
	operation := config.GetAttr("operation").AsString()
	input, err := msgpack.Marshal(config.GetAttr("input"), cty.DynamicPseudoType)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] filename = %s, operation = %s, input = %#v", filename, operation, input)

	return config, nil
}
