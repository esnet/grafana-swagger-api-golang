package main

import (
	"log"

	gapi "github.com/esnet/grafana-swagger-api-golang"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/datasources"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

func main() {
	c, err := gapi.GetClient("http://localhost:3000")
	if err != nil {
		log.Fatalf("failed to get the client: %v", err)
	}

	params := datasources.NewAddDataSourceParams().WithBody(
		&models.AddDataSourceCommand{
			Name:      "foo",
			Type:      "cloudwatch",
			URL:       "http://some-url.com",
			Access:    "access",
			IsDefault: true,
			JSONData: map[string]string{
				"assumeRoleArn":           "arn:aws:iam::123:role/some-role",
				"authType":                "keys",
				"customMetricsNamespaces": "SomeNamespace",
				"defaultRegion":           "us-east-1",
				"tlsSkipVerify":           "true",
			},
			SecureJSONData: map[string]string{
				"accessKey": "123",
				"secretKey": "456",
			},
		},
	)

	res, err := c.Datasources.AddDataSource(params, gapi.BasicAuthenticator{
		Username: "admin",
		Password: "admin",
	})
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	log.Println("response", res.Payload)
}
