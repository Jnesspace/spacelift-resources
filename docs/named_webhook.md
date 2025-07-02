METADATA:
  resource_type: spacelift_named_webhook
  provider: spacelift
  service: notifications
  description: Named webhook endpoint for notification policy routing
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_named_webhook" "RESOURCE_NAME" {
  name      = WEBHOOK_NAME
  endpoint  = ENDPOINT_URL
  space_id  = SPACE_ID
  enabled   = true          # Optional
  secret    = SECRET_KEY    # Optional
  labels    = [LABELS]      # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Webhook identifier
      validation: Must be unique in space
      note: Used for ID generation
      
    endpoint:
      type: String
      description: Destination URL
      validation: Must be valid URL
      
    space_id:
      type: String
      description: Target space identifier
      validation: Must exist in Spacelift
      
    enabled:
      type: Boolean
      description: Controls webhook activation
      default: false

  optional:
    secret:
      type: String
      description: Request signature secret
      sensitive: true
      note: Not retrievable after creation
      
    labels:
      type: Set[String]
      description: Classification tags
      default: []
      note: Used in policies and filtering

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  routing:
    - Referenced in notification policies
    - Routes messages based on policies
    - Supports message filtering
    
  security:
    - Optional request signing
    - Secret never returned in state
    - Uses HTTPS for endpoints
    
  management:
    - Can be enabled/disabled
    - Organized by spaces
    - Labeled for organization
    
  validation:
    - Requires valid endpoint URL
    - Name must be unique in space
    - Space must exist
    - Labels must be valid strings

USAGE:
  notification_policy:
    - Reference webhook by name
    - Configure message routing
    - Apply filtering rules
    - Control message delivery