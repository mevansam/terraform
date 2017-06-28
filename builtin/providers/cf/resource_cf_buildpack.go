package cloudfoundry

import (
	"fmt"

	"github.com/hashicorp/terraform/builtin/providers/cf/cfapi"
	"github.com/hashicorp/terraform/builtin/providers/cf/repo"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceBuildpack() *schema.Resource {

	return &schema.Resource{

		Create: resourceBuildpackCreate,
		Read:   resourceBuildpackRead,
		Update: resourceBuildpackUpdate,
		Delete: resourceBuildpackDelete,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"position": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"locked": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"git", "github_release"},
			},
			"git": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"url", "github_release"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"branch": &schema.Schema{
							Type:          schema.TypeString,
							Optional:      true,
							Default:       "master",
							ConflictsWith: []string{"git.tag"},
						},
						"tag": &schema.Schema{
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"git.branch"},
						},
						"user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"github_release": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				MaxItems:      1,
				ConflictsWith: []string{"url", "git"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"owner": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"repo": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"token": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"filename": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceBuildpackCreate(d *schema.ResourceData, meta interface{}) (err error) {

	session := meta.(*cfapi.Session)
	if session == nil {
		return fmt.Errorf("client is nil")
	}

	var (
		name            string
		position        *int
		enabled, locked *bool

		path       string
		repository repo.Repository

		bp cfapi.CCBuildpack
	)
	name = d.Get("name").(string)
	if v, ok := d.GetOk("position"); ok {
		s := v.(int)
		position = &s
	}
	if v, ok := d.GetOk("enabled"); ok {
		b := v.(bool)
		enabled = &b
	}
	if v, ok := d.GetOk("locked"); ok {
		b := v.(bool)
		locked = &b
	}

	if v, ok := d.GetOk("url"); ok {
		path = v.(string)
	} else {
		if repository, err = getRepositoryFromConfig(d); err != nil {
			return
		}
		path = repository.GetPath()
	}
	if bp, err = session.BuildpackManager().CreateBuildpack(name, position, enabled, locked, path); err != nil {
		return
	}

	d.SetId(bp.ID)
	d.Set("position", bp.Position)
	d.Set("enabled", bp.Enabled)
	d.Set("locked", bp.Locked)

	return
}

func resourceBuildpackRead(d *schema.ResourceData, meta interface{}) (err error) {

	session := meta.(*cfapi.Session)
	if session == nil {
		return fmt.Errorf("client is nil")
	}

	id := d.Id()
	bpm := session.BuildpackManager()

	var bp cfapi.CCBuildpack
	if bp, err = bpm.ReadBuildpack(id); err != nil {
		return
	}

	d.Set("name", bp.Name)
	d.Set("position", bp.Position)
	d.Set("enabled", bp.Enabled)
	d.Set("locked", bp.Locked)

	return
}

func resourceBuildpackUpdate(d *schema.ResourceData, meta interface{}) (err error) {

	session := meta.(*cfapi.Session)
	if session == nil {
		return fmt.Errorf("client is nil")
	}

	id := d.Id()
	bpm := session.BuildpackManager()

	var (
		name            string
		position        *int
		enabled, locked *bool

		update bool

		path       string
		repository repo.Repository

		bp cfapi.CCBuildpack
	)

	if d.HasChange("name") {
		name = d.Get("name").(string)
		update = true
	} else {
		name = d.Get("name").(string)
	}
	if d.HasChange("position") {
		s := d.Get("position").(int)
		position = &s
		update = true
	} else if v, ok := d.GetOk("position"); ok {
		s := v.(int)
		position = &s
	}
	if d.HasChange("enabled") {
		b := d.Get("enabled").(bool)
		enabled = &b
		update = true
	} else if v, ok := d.GetOk("enabled"); ok {
		s := v.(bool)
		enabled = &s
	}
	if d.HasChange("locked") {
		b := d.Get("locked").(bool)
		locked = &b
		update = true
	} else if v, ok := d.GetOk("locked"); ok {
		s := v.(bool)
		locked = &s
	}
	if update {
		if bp, err = bpm.UpdateBuildpack(id, name, position, enabled, locked); err != nil {
			return
		}
		d.Set("position", bp.Position)
		d.Set("enabled", bp.Enabled)
		d.Set("locked", bp.Locked)
	} else {
		bp.Name = name
		bp.Position = position
		bp.Enabled = enabled
		bp.Locked = locked
	}
	if d.HasChange("url") || d.HasChange("git") || d.HasChange("github_release") {

		if v, ok := d.GetOk("url"); ok {
			path = v.(string)
		} else {
			if repository, err = getRepositoryFromConfig(d); err != nil {
				return
			}
			path = repository.GetPath()
		}
		if bp, err = session.BuildpackManager().UploadBuildpackBits(bp, path); err != nil {
			return
		}
	}
	return
}

func resourceBuildpackDelete(d *schema.ResourceData, meta interface{}) (err error) {

	session := meta.(*cfapi.Session)
	if session == nil {
		return fmt.Errorf("client is nil")
	}

	err = session.BuildpackManager().DeleteBuildpack(d.Id())
	return
}
