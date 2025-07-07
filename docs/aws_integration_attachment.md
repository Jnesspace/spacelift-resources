# Resource: spacelift_aws_integration_attachment

## Description
Attaches an AWS integration to a stack or module, enabling AWS credential management with separate read and write permissions.

## Example Usage
```hcl
# Full access attachment
resource "spacelift_aws_integration_attachment" "production" {
  integration_id = spacelift_aws_integration.production.id
  stack_id       = spacelift_stack.infrastructure.id
  read           = true
  write          = true
  
  depends_on = [aws_iam_role.spacelift_role]
}

# Read-only attachment
resource "spacelift_aws_integration_attachment" "monitoring" {
  integration_id = spacelift_aws_integration.production.id
  stack_id       = spacelift_stack.monitoring.id
  read           = true
  write          = false
}

# Module attachment
resource "spacelift_aws_integration_attachment" "vpc_module" {
  integration_id = spacelift_aws_integration.shared.id
  module_id      = spacelift_module.vpc.id
  read           = true
  write          = true
}
```

## Argument Reference

### Required Arguments
* `integration_id` - (Required) ID of the AWS integration to attach

### Optional Arguments
* `stack_id` - (Optional) ID of the stack to attach to. Mutually exclusive with `module_id`
* `module_id` - (Optional) ID of the module to attach to. Mutually exclusive with `stack_id`
* `read` - (Optional) Enable read operations. Defaults to `true`
* `write` - (Optional) Enable write operations. Defaults to `true`

### Read-Only Arguments
* `id` - Unique resource identifier
* `attachment_id` - Internal attachment identifier

## Import
```bash
terraform import spacelift_aws_integration_attachment.example $INTEGRATION_ID/$STACK_ID
```

## Notes
* Either `stack_id` or `module_id` must be specified
* IAM role must exist before attachment (use `depends_on`)
* Role assumption is tested during attachment creation

