# Resource: spacelift_module

## Description
A module is a special type of stack used for testing and versioning Terraform modules. Modules can be shared across multiple stacks and provide reusable infrastructure components.

## Example Usage
```hcl
# Explicit module configuration
resource "spacelift_module" "vpc" {
  name               = "vpc-module"
  terraform_provider = "aws"
  repository         = "terraform-aws-vpc"
  branch             = "main"
  description        = "Reusable VPC module for AWS"
  administrative     = false
}

# Module with auto-inferred name and provider
resource "spacelift_module" "auto_inferred" {
  repository   = "terraform-aws-eks"  # Name and provider inferred
  branch       = "main"
  project_root = "modules/cluster"
  description  = "EKS cluster module"
  public       = true
}
```

## Argument Reference

### Required Arguments
* `repository` - (Required) Repository name without the owner part
* `branch` - (Required) Git branch to apply changes to

### Optional Arguments
* `name` - (Optional) Module name. Auto-inferred from repository if following terraform-{provider}-{name} convention
* `terraform_provider` - (Optional) Terraform provider. Auto-inferred from repository if following naming convention
* `description` - (Optional) Human-readable description of the module
* `project_root` - (Optional) Directory containing the module source code
* `administrative` - (Optional) Whether the module can manage other resources. Defaults to `false`
* `public` - (Optional) Make module publicly accessible. Can only be set at creation. Defaults to `false`
* `protect_from_deletion` - (Optional) Prevent accidental deletion. Defaults to `false`
* `shared_accounts` - (Optional) List of accounts that can access the module
* `space_id` - (Optional) ID of the space the module belongs to
* `labels` - (Optional) Set of labels for module classification
* `worker_pool_id` - (Optional) ID of the worker pool to use for module runs

### Read-Only Arguments
* `id` - Unique resource identifier
* `aws_assume_role_policy_statement` - AWS IAM assume role policy for trust setup

## Import
```bash
terraform import spacelift_module.example $MODULE_ID
```

## Notes
* Repository naming convention terraform-{provider}-{name} enables auto-inference
* Public visibility can only be set during creation
* Modules support version tags for consumption by stacks