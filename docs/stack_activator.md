
spacelift_stack_activator (Resource)

spacelift_stack_activator is used to to enable/disable Spacelift Stack.
Example Usage

resource "spacelift_stack" "app" {
  branch     = "master"
  name       = "Application stack"
  repository = "app"
}

resource "spacelift_stack_activator" "test" {
  enabled  = true
  stack_id = spacelift_stack.app.id
}

Schema
Required

    enabled (Boolean) Enable/disable stack
    stack_id (String) ID of the stack to enable/disable

Read-Only

    id (String) The ID of this resource.

On this page

    Example Usage
    Schema

Report an issue 