---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-app_app Data Source - terraform-provider-epilot-app"
subcategory: ""
description: |-
  App DataSource
---

# epilot-app_app (Data Source)

App DataSource

## Example Usage

```terraform
data "epilot-app_app" "my_app" {
  app_id = "...my_app_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String)

### Read-Only

- `blueprint_ref` (Attributes) (see [below for nested schema](#nestedatt--blueprint_ref))
- `components` (String) List of component configurations for the installed version. Parsed as JSON.
- `enabled` (Boolean) Flag to indicate if the app is enabled. Enabled is set to true when required option values are set.
- `installation_audit` (Attributes) Audit information for the app (see [below for nested schema](#nestedatt--installation_audit))
- `installed_version` (String) Version of the app that is installed
- `installer_org_id` (String) Unique identifier for the organization the app is installed in
- `manifest` (List of String) Manifest ID used to create/update the entity
- `name` (String) Name of the app
- `option_values` (Attributes List) Configuration values for the app components (see [below for nested schema](#nestedatt--option_values))
- `owner_org_id` (String) Organization ID of the app creator
- `role` (String) The name of the role the app can use to access APIs

<a id="nestedatt--blueprint_ref"></a>
### Nested Schema for `blueprint_ref`

Read-Only:

- `job_id` (String) ID of the job that created the blueprint
- `manifest_id` (String) ID of the blueprint


<a id="nestedatt--installation_audit"></a>
### Nested Schema for `installation_audit`

Read-Only:

- `created_at` (String) Timestamp of the creation
- `created_by` (String) User ID of the creator
- `updated_at` (String) Timestamp of the last update
- `updated_by` (String) User ID of the last updater


<a id="nestedatt--option_values"></a>
### Nested Schema for `option_values`

Read-Only:

- `component_id` (String) ID of the component these values are for
- `options` (Attributes List) (see [below for nested schema](#nestedatt--option_values--options))

<a id="nestedatt--option_values--options"></a>
### Nested Schema for `option_values.options`

Read-Only:

- `key` (String) Key matching a config_option from the component
- `value` (Attributes) The configured value for this option (see [below for nested schema](#nestedatt--option_values--options--value))

<a id="nestedatt--option_values--options--value"></a>
### Nested Schema for `option_values.options.value`

Read-Only:

- `boolean` (Boolean)
- `number` (Number)
- `str` (String)
