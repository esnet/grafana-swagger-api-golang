// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetUserByLoginOrEmailReader is a Reader for the GetUserByLoginOrEmail structure.
type GetUserByLoginOrEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserByLoginOrEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserByLoginOrEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserByLoginOrEmailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserByLoginOrEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserByLoginOrEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserByLoginOrEmailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserByLoginOrEmailOK creates a GetUserByLoginOrEmailOK with default headers values
func NewGetUserByLoginOrEmailOK() *GetUserByLoginOrEmailOK {
	return &GetUserByLoginOrEmailOK{}
}

/*
GetUserByLoginOrEmailOK describes a response with status code 200, with default header values.

(empty)
*/
type GetUserByLoginOrEmailOK struct {
	Payload *models.UserProfileDTO
}

func (o *GetUserByLoginOrEmailOK) Error() string {
	return fmt.Sprintf("[GET /users/lookup][%d] getUserByLoginOrEmailOK  %+v", 200, o.Payload)
}
func (o *GetUserByLoginOrEmailOK) GetPayload() *models.UserProfileDTO {
	return o.Payload
}

func (o *GetUserByLoginOrEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserProfileDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByLoginOrEmailUnauthorized creates a GetUserByLoginOrEmailUnauthorized with default headers values
func NewGetUserByLoginOrEmailUnauthorized() *GetUserByLoginOrEmailUnauthorized {
	return &GetUserByLoginOrEmailUnauthorized{}
}

/*
GetUserByLoginOrEmailUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetUserByLoginOrEmailUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserByLoginOrEmailUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/lookup][%d] getUserByLoginOrEmailUnauthorized  %+v", 401, o.Payload)
}
func (o *GetUserByLoginOrEmailUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserByLoginOrEmailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByLoginOrEmailForbidden creates a GetUserByLoginOrEmailForbidden with default headers values
func NewGetUserByLoginOrEmailForbidden() *GetUserByLoginOrEmailForbidden {
	return &GetUserByLoginOrEmailForbidden{}
}

/*
GetUserByLoginOrEmailForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetUserByLoginOrEmailForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserByLoginOrEmailForbidden) Error() string {
	return fmt.Sprintf("[GET /users/lookup][%d] getUserByLoginOrEmailForbidden  %+v", 403, o.Payload)
}
func (o *GetUserByLoginOrEmailForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserByLoginOrEmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByLoginOrEmailNotFound creates a GetUserByLoginOrEmailNotFound with default headers values
func NewGetUserByLoginOrEmailNotFound() *GetUserByLoginOrEmailNotFound {
	return &GetUserByLoginOrEmailNotFound{}
}

/*
GetUserByLoginOrEmailNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetUserByLoginOrEmailNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserByLoginOrEmailNotFound) Error() string {
	return fmt.Sprintf("[GET /users/lookup][%d] getUserByLoginOrEmailNotFound  %+v", 404, o.Payload)
}
func (o *GetUserByLoginOrEmailNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserByLoginOrEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByLoginOrEmailInternalServerError creates a GetUserByLoginOrEmailInternalServerError with default headers values
func NewGetUserByLoginOrEmailInternalServerError() *GetUserByLoginOrEmailInternalServerError {
	return &GetUserByLoginOrEmailInternalServerError{}
}

/*
GetUserByLoginOrEmailInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetUserByLoginOrEmailInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetUserByLoginOrEmailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/lookup][%d] getUserByLoginOrEmailInternalServerError  %+v", 500, o.Payload)
}
func (o *GetUserByLoginOrEmailInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserByLoginOrEmailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
