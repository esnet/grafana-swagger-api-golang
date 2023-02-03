// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// UpdateSignedInUserReader is a Reader for the UpdateSignedInUser structure.
type UpdateSignedInUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSignedInUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSignedInUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateSignedInUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateSignedInUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSignedInUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSignedInUserOK creates a UpdateSignedInUserOK with default headers values
func NewUpdateSignedInUserOK() *UpdateSignedInUserOK {
	return &UpdateSignedInUserOK{}
}

/*
UpdateSignedInUserOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateSignedInUserOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateSignedInUserOK) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateSignedInUserOK  %+v", 200, o.Payload)
}
func (o *UpdateSignedInUserOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateSignedInUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSignedInUserUnauthorized creates a UpdateSignedInUserUnauthorized with default headers values
func NewUpdateSignedInUserUnauthorized() *UpdateSignedInUserUnauthorized {
	return &UpdateSignedInUserUnauthorized{}
}

/*
UpdateSignedInUserUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateSignedInUserUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateSignedInUserUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateSignedInUserUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateSignedInUserUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateSignedInUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSignedInUserForbidden creates a UpdateSignedInUserForbidden with default headers values
func NewUpdateSignedInUserForbidden() *UpdateSignedInUserForbidden {
	return &UpdateSignedInUserForbidden{}
}

/*
UpdateSignedInUserForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateSignedInUserForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateSignedInUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateSignedInUserForbidden  %+v", 403, o.Payload)
}
func (o *UpdateSignedInUserForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateSignedInUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSignedInUserInternalServerError creates a UpdateSignedInUserInternalServerError with default headers values
func NewUpdateSignedInUserInternalServerError() *UpdateSignedInUserInternalServerError {
	return &UpdateSignedInUserInternalServerError{}
}

/*
UpdateSignedInUserInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateSignedInUserInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateSignedInUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /user][%d] updateSignedInUserInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateSignedInUserInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateSignedInUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}