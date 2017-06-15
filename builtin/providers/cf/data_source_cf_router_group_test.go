package cloudfoundry

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/builtin/providers/cf/cfapi"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const routerGroupDataResource = `

data "cf_router_group" "rg" {
    name = "default-tcp"
}
`

func TestAccDataSourceRouterGroup_normal(t *testing.T) {

	ref := "data.cf_router_group.rg"

	resource.Test(t,
		resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{

				resource.TestStep{
					Config: routerGroupDataResource,
					Check: resource.ComposeTestCheckFunc(
						checkDataSourceRouterGroupExists(ref),
						resource.TestCheckResourceAttr(
							ref, "name", "default-tcp"),
					),
				},
			},
		})
}

func checkDataSourceRouterGroupExists(resource string) resource.TestCheckFunc {

	return func(s *terraform.State) error {

		session := testAccProvider.Meta().(*cfapi.Session)

		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("router_group '%s' not found in terraform state", resource)
		}

		session.Log.DebugMessage(
			"terraform state for resource '%s': %# v",
			resource, rs)

		id := rs.Primary.ID
		name := rs.Primary.Attributes["name"]
		rgType := rs.Primary.Attributes["type"]

		routerGroup, err := session.DomainManager().FindRouterGroupByName(name)
		if err != nil {
			return err
		}
		if err := assertSame(id, routerGroup.GUID); err != nil {
			return err
		}
		if err := assertSame(rgType, routerGroup.Type); err != nil {
			return err
		}

		return nil
	}
}
