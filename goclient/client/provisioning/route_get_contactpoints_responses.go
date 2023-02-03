// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// RouteGetContactpointsReader is a Reader for the RouteGetContactpoints structure.
type RouteGetContactpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouteGetContactpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouteGetContactpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRouteGetContactpointsOK creates a RouteGetContactpointsOK with default headers values
func NewRouteGetContactpointsOK() *RouteGetContactpointsOK {
	return &RouteGetContactpointsOK{}
}

/*
RouteGetContactpointsOK describes a response with status code 200, with default header values.

ContactPoints
*/
type RouteGetContactpointsOK struct {
	Payload models.ContactPoints
}

func (o *RouteGetContactpointsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/contact-points][%d] routeGetContactpointsOK  %+v", 200, o.Payload)
}
func (o *RouteGetContactpointsOK) GetPayload() models.ContactPoints {
	return o.Payload
}

func (o *RouteGetContactpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}