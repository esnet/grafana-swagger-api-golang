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

// RemoveOrgUserForCurrentOrgReader is a Reader for the RemoveOrgUserForCurrentOrg structure.
type RemoveOrgUserForCurrentOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveOrgUserForCurrentOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveOrgUserForCurrentOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveOrgUserForCurrentOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveOrgUserForCurrentOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveOrgUserForCurrentOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveOrgUserForCurrentOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveOrgUserForCurrentOrgOK creates a RemoveOrgUserForCurrentOrgOK with default headers values
func NewRemoveOrgUserForCurrentOrgOK() *RemoveOrgUserForCurrentOrgOK {
	return &RemoveOrgUserForCurrentOrgOK{}
}

/*
RemoveOrgUserForCurrentOrgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type RemoveOrgUserForCurrentOrgOK struct {
	Payload *models.SuccessResponseBody
}

func (o *RemoveOrgUserForCurrentOrgOK) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgOK  %+v", 200, o.Payload)
}
func (o *RemoveOrgUserForCurrentOrgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForCurrentOrgBadRequest creates a RemoveOrgUserForCurrentOrgBadRequest with default headers values
func NewRemoveOrgUserForCurrentOrgBadRequest() *RemoveOrgUserForCurrentOrgBadRequest {
	return &RemoveOrgUserForCurrentOrgBadRequest{}
}

/*
RemoveOrgUserForCurrentOrgBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RemoveOrgUserForCurrentOrgBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveOrgUserForCurrentOrgBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveOrgUserForCurrentOrgBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForCurrentOrgUnauthorized creates a RemoveOrgUserForCurrentOrgUnauthorized with default headers values
func NewRemoveOrgUserForCurrentOrgUnauthorized() *RemoveOrgUserForCurrentOrgUnauthorized {
	return &RemoveOrgUserForCurrentOrgUnauthorized{}
}

/*
RemoveOrgUserForCurrentOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RemoveOrgUserForCurrentOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveOrgUserForCurrentOrgUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgUnauthorized  %+v", 401, o.Payload)
}
func (o *RemoveOrgUserForCurrentOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForCurrentOrgForbidden creates a RemoveOrgUserForCurrentOrgForbidden with default headers values
func NewRemoveOrgUserForCurrentOrgForbidden() *RemoveOrgUserForCurrentOrgForbidden {
	return &RemoveOrgUserForCurrentOrgForbidden{}
}

/*
RemoveOrgUserForCurrentOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RemoveOrgUserForCurrentOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveOrgUserForCurrentOrgForbidden) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgForbidden  %+v", 403, o.Payload)
}
func (o *RemoveOrgUserForCurrentOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveOrgUserForCurrentOrgInternalServerError creates a RemoveOrgUserForCurrentOrgInternalServerError with default headers values
func NewRemoveOrgUserForCurrentOrgInternalServerError() *RemoveOrgUserForCurrentOrgInternalServerError {
	return &RemoveOrgUserForCurrentOrgInternalServerError{}
}

/*
RemoveOrgUserForCurrentOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RemoveOrgUserForCurrentOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RemoveOrgUserForCurrentOrgInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /org/users/{user_id}][%d] removeOrgUserForCurrentOrgInternalServerError  %+v", 500, o.Payload)
}
func (o *RemoveOrgUserForCurrentOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RemoveOrgUserForCurrentOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
