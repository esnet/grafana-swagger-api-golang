// Code generated by go-swagger; DO NOT EDIT.

package service_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// UpdateServiceAccountReader is a Reader for the UpdateServiceAccount structure.
type UpdateServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateServiceAccountOK creates a UpdateServiceAccountOK with default headers values
func NewUpdateServiceAccountOK() *UpdateServiceAccountOK {
	return &UpdateServiceAccountOK{}
}

/*
UpdateServiceAccountOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateServiceAccountOK struct {
	Payload *models.UpdateServiceAccountOKBody
}

func (o *UpdateServiceAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountOK  %+v", 200, o.Payload)
}
func (o *UpdateServiceAccountOK) GetPayload() *models.UpdateServiceAccountOKBody {
	return o.Payload
}

func (o *UpdateServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateServiceAccountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountBadRequest creates a UpdateServiceAccountBadRequest with default headers values
func NewUpdateServiceAccountBadRequest() *UpdateServiceAccountBadRequest {
	return &UpdateServiceAccountBadRequest{}
}

/*
UpdateServiceAccountBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateServiceAccountBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateServiceAccountBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateServiceAccountBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountUnauthorized creates a UpdateServiceAccountUnauthorized with default headers values
func NewUpdateServiceAccountUnauthorized() *UpdateServiceAccountUnauthorized {
	return &UpdateServiceAccountUnauthorized{}
}

/*
UpdateServiceAccountUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateServiceAccountUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateServiceAccountUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountForbidden creates a UpdateServiceAccountForbidden with default headers values
func NewUpdateServiceAccountForbidden() *UpdateServiceAccountForbidden {
	return &UpdateServiceAccountForbidden{}
}

/*
UpdateServiceAccountForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateServiceAccountForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountForbidden  %+v", 403, o.Payload)
}
func (o *UpdateServiceAccountForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountNotFound creates a UpdateServiceAccountNotFound with default headers values
func NewUpdateServiceAccountNotFound() *UpdateServiceAccountNotFound {
	return &UpdateServiceAccountNotFound{}
}

/*
UpdateServiceAccountNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateServiceAccountNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountNotFound  %+v", 404, o.Payload)
}
func (o *UpdateServiceAccountNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountInternalServerError creates a UpdateServiceAccountInternalServerError with default headers values
func NewUpdateServiceAccountInternalServerError() *UpdateServiceAccountInternalServerError {
	return &UpdateServiceAccountInternalServerError{}
}

/*
UpdateServiceAccountInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateServiceAccountInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateServiceAccountInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}