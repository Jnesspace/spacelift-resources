spacelift_policy_attachment (Resource)

spacelift_policy_attachment represents a relationship between a policy (spacelift_policy) and a stack (spacelift_stack) or module (spacelift_module). Each policy can only be attached to a stack/module once. LOGIN policies are the exception because they apply globally and not to individual stacks/modules. An attempt to attach one will fail.
Example Usage

resource "spacelift_policy" "no-weekend-deploys" {
  name = "Let's not deploy any changes over the weekend"
  body = file("policies/no-weekend-deploys.rego")
  type = "PLAN"
}

resource "spacelift_stack" "core-infra-production" {
  name       = "Core Infrastructure (production)"
  branch     = "master"
  repository = "core-infra"
}

resource "spacelift_policy_attachment" "no-weekend-deploys" {
  policy_id = spacelift_policy.no-weekend-deploys.id
  stack_id  = spacelift_stack.core-infra-production.id
}

Schema
Required

    policy_id (String) ID of the policy to attach

Optional

    module_id (String) ID of the module to attach the policy to
    stack_id (String) ID of the stack to attach the policy to

Read-Only

    id (String) The ID of this resource.

Import

Import is supported using the following syntax:

terraform import spacelift_policy_attachment.no-weekend-deploys $POLICY_ID/$STACK_ID

On this page

    Example Usage
    Schema
    Import

Report an issue 