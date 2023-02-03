package gapi

import (
	"testing"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/playlists"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

const (
	createAndUpdatePlaylistResponse = `  {
		"id": 1,
		"name": "my playlist",
		"interval": "5m"
	}`

	getPlaylistResponse = `{
		"id" : 2,
		"uid": "uid",
		"name": "my playlist",
		"interval": "5m",
		"orgId": "my org",
		"items": [
			{
				"id": 1,
				"playlistId": 1,
				"type": "dashboard_by_id",
				"value": "3",
				"order": 1,
				"title":"my dasboard"
			},
			{
				"id": 1,
				"playlistId": 1,
				"type": "dashboard_by_id",
				"value": "3",
				"order": 1,
				"title":"my dasboard"
			}
		]
	}`
)

func TestPlaylistCreateAndUpdate(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, createAndUpdatePlaylistResponse)
	defer mocksrv.Close()

	playlist := models.CreatePlaylistCommand{
		Name:     "my playlist",
		Interval: "5m",
		Items: []*models.PlaylistItem{
			{},
		},
	}

	// create
	resp, err := client.Playlists.CreatePlaylist(
		playlists.NewCreatePlaylistParams().
			WithBody(&playlist),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Payload.ID != 1 {
		t.Errorf("Invalid id - %d, Expected %d", resp.Payload.ID, 1)
	}

	// update
	playlist.Items = append(playlist.Items, &models.PlaylistItem{
		Type:  "dashboard_by_id",
		Value: "1",
		Title: "my dashboard",
	})

	_, err = client.Playlists.UpdatePlaylist(
		playlists.NewUpdatePlaylistParams().
			WithUID(resp.Payload.UID).
			WithBody(&models.UpdatePlaylistCommand{
				Items: []*models.PlaylistItem{
					{
						Type:  "dashboard_by_id",
						Value: "1",
						Title: "my dashboard",
					},
				},
			}),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPlaylist(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, getPlaylistResponse)
	defer mocksrv.Close()

	resp, err := client.Playlists.GetPlaylist(
		playlists.NewGetPlaylistParams().WithUID("uid"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Payload.Items) != 2 {
		t.Errorf("Invalid len(items) - %d, Expected %d", len(resp.Payload.Items), 2)
	}
}

func TestDeletePlaylist(t *testing.T) {
	mocksrv, client := gapiTestTools(t, 200, "")
	defer mocksrv.Close()

	_, err := client.Playlists.DeletePlaylist(
		playlists.NewDeletePlaylistParams().WithUID("uid"),
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
}
