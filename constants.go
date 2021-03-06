package main

// Provider Fields
const ADMIN_USER   string = "admin_user"
const ADMIN_PASSWD string = "admin_password"
const HOST         string = "host"
const PORT         string = "port"

// Provider Resource Names
const RSRC_MSG_VPN               string = "solace_msg_vpn"
const RSRC_QUEUE                 string = "solace_queue"
const RSRC_TOPIC_ENDPOINT        string = "solace_topic_endpoint"
const RSRC_CLIENT_USERNAME       string = "solace_client_username"
const RSRC_CLIENT_PROFILE        string = "solace_client_profile"
const RSRC_ACL_PROFILE           string = "solace_acl_profile"
const RSRC_MSG_VPN_BRIDGE        string = "solace_msg_vpn_bridge"
const RSRC_VPN_BRIDGE_REMOTE_VPN string = "solace_msg_vpn_bridge_remote_vpn"
const RSRC_JNDI_CONN_FACTORY     string = "solace_jndi_conn_factory"
const RSRC_JNDI_QUEUE            string = "solace_jndi_queue"
const RSRC_JNDI_TOPIC            string = "solace_jndi_topic"


// Unique Msg-VPN Fields
const AUTHENTICATION_BASIC_ENABLED string = "authentication_basic_enabled"
const AUTHENTICATION_BASIC_PROFILE_NAME string = "authentication_basic_profile_name"
const AUTHENTICATION_BASIC_RADIUS_DOMAIN string = "authentication_basic_radius_domain"
const AUTHENTICATION_BASIC_TYPE string = "authentication_basic_type"
const AUTHENTICATION_CLIENT_CERT_ALLOW_API_PROVIDED_USERNAME_ENABLED string = "authentication_client_cert_allow_api_provided_username_enabled"
const AUTHENTICATION_CLIENT_CERT_ENABLED string = "authentication_client_cert_enabled"
const AUTHENTICATION_CLIENT_CERT_MAX_CHAIN_DEPTH string = "authentication_client_cert_max_chain_depth"
const AUTHENTICATION_CLIENT_CERT_VALIDATE_DATE_ENABLED string = "authentication_client_cert_validate_date_enabled"
const AUTHENTICATION_KERBEROS_ALLOW_API_PROVIDED_USERNAME_ENABLED string = "authentication_kerberos_allow_api_provided_username_enabled"
const AUTHENTICATION_KERBEROS_ENABLED string = "authentication_kerberos_enabled"
const AUTHORIZATION_LDAP_GROUP_MEMBERSHIP_ATTRIBUTE_NAME string = "authorization_ldap_group_membership_attribute_name"
const AUTHORIZATION_PROFILE_NAME string = "authorization_profile_name"
const AUTHORIZATION_TYPE string = "authorization_type"
const BRIDGING_TLS_SERVER_CERT_ENFORCE_TRUSTED_COMMON_NAME_ENABLED string = "bridging_tls_server_cert_enforce_trusted_common_name_enabled"
const BRIDGING_TLS_SERVER_CERT_MAX_CHAIN_DEPTH string = "bridging_tls_server_cert_max_chain_depth"
const BRIDGING_TLS_SERVER_CERT_VALIDATE_DATE_ENABLED string = "bridging_tls_server_cert_validate_date_enabled"
const DISTRIBUTED_CACHE_MANAGEMENT_ENABLED string = "distributed_cache_management_enabled"
const ENABLED string = "enabled"
const EVENT_LARGE_MSG_THRESHOLD string = "event_large_msg_threshold"
const EVENT_LOG_TAG string = "event_log_tag"
const EVENT_PUBLISH_CLIENT_ENABLED string = "event_publish_client_enabled"
const EVENT_PUBLISH_MSG_VPN_ENABLED string = "event_publish_msg_vpn_enabled"
const EVENT_PUBLISH_SUBSCRIPTION_MODE string = "event_publish_subscription_mode"
const EVENT_PUBLISH_TOPIC_FORMAT_MQTT_ENABLED string = "event_publish_topic_format_mqtt_enabled"
const EVENT_PUBLISH_TOPIC_FORMAT_SMF_ENABLED string = "event_publish_topic_format_smf_enabled"
const EXPORT_SUBSCRIPTIONS_ENABLED string = "export_subscriptions_enabled"
const JNDI_ENABLED string = "jndi_enabled"
const MAX_CONNECTION_COUNT string = "max_connection_count"
const MAX_EGRESS_FLOW_COUNT string = "max_egress_flow_count"
const MAX_ENDPOINT_COUNT string = "max_endpoint_count"
const MAX_INGRESS_FLOW_COUNT string = "max_ingress_flow_count"
const MAX_MSG_SPOOL_USAGE string = "max_msg_spool_usage"
const MAX_SUBSCRIPTION_COUNT string = "max_subscription_count"
const MAX_TRANSACTED_SESSION_COUNT string = "max_transacted_session_count"
const MAX_TRANSACTION_COUNT string = "max_transaction_count"
const MSG_VPN_NAME string = "msg_vpn_name"
const REPLICATION_ACK_PROPAGATION_INTERVAL_MSG_COUNT string = "replication_ack_propagation_interval_msg_count"
const REPLICATION_BRIDGE_AUTHENTICATION_BASIC_CLIENT_USERNAME string = "replication_bridge_authentication_basic_client_username"
const REPLICATION_BRIDGE_AUTHENTICATION_BASIC_PASSWORD string = "replication_bridge_authentication_basic_password"
const REPLICATION_BRIDGE_AUTHENTICATION_SCHEME string = "replication_bridge_authentication_scheme"
const REPLICATION_BRIDGE_COMPRESSED_DATA_ENABLED string = "replication_bridge_compressed_data_enabled"
const REPLICATION_BRIDGE_EGRESS_FLOW_WINDOW_SIZE string = "replication_bridge_egress_flow_window_size"
const REPLICATION_BRIDGE_RETRY_DELAY string = "replication_bridge_retry_delay"
const REPLICATION_BRIDGE_TLS_ENABLED string = "replication_bridge_tls_enabled"
const REPLICATION_BRIDGE_UNIDIRECTIONAL_CLIENT_PROFILE_NAME string = "replication_bridge_unidirectional_client_profile_name"
const REPLICATION_ENABLED string = "replication_enabled"
const REPLICATION_ENABLED_QUEUE_BEHAVIOR string = "replication_enabled_queue_behavior"
const REPLICATION_QUEUE_MAX_MSG_SPOOL_USAGE string = "replication_queue_max_msg_spool_usage"
const REPLICATION_QUEUE_REJECT_MSG_TO_SENDER_ON_DISCARD_ENABLED string = "replication_queue_reject_msg_to_sender_on_discard_enabled"
const REPLICATION_REJECT_MSG_WHEN_SYNC_INELIGIBLE_ENABLED string = "replication_reject_msg_when_sync_ineligible_enabled"
const REPLICATION_ROLE string = "replication_role"
const REPLICATION_TRANSACTION_MODE string = "replication_transaction_mode"
const REST_TLS_SERVER_CERT_ENFORCE_TRUSTED_COMMON_NAME_ENABLED string = "rest_tls_server_cert_enforce_trusted_common_name_enabled"
const REST_TLS_SERVER_CERT_MAX_CHAIN_DEPTH string = "rest_tls_server_cert_max_chain_depth"
const REST_TLS_SERVER_CERT_VALIDATE_DATE_ENABLED string = "rest_tls_server_cert_validate_date_enabled"
const SEMP_OVER_MSG_BUS_ADMIN_CLIENT_ENABLED string = "semp_over_msg_bus_admin_client_enabled"
const SEMP_OVER_MSG_BUS_ADMIN_DISTRIBUTED_CACHE_ENABLED string = "semp_over_msg_bus_admin_distributed_cache_enabled"
const SEMP_OVER_MSG_BUS_ADMIN_ENABLED string = "semp_over_msg_bus_admin_enabled"
const SEMP_OVER_MSG_BUS_ENABLED string = "semp_over_msg_bus_enabled"
const SEMP_OVER_MSG_BUS_SHOW_ENABLED string = "semp_over_msg_bus_show_enabled"
const SERVICE_MQTT_MAX_CONNECTION_COUNT string = "service_mqtt_max_connection_count"
const SERVICE_MQTT_PLAIN_TEXT_ENABLED string = "service_mqtt_plain_text_enabled"
const SERVICE_MQTT_PLAIN_TEXT_LISTEN_PORT string = "service_mqtt_plain_text_listen_port"
const SERVICE_MQTT_TLS_ENABLED string = "service_mqtt_tls_enabled"
const SERVICE_MQTT_TLS_LISTEN_PORT string = "service_mqtt_tls_listen_port"
const SERVICE_MQTT_TLS_WEB_SOCKET_ENABLED string = "service_mqtt_tls_web_socket_enabled"
const SERVICE_MQTT_TLS_WEB_SOCKET_LISTEN_PORT string = "service_mqtt_tls_web_socket_listen_port"
const SERVICE_MQTT_WEB_SOCKET_ENABLED string = "service_mqtt_web_socket_enabled"
const SERVICE_MQTT_WEB_SOCKET_LISTEN_PORT string = "service_mqtt_web_socket_listen_port"
const SERVICE_REST_INCOMING_MAX_CONNECTION_COUNT string = "service_rest_incoming_max_connection_count"
const SERVICE_REST_INCOMING_PLAIN_TEXT_ENABLED string = "service_rest_incoming_plain_text_enabled"
const SERVICE_REST_INCOMING_PLAIN_TEXT_LISTEN_PORT string = "service_rest_incoming_plain_text_listen_port"
const SERVICE_REST_INCOMING_TLS_ENABLED string = "service_rest_incoming_tls_enabled"
const SERVICE_REST_INCOMING_TLS_LISTEN_PORT string = "service_rest_incoming_tls_listen_port"
const SERVICE_REST_OUTGOING_MAX_CONNECTION_COUNT string = "service_rest_outgoing_max_connection_count"
const SERVICE_SMF_MAX_CONNECTION_COUNT string = "service_smf_max_connection_count"
const SERVICE_SMF_PLAIN_TEXT_ENABLED string = "service_smf_plain_text_enabled"
const SERVICE_SMF_TLS_ENABLED string = "service_smf_tls_enabled"
const SERVICE_WEB_MAX_CONNECTION_COUNT string = "service_web_max_connection_count"
const SERVICE_WEB_PLAIN_TEXT_ENABLED string = "service_web_plain_text_enabled"
const SERVICE_WEB_TLS_ENABLED string = "service_web_tls_enabled"
const TLS_ALLOW_DOWNGRADE_TO_PLAIN_TEXT_ENABLED string = "tls_allow_downgrade_to_plain_text_enabled"


// Unique Queue Fields
const QUEUE_NAME   string = "queue_name"
const ACCESS_TYPE  string = "access_type"
const PERMISSION   string = "permission"
const CONSUMER_ACK_PROPAGATION_ENABLED string = "consumer_ack_propagation_enabled"
const DEAD_MSG_QUEUE string = "dead_msg_queue"
const EGRESS_ENABLED string = "egress_enabled"
const INGRESS_ENABLED string = "ingress_enabled"
const MAX_BIND_COUNT string = "max_bind_count"
const MAX_DELIVERED_UNACKED_MSGS_PER_FLOW string = "max_delivered_unacked_msgs_per_flow"
const MAX_MSG_SIZE string = "max_msg_size"
const MAX_REDELIVERY_COUNT string = "max_redelivery_count"
const MAX_TTL string = "max_ttl"
const OWNER string = "owner"
const REJECT_LOW_PRIORITY_MSG_ENABLED string = "reject_low_priority_msg_enabled"
const REJECT_LOW_PRIORITY_MSG_LIMIT string = "reject_low_priority_msg_limit"
const REJECT_MSG_TO_SENDER_ON_DISCARD_BEHAVIOR string = "reject_msg_to_sender_on_discard_behavior"
const RESPECT_TTL_ENABLED string = "respect_ttl_enabled"
const TOPIC_SUBSCRIPTION_LIST string = "topic_subscription_list"


// Unique Topic-Endpoint Fields
const MAX_SPOOL_USAGE string = "max_spool_usage"
const TOPIC_ENDPOINT_NAME string = "topic_endpoint_name"


// Unique Client-Username Fields
const ACL_PROFILE_NAME string = "acl_profile_name"
const CLIENT_USERNAME string = "client_username"
const GUARANTEED_ENDPOINT_PERMISSION_OVERRIDE_ENABLED string = "guaranteed_endpoint_permission_override_enabled"
const PASSWORD string = "password"
const SUBSCRIPTION_MANAGER_ENABLED string = "subscription_manager_enabled"


// Unique ACL-Profile Fields
const CLIENT_CONNECT_DEFAULT_ACTION string = "client_connect_default_action"
const PUBLISH_TOPIC_DEFAULT_ACTION string = "publish_topic_default_action"
const SUBSCRIBE_TOPIC_DEFAULT_ACTION string = "subscribe_topic_default_action"
const ACL_CONNECT_EXCEPTION_LIST string = "acl_connect_exception_list"
const ACL_PUBLISH_EXCEPTION_LIST string = "acl_publish_exception_list"
const ACL_SUBSCRIBE_EXCEPTION_LIST string = "acl_subscribe_exception_list"


// Unique Client-Profile Fields
const CLIENT_PROFILE_NAME string = "client_profile_name"
const ALLOW_BRIDGE_CONNECTIONS_ENABLED string = "allow_bridge_connections_enabled"
const ALLOW_CUT_THROUGH_FORWARDING_ENABLED string = "allow_cut_through_forwarding_enabled"
const ALLOW_GUARANTEED_ENDPOINT_CREATE_ENABLED string = "allow_guaranteed_endpoint_create_enabled"
const ALLOW_GUARANTEED_MSG_RECEIVE_ENABLED string = "allow_guaranteed_msg_receive_enabled"
const ALLOW_GUARANTEED_MSG_SEND_ENABLED string = "allow_guaranteed_msg_send_enabled"
const ALLOW_TRANSACTED_SESSIONS_ENABLED string = "allow_transacted_sessions_enabled"
const API_QUEUE_MANAGEMENT_COPY_FROM_ON_CREATE_NAME string = "api_queue_management_copy_from_on_create_name"
const API_TOPIC_ENDPOINT_MANAGEMENT_COPY_FROM_ON_CREATE_NAME string = "api_topic_endpoint_management_copy_from_on_create_name"
const ELIDING_DELAY string = "eliding_delay"
const ELIDING_ENABLED string = "eliding_enabled"
const ELIDING_MAX_TOPIC_COUNT string = "eliding_max_topic_count"
const MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME string = "max_connection_count_per_client_username"
const MAX_ENDPOINT_COUNT_PER_CLIENT_USERNAME string = "max_endpoint_count_per_client_username"
const QUEUE_CONTROL1_MAX_DEPTH string = "queue_control1_max_depth"
const QUEUE_CONTROL1_MIN_MSG_BURST string = "queue_control1_min_msg_burst"
const QUEUE_DIRECT1_MAX_DEPTH string = "queue_direct1_max_depth"
const QUEUE_DIRECT1_MIN_MSG_BURST string = "queue_direct1_min_msg_burst"
const QUEUE_DIRECT2_MAX_DEPTH string = "queue_direct2_max_depth"
const QUEUE_DIRECT2_MIN_MSG_BURST string = "queue_direct2_min_msg_burst"
const QUEUE_DIRECT3_MAX_DEPTH string = "queue_direct3_max_depth"
const QUEUE_DIRECT3_MIN_MSG_BURST string = "queue_direct3_min_msg_burst"
const QUEUE_GUARANTEED1_MAX_DEPTH string = "queue_guaranteed1_max_depth"
const QUEUE_GUARANTEED1_MIN_MSG_BURST string = "queue_guaranteed1_min_msg_burst"
const REJECT_MSG_TO_SENDER_ON_NO_SUBSCRIPTION_MATCH_ENABLED string = "reject_msg_to_sender_on_no_subscription_match_enabled"
const REPLICATION_ALLOW_CLIENT_CONNECT_WHEN_STANDBY_ENABLED string = "replication_allow_client_connect_when_standby_enabled"
const SERVICE_SMF_MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME string = "service_smf_max_connection_count_per_client_username"
const SERVICE_WEB_INACTIVE_TIMEOUT string = "service_web_inactive_timeout"
const SERVICE_WEB_MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME string = "service_web_max_connection_count_per_client_username"
const SERVICE_WEB_MAX_PAYLOAD string = "service_web_max_payload"
const TCP_CONGESTION_WINDOW_SIZE string = "tcp_congestion_window_size"
const TCP_KEEPALIVE_COUNT string = "tcp_keepalive_count"
const TCP_KEEPALIVE_IDLE_TIME string = "tcp_keepalive_idle_time"
const TCP_KEEPALIVE_INTERVAL string = "tcp_keepalive_interval"
const TCP_MAX_SEGMENT_SIZE string = "tcp_max_segment_size"
const TCP_MAX_WINDOW_SIZE string = "tcp_max_window_size"

// Unique Msg-VPN Bridge Fields
const BRIDGE_NAME string = "bridge_name"
const BRIDGE_VIRTUAL_ROUTER string = "bridge_virtual_router"
const REMOTE_AUTHENTICATION_SCHEME string = "remote_authentication_scheme"
const REMOTE_CONNECTION_RETRY_COUNT string = "remote_connection_retry_count"
const REMOTE_CONNECTION_RETRY_DELAY string = "remote_connection_retry_delay"
const REMOTE_DELIVER_TO_ONE_PRIORITY string = "remote_deliver_to_one_priority"
const TLS_CIPHER_SUITE_LIST string = "tls_cipher_suite_list"

// Unique Msg-VPN Bridge Remote Msg-VPN Fields
const COMPRESSED_DATA_ENABLED string = "compressed_data_enabled"
const CONNECT_ORDER string = "connect_order"
const EGRESS_FLOW_WINDOW_SIZE string = "egress_flow_window_size"
const QUEUE_BINDING string = "queue_binding"
const REMOTE_MSG_VPN_INTERFACE string = "remote_msg_vpn_interface"
const REMOTE_MSG_VPN_LOCATION string = "remote_msg_vpn_location"
const REMOTE_MSG_VPN_NAME string = "remote_msg_vpn_name"
const TLS_ENABLED string = "tls_enabled"
const UNIDIRECTIONAL_CLIENT_PROFILE string = "unidirectional_client_profile"

// Unique JCF Fields
const ALLOW_DUPLICATE_CLIENT_ID_ENABLED string = "allow_duplicate_client_id_enabled"
const CLIENT_DESCRIPTION string = "client_description"
const CLIENT_ID string = "client_id"
const CONNECTION_FACTORY_NAME string = "connection_factory_name"
const DTO_RECEIVE_OVERRIDE_ENABLED string = "dto_receive_override_enabled"
const DTO_RECEIVE_SUBSCRIBER_LOCAL_PRIORITY string = "dto_receive_subscriber_local_priority"
const DTO_RECEIVE_SUBSCRIBER_NETWORK_PRIORITY string = "dto_receive_subscriber_network_priority"
const DTO_SEND_ENABLED string = "dto_send_enabled"
const DYNAMIC_ENDPOINT_CREATE_DURABLE_ENABLED string = "dynamic_endpoint_create_durable_enabled"
const DYNAMIC_ENDPOINT_RESPECT_TTL_ENABLED string = "dynamic_endpoint_respect_ttl_enabled"
const GUARANTEED_RECEIVE_ACK_TIMEOUT string = "guaranteed_receive_ack_timeout"
const GUARANTEED_RECEIVE_WINDOW_SIZE string = "guaranteed_receive_window_size"
const GUARANTEED_RECEIVE_WINDOW_SIZE_ACK_THRESHOLD string = "guaranteed_receive_window_size_ack_threshold"
const GUARANTEED_SEND_ACK_TIMEOUT string = "guaranteed_send_ack_timeout"
const GUARANTEED_SEND_WINDOW_SIZE string = "guaranteed_send_window_size"
const MESSAGING_DEFAULT_DELIVERY_MODE string = "messaging_default_delivery_mode"
const MESSAGING_DEFAULT_DMQ_ELIGIBLE_ENABLED string = "messaging_default_dmq_eligible_enabled"
const MESSAGING_DEFAULT_ELIDING_ELIGIBLE_ENABLED string = "messaging_default_eliding_eligible_enabled"
const MESSAGING_JMSX_USER_ID_ENABLED string = "messaging_jmsx_user_id_enabled"
const MESSAGING_TEXT_IN_XML_PAYLOAD_ENABLED string = "messaging_text_in_xml_payload_enabled"
const TRANSPORT_COMPRESSION_LEVEL string = "transport_compression_level"
const TRANSPORT_CONNECT_RETRY_COUNT string = "transport_connect_retry_count"
const TRANSPORT_CONNECT_RETRY_PER_HOST_COUNT string = "transport_connect_retry_per_host_count"
const TRANSPORT_CONNECT_TIMEOUT string = "transport_connect_timeout"
const TRANSPORT_DIRECT_TRANSPORT_ENABLED string = "transport_direct_transport_enabled"
const TRANSPORT_KEEPALIVE_COUNT string = "transport_keepalive_count"
const TRANSPORT_KEEPALIVE_ENABLED string = "transport_keepalive_enabled"
const TRANSPORT_KEEPALIVE_INTERVAL string = "transport_keepalive_interval"
const TRANSPORT_MSG_CALLBACK_ON_IO_THREAD_ENABLED string = "transport_msg_callback_on_io_thread_enabled"
const TRANSPORT_OPTIMIZE_DIRECT_ENABLED string = "transport_optimize_direct_enabled"
const TRANSPORT_PORT string = "transport_port"
const TRANSPORT_READ_TIMEOUT string = "transport_read_timeout"
const TRANSPORT_RECEIVE_BUFFER_SIZE string = "transport_receive_buffer_size"
const TRANSPORT_RECONNECT_RETRY_COUNT string = "transport_reconnect_retry_count"
const TRANSPORT_RECONNECT_RETRY_WAIT string = "transport_reconnect_retry_wait"
const TRANSPORT_SEND_BUFFER_SIZE string = "transport_send_buffer_size"
const TRANSPORT_TCP_NO_DELAY_ENABLED string = "transport_tcp_no_delay_enabled"
const XA_ENABLED string = "xa_enabled"

// Unique JNDI Queue Fields
const PHYSICAL_NAME string = "physical_name"

// Unique JNDI Topic Fields
const TOPIC_NAME string = "topic_name"
