# Resource: spacelift_environment_variable

## Description
Environment variables that can be defined on contexts, stacks, or modules to pass configuration and secrets to Spacelift runs.

## Example Usage
```hcl
# Stack environment variable
resource "spacelift_environment_variable" "api_key" {
  stack_id    = spacelift_stack.api.id
  name        = "API_KEY"
  value       = var.api_key
  write_only  = true
  description = "API key for external service"
}

# Context environment variable
resource "spacelift_environment_variable" "region" {
  context_id  = spacelift_context.production.id
  name        = "AWS_DEFAULT_REGION"
  value       = "us-east-1"
  write_only  = false
  description = "Default AWS region"
}

# Module environment variable
resource "spacelift_environment_variable" "vpc_cidr" {
  module_id   = spacelift_module.vpc.id
  name        = "VPC_CIDR"
  value       = "10.0.0.0/16"
  description = "VPC CIDR block for module testing"
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Environment variable name

### Optional Arguments
* `context_id` - (Optional) Context to attach the variable to. Mutually exclusive with `stack_id` and `module_id`
* `stack_id` - (Optional) Stack to attach the variable to. Mutually exclusive with `context_id` and `module_id`
* `module_id` - (Optional) Module to attach the variable to. Mutually exclusive with `context_id` and `stack_id`
* `value` - (Optional) Environment variable value. Defaults to empty string
* `write_only` - (Optional) Whether the value is secret. Defaults to `true`
* `description` - (Optional) Human-readable description

### Read-Only Arguments
* `id` - Unique resource identifier
* `checksum` - SHA-256 checksum of the value

## Import
```bash
# For context variables
terraform import spacelift_environment_variable.example context/$CONTEXT_ID/$VARIABLE_NAME

# For stack variables  
terraform import spacelift_environment_variable.example stack/$STACK_ID/$VARIABLE_NAME

# For module variables
terraform import spacelift_environment_variable.example module/$MODULE_ID/$VARIABLE_NAME
```

## Notes
* Exactly one of `context_id`, `stack_id`, or `module_id` must be specified
* Write-only variables are masked in logs and outputs
* Variables inherit from attached contexts based on priority