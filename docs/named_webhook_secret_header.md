
spacelift_named_webhook_secret_header (Resource)

spacelift_named_webhook_secret_header represents secret key value combination used as a custom headerwhen delivering webhook requests. It depends on spacelift_named_webhook resource which should exist.
Schema
Required

    key (String) key for the header
    value (String, Sensitive) value for the header
    webhook_id (String) ID of the stack on which the environment variable is defined

Read-Only

    id (String) The ID of this resource.

