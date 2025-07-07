# Resource: spacelift_policy

## Description
A policy is a collection of customer-defined rules written in Rego that are applied at specific decision points within Spacelift to control access, approvals, and automation behavior.

## Example Usage
```hcl
# Plan policy to prevent weekend deployments
resource "spacelift_policy" "no_weekend_deploys" {
  name = "No Weekend Deployments"
  body = file("${path.module}/policies/no-weekend-deploys.rego")
  type = "PLAN"
  description = "Blocks deployments on weekends"
}

# Access policy for team permissions
resource "spacelift_policy" "team_access" {
  name = "Team Access Control"
  body = file("${path.module}/policies/team-access.rego")
  type = "ACCESS"
  space_id = spacelift_space.team.id
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique policy identifier within the account
* `body` - (Required) Rego policy content defining the rules
* `type` - (Required) Policy decision point. Valid values: `ACCESS`, `APPROVAL`, `GIT_PUSH`, `INITIALIZATION`, `LOGIN`, `PLAN`, `TASK`, `TRIGGER`, `NOTIFICATION`

### Optional Arguments
* `description` - (Optional) Human-readable description of the policy
* `labels` - (Optional) Set of labels for policy classification
* `space_id` - (Optional) ID of the space the policy belongs to. Defaults to root space

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_policy.example $POLICY_ID
```

## Notes
* Policies must be written in valid Rego syntax
* Different policy types are evaluated at different points in the Spacelift workflow
* LOGIN policies apply globally and cannot be attached to individual stacks/modules