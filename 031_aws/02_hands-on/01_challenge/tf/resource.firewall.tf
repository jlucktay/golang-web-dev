resource "google_compute_firewall" "main" {
  name = "default-allow-http"

  network       = "default"
  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["http-server"]

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }
}
