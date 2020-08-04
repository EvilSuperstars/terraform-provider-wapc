package provider

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/msgpack"
	"github.com/wapc/wapc-go"
)

// InvokeWapcModule invokes a waPC module.
func InvokeWapcModule(ctx context.Context, config *cty.Value) (*cty.Value, error) {
	src := config.GetAttr("filename").AsString()
	operation := config.GetAttr("operation").AsString()
	input, err := msgpack.Marshal(config.GetAttr("input"), cty.DynamicPseudoType)
	if err != nil {
		return nil, err
	}

	code, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, fmt.Errorf("error reading WebAssembly code (%s): %w", src, err)
	}

	module, err := wapc.New(func(msg string) { log.Printf("%s", msg) }, code, wapc.NoOpHostCallHandler)
	if err != nil {
		return nil, fmt.Errorf("error compiling WebAssembly module: %w", err)
	}
	defer module.Close()

	instance, err := module.Instantiate()
	if err != nil {
		return nil, fmt.Errorf("error instantiating WebAssembly module: %w", err)
	}
	defer instance.Close()

	_, err = instance.Invoke(ctx, operation, input)
	if err != nil {
		return nil, fmt.Errorf("error invoking WebAssembly operation (%s): %w", operation, err)
	}

	return config, nil
}