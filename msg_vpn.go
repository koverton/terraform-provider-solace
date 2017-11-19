package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)


// Main resource definition for Msg-VPN entities
func resourceMsgVpn() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema: schemaMsgVpn(),
		// Provider CRUD functions
		Create: createMsgVpnFunc,
		Read:   readMsgVpnFunc,
		Update: updateMsgVpnFunc,
		Delete: deleteMsgVpnFunc,
	}
}

// List of supported configuration fields for Solace queues
// This method will only get called when initializing the VPN resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func schemaMsgVpn() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		AUTHENTICATION_BASIC_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		AUTHENTICATION_BASIC_PROFILE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		AUTHENTICATION_BASIC_RADIUS_DOMAIN: {
			Type:     schema.TypeString,
			Optional: true,
		},
		AUTHENTICATION_BASIC_TYPE: {
			Type:     schema.TypeString,
			Optional: true,
		},
		AUTHENTICATION_CLIENT_CERT_ALLOW_API_PROVIDED_USERNAME_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		AUTHENTICATION_CLIENT_CERT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		AUTHENTICATION_CLIENT_CERT_MAX_CHAIN_DEPTH: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		AUTHENTICATION_CLIENT_CERT_VALIDATE_DATE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		AUTHENTICATION_KERBEROS_ALLOW_API_PROVIDED_USERNAME_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		AUTHENTICATION_KERBEROS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		AUTHORIZATION_LDAP_GROUP_MEMBERSHIP_ATTRIBUTE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		AUTHORIZATION_PROFILE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		AUTHORIZATION_TYPE: {
			Type:     schema.TypeString,
			Optional: true,
		},
		BRIDGING_TLS_SERVER_CERT_ENFORCE_TRUSTED_COMMON_NAME_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		BRIDGING_TLS_SERVER_CERT_MAX_CHAIN_DEPTH: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		BRIDGING_TLS_SERVER_CERT_VALIDATE_DATE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		DISTRIBUTED_CACHE_MANAGEMENT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		EVENT_LARGE_MSG_THRESHOLD: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		EVENT_LOG_TAG: {
			Type:     schema.TypeString,
			Optional: true,
		},
		EVENT_PUBLISH_CLIENT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		EVENT_PUBLISH_MSG_VPN_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		EVENT_PUBLISH_SUBSCRIPTION_MODE: {
			Type:     schema.TypeString,
			Optional: true,
		},
		EVENT_PUBLISH_TOPIC_FORMAT_MQTT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		EVENT_PUBLISH_TOPIC_FORMAT_SMF_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		EXPORT_SUBSCRIPTIONS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		JNDI_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		MAX_CONNECTION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_EGRESS_FLOW_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_ENDPOINT_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_INGRESS_FLOW_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_MSG_SPOOL_USAGE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_SUBSCRIPTION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_TRANSACTED_SESSION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_TRANSACTION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REPLICATION_ACK_PROPAGATION_INTERVAL_MSG_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REPLICATION_BRIDGE_AUTHENTICATION_BASIC_CLIENT_USERNAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REPLICATION_BRIDGE_AUTHENTICATION_BASIC_PASSWORD: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REPLICATION_BRIDGE_AUTHENTICATION_SCHEME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REPLICATION_BRIDGE_COMPRESSED_DATA_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		REPLICATION_BRIDGE_EGRESS_FLOW_WINDOW_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REPLICATION_BRIDGE_RETRY_DELAY: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REPLICATION_BRIDGE_TLS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		REPLICATION_BRIDGE_UNIDIRECTIONAL_CLIENT_PROFILE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REPLICATION_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		REPLICATION_ENABLED_QUEUE_BEHAVIOR: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REPLICATION_QUEUE_MAX_MSG_SPOOL_USAGE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REPLICATION_QUEUE_REJECT_MSG_TO_SENDER_ON_DISCARD_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		REPLICATION_REJECT_MSG_WHEN_SYNC_INELIGIBLE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		REPLICATION_ROLE: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REPLICATION_TRANSACTION_MODE: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REST_TLS_SERVER_CERT_ENFORCE_TRUSTED_COMMON_NAME_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		REST_TLS_SERVER_CERT_MAX_CHAIN_DEPTH: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REST_TLS_SERVER_CERT_VALIDATE_DATE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SEMP_OVER_MSG_BUS_ADMIN_CLIENT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SEMP_OVER_MSG_BUS_ADMIN_DISTRIBUTED_CACHE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SEMP_OVER_MSG_BUS_ADMIN_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SEMP_OVER_MSG_BUS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SEMP_OVER_MSG_BUS_SHOW_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_MQTT_MAX_CONNECTION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_MQTT_PLAIN_TEXT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_MQTT_PLAIN_TEXT_LISTEN_PORT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_MQTT_TLS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_MQTT_TLS_LISTEN_PORT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_MQTT_TLS_WEB_SOCKET_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_MQTT_TLS_WEB_SOCKET_LISTEN_PORT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_MQTT_WEB_SOCKET_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_MQTT_WEB_SOCKET_LISTEN_PORT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_REST_INCOMING_MAX_CONNECTION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_REST_INCOMING_PLAIN_TEXT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_REST_INCOMING_PLAIN_TEXT_LISTEN_PORT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_REST_INCOMING_TLS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_REST_INCOMING_TLS_LISTEN_PORT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_REST_OUTGOING_MAX_CONNECTION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_SMF_MAX_CONNECTION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_SMF_PLAIN_TEXT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_SMF_TLS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_WEB_MAX_CONNECTION_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_WEB_PLAIN_TEXT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_WEB_TLS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		TLS_ALLOW_DOWNGRADE_TO_PLAIN_TEXT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func populateMsgVpnFromResource(msgVpn *semp_client.MsgVpn, d *schema.ResourceData) {
	v,b := d.GetOk(AUTHENTICATION_BASIC_ENABLED)
	if b {
		msgVpn.AuthenticationBasicEnabled = v.(bool)
	}
	v,b = d.GetOk(AUTHENTICATION_BASIC_PROFILE_NAME)
	if b {
		msgVpn.AuthenticationBasicProfileName = v.(string)
	}
	v,b = d.GetOk(AUTHENTICATION_BASIC_RADIUS_DOMAIN)
	if b {
		msgVpn.AuthenticationBasicRadiusDomain = v.(string)
	}
	v,b = d.GetOk(AUTHENTICATION_BASIC_TYPE)
	if b {
		msgVpn.AuthenticationBasicType = v.(string)
	}
	v,b = d.GetOk(AUTHENTICATION_CLIENT_CERT_ALLOW_API_PROVIDED_USERNAME_ENABLED)
	if b {
		msgVpn.AuthenticationClientCertAllowApiProvidedUsernameEnabled = v.(bool)
	}
	v,b = d.GetOk(AUTHENTICATION_CLIENT_CERT_ENABLED)
	if b {
		msgVpn.AuthenticationClientCertEnabled = v.(bool)
	}
	v,b = d.GetOk(AUTHENTICATION_CLIENT_CERT_MAX_CHAIN_DEPTH)
	if b {
		msgVpn.AuthenticationClientCertMaxChainDepth = int64(v.(int))
	}
	v,b = d.GetOk(AUTHENTICATION_CLIENT_CERT_VALIDATE_DATE_ENABLED)
	if b {
		msgVpn.AuthenticationClientCertValidateDateEnabled = v.(bool)
	}
	v,b = d.GetOk(AUTHENTICATION_KERBEROS_ALLOW_API_PROVIDED_USERNAME_ENABLED)
	if b {
		msgVpn.AuthenticationKerberosAllowApiProvidedUsernameEnabled = v.(bool)
	}
	v,b = d.GetOk(AUTHENTICATION_KERBEROS_ENABLED)
	if b {
		msgVpn.AuthenticationKerberosEnabled = v.(bool)
	}
	v,b = d.GetOk(AUTHORIZATION_LDAP_GROUP_MEMBERSHIP_ATTRIBUTE_NAME)
	if b {
		msgVpn.AuthorizationLdapGroupMembershipAttributeName = v.(string)
	}
	v,b = d.GetOk(AUTHORIZATION_PROFILE_NAME)
	if b {
		msgVpn.AuthorizationProfileName = v.(string)
	}
	v,b = d.GetOk(AUTHORIZATION_TYPE)
	if b {
		msgVpn.AuthorizationType = v.(string)
	}
	v,b = d.GetOk(BRIDGING_TLS_SERVER_CERT_ENFORCE_TRUSTED_COMMON_NAME_ENABLED)
	if b {
		msgVpn.BridgingTlsServerCertEnforceTrustedCommonNameEnabled = v.(bool)
	}
	v,b = d.GetOk(BRIDGING_TLS_SERVER_CERT_MAX_CHAIN_DEPTH)
	if b {
		msgVpn.BridgingTlsServerCertMaxChainDepth = int64(v.(int))
	}
	v,b = d.GetOk(BRIDGING_TLS_SERVER_CERT_VALIDATE_DATE_ENABLED)
	if b {
		msgVpn.BridgingTlsServerCertValidateDateEnabled = v.(bool)
	}
	v,b = d.GetOk(DISTRIBUTED_CACHE_MANAGEMENT_ENABLED)
	if b {
		msgVpn.DistributedCacheManagementEnabled = v.(bool)
	}
	v,b = d.GetOk(ENABLED)
	if b {
		msgVpn.Enabled = v.(bool)
	}
	v,b = d.GetOk(EVENT_LARGE_MSG_THRESHOLD)
	if b {
		msgVpn.EventLargeMsgThreshold = int64(v.(int))
	}
	v,b = d.GetOk(EVENT_LOG_TAG)
	if b {
		msgVpn.EventLogTag = v.(string)
	}
	v,b = d.GetOk(EVENT_PUBLISH_CLIENT_ENABLED)
	if b {
		msgVpn.EventPublishClientEnabled = v.(bool)
	}
	v,b = d.GetOk(EVENT_PUBLISH_MSG_VPN_ENABLED)
	if b {
		msgVpn.EventPublishMsgVpnEnabled = v.(bool)
	}
	v,b = d.GetOk(EVENT_PUBLISH_SUBSCRIPTION_MODE)
	if b {
		msgVpn.EventPublishSubscriptionMode = v.(string)
	}
	v,b = d.GetOk(EVENT_PUBLISH_TOPIC_FORMAT_MQTT_ENABLED)
	if b {
		msgVpn.EventPublishTopicFormatMqttEnabled = v.(bool)
	}
	v,b = d.GetOk(EVENT_PUBLISH_TOPIC_FORMAT_SMF_ENABLED)
	if b {
		msgVpn.EventPublishTopicFormatSmfEnabled = v.(bool)
	}
	v,b = d.GetOk(EXPORT_SUBSCRIPTIONS_ENABLED)
	if b {
		msgVpn.ExportSubscriptionsEnabled = v.(bool)
	}
	v,b = d.GetOk(JNDI_ENABLED)
	if b {
		msgVpn.JndiEnabled = v.(bool)
	}
	v,b = d.GetOk(MAX_CONNECTION_COUNT)
	if b {
		msgVpn.MaxConnectionCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_EGRESS_FLOW_COUNT)
	if b {
		msgVpn.MaxEgressFlowCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_ENDPOINT_COUNT)
	if b {
		msgVpn.MaxEndpointCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_INGRESS_FLOW_COUNT)
	if b {
		msgVpn.MaxIngressFlowCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_MSG_SPOOL_USAGE)
	if b {
		msgVpn.MaxMsgSpoolUsage = int64(v.(int))
	}
	v,b = d.GetOk(MAX_SUBSCRIPTION_COUNT)
	if b {
		msgVpn.MaxSubscriptionCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_TRANSACTED_SESSION_COUNT)
	if b {
		msgVpn.MaxTransactedSessionCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_TRANSACTION_COUNT)
	if b {
		msgVpn.MaxTransactionCount = int64(v.(int))
	}
	v,b = d.GetOk(REPLICATION_ACK_PROPAGATION_INTERVAL_MSG_COUNT)
	if b {
		msgVpn.ReplicationAckPropagationIntervalMsgCount = int64(v.(int))
	}
	v,b = d.GetOk(REPLICATION_BRIDGE_AUTHENTICATION_BASIC_CLIENT_USERNAME)
	if b {
		msgVpn.ReplicationBridgeAuthenticationBasicClientUsername = v.(string)
	}
	v,b = d.GetOk(REPLICATION_BRIDGE_AUTHENTICATION_BASIC_PASSWORD)
	if b {
		msgVpn.ReplicationBridgeAuthenticationBasicPassword = v.(string)
	}
	v,b = d.GetOk(REPLICATION_BRIDGE_AUTHENTICATION_SCHEME)
	if b {
		msgVpn.ReplicationBridgeAuthenticationScheme = v.(string)
	}
	v,b = d.GetOk(REPLICATION_BRIDGE_COMPRESSED_DATA_ENABLED)
	if b {
		msgVpn.ReplicationBridgeCompressedDataEnabled = v.(bool)
	}
	v,b = d.GetOk(REPLICATION_BRIDGE_EGRESS_FLOW_WINDOW_SIZE)
	if b {
		msgVpn.ReplicationBridgeEgressFlowWindowSize = int64(v.(int))
	}
	v,b = d.GetOk(REPLICATION_BRIDGE_RETRY_DELAY)
	if b {
		msgVpn.ReplicationBridgeRetryDelay = int64(v.(int))
	}
	v,b = d.GetOk(REPLICATION_BRIDGE_TLS_ENABLED)
	if b {
		msgVpn.ReplicationBridgeTlsEnabled = v.(bool)
	}
	v,b = d.GetOk(REPLICATION_BRIDGE_UNIDIRECTIONAL_CLIENT_PROFILE_NAME)
	if b {
		msgVpn.ReplicationBridgeUnidirectionalClientProfileName = v.(string)
	}
	v,b = d.GetOk(REPLICATION_ENABLED)
	if b {
		msgVpn.ReplicationEnabled = v.(bool)
	}
	v,b = d.GetOk(REPLICATION_ENABLED_QUEUE_BEHAVIOR)
	if b {
		msgVpn.ReplicationEnabledQueueBehavior = v.(string)
	}
	v,b = d.GetOk(REPLICATION_QUEUE_MAX_MSG_SPOOL_USAGE)
	if b {
		msgVpn.ReplicationQueueMaxMsgSpoolUsage = int64(v.(int))
	}
	v,b = d.GetOk(REPLICATION_QUEUE_REJECT_MSG_TO_SENDER_ON_DISCARD_ENABLED)
	if b {
		msgVpn.ReplicationQueueRejectMsgToSenderOnDiscardEnabled = v.(bool)
	}
	v,b = d.GetOk(REPLICATION_REJECT_MSG_WHEN_SYNC_INELIGIBLE_ENABLED)
	if b {
		msgVpn.ReplicationRejectMsgWhenSyncIneligibleEnabled = v.(bool)
	}
	v,b = d.GetOk(REPLICATION_ROLE)
	if b {
		msgVpn.ReplicationRole = v.(string)
	}
	v,b = d.GetOk(REPLICATION_TRANSACTION_MODE)
	if b {
		msgVpn.ReplicationTransactionMode = v.(string)
	}
	v,b = d.GetOk(REST_TLS_SERVER_CERT_ENFORCE_TRUSTED_COMMON_NAME_ENABLED)
	if b {
		msgVpn.RestTlsServerCertEnforceTrustedCommonNameEnabled = v.(bool)
	}
	v,b = d.GetOk(REST_TLS_SERVER_CERT_MAX_CHAIN_DEPTH)
	if b {
		msgVpn.RestTlsServerCertMaxChainDepth = int64(v.(int))
	}
	v,b = d.GetOk(REST_TLS_SERVER_CERT_VALIDATE_DATE_ENABLED)
	if b {
		msgVpn.RestTlsServerCertValidateDateEnabled = v.(bool)
	}
	v,b = d.GetOk(SEMP_OVER_MSG_BUS_ADMIN_CLIENT_ENABLED)
	if b {
		msgVpn.SempOverMsgBusAdminClientEnabled = v.(bool)
	}
	v,b = d.GetOk(SEMP_OVER_MSG_BUS_ADMIN_DISTRIBUTED_CACHE_ENABLED)
	if b {
		msgVpn.SempOverMsgBusAdminDistributedCacheEnabled = v.(bool)
	}
	v,b = d.GetOk(SEMP_OVER_MSG_BUS_ADMIN_ENABLED)
	if b {
		msgVpn.SempOverMsgBusAdminEnabled = v.(bool)
	}
	v,b = d.GetOk(SEMP_OVER_MSG_BUS_ENABLED)
	if b {
		msgVpn.SempOverMsgBusEnabled = v.(bool)
	}
	v,b = d.GetOk(SEMP_OVER_MSG_BUS_SHOW_ENABLED)
	if b {
		msgVpn.SempOverMsgBusShowEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_MQTT_MAX_CONNECTION_COUNT)
	if b {
		msgVpn.ServiceMqttMaxConnectionCount = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_MQTT_PLAIN_TEXT_ENABLED)
	if b {
		msgVpn.ServiceMqttPlainTextEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_MQTT_PLAIN_TEXT_LISTEN_PORT)
	if b {
		msgVpn.ServiceMqttPlainTextListenPort = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_MQTT_TLS_ENABLED)
	if b {
		msgVpn.ServiceMqttTlsEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_MQTT_TLS_LISTEN_PORT)
	if b {
		msgVpn.ServiceMqttTlsListenPort = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_MQTT_TLS_WEB_SOCKET_ENABLED)
	if b {
		msgVpn.ServiceMqttTlsWebSocketEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_MQTT_TLS_WEB_SOCKET_LISTEN_PORT)
	if b {
		msgVpn.ServiceMqttTlsWebSocketListenPort = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_MQTT_WEB_SOCKET_ENABLED)
	if b {
		msgVpn.ServiceMqttWebSocketEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_MQTT_WEB_SOCKET_LISTEN_PORT)
	if b {
		msgVpn.ServiceMqttWebSocketListenPort = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_REST_INCOMING_MAX_CONNECTION_COUNT)
	if b {
		msgVpn.ServiceRestIncomingMaxConnectionCount = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_REST_INCOMING_PLAIN_TEXT_ENABLED)
	if b {
		msgVpn.ServiceRestIncomingPlainTextEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_REST_INCOMING_PLAIN_TEXT_LISTEN_PORT)
	if b {
		msgVpn.ServiceRestIncomingPlainTextListenPort = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_REST_INCOMING_TLS_ENABLED)
	if b {
		msgVpn.ServiceRestIncomingTlsEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_REST_INCOMING_TLS_LISTEN_PORT)
	if b {
		msgVpn.ServiceRestIncomingTlsListenPort = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_REST_OUTGOING_MAX_CONNECTION_COUNT)
	if b {
		msgVpn.ServiceRestOutgoingMaxConnectionCount = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_SMF_MAX_CONNECTION_COUNT)
	if b {
		msgVpn.ServiceSmfMaxConnectionCount = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_SMF_PLAIN_TEXT_ENABLED)
	if b {
		msgVpn.ServiceSmfPlainTextEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_SMF_TLS_ENABLED)
	if b {
		msgVpn.ServiceSmfTlsEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_WEB_MAX_CONNECTION_COUNT)
	if b {
		msgVpn.ServiceWebMaxConnectionCount = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_WEB_PLAIN_TEXT_ENABLED)
	if b {
		msgVpn.ServiceWebPlainTextEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_WEB_TLS_ENABLED)
	if b {
		msgVpn.ServiceWebTlsEnabled = v.(bool)
	}
	v,b = d.GetOk(TLS_ALLOW_DOWNGRADE_TO_PLAIN_TEXT_ENABLED)
	if b {
		msgVpn.TlsAllowDowngradeToPlainTextEnabled = v.(bool)
	}
}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createMsgVpnFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createMsgVpnFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	msgVpn := semp_client.MsgVpn {
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateMsgVpnFromResource(&msgVpn, d)

	vpn, _, err := semp_client.MsgVpnApi{
		Configuration: state.sempcfg,
	}.CreateMsgVpn(msgVpn, nil)

	if err != nil {
		log.Println("MsgVpnApi.CreateMsgVpn ERROR")
		return err
	}
	logSempMeta("Msg-VPN create", vpn.Meta)
	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + msgVpn.MsgVpnName )

	return nil
}

func readMsgVpnFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	resp, _, err := semp_client.MsgVpnApi {
		Configuration:state.sempcfg,
	}.GetMsgVpn(d.Get(MSG_VPN_NAME).(string), nil)

	if err != nil {
		log.Println("MsgVpnApi.GetMsgVpn ERROR")
		return err
	}
	logSempMeta("Msg-VPN read", resp.Meta)

	return nil
}

func updateMsgVpnFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	msgVpn := semp_client.MsgVpn {
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateMsgVpnFromResource(&msgVpn, d)

	vpn, _, err := semp_client.MsgVpnApi{
		Configuration: state.sempcfg,
	}.UpdateMsgVpn(msgVpn.MsgVpnName, msgVpn, nil)

	if err != nil {
		log.Println("MsgVpnApi.UpdateMsgVpn ERROR")
		return err
	}
	logSempMeta("Msg-VPN update", vpn.Meta)

	return nil
}

func deleteMsgVpnFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	vpn, _, err := semp_client.MsgVpnApi{
		Configuration: state.sempcfg,
	}.DeleteMsgVpn( d.Get(MSG_VPN_NAME).(string) )

	if err != nil {
		log.Println("MsgVpnApi.DeleteMsgVpn ERROR")
		return err
	}
	logSempMeta("Msg-VPN delete", vpn.Meta)

	return nil
}

