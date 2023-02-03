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

// RemoveTeamRoleReader is a Reader for the RemoveTeamRole structure.
type RemoveTeamRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveTeamRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveTeamRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveTeamRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveTeamRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveTeamRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveTeamRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveTeamRoleOK creates a RemoveTeamRoleOK with default headers values
func NewRemoveTeamRoleOK() *RemoveTeamRoleOK {
	return &RemoveTeamRoleOK{}
}

/*
RemoveTeamRoleOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type RemoveTeamRoleOK struct {
	Payload *models.SuccessResponseBody
}

func (o *RemoveTeamRoleOK) Error() string {
	return fmt.Sprintf("[DELETE /access-control/teams/{teamId}/roles/{roleUID}][%d] removeTeamRoleOK  %+v", 200, o.Payload)
}
func (o *RemoveTeamRoleOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *RemoveTeamRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamRoleBadRequest creates a RemoveTeamRoleBadRequest with default headers values
func NewRemoveTeamRoleBadRequest() *RemoveTeamRoleBadRequest {
	return &RemoveTeamRoleBadRequest{}
}

/*
RemoveTeamRoleBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RemoveTeamRoleBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveTeamRoleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /access-control/teams/{teamId}/roles/{roleUID}][%d] removeTeamRoleBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveTeamRoleBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamRoleForbidden creates a RemoveTeamRoleForbidden with default headers values
func NewRemoveTeamRoleForbidden() *RemoveTeamRoleForbidden {
	return &RemoveTeamRoleForbidden{}
}

/*
RemoveTeamRoleForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RemoveTeamRoleForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveTeamRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /access-control/teams/{teamId}/roles/{roleUID}][%d] removeTeamRoleForbidden  %+v", 403, o.Payload)
}
func (o *RemoveTeamRoleForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamRoleNotFound creates a RemoveTeamRoleNotFound with default headers values
func NewRemoveTeamRoleNotFound() *RemoveTeamRoleNotFound {
	return &RemoveTeamRoleNotFound{}
}

/*
RemoveTeamRoleNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RemoveTeamRoleNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveTeamRoleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /access-control/teams/{teamId}/roles/{roleUID}][%d] removeTeamRoleNotFound  %+v", 404, o.Payload)
}
func (o *RemoveTeamRoleNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveTeamRoleInternalServerError creates a RemoveTeamRoleInternalServerError with default headers values
func NewRemoveTeamRoleInternalServerError() *RemoveTeamRoleInternalServerError {
	return &RemoveTeamRoleInternalServerError{}
}

/*
RemoveTeamRoleInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RemoveTeamRoleInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveTeamRoleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /access-control/teams/{teamId}/roles/{roleUID}][%d] removeTeamRoleInternalServerError  %+v", 500, o.Payload)
}
func (o *RemoveTeamRoleInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveTeamRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
