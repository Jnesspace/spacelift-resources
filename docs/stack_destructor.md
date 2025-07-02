
spacelift_stack_destructor (Resource)

spacelift_stack_destructor is used to destroy the resources of a Stack before deleting it. depends_on should be used to make sure that all necessary resources (environment variables, roles, integrations, etc.) are still in place when the destruction run is executed. Note: Destroying this resource will delete the resources in the stack. If this resource needs to be deleted and the resources in the stacks are to be preserved, ensure that the deactivated attribute is set to true.
Example Usage

resource "spacelift_stack" "k8s-core" {
  // ...
}

resource "spacelift_environment_variable" "credentials" {
  // ...
}

resource "spacelift_stack_destructor" "k8s-core" {
  depends_on = [
    spacelift_environment_variable.credentials,
  ]

  stack_id = spacelift_stack.k8s-core.id
}

Schema
Required

    stack_id (String) ID of the stack to delete and destroy on destruction

Optional

    deactivated (Boolean) If set to true, destruction won't delete the stack
    timeouts (Block, Optional) (see below for nested schema)

Read-Only

    id (String) The ID of this resource.

Nested Schema for timeouts

Optional:

    delete (String)

On this page

    Example Usage
    Schema

Report an issue 