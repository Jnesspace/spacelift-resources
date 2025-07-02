METADATA:
  resource_type: spacelift_stack
  provider: spacelift
  service: core
  description: Runtime environment combining source code and configuration for resource management
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_stack" "RESOURCE_NAME" {
  name              = STACK_NAME
  repository        = REPOSITORY_NAME
  branch            = BRANCH_NAME
  project_root      = PROJECT_PATH         # Optional
  description       = STACK_DESCRIPTION    # Optional
  terraform_version = TERRAFORM_VERSION    # Optional
}
```

ATTRIBUTES:
  required:
    branch:
      type: String
      description: Git branch to apply changes to
      validation: Must be valid git branch name
    
    name:
      type: String
      description: Unique stack identifier within account
      validation: Must be unique
    
    repository:
      type: String
      description: Repository name without owner part
      validation: Must exist in VCS

  optional:
    administrative:
      type: Boolean
      description: Stack can manage other stacks
      default: false
      
    autodeploy:
      type: Boolean
      description: Enables automatic deployment of changes
      default: false
      
    project_root:
      type: String
      description: Directory containing stack entrypoint
      default: ""
      
    terraform_version:
      type: String
      description: Terraform version to use
      validation: Valid semver
      
    description:
      type: String
      description: Human-readable stack description
      
    labels:
      type: Set[String]
      description: Stack classification tags
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
    - github (default)
    - github_enterprise
    - gitlab
    - bitbucket_cloud
    - bitbucket_datacenter
    - azure_devops
    - raw_git

  stack_types:
    terraform:
      default: true
      version_required: true
      
    terragrunt:
      config:
        terraform_version: String
        terragrunt_version: String
        use_run_all: Boolean
        use_smart_sanitization: Boolean
        
    cloudformation:
      config:
        entry_template_file: String (required)
        region: String (required)
        stack_name: String (required)
        template_bucket: String (required)
        
    pulumi:
      config:
        login_url: String (required)
        stack_name: String (required)
        
    kubernetes:
      config:
        namespace: String
        kubectl_version: String
        
    ansible:
      config:
        playbook: String (required)

BEHAVIOR:
  state_management:
    - Manages state by default (manage_state=true)
    - Can enable external state access (terraform_external_state_access)
    - Supports state import on creation
    
  security:
    - Deletion protection available
    - Sensitive output handling
    - Well-known secret masking
    
  automation:
    - Supports automatic deployment (autodeploy)
    - Supports local preview runs
    - Configurable worker pools
    - Extensive hook system (before/after scripts)

  hooks:
    before:
      - init
      - plan
      - apply
      - perform
      - destroy
    after:
      - init
      - plan
      - apply
      - perform
      - destroy
      - run