module "mig" {
  source  = "GoogleCloudPlatform/managed-instance-group/google"
  version = "1.1.15"

  hc_path           = "/ping"
  name              = "gunslinger"
  region            = "${local.region}"
  service_port      = "${local.service_port}"
  service_port_name = "http"
  size              = 3
  target_tags       = "${local.target_tags}"
  zone              = "${local.zone}"

  ssh_source_ranges = [
    "0.0.0.0/0",
  ]

  target_pools = [
    "${module.lb.target_pool}",
  ]
}
