package gapi

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/orgs"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
	"github.com/gobs/pretty"
)

const (
	getOrgUsersJSON   = `[{"orgID":1,"userID":1,"email":"admin@localhost","avatarUrl":"/avatar/46d229b033af06a191ff2267bca9ae56","login":"admin","role":"Admin","lastSeenAt":"2018-06-28T14:16:11Z","lastSeenAtAge":"\u003c 1m"}]`
	addOrgUserJSON    = `{"message":"User added to organization"}`
	updateOrgUserJSON = `{"message":"Organization user updated"}`
	removeOrgUserJSON = `{"message":"User removed from organization"}`
)

func TestOrgUsersCurrent(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getOrgUsersJSON)
	defer mocksrv.Close()

	resp, err := client.Org.GetOrgUsersForCurrentOrg(nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	user := &models.OrgUserDTO{
		OrgID:  1,
		UserID: 1,
		Email:  "admin@localhost",
		Login:  "admin",
		Role:   "Admin",
	}

	respUser := resp.Payload[0]
	assert.Equal(t, respUser.OrgID, user.OrgID)
	assert.Equal(t, respUser.UserID, user.UserID)
	assert.Equal(t, respUser.Email, user.Email)
	assert.Equal(t, respUser.Login, user.Login)
	assert.Equal(t, respUser.Role, user.Role)

}

func TestOrgUsers(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getOrgUsersJSON)
	defer mocksrv.Close()

	resp, err := client.Orgs.GetOrgUsers(
		orgs.NewGetOrgUsersParams().WithOrgID(1),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	user := &models.OrgUserDTO{
		OrgID:  1,
		UserID: 1,
		Email:  "admin@localhost",
		Login:  "admin",
		Role:   "Admin",
	}

	respUser := resp.Payload[0]
	assert.Equal(t, respUser.OrgID, user.OrgID)
	assert.Equal(t, respUser.UserID, user.UserID)
	assert.Equal(t, respUser.Email, user.Email)
	assert.Equal(t, respUser.Login, user.Login)
	assert.Equal(t, respUser.Role, user.Role)
}

func TestAddOrgUser(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, addOrgUserJSON)
	defer mocksrv.Close()

	_, err := client.Orgs.AddOrgUser(
		orgs.NewAddOrgUserParams().
			WithOrgID(1).
			WithBody(&models.AddOrgUserCommand{
				LoginOrEmail: "admin@localhost",
				Role:         "Admin",
			}),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateOrgUser(t *testing.T) {
	server, client := gapiTestTools(t, 200, updateOrgUserJSON)
	defer server.Close()

	_, err := client.Orgs.UpdateOrgUser(
		orgs.NewUpdateOrgUserParams().
			WithOrgID(1).
			WithUserID(1).
			WithBody(&models.UpdateOrgUserCommand{
				Role: "Editor",
			}),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveOrgUser(t *testing.T) {
	server, client := gapiTestTools(t, 200, removeOrgUserJSON)
	defer server.Close()

	_, err := client.Orgs.RemoveOrgUser(
		orgs.NewRemoveOrgUserParams().
			WithOrgID(1).
			WithUserID(1),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
