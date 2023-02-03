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

// RoutePostContactpointsReader is a Reader for the RoutePostContactpoints structure.
type RoutePostContactpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoutePostContactpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRoutePostContactpointsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoutePostContactpointsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoutePostContactpointsAccepted creates a RoutePostContactpointsAccepted with default headers values
func NewRoutePostContactpointsAccepted() *RoutePostContactpointsAccepted {
	return &RoutePostContactpointsAccepted{}
}

/*
RoutePostContactpointsAccepted describes a response with status code 202, with default header values.

EmbeddedContactPoint
*/
type RoutePostContactpointsAccepted struct {
	Payload *models.EmbeddedContactPoint
}

func (o *RoutePostContactpointsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/provisioning/contact-points][%d] routePostContactpointsAccepted  %+v", 202, o.Payload)
}
func (o *RoutePostContactpointsAccepted) GetPayload() *models.EmbeddedContactPoint {
	return o.Payload
}

func (o *RoutePostContactpointsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmbeddedContactPoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoutePostContactpointsBadRequest creates a RoutePostContactpointsBadRequest with default headers values
func NewRoutePostContactpointsBadRequest() *RoutePostContactpointsBadRequest {
	return &RoutePostContactpointsBadRequest{}
}

/*
RoutePostContactpointsBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RoutePostContactpointsBadRequest struct {
	Payload *models.ValidationError
}

func (o *RoutePostContactpointsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/provisioning/contact-points][%d] routePostContactpointsBadRequest  %+v", 400, o.Payload)
}
func (o *RoutePostContactpointsBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RoutePostContactpointsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}