METADATA:
  resource_type: spacelift_stack_destructor
  provider: spacelift
  service: lifecycle
  description: Managed destruction of stack resources
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_stack_destructor" "RESOURCE_NAME" {
  stack_id    = STACK_ID
  deactivated = false      # Optional
  
  depends_on = [
    REQUIRED_RESOURCES     # Dependencies needed for cleanup
  ]
}
```

ATTRIBUTES:
  required:
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist in Spacelift

  optional:
    deactivated:
      type: Boolean
      description: Prevents resource deletion
      default: false
      note: Set true to preserve resources
      
    timeouts:
      type: Block
      description: Operation timeouts
      fields:
        delete:
          type: String
          description: Deletion timeout
          default: none

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  destruction:
    - Destroys stack resources
    - Deletes stack afterwards
    - Can be deactivated to preserve
    
  dependencies:
    - Requires explicit depends_on
    - Maintains required resources
    - Ensures clean destruction
    
  safety:
    - Can preserve resources
    - Configurable timeouts
    - Requires dependency setup
    
  workflow:
    - Executes destroy command
    - Waits for completion
    - Handles cleanup order

IMPORTANT_NOTES:
  - Set deactivated=true to preserve resources when removing destructor
  - Use depends_on for required cleanup resources
  - Dependencies might include:
    - Environment variables
    - Role assignments
    - Integrations
    - Mounted files
    
PATTERNS:
  cleanup:
    example:
      - Environment setup
      - Resource destruction
      - Stack deletion
    benefit: Ensures complete cleanup
    
  preservation:
    example:
      - Set deactivated=true
      - Remove destructor
      - Resources preserved
    benefit: Prevents accidental deletion