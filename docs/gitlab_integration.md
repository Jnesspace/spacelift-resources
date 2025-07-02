METADATA:
  resource_type: spacelift_gitlab_integration
  provider: spacelift
  service: vcs_integration
  description: GitLab instance integration configuration
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_gitlab_integration" "RESOURCE_NAME" {
  name             = INTEGRATION_NAME
  api_host         = API_HOST         # URL or private://hostname
  user_facing_host = UI_HOST
  private_token    = GITLAB_TOKEN
  is_default       = false           # Optional
  space_id         = SPACE_ID        # Optional, defaults to root
  vcs_checks       = CHECK_TYPE      # Optional
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Integration identifier
      validation: Must be unique
      
    api_host:
      type: String
      description: API endpoint URL
      validation: Valid URL or private:// scheme
      
    user_facing_host:
      type: String
      description: UI endpoint URL
      validation: Valid HTTPS URL
      
    private_token:
      type: String
      description: GitLab API token
      validation: Valid access token
      sensitive: true

  optional:
    is_default:
      type: Boolean
      description: Default integration status
      default: false
      note: Requires root space if true
      
    space_id:
      type: String
      description: Target space identifier
      default: "root"
      validation: Must exist in Spacelift
      
    description:
      type: String
      description: Human-readable description
      default: ""
      
    labels:
      type: Set[String]
      description: Classification tags
      default: []
      
    vcs_checks:
      type: String
      description: VCS check configuration
      default: "INDIVIDUAL"
      allowed_values:
        - INDIVIDUAL
        - AGGREGATED
        - ALL

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    webhook_url:
      type: String
      description: Repository webhook URL
      generated: true
      
    webhook_secret:
      type: String
      description: Webhook verification secret
      generated: true
      sensitive: true

BEHAVIOR:
  connectivity:
    public:
      - Direct HTTPS access
      - Standard API endpoints
      - UI accessible publicly
      
    private:
      - VCS agent pool required
      - private:// URL scheme
      - Internal network access
      
  authentication:
    - Token-based auth
    - Token stored securely
    - Webhook verification
    
  webhooks:
    - Automatic configuration
    - Secret key generation
    - Event verification
    
  vcs_checks:
    INDIVIDUAL:
      - Separate status checks
      - Per-task reporting
      
    AGGREGATED:
      - Combined status check
      - Single report entry
      
    ALL:
      - Both check types
      - Complete reporting

PATTERNS:
  public_instance:
    example:
      api_host: "https://gitlab.example.com"
      is_default: false
    benefit: Direct connectivity
    
  private_instance:
    example:
      api_host: "private://gitlab"
      is_default: true
    benefit: Secure internal access

IMPORT_FORMAT: $INTEGRATION_ID