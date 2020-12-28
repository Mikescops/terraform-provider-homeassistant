terraform {
  required_providers {
    ha = {
      version = "~> 0.2"
      source = "hashicorp.com/edu/ha"
    }
  }
}

provider "ha" {}

module "magic_ha_lights" {
  source = "./light"
}

output "magic_ha_lights" {
  value = module.magic_ha_lights.light_info
}
