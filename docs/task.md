
spacelift_task (Resource)

spacelift_task represents a task in Spacelift.
Schema
Required

    command (String) Command that will be run.
    stack_id (String) ID of the stack for which to run the task

Optional

    init (Boolean) Whether to initialize the stack or not. Default: true
    keepers (Map of String) Arbitrary map of values that, when changed, will trigger recreation of the resource.
    timeouts (Block, Optional) (see below for nested schema)
    wait (Block List, Max: 1) Wait for the run to finish (see below for nested schema)

Read-Only

    id (String) The ID of this resource.

Nested Schema for timeouts

Optional:

    create (String)

Nested Schema for wait

Optional:

    continue_on_state (Set of String) Continue on the specified states of a finished run. If not specified, the default is [ 'finished' ]. You can use following states: applying, canceled, confirmed, destroying, discarded, failed, finished, initializing, pending_review, performing, planning, preparing_apply, preparing_replan, preparing, queued, ready, replan_requested, skipped, stopped, unconfirmed.
    continue_on_timeout (Boolean) Continue if task timed out, i.e. did not reach any defined end state in time. Default: false
    disabled (Boolean) Whether waiting for the task is disabled or not. Default: false

On this page

    Schema

Report an issue 