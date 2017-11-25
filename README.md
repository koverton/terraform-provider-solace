# Solace Terraform Provider

This is a partial Terraform Provider for provisioning entities on 
Solace Message Brokers. If you are not familiar with Solace you can find 
more details at http://solace.com.

It is capable of connecting to a Solace Message Broker (hardware or VMR) 
and performing basic CRUD operations on that broker for the following 
Solace entity types:

- Message-VPNs
- Queues
- Topic Endpoints
- Client Profiles
- ACL Profiles
- Client Usernames
- VPN-Bridges and bridge Remote-VPNs
- JNDI conn-factories, queues and topics

Some things it can _not_ do yet:
- Set or update alerting thresholds on any entity type
- All other entities: REST delivery points, other JMS entities, MQTT sessions, etc.

## Configuring the Provider

The Solace Provider connects to an existing Solace Message Broker to provision 
entities on that Broker. The properties required for this are those  
properties required to connect to the SEMP port and authenticate as a 
user with administrative privileges:

- __hostname__: the message broker's management address or hostname (NOTE: for hardware brokers, the management interface is different than the data interface)
- __port__: the message broker's management semp management port (typically 8080 for VMRs or 80 for hardware)
- __admin_user__: the administrative user to authenticate as for the management session
- __admin_password__: the administrative user's credentials for the management session

NOTE: This means that an administrative user must be provisioned on the Solace router before the Terraform provider 
can do anything with it.

## Leveraging SEMPv2 Swagger Clients

[SEMPv2 is a Swagger-based API](https://github.com/koverton/semp_client) that supports code-generated client 
DAO's. The Solace Terraform Provider makes use of the GO generated 
client. This means that resource field definitions are generated from the same schema, 
where Swagger camel-case fields like `MsgVpnName` are converted to snake-case 
fields like `msg_vpn_name`. (This is because Terraform's resource parsers restrict us to only 
lowercase characters.) All the [fields available on entities in the swagger API](https://docs.solace.com/API-Developer-Online-Ref-Documentation/swagger-ui/index.html#/bridge) 
can be set in the terraform resources.

The provider does not provision brokers from scratch, it is configured to bind to an existing Message Broker instance.
To provision a Solace broker, for example, on can use the Solace image from the AWS 
AMI Catalog:

```
resource "aws_instance" "my_terraform_vmr_inst" {
  ami = "ami-1541b36f"
  instance_type = "t2.medium"
  key_name = "__YOUR_KEY__"
  security_groups = [ "VMRSecGroup" ]
  tags {
    Name = "my_terraform_vmr_inst"
  }
}
```

Once the VMR is instantiated, the Solace Terraform Provider can provision Msg-VPNs and 
entities inside them.

## Build and Run

This project uses the [Github code layout](https://github.com/golang/go/wiki/GithubCodeLayout) 
and requires the SEMP Swagger client has been code-generated into a package named `semp_client`:
```
$GOPATH/
    src/
        github.com/
            [GITHUBID]/
                terraform-provider-solace/
                    .git/
                    provider.go
                    ...
                semp_client/
                    msg_vpn_api.go
                    ...
```
Note that the code for each resource type imports that `semp_client`. 
If you branch it into your github environment you will need to update all those import statements.

With that in place, steps to build and run the provider are as follows:

```shell
linux> make
       go build
linux> terraform init
       ...
linux> terraform plan
       ...
Plan: 8 to add, 0 to change, 0 to destroy.
       ...
linux> terraform apply
       ...
       Apply complete! Resources: 8 added, 0 changed, 0 destroyed.
```

## Configuring Bridges

There are some challenges you will experience with the semp_client provisioning msg-VPN bridges that 
are documented here. They aren't issues specific to Terraform providers but are inherited from that 
`semp_client` dependency.

`remote_message_vpn_location` is a host or IP-address _with port specified_. If you do not configure the port 
it will fail to provision the resource.

_All client-username configuerations, either on the bridge or on the bridge remote-VPN, must include password 
configuration_. Yes, that is true even if you have disabled authentication and the password on the underlying 
identity is not configured.

## Debugging

Most CRUD functions in this library will log results and errors. To see them though, you need to enable that 
in a TF-specific environment variable:

    export TF_LOG=1
    
An additional debugging tip, whenever things are not working and it's hard to see why, test the matching 
`semp_client` functionality outside of this Terraform provider.

## Example Configuration

```
provider "solace" {
  host = "192.168.56.201"
  port = 8080
  admin_user     = "admin"
  admin_password = "secret"
}

resource "solace_msg_vpn" "test_vpn" {
  msg_vpn_name = "test_vpn"
  max_msg_spool_usage = 1000
  authentication_basic_enabled = true
  authentication_basic_type = "none"
  enabled = true
}

resource "solace_queue" "q_jimmy" {
  queue_name  = "jimmy"
  msg_vpn_name = "test_vpn"
  access_type = "exclusive"
  max_msg_spool_usage = 1000
  permission = "modify-topic"
  topic_subscription_list = "hello,all,you,happy,topics"
  depends_on = ["solace_msg_vpn.test_vpn"]
}

resource "solace_topic_endpoint" "dte_jimmy" {
  topic_endpoint_name  = "jimmy"
  msg_vpn_name = "test_vpn"
  max_spool_usage = 1000
  permission = "modify-topic"
  depends_on = ["solace_msg_vpn.test_vpn"]
}

resource "solace_client_profile"  "profile_jimmy" {
  client_profile_name = "jimmy"
  msg_vpn_name = "test_vpn"
  max_connection_count_per_client_username = 10
  allow_guaranteed_endpoint_create_enabled = true
  allow_guaranteed_msg_send_enabled = true
  allow_guaranteed_msg_receive_enabled = true
}

resource "solace_client_username" "user_jimmy" {
  client_username  = "jimmy"
  msg_vpn_name = "test_vpn"
  client_profile_name = "jimmy"
  enabled = true
  depends_on = ["solace_msg_vpn.test_vpn"]
  depends_on = ["solace_msg_vpn.profile_jimmy"]
}


```

