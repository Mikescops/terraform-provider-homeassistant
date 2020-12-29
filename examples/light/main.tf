terraform {
  required_providers {
    ha = {
      version = "~> 0.2"
      source = "hashicorp.com/edu/ha"
    }
  }
}

locals {
  default_lamp_id = "light.lampe_tv"
}

data "ha_light" "tv_light" {
  entity_id = local.default_lamp_id
}

output "light_info" {
  value = data.ha_light.tv_light
}

resource "ha_light" "tv" {
  entity_id = local.default_lamp_id
  state = "on"
}

output "tv_light_info" {
  value = ha_light.tv
}