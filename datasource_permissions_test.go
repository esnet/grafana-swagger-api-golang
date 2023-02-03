package gapi

import (
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/datasource_permissions"
	"github.com/gobs/pretty"
	"github.com/google/go-cmp/cmp"
	"testing"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/dashboard_permissions"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

const (
	getDatasourcePermissionsJSON = `{
	"datasourceId": 1,
	"enabled": true,
	"permissions": [
		{
			"datasourceId": 1,
			"userId": 1,
			"userLogin": "user",
			"userEmail": "user@test.com",
			"userAvatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56",
			"permission": 1,
			"permissionName": "Query"
		},
		{
			"datasourceId": 2,
			"teamId": 1,
			"team": "A Team",
			"teamAvatarUrl": "/avatar/46d229b033af06a191ff2267bca9ae56",
			"permission": 1,
			"permissionName": "Query"
		}
	]
}`
	addDatasourcePermissionsJSON = `{
	"message": "Datasource permission added"
}`
)

func TestDatasourcePermissions(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDatasourcePermissionsJSON)
	defer server.Close()

	resp, err := client.DatasourcePermissions.GetAllPermissions(
		datasource_permissions.NewGetAllPermissionsParams().WithDatasourceID("1"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expects := []*models.DataSourcePermissionsDTO{
		{
			DatasourceID: 1,
			Enabled:      true,
			Permissions: []*models.DataSourcePermissionRuleDTO{
				{
					UserID:         1,
					DatasourceID:   1,
					TeamID:         0,
					Permission:     1,
					PermissionName: "Query",
					UserAvatarURL:  "/avatar/46d229b033af06a191ff2267bca9ae56",
					UserEmail:      "user@test.com",
					UserLogin:      "user",
				},
			},
		},
		{
			DatasourceID: 2,
			Permissions: []*models.DataSourcePermissionRuleDTO{
				{
					UserID:         0,
					TeamID:         1,
					DatasourceID:   2,
					Permission:     1,
					Team:           "A Team",
					PermissionName: "Query",
					TeamAvatarURL:  "/avatar/46d229b033af06a191ff2267bca9ae56",
				},
			},
		},
	}
	for i := 0; i < 2; i++ {
		if cmp.Diff(resp.Payload.Permissions[i], expects[i].Permissions[0]) != "" {
			t.Error("Not correctly parsing returned datasource permission")
		}
	}

}

func TestAddDatasourcePermissions(t *testing.T) {
	server, client := gapiTestTools(t, 200, addDatasourcePermissionsJSON)
	defer server.Close()

	items := models.UpdateDashboardACLCommand{
		Items: []*models.DashboardACLUpdateItem{
			{
				TeamID:     1,
				Permission: 1,
			},
			{
				UserID:     11,
				Permission: 1,
			},
		},
	}
	_, err := client.DashboardPermissions.UpdateDashboardPermissionsByUID(
		dashboard_permissions.NewUpdateDashboardPermissionsByUIDParams().
			WithUID("uid").WithBody(&items),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
