package gapi

import (
	"testing"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/snapshots"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
	"github.com/gobs/pretty"
)

const (
	createdSnapshotResponse = `{
		"deleteKey":"XXXXXXX",
		"deleteUrl":"myurl/api/snapshots-delete/XXXXXXX",
		"key":"YYYYYYY",
		"url":"myurl/dashboard/snapshot/YYYYYYY",
		"id": 1
	}`
)

func TestSnapshotCreate(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createdSnapshotResponse)
	defer mocksrv.Close()

	snapshot := models.CreateDashboardSnapshotCommand{
		Dashboard: map[string]interface{}{
			"title": "test",
		},
		Expires: 3600,
	}

	resp, err := client.Snapshots.CreateDashboardSnapshot(
		snapshots.NewCreateDashboardSnapshotParams().
			WithBody(&snapshot),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.Payload.DeleteKey != "XXXXXXX" {
		t.Errorf("Invalid key - %s, Expected %s", resp.Payload.DeleteKey, "XXXXXXX")
	}

	for _, code := range []int{400, 401, 403, 412} {
		mocksrv.code = code
		_, err := client.Snapshots.CreateDashboardSnapshot(
			snapshots.NewCreateDashboardSnapshotParams().
				WithBody(&snapshot),
			nil,
		)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}
