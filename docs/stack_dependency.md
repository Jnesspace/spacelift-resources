# Resource: spacelift_stack_dependency

## Description
Creates a dependency relationship between two stacks, ensuring that the dependent stack waits for the dependency stack to complete successfully before executing its own runs.

## Example Usage
```hcl
# Infrastructure dependency
resource "spacelift_stack_dependency" "app_depends_on_infra" {
  stack_id            = spacelift_stack.application.id
  depends_on_stack_id = spacelift_stack.infrastructure.id
}

# Multiple dependencies
resource "spacelift_stack_dependency" "frontend_depends_on_backend" {
  stack_id            = spacelift_stack.frontend.id
  depends_on_stack_id = spacelift_stack.backend.id
}

resource "spacelift_stack_dependency" "backend_depends_on_database" {
  stack_id            = spacelift_stack.backend.id
  depends_on_stack_id = spacelift_stack.database.id
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the dependent stack (the stack that waits)
* `depends_on_stack_id` - (Required) ID of the dependency stack (the stack that must complete first)

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Dependencies are immutable once created
* Circular dependencies are not allowed
* Changes to the dependency stack will trigger runs on the dependent stack
* The dependent stack will not start until the dependency stack finishes successfully