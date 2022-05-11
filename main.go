package main

import (
	"github.com/zigbang/steampipe-plugin-docswave/docswave"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: docswave.Plugin})
}
