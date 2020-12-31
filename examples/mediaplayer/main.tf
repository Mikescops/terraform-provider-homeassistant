terraform {
  required_providers {
    ha = {
      version = "~> 0.2"
      source = "hashicorp.com/edu/ha"
    }
  }
}

locals {
  default_mediaplayer_id = "media_player.android_tv"
}

data "ha_mediaplayer" "tv_cast" {
  entity_id = local.default_mediaplayer_id
}

output "cast_info" {
  value = data.ha_mediaplayer.tv_cast
}

resource "ha_mediaplayer" "tv" {
  entity_id = local.default_mediaplayer_id

#   media_content_id = "https://soundcloud.com/bangtan/christmas-love-by-jimin-of-bts"
#   media_content_type = "music"

  media_content_id = "https://www.youtube.com/watch?v=aAkMkVFwAoo"
  media_content_type = "video/youtube"
}

output "tv_player_info" {
  value = ha_mediaplayer.tv
}