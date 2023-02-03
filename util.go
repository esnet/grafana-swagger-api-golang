package gapi

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type BasicAuthenticator struct {
	Username string
	Password string
}

func (a BasicAuthenticator) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	creds := fmt.Sprintf("%s:%s", a.Username, a.Password)
	return req.SetHeaderParam("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(creds))))
}

type APIKeyAuthenticator struct {
	APIKey string
}

func (a APIKeyAuthenticator) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	return req.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", a.APIKey))
}

func GetClient(serverURL string) (*client.GrafanaHTTPAPI, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := client.NewHTTPClientWithConfig(
		nil,
		client.DefaultTransportConfig().
			WithHost(u.Host).
			WithSchemes([]string{"http"}),
	)
	return c, nil
}
