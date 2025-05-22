resource "epilot-app_app" "my_app" {
  app_id = "...my_app_id..."
  option_values = [
    {
      component_id = "...my_component_id..."
      options = [
        {
          key   = "...my_key..."
          value = "...my_value..."
        }
      ]
    }
  ]
  version = "...my_version..."
}