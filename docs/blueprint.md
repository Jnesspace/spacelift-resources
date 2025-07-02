METADATA:
  resource_type: spacelift_blueprint
  provider: spacelift
  service: templates
  description: Stack template engine for non-Terraform use cases
  version: latest
  note: Prefer spacelift_stack for Terraform users

USAGE_TEMPLATE:
```hcl
resource "spacelift_blueprint" "RESOURCE_NAME" {
  name        = BLUEPRINT_NAME
  space       = SPACE_ID
  state       = STATE          # DRAFT or PUBLISHED
  description = DESCRIPTION    # Optional
  template    = TEMPLATE_BODY  # Required if PUBLISHED
  labels      = [LABELS]      # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Blueprint identifier
      validation: Must be unique in space
      
    space:
      type: String
      description: Target space identifier
      validation: Must exist in Spacelift
      
    state:
      type: String
      description: Blueprint publication state
      allowed_values:
        - DRAFT
        - PUBLISHED

  optional:
    description:
      type: String
      description: Human-readable blueprint description
      
    template:
      type: String
      description: Blueprint template content
      validation: Required when state is PUBLISHED
      
    labels:
      type: Set[String]
      description: Classification tags
      default: []

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  templating:
    - Provides stack creation templates
    - Supports custom configuration
    - Uses templating engine
    
  states:
    DRAFT:
      - Work in progress
      - Template optional
      - Not available for use
    PUBLISHED:
      - Ready for use
      - Template required
      - Available for stack creation
    
  organization:
    - Space-scoped templates
    - Label-based classification
    - Description support
    
  usage:
    primary:
      - Non-Terraform stack creation
      - Template-based deployments
      - Standardized configurations
    alternative:
      - Use spacelift_stack for Terraform

