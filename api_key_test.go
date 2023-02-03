package gapi

import (
	"context"
	"testing"
	"time"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/api_keys"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
	"github.com/gobs/pretty"
)

const (
	createAPIKeyJSON = `{"name":"key-name", "key":"mock-api-key"}` //#nosec
	deleteAPIKeyJSON = `{"message":"API key deleted"}`             //#nosec

	getAPIKeysJSON = `[
		{
			"id": 1,
			"name": "key-name-2",
			"role": "Viewer"
		},
		{
			"id": 2,
			"name": "key-name-2",
			"role": "Admin",
			"expiration": "2021-10-30T10:52:03+03:00"
		}
	]`  //#nosec
)

func TestCreateAPIKey(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createAPIKeyJSON)
	defer mocksrv.Close()

	params := api_keys.AddAPIkeyParams{
		Body: &models.AddCommand{
			Name:          "key-name",
			Role:          "Viewer",
			SecondsToLive: int64(time.Hour),
		},
		Context: context.Background(),
	}

	resp, err := client.APIKeys.AddAPIkey(&params, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(pretty.PrettyFormat(resp))
}

func TestDeleteAPIKey(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, deleteAPIKeyJSON)
	defer mocksrv.Close()

	res, err := client.APIKeys.DeleteAPIkey(api_keys.NewDeleteAPIkeyParams().WithID(1), nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetAPIKeys(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getAPIKeysJSON)
	defer mocksrv.Close()

	res, err := client.APIKeys.GetAPIkeys(api_keys.NewGetAPIkeysParams(), nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}
