# ketch-pulumi-provider

This repository contains code for building a Ketch Pulumi provider which wraps an existing
Terraform provider.

### Build the provider:

    make cleanup
    make build_sdks

    # install localy pulumi package
    make install_pulumi_package


## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript):

Need to install locally @pulumi/ketch package with command:

    make install_nodejs_sdk

### Python:

See readme file in python examples


## Prepare examples

### Sync dependencies

    cd sdk && go mod download

## Run local examples with Golang
    # run example with pulumi
    cd sdk/examples/app/go
    pulumi up

