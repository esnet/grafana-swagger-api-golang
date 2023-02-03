package gapi

import (
	"context"
	"encoding/base64"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/datasources"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
	"strings"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	getDataSource = `
{
	        "id": 10,
	        "orgId": 1,
	        "uid": "rg9qPqP7z",
	        "name": "netsage",
	        "type": "elasticsearch",
	        "typeLogoUrl": "public/app/plugins/datasource/elasticsearch/img/elasticsearch.svg",
	        "access": "proxy",
	        "url": "https://netsage-elk1.grnoc.iu.edu/esproxy2/",
	        "password": "",
	        "user": "",
	        "database": "om-netsage-irnc-*",
	        "basicAuth": true,
	        "readOnly": false,
	        "isDefault": true,
	        "jsonData": {
	        },
	        "secureJsonData": null
	}
`
	getDataSources = `[ ` + getDataSource + `] `
)

func TestNew_basicAuth(t *testing.T) {
	a := BasicAuthenticator{
		Username: "user",
		Password: "pass",
	}

	tr := new(runtime.TestClientRequest)
	err := a.AuthenticateRequest(tr, nil)
	require.NoError(t, err)

	auth, ok := tr.GetHeaderParams()["Authorization"]
	require.True(t, ok)
	require.Len(t, auth, 1)

	assert.True(t, strings.Contains(auth[0], "Basic"))
	blob, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(auth[0], "Basic ", ""))
	require.NoError(t, err)
	assert.Equal(t, blob, []byte("user:pass"))
}

func TestNew_tokenAuth(t *testing.T) {
	a := APIKeyAuthenticator{
		APIKey: "123",
	}

	tr := new(runtime.TestClientRequest)
	err := a.AuthenticateRequest(tr, nil)
	require.NoError(t, err)

	auth, ok := tr.GetHeaderParams()["Authorization"]
	require.True(t, ok)
	require.Len(t, auth, 1)

	assert.Equal(t, auth[0], "Bearer 123")
	//blob, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(auth[0], "Bearer ", ""))
	require.NoError(t, err)
	//assert.Equal(t, blob, []byte("123"))
}

func TestNew_orgID(t *testing.T) {
	const orgID = 456
	c, err := New("http://my-grafana.com", Config{OrgID: orgID})
	if err != nil {
		t.Fatalf("expected error to be nil; got: %s", err.Error())
	}

	expected := "http://my-grafana.com"
	if c.baseURL.String() != expected {
		t.Errorf("expected error: %s; got: %s", expected, c.baseURL.String())
	}

	if c.config.OrgID != orgID {
		t.Errorf("expected error: %d; got: %d", orgID, c.config.OrgID)
	}
}

func TestNew_HTTPHeaders(t *testing.T) {
	const key = "foo"
	headers := map[string]string{key: "bar"}
	c, err := New("http://my-grafana.com", Config{HTTPHeaders: headers})
	if err != nil {
		t.Fatalf("expected error to be nil; got: %s", err.Error())
	}

	value, ok := c.config.HTTPHeaders[key]
	if !ok {
		t.Errorf("expected error: %v; got: %v", headers, c.config.HTTPHeaders)
	}
	if value != headers[key] {
		t.Errorf("expected error: %s; got: %s", headers[key], value)
	}
}

func TestNew_invalidURL(t *testing.T) {
	_, err := New("://my-grafana.com", Config{APIKey: "123"})

	expected := "parse \"://my-grafana.com\": missing protocol scheme"
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err)
	}
}

func TestRequest_200(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDataSources)
	defer server.Close()

	_, err := client.Datasources.GetDataSources(nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestRequest_201(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDataSource)
	defer server.Close()
	v := datasources.AddDataSourceParams{
		Context: context.Background(),
		Body: &models.AddDataSourceCommand{

			Type:      "elasticsearch",
			Name:      "netsage",
			UID:       "rg9qPqP7z",
			URL:       "https://netsage-elk1.grnoc.iu.edu/esproxy2/",
			Access:    "proxy",
			BasicAuth: true,
			IsDefault: true,
			Database:  "om-netsage-irnc-*",
		},
	}

	_, err := client.Datasources.AddDataSource(&v, nil, func(*runtime.ClientOperation) {})
	if err != nil {
		t.Error(err)
	}
}

func TestRequest_500(t *testing.T) {
	server, client := gapiTestTools(t, 500, `{"foo":"bar"}`)
	defer server.Close()

	expected := `[GET /datasources][500] getDataSourcesInternalServerError  &{Error: Message:<nil> Status:}`
	_, err := client.Datasources.GetDataSources(nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err)
	}
}

func TestRequest_badURL(t *testing.T) {
	client, err := GetClient("bad-url")
	require.NoError(t, err)

	expected := `Get "http:///api/datasources": http: no Host in request URL`
	_, err = client.Datasources.GetDataSources(nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err)
	}
}
