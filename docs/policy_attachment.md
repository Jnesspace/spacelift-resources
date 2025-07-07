# Resource: spacelift_policy_attachment

## Description
Attaches a policy to a stack or module, enabling the policy's rules to be evaluated during the target resource's operations.

## Example Usage
```hcl
# Attach plan policy to stack
resource "spacelift_policy_attachment" "no_weekend_deploys" {
  policy_id = spacelift_policy.weekend_block.id
  stack_id  = spacelift_stack.production.id
}

# Attach access policy to module
resource "spacelift_policy_attachment" "module_access" {
  policy_id = spacelift_policy.team_access.id
  module_id = spacelift_module.vpc.id
}
```

## Argument Reference

### Required Arguments
* `policy_id` - (Required) ID of the policy to attach

### Optional Arguments
* `stack_id` - (Optional) ID of the stack to attach the policy to. Mutually exclusive with `module_id`
* `module_id` - (Optional) ID of the module to attach the policy to. Mutually exclusive with `stack_id`

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_policy_attachment.example $POLICY_ID/$STACK_ID
```

## Notes
* Exactly one of `stack_id` or `module_id` must be specified
* LOGIN policies cannot be attached as they apply globally
* Each policy can only be attached once per stack/module