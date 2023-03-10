// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// UpdateTeamMemberReader is a Reader for the UpdateTeamMember structure.
type UpdateTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateTeamMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTeamMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTeamMemberOK creates a UpdateTeamMemberOK with default headers values
func NewUpdateTeamMemberOK() *UpdateTeamMemberOK {
	return &UpdateTeamMemberOK{}
}

/*
UpdateTeamMemberOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateTeamMemberOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateTeamMemberOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/members/{user_id}][%d] updateTeamMemberOK  %+v", 200, o.Payload)
}
func (o *UpdateTeamMemberOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberUnauthorized creates a UpdateTeamMemberUnauthorized with default headers values
func NewUpdateTeamMemberUnauthorized() *UpdateTeamMemberUnauthorized {
	return &UpdateTeamMemberUnauthorized{}
}

/*
UpdateTeamMemberUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateTeamMemberUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateTeamMemberUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/members/{user_id}][%d] updateTeamMemberUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateTeamMemberUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberForbidden creates a UpdateTeamMemberForbidden with default headers values
func NewUpdateTeamMemberForbidden() *UpdateTeamMemberForbidden {
	return &UpdateTeamMemberForbidden{}
}

/*
UpdateTeamMemberForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateTeamMemberForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/members/{user_id}][%d] updateTeamMemberForbidden  %+v", 403, o.Payload)
}
func (o *UpdateTeamMemberForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberNotFound creates a UpdateTeamMemberNotFound with default headers values
func NewUpdateTeamMemberNotFound() *UpdateTeamMemberNotFound {
	return &UpdateTeamMemberNotFound{}
}

/*
UpdateTeamMemberNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateTeamMemberNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/members/{user_id}][%d] updateTeamMemberNotFound  %+v", 404, o.Payload)
}
func (o *UpdateTeamMemberNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberInternalServerError creates a UpdateTeamMemberInternalServerError with default headers values
func NewUpdateTeamMemberInternalServerError() *UpdateTeamMemberInternalServerError {
	return &UpdateTeamMemberInternalServerError{}
}

/*
UpdateTeamMemberInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateTeamMemberInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateTeamMemberInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/members/{user_id}][%d] updateTeamMemberInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateTeamMemberInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
