provider "solace" {
  host           = "192.168.56.201"
  port           = 8080
  admin_user     = "admin"
  admin_password = "admin"
}


# - + - + - + - + - + - + - + - #
# RED_VPN RESOURCE DEFINITIONS  #
# - + - + - + - + - + - + - + - #
resource "solace_msg_vpn" "red_vpn" {
  msg_vpn_name = "red_vpn"
  max_msg_spool_usage = 1000
  authentication_basic_enabled = true
  authentication_basic_type = "none"
  enabled = true
}
resource "solace_queue" "to_blue_bridge_queue" {
  queue_name  = "bridge_queue"
  msg_vpn_name = "red_vpn"
  access_type = "exclusive"
  max_msg_spool_usage = 1000
  permission = "delete"
  egress_enabled = true
  ingress_enabled = true
  topic_subscription_list = "red/topic/1,red/topic/2"
  depends_on = ["solace_msg_vpn.red_vpn"]
}
resource "solace_client_profile"  "red_bridge_profile" {
  client_profile_name = "bridge_profile"
  msg_vpn_name = "red_vpn"
  max_connection_count_per_client_username = 2
  allow_bridge_connections_enabled = true
  allow_guaranteed_endpoint_create_enabled = true
  allow_guaranteed_msg_send_enabled = true
  allow_guaranteed_msg_receive_enabled = true
  depends_on = ["solace_msg_vpn.red_vpn"]
}
resource "solace_client_username" "red_bridge_user" {
  client_username  = "bridge_user"
  msg_vpn_name = "red_vpn"
  client_profile_name = "bridge_profile"
  enabled = true
  depends_on = ["solace_client_profile.red_bridge_profile"]
}


# - + - + - + - + - + - + - + - #
# BLUE_VPN RESOURCE DEFINITIONS #
# - + - + - + - + - + - + - + - #
resource "solace_msg_vpn" "blue_vpn" {
  msg_vpn_name = "blue_vpn"
  max_msg_spool_usage = 1000
  authentication_basic_enabled = true
  authentication_basic_type = "none"
  enabled = true
}
resource "solace_queue" "to_red_bridge_queue" {
  queue_name  = "bridge_queue"
  msg_vpn_name = "blue_vpn"
  access_type = "exclusive"
  max_msg_spool_usage = 1000
  permission = "delete"
  egress_enabled = true
  ingress_enabled = true
  topic_subscription_list = "blue/topic/1,blue/topic/2"
  depends_on = ["solace_msg_vpn.blue_vpn"]
}
resource "solace_client_profile"  "blue_bridge_profile" {
  client_profile_name = "bridge_profile"
  msg_vpn_name = "blue_vpn"
  max_connection_count_per_client_username = 2
  allow_bridge_connections_enabled = true
  allow_guaranteed_endpoint_create_enabled = true
  allow_guaranteed_msg_send_enabled = true
  allow_guaranteed_msg_receive_enabled = true
  depends_on = ["solace_msg_vpn.blue_vpn"]
}
resource "solace_client_username" "blue_bridge_user" {
  client_username  = "bridge_user"
  msg_vpn_name = "blue_vpn"
  client_profile_name = "bridge_profile"
  enabled = true
  depends_on = ["solace_client_profile.blue_bridge_profile"]
}




# - + - + - + - + - + - + - + - + - + - + - #
# BIDIRECTIONAL BRIDGE RESOURCE DEFINITIONS #
# - + - + - + - + - + - + - + - + - + - + - #

## blue_vpn ==> red_vpn ##
resource "solace_msg_vpn_bridge" "blue_red_bi_bridge" {
  msg_vpn_name = "red_vpn"
  bridge_name = "blue_red_bi_bridge"
  bridge_virtual_router = "primary"
  enabled = true
  depends_on = ["solace_client_username.red_bridge_user"]
}
resource "solace_msg_vpn_bridge_remote_vpn" "blue_side" {
  msg_vpn_name = "red_vpn"
  bridge_name = "blue_red_bi_bridge"
  bridge_virtual_router = "primary"
  remote_msg_vpn_name = "blue_vpn"
  remote_msg_vpn_location = "127.0.0.1:55555"
  remote_msg_vpn_interface = "intf0"
  client_username = "bridge_user"
  password = "test"
  queue_binding = "bridge_queue"
  enabled = true
  depends_on = ["solace_msg_vpn_bridge.blue_red_bi_bridge"]
}

## red_vpn ==> blue_vpn ##
resource "solace_msg_vpn_bridge" "red_blue_bi_bridge" {
  msg_vpn_name = "blue_vpn"
  bridge_name = "red_blue_bi_bridge"
  bridge_virtual_router = "primary"
  enabled = true
  depends_on = ["solace_msg_vpn.blue_vpn"]
}
resource "solace_msg_vpn_bridge_remote_vpn" "red_side" {
  msg_vpn_name = "blue_vpn"
  bridge_name = "red_blue_bi_bridge"
  bridge_virtual_router = "primary"
  remote_msg_vpn_name = "red_vpn"
  remote_msg_vpn_location = "127.0.0.1:55555"
  remote_msg_vpn_interface = "intf0"
  client_username = "bridge_user"
  password = "test"
  queue_binding = "bridge_queue"
  enabled = true
  depends_on = ["solace_msg_vpn_bridge.red_blue_bi_bridge"]
}
