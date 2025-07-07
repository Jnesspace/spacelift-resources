# Resource: spacelift_user

## Description
Manages user access and permissions in Spacelift by mapping Identity Provider users to space-level roles and permissions.

## Example Usage
```hcl
# Basic user with single space access
resource "spacelift_user" "developer" {
  username         = "john.doe"
  invitation_email = "john.doe@company.com"
  
  policy {
    space_id = spacelift_space.development.id
    role     = "WRITE"
  }
}

# User with multiple space access
resource "spacelift_user" "admin" {
  username         = "admin.user"
  invitation_email = "admin@company.com"
  
  policy {
    space_id = "root"
    role     = "ADMIN"
  }
  
  policy {
    space_id = spacelift_space.production.id
    role     = "ADMIN"
  }
}
```

## Argument Reference

### Required Arguments
* `username` - (Required) Username from the Identity Provider
* `policy` - (Required) One or more space access policies
  * `space_id` - (Required) ID of the space to grant access to
  * `role` - (Required) Access level. Valid values: `READ`, `WRITE`, `ADMIN`

### Optional Arguments
* `invitation_email` - (Optional) Email address for user invitation. Required for new users

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Users must exist in the configured Identity Provider
* Multiple policies can be defined for access to different spaces
* Role hierarchy: READ < WRITE < ADMIN
* Invitation email is required when creating new users