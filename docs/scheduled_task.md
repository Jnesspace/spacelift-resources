spacelift_scheduled_task (Resource)

spacelift_scheduled_task represents a scheduling configuration for a Stack. It will trigger task on the given schedule or timestamp

Example Usage

resource "spacelift_stack" "k8s-core" {
  // ...
}

// create the resources of a stack on a given schedule
resource "spacelift_scheduled_task" "k8s-core-create" {
  stack_id = spacelift_stack.k8s-core.id

  command  = "terraform apply -auto-approve"
  every    = ["0 7 * * 1-5"]
  timezone = "CET"
}

// destroy the resources of a stack on a given schedule
resource "spacelift_scheduled_task" "k8s-core-destroy" {
  stack_id = spacelift_stack.k8s-core.id

  command  = "terraform destroy -auto-approve"
  every    = ["0 21 * * 1-5"]
  timezone = "CET"
}

// at a given timestamp (unix)
resource "spacelift_scheduled_task" "k8s-core-destroy" {
  stack_id = spacelift_stack.k8s-core.id

  command = "terraform destroy -auto-approve"
  at      = "1663336895"
}

Schema
Required

    command (String) Command that will be run.
    stack_id (String) ID of the stack for which to set up the scheduled task

Optional

    at (Number) Timestamp (unix timestamp) at which time the scheduled task should happen.
    every (List of String) List of cron schedule expressions based on which the scheduled task should be triggered.
    schedule_id (String) ID of the schedule
    timezone (String) Timezone in which the schedule is expressed. Defaults to UTC.

Read-Only

    id (String) The ID of this resource.

Import

Import is supported using the following syntax:

terraform import spacelift_scheduled_task.ireland-kubeconfig $STACK_ID/$SCHEDULED_TASK_ID

On this page

    Example Usage
    Schema
    Import

Report an issue

METADATA:
  resource_type: spacelift_scheduled_task
  provider: spacelift
  service: automation
  description: Scheduled task execution configuration for Spacelift stacks
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_scheduled_task" "RESOURCE_NAME" {
  stack_id = STACK_ID
  command  = COMMAND
  
  # Either at or every must be specified:
  at       = TIMESTAMP      # Optional, unix timestamp
  every    = [CRON_EXPRS]  # Optional, list of cron expressions
  timezone = TIMEZONE      # Optional, defaults to UTC
}
```

ATTRIBUTES:
  required:
    command:
      type: String
      description: Command to execute
      validation: Must be valid shell command
      
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist in Spacelift

  optional:
    at:
      type: Number
      description: Unix timestamp for one-time execution
      validation: Must be future timestamp
      note: Mutually exclusive with every
      
    every:
      type: List[String]
      description: Cron schedule expressions
      validation: Must be valid cron syntax
      note: Mutually exclusive with at
      example: ["0 7 * * 1-5"]
      
    timezone:
      type: String
      description: Schedule timezone
      default: "UTC"
      validation: Must be valid timezone name
      
    schedule_id:
      type: String
      description: Custom schedule identifier
      validation: Must be unique per stack

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  scheduling:
    - Supports one-time execution via timestamp
    - Supports recurring execution via cron
    - Timezone-aware scheduling
    
  execution:
    - Runs in stack's environment
    - Uses stack's worker pool
    - Inherits stack permissions
    
  validation:
    - Must specify either at or every
    - Cannot use both at and every
    - Cron expressions must be valid
    
  examples:
    daily_backup:
      every: ["0 0 * * *"]
      command: "terraform apply -auto-approve"
      
    workday_deploy:
      every: ["0 7 * * 1-5"]
      timezone: "CET"
      command: "terraform apply -auto-approve"
      
    scheduled_cleanup:
      at: "1663336895"
      command: "terraform destroy -auto-approve"

IMPORT_FORMAT: $STACK_ID/$SCHEDULED_TASK_ID