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

// ListUserRolesReader is a Reader for the ListUserRoles structure.
type ListUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUserRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListUserRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListUserRolesOK creates a ListUserRolesOK with default headers values
func NewListUserRolesOK() *ListUserRolesOK {
	return &ListUserRolesOK{}
}

/*
ListUserRolesOK describes a response with status code 200, with default header values.

(empty)
*/
type ListUserRolesOK struct {
	Payload []*models.RoleDTO
}

func (o *ListUserRolesOK) Error() string {
	return fmt.Sprintf("[GET /access-control/users/{userId}/roles][%d] listUserRolesOK  %+v", 200, o.Payload)
}
func (o *ListUserRolesOK) GetPayload() []*models.RoleDTO {
	return o.Payload
}

func (o *ListUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserRolesBadRequest creates a ListUserRolesBadRequest with default headers values
func NewListUserRolesBadRequest() *ListUserRolesBadRequest {
	return &ListUserRolesBadRequest{}
}

/*
ListUserRolesBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type ListUserRolesBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *ListUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /access-control/users/{userId}/roles][%d] listUserRolesBadRequest  %+v", 400, o.Payload)
}
func (o *ListUserRolesBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserRolesForbidden creates a ListUserRolesForbidden with default headers values
func NewListUserRolesForbidden() *ListUserRolesForbidden {
	return &ListUserRolesForbidden{}
}

/*
ListUserRolesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ListUserRolesForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *ListUserRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /access-control/users/{userId}/roles][%d] listUserRolesForbidden  %+v", 403, o.Payload)
}
func (o *ListUserRolesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListUserRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserRolesInternalServerError creates a ListUserRolesInternalServerError with default headers values
func NewListUserRolesInternalServerError() *ListUserRolesInternalServerError {
	return &ListUserRolesInternalServerError{}
}

/*
ListUserRolesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ListUserRolesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *ListUserRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access-control/users/{userId}/roles][%d] listUserRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *ListUserRolesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListUserRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
