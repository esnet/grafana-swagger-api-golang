// Code generated by go-swagger; DO NOT EDIT.

package sync_team_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// AddTeamGroupAPIReader is a Reader for the AddTeamGroupAPI structure.
type AddTeamGroupAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTeamGroupAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTeamGroupAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddTeamGroupAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddTeamGroupAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddTeamGroupAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddTeamGroupAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddTeamGroupAPIInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddTeamGroupAPIOK creates a AddTeamGroupAPIOK with default headers values
func NewAddTeamGroupAPIOK() *AddTeamGroupAPIOK {
	return &AddTeamGroupAPIOK{}
}

/*
AddTeamGroupAPIOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AddTeamGroupAPIOK struct {
	Payload *models.SuccessResponseBody
}

func (o *AddTeamGroupAPIOK) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiOK  %+v", 200, o.Payload)
}
func (o *AddTeamGroupAPIOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPIBadRequest creates a AddTeamGroupAPIBadRequest with default headers values
func NewAddTeamGroupAPIBadRequest() *AddTeamGroupAPIBadRequest {
	return &AddTeamGroupAPIBadRequest{}
}

/*
AddTeamGroupAPIBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AddTeamGroupAPIBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamGroupAPIBadRequest) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiBadRequest  %+v", 400, o.Payload)
}
func (o *AddTeamGroupAPIBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPIUnauthorized creates a AddTeamGroupAPIUnauthorized with default headers values
func NewAddTeamGroupAPIUnauthorized() *AddTeamGroupAPIUnauthorized {
	return &AddTeamGroupAPIUnauthorized{}
}

/*
AddTeamGroupAPIUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AddTeamGroupAPIUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamGroupAPIUnauthorized) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiUnauthorized  %+v", 401, o.Payload)
}
func (o *AddTeamGroupAPIUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPIForbidden creates a AddTeamGroupAPIForbidden with default headers values
func NewAddTeamGroupAPIForbidden() *AddTeamGroupAPIForbidden {
	return &AddTeamGroupAPIForbidden{}
}

/*
AddTeamGroupAPIForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AddTeamGroupAPIForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamGroupAPIForbidden) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiForbidden  %+v", 403, o.Payload)
}
func (o *AddTeamGroupAPIForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPINotFound creates a AddTeamGroupAPINotFound with default headers values
func NewAddTeamGroupAPINotFound() *AddTeamGroupAPINotFound {
	return &AddTeamGroupAPINotFound{}
}

/*
AddTeamGroupAPINotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type AddTeamGroupAPINotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamGroupAPINotFound) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiNotFound  %+v", 404, o.Payload)
}
func (o *AddTeamGroupAPINotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTeamGroupAPIInternalServerError creates a AddTeamGroupAPIInternalServerError with default headers values
func NewAddTeamGroupAPIInternalServerError() *AddTeamGroupAPIInternalServerError {
	return &AddTeamGroupAPIInternalServerError{}
}

/*
AddTeamGroupAPIInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AddTeamGroupAPIInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AddTeamGroupAPIInternalServerError) Error() string {
	return fmt.Sprintf("[POST /teams/{teamId}/groups][%d] addTeamGroupApiInternalServerError  %+v", 500, o.Payload)
}
func (o *AddTeamGroupAPIInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AddTeamGroupAPIInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
