spacelift_context_attachment (Resource)

spacelift_context_attachment represents a Spacelift attachment of a single context to a single stack or module, with a predefined priority.
Example Usage

# For a stack
resource "spacelift_context_attachment" "attachment" {
  context_id = "prod-k8s-ie"
  stack_id   = "k8s-core"
  priority   = 0
}

# For a module
resource "spacelift_context_attachment" "attachment" {
  context_id = "prod-k8s-ie"
  module_id  = "k8s-module"
  priority   = 0
}

Schema
Required

    context_id (String) ID of the context to attach

Optional

    module_id (String) ID of the module to attach the context to
    priority (Number) Priority of the context attachment. All the contexts attached to a stack are sorted by priority (lowest first), though values don't need to be unique. This ordering establishes precedence rules between contexts should there be a conflict and multiple contexts define the same value. Defaults to 0.
    stack_id (String) ID of the stack to attach the context to

Read-Only

    id (String) The ID of this resource.

Import

Import is supported using the following syntax:

terraform import spacelift_context_attachment.test_stack $CONTEXT_ID/$STACK_ID

On this page

    Example Usage
    Schema
    Import

Report an issue

METADATA:
  resource_type: spacelift_context_attachment
  provider: spacelift
  service: configuration
  description: Links contexts to stacks or modules with priority ordering
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_context_attachment" "RESOURCE_NAME" {
  context_id = CONTEXT_ID
  # One of the following must be set:
  stack_id   = STACK_ID    # Optional if module_id is set
  module_id  = MODULE_ID   # Optional if stack_id is set
  priority   = 0          # Optional, defaults to 0
}
```

ATTRIBUTES:
  required:
    context_id:
      type: String
      description: Context to attach
      validation: Must exist in Spacelift

  optional:
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist if module_id not set
      note: Mutually exclusive with module_id
      
    module_id:
      type: String
      description: Target module identifier
      validation: Must exist if stack_id not set
      note: Mutually exclusive with stack_id
      
    priority:
      type: Number
      description: Attachment priority order
      default: 0
      note: Lower values processed first

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  priority:
    - Determines context processing order
    - Lower numbers processed first
    - Values don't need to be unique
    - Used to resolve configuration conflicts
    
  attachment:
    - One context per attachment
    - One attachment per context-target pair
    - Target must be stack or module
    
  inheritance:
    - Configuration inherited by target
    - Priority affects conflict resolution
    - Later contexts override earlier ones
    
  validation:
    - Context must exist
    - Target must exist
    - Cannot attach to both stack and module

IMPORT_FORMAT: $CONTEXT_ID/$STACK_ID