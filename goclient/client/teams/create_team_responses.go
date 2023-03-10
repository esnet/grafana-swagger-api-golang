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

// CreateTeamReader is a Reader for the CreateTeam structure.
type CreateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateTeamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTeamConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTeamOK creates a CreateTeamOK with default headers values
func NewCreateTeamOK() *CreateTeamOK {
	return &CreateTeamOK{}
}

/*
CreateTeamOK describes a response with status code 200, with default header values.

(empty)
*/
type CreateTeamOK struct {
	Payload *models.CreateTeamOKBody
}

func (o *CreateTeamOK) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamOK  %+v", 200, o.Payload)
}
func (o *CreateTeamOK) GetPayload() *models.CreateTeamOKBody {
	return o.Payload
}

func (o *CreateTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateTeamOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamUnauthorized creates a CreateTeamUnauthorized with default headers values
func NewCreateTeamUnauthorized() *CreateTeamUnauthorized {
	return &CreateTeamUnauthorized{}
}

/*
CreateTeamUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreateTeamUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateTeamUnauthorized) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateTeamUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTeamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamForbidden creates a CreateTeamForbidden with default headers values
func NewCreateTeamForbidden() *CreateTeamForbidden {
	return &CreateTeamForbidden{}
}

/*
CreateTeamForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateTeamForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateTeamForbidden) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamForbidden  %+v", 403, o.Payload)
}
func (o *CreateTeamForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamConflict creates a CreateTeamConflict with default headers values
func NewCreateTeamConflict() *CreateTeamConflict {
	return &CreateTeamConflict{}
}

/*
CreateTeamConflict describes a response with status code 409, with default header values.

ConflictError
*/
type CreateTeamConflict struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateTeamConflict) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamConflict  %+v", 409, o.Payload)
}
func (o *CreateTeamConflict) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTeamConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTeamInternalServerError creates a CreateTeamInternalServerError with default headers values
func NewCreateTeamInternalServerError() *CreateTeamInternalServerError {
	return &CreateTeamInternalServerError{}
}

/*
CreateTeamInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateTeamInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateTeamInternalServerError) Error() string {
	return fmt.Sprintf("[POST /teams][%d] createTeamInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateTeamInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
