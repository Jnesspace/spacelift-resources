# Resource: spacelift_run

## Description
Triggers a run on a Spacelift stack programmatically. Runs execute the stack's infrastructure code and can be configured to wait for completion.

## Example Usage
```hcl
# Basic run trigger
resource "spacelift_run" "deploy" {
  stack_id = spacelift_stack.app.id
  
  keepers = {
    version = var.app_version
  }
}

# Run with specific commit and wait configuration
resource "spacelift_run" "rollback" {
  stack_id   = spacelift_stack.app.id
  commit_sha = "abc123def456"
  proposed   = false
  
  keepers = {
    rollback_trigger = timestamp()
  }
  
  wait {
    continue_on_state   = ["finished", "failed"]
    continue_on_timeout = false
    disabled           = false
  }
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to trigger a run on

### Optional Arguments
* `keepers` - (Optional) Map of values that trigger resource recreation when changed
* `commit_sha` - (Optional) Specific commit SHA to run. If not specified, uses the latest commit
* `proposed` - (Optional) Whether to create a proposed run. Defaults to `false`
* `wait` - (Optional) Configuration for waiting for run completion
  * `continue_on_state` - (Optional) States that allow continuation. Defaults to `["finished"]`
  * `continue_on_timeout` - (Optional) Continue if run times out. Defaults to `false`
  * `disabled` - (Optional) Disable waiting for completion. Defaults to `false`
* `timeouts` - (Optional) Resource operation timeouts
  * `create` - (Optional) Run creation timeout

### Read-Only Arguments
* `id` - Unique resource identifier (the triggered run ID)

## Import
```bash
terraform import spacelift_run.example $RUN_ID
```

## Notes
* Runs are triggered when keepers change or when the resource is created
* Proposed runs don't make actual infrastructure changes
* Use wait configuration to control Terraform execution flow based on run completion