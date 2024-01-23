IAC_DIR = ./iac

tf_apply: tf_validate
	cd ${IAC_DIR} && terraform apply

tf_destroy:
	cd ${IAC_DIR} && terraform destroy

tf_validate:
	cd ${IAC_DIR} && terraform fmt && terraform validate