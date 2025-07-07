# Resource: spacelift_scheduled_delete_stack

## Description
Schedules the automatic deletion of a stack at a specified time, useful for temporary environments or cleanup automation.

## Example Usage
```hcl
# Schedule stack deletion in 24 hours
resource "spacelift_scheduled_delete_stack" "cleanup_dev" {
  stack_id   = spacelift_stack.development.id
  at         = timeadd(timestamp(), "24h")
  delete_resources = true
}

# Schedule stack deletion for end of sprint
resource "spacelift_scheduled_delete_stack" "sprint_cleanup" {
  stack_id   = spacelift_stack.feature_branch.id
  at         = "2024-12-31T23:59:59Z"
  delete_resources = false
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to schedule for deletion
* `at` - (Required) Timestamp when the stack should be deleted (RFC3339 format)

### Optional Arguments
* `delete_resources` - (Optional) Whether to run a destroy operation before deleting the stack. Defaults to `true`

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Stack deletion is irreversible once executed
* If `delete_resources` is true, a destroy run will be executed first
* The scheduled deletion can be canceled by removing this resource