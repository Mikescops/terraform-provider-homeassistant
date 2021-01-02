---
page_title: "ha_light Data Source - terraform-provider-homeassistant"
subcategory: ""
description: |-
  
---

# Data Source `ha_light`





## Schema

### Required

- **entity_id** (String) ID of the resource in Home Assistant

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **attributes** (List of Object) Light attributes (see [below for nested schema](#nestedatt--attributes))
- **state** (String) State of the light

<a id="nestedatt--attributes"></a>
### Nested Schema for `attributes`

Read-only:

- **brightness** (Number)
- **rgb_color** (List of Number)


