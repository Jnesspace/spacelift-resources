METADATA:
  resource_type: spacelift_policy
  provider: spacelift
  service: governance
  description: Customer-defined rules for Spacelift decision points
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_policy" "RESOURCE_NAME" {
  name        = POLICY_NAME
  body        = file("${path.module}/policies/policy.rego")
  type        = POLICY_TYPE
  description = DESCRIPTION    # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Policy identifier
      validation: Must be unique in account
      
    body:
      type: String
      description: Rego policy content
      validation: Must be valid Rego syntax
      
    type:
      type: String
      description: Policy decision point
      allowed_values:
        - ACCESS         # Resource access control
        - APPROVAL      # Run approval rules
        - GIT_PUSH      # VCS push handling
        - INITIALIZATION # Stack setup
        - LOGIN         # Authentication rules
        - PLAN          # Terraform plan rules
        - TASK          # Task execution rules
        - TRIGGER       # Run trigger rules
        - NOTIFICATION  # Notification rules
      deprecated_values:
        - STACK_ACCESS  # Use ACCESS instead
        - TASK_RUN     # Use TASK instead
        - TERRAFORM_PLAN # Use PLAN instead

  optional:
    description:
      type: String
      description: Human-readable policy description
      
    labels:
      type: Set[String]
      description: Policy classification tags
      default: []
      
    space_id:
      type: String
      description: Target space identifier
      default: root or legacy space

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  execution:
    - Evaluated at specific decision points
    - Uses Rego policy language
    - Must return valid decision result
    
  application:
    - Requires explicit attachment to resources
    - Can be shared across multiple resources
    - Supports hierarchical organization
    
  scope:
    - Can control multiple resource types
    - Supports space-level isolation
    - Enables granular access control

  integrations:
    - Works with stacks and modules
    - Integrates with VCS operations
    - Controls automation workflows

IMPORT_FORMAT: $POLICY_ID