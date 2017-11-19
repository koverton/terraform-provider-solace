
terraform-provider-solace : 
	go build


install : terraform-provider-solace
	go install

clean : 
	$(RM) -rf terraform-provider-solace terraform.tfstate* .terraform
