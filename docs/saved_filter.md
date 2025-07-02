METADATA:
  resource_type: spacelift_saved_filter
  provider: spacelift
  service: organization
  description: Custom filter criteria for Spacelift views
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_saved_filter" "RESOURCE_NAME" {
  name      = FILTER_NAME
  type      = FILTER_TYPE    # stacks, blueprints, contexts, webhooks
  is_public = true          # Controls visibility
  data      = jsonencode({
    key   = "activeFilters"
    value = jsonencode({
      filters = FILTER_CRITERIA,
      sort    = SORT_CONFIG,
      order   = COLUMN_CONFIG
    })
  })
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Filter identifier
      validation: Must be unique per type
      
    type:
      type: String
      description: Target view type
      allowed_values:
        - stacks
        - blueprints
        - contexts
        - webhooks
        
    is_public:
      type: Boolean
      description: Filter visibility
      validation: true/false
      
    data:
      type: String
      description: JSON filter configuration
      validation: Must be valid JSON
      structure:
        key: String
        value:
          filters: Array of criteria
          sort: Sort configuration
          order: Column visibility

  computed:
    id:
      type: String
      description: Unique resource identifier
      generated: true
      
    created_by:
      type: String
      description: Creator's login
      generated: true

BEHAVIOR:
  filtering:
    - Defines view criteria
    - Supports multiple fields
    - Complex filter logic
    
  visibility:
    - Public or private filters
    - Shared across users
    - View-specific scope
    
  configuration:
    - JSON-based definition
    - Customizable columns
    - Sorting options
    
  views:
    stacks:
      - Filter stack resources
      - Stack-specific fields
      - Stack organization
    blueprints:
      - Template filtering
      - Blueprint management
      - Template organization
    contexts:
      - Context filtering
      - Configuration management
      - Context grouping
    webhooks:
      - Webhook filtering
      - Integration management
      - Endpoint organization

IMPORT_FORMAT: $FILTER_ID

EXAMPLE:
  webhook_filter:
    type: webhooks
    data:
      filters:
        - name:
            type: STRING
            values: ["team_xyz_*"]
      sort:
        direction: ASC
        option: space
      order:
        - name: enabled
          visible: true
        - name: endpoint
          visible: true