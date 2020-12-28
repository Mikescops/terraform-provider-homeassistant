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

You need to have a Home Assistant long live token :

```
export HA_BEARER_TOKEN=xxxxxx
```
