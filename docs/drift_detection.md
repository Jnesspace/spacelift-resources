METADATA:
  resource_type: spacelift_drift_detection
  provider: spacelift
  service: monitoring
  description: Infrastructure drift detection and reconciliation
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_drift_detection" "RESOURCE_NAME" {
  stack_id  = STACK_ID
  schedule  = [CRON_EXPRS]
  reconcile = true          # Optional
  timezone  = TIMEZONE     # Optional, defaults to UTC
}
```

ATTRIBUTES:
  required:
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist in Spacelift
      
    schedule:
      type: List[String]
      description: Cron schedule expressions
      validation: Must be valid cron syntax
      example: ["*/15 * * * *"]

  optional:
    reconcile:
      type: Boolean
      description: Auto-trigger tracked run on drift
      default: false
      
    timezone:
      type: String
      description: Schedule timezone
      default: "UTC"
      validation: Must be valid timezone name
      
    ignore_state:
      type: Boolean
      description: Check drift in any final state
      default: false
      note: If false, only checks 'Finished' state

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  detection:
    - Runs proposed run on schedule
    - Checks infrastructure state
    - Reports detected drift
    - Can monitor any final state
    
  reconciliation:
    - Optional automatic fixes
    - Triggers tracked run on drift
    - Uses stack's configuration
    - Inherits stack permissions
    
  monitoring:
    - Regular schedule checks
    - Timezone-aware scheduling
    - Integrates with webhooks
    - Reports drift status
    
  examples:
    frequent_check:
      schedule: ["*/15 * * * *"]
      reconcile: true
      description: "Check every 15 minutes, auto-fix"
      
    daily_monitor:
      schedule: ["0 0 * * *"]
      reconcile: false
      description: "Daily check, manual fixes"

IMPORT_FORMAT: stack/$STACK_ID