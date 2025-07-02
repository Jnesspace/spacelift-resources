spacelift_aws_role (Resource)

NOTE: while this resource continues to work, we have replaced it with the spacelift_aws_integration resource. The new resource allows integrations to be shared by multiple stacks/modules and also supports separate read vs write roles. Please use the spacelift_aws_integration resource instead.

spacelift_aws_role represents cross-account IAM role delegation between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.

If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).

Note: when assuming credentials for shared worker, Spacelift will use $accountName@$stackID or $accountName@$moduleID as external ID and $runID@$stackID@$accountName truncated to 64 characters as session ID.
Example Usage

# Assuming the role in Spacelift
resource "aws_iam_role" "spacelift" {
  name = "spacelift"

  assume_role_policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [jsondecode(spacelift_stack.k8s-core.aws_assume_role_policy_statement)]
  })
}

resource "aws_iam_role_policy_attachment" "spacelift" {
  role       = aws_iam_role.spacelift.name
  policy_arn = "arn:aws:iam::aws:policy/PowerUserAccess"
}

# Role assumed by a Stack
resource "spacelift_aws_role" "spacelift-stack" {
  stack_id = "k8s-core"
  role_arn = aws_iam_role.spacelift.arn
}

# Role assumed by a Module
resource "spacelift_aws_role" "spacelift-module" {
  module_id = "k8s-core"
  role_arn  = aws_iam_role.spacelift.arn
}

# Assuming the role in the private worker, for a stack.
resource "spacelift_aws_role" "k8s-core" {
  stack_id                       = "k8s-core"
  role_arn                       = "arn:aws:iam::123456789012:custom/role"
  generate_credentials_in_worker = true
}

# Assuming the role in the private worker, for a module.
resource "spacelift_aws_role" "k8s-core" {
  module_id                      = "k8s-core"
  role_arn                       = "arn:aws:iam::123456789012:custom/role"
  generate_credentials_in_worker = true
}

Schema
Required

    role_arn (String) ARN of the AWS IAM role to attach

Optional

    duration_seconds (Number) AWS IAM role session duration in seconds
    external_id (String) Custom external ID (works only for private workers).
    generate_credentials_in_worker (Boolean) Generate AWS credentials in the private worker. Defaults to false.
    module_id (String) ID of the module which assumes the AWS IAM role
    region (String) AWS region to select a regional AWS STS endpoint.
    stack_id (String) ID of the stack which assumes the AWS IAM role

Read-Only

    id (String) The ID of this resource.

Import

Import is supported using the following syntax:

terraform import spacelift_aws_role.k8s-core stack/$STACK_ID

terraform import spacelift_aws_role.k8s-core module/$MODULE_ID

On this page

    Example Usage
    Schema
    Import

Report an issue

METADATA:
  resource_type: spacelift_aws_role
  provider: spacelift
  service: cloud_integration
  description: AWS IAM role delegation for Spacelift workers
  version: latest
  deprecation_notice: Use spacelift_aws_integration instead for improved functionality

USAGE_TEMPLATE:
```hcl
resource "spacelift_aws_role" "RESOURCE_NAME" {
  role_arn = IAM_ROLE_ARN
  # One of the following must be set:
  stack_id  = STACK_ID    # Optional if module_id is set
  module_id = MODULE_ID   # Optional if stack_id is set
  
  # Optional configuration:
  generate_credentials_in_worker = false  # For private workers
  duration_seconds = 3600               # Session duration
  external_id = "custom-id"            # For private workers
  region = "us-west-2"                # AWS region for STS
}
```

ATTRIBUTES:
  required:
    role_arn:
      type: String
      description: AWS IAM role ARN to assume
      validation: Must be valid ARN format

  optional:
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist if module_id not set
      note: Mutually exclusive with module_id
      
    module_id:
      type: String
      description: Target module identifier
      validation: Must exist if stack_id not set
      note: Mutually exclusive with stack_id
      
    generate_credentials_in_worker:
      type: Boolean
      description: Generate credentials in private worker
      default: false
      applies_to: Private workers only
      
    duration_seconds:
      type: Number
      description: Role session duration
      validation: Valid session duration
      
    external_id:
      type: String
      description: Custom external ID
      applies_to: Private workers only
      
    region:
      type: String
      description: AWS region for STS endpoint
      validation: Valid AWS region name

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  session:
    shared_workers:
      external_id: "$accountName@$stackID" or "$accountName@$moduleID"
      session_id: "$runID@$stackID@$accountName" (max 64 chars)
    
    private_workers:
      - Can use custom external ID
      - Uses worker's AWS credentials
      - Supports regional endpoints
    
  credentials:
    - Temporary STS credentials
    - Configurable session duration
    - Available in runtime environment
    
  limitations:
    - Deprecated in favor of aws_integration
    - One role per stack/module
    - No separate read/write roles

IMPORT_FORMAT:
  stack: stack/$STACK_ID
  module: module/$MODULE_ID