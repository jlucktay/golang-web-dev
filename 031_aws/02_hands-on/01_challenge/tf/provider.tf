provider "google" {
  credentials = "${file("golang-web-dev-b8f5f688f0c9.json")}"
  project     = "golang-web-dev-227919"
  region      = "europe-west2"                                # London
  zone        = "europe-west2-a"
}
