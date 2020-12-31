terraform {
  required_providers {
    ha = {
      version = "~> 0.2"
      source = "hashicorp.com/edu/ha"
    }
  }
}

locals {
  allowed_holiday_theme = { for v in ["christmas", "halloween", "summer", "easter"] : v => v }
  theme = local.allowed_holiday_theme[var.desired_holiday_mood] # will fail if var.desired_holiday_mood is invalid
  lamp_color = {
    "christmas" = [255, 0, 0], // Red
    "halloween" = [255, 165, 0], // Orange
    "summer" = [0, 255, 255], // Aqua blue / cyan
    "easter" = [128, 0 , 128] // Purple
  }
  tv_video = {
    "christmas" = "https://www.youtube.com/watch?v=L_LUpnjgPso", // Fireplace video
    "easter" = "https://www.youtube.com/watch?v=OEpZlj4T9U0", // Easter holiday animation for kid
    "summer" = "https://www.youtube.com/watch?v=qREKP9oijWI" // Beach video
    "halloween" = "https://www.youtube.com/watch?v=VB-tVj9EZPk" // Pumpkin video
  }

  song = {
    "christmas" = "https://www.youtube.com/watch?v=aAkMkVFwAoo", // Mariah Carey - All I want for christmas is you
    "easter" = "https://www.youtube.com/watch?v=IYV5fIOVWiU", // bell song
    "halloween" = "https://www.youtube.com/watch?v=WZt7YEuLQ1U", // Creepy horror song
    "summer" = "https://www.youtube.com/watch?v=X-77txuiVXs" // Loona - Vamos a la playa
  }
}

// Display video on tv if enabled
resource "ha_mediaplayer" "tv_video" {
  count             = var.enable_tv ? 1 : 0
  media_content_id = local.tv_video[local.theme]
  media_content_type = "youtube"
  entity_id = var.tv_name
}

// Put music on google home if enabled
resource "ha_mediaplayer" "ghome_song" {
  count             = var.enable_google_home ? 1 : 0
  media_content_id = local.song[local.theme]
  media_content_type = "youtube"
  entity_id = var.google_home_name
}

resource "ha_light" "lights" {
  count = length(var.lamps_name)
  entity_id = var.lamps_name[count.index]
  rgb_color = local.lamp_color[local.theme]
  brightness = 200
  state = "on"
}


