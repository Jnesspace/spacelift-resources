
spacelift_terraform_provider (Resource)

spacelift_terraform_provider represents a Terraform provider in Spacelift's own provider registry.
Example Usage

resource "spacelift_terraform_provider" "datadog" {
  type     = "datadog"
  space_id = "root"

  description = "Our fork of the Datadog provider"
  labels      = ["fork"]
  public      = false
}

Schema
Required

    space_id (String) ID (slug) of the space the provider is in
    type (String) Type of the provider - should be unique in one account

Optional

    description (String) Free-form description for human users, supports Markdown
    labels (Set of String)
    public (Boolean) Whether the provider is public or not, defaults to false (private)

Read-Only

    id (String) The ID of this resource.

On this page

    Example Usage
    Schema

Report an issue 