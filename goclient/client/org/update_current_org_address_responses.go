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

// UpdateCurrentOrgAddressReader is a Reader for the UpdateCurrentOrgAddress structure.
type UpdateCurrentOrgAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCurrentOrgAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCurrentOrgAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCurrentOrgAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCurrentOrgAddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCurrentOrgAddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCurrentOrgAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCurrentOrgAddressOK creates a UpdateCurrentOrgAddressOK with default headers values
func NewUpdateCurrentOrgAddressOK() *UpdateCurrentOrgAddressOK {
	return &UpdateCurrentOrgAddressOK{}
}

/*
UpdateCurrentOrgAddressOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateCurrentOrgAddressOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateCurrentOrgAddressOK) Error() string {
	return fmt.Sprintf("[PUT /org/address][%d] updateCurrentOrgAddressOK  %+v", 200, o.Payload)
}
func (o *UpdateCurrentOrgAddressOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCurrentOrgAddressBadRequest creates a UpdateCurrentOrgAddressBadRequest with default headers values
func NewUpdateCurrentOrgAddressBadRequest() *UpdateCurrentOrgAddressBadRequest {
	return &UpdateCurrentOrgAddressBadRequest{}
}

/*
UpdateCurrentOrgAddressBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateCurrentOrgAddressBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateCurrentOrgAddressBadRequest) Error() string {
	return fmt.Sprintf("[PUT /org/address][%d] updateCurrentOrgAddressBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCurrentOrgAddressBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCurrentOrgAddressUnauthorized creates a UpdateCurrentOrgAddressUnauthorized with default headers values
func NewUpdateCurrentOrgAddressUnauthorized() *UpdateCurrentOrgAddressUnauthorized {
	return &UpdateCurrentOrgAddressUnauthorized{}
}

/*
UpdateCurrentOrgAddressUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateCurrentOrgAddressUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateCurrentOrgAddressUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /org/address][%d] updateCurrentOrgAddressUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateCurrentOrgAddressUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgAddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCurrentOrgAddressForbidden creates a UpdateCurrentOrgAddressForbidden with default headers values
func NewUpdateCurrentOrgAddressForbidden() *UpdateCurrentOrgAddressForbidden {
	return &UpdateCurrentOrgAddressForbidden{}
}

/*
UpdateCurrentOrgAddressForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateCurrentOrgAddressForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateCurrentOrgAddressForbidden) Error() string {
	return fmt.Sprintf("[PUT /org/address][%d] updateCurrentOrgAddressForbidden  %+v", 403, o.Payload)
}
func (o *UpdateCurrentOrgAddressForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgAddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCurrentOrgAddressInternalServerError creates a UpdateCurrentOrgAddressInternalServerError with default headers values
func NewUpdateCurrentOrgAddressInternalServerError() *UpdateCurrentOrgAddressInternalServerError {
	return &UpdateCurrentOrgAddressInternalServerError{}
}

/*
UpdateCurrentOrgAddressInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateCurrentOrgAddressInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateCurrentOrgAddressInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /org/address][%d] updateCurrentOrgAddressInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateCurrentOrgAddressInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateCurrentOrgAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}