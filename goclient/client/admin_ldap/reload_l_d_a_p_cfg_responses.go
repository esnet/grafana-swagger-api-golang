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

// ReloadLDAPCfgReader is a Reader for the ReloadLDAPCfg structure.
type ReloadLDAPCfgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReloadLDAPCfgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReloadLDAPCfgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReloadLDAPCfgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReloadLDAPCfgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReloadLDAPCfgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReloadLDAPCfgOK creates a ReloadLDAPCfgOK with default headers values
func NewReloadLDAPCfgOK() *ReloadLDAPCfgOK {
	return &ReloadLDAPCfgOK{}
}

/*
ReloadLDAPCfgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type ReloadLDAPCfgOK struct {
	Payload *models.SuccessResponseBody
}

func (o *ReloadLDAPCfgOK) Error() string {
	return fmt.Sprintf("[POST /admin/ldap/reload][%d] reloadLDAPCfgOK  %+v", 200, o.Payload)
}
func (o *ReloadLDAPCfgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *ReloadLDAPCfgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReloadLDAPCfgUnauthorized creates a ReloadLDAPCfgUnauthorized with default headers values
func NewReloadLDAPCfgUnauthorized() *ReloadLDAPCfgUnauthorized {
	return &ReloadLDAPCfgUnauthorized{}
}

/*
ReloadLDAPCfgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type ReloadLDAPCfgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *ReloadLDAPCfgUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/ldap/reload][%d] reloadLDAPCfgUnauthorized  %+v", 401, o.Payload)
}
func (o *ReloadLDAPCfgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ReloadLDAPCfgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReloadLDAPCfgForbidden creates a ReloadLDAPCfgForbidden with default headers values
func NewReloadLDAPCfgForbidden() *ReloadLDAPCfgForbidden {
	return &ReloadLDAPCfgForbidden{}
}

/*
ReloadLDAPCfgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ReloadLDAPCfgForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *ReloadLDAPCfgForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/ldap/reload][%d] reloadLDAPCfgForbidden  %+v", 403, o.Payload)
}
func (o *ReloadLDAPCfgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ReloadLDAPCfgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReloadLDAPCfgInternalServerError creates a ReloadLDAPCfgInternalServerError with default headers values
func NewReloadLDAPCfgInternalServerError() *ReloadLDAPCfgInternalServerError {
	return &ReloadLDAPCfgInternalServerError{}
}

/*
ReloadLDAPCfgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ReloadLDAPCfgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *ReloadLDAPCfgInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/ldap/reload][%d] reloadLDAPCfgInternalServerError  %+v", 500, o.Payload)
}
func (o *ReloadLDAPCfgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ReloadLDAPCfgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}