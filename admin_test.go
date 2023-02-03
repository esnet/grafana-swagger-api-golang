package gapi

import (
	"strings"
	"testing"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/admin"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/admin_users"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
	"github.com/gobs/pretty"
)

const (
	createUserJSON            = `{"id":1,"message":"User created"}`
	deleteUserJSON            = `{"message":"User deleted"}`
	updateUserPasswordJSON    = `{"message":"User password updated"}`
	updateUserPermissionsJSON = `{"message":"User permissions updated"}`

	pauseAllAlertsJSON = `{
		"alertsAffected": 1,
		"state": "Paused",
		"message": "alert paused"
	}`
)

func TestCreateUser(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createUserJSON)
	defer mocksrv.Close()

	resp, err := client.AdminUsers.AdminCreateUser(admin_users.NewAdminCreateUserParams().WithBody(
		&models.AdminCreateUserForm{
			Email:    "admin@localhost",
			Login:    "admin",
			Name:     "Administrator",
			Password: "password",
		},
	), nil)
	if err != nil {
		t.Error(err)
	}

	if resp.Payload.ID != 1 {
		t.Error("Not correctly parsing returned user message.")
	}
}

func TestDeleteUser(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, deleteUserJSON)
	defer mocksrv.Close()

	_, err := client.AdminUsers.AdminDeleteUser(admin_users.NewAdminDeleteUserParams().WithUserID(1), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, updateUserPasswordJSON)
	defer mocksrv.Close()

	_, err := client.AdminUsers.AdminUpdateUserPassword(admin_users.NewAdminUpdateUserPasswordParams().
		WithUserID(1).
		WithBody(&models.AdminUpdateUserPasswordForm{
			Password: "new-password",
		},
		), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUserPermissions(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, updateUserPermissionsJSON)
	defer mocksrv.Close()

	_, err := client.AdminUsers.AdminUpdateUserPermissions(admin_users.NewAdminUpdateUserPermissionsParams().
		WithUserID(1).
		WithBody(&models.AdminUpdateUserPermissionsForm{
			IsGrafanaAdmin: false,
		},
		), nil)
	if err != nil {
		t.Error(err)
	}
}

func TestPauseAllAlerts(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, pauseAllAlertsJSON)
	defer mocksrv.Close()

	res, err := client.Admin.PauseAllAlerts(admin.NewPauseAllAlertsParams().WithBody(
		&models.PauseAllAlertsCommand{
			Paused: true,
		},
	), nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))

	if res.Payload.State != "Paused" {
		t.Error("pause all alerts response should contain the correct response message")
	}
}

func TestPauseAllAlerts_500(t *testing.T) {
	server, client := gapiTestTools(t, 500, pauseAllAlertsJSON)
	defer server.Close()

	_, err := client.Admin.PauseAllAlerts(admin.NewPauseAllAlertsParams(), nil)
	if !strings.Contains(*err.(*admin.PauseAllAlertsInternalServerError).Payload.Message, "alert paused") {
		t.Errorf("expected error to contain 'status: 500'; got: %s", err.Error())
	}

}
