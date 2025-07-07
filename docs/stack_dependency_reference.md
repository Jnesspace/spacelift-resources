# Resource: spacelift_stack_dependency_reference

## Description
Creates a reference to outputs from another stack, allowing one stack to consume the outputs of its dependency stack without creating a hard dependency relationship.

## Example Usage
```hcl
# Reference infrastructure stack outputs
resource "spacelift_stack_dependency_reference" "vpc_outputs" {
  stack_dependency_id = spacelift_stack_dependency.app_depends_on_infra.id
  output_name         = "vpc_id"
}

# Use the referenced output in your stack
resource "aws_instance" "app" {
  ami           = "ami-12345678"
  instance_type = "t3.micro"
  subnet_id     = spacelift_stack_dependency_reference.vpc_outputs.value
}
```

## Argument Reference

### Required Arguments
* `stack_dependency_id` - (Required) ID of the stack dependency relationship
* `output_name` - (Required) Name of the output to reference from the dependency stack

### Read-Only Arguments
* `id` - Unique resource identifier
* `value` - The actual value of the referenced output

## Notes
* The dependency stack must have a successful run with the specified output
* Referenced outputs are available as computed values in your configuration
* Changes to the dependency stack's outputs will trigger updates to dependent resources