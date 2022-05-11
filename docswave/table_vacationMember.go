package docswave

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableDocswaveVacationMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "docswave_vacation_member",
		List: &plugin.ListConfig{
			Hydrate: listVacationMember,
		},
		Columns: []*plugin.Column{
			{Name: "member_name", Type: proto.ColumnType_STRING},
			{Name: "vacation_item_name", Type: proto.ColumnType_STRING},
			{Name: "vacation_start_date", Type: proto.ColumnType_STRING},
			{Name: "vacation_end_date", Type: proto.ColumnType_STRING},
		},
	}
}

type VacationMember struct {
	MemberName        string `json:"memberName"`
	VacationItemName  string `json:"vacationItemName"`
	VacationStartDate string `json:"vacationStartDate"`
	VacationEndDate   string `json:"vacationEndDate"`
}

func listVacationMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	client, token, err := connect(ctx, d)
	if err != nil {
		logger.Error("Error-Docswave-listVacationMember-connect:", err)
		return nil, nil
	}

	url := "https://openapi.docswave.com/vacationMembers/zigbang?openApiKey=" + token

	resp, err := client.Get(url)
	if err != nil {
		logger.Error("Error-Docswave-listVacationMember-client.Get:", err)
		return nil, nil
	}
	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error-Docswave-listVacationMember-ioutil.ReadAll:", err)
		return nil, nil
	}

	var vacationMembers []VacationMember

	json.Unmarshal([]byte(htmlData), &vacationMembers)

	for _, member := range vacationMembers {
		d.StreamListItem(ctx, member)
	}

	return nil, nil
}
