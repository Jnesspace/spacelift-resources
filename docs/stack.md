# Resource: spacelift_stack

## Description
A stack is a combination of source code and configuration that allows Spacelift to manage a piece of infrastructure. Stacks are the primary execution unit in Spacelift, similar to a CloudFormation stack or a CI/CD project.

## Example Usage
```hcl
# Basic Terraform stack
resource "spacelift_stack" "example" {
  name              = "my-infrastructure"
  repository        = "terraform-infrastructure"
  branch            = "main"
  project_root      = "environments/production"
  description       = "Production infrastructure stack"
  terraform_version = "1.5.0"
  autodeploy        = true
}

# Administrative stack
resource "spacelift_stack" "admin" {
  name           = "spacelift-admin"
  repository     = "spacelift-config"
  branch         = "main"
  administrative = true
  description    = "Manages other Spacelift resources"
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique stack identifier within the account
* `repository` - (Required) Repository name without the owner part
* `branch` - (Required) Git branch to apply changes to

### Optional Arguments
* `administrative` - (Optional) Whether the stack can manage other stacks. Defaults to `false`
* `autodeploy` - (Optional) Enable automatic deployment of changes. Defaults to `false`
* `description` - (Optional) Human-readable description of the stack
* `project_root` - (Optional) Directory containing the stack's entrypoint. Defaults to repository root
* `terraform_version` - (Optional) Terraform version to use
* `labels` - (Optional) Set of labels for stack classification
* `space_id` - (Optional) ID of the space the stack belongs to. Defaults to root space
* `worker_pool_id` - (Optional) ID of the worker pool to use for runs

### Read-Only Arguments
* `id` - Unique resource identifier
* `aws_assume_role_policy_statement` - AWS IAM assume role policy for trust setup

## Import
```bash
terraform import spacelift_stack.example $STACK_ID
```

## Notes
* Stack names must be unique within the account
* Administrative stacks have elevated privileges and can manage other Spacelift resources
* Autodeploy requires proper permissions and policies to be configured