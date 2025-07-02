METADATA:
  resource_type: spacelift_environment_variable
  provider: spacelift
  service: configuration
  description: Environment variable configuration for Spacelift entities
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_environment_variable" "RESOURCE_NAME" {
  name        = ENV_VAR_NAME
  value       = ENV_VAR_VALUE
  # One of the following must be set:
  context_id  = CONTEXT_ID    # Optional
  module_id   = MODULE_ID     # Optional
  stack_id    = STACK_ID      # Optional
  write_only  = true         # Optional
  description = DESCRIPTION   # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Environment variable identifier
      validation: Must be valid environment variable name

  optional:
    context_id:
      type: String
      description: Target context identifier
      validation: Must exist if module_id and stack_id not set
      note: Mutually exclusive with module_id and stack_id
      
    module_id:
      type: String
      description: Target module identifier
      validation: Must exist if context_id and stack_id not set
      note: Mutually exclusive with context_id and stack_id
      
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist if context_id and module_id not set
      note: Mutually exclusive with context_id and module_id
      
    value:
      type: String
      description: Environment variable value
      default: ""
      sensitive: true
      
    write_only:
      type: Boolean
      description: Marks value as secret
      default: true
      
    description:
      type: String
      description: Human-readable variable description
      default: ""

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    checksum:
      type: String
      description: SHA-256 hash of value
      generated: true

BEHAVIOR:
  scope:
    - Can be defined on contexts, modules, or stacks
    - Must specify exactly one target entity
    - Inheritance follows context attachment hierarchy
    
  security:
    - Write-only values are masked in logs
    - Values are encrypted at rest
    - Checksums enable value verification
    
  validation:
    - Requires valid environment variable name
    - Cannot have duplicate names in same scope
    - Target entity must exist

IMPORT_FORMAT:
  context: context/$CONTEXT_ID/$ENV_VAR_NAME
  module: module/$MODULE_ID/$ENV_VAR_NAME
  stack: stack/$STACK_ID/$ENV_VAR_NAME