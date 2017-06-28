package cloudfoundry

import (
	"fmt"

	"github.com/hashicorp/terraform/builtin/providers/cf/cfapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceApp() *schema.Resource {

	return &schema.Resource{

		Create: resourceAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"quota": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAppCreate(d *schema.ResourceData, meta interface{}) (err error) {

	session := meta.(*cfapi.Session)
	if session == nil {
		return fmt.Errorf("client is nil")
	}

	var (
		name, quota string
		org         cfapi.CCOrg
	)
	name = d.Get("name").(string)
	if v, ok := d.GetOk("quota"); ok {
		quota = v.(string)
	}

	om := session.OrgManager()
	if org, err = om.CreateOrg(name, quota); err != nil {
		return err
	}
	if len(quota) == 0 {
		d.Set("quota", org.QuotaGUID)
	}
	d.SetId(org.ID)
	return resourceAppUpdate(d, NewResourceMeta{meta})
}

func resourceAppRead(d *schema.ResourceData, meta interface{}) (err error) {

	session := meta.(*cfapi.Session)
	if session == nil {
		return fmt.Errorf("client is nil")
	}

	id := d.Id()
	om := session.OrgManager()

	var org cfapi.CCOrg
	if org, err = om.ReadOrg(id); err != nil {
		return
	}

	d.Set("name", org.Name)
	d.Set("quota", org.QuotaGUID)
	return
}

func resourceAppUpdate(d *schema.ResourceData, meta interface{}) (err error) {

	var (
		newResource bool
		session     *cfapi.Session
	)

	if m, ok := meta.(NewResourceMeta); ok {
		session = m.meta.(*cfapi.Session)
		newResource = true
	} else {
		session = meta.(*cfapi.Session)
		if session == nil {
			return fmt.Errorf("client is nil")
		}
		newResource = false
	}

	id := d.Id()
	om := session.OrgManager()

	if !newResource {

		org := cfapi.CCOrg{
			ID:   id,
			Name: d.Get("name").(string),
		}
		if v, ok := d.GetOk("quota"); ok {
			org.QuotaGUID = v.(string)
		}

		if err = om.UpdateOrg(org); err != nil {
			return err
		}
	}

	return
}

func resourceAppDelete(d *schema.ResourceData, meta interface{}) (err error) {

	session := meta.(*cfapi.Session)
	if session == nil {
		return fmt.Errorf("client is nil")
	}

	err = session.OrgManager().DeleteOrg(d.Id())
	return
}
