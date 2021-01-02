terraform {
  required_providers {
    ha = {
      version = "~> 0.2"
      source  = "pixelswap.fr/iot/ha"
    }

  }
}
provider "vault" {
  address         = "https://192.168.1.77:8200"
  skip_tls_verify = true
}

data "vault_generic_secret" "HA_API_token" {
  path = "kv/ha_api_key"
}

data "vault_generic_secret" "HA_URL" {
  path = "kv/ha_url"
}

provider "ha" {
  host_url     = data.vault_generic_secret.HA_URL.data["url"]
  bearer_token = data.vault_generic_secret.HA_API_token.data["key"]
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

output "magic_ha_lights_next" {
  value = module.magic_ha_lights.tv_light_info
}

module "holiday_mood" {
  source               = "./holiday_mood"
  enable_lamps         = false
  enable_google_home   = true
  enable_tv            = true
  desired_holiday_mood = "summer"
  volume_level         = 0.5

}

