# TODO: look into switching over to this:
# https://cloud.google.com/compute/docs/instances/managing-instance-access

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
    minimal_action = "REPLACE"
    type           = "PROACTIVE"
  }]

  ssh_source_ranges = [
    "0.0.0.0/0",
  ]

  target_pools = [
    "${module.lb.target_pool}",
  ]
}
