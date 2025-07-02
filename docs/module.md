METADATA:
  resource_type: spacelift_module
  provider: spacelift
  service: terraform_modules
  description: Special stack type for testing and versioning Terraform modules
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_module" "RESOURCE_NAME" {
  # Required fields
  repository = REPOSITORY_NAME
  branch     = BRANCH_NAME

  # Optional fields with auto-inference
  name               = MODULE_NAME               # Optional, inferred from repository
  terraform_provider = PROVIDER_NAME            # Optional, inferred from repository
  
  # Common optional fields
  description        = DESCRIPTION              # Optional
  project_root      = MODULE_SOURCE_PATH       # Optional
  administrative    = false                    # Optional
  public           = false                    # Optional
}
```

ATTRIBUTES:
  required:
    branch:
      type: String
      description: Git branch to apply changes to
      validation: Must be valid git branch name
    
    repository:
      type: String
      description: Repository name without owner part
      validation: Must exist in VCS
      note: Repository name format terraform-PROVIDER-NAME enables auto-inference

  optional:
    name:
      type: String
      description: Module identifier
      validation: Must be unique in account
      note: Auto-inferred from repository if following naming convention
      
    terraform_provider:
      type: String
      description: Terraform provider used by module
      note: Auto-inferred from repository if following naming convention
      
    administrative:
      type: Boolean
      description: Module can manage other resources
      default: false
      
    description:
      type: String
      description: Human-readable module description
      
    project_root:
      type: String
      description: Directory containing module source
      default: ""
      
    public:
      type: Boolean
      description: Public accessibility flag
      default: false
      immutable: true
      
    protect_from_deletion:
      type: Boolean
      description: Prevents accidental deletion
      default: false
      
    shared_accounts:
      type: Set[String]
      description: Accounts with module access
      default: []
      
    space_id:
      type: String
      description: Target space identifier
      default: root or legacy space
      
    workflow_tool:
      type: String
      description: IaC execution tool
      allowed_values: ["OPEN_TOFU", "TERRAFORM_FOSS", "CUSTOM"]
      default: "TERRAFORM_FOSS"

    worker_pool_id:
      type: String
      description: Worker pool for execution
      note: Required for self-hosted instances
      
    labels:
      type: Set[String]
      description: Module classification tags
      default: []

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    aws_assume_role_policy_statement:
      type: String
      description: AWS IAM assume role policy for trust setup
      generated: true

INTEGRATIONS:
  vcs_providers:
    github:
      default: true
      
    github_enterprise:
      required:
        namespace: String # GitHub org/user
      optional:
        id: String # Integration ID
        
    gitlab:
      required:
        namespace: String # GitLab namespace
      optional:
        id: String # Integration ID
        
    bitbucket_cloud:
      required:
        namespace: String # Bitbucket project
      optional:
        id: String # Integration ID
        
    bitbucket_datacenter:
      required:
        namespace: String # Bitbucket project
      optional:
        id: String # Integration ID
        
    azure_devops:
      required:
        project: String # Azure DevOps project
      optional:
        id: String # Integration ID
        
    raw_git:
      required:
        url: String # HTTPS repository URL
        namespace: String # Display name

BEHAVIOR:
  naming:
    - Auto-infers name and provider from repository name
    - Repository format: terraform-PROVIDER-NAME
    - Custom names override auto-inference
    
  security:
    - Public/private visibility set at creation
    - Supports account sharing
    - Deletion protection available
    
  execution:
    - Runs in specified worker pool
    - Supports different IaC tools
    - Can initialize and test module code

IMPORT_FORMAT: $MODULE_ID