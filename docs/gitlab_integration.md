# Resource: spacelift_gitlab_integration

## Description
Configures integration with a GitLab instance (GitLab.com or self-hosted) to enable repository access and webhook functionality for Spacelift stacks and modules.

## Example Usage
```hcl
# Public GitLab.com integration
resource "spacelift_gitlab_integration" "gitlab_com" {
  name             = "gitlab-com-integration"
  api_host         = "https://gitlab.com"
  user_facing_host = "https://gitlab.com"
  private_token    = var.gitlab_token
  space_id         = "root"
  is_default       = true
}

# Self-hosted GitLab integration
resource "spacelift_gitlab_integration" "gitlab_internal" {
  name             = "gitlab-internal"
  api_host         = "https://gitlab.company.com"
  user_facing_host = "https://gitlab.company.com"
  private_token    = var.gitlab_internal_token
  description      = "Internal GitLab instance"
  labels           = ["internal", "gitlab"]
  vcs_checks       = "AGGREGATED"
}

# Private GitLab with VCS agent
resource "spacelift_gitlab_integration" "gitlab_private" {
  name             = "gitlab-private"
  api_host         = "private://gitlab-internal"
  user_facing_host = "https://gitlab.internal.company.com"
  private_token    = var.gitlab_private_token
  vcs_checks       = "ALL"
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique integration identifier
* `api_host` - (Required) GitLab API endpoint URL or `private://hostname` for VCS agent access
* `user_facing_host` - (Required) GitLab UI URL for user-facing links
* `private_token` - (Required) GitLab personal access token

### Optional Arguments
* `description` - (Optional) Human-readable description of the integration
* `is_default` - (Optional) Whether this is the default GitLab integration. Defaults to `false`
* `labels` - (Optional) Set of labels for integration classification
* `space_id` - (Optional) ID of the space the integration belongs to. Defaults to `"root"`
* `vcs_checks` - (Optional) VCS status check configuration. Valid values: `INDIVIDUAL`, `AGGREGATED`, `ALL`. Defaults to `INDIVIDUAL`

### Read-Only Arguments
* `id` - Unique resource identifier
* `webhook_url` - URL for GitLab webhooks
* `webhook_secret` - Secret for webhook verification (sensitive)

## Import
```bash
terraform import spacelift_gitlab_integration.example $INTEGRATION_ID
```

## Notes
* Default integrations must be in the root space
* Private access requires VCS agent pool configuration
* Webhook URL and secret are automatically generated for repository configuration