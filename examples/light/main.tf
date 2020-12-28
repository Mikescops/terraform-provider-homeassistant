terraform {
  required_providers {
    ha = {
      version = "~> 0.2"
      source = "hashicorp.com/edu/ha"
    }
  }
}

data "ha_light" "tv_light" {
  id = "lampe_tv"
}

output "light_info" {
  value = data.ha_light.tv_light
}
