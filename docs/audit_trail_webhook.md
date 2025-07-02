# spacelift_audit_trail_webhook

METADATA:
  resource_type: spacelift_audit_trail_webhook
  provider: spacelift
  service: webhooks
  description: Webhook endpoint for Spacelift audit event notifications
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_audit_trail_webhook" "RESOURCE_NAME" {
  endpoint = URL_ENDPOINT
  enabled  = BOOLEAN
  secret   = SECRET_KEY
}
```

ATTRIBUTES:
  required:
    enabled:
      type: Boolean
      description: Controls webhook activation status
      default: false
      validation: none
    
    endpoint:
      type: String
      description: Destination URL for POST requests
      validation: Must be valid URL format
    
    secret:
      type: String
      description: Secret key for request authentication
      sensitive: true
      validation: none

  optional:
    custom_headers:
      type: Map[String]
      description: Additional HTTP headers for requests
      default: {}
      validation: Valid HTTP headers
    
    include_runs:
      type: Boolean
      description: Include run data in webhook payload
      default: false
      validation: none

  computed:
    id:
      type: String
      description: Unique identifier for the resource
      generated: true

BEHAVIOR:
  - Webhook is inactive until enabled=true is set
  - Secret is never returned in state or outputs
  - POST requests include X-Spacelift-Signature header
