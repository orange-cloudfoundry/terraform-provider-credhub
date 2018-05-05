package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/orange-cloudfoundry/terraform-provider-credhub/credhub"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: credhub.Provider})
}
