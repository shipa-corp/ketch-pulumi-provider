package ketch

import (
	"context"
	"log"

	"github.com/brunoa19/ketch-pulumi-provider/provider/terraform/client"
	"github.com/brunoa19/ketch-pulumi-provider/provider/terraform/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var (
	schemaIngressController = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"class_name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"service_endpoint": {
					Type:     schema.TypeString,
					Required: true,
				},
				"type": {
					Type:     schema.TypeString,
					Required: true,
				},
				"cluster_issuer": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
)

func resourceFramework() *schema.Resource {
	return &schema.Resource{
		Create: resourceFrameworkCreate,
		Read:   resourceFrameworkRead,
		Update: resourceFrameworkUpdate,
		Delete: resourceFrameworkDelete,
		Schema: map[string]*schema.Schema{
			"framework": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"app_quota_limit": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ingress_controller": schemaIngressController,
					},
				},
			},
		},
	}
}

func resourceFrameworkCreate(d *schema.ResourceData, m interface{}) error {
	raw := d.Get("framework").([]interface{})[0].(map[string]interface{})

	framework := &client.Framework{}
	helper.TerraformToStruct(raw, framework)

	log.Printf("RAW framework: %+v\n", raw)
	log.Printf("CONVERTED framework: %+v\n", *framework)

	c := m.(*client.Client)
	err := c.CreateFramework(context.Background(), framework)
	if err != nil {
		return err
	}

	d.SetId(framework.Name)

	resourceFrameworkRead(d, m)

	return nil
}

func resourceFrameworkRead(d *schema.ResourceData, m interface{}) error {
	name := d.Id()

	c := m.(*client.Client)
	framework, err := c.GetFramework(context.Background(), name)
	if err != nil {
		return err
	}

	if err = d.Set("framework", helper.StructToTerraform(framework)); err != nil {
		return err
	}

	return nil
}

func resourceFrameworkUpdate(d *schema.ResourceData, m interface{}) error {
	if !d.HasChange("framework") {
		return resourceFrameworkRead(d, m)
	}

	raw := d.Get("framework").([]interface{})[0].(map[string]interface{})

	framework := &client.Framework{}
	helper.TerraformToStruct(raw, framework)

	log.Printf("RAW framework: %+v\n", raw)

	c := m.(*client.Client)
	err := c.UpdateFramework(context.Background(), framework)
	if err != nil {
		return err
	}

	return resourceFrameworkRead(d, m)
}

func resourceFrameworkDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Id()
	c := m.(*client.Client)
	err := c.DeleteFramework(context.Background(), name)
	if err != nil {
		return err
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return nil
}
