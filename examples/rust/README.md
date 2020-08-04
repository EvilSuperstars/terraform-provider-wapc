# Example Rust Module

This directory contains an example [Rust](https://www.rust-lang.org/) module for use with the Terraform waPC provider.

Use the [waPC Guest SDK](https://github.com/wapc/wapc-guest-rust).

## Setup

Install a recent version of Rust.

### Build

```
make release
```

A successful build creates a waPC-compliant Wasm module `target/wasm32-unknown-unknown/release/wapc_rust_example.wasm`.
This module can be specified as the `filename` attribute of the Terraform waPC provider.
