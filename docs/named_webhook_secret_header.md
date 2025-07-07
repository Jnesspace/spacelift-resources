# Resource: spacelift_named_webhook_secret_header

## Description
Configures custom HTTP headers with secrets for named webhooks, allowing secure authentication and custom metadata to be sent with webhook requests.

## Example Usage
```hcl
# API key header for webhook authentication
resource "spacelift_named_webhook_secret_header" "api_key" {
  webhook_id = spacelift_named_webhook.monitoring.id
  key        = "X-API-Key"
  value      = var.webhook_api_key
}

# Custom authorization header
resource "spacelift_named_webhook_secret_header" "auth" {
  webhook_id = spacelift_named_webhook.slack_alerts.id
  key        = "Authorization"
  value      = "Bearer ${var.slack_bearer_token}"
}

# Custom metadata header
resource "spacelift_named_webhook_secret_header" "source" {
  webhook_id = spacelift_named_webhook.monitoring.id
  key        = "X-Source"
  value      = "spacelift-production"
}
```

## Argument Reference

### Required Arguments
* `webhook_id` - (Required) ID of the named webhook to add the header to
* `key` - (Required) HTTP header name
* `value` - (Required) HTTP header value (sensitive)

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Header values are stored securely and treated as sensitive
* Multiple headers can be added to the same webhook
* Headers are included in all webhook requests from the named webhook

