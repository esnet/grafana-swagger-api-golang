package gapi

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/users"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
	"github.com/gobs/pretty"
)

const (
	getUsersJSON       = `{ "users": [ {"id":1,"email":"users@localhost","isAdmin":true} ] } `
	getUserJSON        = `{"id":2,"email":"user@localhost","isGrafanaAdmin":false}`
	getUserByEmailJSON = `{"id":3,"email":"userByEmail@localhost","isGrafanaAdmin":true}`
	getUserUpdateJSON  = `{"id":4,"email":"userUpdate@localhost","isGrafanaAdmin":false}`
)

func TestSamir(t *testing.T) {
	v := models.SearchUserQueryResult{}
	err := json.Unmarshal([]byte(getUsersJSON), &v)
	assert.Nil(t, err)
}

func TestUsers(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getUsersJSON)
	defer mocksrv.Close()

	resp, err := client.Users.SearchUsers(
		users.NewSearchUsersParams(),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if len(resp.Payload.Users) != 1 {
		t.Fatal("No users were returned.")
	}

	user := resp.Payload.Users[0]

	if user.Email != "users@localhost" ||
		user.ID != 1 ||
		!user.IsAdmin {
		t.Error("Not correctly parsing returned users.")
	}
}

func TestUser(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getUserJSON)
	defer mocksrv.Close()

	user, err := client.Users.GetUserByID(
		users.NewGetUserByIDParams().
			WithUserID(1),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(user))

	if user.Payload.Email != "user@localhost" ||
		user.Payload.ID != 2 ||
		user.Payload.IsGrafanaAdmin != false {
		t.Error("Not correctly parsing returned user.")
	}
}

func TestUserByEmail(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getUserByEmailJSON)
	defer mocksrv.Close()

	user, err := client.Users.GetUserByLoginOrEmail(
		users.NewGetUserByLoginOrEmailParams().
			WithLoginOrEmail("admin@localhost"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(user))

	if user.Payload.Email != "userByEmail@localhost" ||
		user.Payload.ID != 3 ||
		user.Payload.IsGrafanaAdmin != true {
		t.Error("Not correctly parsing returned user.")
	}
}

func TestUserUpdate(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getUserUpdateJSON)
	defer mocksrv.Close()

	user, err := client.Users.GetUserByID(
		users.NewGetUserByIDParams().
			WithUserID(4),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Users.UpdateUser(
		users.NewUpdateUserParams().
			WithUserID(user.Payload.ID).
			WithBody(&models.UpdateUserCommand{
				Email: user.Payload.Email,
				Name:  user.Payload.Name,
				Login: user.Payload.Login,
				Theme: "dark",
			}),
		nil,
	)
	if err != nil {
		t.Error(err)
	}
}
