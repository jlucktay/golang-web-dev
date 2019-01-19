data "template_file" startup_script {
  template = "${file("${path.module}/gce.startup.sh")}"

  vars = {
    mysql_ip       = ""
    mysql_password = "${module.sql.generated_user_password}"
  }
}
