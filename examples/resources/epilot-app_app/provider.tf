terraform {
  required_providers {
    epilot-app = {
      source  = "epilot-dev/epilot-app"
      version = "0.6.2"
    }
  }
}

provider "epilot-app" {
  # Configuration options
  epilot_auth = "my-auth-token"
  server_url = "http://localhost:8081"
}