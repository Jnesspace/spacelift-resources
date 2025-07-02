METADATA:
  resource_type: spacelift_policy_attachment
  provider: spacelift
  service: governance
  description: Links policies to stacks or modules for rule enforcement
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_policy_attachment" "RESOURCE_NAME" {
  policy_id = POLICY_ID
  # One of the following must be set:
  stack_id  = STACK_ID   # Optional if module_id is set
  module_id = MODULE_ID  # Optional if stack_id is set
}
```

ATTRIBUTES:
  required:
    policy_id:
      type: String
      description: Policy to attach
      validation: Must exist and not be LOGIN type

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

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  attachment:
    - One policy per stack/module
    - LOGIN policies cannot be attached
    - Target must be either stack or module
    
  validation:
    - Policy must exist
    - Target must exist
    - Cannot attach same policy twice
    
  scope:
    - Policies affect attached resource
    - Stack attachments affect stack runs
    - Module attachments affect module tests

LIMITATIONS:
  - LOGIN policies are global only
  - Cannot attach to both stack and module
  - One policy per target resource

IMPORT_FORMAT: $POLICY_ID/$STACK_ID