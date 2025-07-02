METADATA:
  resource_type: spacelift_task
  provider: spacelift
  service: core
  description: Task execution configuration for Spacelift stacks
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_task" "RESOURCE_NAME" {
  stack_id = STACK_ID
  command  = COMMAND
  init     = true    # Optional
}
```

ATTRIBUTES:
  required:
    command:
      type: String
      description: Command to execute
      validation: Must be valid shell command
    
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist in Spacelift

  optional:
    init:
      type: Boolean
      description: Controls stack initialization
      default: true
      
    keepers:
      type: Map[String]
      description: Values that trigger resource recreation
      default: {}

    timeouts:
      type: Block
      description: Operation timeout settings
      fields:
        create:
          type: String
          description: Resource creation timeout
          default: none

    wait:
      type: Block
      max: 1
      description: Task completion wait configuration
      fields:
        continue_on_state:
          type: Set[String]
          description: States allowing continuation
          default: ["finished"]
          allowed_values:
            - applying
            - canceled
            - confirmed
            - destroying
            - discarded
            - failed
            - finished
            - initializing
            - pending_review
            - performing
            - planning
            - preparing_apply
            - preparing_replan
            - preparing
            - queued
            - ready
            - replan_requested
            - skipped
            - stopped
            - unconfirmed
        
        continue_on_timeout:
          type: Boolean
          description: Continue after timeout
          default: false
        
        disabled:
          type: Boolean
          description: Disable wait behavior
          default: false

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  execution:
    - Runs in stack's environment
    - Can initialize stack before execution
    - Supports command execution monitoring
    
  lifecycle:
    - Can be triggered by keeper changes
    - Supports timeout configuration
    - Can wait for completion states
    
  states:
    terminal:
      - finished
      - failed
      - discarded
      - stopped
    transitional:
      - initializing
      - planning
      - applying
      - performing
    review:
      - pending_review
      - unconfirmed