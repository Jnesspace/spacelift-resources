# Resource: spacelift_stack_destructor

## Description
Manages the safe destruction of stack resources before deleting the stack itself. This resource ensures proper cleanup order and prevents resource leaks.

## Example Usage
```hcl
# Basic stack destructor
resource "spacelift_stack_destructor" "app_cleanup" {
  stack_id = spacelift_stack.app.id
  
  depends_on = [
    spacelift_environment_variable.credentials,
    spacelift_aws_role.deployment_role,
    spacelift_context_attachment.production
  ]
}

# Destructor with preservation option
resource "spacelift_stack_destructor" "safe_cleanup" {
  stack_id    = spacelift_stack.database.id
  deactivated = true  # Prevents actual destruction
  
  timeouts {
    delete = "30m"
  }
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to destroy

### Optional Arguments
* `deactivated` - (Optional) If true, prevents actual resource destruction. Defaults to `false`
* `timeouts` - (Optional) Resource operation timeouts
  * `delete` - (Optional) Destruction timeout

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Use `depends_on` to ensure required resources exist during destruction
* Set `deactivated = true` to preserve resources when removing the destructor
* The destructor runs a destroy operation before deleting the stack
* Destruction order is critical - define dependencies carefully