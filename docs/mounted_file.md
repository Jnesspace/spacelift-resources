# Resource: spacelift_mounted_file

## Description
Files that are mounted in the workspace during runs, allowing you to provide configuration files, certificates, or other assets to your infrastructure code.

## Example Usage
```hcl
# Stack mounted file
resource "spacelift_mounted_file" "kubeconfig" {
  stack_id      = spacelift_stack.k8s.id
  relative_path = "kubeconfig"
  content       = filebase64("${path.module}/kubeconfig.yaml")
  write_only    = true
  description   = "Kubernetes configuration file"
}

# Context mounted file
resource "spacelift_mounted_file" "ca_cert" {
  context_id    = spacelift_context.production.id
  relative_path = "certs/ca.pem"
  content       = filebase64("${path.module}/ca.pem")
  write_only    = false
  description   = "CA certificate for production environment"
}

# Module mounted file
resource "spacelift_mounted_file" "module_config" {
  module_id     = spacelift_module.vpc.id
  relative_path = "config.json"
  content       = filebase64("${path.module}/module-config.json")
  description   = "Module test configuration"
}
```

## Argument Reference

### Required Arguments
* `relative_path` - (Required) Path where the file will be mounted in the workspace (relative to `/mnt/workspace/`)
* `content` - (Required) Base64-encoded file content

### Optional Arguments
* `context_id` - (Optional) Context to attach the file to. Mutually exclusive with `stack_id` and `module_id`
* `stack_id` - (Optional) Stack to attach the file to. Mutually exclusive with `context_id` and `module_id`
* `module_id` - (Optional) Module to attach the file to. Mutually exclusive with `context_id` and `stack_id`
* `write_only` - (Optional) Whether the content can be read back outside runs. Defaults to `true`
* `description` - (Optional) Human-readable description

### Read-Only Arguments
* `id` - Unique resource identifier
* `checksum` - SHA-256 checksum of the content

## Import
```bash
# For context files
terraform import spacelift_mounted_file.example context/$CONTEXT_ID/$FILE_ID

# For stack files
terraform import spacelift_mounted_file.example stack/$STACK_ID/$FILE_ID

# For module files
terraform import spacelift_mounted_file.example module/$MODULE_ID/$FILE_ID
```

## Notes
* Exactly one of `context_id`, `stack_id`, or `module_id` must be specified
* Files are mounted under `/mnt/workspace/` in the run environment
* Use `filebase64()` function to encode file contents