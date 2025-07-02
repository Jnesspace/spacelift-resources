METADATA:
  resource_type: spacelift_stack_activator
  provider: spacelift
  service: lifecycle
  description: Controls stack activation state
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_stack_activator" "RESOURCE_NAME" {
  stack_id = STACK_ID
  enabled  = true    # Required, controls stack state
}
```

ATTRIBUTES:
  required:
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist in Spacelift
      
    enabled:
      type: Boolean
      description: Stack activation state
      validation: Must be true or false

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  activation:
    - Controls stack processing
    - Enables/disables operations
    - Affects all stack functions
    
  state_management:
    - Independent of stack definition
    - Immediate state change
    - Reversible operation
    
  effects:
    enabled_true:
      - Allows stack operations
      - Enables scheduled runs
      - Processes webhooks
      - Allows manual triggers
      
    enabled_false:
      - Blocks stack operations
      - Disables scheduled runs
      - Ignores webhooks
      - Prevents manual triggers

PATTERNS:
  maintenance:
    example:
      - Disable for maintenance
      - Perform updates
      - Re-enable when ready
    benefit: Controlled maintenance windows
    
  automation:
    example:
      - Conditional activation
      - Environment-based control
      - Scheduled availability
    benefit: Automated stack management