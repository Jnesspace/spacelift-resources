METADATA:
  resource_type: spacelift_run
  provider: spacelift
  service: execution
  description: Programmatic run trigger for Spacelift stacks
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_run" "RESOURCE_NAME" {
  stack_id    = STACK_ID
  keepers     = {             # Optional
    key = VALUE
  }
  commit_sha  = COMMIT_SHA    # Optional
  proposed    = false         # Optional
}
```

ATTRIBUTES:
  required:
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist in Spacelift

  optional:
    keepers:
      type: Map[String]
      description: Change detection map
      note: Changes trigger resource recreation
      
    commit_sha:
      type: String
      description: Specific commit to run
      validation: Must be valid git SHA
      
    proposed:
      type: Boolean
      description: Run in proposed state
      default: false

    timeouts:
      type: Block
      description: Operation timeouts
      fields:
        create:
          type: String
          description: Creation timeout
          default: none

    wait:
      type: Block
      max: 1
      description: Run completion settings
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
      description: Triggered run identifier
      generated: true

BEHAVIOR:
  triggering:
    - Triggered by keeper changes
    - Can target specific commits
    - Supports proposed state
    
  execution:
    - Runs in stack's environment
    - Uses stack's worker pool
    - Inherits stack configuration
    
  states:
    terminal:
      - finished
      - failed
      - discarded
      - stopped
    review:
      - pending_review
      - unconfirmed
    active:
      - applying
      - destroying
      - performing
      - planning
    preparation:
      - initializing
      - preparing
      - preparing_apply
      - preparing_replan
    queuing:
      - queued
      - ready
      - replan_requested