resource "epilot-app_app" "my_app" {
  app_id = "...my_app_id..."
  option_values = [
    {
      component_id = "...my_component_id..."
      options = [
        {
          key = "...my_key..."
          value = {
            boolean = false
            number  = 4.53
            str     = "...my_str..."
          }
        }
      ]
    }
  ]
  version = "...my_version..."
}