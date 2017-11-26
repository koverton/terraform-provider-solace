package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/koverton/semp_client"
	"log"
	"strings"
	"strconv"
)

// Main resource definition for Client-Username entities
func resourceAclProfile() *schema.Resource {
	return &schema.Resource {
		SchemaVersion: 1,
		// List of supported configuration fields for your resource
		Schema:        aclProfileSchema(),
		// Provider CRUD functions
		Create:        createAclProfileFunc,
		Read:          readAclProfileFunc,
		Update:        updateAclProfileFunc,
		Delete:        deleteAclProfileFunc,
	}
}

// List of supported configuration fields for Solace aclProfiles
// This method will only get called when initializing the aclProfile resource
// Ideally this code should be generated from the Swagger schema or generated DAO.
func aclProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		ACL_PROFILE_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		MSG_VPN_NAME: {
			Type:     schema.TypeString,
			Required: true,
		},
		CLIENT_CONNECT_DEFAULT_ACTION: {
			Type:     schema.TypeString,
			Optional: true,
		},
		PUBLISH_TOPIC_DEFAULT_ACTION: {
			Type:     schema.TypeString,
			Optional: true,
		},
		SUBSCRIBE_TOPIC_DEFAULT_ACTION: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ACL_CONNECT_EXCEPTION_LIST: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ACL_PUBLISH_EXCEPTION_LIST: {
			Type:     schema.TypeString,
			Optional: true,
		},
		ACL_SUBSCRIBE_EXCEPTION_LIST: {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

// Generated from a field-list, either Swagger schema or generated GO client code.
// Given a DAO instance, populate it with any fields found in the resource data.
func populateAclProfileFromResource(acl *semp_client.MsgVpnAclProfile, d *schema.ResourceData) {
	// All optional fields should only be set if present in the resource data
	v,b := d.GetOk(ACL_PROFILE_NAME)
	if b {
		acl.AclProfileName = v.(string)
	}
	v,b = d.GetOk(MSG_VPN_NAME)
	if b {
		acl.MsgVpnName = v.(string)
	}
	v,b = d.GetOk(CLIENT_CONNECT_DEFAULT_ACTION)
	if b {
		acl.ClientConnectDefaultAction = v.(string)
	}
	v,b = d.GetOk(PUBLISH_TOPIC_DEFAULT_ACTION)
	if b {
		acl.PublishTopicDefaultAction = v.(string)
	}
	v,b = d.GetOk(SUBSCRIBE_TOPIC_DEFAULT_ACTION)
	if b {
		acl.SubscribeTopicDefaultAction = v.(string)
	}

}

func populateAclProfileExceptions(api *semp_client.AclProfileApi, d *schema.ResourceData) error {
	aclname := d.Get(ACL_PROFILE_NAME).(string)
	vpnname := d.Get(MSG_VPN_NAME).(string)

	exclist, b := d.GetOk(ACL_CONNECT_EXCEPTION_LIST)
	if b {
		for i, s := range strings.Split(exclist.(string), ",") {
			exc := semp_client.MsgVpnAclProfileClientConnectException {
				MsgVpnName:                    vpnname,
				AclProfileName:                aclname,
				ClientConnectExceptionAddress: s,
			}
			r, _, err := api.CreateMsgVpnAclProfileClientConnectException(vpnname, aclname, exc, nil)
			if err != nil {
				log.Printf("AclProfileApi.CreateMsgVpnAclProfileClientConnectException(%d) ERROR\n", i)
				return err
			}
			logSempMeta("AclProfileApi Connect Exception Create "+strconv.Itoa(i), r.Meta)
		}
	}
	exclist, b = d.GetOk(ACL_PUBLISH_EXCEPTION_LIST)
	if b {
		for i, s := range strings.Split(exclist.(string), ",") {
			exc := semp_client.MsgVpnAclProfilePublishException {
				MsgVpnName:            vpnname,
				AclProfileName:        aclname,
				PublishExceptionTopic: s,
				TopicSyntax:           "smf",
			}
			r, _, err := api.CreateMsgVpnAclProfilePublishException(vpnname, aclname, exc, nil)
			if err != nil {
				log.Printf("AclProfileApi.CreateMsgVpnAclProfilePublishException(%d) ERROR\n", i)
				return err
			}
			logSempMeta("AclProfileApi Publish Exception Create "+strconv.Itoa(i), r.Meta)
		}
	}
	exclist, b = d.GetOk(ACL_SUBSCRIBE_EXCEPTION_LIST)
	if b {
		for i, s := range strings.Split(exclist.(string), ",") {
			exc := semp_client.MsgVpnAclProfileSubscribeException {
				MsgVpnName:              vpnname,
				AclProfileName:          aclname,
				SubscribeExceptionTopic: s,
				TopicSyntax:           "smf",
			}
			r, _, err := api.CreateMsgVpnAclProfileSubscribeException(vpnname, aclname, exc, nil)
			if err != nil {
				log.Printf("AclProfileApi.CreateMsgVpnAclProfileSubscribeException(%d) ERROR\n", i)
				return err
			}
			logSempMeta("AclProfileApi Subscribe Exception Create "+strconv.Itoa(i), r.Meta)
		}
	}
	return nil
}

// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +
//		CRUD Callback Functions
//
// The methods defined below will get called for each resource that needs to
// get created, read, updated and deleted.
// For example, if 10 msg-VPNs need to be created then `createAclProfileFunc`
// will get called 10 times every time with the information for the proper
// resource that is being mapped.
//
// If at some point any of these functions returns an error, Terraform will
// imply that something went wrong with the modification of the resource and it
// will prevent the execution of further calls that depend on that resource
// that failed to be created/updated/deleted.
// + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - + - +

func createAclProfileFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	acl := semp_client.MsgVpnAclProfile{
		AclProfileName:  d.Get(ACL_PROFILE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateAclProfileFromResource(&acl, d)

	api := semp_client.AclProfileApi{Configuration: state.sempcfg }

	cu, _, err := api.CreateMsgVpnAclProfile(acl.MsgVpnName, acl, nil)
	if err != nil {
		log.Println("AclProfileApi.CreateMsgVpnAclProfile ERROR")
		return err
	}
	logSempMeta("ACL-Profile create", cu.Meta)
	// Must uniquely identify the resource within terraform
	d.SetId( state.host + "_" + acl.MsgVpnName + "_aclprofile_" + acl.AclProfileName )

	return populateAclProfileExceptions(&api, d)
}

func readAclProfileFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)

	resp, _, err := semp_client.AclProfileApi {
			Configuration:state.sempcfg,
		}.GetMsgVpnAclProfile(d.Get(MSG_VPN_NAME).(string), d.Get(ACL_PROFILE_NAME).(string), []string { "*" })

	if err != nil {
		log.Println("AclProfileApi.GetMsgVpnAclProfile ERROR")
		return err
	}
	logSempMeta("ACL-Profile read", resp.Meta)

	return nil
}

func updateAclProfileFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	// These are the only required fields, so init them upfront
	acl := semp_client.MsgVpnAclProfile{
		AclProfileName:  d.Get(ACL_PROFILE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}
	populateAclProfileFromResource(&acl, d)

	api := semp_client.AclProfileApi{Configuration: state.sempcfg }

	cu, _, err := api.UpdateMsgVpnAclProfile(acl.MsgVpnName, acl.AclProfileName, acl, nil)

	if err != nil {
		log.Println("AclProfileApi.UpdateMsgVpnAclProfile ERROR")
		return err
	}
	logSempMeta("ACL-Profile update", cu.Meta)

	return populateAclProfileExceptions(&api, d)
}

func deleteAclProfileFunc(d *schema.ResourceData, meta interface{}) error {
	state := meta.(*ProviderState)
	acl := semp_client.MsgVpnAclProfile{
		AclProfileName: d.Get(ACL_PROFILE_NAME).(string),
		MsgVpnName: d.Get(MSG_VPN_NAME).(string),
	}


	cu, _, err := semp_client.AclProfileApi{
		Configuration: state.sempcfg,
		}.DeleteMsgVpnAclProfile(acl.MsgVpnName, acl.AclProfileName)

	if err != nil {
		log.Println("AclProfileApi.DeleteMsgVpnAclProfile ERROR")
		return err
	}
	logSempMeta("ACL-Profile delete", cu.Meta)

	return nil
}

