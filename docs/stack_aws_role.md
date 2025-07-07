# Resource: spacelift_stack_aws_role

## Description
**DEPRECATED**: AWS IAM role linked to a specific stack. This resource is deprecated in favor of `spacelift_aws_integration` and `spacelift_aws_integration_attachment` which provide better functionality including shared integrations and separate read/write roles.

## Example Usage
```hcl
# Stack AWS role (deprecated)
resource "spacelift_stack_aws_role" "production" {
  stack_id = spacelift_stack.infrastructure.id
  role_arn = aws_iam_role.spacelift.arn
}

# Private worker configuration (deprecated)
resource "spacelift_stack_aws_role" "private_worker" {
  stack_id                       = spacelift_stack.app.id
  role_arn                       = "arn:aws:iam::123456789012:role/spacelift"
  generate_credentials_in_worker = true
  external_id                    = "custom-external-id"
}
```

## Argument Reference

### Required Arguments
* `stack_id` - (Required) ID of the stack to attach the role to
* `role_arn` - (Required) ARN of the AWS IAM role to assume

### Optional Arguments
* `generate_credentials_in_worker` - (Optional) Generate credentials in private worker. Defaults to `false`
* `duration_seconds` - (Optional) Role session duration in seconds
* `external_id` - (Optional) Custom external ID for private workers
* `region` - (Optional) AWS region for STS endpoint

### Read-Only Arguments
* `id` - Unique resource identifier

## Migration to spacelift_aws_integration
```hcl
# Old (deprecated)
resource "spacelift_stack_aws_role" "old" {
  stack_id = spacelift_stack.app.id
  role_arn = aws_iam_role.spacelift.arn
}

# New (recommended)
resource "spacelift_aws_integration" "new" {
  name     = "production-aws"
  role_arn = aws_iam_role.spacelift.arn
}

resource "spacelift_aws_integration_attachment" "new" {
  integration_id = spacelift_aws_integration.new.id
  stack_id       = spacelift_stack.app.id
}
```

## Notes
* **This resource is deprecated** - use `spacelift_aws_integration` instead
* For shared workers, external ID is automatically generated
* Migration provides better security and reusability