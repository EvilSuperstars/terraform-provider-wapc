package provider

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/go-cty/cty"
	ctyjson "github.com/hashicorp/go-cty/cty/json"
	"github.com/wapc/wapc-go"
)

// InvokeWapcModule invokes a waPC module.
func InvokeWapcModule(ctx context.Context, config *cty.Value) (*cty.Value, error) {
	src := config.GetAttr("filename").AsString()
	operation := config.GetAttr("operation").AsString()
	input, err := ctyValueToJsonBytes(config.GetAttr("input"))
	if err != nil {
		return nil, fmt.Errorf("error marshaling input to JSON: %w", err)
	}

	log.Printf("[DEBUG] Reading WebAssembly module %s", src)
	code, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, fmt.Errorf("error reading WebAssembly module (%s): %w", src, err)
	}

	module, err := wapc.New(code, wapc.NoOpHostCallHandler)
	if err != nil {
		return nil, fmt.Errorf("error compiling WebAssembly module: %w", err)
	}
	module.SetLogger(func(msg string) { log.Println(msg) })
	module.SetWriter(func(msg string) { log.Print(msg) })
	defer module.Close()

	instance, err := module.Instantiate()
	if err != nil {
		return nil, fmt.Errorf("error instantiating WebAssembly module: %w", err)
	}
	defer instance.Close()

	log.Printf("[DEBUG] Invoking WebAssembly operation %s", operation)
	result, err := instance.Invoke(ctx, operation, input)
	if err != nil {
		return nil, fmt.Errorf("error invoking WebAssembly operation (%s): %w", operation, err)
	}

	resultValue, err := jsonBytesToCtyValue(result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling result from JSON: %w", err)
	}

	state := cty.ObjectVal(map[string]cty.Value{
		"filename":  config.GetAttr("filename"),
		"operation": config.GetAttr("operation"),
		"input":     config.GetAttr("input"),
		"result":    resultValue,
	})

	return &state, nil
}

// ctyValueToJsonBytes marshals a cty.Value into JSON bytes.
func ctyValueToJsonBytes(value cty.Value) ([]byte, error) {
	simple := &ctyjson.SimpleJSONValue{Value: value}

	return simple.MarshalJSON()
}

// jsonBytesToCtyValue unmarshals JSON bytes into a cty.Value.
func jsonBytesToCtyValue(buf []byte) (cty.Value, error) {
	if len(buf) == 0 {
		return cty.NullVal(cty.DynamicPseudoType), nil
	}

	simple := &ctyjson.SimpleJSONValue{}
	err := simple.UnmarshalJSON(buf)
	if err != nil {
		return cty.NilVal, err
	}

	return simple.Value, nil
}
