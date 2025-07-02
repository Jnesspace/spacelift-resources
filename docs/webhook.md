
spacelift_webhook (Resource)

spacelift_webhook represents a webhook endpoint to which Spacelift sends the POST request about run state changes.
Example Usage

resource "spacelift_webhook" "webhook" {
  endpoint = "https://example.com/webhooks"
  stack_id = "k8s-core"
}

Schema
Required

    endpoint (String) endpoint to send the POST request to

Optional

    enabled (Boolean) enables or disables sending webhooks. Defaults to true.
    module_id (String) ID of the module which triggers the webhooks
    secret (String, Sensitive) secret used to sign each POST request so you're able to verify that the request comes from us. Defaults to an empty value. Note that once it's created, it will be just an empty string in the state due to security reasons.
    stack_id (String) ID of the stack which triggers the webhooks

Read-Only

    id (String) The ID of this resource.

Import

Import is supported using the following syntax:

terraform import spacelift_webhook.webhook stack/$STACK_ID/$WEBHOOK_ID

On this page

    Example Usage
    Schema
    Import

Report an issue 