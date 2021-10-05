package ketch

import (
	"fmt"

	"github.com/brunoa19/ketch-pulumi-provider/provider/terraform/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"ketch_app":       resourceApp(),
			"ketch_job":       resourceJob(),
			"ketch_framework": resourceFramework(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c, err := client.NewClient()
	if err != nil {
		return nil, fmt.Errorf("unable to create Ketch client: %s", err.Error())
	}

	return c, nil
}
