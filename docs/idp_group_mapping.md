METADATA:
  resource_type: spacelift_idp_group_mapping
  provider: spacelift
  service: access_control
  description: Maps Identity Provider groups to Spacelift permissions
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_idp_group_mapping" "RESOURCE_NAME" {
  name        = GROUP_NAME
  description = DESCRIPTION    # Optional
  
  policy {
    space_id = SPACE_ID
    role     = ROLE_TYPE     # READ, WRITE, or ADMIN
  }
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: IdP group identifier
      validation: Must be unique in account
      
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
    description:
      type: String
      description: Human-readable group description
      default: ""

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  permission_inheritance:
    - Group permissions apply to all members
    - Individual permissions take precedence
    - Higher individual roles override group roles
    
  roles:
    READ:
      - View resources
      - Access configurations
      - Read logs
    WRITE:
      - Modify resources
      - Update configurations
      - Trigger operations
    ADMIN:
      - Full control
      - User management
      - Policy configuration
    
  validation:
    - Group name must be unique
    - Space must exist
    - Role must be valid
    - At least one policy required

INTEGRATION:
  identity_provider:
    - Syncs group memberships
    - Manages user associations
    - Updates permissions
    
  policies:
    - Define group access rights
    - Control space permissions
    - Set role-based access