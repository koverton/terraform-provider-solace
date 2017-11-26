package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)

// Main resource definition for Client-Username entities
func resourceJndiQueue() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema:        jndiQueueSchema(),
		// Provider CRUD functions
		Create:        createJndiQueueFunc,
		Read:          readJndiQueueFunc,
		Update:        updateJndiQueueFunc,
		Delete:        deleteJndiQueueFunc,
	}
}

// List of supported configuration fields for Solace jndiQueues
// This method will only get called when initializing the jndiQueue resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func jndiQueueSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		PHYSICAL_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},
		QUEUE_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},

	}
}

// Generated from a field-list, either Swagger schema or generated GO client code.
// Given a DAO instance, populate it with any fields found in the resource data.
func populateJndiQueueFromResource(jndiq *semp_client.MsgVpnJndiQueue, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(MSG_VPN_NAME)
	if b {
		jndiq.MsgVpnName = v.(string)
	}
	v,b = d.GetOk(PHYSICAL_NAME)
	if b {
		jndiq.PhysicalName = v.(string)
	}
	v,b = d.GetOk(QUEUE_NAME)
	if b {
		jndiq.QueueName = v.(string)
	}

}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createJndiQueueFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createJndiQueueFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	jndiq := semp_client.MsgVpnJndiQueue{
		QueueName:  d.Get(QUEUE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateJndiQueueFromResource(&jndiq, d)

	cu, _, err := semp_client.JndiApi{
		Configuration: state.sempcfg,
	}.CreateMsgVpnJndiQueue(jndiq.MsgVpnName, jndiq, nil)

	if err != nil {
		log.Println("JndiQueueApi.CreateMsgVpnJndiQueue ERROR")
		return err
	}
	logSempMeta("JndiQueue create", cu.Meta)
	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + jndiq.MsgVpnName + "_jndiQueue_" + jndiq.QueueName )

	return nil
}

func readJndiQueueFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	resp, _, err := semp_client.JndiApi {
		Configuration:state.sempcfg,
	}.GetMsgVpnJndiQueue(d.Get(MSG_VPN_NAME).(string), d.Get(QUEUE_NAME).(string), []string { "*" })

	if err != nil {
		log.Println("JndiQueueApi.GetMsgVpnJndiQueue ERROR")
		return err
	}
	logSempMeta("JndiQueue read", resp.Meta)

	return nil
}

func updateJndiQueueFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	jndiq := semp_client.MsgVpnJndiQueue{
		QueueName:  d.Get(QUEUE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateJndiQueueFromResource(&jndiq, d)

	cu, _, err := semp_client.JndiApi {
		Configuration: state.sempcfg,
	}.UpdateMsgVpnJndiQueue(jndiq.MsgVpnName, jndiq.QueueName, jndiq, nil)

	if err != nil {
		log.Println("JndiQueueApi.UpdateMsgVpnJndiQueue ERROR")
		return err
	}
	logSempMeta("JndiQueue update", cu.Meta)

	return nil
}

func deleteJndiQueueFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	jndiq := semp_client.MsgVpnJndiQueue{
		QueueName: d.Get(QUEUE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}

	cu, _, err := semp_client.JndiApi {
		Configuration: state.sempcfg,
	}.DeleteMsgVpnJndiQueue(jndiq.MsgVpnName, jndiq.QueueName)

	if err != nil {
		log.Println("JndiQueueApi.DeleteMsgVpnJndiQueue ERROR")
		return err
	}
	logSempMeta("JndiQueue delete", cu.Meta)

	return nil
}


