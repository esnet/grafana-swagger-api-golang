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

// PostACSReader is a Reader for the PostACS structure.
type PostACSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostACSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewPostACSFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostACSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostACSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostACSFound creates a PostACSFound with default headers values
func NewPostACSFound() *PostACSFound {
	return &PostACSFound{}
}

/*
PostACSFound describes a response with status code 302, with default header values.

(empty)
*/
type PostACSFound struct {
}

func (o *PostACSFound) Error() string {
	return fmt.Sprintf("[POST /saml/acs][%d] postACSFound ", 302)
}

func (o *PostACSFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostACSForbidden creates a PostACSForbidden with default headers values
func NewPostACSForbidden() *PostACSForbidden {
	return &PostACSForbidden{}
}

/*
PostACSForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type PostACSForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *PostACSForbidden) Error() string {
	return fmt.Sprintf("[POST /saml/acs][%d] postACSForbidden  %+v", 403, o.Payload)
}
func (o *PostACSForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostACSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostACSInternalServerError creates a PostACSInternalServerError with default headers values
func NewPostACSInternalServerError() *PostACSInternalServerError {
	return &PostACSInternalServerError{}
}

/*
PostACSInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PostACSInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *PostACSInternalServerError) Error() string {
	return fmt.Sprintf("[POST /saml/acs][%d] postACSInternalServerError  %+v", 500, o.Payload)
}
func (o *PostACSInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostACSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}