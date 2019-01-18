module "lb" {
  source  = "GoogleCloudPlatform/lb/google"
  version = "1.0.3"

  name         = "stagecoach"
  region       = "${local.region}"
  service_port = "${module.mig.service_port}"

  target_tags = [
    "${module.mig.target_tags}",
  ]
}
