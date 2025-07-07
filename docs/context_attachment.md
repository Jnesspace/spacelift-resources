# Resource: spacelift_context_attachment

## Description
Attaches a context to a stack or module with a specified priority, allowing the target resource to inherit the context's configuration (environment variables, mounted files, and hooks).

## Example Usage
```hcl
# Attach context to stack
resource "spacelift_context_attachment" "production_stack" {
  context_id = spacelift_context.production.id
  stack_id   = spacelift_stack.api.id
  priority   = 0
}

# Attach context to module with higher priority
resource "spacelift_context_attachment" "shared_config" {
  context_id = spacelift_context.shared.id
  module_id  = spacelift_module.vpc.id
  priority   = 10
}
```

## Argument Reference

### Required Arguments
* `context_id` - (Required) ID of the context to attach

### Optional Arguments
* `stack_id` - (Optional) ID of the stack to attach the context to. Mutually exclusive with `module_id`
* `module_id` - (Optional) ID of the module to attach the context to. Mutually exclusive with `stack_id`
* `priority` - (Optional) Priority of the attachment (lower numbers processed first). Defaults to `0`

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_context_attachment.example $CONTEXT_ID/$STACK_ID
```

## Notes
* Exactly one of `stack_id` or `module_id` must be specified
* Lower priority numbers are processed first in case of conflicts
* Multiple contexts can be attached to the same stack/module with different priorities