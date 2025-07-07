# Resource: spacelift_blueprint

## Description
A blueprint is a template for creating stacks using a templating engine. Primarily useful for non-Terraform users who need to create stacks programmatically.

## Example Usage
```hcl
# Published blueprint
resource "spacelift_blueprint" "vpc_template" {
  name        = "vpc-blueprint"
  space       = spacelift_space.templates.id
  state       = "PUBLISHED"
  description = "Standard VPC deployment template"
  template    = file("${path.module}/templates/vpc-blueprint.yaml")
  labels      = ["aws", "vpc", "template"]
}

# Draft blueprint
resource "spacelift_blueprint" "experimental" {
  name        = "experimental-blueprint"
  space       = "root"
  state       = "DRAFT"
  description = "Experimental infrastructure template"
  labels      = ["experimental"]
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique blueprint identifier within the space
* `space` - (Required) ID of the space the blueprint belongs to
* `state` - (Required) Blueprint state. Valid values: `DRAFT`, `PUBLISHED`

### Optional Arguments
* `description` - (Optional) Human-readable description of the blueprint
* `template` - (Optional) Blueprint template content. Required when state is `PUBLISHED`
* `labels` - (Optional) Set of labels for blueprint classification

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* For Terraform users, prefer using `spacelift_stack` directly
* Template content is required when publishing a blueprint
* Blueprints in DRAFT state cannot be used to create stacks

