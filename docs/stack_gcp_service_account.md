# Resource: spacelift_stack_gcp_service_account

## Description
**DEPRECATED**: GCP service account linked to a specific stack. This resource is deprecated in favor of `spacelift_gcp_service_account` which provides the same functionality with a clearer name.

## Example Usage
```hcl
# Stack GCP service account (deprecated)
resource "spacelift_stack_gcp_service_account" "production" {
  stack_id = spacelift_stack.gcp_infrastructure.id
  
  token_scopes = [
    "https://www.googleapis.com/auth/compute",
    "https://www.googleapis.com/auth/cloud-platform"
  ]
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to link the service account to
* `token_scopes` - (Required) List of OAuth 2.0 scopes for the service account tokens

### Read-Only Arguments
* `id` - Unique resource identifier
* `service_account_email` - Email address of the created GCP service account

## Migration to spacelift_gcp_service_account
```hcl
# Old (deprecated)
resource "spacelift_stack_gcp_service_account" "old" {
  stack_id = spacelift_stack.app.id
  token_scopes = ["https://www.googleapis.com/auth/compute"]
}

# New (recommended)
resource "spacelift_gcp_service_account" "new" {
  stack_id = spacelift_stack.app.id
  token_scopes = ["https://www.googleapis.com/auth/compute"]
}
```

## Notes
* **This resource is deprecated** - use `spacelift_gcp_service_account` instead
* Functionality is identical between the old and new resources
* Service accounts are created automatically by Spacelift