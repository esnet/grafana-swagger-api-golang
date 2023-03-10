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

// PostSLOReader is a Reader for the PostSLO structure.
type PostSLOReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSLOReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewPostSLOFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewPostSLOBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSLOForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSLOInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSLOFound creates a PostSLOFound with default headers values
func NewPostSLOFound() *PostSLOFound {
	return &PostSLOFound{}
}

/*
PostSLOFound describes a response with status code 302, with default header values.

(empty)
*/
type PostSLOFound struct {
}

func (o *PostSLOFound) Error() string {
	return fmt.Sprintf("[POST /saml/slo][%d] postSLOFound ", 302)
}

func (o *PostSLOFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSLOBadRequest creates a PostSLOBadRequest with default headers values
func NewPostSLOBadRequest() *PostSLOBadRequest {
	return &PostSLOBadRequest{}
}

/*
PostSLOBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type PostSLOBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *PostSLOBadRequest) Error() string {
	return fmt.Sprintf("[POST /saml/slo][%d] postSLOBadRequest  %+v", 400, o.Payload)
}
func (o *PostSLOBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostSLOBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSLOForbidden creates a PostSLOForbidden with default headers values
func NewPostSLOForbidden() *PostSLOForbidden {
	return &PostSLOForbidden{}
}

/*
PostSLOForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type PostSLOForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *PostSLOForbidden) Error() string {
	return fmt.Sprintf("[POST /saml/slo][%d] postSLOForbidden  %+v", 403, o.Payload)
}
func (o *PostSLOForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostSLOForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSLOInternalServerError creates a PostSLOInternalServerError with default headers values
func NewPostSLOInternalServerError() *PostSLOInternalServerError {
	return &PostSLOInternalServerError{}
}

/*
PostSLOInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PostSLOInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *PostSLOInternalServerError) Error() string {
	return fmt.Sprintf("[POST /saml/slo][%d] postSLOInternalServerError  %+v", 500, o.Payload)
}
func (o *PostSLOInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostSLOInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
