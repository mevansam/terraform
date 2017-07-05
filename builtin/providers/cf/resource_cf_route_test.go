package cloudfoundry

import (
	"fmt"
	"testing"

	"code.cloudfoundry.org/cli/cf/errors"

	"github.com/hashicorp/terraform/builtin/providers/cf/cfapi"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const routeResource = `

data "cf_domain" "local" {
    name = "local.pcfdev.io"
}
data "cf_org" "org" {
    name = "pcfdev-org"
}
data "cf_space" "space" {
    name = "pcfdev-space"
	org = "${data.cf_org.org.id}"
}

resource "cf_app" "test-app" {
	name = "test-app"
	space = "${data.cf_space.space.id}"
	command = "test-app --ports=8080"

	git {
		url = "https://github.com/mevansam/test-app.git"
	}
}
resource "cf_route" "test-app" {
	domain = "${data.cf_domain.local.id}"
	space = "${data.cf_space.space.id}"
	hostname = "test-app"

	target {
		app = "${cf_app.test-app.id}"
	}
}
`

const routeResourceUpdate = `

data "cf_domain" "local" {
    name = "local.pcfdev.io"
}
data "cf_org" "org" {
    name = "pcfdev-org"
}
data "cf_space" "space" {
    name = "pcfdev-space"
	org = "${data.cf_org.org.id}"
}

resource "cf_app" "test-app" {
	name = "test-app"
	space = "${data.cf_space.space.id}"
	command = "test-app --ports=8080"

	git {
		url = "https://github.com/mevansam/test-app.git"
	}
}
resource "cf_app" "test-app-8888" {
	name = "test-app-8888"
	space = "${data.cf_space.space.id}"
	ports = [ 8888 ]
	command = "test-app --ports=8888"

	git {
		url = "https://github.com/mevansam/test-app.git"
	}
}
resource "cf_app" "test-app-9999" {
	name = "test-app-9999"
	space = "${data.cf_space.space.id}"
	ports = [ 9999 ]
	command = "test-app --ports=9999"

	git {
		url = "https://github.com/mevansam/test-app.git"
	}
}
resource "cf_route" "test-app" {
	domain = "${data.cf_domain.local.id}"
	space = "${data.cf_space.space.id}"
	hostname = "test-app-multi"

	target {
		app = "${cf_app.test-app-9999.id}"
		port = 9999
	}
	target {
		app = "${cf_app.test-app-8888.id}"
		port = 8888
	}
	target {
		app = "${cf_app.test-app.id}"
	}
}
`

func TestAccRoute_normal(t *testing.T) {

	refRoute := "cf_route.test-app"

	resource.Test(t,
		resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: testAccCheckRouteDestroyed([]string{"test-app", "test-app-updated"}, "local.pcfdev.io"),
			Steps: []resource.TestStep{

				resource.TestStep{
					Config: routeResource,
					Check: resource.ComposeTestCheckFunc(
						testAccCheckRouteExists(refRoute, func() (err error) {

							responses := []string{"8080"}
							if err = assertHTTPResponse("http://test-app.local.pcfdev.io/port", 200, &responses); err != nil {
								return err
							}
							return
						}),
						resource.TestCheckResourceAttr(
							refRoute, "hostname", "test-app"),
						resource.TestCheckResourceAttr(
							refRoute, "target.#", "1"),
					),
				},

				resource.TestStep{
					Config: routeResourceUpdate,
					Check: resource.ComposeTestCheckFunc(
						testAccCheckRouteExists(refRoute, func() (err error) {

							responses := []string{"8080", "8888", "9999"}
							for i := 1; i <= 9; i++ {
								if err = assertHTTPResponse("http://test-app-multi.local.pcfdev.io/port", 200, &responses); err != nil {
									return err
								}
							}
							return
						}),
						resource.TestCheckResourceAttr(
							refRoute, "hostname", "test-app-multi"),
						resource.TestCheckResourceAttr(
							refRoute, "target.#", "3"),
					),
				},
			},
		})
}

func testAccCheckRouteExists(resRoute string, validate func() error) resource.TestCheckFunc {

	return func(s *terraform.State) (err error) {

		session := testAccProvider.Meta().(*cfapi.Session)

		rs, ok := s.RootModule().Resources[resRoute]
		if !ok {
			return fmt.Errorf("route '%s' not found in terraform state", resRoute)
		}

		session.Log.DebugMessage(
			"terraform state for resource '%s': %# v",
			resRoute, rs)

		id := rs.Primary.ID
		attributes := rs.Primary.Attributes

		var route cfapi.CCRoute
		rm := session.RouteManager()
		if route, err = rm.ReadRoute(id); err != nil {
			return
		}
		session.Log.DebugMessage(
			"retrieved route for resource '%s' with id '%s': %# v",
			resRoute, id, route)

		if err = assertEquals(attributes, "domain", route.DomainGUID); err != nil {
			return
		}
		if err = assertEquals(attributes, "space", route.SpaceGUID); err != nil {
			return
		}
		if err = assertEquals(attributes, "hostname", route.Hostname); err != nil {
			return
		}
		if err = assertEquals(attributes, "port", route.Port); err != nil {
			return
		}
		if err = assertEquals(attributes, "path", route.Path); err != nil {
			return
		}

		err = validate()
		return
	}
}

func testAccCheckRouteDestroyed(hostnames []string, domain string) resource.TestCheckFunc {

	return func(s *terraform.State) error {

		session := testAccProvider.Meta().(*cfapi.Session)
		for _, h := range hostnames {
			if _, err := session.RouteManager().FindRoute(domain, &h, nil, nil); err != nil {
				switch err.(type) {
				case *errors.ModelNotFoundError:
					continue
				default:
					return err
				}
			}
			return fmt.Errorf("route with hostname '%s' and domain '%s' still exists in cloud foundry", h, domain)
		}
		return nil
	}
}
