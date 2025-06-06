package equinix_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/equinix/terraform-provider-equinix/equinix"
	"github.com/equinix/terraform-provider-equinix/internal/acceptance"
	testinghelpers "github.com/equinix/terraform-provider-equinix/internal/fabric/testing_helpers"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Note:
// Keeping data "equinix_fabric_routing_protocol" tests in this file
// due to the long setup times required for RP tests.
// The FCR, Connection and RPs will already be created in the resource test, so the
// data_source tests will just leverage the RPs there to retrieve the data and check results

func TestAccFabricCreateRoutingProtocols_PFCR(t *testing.T) {
	ports := testinghelpers.GetFabricEnvPorts(t)

	var portUUID string

	if len(ports) > 0 {
		portUUID = ports["pfcr"]["dot1q"][1].GetUuid()
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheck(t); acceptance.TestAccPreCheckProviderConfigured(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: checkRoutingProtocolDelete,
		Steps: []resource.TestStep{
			{
				Config: testAccFabricCreateRoutingProtocolConfig("RP_Conn_Test_PFCR", portUUID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("equinix_fabric_routing_protocol.direct", "id"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.direct", "type", "DIRECT"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.direct", "state", "PROVISIONED"),
					resource.TestCheckResourceAttrSet("equinix_fabric_routing_protocol.direct", "change.0.uuid"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.direct", "change.0.type", "ROUTING_PROTOCOL_CREATION"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.direct", "direct_ipv4.0.equinix_iface_ip", "190.1.1.1/30"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.direct", "direct_ipv6.0.equinix_iface_ip", "190::1:1/126"),

					resource.TestCheckResourceAttrSet("equinix_fabric_routing_protocol.bgp", "id"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "type", "BGP"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "state", "PROVISIONED"),
					resource.TestCheckResourceAttrSet("equinix_fabric_routing_protocol.bgp", "change.0.uuid"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "change.0.type", "ROUTING_PROTOCOL_CREATION"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.customer_peer_ip", "190.1.1.2"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.equinix_peer_ip", "190.1.1.1"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.enabled", "true"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.outbound_as_prepend_count", "1"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.inbound_med", "4"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.outbound_med", "7"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.customer_peer_ip", "190::1:2"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.equinix_peer_ip", "190::1:1"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.enabled", "true"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.outbound_as_prepend_count", "1"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.inbound_med", "4"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.outbound_med", "7"),
					resource.TestCheckResourceAttr("equinix_fabric_routing_protocol.bgp", "customer_asn", "100"),

					resource.TestCheckResourceAttrSet("data.equinix_fabric_routing_protocol.direct", "id"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.direct", "type", "DIRECT"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.direct", "state", "PROVISIONED"),
					resource.TestCheckResourceAttrSet("data.equinix_fabric_routing_protocol.direct", "change.0.uuid"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.direct", "change.0.type", "ROUTING_PROTOCOL_CREATION"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.direct", "direct_ipv4.0.equinix_iface_ip", "190.1.1.1/30"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.direct", "direct_ipv6.0.equinix_iface_ip", "190::1:1/126"),

					resource.TestCheckResourceAttrSet("data.equinix_fabric_routing_protocol.bgp", "id"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "type", "BGP"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "state", "PROVISIONED"),
					resource.TestCheckResourceAttrSet("data.equinix_fabric_routing_protocol.bgp", "change.0.uuid"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "change.0.type", "ROUTING_PROTOCOL_CREATION"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.customer_peer_ip", "190.1.1.2"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.equinix_peer_ip", "190.1.1.1"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "bgp_ipv4.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.customer_peer_ip", "190::1:2"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.equinix_peer_ip", "190::1:1"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "bgp_ipv6.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.equinix_fabric_routing_protocol.bgp", "customer_asn", "100"),
				),
			},
		},
	})
}

func testAccFabricCreateRoutingProtocolConfig(name, portUUID string) string {
	return fmt.Sprintf(`

resource "equinix_fabric_cloud_router" "this" {
	type = "XF_ROUTER"
	name = "RP_Test_PFCR"
	location{
		metro_code  = "SV"
	}
    order {
		purchase_order_number = "123485"
		term_length = 1
	}
	notifications{
		type = "ALL"
		emails = ["test@equinix.com", "test1@equinix.com"]
	}
	project{
		project_id = "33ec651f-cc99-48e0-94d3-47466899cdc7"
	}
	account {
		account_number = 201257
	}
	package {
		code = "STANDARD"
	}
}

resource "equinix_fabric_connection" "this" {
	type = "IP_VC"
	name = "%s"
	notifications{
		type = "ALL"
		emails = ["test@equinix.com","test1@equinix.com"]
	}
    order {
		purchase_order_number = "123485"
		term_length = 1
	}
	bandwidth = 50
	redundancy {
		priority= "PRIMARY"
	}
	a_side {
		access_point {
			type = "CLOUD_ROUTER"
			router {
				uuid = equinix_fabric_cloud_router.this.id
			}
		}
	}
	project{
		project_id = "33ec651f-cc99-48e0-94d3-47466899cdc7"
	}
	z_side {
		access_point {
			type = "COLO"
			port{
				uuid = "%s"
			}
			link_protocol {
				type= "DOT1Q"
				vlan_tag= 1501
			}
			location {
				metro_code = "SV"
			}
		}
	}
}

resource "equinix_fabric_routing_protocol" "direct" {
	connection_uuid = equinix_fabric_connection.this.id
	type = "DIRECT"
	name = "rp_direct_PFCR"
	direct_ipv4{
		equinix_iface_ip = "190.1.1.1/30"
	}
	direct_ipv6{
		equinix_iface_ip = "190::1:1/126"
	}
}

resource "equinix_fabric_routing_protocol" "bgp" {
	depends_on = [
      equinix_fabric_routing_protocol.direct
  	]
	connection_uuid = equinix_fabric_connection.this.id
	type = "BGP"
	name = "rp_bgp_PFCR"
	bgp_ipv4{
		customer_peer_ip = "190.1.1.2"
		outbound_as_prepend_count = "1"
		inbound_med = 4
		outbound_med = 7
	}
	bgp_ipv6{
		customer_peer_ip = "190::1:2"
		outbound_as_prepend_count = "1"
		inbound_med = 4
		outbound_med = 7
	}
	customer_asn = 100
}

data "equinix_fabric_routing_protocol" "direct" {
	connection_uuid = equinix_fabric_connection.this.id
	uuid = equinix_fabric_routing_protocol.direct.id
}

data "equinix_fabric_routing_protocol" "bgp" {
	connection_uuid = equinix_fabric_connection.this.id
	uuid = equinix_fabric_routing_protocol.bgp.id
}

`, name, portUUID)
}

func checkRoutingProtocolDelete(s *terraform.State) error {
	ctx := context.Background()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "equinix_fabric_routing_protocol" {
			continue
		}
		err := equinix.WaitUntilRoutingProtocolIsDeprovisioned(rs.Primary.ID, rs.Primary.Attributes["connection_uuid"], acceptance.TestAccProvider.Meta(), &schema.ResourceData{}, ctx, 10*time.Minute)
		if err != nil {
			return fmt.Errorf("API call failed while waiting for resource deletion")
		}
	}
	return nil
}
