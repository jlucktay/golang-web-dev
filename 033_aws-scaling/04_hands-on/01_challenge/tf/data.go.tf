data "template_file" "gce_startup" {
  template = "${file("${path.module}/gce.startup.sh")}"

  vars {
    main_go = "${file("${path.module}/go/main.go")}}"
  }
}
