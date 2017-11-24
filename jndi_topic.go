package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)

// Main resource definition for Client-Username entities
func resourceJndiTopic() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema:        jndiTopicSchema(),
		// Provider CRUD functions
		Create:        createJndiTopicFunc,
		Read:          readJndiTopicFunc,
		Update:        updateJndiTopicFunc,
		Delete:        deleteJndiTopicFunc,
	}
}

// List of supported configuration fields for Solace jndiTopics
// This method will only get called when initializing the jndiTopic resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func jndiTopicSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		PHYSICAL_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		TOPIC_NAME: {
			Type:     schema.TypeString,
			Optional: true,
		},

	}
}

// Generated from a field-list, either Swagger schema or generated GO client code.
// Given a DAO instance, populate it with any fields found in the resource data.
func populateJndiTopicFromResource(jnditopic *semp_client.MsgVpnJndiTopic, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(MSG_VPN_NAME)
	if b {
		jnditopic.MsgVpnName = v.(string)
	}
	v,b = d.GetOk(PHYSICAL_NAME)
	if b {
		jnditopic.PhysicalName = v.(string)
	}
	v,b = d.GetOk(TOPIC_NAME)
	if b {
		jnditopic.TopicName = v.(string)
	}

}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createJndiTopicFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createJndiTopicFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	jnditopic := semp_client.MsgVpnJndiTopic{
		TopicName:  d.Get(TOPIC_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateJndiTopicFromResource(&jnditopic, d)

	resp, _, err := semp_client.JndiApi {
		Configuration: state.sempcfg,
	}.CreateMsgVpnJndiTopic(jnditopic.MsgVpnName, jnditopic, nil)

	if err != nil {
		log.Println("JndiTopicApi.CreateMsgVpnJndiTopic ERROR")
		return err
	}
	logSempMeta("JndiTopic create", resp.Meta)
	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + jnditopic.MsgVpnName + "_jndiTopic_" + jnditopic.TopicName )

	return nil
}

func readJndiTopicFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	resp, _, err := semp_client.JndiApi {
		Configuration:state.sempcfg,
	}.GetMsgVpnJndiTopic(d.Get(MSG_VPN_NAME).(string), d.Get(TOPIC_NAME).(string), nil)

	if err != nil {
		log.Println("JndiTopicApi.GetMsgVpnJndiTopic ERROR")
		return err
	}
	logSempMeta("JndiTopic read", resp.Meta)

	return nil
}

func updateJndiTopicFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	jnditopic := semp_client.MsgVpnJndiTopic{
		TopicName:  d.Get(TOPIC_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateJndiTopicFromResource(&jnditopic, d)

	cu, _, err := semp_client.JndiApi {
		Configuration: state.sempcfg,
	}.UpdateMsgVpnJndiTopic(jnditopic.MsgVpnName, jnditopic.TopicName, jnditopic, nil)

	if err != nil {
		log.Println("JndiTopicApi.UpdateMsgVpnJndiTopic ERROR")
		return err
	}
	logSempMeta("JndiTopic update", cu.Meta)

	return nil
}

func deleteJndiTopicFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	jnditopic := semp_client.MsgVpnJndiTopic{
		TopicName:  d.Get(TOPIC_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}

	cu, _, err := semp_client.JndiApi {
		Configuration: state.sempcfg,
	}.DeleteMsgVpnJndiTopic(jnditopic.MsgVpnName, jnditopic.TopicName)

	if err != nil {
		log.Println("JndiTopicApi.DeleteMsgVpnJndiTopic ERROR")
		return err
	}
	logSempMeta("JndiTopic delete", cu.Meta)

	return nil
}


