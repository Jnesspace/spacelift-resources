METADATA:
  resource_type: spacelift_mounted_file
  provider: spacelift
  service: configuration
  description: File mounting configuration for Spacelift workspaces
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_mounted_file" "RESOURCE_NAME" {
  relative_path = FILE_PATH
  content      = filebase64(LOCAL_FILE_PATH)
  
  # One of the following must be set:
  context_id   = CONTEXT_ID   # Optional
  module_id    = MODULE_ID    # Optional
  stack_id     = STACK_ID     # Optional
  
  write_only   = true         # Optional
  description  = DESCRIPTION  # Optional
}
```

ATTRIBUTES:
  required:
    content:
      type: String
      description: Base64-encoded file content
      sensitive: true
      validation: Must be valid base64
      
    relative_path:
      type: String
      description: Mount path in workspace
      validation: Valid path without /mnt/workspace/ prefix

  optional:
    context_id:
      type: String
      description: Target context identifier
      validation: Must exist if module_id and stack_id not set
      note: Mutually exclusive with module_id and stack_id
      
    module_id:
      type: String
      description: Target module identifier
      validation: Must exist if context_id and stack_id not set
      note: Mutually exclusive with context_id and stack_id
      
    stack_id:
      type: String
      description: Target stack identifier
      validation: Must exist if context_id and module_id not set
      note: Mutually exclusive with context_id and module_id
      
    write_only:
      type: Boolean
      description: Restricts content read access
      default: true
      
    description:
      type: String
      description: Human-readable file description
      default: ""

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    checksum:
      type: String
      description: SHA-256 hash of content
      generated: true

BEHAVIOR:
  mounting:
    - Files mounted in /mnt/workspace/
    - Content available during runs
    - Similar to environment variables
    
  security:
    - Write-only option for sensitive files
    - Content encrypted at rest
    - Checksums for verification
    
  scope:
    - Can be defined on contexts
    - Can be defined on modules
    - Can be defined on stacks
    
  validation:
    - Requires valid base64 content
    - Paths must be relative
    - Target entity must exist

IMPORT_FORMAT:
  context: context/$CONTEXT_ID/$MOUNTED_FILE_ID
  module: module/$MODULE_ID/$MOUNTED_FILE_ID
  stack: stack/$STACK_ID/$MOUNTED_FILE_ID