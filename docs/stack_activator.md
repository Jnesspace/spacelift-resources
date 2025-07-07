# Resource: spacelift_stack_activator

## Description
Controls the activation state of a Spacelift stack, allowing you to enable or disable stack operations programmatically.

## Example Usage
```hcl
# Enable a stack
resource "spacelift_stack_activator" "production" {
  stack_id = spacelift_stack.production.id
  enabled  = true
}

# Conditionally activate stack based on environment
resource "spacelift_stack_activator" "feature_stack" {
  stack_id = spacelift_stack.feature.id
  enabled  = var.enable_feature_stack
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to control
* `enabled` - (Required) Whether the stack should be enabled

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Disabled stacks cannot execute runs, tasks, or scheduled operations
* This resource provides independent control over stack state
* Useful for maintenance windows or conditional stack activation