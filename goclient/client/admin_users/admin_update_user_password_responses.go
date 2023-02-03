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

// AdminUpdateUserPasswordReader is a Reader for the AdminUpdateUserPassword structure.
type AdminUpdateUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateUserPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateUserPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateUserPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminUpdateUserPasswordOK creates a AdminUpdateUserPasswordOK with default headers values
func NewAdminUpdateUserPasswordOK() *AdminUpdateUserPasswordOK {
	return &AdminUpdateUserPasswordOK{}
}

/*
AdminUpdateUserPasswordOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminUpdateUserPasswordOK struct {
	Payload *models.SuccessResponseBody
}

func (o *AdminUpdateUserPasswordOK) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordOK  %+v", 200, o.Payload)
}
func (o *AdminUpdateUserPasswordOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordBadRequest creates a AdminUpdateUserPasswordBadRequest with default headers values
func NewAdminUpdateUserPasswordBadRequest() *AdminUpdateUserPasswordBadRequest {
	return &AdminUpdateUserPasswordBadRequest{}
}

/*
AdminUpdateUserPasswordBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AdminUpdateUserPasswordBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminUpdateUserPasswordBadRequest) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordBadRequest  %+v", 400, o.Payload)
}
func (o *AdminUpdateUserPasswordBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordUnauthorized creates a AdminUpdateUserPasswordUnauthorized with default headers values
func NewAdminUpdateUserPasswordUnauthorized() *AdminUpdateUserPasswordUnauthorized {
	return &AdminUpdateUserPasswordUnauthorized{}
}

/*
AdminUpdateUserPasswordUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminUpdateUserPasswordUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminUpdateUserPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordUnauthorized  %+v", 401, o.Payload)
}
func (o *AdminUpdateUserPasswordUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordForbidden creates a AdminUpdateUserPasswordForbidden with default headers values
func NewAdminUpdateUserPasswordForbidden() *AdminUpdateUserPasswordForbidden {
	return &AdminUpdateUserPasswordForbidden{}
}

/*
AdminUpdateUserPasswordForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminUpdateUserPasswordForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminUpdateUserPasswordForbidden) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordForbidden  %+v", 403, o.Payload)
}
func (o *AdminUpdateUserPasswordForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordInternalServerError creates a AdminUpdateUserPasswordInternalServerError with default headers values
func NewAdminUpdateUserPasswordInternalServerError() *AdminUpdateUserPasswordInternalServerError {
	return &AdminUpdateUserPasswordInternalServerError{}
}

/*
AdminUpdateUserPasswordInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminUpdateUserPasswordInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *AdminUpdateUserPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordInternalServerError  %+v", 500, o.Payload)
}
func (o *AdminUpdateUserPasswordInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}