// Code generated by go-swagger; DO NOT EDIT.

package admin_ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetLDAPStatusReader is a Reader for the GetLDAPStatus structure.
type GetLDAPStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLDAPStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLDAPStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLDAPStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLDAPStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLDAPStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLDAPStatusOK creates a GetLDAPStatusOK with default headers values
func NewGetLDAPStatusOK() *GetLDAPStatusOK {
	return &GetLDAPStatusOK{}
}

/*
GetLDAPStatusOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type GetLDAPStatusOK struct {
	Payload *models.SuccessResponseBody
}

func (o *GetLDAPStatusOK) Error() string {
	return fmt.Sprintf("[GET /admin/ldap/status][%d] getLDAPStatusOK  %+v", 200, o.Payload)
}
func (o *GetLDAPStatusOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *GetLDAPStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLDAPStatusUnauthorized creates a GetLDAPStatusUnauthorized with default headers values
func NewGetLDAPStatusUnauthorized() *GetLDAPStatusUnauthorized {
	return &GetLDAPStatusUnauthorized{}
}

/*
GetLDAPStatusUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetLDAPStatusUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetLDAPStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /admin/ldap/status][%d] getLDAPStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *GetLDAPStatusUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLDAPStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLDAPStatusForbidden creates a GetLDAPStatusForbidden with default headers values
func NewGetLDAPStatusForbidden() *GetLDAPStatusForbidden {
	return &GetLDAPStatusForbidden{}
}

/*
GetLDAPStatusForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetLDAPStatusForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetLDAPStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/ldap/status][%d] getLDAPStatusForbidden  %+v", 403, o.Payload)
}
func (o *GetLDAPStatusForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLDAPStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLDAPStatusInternalServerError creates a GetLDAPStatusInternalServerError with default headers values
func NewGetLDAPStatusInternalServerError() *GetLDAPStatusInternalServerError {
	return &GetLDAPStatusInternalServerError{}
}

/*
GetLDAPStatusInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetLDAPStatusInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetLDAPStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/ldap/status][%d] getLDAPStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLDAPStatusInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLDAPStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
