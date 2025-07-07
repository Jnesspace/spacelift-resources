# Resource: spacelift_scheduled_run

## Description
Configures scheduled runs for a stack, allowing automatic execution of infrastructure code at specified times or intervals.

## Example Usage
```hcl
# Daily scheduled run
resource "spacelift_scheduled_run" "daily_apply" {
  stack_id = spacelift_stack.production.id
  name     = "daily-production-sync"
  every    = ["0 2 * * *"]  # Every day at 2 AM
  timezone = "UTC"
}

# Workday deployment schedule
resource "spacelift_scheduled_run" "workday_deploy" {
  stack_id = spacelift_stack.staging.id
  name     = "workday-deployment"
  every    = ["0 9 * * 1-5"]  # Weekdays at 9 AM
  timezone = "America/New_York"
  
  runtime_config {
    project_root = "environments/staging"
    environment {
      key   = "DEPLOY_MODE"
      value = "automated"
    }
  }
}

# One-time scheduled run
resource "spacelift_scheduled_run" "maintenance" {
  stack_id = spacelift_stack.maintenance.id
  name     = "maintenance-window"
  at       = 1699747200  # Specific timestamp
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to schedule runs for

### Optional Arguments
* `name` - (Optional) Human-readable name for the scheduled run
* `every` - (Optional) List of cron expressions for recurring runs. Mutually exclusive with `at`
* `at` - (Optional) Unix timestamp for one-time run. Mutually exclusive with `every`
* `timezone` - (Optional) Timezone for schedule evaluation. Defaults to `"UTC"`
* `schedule_id` - (Optional) Custom schedule identifier
* `runtime_config` - (Optional) Custom runtime configuration for the scheduled runs
  * `project_root` - (Optional) Project root directory override
  * `runner_image` - (Optional) Custom Docker image for the run
  * `environment` - (Optional) Environment variables for the run
    * `key` - (Required) Environment variable name
    * `value` - (Required) Environment variable value

### Read-Only Arguments
* `id` - Unique resource identifier
* `next_schedule` - Unix timestamp of the next scheduled run

## Import
```bash
terraform import spacelift_scheduled_run.example $STACK_ID/$SCHEDULED_RUN_ID
```

## Notes
* Either `every` or `at` must be specified, but not both
* Cron expressions follow standard cron syntax
* Runtime configuration allows customization per scheduled run