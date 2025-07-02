METADATA:
  resource_type: spacelift_azure_integration_attachment
  provider: spacelift
  service: cloud_integration
  description: Links Azure integrations to stacks or modules with access control
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_azure_integration_attachment" "RESOURCE_NAME" {
  integration_id  = INTEGRATION_ID
  # One of the following must be set:
  stack_id       = STACK_ID       # Optional if module_id is set
  module_id      = MODULE_ID      # Optional if stack_id is set
  
  # Access control (both default to true):
  read           = true           # Optional
  write          = true          # Optional
  
  # Optional override:
  subscription_id = SUBSCRIPTION_ID  # Overrides integration default
}
```

ATTRIBUTES:
  required:
    integration_id:
      type: String
      description: Azure integration identifier
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
      
    subscription_id:
      type: String
      description: Azure subscription override
      validation: Valid UUID format
      note: Overrides integration default

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
    
  subscription:
    - Uses integration default
    - Supports per-attachment override
    - Must be valid subscription
    
  validation:
    - Integration must exist
    - Target must exist
    - Cannot attach to both stack and module
    - Subscription must be accessible
    
  usage:
    stack:
      - Infrastructure deployments
      - Resource management
      - State storage
    module:
      - Module testing
      - Resource verification
      - Dependency checks

PATTERNS:
  readonly:
    example:
      read: true
      write: false
    benefit: Safe resource inspection
    
  writeonly:
    example:
      read: false
      write: true
    benefit: Controlled resource modification
    
  subscription:
    example:
      subscription_id override
    benefit: Resource isolation

IMPORT_FORMAT:
  stack: $INTEGRATION_ID/$STACK_ID
  module: $INTEGRATION_ID/$MODULE_ID