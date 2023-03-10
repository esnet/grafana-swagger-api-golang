// Code generated by go-swagger; DO NOT EDIT.

package org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// UpdateCurrentOrgReader is a Reader for the UpdateCurrentOrg structure.
type UpdateCurrentOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCurrentOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCurrentOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCurrentOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCurrentOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCurrentOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCurrentOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCurrentOrgOK creates a UpdateCurrentOrgOK with default headers values
func NewUpdateCurrentOrgOK() *UpdateCurrentOrgOK {
	return &UpdateCurrentOrgOK{}
}

/*
UpdateCurrentOrgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateCurrentOrgOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateCurrentOrgOK) Error() string {
	return fmt.Sprintf("[PUT /org][%d] updateCurrentOrgOK  %+v", 200, o.Payload)
}
func (o *UpdateCurrentOrgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCurrentOrgBadRequest creates a UpdateCurrentOrgBadRequest with default headers values
func NewUpdateCurrentOrgBadRequest() *UpdateCurrentOrgBadRequest {
	return &UpdateCurrentOrgBadRequest{}
}

/*
UpdateCurrentOrgBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateCurrentOrgBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateCurrentOrgBadRequest) Error() string {
	return fmt.Sprintf("[PUT /org][%d] updateCurrentOrgBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCurrentOrgBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCurrentOrgUnauthorized creates a UpdateCurrentOrgUnauthorized with default headers values
func NewUpdateCurrentOrgUnauthorized() *UpdateCurrentOrgUnauthorized {
	return &UpdateCurrentOrgUnauthorized{}
}

/*
UpdateCurrentOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateCurrentOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateCurrentOrgUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /org][%d] updateCurrentOrgUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateCurrentOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCurrentOrgForbidden creates a UpdateCurrentOrgForbidden with default headers values
func NewUpdateCurrentOrgForbidden() *UpdateCurrentOrgForbidden {
	return &UpdateCurrentOrgForbidden{}
}

/*
UpdateCurrentOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateCurrentOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateCurrentOrgForbidden) Error() string {
	return fmt.Sprintf("[PUT /org][%d] updateCurrentOrgForbidden  %+v", 403, o.Payload)
}
func (o *UpdateCurrentOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCurrentOrgInternalServerError creates a UpdateCurrentOrgInternalServerError with default headers values
func NewUpdateCurrentOrgInternalServerError() *UpdateCurrentOrgInternalServerError {
	return &UpdateCurrentOrgInternalServerError{}
}

/*
UpdateCurrentOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateCurrentOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateCurrentOrgInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /org][%d] updateCurrentOrgInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateCurrentOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
