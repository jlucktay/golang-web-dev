module "sql" {
  source  = "GoogleCloudPlatform/sql-db/google"
  version = "1.0.1"

  name   = "saddlebag-${random_id.sql_name.hex}"
  region = "${local.region}"
}
