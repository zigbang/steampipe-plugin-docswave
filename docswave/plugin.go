package docswave

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-docswave",
		DefaultTransform: transform.FromCamel(),
		DefaultConcurrency: &plugin.DefaultConcurrencyConfig{
			TotalMaxConcurrency: 10,
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"docswave_team":            tableDocswaveTeam(ctx),
			"docswave_member":          tableDocswaveMember(ctx),
			"docswave_vacation_member": tableDocswaveVacationMember(ctx),
		},
	}
	return p
}
