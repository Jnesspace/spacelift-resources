# Resource: spacelift_saved_filter

## Description
A saved filter defines custom criteria for filtering resources in Spacelift views, allowing users to create reusable filters for stacks, modules, contexts, and webhooks.

## Example Usage
```hcl
# Public filter for team stacks
resource "spacelift_saved_filter" "team_stacks" {
  name      = "team-xyz-stacks"
  type      = "stacks"
  is_public = true
  
  data = jsonencode({
    key = "activeFilters"
    value = jsonencode({
      filters = [
        [
          "name",
          {
            key         = "name"
            filterName  = "name"
            type        = "STRING"
            values      = ["team_xyz_*"]
          }
        ]
      ]
      sort = {
        direction = "ASC"
        option    = "space"
      }
      order = [
        {
          name    = "name"
          visible = true
        },
        {
          name    = "space"
          visible = true
        }
      ]
    })
  })
}

# Private webhook filter
resource "spacelift_saved_filter" "production_webhooks" {
  name      = "production-webhooks"
  type      = "webhooks"
  is_public = false
  
  data = jsonencode({
    key = "activeFilters"
    value = jsonencode({
      filters = [
        [
          "label",
          {
            key         = "label"
            filterName  = "label"
            type        = "STRING"
            values      = ["production"]
          }
        ]
      ]
    })
  })
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique filter identifier
* `type` - (Required) Target view type. Valid values: `stacks`, `blueprints`, `contexts`, `webhooks`
* `is_public` - (Required) Whether the filter is visible to all users
* `data` - (Required) JSON-encoded filter configuration

### Read-Only Arguments
* `id` - Unique resource identifier
* `created_by` - Username of the filter creator

## Import
```bash
terraform import spacelift_saved_filter.example $FILTER_ID
```

## Notes
* Filter data must be valid JSON with specific structure
* Public filters are available to all users in the account
* Different types support different filter criteria