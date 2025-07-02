
spacelift_user (Resource)

spacelift_user represents a mapping between a Spacelift user (managed using an Identity Provider) and a Policy. A Policy defines what access rights the user has to a given Space.
Schema
Required

    policy (Block Set, Min: 1) (see below for nested schema)
    username (String) Username of the user

Optional

    invitation_email (String) invitation_email will be used to send an invitation to the specified email address. This property is required when creating a new user. This property is optional when importing an existing user.

Read-Only

    id (String) The ID of this resource.

Nested Schema for policy

Required:

    role (String) Type of access to the space. Possible values are: READ, WRITE, ADMIN
    space_id (String) ID (slug) of the space the user has access to

On this page

    Schema

Report an issue 