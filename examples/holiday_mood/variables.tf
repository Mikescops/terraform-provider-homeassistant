variable "enable_lamps" {
  type        = bool
  description = "If set to true, the lamps will change to fit with the desired holiday theme"
  default     = true
}

variable "lamps_name" {
  type        = list
  description = "List of the lamp names, enable_lamps must be set to true in order to control these lamps"
  default     = ["light.Lampe_TV", "light.Lampe_Cuisine", "light.Lampe_Salon"]
}

variable "enable_tv" {
  type        = bool
  description = "If set to true, the tv will cast a holiday themed video"
  default     = true
}

variable "tv_name" {
  type        = string
  description = "Name of the tv, enable_tv must be set to true in order to control this tv"
  default     = "media_player.android_tv"
}

variable "enable_google_home" {
  type        = bool
  description = "If set to true, the tv google home will cast a holiday song"
  default     = true
}

variable "google_home_name" {
  type        = string
  description = "Name of the google home, enable_google_home must be set to true in order to control this google home"
  default     = "media_player.esclave_gentil"
}

variable desired_holiday_mood {
  type = string
  description = "Pick your favorite holiday between: christmas, halloween, summer and easter"
  default = "christmas"
}
