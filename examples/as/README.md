# Example AssemblyScript Module

This directory contains an example [AssemblyScript](https://www.assemblyscript.org/) module for use with the Terraform waPC provider.

Use the [waPC Guest Library for AssemblyScript](https://github.com/wapc/wapc-guest-as).

## Setup

Install a recent version of [Node.js](https://nodejs.org/).

### Build

```
npm run asbuild
```

A successful build creates a waPC-compliant Wasm module `build/optimized.wasm`.
This module can be specified as the `filename` attribute of the Terraform waPC provider.
