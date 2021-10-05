package ketch

import (
	"context"
	"log"

	"github.com/brunoa19/ketch-pulumi-provider/provider/terraform/client"
	"github.com/brunoa19/ketch-pulumi-provider/provider/terraform/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var (
	processesSchema = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"cmd": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	routingSettingsSchema = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"weight": {
					Type:     schema.TypeInt,
					Optional: true,
				},
			},
		},
	}

	schemaApp = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				// Required
				"name": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
				"image": {
					Type:     schema.TypeString,
					Required: true,
				},
				"framework": {
					Type:     schema.TypeString,
					Required: true,
				},

				// Optional
				"cnames": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"port": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeInt,
					},
				},
				"units": {
					Type:     schema.TypeInt,
					Optional: true,
				},

				"processes": processesSchema,

				"routing_settings": routingSettingsSchema,

				"version": {
					Type:     schema.TypeInt,
					Optional: true,
				},
			},
		},
	}
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,
		Schema: map[string]*schema.Schema{
			"app": schemaApp,
		},
	}
}

func resourceAppCreate(d *schema.ResourceData, m interface{}) error {
	raw := d.Get("app").([]interface{})[0].(map[string]interface{})
	app := &client.App{}

	helper.TerraformToStruct(raw, app)

	c := m.(*client.Client)
	err := c.CreateApp(context.Background(), app)
	if err != nil {
		return err
	}

	d.SetId(app.Name)

	resourceAppRead(d, m)

	return nil
}

func resourceAppRead(d *schema.ResourceData, m interface{}) error {
	name := d.Id()

	c := m.(*client.Client)
	app, err := c.GetApp(context.Background(), name)
	if err != nil {
		return err
	}

	if err = d.Set("app", helper.StructToTerraform(app)); err != nil {
		return err
	}

	return nil
}

func resourceAppUpdate(d *schema.ResourceData, m interface{}) error {
	if !d.HasChange("app") {
		return resourceAppRead(d, m)
	}

	raw := d.Get("app").([]interface{})[0].(map[string]interface{})
	app := &client.App{}
	helper.TerraformToStruct(raw, app)

	log.Printf(" ### RAW app data: %+v\n", raw)
	log.Printf(" ### CONVERTED app data: %+v\n", *app)

	c := m.(*client.Client)
	err := c.UpdateApp(context.Background(), app)
	if err != nil {
		return err
	}

	return resourceAppRead(d, m)
}

func resourceAppDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Id()
	c := m.(*client.Client)
	err := c.DeleteApp(context.Background(), name)
	if err != nil {
		return err
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return nil
}
