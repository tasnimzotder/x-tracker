#provider "local" {}
#
#resource "null_resource" "setup-docker" {
#  depends_on = [
#    aws_instance.x-tracker-instances
#  ]
#  count = length(local.instances)
#
#  provisioner "local-exec" {
#    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i admin@${aws_instance.x-tracker-instances[count.index].public_dns}, playbooks/setup_ec2_all.yml"
#  }
#}