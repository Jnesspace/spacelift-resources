# Resource: spacelift_scheduled_task

## Description
Configures scheduled task execution for a stack, allowing automatic execution of custom commands at specified times or intervals.

## Example Usage
```hcl
# Daily cleanup task
resource "spacelift_scheduled_task" "cleanup" {
  stack_id = spacelift_stack.production.id
  command  = "./cleanup.sh"
  every    = ["0 3 * * *"]  # Every day at 3 AM
  timezone = "UTC"
}

# Workday deployment task
resource "spacelift_scheduled_task" "deploy" {
  stack_id = spacelift_stack.staging.id
  command  = "terraform apply -auto-approve"
  every    = ["0 9 * * 1-5"]  # Weekdays at 9 AM
  timezone = "America/New_York"
}

# One-time maintenance task
resource "spacelift_scheduled_task" "maintenance" {
  stack_id = spacelift_stack.app.id
  command  = "./maintenance.sh"
  at       = 1699747200  # Specific timestamp
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to run the task on
* `command` - (Required) Command to execute

### Optional Arguments
* `every` - (Optional) List of cron expressions for recurring tasks. Mutually exclusive with `at`
* `at` - (Optional) Unix timestamp for one-time task. Mutually exclusive with `every`
* `timezone` - (Optional) Timezone for schedule evaluation. Defaults to `"UTC"`
* `schedule_id` - (Optional) Custom schedule identifier

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_scheduled_task.example $STACK_ID/$SCHEDULED_TASK_ID
```

## Notes
* Either `every` or `at` must be specified, but not both
* Tasks run in the stack's environment with access to mounted files and variables
* Cron expressions follow standard cron syntax