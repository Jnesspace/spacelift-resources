# Resource: spacelift_named_webhook

## Description
A named webhook endpoint that can be referenced in notification policies to route messages to external systems for alerts, integrations, and monitoring.

## Example Usage
```hcl
# Slack webhook
resource "spacelift_named_webhook" "slack_alerts" {
  name     = "slack-notifications"
  endpoint = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  space_id = spacelift_space.production.id
  enabled  = true
  secret   = var.slack_webhook_secret
  labels   = ["slack", "alerts"]
}

# Custom webhook endpoint
resource "spacelift_named_webhook" "monitoring" {
  name     = "monitoring-webhook"
  endpoint = "https://monitoring.company.com/spacelift-webhook"
  space_id = "root"
  enabled  = true
  labels   = ["monitoring", "production"]
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique webhook identifier within the space
* `endpoint` - (Required) Destination URL for webhook requests
* `space_id` - (Required) ID of the space the webhook belongs to
* `enabled` - (Required) Whether the webhook is active

### Optional Arguments
* `secret` - (Optional) Secret for request signature verification
* `labels` - (Optional) Set of labels for webhook classification

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Named webhooks are referenced in notification policies by name
* Secret is used for request signing and is not returned after creation
* Labels can be used for filtering and organization