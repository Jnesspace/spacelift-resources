METADATA:
  resource_type: spacelift_aws_integration_attachment
  provider: spacelift
  service: cloud_integration
  description: Links AWS integrations to stacks or modules with access control
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_aws_integration_attachment" "RESOURCE_NAME" {
  integration_id = INTEGRATION_ID
  # One of the following must be set:
  stack_id      = STACK_ID       # Optional if module_id is set
  module_id     = MODULE_ID      # Optional if stack_id is set
  
  # Access control (both default to true):
  read          = true           # Optional
  write         = true          # Optional
  
  depends_on = [
    REQUIRED_RESOURCES          # IAM role must exist first
  ]
}
```

ATTRIBUTES:
  required:
    integration_id:
      type: String
      description: AWS integration identifier
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
      
    read:
      type: Boolean
      description: Enable read operations
      default: true
      
    write:
      type: Boolean
      description: Enable write operations
      default: true

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    attachment_id:
      type: String
      description: Internal attachment identifier
      generated: true

BEHAVIOR:
  permissions:
    - Separate read/write control
    - Both default to enabled
    - Can restrict to read-only
    - Can restrict to write-only
    
  validation:
    - Tests role assumption
    - Requires existing role
    - Verifies permissions
    - Checks target exists
    
  dependencies:
    - IAM role must exist first
    - Integration must be valid
    - Target must be valid
    
  usage:
    stack:
      - Infrastructure deployments
      - Resource management
      - State storage
    module:
      - Module testing
      - Resource verification
      - Dependency checks

IMPORTANT_NOTES:
  - Use depends_on to ensure role exists
  - Role assumption tested during attachment
  - One attachment per integration-target pair
  - Cannot attach to both stack and module

IMPORT_FORMAT: $INTEGRATION_ID/$TARGET_TYPE/$TARGET_ID

