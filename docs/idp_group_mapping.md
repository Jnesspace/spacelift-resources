# Resource: spacelift_idp_group_mapping

## Description
Maps Identity Provider user groups to Spacelift permissions, allowing group-based access control where all users in a group inherit the specified permissions.

## Example Usage
```hcl
# Admin group mapping
resource "spacelift_idp_group_mapping" "admins" {
  name        = "spacelift-admins"
  description = "Administrator access for Spacelift admins group"
  
  policy {
    space_id = "root"
    role     = "ADMIN"
  }
}

# Development team mapping
resource "spacelift_idp_group_mapping" "developers" {
  name        = "development-team"
  description = "Development team access"
  
  policy {
    space_id = spacelift_space.development.id
    role     = "WRITE"
  }
  
  policy {
    space_id = spacelift_space.staging.id
    role     = "READ"
  }
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique group name within the account (should match IdP group name)
* `policy` - (Required) One or more space access policies
  * `space_id` - (Required) ID of the space to grant access to
  * `role` - (Required) Access level. Valid values: `READ`, `WRITE`, `ADMIN`

### Optional Arguments
* `description` - (Optional) Human-readable description of the group mapping

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Group name should match the group name from your Identity Provider
* Multiple policies can be defined for access to different spaces
* User permissions take precedence over group permissions if higher
* All group members inherit the defined permissions