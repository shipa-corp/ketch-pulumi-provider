package ketch

import (
	"context"
	"log"

	"github.com/brunoa19/ketch-pulumi-provider/provider/terraform/client"
	"github.com/brunoa19/ketch-pulumi-provider/provider/terraform/helper"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var (
	schemaJobContainers = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"image": {
					Type:     schema.TypeString,
					Required: true,
				},
				"command": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	schemaJobPolicy = &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"restart_policy": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
)

func resourceJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobCreate,
		Read:   resourceJobRead,
		Update: resourceJobUpdate,
		Delete: resourceJobDelete,
		Schema: map[string]*schema.Schema{
			"job": {
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
						"type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"framework": {
							Type:     schema.TypeString,
							Required: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"parallelism": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"completions": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"suspend": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"backoff_limit": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"containers": schemaJobContainers,
						"policy":     schemaJobPolicy,
					},
				},
			},
		},
	}
}

func resourceJobCreate(d *schema.ResourceData, m interface{}) error {
	raw := d.Get("job").([]interface{})[0].(map[string]interface{})

	job := &client.Job{}
	helper.TerraformToStruct(raw, job)

	log.Printf("RAW job: %+v\n", raw)
	log.Printf("CONVERTED job: %+v\n", *job)

	c := m.(*client.Client)
	err := c.CreateJob(context.Background(), job)
	if err != nil {
		return err
	}

	d.SetId(job.Name)

	resourceJobRead(d, m)

	return nil
}

func resourceJobRead(d *schema.ResourceData, m interface{}) error {
	name := d.Id()

	c := m.(*client.Client)
	job, err := c.GetJob(context.Background(), name)
	if err != nil {
		return err
	}

	return d.Set("job", helper.StructToTerraform(job))
}

func resourceJobUpdate(d *schema.ResourceData, m interface{}) error {
	if !d.HasChange("job") {
		return resourceJobRead(d, m)
	}

	raw := d.Get("job").([]interface{})[0].(map[string]interface{})

	job := &client.Job{}
	helper.TerraformToStruct(raw, job)

	log.Printf("RAW job: %+v\n", raw)

	c := m.(*client.Client)
	err := c.UpdateJob(context.Background(), job)
	if err != nil {
		return err
	}

	return resourceJobRead(d, m)
}

func resourceJobDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Id()
	c := m.(*client.Client)
	err := c.DeleteJob(context.Background(), name)
	if err != nil {
		return err
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return nil
}
