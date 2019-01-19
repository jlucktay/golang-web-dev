module "sql" {
  source  = "GoogleCloudPlatform/sql-db/google"
  version = "1.0.1"

  db_charset = "utf8mb4"                             // https://medium.com/@adamhooper/in-mysql-never-use-utf8-use-utf8mb4-11761243e434
  db_name    = "pocket01"
  name       = "saddlebag-${random_id.sql_name.hex}"
  region     = "${local.region}"

  ip_configuration = [
    {
      private_network = "${data.google_compute_network.default.self_link}"
    },
  ]
}

# https://github.com/GoogleCloudPlatform/terraform-google-sql-db/blob/master/variables.tf
/*
module "mysql-db" {
  database_flags = [
    {
      name  = "log_bin_trust_function_creators"
      value = "on"
    },
  ]
*/
