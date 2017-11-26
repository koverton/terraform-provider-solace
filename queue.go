package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
	"strings"
	"strconv"
)


// Main resource definition for Queue entities
func resourceQueue() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema: schemaQueue(),
		// Provider CRUD functions
		Create: createQueueFunc,
		Read:   readQueueFunc,
		Update: updateQueueFunc,
		Delete: deleteQueueFunc,
	}
}

// List of supported configuration fields for Solace queues
// This method will only get called when initializing the queue resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func schemaQueue() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		QUEUE_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		ACCESS_TYPE: {
			Type:     schema.TypeString,
			Optional: true,
			//Default:  "exclusive",
		},
		CONSUMER_ACK_PROPAGATION_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
			//Default:  true,
		},
		DEAD_MSG_QUEUE: {
			Type:     schema.TypeString,
			Optional: true,
		},
		EGRESS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		INGRESS_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		MAX_BIND_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
			//Default:  1000,
		},
		MAX_DELIVERED_UNACKED_MSGS_PER_FLOW: {
			Type:     schema.TypeInt,
			Optional: true,
			//Default:  10000,
		},
		MAX_MSG_SIZE: {
			Type:     schema.TypeInt,
			Optional: true,
			//Default:  10000000,
		},
		MAX_MSG_SPOOL_USAGE: {
			Type:     schema.TypeInt,
			Optional: true,
			//Default:  1500,
		},
		MAX_REDELIVERY_COUNT: {
			Type:     schema.TypeInt,
			Optional: true,
			//Default:  0,
		},
		MAX_TTL: {
			Type:     schema.TypeInt,
			Optional: true,
			// Default:  0,
		},
		OWNER: {
			Type:     schema.TypeString,
			Optional: true,
		},
		PERMISSION: {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "delete",
		},
		REJECT_LOW_PRIORITY_MSG_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
			//Default:  false,
		},
		REJECT_LOW_PRIORITY_MSG_LIMIT: {
			Type:     schema.TypeInt,
			Optional: true,
			//Default:  false,
		},
		REJECT_MSG_TO_SENDER_ON_DISCARD_BEHAVIOR: {
			Type:     schema.TypeString,
			Optional: true,
			//Default:  "when-queue-enabled",
		},
		RESPECT_TTL_ENABLED: {
			Type:     schema.TypeBool,
			Optional: true,
			//Default:  false,
		},
		// MANUALLY ADDED for queue-subscription specs
		TOPIC_SUBSCRIPTION_LIST: {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Generated from a field-list, either Swagger schema or generated GO client code.
// Given a DAO instance, populate it with any fields found in the resource data.
func populateQueueFromResource(queue *semp_client.MsgVpnQueue, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(ACCESS_TYPE)
	if b {
		queue.AccessType = v.(string)
	}
	v,b = d.GetOk(CONSUMER_ACK_PROPAGATION_ENABLED)
	if b {
		queue.ConsumerAckPropagationEnabled = v.(bool)
	}
	v,b = d.GetOk(DEAD_MSG_QUEUE)
	if b {
		queue.DeadMsgQueue = v.(string)
	}
	v,b = d.GetOk(EGRESS_ENABLED)
	if b {
		queue.EgressEnabled = v.(bool)
	}
	v,b = d.GetOk(INGRESS_ENABLED)
	if b {
		queue.IngressEnabled = v.(bool)
	}
	v,b = d.GetOk(MAX_BIND_COUNT)
	if b {
		queue.MaxBindCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_DELIVERED_UNACKED_MSGS_PER_FLOW)
	if b {
		queue.MaxDeliveredUnackedMsgsPerFlow = int64(v.(int))
	}
	v,b = d.GetOk(MAX_MSG_SIZE)
	if b {
		queue.MaxMsgSize = int32(v.(int))
	}
	v,b = d.GetOk(MAX_MSG_SPOOL_USAGE)
	if b {
		queue.MaxMsgSpoolUsage = int64(v.(int))
	}
	v,b = d.GetOk(MAX_REDELIVERY_COUNT)
	if b {
		queue.MaxRedeliveryCount = int64(v.(int))
	}
	v,b = d.GetOk(MAX_TTL)
	if b {
		queue.MaxTtl = int64(v.(int))
	}
	v,b = d.GetOk(OWNER)
	if b {
		queue.Owner = v.(string)
	}
	v,b = d.GetOk(PERMISSION)
	if b {
		queue.Permission = v.(string)
	}
	v,b = d.GetOk(REJECT_LOW_PRIORITY_MSG_ENABLED)
	if b {
		queue.RejectLowPriorityMsgEnabled = v.(bool)
	}
	v,b = d.GetOk(REJECT_LOW_PRIORITY_MSG_LIMIT)
	if b {
		queue.RejectLowPriorityMsgLimit = int64(v.(int))
	}
	v,b = d.GetOk(REJECT_MSG_TO_SENDER_ON_DISCARD_BEHAVIOR)
	if b {
		queue.RejectMsgToSenderOnDiscardBehavior = v.(string)
	}
	v,b = d.GetOk(RESPECT_TTL_ENABLED)
	if b {
		queue.RespectTtlEnabled = v.(bool)
	}
}

// Function to take subscription-list from the queue resource def and add them to the queue subscriptions
func addQueueTopicSubscriptions(cfg *semp_client.Configuration, d *schema.ResourceData) error {
	tlist, b := d.GetOk(TOPIC_SUBSCRIPTION_LIST)
	if b {
		qname := d.Get(QUEUE_NAME).(string)
		vpnname := d.Get(MSG_VPN_NAME).(string)
		for i, topicStr := range strings.Split(tlist.(string), ",") {
			sub := semp_client.MsgVpnQueueSubscription{
				MsgVpnName: vpnname,
				QueueName:  qname,
				SubscriptionTopic: topicStr,
			}

			qr, _, err := semp_client.QueueApi{
				Configuration: cfg,
			}.CreateMsgVpnQueueSubscription(vpnname, qname, sub, nil)
			if err != nil {
				log.Printf("QueueApi.CreateMsgVpnQueueSubscription(%d) ERROR\n", i)
				return err
			}
			logSempMeta("Queue Subscription Create "+strconv.Itoa(i), qr.Meta)
		}
	}
	return nil
}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createQueueFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createQueueFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	queue  := semp_client.MsgVpnQueue{
		QueueName:  d.Get(QUEUE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateQueueFromResource(&queue, d)

	q, _, err := semp_client.QueueApi{
			Configuration: state.sempcfg,
		}.CreateMsgVpnQueue(queue.MsgVpnName, queue, nil)

	if err != nil {
		log.Println("QueueApi.CreateMsgVpnQueue ERROR")
		return err
	}
	logSempMeta("Queue Create", q.Meta)

	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + queue.MsgVpnName + "_q_" + queue.QueueName )

	return addQueueTopicSubscriptions(state.sempcfg, d)
}

func readQueueFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	q, _, err := semp_client.QueueApi {
			Configuration:state.sempcfg,
		}.GetMsgVpnQueue(d.Get(QUEUE_NAME).(string), d.Get(MSG_VPN_NAME).(string), []string { "*" })

	if err != nil {
		log.Println("QueueApi.GetMsgVpnQueue ERROR")
		return err
	}
	logSempMeta("Queue read", q.Meta)

	return addQueueTopicSubscriptions(state.sempcfg, d)
}

func updateQueueFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	queue  := semp_client.MsgVpnQueue{
		QueueName:  d.Get(QUEUE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateQueueFromResource(&queue, d)

	q, _, err := semp_client.QueueApi{
		Configuration: state.sempcfg,
	}.UpdateMsgVpnQueue(queue.MsgVpnName, queue.QueueName, queue, nil)

	if err != nil {
		log.Println("QueueApi.UpdateMsgVpnQueue ERROR")
		return err
	}
	logSempMeta("Queue update", q.Meta)

	return nil
}

func deleteQueueFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	queue  := semp_client.MsgVpnQueue{
		QueueName: d.Get(QUEUE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}

	q, _, err := semp_client.QueueApi{
		Configuration: state.sempcfg,
		}.DeleteMsgVpnQueue(queue.MsgVpnName, queue.QueueName)

	if err != nil {
		log.Println("QueueApi.DeleteMsgVpnQueue ERROR")
		return err
	}
	logSempMeta("Queue delete", q.Meta)

	return nil
}

