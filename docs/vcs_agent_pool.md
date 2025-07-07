# Resource: spacelift_vcs_agent_pool

## Description
A VCS agent pool manages a logical group of agents that provide secure access to private VCS installations, enabling Spacelift to connect to internal Git repositories.

## Example Usage
```hcl
# Basic VCS agent pool
resource "spacelift_vcs_agent_pool" "github_enterprise" {
  name        = "github-enterprise-pool"
  description = "Agent pool for GitHub Enterprise Server access"
}

# VCS agent pool for GitLab self-hosted
resource "spacelift_vcs_agent_pool" "gitlab_internal" {
  name        = "gitlab-internal"
  description = "Internal GitLab instance access"
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique agent pool identifier within the account

### Optional Arguments
* `description` - (Optional) Human-readable description of the agent pool

### Read-Only Arguments
* `id` - Unique resource identifier
* `config` - Base64-encoded agent configuration (sensitive)

## Import
```bash
terraform import spacelift_vcs_agent_pool.example $VCS_AGENT_POOL_ID
```

## Notes
* Agent pools enable access to private VCS installations
* Configuration is automatically generated and sensitive
* Agents must be deployed and configured separately to join the pool