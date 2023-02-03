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

// UpdateUserQuotaReader is a Reader for the UpdateUserQuota structure.
type UpdateUserQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserQuotaOK creates a UpdateUserQuotaOK with default headers values
func NewUpdateUserQuotaOK() *UpdateUserQuotaOK {
	return &UpdateUserQuotaOK{}
}

/*
UpdateUserQuotaOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateUserQuotaOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateUserQuotaOK) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/quotas/{quota_target}][%d] updateUserQuotaOK  %+v", 200, o.Payload)
}
func (o *UpdateUserQuotaOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateUserQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserQuotaUnauthorized creates a UpdateUserQuotaUnauthorized with default headers values
func NewUpdateUserQuotaUnauthorized() *UpdateUserQuotaUnauthorized {
	return &UpdateUserQuotaUnauthorized{}
}

/*
UpdateUserQuotaUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateUserQuotaUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateUserQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/quotas/{quota_target}][%d] updateUserQuotaUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateUserQuotaUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateUserQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserQuotaForbidden creates a UpdateUserQuotaForbidden with default headers values
func NewUpdateUserQuotaForbidden() *UpdateUserQuotaForbidden {
	return &UpdateUserQuotaForbidden{}
}

/*
UpdateUserQuotaForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateUserQuotaForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateUserQuotaForbidden) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/quotas/{quota_target}][%d] updateUserQuotaForbidden  %+v", 403, o.Payload)
}
func (o *UpdateUserQuotaForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateUserQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserQuotaNotFound creates a UpdateUserQuotaNotFound with default headers values
func NewUpdateUserQuotaNotFound() *UpdateUserQuotaNotFound {
	return &UpdateUserQuotaNotFound{}
}

/*
UpdateUserQuotaNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateUserQuotaNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateUserQuotaNotFound) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/quotas/{quota_target}][%d] updateUserQuotaNotFound  %+v", 404, o.Payload)
}
func (o *UpdateUserQuotaNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateUserQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserQuotaInternalServerError creates a UpdateUserQuotaInternalServerError with default headers values
func NewUpdateUserQuotaInternalServerError() *UpdateUserQuotaInternalServerError {
	return &UpdateUserQuotaInternalServerError{}
}

/*
UpdateUserQuotaInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateUserQuotaInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateUserQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /admin/users/{user_id}/quotas/{quota_target}][%d] updateUserQuotaInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateUserQuotaInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateUserQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}