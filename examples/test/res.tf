# epilot-app_file.my_file:
# resource "epilot-app_file" "my_file" {
    # access_control = "private"
    # filename       = "8z87al.jpg"
    # mime_type      = "image/jpeg"
    # source_url     = "https://file.dev.sls.epilot.io/v1/files/public/links/739224-taqDBUyNjSX7c_n2v6Zvv/8z87al.jpg"
    # title          = "8z87al.jpg"
    # type           = "image"
# }

terraform {
  required_providers {
    epilot-app = {
      source  = "epilot-dev/epilot-app"
      version = "2.1.1"
    }
  }
}

variable "epilot_auth" {
  type        = string
  description = "epilot_auth"
}


provider "epilot-app" {
  # Configuration options
  epilot_auth = var.epilot_auth
  server_url  = "https://file.sls.epilot.io"
}
