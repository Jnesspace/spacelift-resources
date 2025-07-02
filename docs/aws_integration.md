# spacelift_aws_integration

METADATA:
  resource_type: spacelift_aws_integration
  provider: spacelift
  service: cloud_integration
  description: AWS account integration for Spacelift resource management
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_aws_integration" "RESOURCE_NAME" {
  name = INTEGRATION_NAME
  role_arn = IAM_ROLE_ARN
  generate_credentials_in_worker = BOOLEAN
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Friendly identifier for the integration
      validation: none
    
    role_arn:
      type: String
      description: AWS IAM role ARN to be assumed
      validation: Must be valid ARN format

  optional:
    duration_seconds:
      type: Number
      description: Validity period for assumed credentials
      default: 900
      validation: 900-3600 seconds
    
    external_id:
      type: String
      description: Custom external ID for role assumption
      applies_to: Private workers only
      validation: none
    
    generate_credentials_in_worker:
      type: Boolean
      description: Enable AWS credential generation in private worker
      default: false
      validation: none
    
    labels:
      type: Set[String]
      description: Resource classification tags
      default: []
      validation: none
    
    region:
      type: String
      description: AWS region for STS endpoint selection
      validation: Valid AWS region
    
    space_id:
      type: String
      description: Target space identifier
      validation: Must exist in Spacelift

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  session_naming:
    format: "$runID@$stackID@$accountName"
    max_length: 64

  external_id_format:
    stack: "$accountName@$integrationID@$stackID@$suffix"
    module: "$accountName@$integrationID@$moduleID@$suffix"
    suffix_values: ["read", "write"]

  requirements:
    - Explicit stack attachment required after creation
    - IAM role must trust Spacelift's AWS account
    - Role must have necessary permissions for intended operations