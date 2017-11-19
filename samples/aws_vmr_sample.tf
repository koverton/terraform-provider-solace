
provider "aws" {
  access_key = "AKIAIORZGWWVDPM3RQQQ"
  secret_key = "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"
  region     = "us-east-1"
}

resource "aws_instance" "terraform_vmr_inst" {
  ami = "ami-1541b36f"
  instance_type = "t2.medium"
  key_name = "koverton_aws_east"
  security_groups = [ "Default VMR Group" ]
  tags {
    Name = "terraform_vmr_inst"
  }
}


#resource "aws_eip" "ip" {
#  instance = "${aws_instance.jpmc_terraform_vmr_inst.id}"
#}


# host =  "192.168.56.201"
provider "solace" {
  host           = "${aws_instance.terraform_vmr_inst.public_ip}"
  port           = 8080
  admin_user     = "admin"
  admin_password = "admin"
}

resource "solace_msg_vpn" "test_vpn" {
  msg_vpn_name = "test_vpn"
  max_msg_spool_usage = 1000
  authentication_basic_enabled = true
  authentication_basic_type = "none"
  enabled = true
}

resource "solace_queue" "q_jimmy" {
  queue_name  = "jimmy"
  msg_vpn_name = "test_vpn"
  access_type = "exclusive"
  max_msg_spool_usage = 1000
  permission = "modify-topic"
  depends_on = ["solace_msg_vpn.test_vpn"]
}

resource "solace_topic_endpoint" "dte_jimmy" {
  topic_endpoint_name  = "jimmy"
  msg_vpn_name = "test_vpn"
  max_spool_usage = 1000
  permission = "modify-topic"
  depends_on = ["solace_msg_vpn.test_vpn"]
}

resource "solace_client_profile"  "profile_jimmy" {
  client_profile_name = "jimmy"
  msg_vpn_name = "test_vpn"
  max_connection_count_per_client_username = 10
  allow_guaranteed_endpoint_create_enabled = true
  allow_guaranteed_msg_send_enabled = true
  allow_guaranteed_msg_receive_enabled = true
}

resource "solace_client_username" "user_jimmy" {
  client_username  = "jimmy"
  msg_vpn_name = "test_vpn"
  client_profile_name = "jimmy"
  enabled = true
  depends_on = ["solace_msg_vpn.test_vpn"]
  depends_on = ["solace_msg_vpn.profile_jimmy"]
}

