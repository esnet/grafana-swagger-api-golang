// Code generated by go-swagger; DO NOT EDIT.

package saml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetSLOReader is a Reader for the GetSLO structure.
type GetSLOReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSLOReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewGetSLOFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetSLOBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSLOForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSLOInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSLOFound creates a GetSLOFound with default headers values
func NewGetSLOFound() *GetSLOFound {
	return &GetSLOFound{}
}

/*
GetSLOFound describes a response with status code 302, with default header values.

(empty)
*/
type GetSLOFound struct {
}

func (o *GetSLOFound) Error() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOFound ", 302)
}

func (o *GetSLOFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSLOBadRequest creates a GetSLOBadRequest with default headers values
func NewGetSLOBadRequest() *GetSLOBadRequest {
	return &GetSLOBadRequest{}
}

/*
GetSLOBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type GetSLOBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSLOBadRequest) Error() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOBadRequest  %+v", 400, o.Payload)
}
func (o *GetSLOBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSLOBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSLOForbidden creates a GetSLOForbidden with default headers values
func NewGetSLOForbidden() *GetSLOForbidden {
	return &GetSLOForbidden{}
}

/*
GetSLOForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetSLOForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSLOForbidden) Error() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOForbidden  %+v", 403, o.Payload)
}
func (o *GetSLOForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSLOForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSLOInternalServerError creates a GetSLOInternalServerError with default headers values
func NewGetSLOInternalServerError() *GetSLOInternalServerError {
	return &GetSLOInternalServerError{}
}

/*
GetSLOInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetSLOInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSLOInternalServerError) Error() string {
	return fmt.Sprintf("[GET /saml/slo][%d] getSLOInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSLOInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSLOInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}