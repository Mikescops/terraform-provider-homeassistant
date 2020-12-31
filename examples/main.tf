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
//  bearer_token = "xx.xxx.xx"
}

module "magic_ha_lights" {
  source = "./light"
}

module "magic_ha_mediaplayer" {
  source = "./mediaplayer"
}

output "magic_ha_lights_initial" {
  value = module.magic_ha_lights.light_info
}

output "magic_ha_mediaplayer_initial" {
  value = module.magic_ha_mediaplayer.cast_info
}

// output "magic_ha_lights_next" {
//   value = module.magic_ha_lights.tv_light_info
// }


module "holiday_mood" {
  source = "./holiday_mood"
  enable_lamps = false
  enable_google_home = true
  enable_tv = true
  desired_holiday_mood = "christmas"
}
