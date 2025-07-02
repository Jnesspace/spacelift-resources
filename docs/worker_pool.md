spacelift_worker_pool (Resource)

spacelift_worker_pool represents a worker pool assigned to the Spacelift account.
Example Usage

resource "spacelift_worker_pool" "k8s-core" {
  name        = "Main worker"
  csr         = filebase64("/path/to/csr")
  description = "Used for all type jobs"
}

Schema
Required

    name (String) name of the worker pool

Optional

    csr (String, Sensitive) certificate signing request in base64. Changing this value will trigger a token reset.
    description (String) description of the worker pool
    labels (Set of String)
    space_id (String) ID (slug) of the space the worker pool is in

Read-Only

    config (String, Sensitive) credentials necessary to connect WorkerPool's workers to the control plane
    id (String) The ID of this resource.
    private_key (String, Sensitive) private key in base64

Import

Import is supported using the following syntax:

terraform import spacelift_worker_pool.k8s-core $WORKER_POOL_ID