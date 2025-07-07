# Resource: spacelift_audit_trail_webhook

## Description
Configures a webhook endpoint to receive audit trail events from Spacelift, enabling integration with external logging, monitoring, or compliance systems.

## Example Usage
```hcl
# Audit trail webhook for compliance logging
resource "spacelift_audit_trail_webhook" "compliance" {
  endpoint    = "https://compliance.company.com/audit/spacelift"
  secret      = var.audit_webhook_secret
  include_runs = true
  enabled     = true
}

# Security monitoring webhook
resource "spacelift_audit_trail_webhook" "security" {
  endpoint    = "https://security-monitoring.company.com/webhooks/spacelift"
  secret      = var.security_webhook_secret
  include_runs = false
  enabled     = true
}
```

## Argument Reference

### Required Arguments
* `endpoint` - (Required) Destination URL for audit trail webhook requests
* `enabled` - (Required) Whether the webhook is active

### Optional Arguments
* `secret` - (Optional) Secret for request signature verification
* `include_runs` - (Optional) Whether to include run events in audit trail. Defaults to `true`

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Audit trail webhooks receive all account-level security and administrative events
* Secret is used for request signing and is not returned after creation
* Run events can be excluded to reduce webhook volume
