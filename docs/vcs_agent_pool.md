METADATA:
  resource_type: spacelift_vcs_agent_pool
  provider: spacelift
  service: vcs_integration
  description: Agent pool management for private VCS access
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_vcs_agent_pool" "RESOURCE_NAME" {
  name        = POOL_NAME
  description = DESCRIPTION    # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Pool identifier
      validation: Must be unique in account

  optional:
    description:
      type: String
      description: Human-readable pool description
      default: ""

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    config:
      type: String
      description: Base64 encoded pool configuration
      sensitive: true
      generated: true

BEHAVIOR:
  access:
    - Enables private VCS access
    - Provides proxy functionality
    - Manages agent groups
    
  security:
    - Secure configuration delivery
    - Protected agent credentials
    - Encrypted communication
    
  management:
    - Logical agent grouping
    - Pool-level configuration
    - Agent coordination
    
  usage:
    enterprise:
      - GitHub Enterprise access
      - GitLab self-hosted
      - Bitbucket Server/DC
    private:
      - Internal Git servers
      - Custom VCS installations
      - Protected repositories

IMPORT_FORMAT: $VCS_AGENT_POOL_ID