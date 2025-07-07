# Resource: spacelift_gcp_service_account

## Description
A GCP service account that's automatically created and linked to a specific stack or module, providing credential-less authentication for Google Cloud Platform resources.

## Example Usage
```hcl
# Stack GCP service account
resource "spacelift_gcp_service_account" "production" {
  stack_id = spacelift_stack.gcp_infrastructure.id
  
  token_scopes = [
    "https://www.googleapis.com/auth/compute",
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/devstorage.full_control"
  ]
}

# Module GCP service account
resource "spacelift_gcp_service_account" "vpc_module" {
  module_id = spacelift_module.gcp_vpc.id
  
  token_scopes = [
    "https://www.googleapis.com/auth/compute"
  ]
}

# Grant permissions to the service account
resource "google_project_iam_member" "spacelift" {
  project = var.gcp_project_id
  role    = "roles/editor"
  member  = "serviceAccount:${spacelift_gcp_service_account.production.service_account_email}"
}
```

## Argument Reference

### Required Arguments
* `token_scopes` - (Required) List of OAuth 2.0 scopes for the service account tokens

### Optional Arguments
* `stack_id` - (Optional) ID of the stack to link the service account to. Mutually exclusive with `module_id`
* `module_id` - (Optional) ID of the module to link the service account to. Mutually exclusive with `stack_id`

### Read-Only Arguments
* `id` - Unique resource identifier
* `service_account_email` - Email address of the created GCP service account

## Import
```bash
# For stacks
terraform import spacelift_gcp_service_account.example stack/$STACK_ID

# For modules
terraform import spacelift_gcp_service_account.example module/$MODULE_ID
```

## Notes
* Either `stack_id` or `module_id` must be specified
* Service accounts are created automatically by Spacelift
* Use the service account email to grant GCP permissions
* Temporary credentials are injected during runs