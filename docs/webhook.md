METADATA:
  resource_type: spacelift_webhook
  provider: spacelift
  service: notifications
  description: Webhook endpoint for Spacelift run state change notifications
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_webhook" "RESOURCE_NAME" {
  endpoint  = WEBHOOK_URL
  stack_id  = STACK_ID      # Optional if module_id is set
  module_id = MODULE_ID     # Optional if stack_id is set
  enabled   = true         # Optional
  secret    = "mysecret"   # Optional
}
```

ATTRIBUTES:
  required:
    endpoint:
      type: String
      description: Destination URL for POST requests
      validation: Must be valid URL

  optional:
    enabled:
      type: Boolean
      description: Controls webhook activation
      default: true
      
    module_id:
      type: String
      description: Source module identifier
      validation: Must exist if stack_id not set
      note: Mutually exclusive with stack_id
      
    stack_id:
      type: String
      description: Source stack identifier
      validation: Must exist if module_id not set
      note: Mutually exclusive with module_id
      
    secret:
      type: String
      description: Request signature secret
      sensitive: true
      note: Not retrievable after creation

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  triggers:
    - Sends POST requests on run state changes
    - Only triggers for specified stack or module
    
  security:
    - Supports request signing with secret
    - Secret is never returned in state
    - Uses HTTPS for endpoint communication
    
  validation:
    - Requires either stack_id or module_id
    - Cannot specify both stack_id and module_id
    - Endpoint must be accessible

IMPORT_FORMAT: stack/$STACK_ID/$WEBHOOK_ID