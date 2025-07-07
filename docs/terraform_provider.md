# Resource: spacelift_terraform_provider

## Description
Configures a custom Terraform provider for use in Spacelift stacks, allowing integration with provider registries or custom provider builds.

## Example Usage
```hcl
# Public registry provider
resource "spacelift_terraform_provider" "aws_custom" {
  type         = "aws"
  space_id     = spacelift_space.production.id
  public       = true
  description  = "Custom AWS provider configuration"
  labels       = ["aws", "custom"]
}

# Private provider with GPG key
resource "spacelift_terraform_provider" "internal" {
  type              = "internal"
  space_id          = "root"
  public            = false
  description       = "Internal company provider"
  gpg_key_id        = "1234567890ABCDEF"
  gpg_ascii_armor   = file("${path.module}/provider-key.asc")
}
```

## Argument Reference

### Required Arguments
* `type` - (Required) Provider type/name

### Optional Arguments
* `space_id` - (Optional) ID of the space the provider belongs to. Defaults to `"root"`
* `public` - (Optional) Whether the provider is publicly accessible. Defaults to `false`
* `description` - (Optional) Human-readable description of the provider
* `labels` - (Optional) Set of labels for provider classification
* `gpg_key_id` - (Optional) GPG key ID for provider verification
* `gpg_ascii_armor` - (Optional) GPG public key in ASCII armor format

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Custom providers can override default registry providers
* GPG verification is recommended for security
* Public providers can be shared across spaces