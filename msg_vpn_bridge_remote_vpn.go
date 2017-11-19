package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)


// Main resource definition for VPN-Bridge entities
func resourceMsgVpnBridgeRemoteVpn() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema: schemaMsgVpnBridgeRemoteVpn(),
		// Provider CRUD functions
		Create: createMsgVpnBridgeRemoteVpnFunc,
		Read:   readMsgVpnBridgeRemoteVpnFunc,
		Update: updateMsgVpnBridgeRemoteVpnFunc,
		Delete: deleteMsgVpnBridgeRemoteVpnFunc,
	}
}

// List of supported configuration fields for Solace msgVpnBridgeRemoteVpns
// This method will only get called when initializing the msgVpnBridgeRemoteVpn resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func schemaMsgVpnBridgeRemoteVpn() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		BRIDGE_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		BRIDGE_VIRTUAL_ROUTER: {
			Type:     schema.TypeString,
			Required: true,
		},
		REMOTE_MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		REMOTE_MSG_VPN_LOCATION: {
			Type:     schema.TypeString,
			Required: true,
		},
		REMOTE_MSG_VPN_INTERFACE: {
			Type:     schema.TypeString,
			Required: true,
		},
		CLIENT_USERNAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		COMPRESSED_DATA_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		CONNECT_ORDER: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		EGRESS_FLOW_WINDOW_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		PASSWORD: {
			Type:     schema.TypeString,
			Optional: true,
		},
		QUEUE_BINDING: {
			Type:     schema.TypeString,
			Optional: true,
		},
		TLS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		UNIDIRECTIONAL_CLIENT_PROFILE: {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Function to populate all optional fields on a msgVpnBridgeRemoteVpn; ideally this code should be
// generated from the Swagger schema or generated DAO.
func populateMsgVpnBridgeRemoteVpnFromResource(remote *semp_client.MsgVpnBridgeRemoteMsgVpn, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(BRIDGE_NAME)
	if b {
		remote.BridgeName = v.(string)
	}
	v,b = d.GetOk(MSG_VPN_NAME)
	if b {
		remote.MsgVpnName = v.(string)
	}
	v,b = d.GetOk(BRIDGE_VIRTUAL_ROUTER)
	if b {
		remote.BridgeVirtualRouter = v.(string)
	}
	v,b = d.GetOk(REMOTE_MSG_VPN_NAME)
	if b {
		remote.RemoteMsgVpnName = v.(string)
	}
	v,b = d.GetOk(REMOTE_MSG_VPN_LOCATION)
	if b {
		remote.RemoteMsgVpnLocation = v.(string)
	}
	v,b = d.GetOk(REMOTE_MSG_VPN_INTERFACE)
	if b {
		remote.RemoteMsgVpnInterface = v.(string)
	}
	v,b = d.GetOk(CLIENT_USERNAME)
	if b {
		remote.ClientUsername = v.(string)
	}
	v,b = d.GetOk(COMPRESSED_DATA_ENABLED)
	if b {
		remote.CompressedDataEnabled = v.(bool)
	}
	v,b = d.GetOk(CONNECT_ORDER)
	if b {
		remote.ConnectOrder = v.(int32)
	}
	v,b = d.GetOk(EGRESS_FLOW_WINDOW_SIZE)
	if b {
		remote.EgressFlowWindowSize = v.(int64)
	}
	v,b = d.GetOk(ENABLED)
	if b {
		remote.Enabled = v.(bool)
	}
	v,b = d.GetOk(PASSWORD)
	if b {
		remote.Password = v.(string)
	}
	v,b = d.GetOk(QUEUE_BINDING)
	if b {
		remote.QueueBinding = v.(string)
	}
	v,b = d.GetOk(TLS_ENABLED)
	if b {
		remote.TlsEnabled = v.(bool)
	}
	v,b = d.GetOk(UNIDIRECTIONAL_CLIENT_PROFILE)
	if b {
		remote.UnidirectionalClientProfile = v.(string)
	}
}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createMsgVpnBridgeRemoteVpnFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createMsgVpnBridgeRemoteVpnFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	remote := semp_client.MsgVpnBridgeRemoteMsgVpn{
		BridgeName:  d.Get(BRIDGE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
		BridgeVirtualRouter: d.Get(BRIDGE_VIRTUAL_ROUTER).(string),
		RemoteMsgVpnName: d.Get(REMOTE_MSG_VPN_NAME).(string),
		RemoteMsgVpnLocation: d.Get(REMOTE_MSG_VPN_LOCATION).(string),
		RemoteMsgVpnInterface: d.Get(REMOTE_MSG_VPN_INTERFACE).(string),
	}
	populateMsgVpnBridgeRemoteVpnFromResource(&remote, d)

	r, _, err := semp_client.BridgeApi {
			Configuration: state.sempcfg,
		}.CreateMsgVpnBridgeRemoteMsgVpn(
			remote.MsgVpnName,
			remote.BridgeName,
			remote.BridgeVirtualRouter,
			remote,
			nil)
	if err != nil {
		log.Println("BridgeApi.CreateMsgVpnBridgeRemoteMsgVpn ERROR")
		return err
	}

	logSempMeta("VPN-Bridge Remote-VPN create", r.Meta)

	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + remote.MsgVpnName +
		"_vpnbridge_" + remote.BridgeName +
			"_remote_" + remote.RemoteMsgVpnName +
				"_virt_" + remote.BridgeVirtualRouter)

	return nil
}

func readMsgVpnBridgeRemoteVpnFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	resp, _, err := semp_client.BridgeApi {
			Configuration:state.sempcfg,
		}.GetMsgVpnBridgeRemoteMsgVpn(
			d.Get(BRIDGE_NAME).(string),
			d.Get(MSG_VPN_NAME).(string),
			d.Get(BRIDGE_VIRTUAL_ROUTER).(string),
			d.Get(REMOTE_MSG_VPN_NAME).(string),
			d.Get(REMOTE_MSG_VPN_LOCATION).(string),
			d.Get(REMOTE_MSG_VPN_INTERFACE).(string),
			nil)

	if err != nil {
		log.Println("BridgeApi.GetMsgVpnBridgeRemoteMsgVpn ERROR")
		return err
	}

	logSempMeta("VPN-Bridge Remote-VPN read", resp.Meta)

	return nil
}

func updateMsgVpnBridgeRemoteVpnFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	remote := semp_client.MsgVpnBridgeRemoteMsgVpn{
		BridgeName:  d.Get(BRIDGE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
		BridgeVirtualRouter: d.Get(BRIDGE_VIRTUAL_ROUTER).(string),
		RemoteMsgVpnName: d.Get(REMOTE_MSG_VPN_NAME).(string),
		RemoteMsgVpnLocation: d.Get(REMOTE_MSG_VPN_LOCATION).(string),
		RemoteMsgVpnInterface: d.Get(REMOTE_MSG_VPN_INTERFACE).(string),
	}
	populateMsgVpnBridgeRemoteVpnFromResource(&remote, d)

	r, _, err := semp_client.BridgeApi {
			Configuration: state.sempcfg,
		}.UpdateMsgVpnBridgeRemoteMsgVpn(
			remote.MsgVpnName,
			remote.BridgeName,
			remote.BridgeVirtualRouter,
			remote.RemoteMsgVpnName,
			remote.RemoteMsgVpnLocation,
			remote.RemoteMsgVpnInterface,
			remote,
			nil)

	if err != nil {
		log.Println("BridgeApi.UpdateMsgVpnBridgeRemoteMsgVpn ERROR")
		return err
	}
	logSempMeta("VPN-Bridge Remote-VPN update", r.Meta)

	return nil
}

func deleteMsgVpnBridgeRemoteVpnFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	remote := semp_client.MsgVpnBridgeRemoteMsgVpn{
		BridgeName:  d.Get(BRIDGE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
		BridgeVirtualRouter: d.Get(BRIDGE_VIRTUAL_ROUTER).(string),
		RemoteMsgVpnName: d.Get(REMOTE_MSG_VPN_NAME).(string),
		RemoteMsgVpnLocation: d.Get(REMOTE_MSG_VPN_LOCATION).(string),
		RemoteMsgVpnInterface: d.Get(REMOTE_MSG_VPN_INTERFACE).(string),
	}

	r, _, err := semp_client.BridgeApi {
			Configuration: state.sempcfg,
		}.DeleteMsgVpnBridgeRemoteMsgVpn(
			remote.MsgVpnName,
			remote.BridgeName,
			remote.BridgeVirtualRouter,
			remote.RemoteMsgVpnName,
			remote.RemoteMsgVpnLocation,
			remote.RemoteMsgVpnInterface)

	if err != nil {
		log.Println("BridgeApi.DeleteMsgVpnBridgeRemoteMsgVpn ERROR")
		return err
	}
	logSempMeta("VPN-Bridge Remote-VPN delete", r.Meta)

	return nil
}

