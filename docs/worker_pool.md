# Resource: spacelift_worker_pool

## Description
A worker pool is a logical group of workers that execute Spacelift runs. Worker pools allow you to control where and how your infrastructure code is executed.

## Example Usage
```hcl
# Basic worker pool
resource "spacelift_worker_pool" "main" {
  name        = "main-workers"
  description = "Primary worker pool for production workloads"
}

# Worker pool with CSR for private workers
resource "spacelift_worker_pool" "private" {
  name        = "private-workers"
  description = "Private worker pool for secure workloads"
  csr         = filebase64("${path.module}/worker.csr")
  labels      = ["private", "secure"]
}
```

## Argument Reference

### Required Arguments
* `name` - (Required) Unique worker pool identifier within the account

### Optional Arguments
* `description` - (Optional) Human-readable description of the worker pool
* `csr` - (Optional) Base64-encoded certificate signing request for private workers
* `labels` - (Optional) Set of labels for worker pool classification
* `space_id` - (Optional) ID of the space the worker pool belongs to

### Read-Only Arguments
* `id` - Unique resource identifier
* `config` - Base64-encoded worker configuration (sensitive)
* `private_key` - Base64-encoded private key for worker authentication (sensitive)

## Import
```bash
terraform import spacelift_worker_pool.example $WORKER_POOL_ID
```

## Notes
* Private workers require a CSR for certificate-based authentication
* Configuration and private key are generated automatically and are sensitive
* Worker pools can be assigned to specific stacks or modules