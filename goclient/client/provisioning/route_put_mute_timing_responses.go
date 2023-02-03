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

// RoutePutMuteTimingReader is a Reader for the RoutePutMuteTiming structure.
type RoutePutMuteTimingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoutePutMuteTimingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoutePutMuteTimingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoutePutMuteTimingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoutePutMuteTimingOK creates a RoutePutMuteTimingOK with default headers values
func NewRoutePutMuteTimingOK() *RoutePutMuteTimingOK {
	return &RoutePutMuteTimingOK{}
}

/*
RoutePutMuteTimingOK describes a response with status code 200, with default header values.

MuteTimeInterval
*/
type RoutePutMuteTimingOK struct {
	Payload *models.MuteTimeInterval
}

func (o *RoutePutMuteTimingOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/mute-timings/{name}][%d] routePutMuteTimingOK  %+v", 200, o.Payload)
}
func (o *RoutePutMuteTimingOK) GetPayload() *models.MuteTimeInterval {
	return o.Payload
}

func (o *RoutePutMuteTimingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MuteTimeInterval)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoutePutMuteTimingBadRequest creates a RoutePutMuteTimingBadRequest with default headers values
func NewRoutePutMuteTimingBadRequest() *RoutePutMuteTimingBadRequest {
	return &RoutePutMuteTimingBadRequest{}
}

/*
RoutePutMuteTimingBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RoutePutMuteTimingBadRequest struct {
	Payload *models.ValidationError
}

func (o *RoutePutMuteTimingBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/mute-timings/{name}][%d] routePutMuteTimingBadRequest  %+v", 400, o.Payload)
}
func (o *RoutePutMuteTimingBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RoutePutMuteTimingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}