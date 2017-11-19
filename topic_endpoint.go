package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
)

// Main resource definition for Topic-Endpoint entities
func resourceTopicEndpoint() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema: schemaTopicEndpoint(),
		// Provider CRUD functions
		Create: createTopicEndpointFunc,
		Read:   readTopicEndpointFunc,
		Update: updateTopicEndpointFunc,
		Delete: deleteTopicEndpointFunc,
	}
}

// List of supported configuration fields for Solace topicEndpoints
// This method will only get called when initializing the topicEndpoint resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func schemaTopicEndpoint() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		TOPIC_ENDPOINT_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		CONSUMER_ACK_PROPAGATION_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		DEAD_MSG_QUEUE: {
			Type:     schema.TypeString,
			Optional: true,
		},
		EGRESS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		INGRESS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		MAX_DELIVERED_UNACKED_MSGS_PER_FLOW: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_MSG_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_REDELIVERY_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_SPOOL_USAGE: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		MAX_TTL: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		OWNER: {
			Type:     schema.TypeString,
			Optional: true,
		},
		PERMISSION: {
			Type:     schema.TypeString,
			Optional: true,
		},
		REJECT_LOW_PRIORITY_MSG_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
		REJECT_LOW_PRIORITY_MSG_LIMIT: {
			Type:     schema.TypeInt,
			Optional: true,
		},
		REJECT_MSG_TO_SENDER_ON_DISCARD_BEHAVIOR: {
			Type:     schema.TypeString,
			Optional: true,
		},
		RESPECT_TTL_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

// Function to populate all optional fields on a topicEndpoint; ideally this code should be
// generated from the Swagger schema or generated DAO.
func populateTopicEndpointFromResource(dte *semp_client.MsgVpnTopicEndpoint, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(CONSUMER_ACK_PROPAGATION_ENABLED)
	if b {
		dte.ConsumerAckPropagationEnabled = v.(bool)
	}
	v,b = d.GetOk(DEAD_MSG_QUEUE)
	if b {
		dte.DeadMsgQueue = v.(string)
	}
	v,b = d.GetOk(EGRESS_ENABLED)
	if b {
		dte.EgressEnabled = v.(bool)
	}
	v,b = d.GetOk(INGRESS_ENABLED)
	if b {
		dte.IngressEnabled = v.(bool)
	}
	v,b = d.GetOk(MAX_DELIVERED_UNACKED_MSGS_PER_FLOW)
	if b {
		dte.MaxDeliveredUnackedMsgsPerFlow = int64(v.(int))
	}
	v,b = d.GetOk(MAX_MSG_SIZE)
	if b {
		dte.MaxMsgSize = int32(v.(int))
	}
	v,b = d.GetOk(MAX_REDELIVERY_COUNT)
	if b {
		dte.MaxRedeliveryCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_SPOOL_USAGE)
	if b {
		dte.MaxSpoolUsage = int64(v.(int))
	}
	v,b = d.GetOk(MAX_TTL)
	if b {
		dte.MaxTtl = int64(v.(int))
	}
	v,b = d.GetOk(OWNER)
	if b {
		dte.Owner = v.(string)
	}
	v,b = d.GetOk(PERMISSION)
	if b {
		dte.Permission = v.(string)
	}
	v,b = d.GetOk(REJECT_LOW_PRIORITY_MSG_ENABLED)
	if b {
		dte.RejectLowPriorityMsgEnabled = v.(bool)
	}
	v,b = d.GetOk(REJECT_LOW_PRIORITY_MSG_LIMIT)
	if b {
		dte.RejectLowPriorityMsgLimit = int64(v.(int))
	}
	v,b = d.GetOk(REJECT_MSG_TO_SENDER_ON_DISCARD_BEHAVIOR)
	if b {
		dte.RejectMsgToSenderOnDiscardBehavior = v.(string)
	}
	v,b = d.GetOk(RESPECT_TTL_ENABLED)
	if b {
		dte.RespectTtlEnabled = v.(bool)
	}
}


// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createTopicEndpointFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createTopicEndpointFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	dte := semp_client.MsgVpnTopicEndpoint{
		TopicEndpointName:  d.Get(TOPIC_ENDPOINT_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateTopicEndpointFromResource(&dte, d)

	e, _, err := semp_client.TopicEndpointApi{
			Configuration: state.sempcfg,
		}.CreateMsgVpnTopicEndpoint(dte.MsgVpnName, dte, nil)

	if err != nil {
		log.Println("TopicEndpointApi.CreateMsgVpnTopicEndpoint ERROR")
		return err
	}
	logSempMeta("Topic-Endpoint create", e.Meta)

	d.SetId( state.host + "_" + dte.MsgVpnName + "_dte_" + dte.TopicEndpointName )

	return nil
}

func readTopicEndpointFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	resp, _, err := semp_client.TopicEndpointApi {
			Configuration:state.sempcfg,
		}.GetMsgVpnTopicEndpoint(d.Get(MSG_VPN_NAME).(string), d.Get(TOPIC_ENDPOINT_NAME).(string), nil)

	if err != nil {
		log.Println("TopicEndpointApi.GetMsgVpnTopicEndpoint ERROR")
		return err
	}
	logSempMeta("Topic-Endpoint read", resp.Meta)

	return nil
}

func updateTopicEndpointFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	dte := semp_client.MsgVpnTopicEndpoint{
		TopicEndpointName:  d.Get(TOPIC_ENDPOINT_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateTopicEndpointFromResource(&dte, d)

	e, _, err := semp_client.TopicEndpointApi{
		Configuration: state.sempcfg,
	}.UpdateMsgVpnTopicEndpoint(dte.MsgVpnName, dte.TopicEndpointName, dte, nil)

	if err != nil {
		log.Println("TopicEndpointApi.UpdateMsgVpnTopicEndpoint ERROR")
		return err
	}
	logSempMeta("Topic-Endpoint update", e.Meta)

	return nil
}

func deleteTopicEndpointFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	dte := semp_client.MsgVpnTopicEndpoint{
		TopicEndpointName: d.Get(TOPIC_ENDPOINT_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}

	e, _, err := semp_client.TopicEndpointApi{
		Configuration: state.sempcfg,
		}.DeleteMsgVpnTopicEndpoint(dte.MsgVpnName, dte.TopicEndpointName)

	if err != nil {
		log.Println("TopicEndpointApi.DeleteMsgVpnTopicEndpoint ERROR")
		return err
	}
	logSempMeta("Topic-Endpoint delete", e.Meta)

	return nil
}

