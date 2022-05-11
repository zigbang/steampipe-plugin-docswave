package docswave

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableDocswaveMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "docswave_member",
		List: &plugin.ListConfig{
			Hydrate: listMember,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("member_id"),
			Hydrate:    getMember,
		},
		// HydrateDependencies: []plugin.HydrateDependencies{
		// 	{
		// 		Func:    getMemberTeam,
		// 		Depends: []plugin.HydrateFunc{listMember, getMember},
		// 	},
		// },
		Columns: []*plugin.Column{
			{Name: "member_id", Type: proto.ColumnType_STRING},
			{Name: "member_name", Type: proto.ColumnType_STRING},
			{Name: "member_email", Type: proto.ColumnType_STRING},
			{Name: "member_tel_no", Type: proto.ColumnType_STRING},
			{Name: "member_tel_no_office", Type: proto.ColumnType_STRING},
			{Name: "member_no", Type: proto.ColumnType_STRING},
			{Name: "member_address", Type: proto.ColumnType_STRING},
			{Name: "member_image", Type: proto.ColumnType_STRING},
			{Name: "member_language", Type: proto.ColumnType_STRING},
			{Name: "member_time_zone", Type: proto.ColumnType_STRING},
			{Name: "member_country", Type: proto.ColumnType_STRING},
			{Name: "member_status", Type: proto.ColumnType_STRING},
			{Name: "member_entry_date", Type: proto.ColumnType_STRING},
			{Name: "member_role", Type: proto.ColumnType_STRING},
			{Name: "team_id", Type: proto.ColumnType_STRING, Hydrate: getMemberTeam},
			{Name: "created_date", Type: proto.ColumnType_STRING},
			{Name: "update_date", Type: proto.ColumnType_STRING},
		},
	}
}

type Member struct {
	MemberId          string `json:"memberId"`
	MemberName        string `json:"memberName"`
	MemberEmail       string `json:"memberEmail"`
	MemberTelNo       string `json:"memberTelNo"`
	MemberTelNoOffice string `json:"memberTelNoOffice"`
	MemberNo          string `json:"memberNo"`
	MemberAddress     string `json:"memberAddress"`
	MemberImage       string `json:"memberImage"`
	MemberLanguage    string `json:"memberLanguage"`
	MemberTimeZone    string `json:"memberTimeZone"`
	MemberCountry     string `json:"memberCountry"`
	MemberStatus      string `json:"memberStatus"`
	MemberEntryDate   string `json:"memberEntryDate"`
	MemberRole        string `json:"memberRole"`
	TeamId            string `json:"teamId"`
	CreatedDate       string `json:"createdDate"`
	UpdateDate        string `json:"updateDate"`
}

type Team struct {
	TeamId          string `json:"teamId"`
	TeamName        string `json:"teamName"`
	ParentTeamId    string `json:"parentTeamId"`
	ChildTeamModels string `json:"childTeamModels"`
	CreatedDate     string `json:"createdDate"`
	UpdateDate      string `json:"updateDate"`
}

func listMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("Warn-Docswave-listMember-Start:")
	client, token, err := connect(ctx, d)
	if err != nil {
		logger.Error("Error-Docswave-listMember-connect:", err)
		return nil, nil
	}

	url := "https://openapi.docswave.com/members?openApiKey=" + token

	resp, err := client.Get(url)
	if err != nil {
		logger.Error("Error-Docswave-listMember-client.Get:", err)
		return nil, nil
	}
	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error-Docswave-listMember-ioutil.ReadAll:", err)
		return nil, nil
	}

	var members []Member

	json.Unmarshal([]byte(htmlData), &members)

	for _, member := range members {
		d.StreamListItem(ctx, member)
	}

	return nil, nil
}

func getMember(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("Warn-Docswave-getMember-Start:")
	quals := d.KeyColumnQuals
	logger.Trace("getMember", "quals", quals)
	memberId := quals["member_id"].GetStringValue()
	logger.Trace("getMember-memberId", "quals", memberId)

	client, token, err := connect(ctx, d)
	if err != nil {
		logger.Error("Error-Docswave-getMember-connect:", err)
		return nil, nil
	}

	url := "https://openapi.docswave.com/members/" + memberId + "?openApiKey=" + token

	resp, err := client.Get(url)
	if err != nil {
		logger.Error("Error-Docswave-getMember-client.Get:", err)
		return nil, nil
	}
	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error-Docswave-getMember-ioutil.ReadAll:", err)
		return nil, nil
	}

	var member Member

	json.Unmarshal([]byte(htmlData), &member)

	return member, nil
}

func getMemberTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Trace("Warn-Docswave-getMemberTeam-Start:")
	hydrateData := h.HydrateResults
	logger.Trace("getMemberTeam", "hydrateData", hydrateData)
	var memberId string
	if hydrateData["listMember"] != nil {
		if hydrateData["listMember"].(Member).MemberStatus == "WORKING" {
			memberId = hydrateData["listMember"].(Member).MemberId
		} else {
			return nil, nil
		}
	} else {
		if hydrateData["getMember"].(Member).MemberStatus == "WORKING" {
			memberId = hydrateData["getMember"].(Member).MemberId
		} else {
			return nil, nil
		}
	}
	logger.Trace("getMember-memberId", "value", memberId)

	client, token, err := connect(ctx, d)
	if err != nil {
		logger.Error("Error-Docswave-getMemberTeam-connect:", err)
		return nil, nil
	}

	url := "https://openapi.docswave.com/members/" + memberId + "/teams?openApiKey=" + token

	resp, err := client.Get(url)
	if err != nil {
		logger.Error("Error-Docswave-getMemberTeam-client.Get:", err)
		return nil, nil
	}
	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error-Docswave-getMemberTeam-ioutil.ReadAll:", err)
		return nil, nil
	}

	var teams []Team

	json.Unmarshal([]byte(htmlData), &teams)

	logger.Trace("Warn-Docswave-getMemberTeam-Unmarshal:", htmlData, teams)

	if len(teams) == 0 {
		return nil, nil
	} else {
		return teams[0], nil
	}
}
