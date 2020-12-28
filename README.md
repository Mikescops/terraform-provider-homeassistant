# Terraform Provider Home Assistant

Run the following command to build the provider

```shell
go build -o terraform-provider-ha
```

## Install the addon

First, build and install the provider.

```shell
make install
```

Go to the examples directory :

```shell
terraform init && terraform apply
```

## Requirements

- Terraform 0.14.3
- go 1.15.6

You can setup the provider in TF or use Env variables like:

```
export HA_BEARER_TOKEN=xxxxxx
export HA_HOST_URL=https://xxxxxx
```
