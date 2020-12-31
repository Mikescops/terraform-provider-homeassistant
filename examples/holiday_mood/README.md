## Requirements

| Name | Version |
|------|---------|
| ha | ~> 0.2 |

## Providers

| Name | Version |
|------|---------|
| ha | ~> 0.2 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| desired\_holiday\_mood | Pick your favorite holiday between: christmas, halloween, summer and easter | `string` | `"christmas"` | no |
| enable\_google\_home | If set to true, the tv google home will cast a holiday song | `bool` | `true` | no |
| enable\_lamps | If set to true, the lamps will change to fit with the desired holiday theme | `bool` | `true` | no |
| enable\_tv | If set to true, the tv will cast a holiday themed video | `bool` | `true` | no |
| google\_home\_name | Name of the google home, enable\_google\_home must be set to true in order to control this google home | `string` | `"media_player.esclave_gentil"` | no |
| lamps\_name | List of the lamp names, enable\_lamps must be set to true in order to control these lamps | `list` | <pre>[<br>  "light.Lampe_TV",<br>  "light.Lampe_Cuisine",<br>  "light.Lampe_Salon"<br>]</pre> | no |
| tv\_name | Name of the tv, enable\_tv must be set to true in order to control this tv | `string` | `"media_player.android_tv"` | no |

## Outputs

No output.

