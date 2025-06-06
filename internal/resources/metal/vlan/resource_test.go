package vlan_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/equinix/terraform-provider-equinix/internal/acceptance"
	"github.com/equinix/terraform-provider-equinix/internal/config"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/packethost/packngo"
)

func testAccCheckMetalVlanConfig_metro(projSuffix, metro, desc string) string {
	return fmt.Sprintf(`
resource "equinix_metal_project" "foobar" {
    name = "tfacc-vlan-%s"
}

resource "equinix_metal_vlan" "foovlan" {
    project_id = equinix_metal_project.foobar.id
    metro = "%s"
    description = "%s"
}
`, projSuffix, metro, desc)
}

func testAccCheckMetalVlanConfig_NoDescription(projSuffix, metro string) string {
	return fmt.Sprintf(`
resource "equinix_metal_project" "foobar" {
    name = "tfacc-vlan-%s"
}

resource "equinix_metal_vlan" "foovlan" {
    project_id = equinix_metal_project.foobar.id
    metro = "%s"
}
`, projSuffix, metro)
}

func TestAccMetalVlan_metro(t *testing.T) {
	var vlan packngo.VirtualNetwork
	rs := acctest.RandString(10)
	lowerSiliconValley := "sv"
	upperDallas := "DA"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acceptance.TestAccPreCheckMetal(t) },
		ExternalProviders:        acceptance.TestExternalProviders,
		ProtoV6ProviderFactories: acceptance.ProtoV6ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				// Create VLAN with metro "sv" (lower-case)
				Config: testAccCheckMetalVlanConfig_metro(rs, lowerSiliconValley, "tfacc-vlan"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", "tfacc-vlan"),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", lowerSiliconValley),
				),
			},
			{
				// Confirm no changes if metro is changed to "SV" (upper-case)
				Config:   testAccCheckMetalVlanConfig_metro(rs, strings.ToUpper(lowerSiliconValley), "tfacc-vlan"),
				PlanOnly: true,
			},
			{
				// Recreate VLAN with metro "DA" (upper-case)
				Config: testAccCheckMetalVlanConfig_metro(rs, upperDallas, "tfacc-vlan"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("equinix_metal_vlan.foovlan", plancheck.ResourceActionDestroyBeforeCreate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", "tfacc-vlan"),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", upperDallas),
				),
			},
			{
				// Confirm no changes if metro is changed to "da" (lower-case)
				Config:   testAccCheckMetalVlanConfig_metro(rs, strings.ToLower(upperDallas), "tfacc-vlan"),
				PlanOnly: true,
			},
		},
	})
}

func TestAccMetalVlan_descriptionUpdate(t *testing.T) {
	var vlan packngo.VirtualNetwork
	rs := acctest.RandString(10)
	metro := "sv"
	description := "tfacc-vlan"
	updatedDescription := "tfacc-vlan-updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acceptance.TestAccPreCheckMetal(t) },
		ExternalProviders:        acceptance.TestExternalProviders,
		ProtoV6ProviderFactories: acceptance.ProtoV6ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				// Create VLAN with description
				Config: testAccCheckMetalVlanConfig_metro(rs, metro, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", description),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", metro),
				),
			},
			{
				// Update VLAN with description "tfacc-vlan-updated"
				Config: testAccCheckMetalVlanConfig_metro(rs, metro, updatedDescription),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("equinix_metal_vlan.foovlan", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", updatedDescription),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", metro),
				),
			},
		},
	})
}

func TestAccMetalVlan_NoDescription(t *testing.T) {
	var vlan packngo.VirtualNetwork
	rs := acctest.RandString(10)
	metro := "sv"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acceptance.TestAccPreCheckMetal(t) },
		ExternalProviders:        acceptance.TestExternalProviders,
		ProtoV6ProviderFactories: acceptance.ProtoV6ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMetalVlanConfig_NoDescription(rs, metro),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", ""),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", metro),
				),
			},
		},
	})
}

func TestAccMetalVlan_RemoveDescription(t *testing.T) {
	var vlan packngo.VirtualNetwork
	rs := acctest.RandString(10)
	metro := "sv"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acceptance.TestAccPreCheckMetal(t) },
		ExternalProviders:        acceptance.TestExternalProviders,
		ProtoV6ProviderFactories: acceptance.ProtoV6ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMetalVlanConfig_metro(rs, metro, "tfacc-vlan"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", "tfacc-vlan"),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", metro),
				),
			},
			{
				Config: testAccCheckMetalVlanConfig_NoDescription(rs, metro),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", ""),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", metro),
				),
			},
		},
	})
}

func testAccCheckMetalVlanExists(n string, vlan *packngo.VirtualNetwork) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		client := acceptance.TestAccProvider.Meta().(*config.Config).Metal

		foundVlan, _, err := client.ProjectVirtualNetworks.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}
		if foundVlan.ID != rs.Primary.ID {
			return fmt.Errorf("Record not found: %v - %v", rs.Primary.ID, foundVlan)
		}

		*vlan = *foundVlan

		return nil
	}
}

func testAccMetalVlanCheckDestroyed(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*config.Config).Metal

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "equinix_metal_vlan" {
			continue
		}
		if _, _, err := client.ProjectVirtualNetworks.Get(rs.Primary.ID, nil); err == nil {
			return fmt.Errorf("Metal Vlan still exists")
		}
	}

	return nil
}

func TestAccMetalVlan_importBasic(t *testing.T) {
	rs := acctest.RandString(10)
	metro := "sv"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acceptance.TestAccPreCheckMetal(t) },
		ExternalProviders:        acceptance.TestExternalProviders,
		ProtoV6ProviderFactories: acceptance.ProtoV6ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMetalVlanConfig_metro(rs, metro, "tfacc-vlan"),
			},
			{
				ResourceName:      "equinix_metal_vlan.foovlan",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccMetalVlan_metro_upgradeFromVersion(t *testing.T) {
	var vlan packngo.VirtualNetwork
	rs := acctest.RandString(10)
	metro := "sv"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.TestAccPreCheckMetal(t) },
		CheckDestroy: testAccMetalDatasourceVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"equinix": {
						VersionConstraint: "1.29.0", // latest version with resource defined on SDKv2
						Source:            "equinix/equinix",
					},
				},
				Config: testAccCheckMetalVlanConfig_metro(rs, metro, "tfacc-vlan"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", "tfacc-vlan"),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", metro),
				),
			},
			{
				ProtoV6ProviderFactories: acceptance.ProtoV6ProviderFactories,
				Config:                   testAccCheckMetalVlanConfig_metro(rs, metro, "tfacc-vlan"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccMetalVlan_metro_suppress_diff(t *testing.T) {
	var vlan packngo.VirtualNetwork
	rs := acctest.RandString(10)
	metro := "sv"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acceptance.TestAccPreCheckMetal(t) },
		ExternalProviders:        acceptance.TestExternalProviders,
		ProtoV6ProviderFactories: acceptance.ProtoV6ProviderFactories,
		CheckDestroy:             testAccMetalVlanCheckDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMetalVlanConfig_metro(rs, metro, "tfacc-vlan"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetalVlanExists("equinix_metal_vlan.foovlan", &vlan),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "description", "tfacc-vlan"),
					resource.TestCheckResourceAttr(
						"equinix_metal_vlan.foovlan", "metro", metro),
				),
			},
			{
				Config: testAccCheckMetalVlanConfig_metro(rs, strings.ToUpper(metro), "tfacc-vlan"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}
