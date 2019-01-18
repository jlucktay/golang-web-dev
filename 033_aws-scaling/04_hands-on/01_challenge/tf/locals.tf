locals {
  region       = "europe-west2"   # London
  service_port = 80
  zone         = "europe-west2-a"

  target_tags = [
    "http-server",
  ]
}
