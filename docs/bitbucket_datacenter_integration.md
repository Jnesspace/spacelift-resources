# Resource: spacelift_bitbucket_datacenter_integration

## Description
Configures integration with Bitbucket Data Center (self-hosted Bitbucket) to enable repository access and webhook functionality for Spacelift stacks and modules.

## Example Usage
```hcl
# Public Bitbucket Data Center integration
resource "spacelift_bitbucket_datacenter_integration" "bitbucket_dc" {
  name             = "bitbucket-datacenter"
  api_host         = "https://bitbucket.company.com"
  user_facing_host = "https://bitbucket.company.com"
  username         = "spacelift-user"
  access_token     = var.bitbucket_token
  is_default       = true
  space_id         = "root"
}

# Private Bitbucket Data Center with VCS agent
resource "spacelift_bitbucket_datacenter_integration" "bitbucket_private" {
  name             = "bitbucket-private"
  api_host         = "private://bitbucket-internal"
  user_facing_host = "https://bitbucket.internal.company.com"
  username         = "spacelift-bot"
  access_token     = var.bitbucket_private_token
  description      = "Internal Bitbucket Data Center"
  labels           = ["internal", "bitbucket"]
  vcs_checks       = "AGGREGATED"
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique integration identifier
* `api_host` - (Required) Bitbucket API endpoint URL or `private://hostname` for VCS agent access
* `user_facing_host` - (Required) Bitbucket UI URL for user-facing links
* `username` - (Required) Bitbucket username for authentication
* `access_token` - (Required) Bitbucket access token or app password
* `is_default` - (Required) Whether this is the default Bitbucket integration

### Optional Arguments
* `description` - (Optional) Human-readable description of the integration
* `labels` - (Optional) Set of labels for integration classification
* `space_id` - (Optional) ID of the space the integration belongs to. Defaults to `"root"`
* `vcs_checks` - (Optional) VCS status check configuration. Valid values: `INDIVIDUAL`, `AGGREGATED`, `ALL`. Defaults to `INDIVIDUAL`

### Read-Only Arguments
* `id` - Unique resource identifier
* `webhook_url` - URL for Bitbucket webhooks
* `webhook_secret` - Secret for webhook verification (sensitive)

## Import
```bash
terraform import spacelift_bitbucket_datacenter_integration.example $INTEGRATION_ID
```

## Notes
* Default integrations must be in the root space
* Private access requires VCS agent pool configuration
* Webhook URL and secret are automatically generated for repository configuration