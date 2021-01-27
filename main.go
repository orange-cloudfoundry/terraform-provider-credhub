package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/orange-cloudfoundry/terraform-provider-credhub/credhub"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Version(version.Print("terraform-provider-credhub"))
	kingpin.HelpFlag.Short('h')
	usage := strings.ReplaceAll(kingpin.DefaultUsageTemplate, "usage: ", "usage: CLOUD_FILE=config.yml ")
	kingpin.UsageTemplate(usage)
	kingpin.Parse()
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: credhub.Provider})
}
