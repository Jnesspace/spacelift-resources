METADATA:
  resource_type: spacelift_space
  provider: spacelift
  service: organization
  description: Hierarchical resource collection for granular access control
  version: latest
  permissions_required: root Admin

USAGE_TEMPLATE:
```hcl
resource "spacelift_space" "RESOURCE_NAME" {
  name            = SPACE_NAME
  parent_space_id = PARENT_ID          # Optional, defaults to "root"
  description     = DESCRIPTION        # Optional
  inherit_entities = true             # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Space identifier
      validation: Must be unique in account

  optional:
    parent_space_id:
      type: String
      description: Parent space identifier
      default: "root"
      immutable: true
      
    description:
      type: String
      description: Human-readable space description
      
    inherit_entities:
      type: Boolean
      description: Inherit parent space read access
      default: false
      
    labels:
      type: Set[String]
      description: Space classification tags
      default: []

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  hierarchy:
    - Forms tree structure with root space
    - All non-root spaces must have parent
    - Supports entity inheritance from parent
    
  access_control:
    - Enables granular resource access
    - Can inherit read access from parent
    - Requires root Admin permissions
    
  resource_management:
    - Contains stacks, modules, policies
    - Supports resource organization
    - Enables team-based resource isolation

  limitations:
    - Creation requires root Admin access
    - Must be managed from root space or API
    - Parent relationship is immutable

IMPORT_FORMAT: $SPACE_ID