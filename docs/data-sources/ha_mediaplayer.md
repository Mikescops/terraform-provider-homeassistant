---
page_title: "ha_mediaplayer Data Source - terraform-provider-homeassistant"
subcategory: ""
description: |-
  
---

# Data Source `ha_mediaplayer`





## Schema

### Required

- **entity_id** (String) ID of the resource in Home Assistant

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **attributes** (List of Object) Device attributes (see [below for nested schema](#nestedatt--attributes))
- **state** (String) State of the device

<a id="nestedatt--attributes"></a>
### Nested Schema for `attributes`

Read-only:

- **media_title** (String)
- **volume_level** (Number)


