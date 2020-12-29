terraform {
  required_providers {
    ha = {
      version = "~> 0.2"
      source = "hashicorp.com/edu/ha"
    }
  }
}

provider "ha" {
  host_url = "https://home.pixelswap.fr:8888/api"
  # bearer_token = "xxxx"
}

module "magic_ha_lights" {
  source = "./light"
}

output "magic_ha_lights_initial" {
  value = module.magic_ha_lights.light_info
}

// output "magic_ha_lights_next" {
//   value = module.magic_ha_lights.tv_light_info
// }
