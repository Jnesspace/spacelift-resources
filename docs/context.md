METADATA:
  resource_type: spacelift_context
  provider: spacelift
  service: configuration
  description: Reusable configuration collection for Spacelift resources
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_context" "RESOURCE_NAME" {
  name        = CONTEXT_NAME
  description = DESCRIPTION    # Optional
  labels      = [LABELS]      # Optional
  space_id    = SPACE_ID      # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Context identifier
      validation: Must be unique in account

  optional:
    description:
      type: String
      description: Human-readable context description
      
    labels:
      type: Set[String]
      description: Classification and automation tags
      note: Use autoattach:<label> format for auto-attachment
      default: []
      
    space_id:
      type: String
      description: Target space identifier
      default: root or legacy space

    # Hook Scripts
    before_init:
      type: List[String]
      description: Pre-initialization commands
      
    after_init:
      type: List[String]
      description: Post-initialization commands
      
    before_plan:
      type: List[String]
      description: Pre-plan commands
      
    after_plan:
      type: List[String]
      description: Post-plan commands
      
    before_apply:
      type: List[String]
      description: Pre-apply commands
      
    after_apply:
      type: List[String]
      description: Post-apply commands
      
    before_perform:
      type: List[String]
      description: Pre-perform commands
      
    after_perform:
      type: List[String]
      description: Post-perform commands
      
    before_destroy:
      type: List[String]
      description: Pre-destroy commands
      
    after_destroy:
      type: List[String]
      description: Post-destroy commands
      
    after_run:
      type: List[String]
      description: Post-run commands

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  sharing:
    - Can be attached to multiple stacks and modules
    - Requires explicit attachment via spacelift_context_attachment
    - Supports inheritance in space hierarchy
    
  automation:
    - Supports autoattach via labels
    - Provides extensive hook system
    - Hooks run in defined order
    
  configuration:
    - Can contain environment variables
    - Can contain mounted files
    - Configuration inherited by attached resources

IMPORT_FORMAT: $CONTEXT_ID