package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/koverton/semp_client"
	"strconv"
	"log"
)

// This is boilerplate Terraform bootstrapping code
// Note that it totally ignores commandline args because
// when terraform starts your provider it only passes the
// resource definition, not any cmdline arguments.
func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}

// This is the Provider closure passed into each resource
// CRUD function invocation
type ProviderState struct {
	// All SEMP resource APIs use this configuration to bind to the broker
	sempcfg  *semp_client.Configuration
	// The broker host is used to uniquely identify each resource
	host     string
}

// Main Resource Provider definition;
func Provider() terraform.ResourceProvider {
	log.Printf("Creating Terraform Solace Provider")
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResources(),
		ConfigureFunc: providerConfigure,
	}
}

// List of supported configuration fields for the Solace provider.
func providerSchema() map[string]*schema.Schema {
	log.Printf("Returning Solace Provider Schema")
	return map[string]*schema.Schema{
		HOST: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The address of the Solace msg broker",
		},
		PORT: {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "The port of the Solace msg broker",
		},
		ADMIN_USER: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Admin identity to login to the Solace VMR.",
		},
		ADMIN_PASSWD: {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Password of the admin identity used to login to the Solace VMR.",
		},
	}
}

// List of supported resources and their configuration fields.
// Each resource's complete definition is contained in a separate <resource>.go code file.
func providerResources() map[string]*schema.Resource {
	log.Printf("Returning Solace Provider Resources Definition")
	return map[string]*schema.Resource{
		RSRC_MSG_VPN:               resourceMsgVpn(),        // @see msg_vpn.go
		RSRC_QUEUE:                 resourceQueue(),         // @see queue.go
		RSRC_TOPIC_ENDPOINT:        resourceTopicEndpoint(), // @see topic_endpoint.go
		RSRC_CLIENT_USERNAME:       resourceClientUsername(),// @see client_username.go
		RSRC_CLIENT_PROFILE:        resourceClientProfile(), // @see client_profile.go
		RSRC_MSG_VPN_BRIDGE:        resourceMsgVpnBridge(),  // @see msg_vpn_bridge.go
		RSRC_VPN_BRIDGE_REMOTE_VPN: resourceMsgVpnBridgeRemoteVpn(), // @see msg_vpn_bridge_remote_vpn.go
	}
}

// Terraform calls your provider factory function to create and configure your provider
// before use. You can return any object you like here but, since this is the object
// passed to our resource functions, it  contains the semp_client configuration used
// by the semp_client backend API.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	cfg := semp_client.NewConfiguration()
	cfg.Username = d.Get(ADMIN_USER).(string)
	cfg.Password = d.Get(ADMIN_PASSWD).(string)
	host := d.Get(HOST).(string)
	port := d.Get(PORT).(int)
	cfg.Host     = host
	cfg.BasePath = "http://" + host + ":" + strconv.Itoa(port) + "/SEMP/v2/config"

	log.Printf("Configured semp_client: basePath:%s, user:%s, pass:%s\n",
		cfg.BasePath, cfg.Username, cfg.Password)

	state := ProviderState{
		sempcfg: cfg,
		host: host,
	}

	return &state, nil
}
