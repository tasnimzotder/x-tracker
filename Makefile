IAC_DIR = ./iac
PLAYBOOKS_DIR = ./iac/playbooks

tf_apply: tf_validate
	cd ${IAC_DIR} && terraform apply

tf_destroy:
	cd ${IAC_DIR} && terraform destroy

tf_validate:
	cd ${IAC_DIR} && terraform fmt && terraform validate

ansible_verify:
	cd $(PLAYBOOKS_DIR) && ansible-inventory -i inventory.ini --list

ansible_ping:
	cd $(PLAYBOOKS_DIR) && ansible ec2_instance_fe -m ping -i inventory.ini

ansible_run:
	cd $(PLAYBOOKS_DIR) && ansible-playbook -i inventory.ini setup_ec2_all.yml

act_deploy:
	act --secret-file act.env --container-architecture linux/amd64 -W ./.github/workflows/deploy-fe.yml