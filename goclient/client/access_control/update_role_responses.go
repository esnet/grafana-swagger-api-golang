// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// UpdateRoleReader is a Reader for the UpdateRole structure.
type UpdateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRoleOK creates a UpdateRoleOK with default headers values
func NewUpdateRoleOK() *UpdateRoleOK {
	return &UpdateRoleOK{}
}

/*
UpdateRoleOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateRoleOK struct {
	Payload *models.RoleDTO
}

func (o *UpdateRoleOK) Error() string {
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}][%d] updateRoleOK  %+v", 200, o.Payload)
}
func (o *UpdateRoleOK) GetPayload() *models.RoleDTO {
	return o.Payload
}

func (o *UpdateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleBadRequest creates a UpdateRoleBadRequest with default headers values
func NewUpdateRoleBadRequest() *UpdateRoleBadRequest {
	return &UpdateRoleBadRequest{}
}

/*
UpdateRoleBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateRoleBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateRoleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}][%d] updateRoleBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateRoleBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleForbidden creates a UpdateRoleForbidden with default headers values
func NewUpdateRoleForbidden() *UpdateRoleForbidden {
	return &UpdateRoleForbidden{}
}

/*
UpdateRoleForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateRoleForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateRoleForbidden) Error() string {
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}][%d] updateRoleForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRoleForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleNotFound creates a UpdateRoleNotFound with default headers values
func NewUpdateRoleNotFound() *UpdateRoleNotFound {
	return &UpdateRoleNotFound{}
}

/*
UpdateRoleNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateRoleNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateRoleNotFound) Error() string {
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}][%d] updateRoleNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRoleNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleInternalServerError creates a UpdateRoleInternalServerError with default headers values
func NewUpdateRoleInternalServerError() *UpdateRoleInternalServerError {
	return &UpdateRoleInternalServerError{}
}

/*
UpdateRoleInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateRoleInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateRoleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}][%d] updateRoleInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateRoleInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}