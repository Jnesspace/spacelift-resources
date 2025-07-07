# Resource: spacelift_security_email

## Description
Configures security notification email addresses to receive alerts about important security events in your Spacelift account.

## Example Usage
```hcl
# Primary security team email
resource "spacelift_security_email" "security_team" {
  email = "security@company.com"
}

# Secondary security contact
resource "spacelift_security_email" "security_backup" {
  email = "security-backup@company.com"
}
```

## Argument Reference

### Required Arguments
* `email` - (Required) Email address to receive security notifications

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Multiple security emails can be configured
* Security emails receive notifications about account security events
* Email addresses must be valid and accessible