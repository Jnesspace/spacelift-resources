
spacelift_vcs_agent_pool (Resource)

spacelift_vcs_agent_pool represents a Spacelift VCS agent pool - a logical group of proxies allowing Spacelift to access private VCS installations
Example Usage

resource "spacelift_vcs_agent_pool" "ghe" {
  name        = "ghe"
  description = "VCS agent pool for our internal GitHub Enterprise"
}

Schema
Required

    name (String) Name of the VCS agent pool, must be unique within an account

Optional

    description (String) Free-form VCS agent pool description for users

Read-Only

    config (String, Sensitive) VCS agent pool configuration, encoded using base64
    id (String) The ID of this resource.

Import

Import is supported using the following syntax:

terraform import spacelift_vcs_agent_pool.ghe $VCS_AGENT_POOL_ID

On this page

    Example Usage
    Schema
    Import

Report an issue 