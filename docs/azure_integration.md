# Resource: spacelift_azure_integration

## Description
Configures Azure integration for secure, credential-less access to Azure resources using Azure Active Directory service principals. Integrations can be shared across multiple stacks and modules.

## Example Usage
```hcl
# Basic Azure integration
resource "spacelift_azure_integration" "production" {
  name         = "production-azure"
  tenant_id    = "12345678-1234-1234-1234-123456789012"
  client_id    = "87654321-4321-4321-4321-210987654321"
  client_secret = var.azure_client_secret
  space_id     = spacelift_space.production.id
}

# Azure integration with custom configuration
resource "spacelift_azure_integration" "development" {
  name         = "dev-azure"
  tenant_id    = var.azure_tenant_id
  client_id    = var.azure_client_id
  client_secret = var.azure_client_secret
  labels       = ["azure", "development"]
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique integration identifier
* `tenant_id` - (Required) Azure Active Directory tenant ID
* `client_id` - (Required) Azure service principal client ID
* `client_secret` - (Required) Azure service principal client secret

### Optional Arguments
* `space_id` - (Optional) ID of the space the integration belongs to. Defaults to `"root"`
* `labels` - (Optional) Set of labels for integration classification

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_azure_integration.example $INTEGRATION_ID
```

## Notes
* Integrations can be attached to multiple stacks/modules using `spacelift_azure_integration_attachment`
* Service principal must have appropriate permissions for target Azure resources
* Client secret is stored securely and not returned after creation