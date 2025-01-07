resource "epilot-app_app" "my_app" {
  app_id = "...my_app_id..."
  config_options = {
    component_type = "CUSTOM_JOURNEY_BLOCK"
    configuration = [
      {
        key   = "...my_key..."
        type  = "number"
        value = "...my_value..."
      }
    ]
  }
  enabled = false
}