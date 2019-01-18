data "google_compute_image" "ubuntu" {
  family  = "ubuntu-1804-lts"
  project = "gce-uefi-images"
}
