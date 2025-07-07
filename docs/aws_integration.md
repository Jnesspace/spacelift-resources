# Resource: spacelift_aws_integration

## Description
Configures AWS integration for secure, credential-less access to AWS resources. AWS integrations can be shared across multiple stacks and modules, with separate read and write IAM roles for enhanced security.

## Example Usage
```hcl
# Basic AWS integration
resource "spacelift_aws_integration" "production" {
  name                           = "production-aws"
  role_arn                       = aws_iam_role.spacelift_production.arn
  generate_credentials_in_worker = false
  space_id                       = spacelift_space.production.id
}

# AWS integration with separate read/write roles
resource "spacelift_aws_integration" "secure" {
  name                           = "secure-aws"
  role_arn                       = aws_iam_role.spacelift_write.arn
  read_role_arn                  = aws_iam_role.spacelift_read.arn
  generate_credentials_in_worker = false
  external_id                    = "custom-external-id"
  duration_seconds               = 3600
}

# Private worker integration
resource "spacelift_aws_integration" "private" {
  name                           = "private-worker-aws"
  role_arn                       = aws_iam_role.spacelift_private.arn
  generate_credentials_in_worker = true
  external_id                    = var.private_worker_external_id
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique integration identifier
* `role_arn` - (Required) ARN of the AWS IAM role for write operations

### Optional Arguments
* `read_role_arn` - (Optional) ARN of the AWS IAM role for read operations. If not specified, `role_arn` is used for both
* `generate_credentials_in_worker` - (Optional) Generate credentials in private worker. Defaults to `false`
* `external_id` - (Optional) Custom external ID for enhanced security. Auto-generated if not specified
* `duration_seconds` - (Optional) Session duration in seconds. Defaults to 3600
* `space_id` - (Optional) ID of the space the integration belongs to. Defaults to `"root"`
* `labels` - (Optional) Set of labels for integration classification

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_aws_integration.example $INTEGRATION_ID
```

## Notes
* Integrations can be attached to multiple stacks/modules using `spacelift_aws_integration_attachment`
* Separate read/write roles enable principle of least privilege
* External ID enhances security for cross-account access