package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)

// Main resource definition for Client-Profile entities
func resourceClientProfile() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema:        clientProfileSchema(),
		// Provider CRUD functions
		Create:        createClientProfileFunc,
		Read:          readClientProfileFunc,
		Update:        updateClientProfileFunc,
		Delete:        deleteClientProfileFunc,
	}
}

// List of supported configuration fields for Solace clientProfiles
// This method will only get called when initializing the clientProfile resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func clientProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		ALLOW_BRIDGE_CONNECTIONS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ALLOW_CUT_THROUGH_FORWARDING_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ALLOW_GUARANTEED_ENDPOINT_CREATE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ALLOW_GUARANTEED_MSG_RECEIVE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ALLOW_GUARANTEED_MSG_SEND_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ALLOW_TRANSACTED_SESSIONS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		API_QUEUE_MANAGEMENT_COPY_FROM_ON_CREATE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		API_TOPIC_ENDPOINT_MANAGEMENT_COPY_FROM_ON_CREATE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		CLIENT_PROFILE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ELIDING_DELAY: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		ELIDING_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		ELIDING_MAX_TOPIC_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_EGRESS_FLOW_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_ENDPOINT_COUNT_PER_CLIENT_USERNAME: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_INGRESS_FLOW_COUNT: {
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
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		QUEUE_CONTROL1_MAX_DEPTH: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_CONTROL1_MIN_MSG_BURST: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_DIRECT1_MAX_DEPTH: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_DIRECT1_MIN_MSG_BURST: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_DIRECT2_MAX_DEPTH: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_DIRECT2_MIN_MSG_BURST: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_DIRECT3_MAX_DEPTH: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_DIRECT3_MIN_MSG_BURST: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_GUARANTEED1_MAX_DEPTH: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		QUEUE_GUARANTEED1_MIN_MSG_BURST: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REJECT_MSG_TO_SENDER_ON_NO_SUBSCRIPTION_MATCH_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		REPLICATION_ALLOW_CLIENT_CONNECT_WHEN_STANDBY_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		SERVICE_SMF_MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_WEB_INACTIVE_TIMEOUT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_WEB_MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		SERVICE_WEB_MAX_PAYLOAD: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TCP_CONGESTION_WINDOW_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TCP_KEEPALIVE_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TCP_KEEPALIVE_IDLE_TIME: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TCP_KEEPALIVE_INTERVAL: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TCP_MAX_SEGMENT_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TCP_MAX_WINDOW_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

// Generated from a field-list, either Swagger schema or generated GO client code.
// Given a DAO instance, populate it with any fields found in the resource data.
func populateClientProfileFromResource(profile *semp_client.MsgVpnClientProfile, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(ALLOW_BRIDGE_CONNECTIONS_ENABLED)
	if b {
		profile.AllowBridgeConnectionsEnabled = v.(bool)
	}
	v,b = d.GetOk(ALLOW_CUT_THROUGH_FORWARDING_ENABLED)
	if b {
		profile.AllowCutThroughForwardingEnabled = v.(bool)
	}
	v,b = d.GetOk(ALLOW_GUARANTEED_ENDPOINT_CREATE_ENABLED)
	if b {
		profile.AllowGuaranteedEndpointCreateEnabled = v.(bool)
	}
	v,b = d.GetOk(ALLOW_GUARANTEED_MSG_RECEIVE_ENABLED)
	if b {
		profile.AllowGuaranteedMsgReceiveEnabled = v.(bool)
	}
	v,b = d.GetOk(ALLOW_GUARANTEED_MSG_SEND_ENABLED)
	if b {
		profile.AllowGuaranteedMsgSendEnabled = v.(bool)
	}
	v,b = d.GetOk(ALLOW_TRANSACTED_SESSIONS_ENABLED)
	if b {
		profile.AllowTransactedSessionsEnabled = v.(bool)
	}
	v,b = d.GetOk(API_QUEUE_MANAGEMENT_COPY_FROM_ON_CREATE_NAME)
	if b {
		profile.ApiQueueManagementCopyFromOnCreateName = v.(string)
	}
	v,b = d.GetOk(API_TOPIC_ENDPOINT_MANAGEMENT_COPY_FROM_ON_CREATE_NAME)
	if b {
		profile.ApiTopicEndpointManagementCopyFromOnCreateName = v.(string)
	}
	v,b = d.GetOk(CLIENT_PROFILE_NAME)
	if b {
		profile.ClientProfileName = v.(string)
	}
	v,b = d.GetOk(ELIDING_DELAY)
	if b {
		profile.ElidingDelay = int64(v.(int))
	}
	v,b = d.GetOk(ELIDING_ENABLED)
	if b {
		profile.ElidingEnabled = v.(bool)
	}
	v,b = d.GetOk(ELIDING_MAX_TOPIC_COUNT)
	if b {
		profile.ElidingMaxTopicCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME)
	if b {
		profile.MaxConnectionCountPerClientUsername = int64(v.(int))
	}
	v,b = d.GetOk(MAX_EGRESS_FLOW_COUNT)
	if b {
		profile.MaxEgressFlowCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_ENDPOINT_COUNT_PER_CLIENT_USERNAME)
	if b {
		profile.MaxEndpointCountPerClientUsername = int64(v.(int))
	}
	v,b = d.GetOk(MAX_INGRESS_FLOW_COUNT)
	if b {
		profile.MaxIngressFlowCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_SUBSCRIPTION_COUNT)
	if b {
		profile.MaxSubscriptionCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_TRANSACTED_SESSION_COUNT)
	if b {
		profile.MaxTransactedSessionCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_TRANSACTION_COUNT)
	if b {
		profile.MaxTransactionCount = int64(v.(int))
	}
	v,b = d.GetOk(MSG_VPN_NAME)
	if b {
		profile.MsgVpnName = v.(string)
	}
	v,b = d.GetOk(QUEUE_CONTROL1_MAX_DEPTH)
	if b {
		profile.QueueControl1MaxDepth = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_CONTROL1_MIN_MSG_BURST)
	if b {
		profile.QueueControl1MinMsgBurst = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_DIRECT1_MAX_DEPTH)
	if b {
		profile.QueueDirect1MaxDepth = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_DIRECT1_MIN_MSG_BURST)
	if b {
		profile.QueueDirect1MinMsgBurst = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_DIRECT2_MAX_DEPTH)
	if b {
		profile.QueueDirect2MaxDepth = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_DIRECT2_MIN_MSG_BURST)
	if b {
		profile.QueueDirect2MinMsgBurst = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_DIRECT3_MAX_DEPTH)
	if b {
		profile.QueueDirect3MaxDepth = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_DIRECT3_MIN_MSG_BURST)
	if b {
		profile.QueueDirect3MinMsgBurst = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_GUARANTEED1_MAX_DEPTH)
	if b {
		profile.QueueGuaranteed1MaxDepth = int32(v.(int))
	}
	v,b = d.GetOk(QUEUE_GUARANTEED1_MIN_MSG_BURST)
	if b {
		profile.QueueGuaranteed1MinMsgBurst = int32(v.(int))
	}
	v,b = d.GetOk(REJECT_MSG_TO_SENDER_ON_NO_SUBSCRIPTION_MATCH_ENABLED)
	if b {
		profile.RejectMsgToSenderOnNoSubscriptionMatchEnabled = v.(bool)
	}
	v,b = d.GetOk(REPLICATION_ALLOW_CLIENT_CONNECT_WHEN_STANDBY_ENABLED)
	if b {
		profile.ReplicationAllowClientConnectWhenStandbyEnabled = v.(bool)
	}
	v,b = d.GetOk(SERVICE_SMF_MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME)
	if b {
		profile.ServiceSmfMaxConnectionCountPerClientUsername = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_WEB_INACTIVE_TIMEOUT)
	if b {
		profile.ServiceWebInactiveTimeout = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_WEB_MAX_CONNECTION_COUNT_PER_CLIENT_USERNAME)
	if b {
		profile.ServiceWebMaxConnectionCountPerClientUsername = int64(v.(int))
	}
	v,b = d.GetOk(SERVICE_WEB_MAX_PAYLOAD)
	if b {
		profile.ServiceWebMaxPayload = int64(v.(int))
	}
	v,b = d.GetOk(TCP_CONGESTION_WINDOW_SIZE)
	if b {
		profile.TcpCongestionWindowSize = int64(v.(int))
	}
	v,b = d.GetOk(TCP_KEEPALIVE_COUNT)
	if b {
		profile.TcpKeepaliveCount = int64(v.(int))
	}
	v,b = d.GetOk(TCP_KEEPALIVE_IDLE_TIME)
	if b {
		profile.TcpKeepaliveIdleTime = int64(v.(int))
	}
	v,b = d.GetOk(TCP_KEEPALIVE_INTERVAL)
	if b {
		profile.TcpKeepaliveInterval = int64(v.(int))
	}
	v,b = d.GetOk(TCP_MAX_SEGMENT_SIZE)
	if b {
		profile.TcpMaxSegmentSize = int64(v.(int))
	}
	v,b = d.GetOk(TCP_MAX_WINDOW_SIZE)
	if b {
		profile.TcpMaxWindowSize = int64(v.(int))
	}
}


// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createClientProfileFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createClientProfileFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	profile := semp_client.MsgVpnClientProfile{
		ClientProfileName:  d.Get(CLIENT_PROFILE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateClientProfileFromResource(&profile, d)

	cp, _, err := semp_client.ClientProfileApi{
			Configuration: state.sempcfg,
		}.CreateMsgVpnClientProfile(profile.MsgVpnName, profile, nil)

	if err != nil {
		log.Println("ClientProfileApi.CreateMsgVpnClientProfile ERROR")
		return err
	}
	logSempMeta("client-profile create", cp.Meta)

	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + profile.MsgVpnName + "_clientprofile_" + profile.ClientProfileName )

	return nil
}

func readClientProfileFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	resp, _, err := semp_client.ClientProfileApi {
			Configuration:state.sempcfg,
		}.GetMsgVpnClientProfile(d.Get(MSG_VPN_NAME).(string), d.Get(CLIENT_PROFILE_NAME).(string), []string { "*" })

	if err != nil {
		log.Println("ClientProfileApi.GetMsgVpnClientProfile ERROR")
		return err
	}
	logSempMeta("client-profile read", resp.Meta)

	return nil
}

func updateClientProfileFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	profile := semp_client.MsgVpnClientProfile{
		ClientProfileName:  d.Get(CLIENT_PROFILE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateClientProfileFromResource(&profile, d)

	cp, _, err := semp_client.ClientProfileApi{
		Configuration: state.sempcfg,
	}.UpdateMsgVpnClientProfile(profile.MsgVpnName, profile.ClientProfileName, profile, nil)

	if err != nil {
		log.Println("ClientProfileApi.UpdateMsgVpnClientProfile ERROR")
		return err
	}
	logSempMeta("client-profile update", cp.Meta)

	return nil
}

func deleteClientProfileFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	profile := semp_client.MsgVpnClientProfile{
		ClientProfileName: d.Get(CLIENT_PROFILE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}

	cp, _, err := semp_client.ClientProfileApi{
		Configuration: state.sempcfg,
		}.DeleteMsgVpnClientProfile(profile.MsgVpnName, profile.ClientProfileName)

	if err != nil {
		log.Println("ClientProfileApi.DeleteMsgVpnClientProfile ERROR")
		return err
	}
	logSempMeta("client-profile delete", cp.Meta)

	return nil
}

