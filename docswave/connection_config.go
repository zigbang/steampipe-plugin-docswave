package docswave

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type docswaveConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {Type: schema.TypeString},
}

func ConfigInstance() interface{} {
	return &docswaveConfig{}
}

func GetConfig(connection *plugin.Connection) docswaveConfig {
	if connection == nil || connection.Config == nil {
		return docswaveConfig{}
	}
	config := connection.Config.(docswaveConfig)
	return config
}
