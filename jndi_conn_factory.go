package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)

// Main resource definition for Client-Username entities
func resourceJndiConnectionFactory() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema:        jndiConnectionFactorySchema(),
		// Provider CRUD functions
		Create:        createJndiConnectionFactoryFunc,
		Read:          readJndiConnectionFactoryFunc,
		Update:        updateJndiConnectionFactoryFunc,
		Delete:        deleteJndiConnectionFactoryFunc,
	}
}

// List of supported configuration fields for Solace jndiConnectionFactorys
// This method will only get called when initializing the jndiConnectionFactory resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func jndiConnectionFactorySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		CONNECTION_FACTORY_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		ALLOW_DUPLICATE_CLIENT_ID_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		CLIENT_DESCRIPTION: {
			Type:     schema.TypeString,
			Optional: true,
		},
		CLIENT_ID: {
			Type:     schema.TypeString,
			Optional: true,
		},
		DTO_RECEIVE_OVERRIDE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		DTO_RECEIVE_SUBSCRIBER_LOCAL_PRIORITY: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		DTO_RECEIVE_SUBSCRIBER_NETWORK_PRIORITY: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		DTO_SEND_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		DYNAMIC_ENDPOINT_CREATE_DURABLE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		DYNAMIC_ENDPOINT_RESPECT_TTL_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		GUARANTEED_RECEIVE_ACK_TIMEOUT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		GUARANTEED_RECEIVE_WINDOW_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		GUARANTEED_RECEIVE_WINDOW_SIZE_ACK_THRESHOLD: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		GUARANTEED_SEND_ACK_TIMEOUT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		GUARANTEED_SEND_WINDOW_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MESSAGING_DEFAULT_DELIVERY_MODE: {
			Type:     schema.TypeString,
			Optional: true,
		},
		MESSAGING_DEFAULT_DMQ_ELIGIBLE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		MESSAGING_DEFAULT_ELIDING_ELIGIBLE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		MESSAGING_JMSX_USER_ID_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		MESSAGING_TEXT_IN_XML_PAYLOAD_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		TRANSPORT_COMPRESSION_LEVEL: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_CONNECT_RETRY_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_CONNECT_RETRY_PER_HOST_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_CONNECT_TIMEOUT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_DIRECT_TRANSPORT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		TRANSPORT_KEEPALIVE_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_KEEPALIVE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		TRANSPORT_KEEPALIVE_INTERVAL: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_MSG_CALLBACK_ON_IO_THREAD_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		TRANSPORT_OPTIMIZE_DIRECT_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		TRANSPORT_PORT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_READ_TIMEOUT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_RECEIVE_BUFFER_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_RECONNECT_RETRY_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_RECONNECT_RETRY_WAIT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_SEND_BUFFER_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		TRANSPORT_TCP_NO_DELAY_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		XA_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},

	}
}

// Generated from a field-list, either Swagger schema or generated GO client code.
// Given a DAO instance, populate it with any fields found in the resource data.
func populateJndiConnectionFactoryFromResource(jcf *semp_client.MsgVpnJndiConnectionFactory, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(ALLOW_DUPLICATE_CLIENT_ID_ENABLED)
	if b {
		jcf.AllowDuplicateClientIdEnabled = v.(bool)
	}
	v,b = d.GetOk(CLIENT_DESCRIPTION)
	if b {
		jcf.ClientDescription = v.(string)
	}
	v,b = d.GetOk(CLIENT_ID)
	if b {
		jcf.ClientId = v.(string)
	}
	v,b = d.GetOk(CONNECTION_FACTORY_NAME)
	if b {
		jcf.ConnectionFactoryName = v.(string)
	}
	v,b = d.GetOk(DTO_RECEIVE_OVERRIDE_ENABLED)
	if b {
		jcf.DtoReceiveOverrideEnabled = v.(bool)
	}
	v,b = d.GetOk(DTO_RECEIVE_SUBSCRIBER_LOCAL_PRIORITY)
	if b {
		jcf.DtoReceiveSubscriberLocalPriority = int32(v.(int))
	}
	v,b = d.GetOk(DTO_RECEIVE_SUBSCRIBER_NETWORK_PRIORITY)
	if b {
		jcf.DtoReceiveSubscriberNetworkPriority = int32(v.(int))
	}
	v,b = d.GetOk(DTO_SEND_ENABLED)
	if b {
		jcf.DtoSendEnabled = v.(bool)
	}
	v,b = d.GetOk(DYNAMIC_ENDPOINT_CREATE_DURABLE_ENABLED)
	if b {
		jcf.DynamicEndpointCreateDurableEnabled = v.(bool)
	}
	v,b = d.GetOk(DYNAMIC_ENDPOINT_RESPECT_TTL_ENABLED)
	if b {
		jcf.DynamicEndpointRespectTtlEnabled = v.(bool)
	}
	v,b = d.GetOk(GUARANTEED_RECEIVE_ACK_TIMEOUT)
	if b {
		jcf.GuaranteedReceiveAckTimeout = int32(v.(int))
	}
	v,b = d.GetOk(GUARANTEED_RECEIVE_WINDOW_SIZE)
	if b {
		jcf.GuaranteedReceiveWindowSize = int32(v.(int))
	}
	v,b = d.GetOk(GUARANTEED_RECEIVE_WINDOW_SIZE_ACK_THRESHOLD)
	if b {
		jcf.GuaranteedReceiveWindowSizeAckThreshold = int32(v.(int))
	}
	v,b = d.GetOk(GUARANTEED_SEND_ACK_TIMEOUT)
	if b {
		jcf.GuaranteedSendAckTimeout = int32(v.(int))
	}
	v,b = d.GetOk(GUARANTEED_SEND_WINDOW_SIZE)
	if b {
		jcf.GuaranteedSendWindowSize = int32(v.(int))
	}
	v,b = d.GetOk(MESSAGING_DEFAULT_DELIVERY_MODE)
	if b {
		jcf.MessagingDefaultDeliveryMode = v.(string)
	}
	v,b = d.GetOk(MESSAGING_DEFAULT_DMQ_ELIGIBLE_ENABLED)
	if b {
		jcf.MessagingDefaultDmqEligibleEnabled = v.(bool)
	}
	v,b = d.GetOk(MESSAGING_DEFAULT_ELIDING_ELIGIBLE_ENABLED)
	if b {
		jcf.MessagingDefaultElidingEligibleEnabled = v.(bool)
	}
	v,b = d.GetOk(MESSAGING_JMSX_USER_ID_ENABLED)
	if b {
		jcf.MessagingJmsxUserIdEnabled = v.(bool)
	}
	v,b = d.GetOk(MESSAGING_TEXT_IN_XML_PAYLOAD_ENABLED)
	if b {
		jcf.MessagingTextInXmlPayloadEnabled = v.(bool)
	}
	v,b = d.GetOk(MSG_VPN_NAME)
	if b {
		jcf.MsgVpnName = v.(string)
	}
	v,b = d.GetOk(TRANSPORT_COMPRESSION_LEVEL)
	if b {
		jcf.TransportCompressionLevel = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_CONNECT_RETRY_COUNT)
	if b {
		jcf.TransportConnectRetryCount = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_CONNECT_RETRY_PER_HOST_COUNT)
	if b {
		jcf.TransportConnectRetryPerHostCount = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_CONNECT_TIMEOUT)
	if b {
		jcf.TransportConnectTimeout = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_DIRECT_TRANSPORT_ENABLED)
	if b {
		jcf.TransportDirectTransportEnabled = v.(bool)
	}
	v,b = d.GetOk(TRANSPORT_KEEPALIVE_COUNT)
	if b {
		jcf.TransportKeepaliveCount = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_KEEPALIVE_ENABLED)
	if b {
		jcf.TransportKeepaliveEnabled = v.(bool)
	}
	v,b = d.GetOk(TRANSPORT_KEEPALIVE_INTERVAL)
	if b {
		jcf.TransportKeepaliveInterval = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_MSG_CALLBACK_ON_IO_THREAD_ENABLED)
	if b {
		jcf.TransportMsgCallbackOnIoThreadEnabled = v.(bool)
	}
	v,b = d.GetOk(TRANSPORT_OPTIMIZE_DIRECT_ENABLED)
	if b {
		jcf.TransportOptimizeDirectEnabled = v.(bool)
	}
	v,b = d.GetOk(TRANSPORT_PORT)
	if b {
		jcf.TransportPort = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_READ_TIMEOUT)
	if b {
		jcf.TransportReadTimeout = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_RECEIVE_BUFFER_SIZE)
	if b {
		jcf.TransportReceiveBufferSize = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_RECONNECT_RETRY_COUNT)
	if b {
		jcf.TransportReconnectRetryCount = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_RECONNECT_RETRY_WAIT)
	if b {
		jcf.TransportReconnectRetryWait = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_SEND_BUFFER_SIZE)
	if b {
		jcf.TransportSendBufferSize = int32(v.(int))
	}
	v,b = d.GetOk(TRANSPORT_TCP_NO_DELAY_ENABLED)
	if b {
		jcf.TransportTcpNoDelayEnabled = v.(bool)
	}
	v,b = d.GetOk(XA_ENABLED)
	if b {
		jcf.XaEnabled = v.(bool)
	}

}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createJndiConnectionFactoryFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createJndiConnectionFactoryFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	jcf := semp_client.MsgVpnJndiConnectionFactory{
		ConnectionFactoryName:  d.Get(CONNECTION_FACTORY_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateJndiConnectionFactoryFromResource(&jcf, d)

	cu, _, err := semp_client.JndiApi {
		Configuration: state.sempcfg,
	}.CreateMsgVpnJndiConnectionFactory(jcf.MsgVpnName, jcf, nil)

	if err != nil {
		log.Println("JndiApi.CreateMsgVpnJndiConnectionFactory ERROR")
		return err
	}
	logSempMeta("JndiConnectionFactory create", cu.Meta)
	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + jcf.MsgVpnName + "_jndiConnectionFactory_" + jcf.ConnectionFactoryName )

	return nil
}

func readJndiConnectionFactoryFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	userresp, _, err := semp_client.JndiApi {
		Configuration:state.sempcfg,
	}.GetMsgVpnJndiConnectionFactory(d.Get(MSG_VPN_NAME).(string), d.Get(CONNECTION_FACTORY_NAME).(string), nil)

	if err != nil {
		log.Println("JndiApi.GetMsgVpnJndiConnectionFactory ERROR")
		return err
	}
	logSempMeta("JndiConnectionFactory read", userresp.Meta)

	return nil
}

func updateJndiConnectionFactoryFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	user := semp_client.MsgVpnJndiConnectionFactory{
		ConnectionFactoryName:  d.Get(CONNECTION_FACTORY_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateJndiConnectionFactoryFromResource(&user, d)

	cu, _, err := semp_client.JndiApi{
		Configuration: state.sempcfg,
	}.UpdateMsgVpnJndiConnectionFactory(user.MsgVpnName, user.ConnectionFactoryName, user, nil)

	if err != nil {
		log.Println("JndiApi.UpdateMsgVpnJndiConnectionFactory ERROR")
		return err
	}
	logSempMeta("JndiConnectionFactory update", cu.Meta)

	return nil
}

func deleteJndiConnectionFactoryFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	user := semp_client.MsgVpnJndiConnectionFactory{
		ConnectionFactoryName:  d.Get(CONNECTION_FACTORY_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}

	cu, _, err := semp_client.JndiApi {
		Configuration: state.sempcfg,
	}.DeleteMsgVpnJndiConnectionFactory(user.MsgVpnName, user.ConnectionFactoryName)

	if err != nil {
		log.Println("JndiApi.DeleteMsgVpnJndiConnectionFactory ERROR")
		return err
	}
	logSempMeta("JndiConnectionFactory delete", cu.Meta)

	return nil
}


