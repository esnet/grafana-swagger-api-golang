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

// ListTeamRolesReader is a Reader for the ListTeamRoles structure.
type ListTeamRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTeamRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTeamRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListTeamRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTeamRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTeamRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTeamRolesOK creates a ListTeamRolesOK with default headers values
func NewListTeamRolesOK() *ListTeamRolesOK {
	return &ListTeamRolesOK{}
}

/*
ListTeamRolesOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type ListTeamRolesOK struct {
	Payload *models.SuccessResponseBody
}

func (o *ListTeamRolesOK) Error() string {
	return fmt.Sprintf("[GET /access-control/teams/{teamId}/roles][%d] listTeamRolesOK  %+v", 200, o.Payload)
}
func (o *ListTeamRolesOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *ListTeamRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamRolesBadRequest creates a ListTeamRolesBadRequest with default headers values
func NewListTeamRolesBadRequest() *ListTeamRolesBadRequest {
	return &ListTeamRolesBadRequest{}
}

/*
ListTeamRolesBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type ListTeamRolesBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *ListTeamRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /access-control/teams/{teamId}/roles][%d] listTeamRolesBadRequest  %+v", 400, o.Payload)
}
func (o *ListTeamRolesBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListTeamRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamRolesForbidden creates a ListTeamRolesForbidden with default headers values
func NewListTeamRolesForbidden() *ListTeamRolesForbidden {
	return &ListTeamRolesForbidden{}
}

/*
ListTeamRolesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ListTeamRolesForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *ListTeamRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /access-control/teams/{teamId}/roles][%d] listTeamRolesForbidden  %+v", 403, o.Payload)
}
func (o *ListTeamRolesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListTeamRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTeamRolesInternalServerError creates a ListTeamRolesInternalServerError with default headers values
func NewListTeamRolesInternalServerError() *ListTeamRolesInternalServerError {
	return &ListTeamRolesInternalServerError{}
}

/*
ListTeamRolesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ListTeamRolesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *ListTeamRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access-control/teams/{teamId}/roles][%d] listTeamRolesInternalServerError  %+v", 500, o.Payload)
}
func (o *ListTeamRolesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListTeamRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
