// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetAPIkeysReader is a Reader for the GetAPIkeys structure.
type GetAPIkeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIkeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIkeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAPIkeysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIkeysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIkeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIkeysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIkeysOK creates a GetAPIkeysOK with default headers values
func NewGetAPIkeysOK() *GetAPIkeysOK {
	return &GetAPIkeysOK{}
}

/*
GetAPIkeysOK describes a response with status code 200, with default header values.

(empty)
*/
type GetAPIkeysOK struct {
	Payload []*models.APIKeyDTO
}

func (o *GetAPIkeysOK) Error() string {
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysOK  %+v", 200, o.Payload)
}
func (o *GetAPIkeysOK) GetPayload() []*models.APIKeyDTO {
	return o.Payload
}

func (o *GetAPIkeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIkeysUnauthorized creates a GetAPIkeysUnauthorized with default headers values
func NewGetAPIkeysUnauthorized() *GetAPIkeysUnauthorized {
	return &GetAPIkeysUnauthorized{}
}

/*
GetAPIkeysUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetAPIkeysUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetAPIkeysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAPIkeysUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAPIkeysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIkeysForbidden creates a GetAPIkeysForbidden with default headers values
func NewGetAPIkeysForbidden() *GetAPIkeysForbidden {
	return &GetAPIkeysForbidden{}
}

/*
GetAPIkeysForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetAPIkeysForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetAPIkeysForbidden) Error() string {
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysForbidden  %+v", 403, o.Payload)
}
func (o *GetAPIkeysForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAPIkeysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIkeysNotFound creates a GetAPIkeysNotFound with default headers values
func NewGetAPIkeysNotFound() *GetAPIkeysNotFound {
	return &GetAPIkeysNotFound{}
}

/*
GetAPIkeysNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetAPIkeysNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetAPIkeysNotFound) Error() string {
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysNotFound  %+v", 404, o.Payload)
}
func (o *GetAPIkeysNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAPIkeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIkeysInternalServerError creates a GetAPIkeysInternalServerError with default headers values
func NewGetAPIkeysInternalServerError() *GetAPIkeysInternalServerError {
	return &GetAPIkeysInternalServerError{}
}

/*
GetAPIkeysInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetAPIkeysInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetAPIkeysInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAPIkeysInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAPIkeysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}