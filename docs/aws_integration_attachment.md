spacelift_aws_integration_attachment (Resource)

spacelift_aws_integration_attachment represents the attachment between a reusable AWS integration and a single stack or module.
Example Usage

# For a stack
resource "spacelift_aws_integration_attachment" "this" {
  integration_id = spacelift_aws_integration.this.id
  stack_id       = "my-stack-id"
  read           = true
  write          = true

  # The role needs to exist before we attach since we test role assumption during attachment.
  depends_on = [
    aws_iam_role.this
  ]
}

# For a module
resource "spacelift_aws_integration_attachment" "this" {
  integration_id = spacelift_aws_integration.this.id
  module_id      = "my-module-id"
  read           = true
  write          = true

  # The role needs to exist before we attach since we test role assumption during attachment.
  depends_on = [
    aws_iam_role.this
  ]
}

Schema
Required

    integration_id (String) ID of the integration to attach

Optional

    module_id (String) ID of the module to attach the integration to
    read (Boolean) Indicates whether this attachment is used for read operations. Defaults to true.
    stack_id (String) ID of the stack to attach the integration to
    write (Boolean) Indicates whether this attachment is used for write operations. Defaults to true.

Read-Only

    attachment_id (String) Internal ID of the attachment entity
    id (String) The ID of this resource.

