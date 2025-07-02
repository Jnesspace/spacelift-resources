
spacelift_security_email (Resource)

spacelift_security_email represents an email address that receives notifications about security issues in Spacelift.
Example Usage

resource "spacelift_security_email" "example" {
  email = "user@example.com"
}

Schema
Required

    email (String) Email address to which the security notifications are sent

Read-Only

    id (String) The ID of this resource.

On this page

    Example Usage
    Schema

Report an issue 