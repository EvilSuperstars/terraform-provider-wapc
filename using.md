# `waPC` Provider

The Terraform [waPC](https://github.com/EvilSuperstars/terraform-provider-wapc) provider is used to interact with waPC-compliant WebAssembly modules.
The provider uses the [wapc-go library](https://github.com/wapc/wapc-go).

This provider requires no configuration.

### Example Usage

```hcl
provider "wapc" {}

data "wapc_module" "example" {
  filename  = "hello.wasm"
  operation = "hello"
  input     = "example"
}
```

## Data Sources

### `wapc_module`

#### Argument Reference

The following arguments are supported:

* `filename` - (Required, string) The filename of the waPC-compliant WebAssembly module.
* `operation` - (Required, string) The name of the operation to invoke.
* `input` - (Required, string) The operation's input.

#### Attributes Reference

The following attributes are exported in addition to the above configuration:

* `result` - (string) The operation's result.
