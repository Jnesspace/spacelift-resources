# Resource: spacelift_context

## Description
A context is a collection of configuration elements (environment variables, mounted files, and hooks) that can be attached to multiple stacks or modules to share common configuration.

## Example Usage
```hcl
# Production context with shared configuration
resource "spacelift_context" "production" {
  name        = "production-config"
  description = "Shared production environment configuration"
  labels      = ["production", "shared"]
  
  # Lifecycle hooks
  before_init = ["echo 'Starting production deployment'"]
  after_apply = ["./notify-deployment.sh"]
}

# Development context
resource "spacelift_context" "development" {
  name        = "dev-config"
  description = "Development environment configuration"
  space_id    = spacelift_space.development.id
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique context identifier within the account

### Optional Arguments
* `description` - (Optional) Human-readable description of the context
* `labels` - (Optional) Set of labels for context classification and auto-attachment
* `space_id` - (Optional) ID of the space the context belongs to. Defaults to root space
* `before_init` - (Optional) List of commands to run before initialization
* `after_init` - (Optional) List of commands to run after initialization
* `before_plan` - (Optional) List of commands to run before planning
* `after_plan` - (Optional) List of commands to run after planning
* `before_apply` - (Optional) List of commands to run before applying
* `after_apply` - (Optional) List of commands to run after applying
* `before_destroy` - (Optional) List of commands to run before destroying
* `after_destroy` - (Optional) List of commands to run after destroying
* `before_perform` - (Optional) List of commands to run before performing
* `after_perform` - (Optional) List of commands to run after performing
* `after_run` - (Optional) List of commands to run after any run completes

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_context.example $CONTEXT_ID
```

## Notes
* Contexts can be attached to multiple stacks/modules using `spacelift_context_attachment`
* Labels starting with "autoattach:" enable automatic attachment to matching resources
* Hook scripts run in the order they are defined