package cloudfoundry

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/builtin/providers/cloudfoundry/cfapi"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

var pcfDevOrgID string

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"cf": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {

	if !testAccEnvironmentSet() {
		t.Fatal("Acceptance environment has not been set.")
	}
}

func testAccEnvironmentSet() bool {

	endpoint := os.Getenv("CF_API_URL")
	user := os.Getenv("CF_USER")
	password := os.Getenv("CF_PASSWORD")
	uaaClientID := os.Getenv("CF_UAA_CLIENT_ID")
	uaaClientSecret := os.Getenv("CF_UAA_CLIENT_SECRET")
	skipSslValidation := strings.ToLower(os.Getenv("CF_SKIP_SSL_VALIDATION"))

	if endpoint == "" ||
		user == "" ||
		password == "" ||
		uaaClientID == "" ||
		uaaClientSecret == "" ||
		skipSslValidation == "" {

		fmt.Println("CF_API_URL, CF_USER, CF_PASSWORD, CF_UAA_CLIENT_ID, CF_UAA_CLIENT_SECRET " +
			"and CF_SKIP_SSL_VALIDATION must be set for acceptance tests to work.")
		return false

	}
	return true
}

func defaultPcfDevOrgID() string {

	if len(pcfDevOrgID) == 0 && testAccEnvironmentSet() {
		var (
			err       error
			pcfDevOrg cfapi.CCOrg
		)
		meta := testAccProvider.Meta()
		if meta != nil {
			session := meta.(*cfapi.Session)
			if pcfDevOrg, err = session.OrgManager().FindOrg("pcfdev-org"); err != nil {
				panic(err.Error())
			}
			pcfDevOrgID = pcfDevOrg.ID
		}
	}
	return pcfDevOrgID
}

func assertEquals(attributes map[string]string,
	key string, expected interface{}) error {
	v := attributes[key]
	if v != expected {
		return fmt.Errorf("expected resource '%s' to be '%s' but it was '%s'", key, v, expected)
	}
	return nil
}

func assertListEquals(attributes map[string]string,
	key string, actualLen int,
	match func(map[string]string, int) bool) (err error) {

	var n int

	num := attributes[key+".#"]
	if num != "" {
		n, err = strconv.Atoi(num)
		if err != nil {
			return
		}
	} else {
		n = 0
	}

	if actualLen > 0 && n == 0 {
		return fmt.Errorf(
			"expected resource '%s' to be empty but it has '%d' elements", key, actualLen)
	}
	if actualLen != n {
		return fmt.Errorf(
			"expected resource '%s' to have '%d' elements but it has '%d' elements",
			key, n, actualLen)
	}
	if n > 0 {
		found := 0

		var (
			values map[string]string
			ok     bool
		)

		keyValues := make(map[string]map[string]string)
		for k, v := range attributes {
			keyParts := strings.Split(k, ".")
			if key == keyParts[0] && keyParts[1] != "#" {
				i := keyParts[1]
				if values, ok = keyValues[i]; !ok {
					values = make(map[string]string)
					keyValues[i] = values
				}
				if len(keyParts) == 2 {
					values["value"] = v
				} else {
					values[strings.Join(keyParts[2:], ".")] = v
				}
			}
		}

		for _, values := range keyValues {
			for j := 0; j < actualLen; j++ {
				if match(values, j) {
					found++
				}
			}
		}
		if n != found {
			return fmt.Errorf(
				"expected list resource '%s' to match '%d' elements but matched only '%d' elements",
				key, n, found)
		}
	}
	return nil
}

func assertSetEquals(attributes map[string]string,
	key string, expected []interface{}) (err error) {

	var n int

	num := attributes[key+".#"]
	if num != "" {
		n, err = strconv.Atoi(num)
		if err != nil {
			return
		}
	} else {
		n = 0
	}

	if len(expected) > 0 && n == 0 {
		return fmt.Errorf(
			"expected resource '%s' to be '%v' but it was empty", key, expected)
	}
	if len(expected) != n {
		return fmt.Errorf(
			"expected resource '%s' to have '%d' elements but it has '%d' elements",
			key, len(expected), n)
	}
	if n > 0 {
		found := 0
		for _, e := range expected {
			if _, ok := attributes[key+"."+strconv.Itoa(hashcode.String(e.(string)))]; ok {
				found++
			}
		}
		if n != found {
			return fmt.Errorf(
				"expected set resource '%s' to have elements '%v' but matched only '%d' elements",
				key, expected, found)
		}
	}
	return
}

func asserMapEquals(key string, attributes map[string]string, actual map[string]interface{}) (err error) {

	expected := make(map[string]interface{})
	for k, v := range attributes {
		keyParts := strings.Split(k, ".")
		if keyParts[0] == key && keyParts[1] != "%" {

			l := len(keyParts)
			m := expected
			for _, kk := range keyParts[1 : l-1] {
				if _, ok := m[kk]; !ok {
					m[kk] = make(map[string]interface{})
				}
				m = m[kk].(map[string]interface{})
			}
			m[keyParts[l-1]] = v
		}
	}
	if !reflect.DeepEqual(expected, actual) {
		err = fmt.Errorf("map with key '%s' expected to be %#v but was %#v", key, expected, actual)
	}
	return nil
}
