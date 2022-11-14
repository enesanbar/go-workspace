package main

import (
	"github.com/enesanbar/terraform-provider-petstore/petstore"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return petstore.Provider()
		},
	})
}
