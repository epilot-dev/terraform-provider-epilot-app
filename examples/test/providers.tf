# Provider definitions
terraform {
  required_providers {
    epilot-app = {
      source  = "epilot-dev/epilot-app"
      version = "0.0.2"
    }
  }
}

# Variables
variable "epilot_auth" {
  type = string
}

variable "app_api_url" {
  type    = string
  default = "https://app.sls.epilot.io"
}