/* module "sql" {
  source  = "GoogleCloudPlatform/sql-db/google"
  version = "1.0.1"

  name   = "saddlebag-${random_id.sql_name.hex}"
  region = "${local.region}"
}
 */
# https://github.com/GoogleCloudPlatform/terraform-google-sql-db/blob/master/variables.tf
/*
module "mysql-db" {
  ip_configuration = [{
    authorized_networks = [{
      name  = "${var.network_name}"
      value = "${google_compute_subnetwork.default.ip_cidr_range}"
    }]
  }]

  database_flags = [
    {
      name  = "log_bin_trust_function_creators"
      value = "on"
    },
  ]
*/
