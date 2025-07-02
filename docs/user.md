METADATA:
  resource_type: spacelift_user
  provider: spacelift
  service: access_control
  description: User access management with Identity Provider integration
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_user" "RESOURCE_NAME" {
  username         = USERNAME
  invitation_email = EMAIL_ADDRESS   # Required for new users
  
  policy {
    space_id = SPACE_ID
    role     = ROLE_TYPE    # READ, WRITE, or ADMIN
  }
}
```

ATTRIBUTES:
  required:
    username:
      type: String
      description: User identifier
      validation: Must match IdP username
      
    policy:
      type: Block Set
      min: 1
      description: Space access policies
      fields:
        space_id:
          type: String
          description: Target space identifier
          required: true
          validation: Must exist in Spacelift
        role:
          type: String
          description: Access level in space
          required: true
          allowed_values:
            - READ
            - WRITE
            - ADMIN

  optional:
    invitation_email:
      type: String
      description: User invitation email
      validation: Must be valid email
      note: Required for new users, optional for existing

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  access_control:
    - Managed through IdP
    - Role-based permissions
    - Space-level granularity
    
  roles:
    READ:
      - View resources
      - Read configurations
      - Access logs
    WRITE:
      - Create resources
      - Modify configurations
      - Trigger runs
    ADMIN:
      - Manage users
      - Configure policies
      - Full control
    
  invitation:
    - Email-based onboarding
    - Required for new users
    - Optional for existing
    
  validation:
    - Username must exist in IdP
    - Valid email required
    - At least one policy needed
    - Space must exist
    - Role must be valid

INTEGRATION:
  identity_provider:
    - Manages user identities
    - Handles authentication
    - Syncs user information
    
  policies:
    - Define access rights
    - Control space access
    - Set permission levels