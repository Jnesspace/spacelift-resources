# spacelift_azure_integration

METADATA:
  resource_type: spacelift_azure_integration
  provider: spacelift
  service: cloud_integration
  description: Azure AD tenant integration for Spacelift resource management
  version: latest

USAGE_TEMPLATE:
```hcl
resource "spacelift_azure_integration" "RESOURCE_NAME" {
  name                    = INTEGRATION_NAME
  tenant_id               = AZURE_TENANT_ID
  default_subscription_id = SUBSCRIPTION_ID
  labels                  = [LABEL_VALUES]
}
```

ATTRIBUTES:
  required:
    name:
      type: String
      description: Friendly identifier for the integration
      validation: none
    
    tenant_id:
      type: String
      description: Azure AD tenant identifier
      validation: Valid UUID format

  optional:
    default_subscription_id:
      type: String
      description: Default Azure subscription for resource management
      validation: Valid UUID format
    
    labels:
      type: Set[String]
      description: Resource classification tags
      default: []
      validation: none
    
    space_id:
      type: String
      description: Target space identifier
      validation: Must exist in Spacelift

  computed:
    admin_consent_provided:
      type: Boolean
      description: Confirmation of admin consent status
      generated: true
    
    admin_consent_url:
      type: String
      description: URL for granting admin consent
      generated: true
    
    application_id:
      type: String
      description: Azure AD application identifier
      generated: true
    
    display_name:
      type: String
      description: Auto-generated Azure application name
      generated: true
      immutable: true
    
    id:
      type: String
      description: Unique resource identifier
      generated: true

BEHAVIOR:
  requirements:
    - Explicit stack attachment required after creation
    - Manual admin consent required via admin_consent_url
    - Application display name cannot be modified without recreation

  permissions:
    - Azure AD application permissions must be granted
    - Subscription access must be configured if specified
    
  lifecycle:
    create:
      - Azure AD application is created
      - Admin consent URL is generated
      - Awaits manual consent
    
    update:
      - Most fields can be modified
      - display_name changes require recreation
      
    delete:
      - Removes Azure AD application
      - Revokes all access