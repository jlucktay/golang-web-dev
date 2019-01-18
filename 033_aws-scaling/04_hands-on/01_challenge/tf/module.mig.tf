module "mig" {
  source  = "GoogleCloudPlatform/managed-instance-group/google"
  version = "1.1.15"

  compute_image     = "${data.google_compute_image.ubuntu.self_link}"
  hc_path           = "/ping"
  name              = "gunslinger"
  region            = "${local.region}"
  service_port      = "${local.service_port}"
  service_port_name = "http"
  startup_script    = "${data.template_file.gce_startup.rendered}"
  size              = 3
  target_tags       = "${local.target_tags}"
  update_strategy   = "ROLLING_UPDATE"
  zone              = "${local.zone}"

  rolling_update_policy = [{
    max_unavailable_percent = 100
    minimal_action          = "REPLACE"
    type                    = "PROACTIVE"
  }]

  ssh_source_ranges = [
    "0.0.0.0/0",
  ]

  target_pools = [
    "${module.lb.target_pool}",
  ]
}

# startup_script
## get the go server running on the instance
### build it locally
### copy the binary
