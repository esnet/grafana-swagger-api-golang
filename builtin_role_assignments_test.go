package gapi

import (
	"github.com/google/go-cmp/cmp"
	"testing"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/access_control"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

const (
	newBuiltInRoleAssignmentResponse = `
{
    "message": "Built-in role grant added"
}
`
	getBuiltInRoleAssignmentsResponse = `
	[
    	{
				"displayName": "Grafana Admin",
    	        "description": "",
    	        "name": "grafana:roles:users:admin:read",
    	        "uid": "tJTyTNqMk",
    	        "version": 1
    	},
		{
				"displayName": "Viewer",
    	        "description": "Role to allow users to create/read reports",
    	        "name": "custom:reports:editor",
    	        "uid": "tJTyTNqMk1",
    	        "version": 2
    	}

]
`

	removeBuiltInRoleAssignmentResponse = `
{
    "message": "Built-in role grant removed"
}
`
)

func TestNewBuiltInRoleAssignment(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, newBuiltInRoleAssignmentResponse)
	t.Cleanup(func() {
		mocksrv.Close()
	})

	br := models.AddUserRoleCommand{
		Global:  false,
		RoleUID: "test:policy",
	}

	_, err := client.AccessControl.AddUserRole(access_control.NewAddUserRoleParams().WithBody(&br), nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetBuiltInRoleAssignments(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getBuiltInRoleAssignmentsResponse)
	t.Cleanup(func() {
		mocksrv.Close()
	})

	resp, err := client.AccessControl.ListRoles(access_control.NewListRolesParams(), nil)

	if err != nil {
		t.Error(err)
	}

	expected := map[string][]*models.RoleDTO{
		"Grafana Admin": {
			{
				Version:     1,
				DisplayName: "Grafana Admin",
				Name:        "grafana:roles:users:admin:read",
				UID:         "tJTyTNqMk",
				Description: "",
			},
		},
		"Viewer": {
			{
				Version:     2,
				DisplayName: "Viewer",
				Name:        "custom:reports:editor",
				UID:         "tJTyTNqMk1",
				Description: "Role to allow users to create/read reports",
			},
		},
	}
	if cmp.Diff(expected["Grafana Admin"][0], resp.Payload[0]) != "" {
		t.Error("Unexpected built-in role assignments.")
	}
	if cmp.Diff(expected["Viewer"][0], resp.Payload[1]) != "" {
		t.Error("Unexpected built-in role assignments.")
	}

}

func TestDeleteBuiltInRoleAssignment(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, removeBuiltInRoleAssignmentResponse)
	t.Cleanup(func() {
		mocksrv.Close()
	})

	global := false

	_, err := client.AccessControl.RemoveUserRole(
		access_control.NewRemoveUserRoleParams().
			WithRoleUID("test:policy").
			WithGlobal(&global),
		nil,
	)

	if err != nil {
		t.Error(err)
	}
}
