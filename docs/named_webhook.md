
spacelift_named_webhook (Resource)

spacelift_named_webhook represents a named webhook endpoint used for creating webhookswhich are referred to in Notification policies to route messages.
Schema
Required

    enabled (Boolean) enables or disables sending webhooks.
    endpoint (String) endpoint to send the requests to
    name (String) the name for the webhook which will also be used to generate the id
    space_id (String) ID of the space the webhook is in

Optional

    labels (Set of String) labels for the webhook to use when referring in policies or filtering them
    secret (String, Sensitive) secret used to sign each request so you're able to verify that the request comes from us. Defaults to an empty value. Note that once it's created, it will be just an empty string in the state due to security reasons.

Read-Only

    id (String) The ID of this resource.

On this page

    Schema

Report an issue 