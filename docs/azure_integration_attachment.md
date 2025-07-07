# Resource: spacelift_azure_integration_attachment

## Description
Attaches an Azure integration to a stack or module, enabling Azure credential management for Azure resource deployments.

## Example Usage
```hcl
# Attach Azure integration to stack
resource "spacelift_azure_integration_attachment" "production" {
  integration_id = spacelift_azure_integration.production.id
  stack_id       = spacelift_stack.azure_infrastructure.id
}

# Attach Azure integration to module
resource "spacelift_azure_integration_attachment" "networking_module" {
  integration_id = spacelift_azure_integration.shared.id
  module_id      = spacelift_module.azure_network.id
}
```

## Argument Reference

### Required Arguments
* `integration_id` - (Required) ID of the Azure integration to attach

### Optional Arguments
* `stack_id` - (Optional) ID of the stack to attach to. Mutually exclusive with `module_id`
* `module_id` - (Optional) ID of the module to attach to. Mutually exclusive with `stack_id`

### Read-Only Arguments
* `id` - Unique resource identifier
* `attachment_id` - Internal attachment identifier

## Import
```bash
terraform import spacelift_azure_integration_attachment.example $INTEGRATION_ID/$STACK_ID
```

## Notes
* Either `stack_id` or `module_id` must be specified
* Azure credentials are automatically injected during runs
* Service principal permissions are tested during attachment creation