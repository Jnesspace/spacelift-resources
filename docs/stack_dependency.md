METADATA:
  resource_type: spacelift_stack_dependency
  provider: spacelift
  service: orchestration
  description: Stack dependency management for execution ordering
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_stack_dependency" "RESOURCE_NAME" {
  stack_id            = DEPENDENT_STACK_ID
  depends_on_stack_id = DEPENDENCY_STACK_ID
}
```

ATTRIBUTES:
  required:
    stack_id:
      type: String
      description: Dependent stack identifier
      validation: Must exist in Spacelift
      immutable: true
      
    depends_on_stack_id:
      type: String
      description: Dependency stack identifier
      validation: Must exist in Spacelift
      immutable: true

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  execution:
    - Blocks dependent stack runs
    - Waits for dependency completion
    - Requires successful dependency state
    
  triggering:
    - Changes in dependency trigger dependent
    - Ensures infrastructure consistency
    - Maintains deployment order
    
  validation:
    - Both stacks must exist
    - Cannot create circular dependencies
    - Relationship is immutable
    
  states:
    dependent:
      - Waits for dependency
      - Triggered by changes
      - Requires dependency success
    dependency:
      - Must finish successfully
      - Changes affect dependents
      - Status affects dependent runs

PATTERNS:
  infrastructure:
    example:
      - Base infrastructure stack
      - Application stack depends on base
    benefit: Ensures proper deployment order
    
  microservices:
    example:
      - Shared services stack
      - Individual service stacks
    benefit: Maintains service dependencies