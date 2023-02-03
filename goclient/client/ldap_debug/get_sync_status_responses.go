// Code generated by go-swagger; DO NOT EDIT.

package ldap_debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetSyncStatusReader is a Reader for the GetSyncStatus structure.
type GetSyncStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyncStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSyncStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSyncStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSyncStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSyncStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSyncStatusOK creates a GetSyncStatusOK with default headers values
func NewGetSyncStatusOK() *GetSyncStatusOK {
	return &GetSyncStatusOK{}
}

/*
GetSyncStatusOK describes a response with status code 200, with default header values.

(empty)
*/
type GetSyncStatusOK struct {
	Payload *models.ActiveSyncStatusDTO
}

func (o *GetSyncStatusOK) Error() string {
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusOK  %+v", 200, o.Payload)
}
func (o *GetSyncStatusOK) GetPayload() *models.ActiveSyncStatusDTO {
	return o.Payload
}

func (o *GetSyncStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveSyncStatusDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyncStatusUnauthorized creates a GetSyncStatusUnauthorized with default headers values
func NewGetSyncStatusUnauthorized() *GetSyncStatusUnauthorized {
	return &GetSyncStatusUnauthorized{}
}

/*
GetSyncStatusUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetSyncStatusUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSyncStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSyncStatusUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSyncStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyncStatusForbidden creates a GetSyncStatusForbidden with default headers values
func NewGetSyncStatusForbidden() *GetSyncStatusForbidden {
	return &GetSyncStatusForbidden{}
}

/*
GetSyncStatusForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetSyncStatusForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSyncStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusForbidden  %+v", 403, o.Payload)
}
func (o *GetSyncStatusForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSyncStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyncStatusInternalServerError creates a GetSyncStatusInternalServerError with default headers values
func NewGetSyncStatusInternalServerError() *GetSyncStatusInternalServerError {
	return &GetSyncStatusInternalServerError{}
}

/*
GetSyncStatusInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetSyncStatusInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSyncStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /admin/ldap-sync-status][%d] getSyncStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSyncStatusInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSyncStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
