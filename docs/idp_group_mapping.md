
spacelift_idp_group_mapping (Resource)

spacelift_idp_group_mapping represents a mapping (binding) between a user group (as provided by IdP) and a Spacelift User Management Policy. If you assign permissions (a Policy) to a user group, all users in the group will have those permissions unless the user's permissions are higher than the group's permissions.
Example Usage

resource "spacelift_idp_group_mapping" "test" {
  name = "test"
  policy {
    space_id = "root"
    role     = "ADMIN"
  }
  description = "test description"
}

Schema
Required

    name (String) Name of the user group - should be unique in one account
    policy (Block Set, Min: 1) (see below for nested schema)

Optional

    description (String) Description of the user group

Read-Only

    id (String) The ID of this resource.

Nested Schema for policy

Required:

    role (String) Type of access to the space. Possible values are: READ, WRITE, ADMIN
    space_id (String) ID (slug) of the space the user group has access to

On this page

    Example Usage
    Schema

Report an issue 