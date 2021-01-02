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

- [Terraform](https://www.terraform.io/downloads.html) 0.14.3
- [Go](https://golang.org/doc/install) 1.15.6

You can setup the provider in TF or use env variables like:

```bash
export HA_BEARER_TOKEN=xxxxxx
export HA_HOST_URL=https://<insert_domain_or_ip>/api
```

## How to run tests

You need to have a Home Assistant instance running with a dummy light setup.
To do so, add the following configuration to your configuration.yaml file :

```yaml
# Dummy light
light:
  - platform: template
    lights:
      dummy_light:
        friendly_name: 'Dummy Light'
        turn_on:
        turn_off:
        set_level:
```

You also need to setup env variables in your terminal :

```bash
export HA_BEARER_TOKEN=xxxxxx
export HA_HOST_URL=https://<insert_domain_or_ip>/api
```

Then run the following command :

```bash
make testacc
```
