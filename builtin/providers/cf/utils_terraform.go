package cloudfoundry

import (
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
)

// getResourceChange -
func getResourceChange(key string, d *schema.ResourceData) (bool, string, string) {
	old, new := d.GetChange(key)
	return old != new, old.(string), new.(string)
}

// getListChanges -
func getListChanges(old interface{}, new interface{}) (remove []string, add []string) {

	var a bool

	for _, o := range old.(*schema.Set).List() {
		remove = append(remove, o.(string))
	}
	for _, n := range new.(*schema.Set).List() {
		nn := n.(string)
		a = true
		for i, r := range remove {
			if nn == r {
				remove = append(remove[:i], remove[i+1:]...)
				a = false
				break
			}
		}
		if a {
			add = append(add, nn)
		}
	}
	return
}

// getListChangedSchemaLists -
func getListChangedSchemaLists(old interface{}, new interface{}) (remove []map[string]interface{}, add []map[string]interface{}) {

	var a bool

	for _, o := range old.(*schema.Set).List() {
		remove = append(remove, o.(map[string]interface{}))
	}
	for _, n := range new.(*schema.Set).List() {
		nn := n.(map[string]interface{})
		a = true
		for i, r := range remove {
			if reflect.DeepEqual(nn, r) {
				remove = append(remove[:i], remove[i+1:]...)
				a = false
				break
			}
		}
		if a {
			add = append(add, nn)
		}
	}
	return
}
