# Resource: spacelift_webhook

## Description
**DEPRECATED**: Configures webhook endpoints for stack notifications. This resource is deprecated in favor of `spacelift_named_webhook` which provides better organization and management capabilities.

## Example Usage
```hcl
# Stack webhook (deprecated)
resource "spacelift_webhook" "slack_notifications" {
  stack_id = spacelift_stack.production.id
  endpoint = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  secret   = var.slack_webhook_secret
}

# Module webhook (deprecated)
resource "spacelift_webhook" "module_alerts" {
  module_id = spacelift_module.vpc.id
  endpoint  = "https://monitoring.company.com/webhook"
  secret    = var.monitoring_secret
}
```

## Argument Reference

### Required Arguments
* `endpoint` - (Required) Destination URL for webhook requests

### Optional Arguments
* `stack_id` - (Optional) ID of the stack to attach webhook to. Mutually exclusive with `module_id`
* `module_id` - (Optional) ID of the module to attach webhook to. Mutually exclusive with `stack_id`
* `secret` - (Optional) Secret for request signature verification

### Read-Only Arguments
* `id` - Unique resource identifier

## Migration to spacelift_named_webhook
```hcl
# Old (deprecated)
resource "spacelift_webhook" "old" {
  stack_id = spacelift_stack.app.id
  endpoint = "https://hooks.slack.com/webhook"
  secret   = var.webhook_secret
}

# New (recommended)
resource "spacelift_named_webhook" "new" {
  name     = "slack-notifications"
  endpoint = "https://hooks.slack.com/webhook"
  space_id = "root"
  enabled  = true
  secret   = var.webhook_secret
}

# Then reference in notification policies
resource "spacelift_policy" "notification" {
  type = "NOTIFICATION"
  name = "slack-notifications"
  body = file("${path.module}/notification-policy.rego")
}
```

## Notes
* **This resource is deprecated** - use `spacelift_named_webhook` instead
* Either `stack_id` or `module_id` must be specified
* Named webhooks provide better reusability and organization