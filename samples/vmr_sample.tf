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


