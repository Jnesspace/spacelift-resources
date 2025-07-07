# Resource: spacelift_task

## Description
A task represents a command execution on a Spacelift stack. Tasks allow you to run arbitrary commands within the stack's environment and configuration.

## Example Usage
```hcl
# Basic task
resource "spacelift_task" "deploy" {
  stack_id = spacelift_stack.app.id
  command  = "terraform apply -auto-approve"
  init     = true
}

# Task with custom configuration
resource "spacelift_task" "cleanup" {
  stack_id = spacelift_stack.app.id
  command  = "./cleanup.sh"
  init     = false
  
  keepers = {
    trigger = timestamp()
  }
  
  wait {
    continue_on_state = ["finished", "failed"]
    disabled          = false
  }
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to run the task on
* `command` - (Required) Command to execute

### Optional Arguments
* `init` - (Optional) Whether to initialize the stack before running. Defaults to `true`
* `keepers` - (Optional) Map of values that trigger resource recreation when changed
* `wait` - (Optional) Configuration for waiting for task completion
  * `continue_on_state` - (Optional) States that allow continuation. Defaults to `["finished"]`
  * `continue_on_timeout` - (Optional) Continue if task times out. Defaults to `false`
  * `disabled` - (Optional) Disable waiting for completion. Defaults to `false`
* `timeouts` - (Optional) Resource operation timeouts
  * `create` - (Optional) Task creation timeout

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_task.example $STACK_ID/$TASK_ID
```

## Notes
* Tasks run in the stack's configured environment with access to mounted files and environment variables
* Use keepers to trigger task re-execution when dependencies change
* Tasks can be configured to wait for specific completion states