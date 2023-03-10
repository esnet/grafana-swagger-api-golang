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

// RoutePutPolicyTreeReader is a Reader for the RoutePutPolicyTree structure.
type RoutePutPolicyTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoutePutPolicyTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRoutePutPolicyTreeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoutePutPolicyTreeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoutePutPolicyTreeAccepted creates a RoutePutPolicyTreeAccepted with default headers values
func NewRoutePutPolicyTreeAccepted() *RoutePutPolicyTreeAccepted {
	return &RoutePutPolicyTreeAccepted{}
}

/*
RoutePutPolicyTreeAccepted describes a response with status code 202, with default header values.

Ack
*/
type RoutePutPolicyTreeAccepted struct {
	Payload models.Ack
}

func (o *RoutePutPolicyTreeAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/policies][%d] routePutPolicyTreeAccepted  %+v", 202, o.Payload)
}
func (o *RoutePutPolicyTreeAccepted) GetPayload() models.Ack {
	return o.Payload
}

func (o *RoutePutPolicyTreeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoutePutPolicyTreeBadRequest creates a RoutePutPolicyTreeBadRequest with default headers values
func NewRoutePutPolicyTreeBadRequest() *RoutePutPolicyTreeBadRequest {
	return &RoutePutPolicyTreeBadRequest{}
}

/*
RoutePutPolicyTreeBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RoutePutPolicyTreeBadRequest struct {
	Payload *models.ValidationError
}

func (o *RoutePutPolicyTreeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/policies][%d] routePutPolicyTreeBadRequest  %+v", 400, o.Payload)
}
func (o *RoutePutPolicyTreeBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RoutePutPolicyTreeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
