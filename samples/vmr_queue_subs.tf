provider "solace" {
  host           = "192.168.56.201"
  port           = 8080
  admin_user     = "admin"
  admin_password = "admin"
}

resource "solace_msg_vpn" "red_vpn" {
  msg_vpn_name = "red_vpn"
  max_msg_spool_usage = 1000
  authentication_basic_enabled = true
  authentication_basic_type = "none"
  enabled = true
}


resource "solace_queue" "jimmy" {
  queue_name  = "jimmy"
  msg_vpn_name = "red_vpn"
  access_type = "exclusive"
  max_msg_spool_usage = 1000
  permission = "delete"
  topic_subscription_list = "hello,all,you,happy,topics"
}
