provider "google" {
  credentials = "${file("golang-web-dev-b8f5f688f0c9.json")}"
  project     = "golang-web-dev-227919"
  region      = "${local.region}"
  zone        = "${local.zone}"
}
