METADATA:
  resource_type: spacelift_scheduled_run
  provider: spacelift
  service: automation
  description: Scheduled run configuration for Spacelift stacks
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_scheduled_run" "RESOURCE_NAME" {
  stack_id = STACK_ID
  name     = SCHEDULE_NAME      # Optional
  
  # Either at or every must be specified:
  at       = TIMESTAMP         # Optional, unix timestamp
  every    = [CRON_EXPRS]     # Optional, list of cron expressions
  timezone = TIMEZONE         # Optional, defaults to UTC
  
  # Optional runtime configuration:
  runtime_config {
    project_root  = "path/to/root"
    runner_image  = "custom/image:tag"
    environment {
      key   = "ENV_VAR"
      value = "value"
    }
  }
}
```

ATTRIBUTES:
  required:
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist in Spacelift

  optional:
    name:
      type: String
      description: Schedule identifier
      validation: Must be unique per stack
      
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
      
    runtime_config:
      type: Block
      max: 1
      description: Run-specific configuration
      fields:
        project_root:
          type: String
          description: Stack entrypoint path
          
        runner_image:
          type: String
          description: Custom Docker image
          
        environment:
          type: Block Set
          description: Run environment variables
          fields:
            key:
              type: String
              description: Variable name
              required: true
            value:
              type: String
              description: Variable value
              required: true
              
        hooks:
          description: Run lifecycle hooks
          types:
            - before_init
            - after_init
            - before_plan
            - after_plan
            - before_apply
            - after_apply
            - before_destroy
            - after_destroy
            - before_perform
            - after_perform
            - after_run

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    next_schedule:
      type: Number
      description: Next run timestamp
      generated: true
      
    terraform_version:
      type: String
      description: Terraform version
      generated: true
      
    terraform_workflow_tool:
      type: String
      description: IaC execution tool
      generated: true
      allowed_values: ["OPEN_TOFU", "TERRAFORM_FOSS", "CUSTOM"]
      default: "TERRAFORM_FOSS"

BEHAVIOR:
  scheduling:
    - Supports one-time execution via timestamp
    - Supports recurring execution via cron
    - Timezone-aware scheduling
    
  runtime:
    - Customizable environment variables
    - Custom Docker image support
    - Extensive hook system
    - Project root configuration
    
  execution:
    - Runs in stack's environment
    - Uses configured worker pool
    - Inherits stack permissions
    
  examples:
    workday_apply:
      every: ["0 7 * * 1-5"]
      timezone: "CET"
      name: "apply-workdays"
      
    one_time_run:
      at: "1663336895"
      name: "one-off-apply"
      
    custom_runtime:
      every: ["0 21 * * 1-5"]
      runtime_config:
        terraform_version: "1.5.7"

IMPORT_FORMAT: $STACK_ID/$SCHEDULED_RUN_ID