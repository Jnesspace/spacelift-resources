# Resource: spacelift_drift_detection

## Description
Configures automatic drift detection for a stack, enabling periodic checks to identify when the actual infrastructure state differs from the Terraform state.

## Example Usage
```hcl
# Daily drift detection
resource "spacelift_drift_detection" "production" {
  stack_id  = spacelift_stack.production.id
  reconcile = true
  schedule  = ["0 9 * * *"]  # Every day at 9 AM
  timezone  = "UTC"
}

# Workday drift detection with reconciliation disabled
resource "spacelift_drift_detection" "staging" {
  stack_id  = spacelift_stack.staging.id
  reconcile = false
  schedule  = ["0 9 * * 1-5"]  # Weekdays at 9 AM
  timezone  = "America/New_York"
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to configure drift detection for
* `reconcile` - (Required) Whether to automatically reconcile detected drift
* `schedule` - (Required) List of cron expressions for drift detection runs

### Optional Arguments
* `timezone` - (Optional) Timezone for schedule evaluation. Defaults to `"UTC"`

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Drift detection runs a plan to identify differences between desired and actual state
* When `reconcile` is true, detected drift triggers an automatic apply
* Schedule follows standard cron syntax for flexible timing configuration