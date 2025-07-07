# Resource: spacelift_space

## Description
A space is a logical collection of Spacelift resources that enables hierarchical organization and access control. Spaces can contain stacks, modules, contexts, policies, and other resources.

## Example Usage
```hcl
# Development space
resource "spacelift_space" "development" {
  name            = "development"
  parent_space_id = "root"
  description     = "Development environment resources"
  inherit_entities = true
}

# Team-specific space
resource "spacelift_space" "frontend_team" {
  name            = "frontend-team"
  parent_space_id = spacelift_space.development.id
  description     = "Frontend team development resources"
  labels          = ["team:frontend", "env:dev"]
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique space identifier within the account

### Optional Arguments
* `parent_space_id` - (Optional) Parent space ID. Defaults to "root". Cannot be changed after creation
* `description` - (Optional) Human-readable description of the space
* `inherit_entities` - (Optional) Whether to inherit read access from parent space. Defaults to `false`
* `labels` - (Optional) Set of labels for space classification

### Read-Only Arguments
* `id` - Unique resource identifier

## Import
```bash
terraform import spacelift_space.example $SPACE_ID
```

## Notes
* Requires root admin permissions to create or manage spaces
* Parent space relationship cannot be changed after creation
* All spaces except root must have a parent space
* Entity inheritance allows child spaces to read parent space resources