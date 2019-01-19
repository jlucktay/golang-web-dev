data "template_file" startup_script {
  template = "${file("${path.module}/gce.startup.sh")}"

  vars = {
    mysql_ip       = "${data.external.gcloud_sql.result.ipAddress}"
    mysql_password = "${module.sql.generated_user_password}"
  }
}

data "external" gcloud_sql {
  program = [
    "/usr/local/bin/bash",
    "-c",
    "${format("gcloud sql instances list --project golang-web-dev-227919 --format json | jq -r '.[].ipAddresses[] | select( .type == \"PRIVATE\" )'")}",
  ]
}
