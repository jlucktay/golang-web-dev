data "template_file" "gce_startup" {
  template = "${file("${path.module}/gce.startup.sh")}"

  vars {
    toprc = "" #"${file("${path.module}/.toprc")}}"
  }
}
