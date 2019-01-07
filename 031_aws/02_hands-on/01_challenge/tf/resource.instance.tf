resource "google_compute_instance" "main" {
  name = "hello-go"

  machine_type = "f1-micro"
  tags         = ["http-server"]

  boot_disk {
    initialize_params {
      image = "ubuntu-1810"
    }
  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral IP
    }
  }

  # metadata {
  #   foo = "bar"
  # }

  # metadata_startup_script = "echo hi > /test.txt"
}
