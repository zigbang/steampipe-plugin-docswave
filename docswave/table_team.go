package docswave

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableDocswaveTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "docswave_team",
		List: &plugin.ListConfig{
			Hydrate: listTeam,
		},
		Columns: []*plugin.Column{
			{Name: "team_id", Type: proto.ColumnType_STRING},
			{Name: "team_name", Type: proto.ColumnType_STRING},
			{Name: "parent_team_id", Type: proto.ColumnType_STRING},
			{Name: "child_team_models", Type: proto.ColumnType_STRING},
			{Name: "created_date", Type: proto.ColumnType_STRING},
			{Name: "update_date", Type: proto.ColumnType_STRING},
		},
	}
}

func listTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	client, token, err := connect(ctx, d)
	if err != nil {
		logger.Error("Error-Docswave-listTeam-connect:", err)
		return nil, nil
	}

	url := "https://openapi.docswave.com/teams?openApiKey=" + token

	resp, err := client.Get(url)
	if err != nil {
		logger.Error("Error-Docswave-listTeam-client.Get:", err)
		return nil, nil
	}
	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error-Docswave-listTeam-ioutil.ReadAll:", err)
		return nil, nil
	}

	var teams []Team

	json.Unmarshal([]byte(htmlData), &teams)

	for _, team := range teams {
		d.StreamListItem(ctx, team)
	}

	return nil, nil
}
