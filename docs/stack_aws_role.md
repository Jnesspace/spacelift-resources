spacelift_stack_aws_role (Resource)
Note:

spacelift_stack_aws_role is deprecated. Please use spacelift_aws_role instead. The functionality is identical.

NOTE: while this resource continues to work, we have replaced it with the spacelift_aws_integration resource. The new resource allows integrations to be shared by multiple stacks/modules and also supports separate read vs write roles. Please use the spacelift_aws_integration resource instead.

spacelift_stack_aws_role represents cross-account IAM role delegation between the Spacelift worker and an individual stack or module. If this is set, Spacelift will use AWS STS to assume the supplied IAM role and put its temporary credentials in the runtime environment.

If you use private workers, you can also assume IAM role on the worker side using your own AWS credentials (e.g. from EC2 instance profile).

Note: when assuming credentials for shared worker, Spacelift will use $accountName@$stackID or $accountName@$moduleID as external ID and $runID@$stackID@$accountName truncated to 64 characters as session ID.

METADATA:
  resource_type: spacelift_stack_aws_role
  provider: spacelift
  service: cloud_integration
  description: Legacy AWS IAM role delegation for stacks
  version: latest
  deprecation_notice: Use spacelift_aws_integration for improved functionality including shared integrations and separate read/write roles

USAGE_TEMPLATE:
```hcl
resource "spacelift_stack_aws_role" "RESOURCE_NAME" {
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
    - Deprecated resource type
    - One role per stack/module
    - No separate read/write roles
    - Use aws_integration instead

MIGRATION:
  target: spacelift_aws_integration
  benefits:
    - Shared across resources
    - Separate read/write roles
    - Improved management
    - Better security model

EXAMPLE_USAGE:
  shared_worker:
    purpose: Basic stack integration
    configuration:
      role_arn: AWS role ARN
      stack_id: Target stack
      
  private_worker:
    purpose: Custom credential generation
    configuration:
      role_arn: AWS role ARN
      stack_id: Target stack
      generate_credentials_in_worker: true