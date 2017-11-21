package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)


// Main resource definition for VPN-Bridge entities
func resourceMsgVpnBridge() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema: schemaMsgVpnBridge(),
		// Provider CRUD functions
		Create: createMsgVpnBridgeFunc,
		Read:   readMsgVpnBridgeFunc,
		Update: updateMsgVpnBridgeFunc,
		Delete: deleteMsgVpnBridgeFunc,
	}
}

// List of supported configuration fields for Solace msgVpnBridges
// This method will only get called when initializing the msgVpnBridge resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func schemaMsgVpnBridge() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		BRIDGE_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		BRIDGE_VIRTUAL_ROUTER: {
			Type:     schema.TypeString,
			Required: true,
		},
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		MAX_TTL: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REMOTE_AUTHENTICATION_SCHEME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REMOTE_CONNECTION_RETRY_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REMOTE_CONNECTION_RETRY_DELAY: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REMOTE_DELIVER_TO_ONE_PRIORITY: {
			Type:     schema.TypeString,
			Optional: true,
		},
		TLS_CIPHER_SUITE_LIST: {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Generated from a field-list, either Swagger schema or generated GO client code.
// Given a DAO instance, populate it with any fields found in the resource data.
func populateMsgVpnBridgeFromResource(bridge *semp_client.MsgVpnBridge, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(BRIDGE_NAME)
	if b {
		bridge.BridgeName = v.(string)
	}
	v,b = d.GetOk(BRIDGE_VIRTUAL_ROUTER)
	if b {
		bridge.BridgeVirtualRouter = v.(string)
	}
	v,b = d.GetOk(MSG_VPN_NAME)
	if b {
		bridge.MsgVpnName = v.(string)
	}
	v,b = d.GetOk(ENABLED)
	if b {
		bridge.Enabled = v.(bool)
	}
	v,b = d.GetOk(MAX_TTL)
	if b {
		bridge.MaxTtl = v.(int64)
	}
	v,b = d.GetOk(REMOTE_AUTHENTICATION_SCHEME)
	if b {
		bridge.RemoteAuthenticationScheme = v.(string)
	}
	v,b = d.GetOk(REMOTE_CONNECTION_RETRY_COUNT)
	if b {
		bridge.RemoteConnectionRetryCount = v.(int64)
	}
	v,b = d.GetOk(REMOTE_CONNECTION_RETRY_DELAY)
	if b {
		bridge.RemoteConnectionRetryDelay = v.(int64)
	}
	v,b = d.GetOk(REMOTE_DELIVER_TO_ONE_PRIORITY)
	if b {
		bridge.RemoteDeliverToOnePriority = v.(string)
	}
	v,b = d.GetOk(TLS_CIPHER_SUITE_LIST)
	if b {
		bridge.TlsCipherSuiteList = v.(string)
	}
}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createMsgVpnBridgeFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createMsgVpnBridgeFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	bridge := semp_client.MsgVpnBridge{
		BridgeName:  d.Get(BRIDGE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
		BridgeVirtualRouter: d.Get(BRIDGE_VIRTUAL_ROUTER).(string),
	}
	populateMsgVpnBridgeFromResource(&bridge, d)

	b, _, err := semp_client.BridgeApi {
			Configuration: state.sempcfg,
		}.CreateMsgVpnBridge(bridge.MsgVpnName, bridge, nil)

	if err != nil {
		log.Println("BridgeApi.CreateMsgVpnBridge ERROR")
		return err
	}
	logSempMeta("VPN-Bridge create", b.Meta)

	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + bridge.MsgVpnName + "_vpnbridge_" + bridge.BridgeName )

	return nil
}

func readMsgVpnBridgeFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	resp, _, err := semp_client.BridgeApi {
			Configuration:state.sempcfg,
		}.GetMsgVpnBridge(d.Get(MSG_VPN_NAME).(string), d.Get(BRIDGE_NAME).(string), d.Get(BRIDGE_VIRTUAL_ROUTER).(string), nil)

	if err != nil {
		log.Println("BridgeApi.GetMsgVpnBridge ERROR")
		return err
	}

	logSempMeta("VPN-Bridge read", resp.Meta)

	return nil
}

func updateMsgVpnBridgeFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	bridge := semp_client.MsgVpnBridge{
		BridgeName:  d.Get(BRIDGE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
		BridgeVirtualRouter: d.Get(BRIDGE_VIRTUAL_ROUTER).(string),
	}
	populateMsgVpnBridgeFromResource(&bridge, d)

	b, _, err := semp_client.BridgeApi {
		Configuration: state.sempcfg,
	}.UpdateMsgVpnBridge(bridge.MsgVpnName, bridge.BridgeName, bridge.BridgeVirtualRouter, bridge, nil)

	if err != nil {
		log.Println("BridgeApi.UpdateMsgVpnBridge ERROR")
		return err
	}
	logSempMeta("VPN-Bridge update", b.Meta)

	return nil
}

func deleteMsgVpnBridgeFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	bridge := semp_client.MsgVpnBridge{
		BridgeName: d.Get(BRIDGE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
		BridgeVirtualRouter: d.Get(BRIDGE_VIRTUAL_ROUTER).(string),
	}

	b, _, err := semp_client.BridgeApi {
		Configuration: state.sempcfg,
		}.DeleteMsgVpnBridge(bridge.MsgVpnName, bridge.BridgeName, bridge.BridgeVirtualRouter)

	if err != nil {
		log.Println("BridgeApi.DeleteMsgVpnBridge ERROR")
		return err
	}
	logSempMeta("VPN-Bridge delete", b.Meta)

	return nil
}

