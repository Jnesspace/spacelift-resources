# Resource: spacelift_version

## Description
Manages version tags for modules, allowing semantic versioning and controlled releases of Terraform modules in Spacelift.

## Example Usage
```hcl
# Create a version tag for a module
resource "spacelift_version" "vpc_v1_0_0" {
  module_id   = spacelift_module.vpc.id
  version_number = "1.0.0"
  commit_sha  = "abc123def456789"
  description = "Initial stable release of VPC module"
}

# Pre-release version
resource "spacelift_version" "vpc_beta" {
  module_id      = spacelift_module.vpc.id
  version_number = "1.1.0-beta.1"
  commit_sha     = "def456abc123789"
  description    = "Beta release with new features"
}
```

## Argument Reference

### Required Arguments
* `module_id` - (Required) ID of the module to create a version for
* `version_number` - (Required) Semantic version number (e.g., "1.0.0")
* `commit_sha` - (Required) Git commit SHA to tag as this version

### Optional Arguments
* `description` - (Optional) Human-readable description of this version

### Read-Only Arguments
* `id` - Unique resource identifier

## Notes
* Version numbers should follow semantic versioning conventions
* Commit SHA must exist in the module's repository
* Versions enable controlled consumption of modules by stacks