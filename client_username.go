package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)

// Main resource definition for Client-Username entities
func resourceClientUsername() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema:        clientUsernameSchema(),
		// Provider CRUD functions
		Create:        createClientUsernameFunc,
		Read:          readClientUsernameFunc,
		Update:        updateClientUsernameFunc,
		Delete:        deleteClientUsernameFunc,
	}
}

// List of supported configuration fields for Solace clientUsernames
// This method will only get called when initializing the clientUsername resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func clientUsernameSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		CLIENT_USERNAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		ACL_PROFILE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
		},
		CLIENT_PROFILE_NAME: {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
		},
		ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		GUARANTEED_ENDPOINT_PERMISSION_OVERRIDE_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		PASSWORD: {
			Type:     schema.TypeString,
			Optional: true,
		},
		SUBSCRIPTION_MANAGER_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

// Generated from a field-list, either Swagger schema or generated GO client code.
// Given a DAO instance, populate it with any fields found in the resource data.
func populateClientUsernameFromResource(user *semp_client.MsgVpnClientUsername, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(ACL_PROFILE_NAME)
	if b {
		user.AclProfileName = v.(string)
	}
	v,b = d.GetOk(CLIENT_PROFILE_NAME)
	if b {
		user.ClientProfileName = v.(string)
	}
	v,b = d.GetOk(CLIENT_USERNAME)
	if b {
		user.ClientUsername = v.(string)
	}
	v,b = d.GetOk(ENABLED)
	if b {
		user.Enabled = v.(bool)
	}
	v,b = d.GetOk(GUARANTEED_ENDPOINT_PERMISSION_OVERRIDE_ENABLED)
	if b {
		user.GuaranteedEndpointPermissionOverrideEnabled = v.(bool)
	}
	v,b = d.GetOk(MSG_VPN_NAME)
	if b {
		user.MsgVpnName = v.(string)
	}
	v,b = d.GetOk(PASSWORD)
	if b {
		user.Password = v.(string)
	}
	v,b = d.GetOk(SUBSCRIPTION_MANAGER_ENABLED)
	if b {
		user.SubscriptionManagerEnabled = v.(bool)
	}
}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createClientUsernameFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createClientUsernameFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	user := semp_client.MsgVpnClientUsername{
		ClientUsername:  d.Get(CLIENT_USERNAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateClientUsernameFromResource(&user, d)

	cu, _, err := semp_client.ClientUsernameApi{
			Configuration: state.sempcfg,
		}.CreateMsgVpnClientUsername(user.MsgVpnName, user, nil)

	if err != nil {
		log.Println("ClientUsernameApi.CreateMsgVpnClientUsername ERROR")
		return err
	}
	logSempMeta("client-username create", cu.Meta)
	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + user.MsgVpnName + "_clientusername_" + user.ClientUsername )

	return nil
}

func readClientUsernameFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	userresp, _, err := semp_client.ClientUsernameApi {
			Configuration:state.sempcfg,
		}.GetMsgVpnClientUsername(d.Get(MSG_VPN_NAME).(string), d.Get(CLIENT_USERNAME).(string), nil)

	if err != nil {
		log.Println("ClientUsernameApi.GetMsgVpnClientUsername ERROR")
		return err
	}
	logSempMeta("client-username read", userresp.Meta)

	return nil
}

func updateClientUsernameFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	user := semp_client.MsgVpnClientUsername{
		ClientUsername:  d.Get(CLIENT_USERNAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateClientUsernameFromResource(&user, d)

	cu, _, err := semp_client.ClientUsernameApi{
		Configuration: state.sempcfg,
	}.UpdateMsgVpnClientUsername(user.MsgVpnName, user.ClientUsername, user, nil)

	if err != nil {
		log.Println("ClientUsernameApi.UpdateMsgVpnClientUsername ERROR")
		return err
	}
	logSempMeta("client-username update", cu.Meta)

	return nil
}

func deleteClientUsernameFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	user := semp_client.MsgVpnClientUsername{
		ClientUsername: d.Get(CLIENT_USERNAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}

	cu, _, err := semp_client.ClientUsernameApi{
		Configuration: state.sempcfg,
		}.DeleteMsgVpnClientUsername(user.MsgVpnName, user.ClientUsername)

	if err != nil {
		log.Println("ClientUsernameApi.DeleteMsgVpnClientUsername ERROR")
		return err
	}
	logSempMeta("client-username delete", cu.Meta)

	return nil
}

