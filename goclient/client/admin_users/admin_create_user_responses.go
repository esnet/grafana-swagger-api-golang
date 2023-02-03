// Code generated by go-swagger; DO NOT EDIT.

package admin_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// AdminCreateUserReader is a Reader for the AdminCreateUser structure.
type AdminCreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminCreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewAdminCreateUserPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminCreateUserOK creates a AdminCreateUserOK with default headers values
func NewAdminCreateUserOK() *AdminCreateUserOK {
	return &AdminCreateUserOK{}
}

/*
AdminCreateUserOK describes a response with status code 200, with default header values.

(empty)
*/
type AdminCreateUserOK struct {
	Payload *models.UserIDDTO
}

func (o *AdminCreateUserOK) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserOK  %+v", 200, o.Payload)
}
func (o *AdminCreateUserOK) GetPayload() *models.UserIDDTO {
	return o.Payload
}

func (o *AdminCreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserIDDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserBadRequest creates a AdminCreateUserBadRequest with default headers values
func NewAdminCreateUserBadRequest() *AdminCreateUserBadRequest {
	return &AdminCreateUserBadRequest{}
}

/*
AdminCreateUserBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AdminCreateUserBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminCreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserBadRequest  %+v", 400, o.Payload)
}
func (o *AdminCreateUserBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserUnauthorized creates a AdminCreateUserUnauthorized with default headers values
func NewAdminCreateUserUnauthorized() *AdminCreateUserUnauthorized {
	return &AdminCreateUserUnauthorized{}
}

/*
AdminCreateUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminCreateUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminCreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminCreateUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserForbidden creates a AdminCreateUserForbidden with default headers values
func NewAdminCreateUserForbidden() *AdminCreateUserForbidden {
	return &AdminCreateUserForbidden{}
}

/*
AdminCreateUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminCreateUserForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminCreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserForbidden  %+v", 403, o.Payload)
}
func (o *AdminCreateUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserPreconditionFailed creates a AdminCreateUserPreconditionFailed with default headers values
func NewAdminCreateUserPreconditionFailed() *AdminCreateUserPreconditionFailed {
	return &AdminCreateUserPreconditionFailed{}
}

/*
AdminCreateUserPreconditionFailed describes a response with status code 412, with default header values.

PreconditionFailedError
*/
type AdminCreateUserPreconditionFailed struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminCreateUserPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserPreconditionFailed  %+v", 412, o.Payload)
}
func (o *AdminCreateUserPreconditionFailed) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserInternalServerError creates a AdminCreateUserInternalServerError with default headers values
func NewAdminCreateUserInternalServerError() *AdminCreateUserInternalServerError {
	return &AdminCreateUserInternalServerError{}
}

/*
AdminCreateUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminCreateUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminCreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /admin/users][%d] adminCreateUserInternalServerError  %+v", 500, o.Payload)
}
func (o *AdminCreateUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminCreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}