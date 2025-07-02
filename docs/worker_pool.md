METADATA:
  resource_type: spacelift_worker_pool
  provider: spacelift
  service: execution
  description: Worker pool for executing Spacelift jobs
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_worker_pool" "RESOURCE_NAME" {
  name        = POOL_NAME
  csr         = filebase64(CSR_PATH)    # Optional
  description = DESCRIPTION             # Optional
  labels      = [LABELS]               # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Worker pool identifier
      validation: Must be unique in account

  optional:
    csr:
      type: String
      description: Base64 encoded certificate signing request
      sensitive: true
      note: Changes trigger token reset
      
    description:
      type: String
      description: Human-readable pool description
      
    labels:
      type: Set[String]
      description: Pool classification tags
      default: []
      
    space_id:
      type: String
      description: Target space identifier
      default: root or legacy space

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    config:
      type: String
      description: Worker connection credentials
      sensitive: true
      generated: true
      
    private_key:
      type: String
      description: Base64 encoded private key
      sensitive: true
      generated: true

BEHAVIOR:
  security:
    - Uses certificate-based authentication
    - Generates secure connection credentials
    - Private key never transmitted after creation
    
  configuration:
    - Can be assigned to specific spaces
    - Supports labeling for organization
    - CSR changes trigger security resets
    
  execution:
    - Provides isolated execution environment
    - Can be shared across stacks/modules
    - Supports custom worker configurations

IMPORT_FORMAT: $WORKER_POOL_ID