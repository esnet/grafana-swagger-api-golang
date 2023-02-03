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

// RetrieveServiceAccountReader is a Reader for the RetrieveServiceAccount structure.
type RetrieveServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRetrieveServiceAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRetrieveServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRetrieveServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRetrieveServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRetrieveServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrieveServiceAccountOK creates a RetrieveServiceAccountOK with default headers values
func NewRetrieveServiceAccountOK() *RetrieveServiceAccountOK {
	return &RetrieveServiceAccountOK{}
}

/*
RetrieveServiceAccountOK describes a response with status code 200, with default header values.

(empty)
*/
type RetrieveServiceAccountOK struct {
	Payload *models.ServiceAccountDTO
}

func (o *RetrieveServiceAccountOK) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}][%d] retrieveServiceAccountOK  %+v", 200, o.Payload)
}
func (o *RetrieveServiceAccountOK) GetPayload() *models.ServiceAccountDTO {
	return o.Payload
}

func (o *RetrieveServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceAccountDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveServiceAccountBadRequest creates a RetrieveServiceAccountBadRequest with default headers values
func NewRetrieveServiceAccountBadRequest() *RetrieveServiceAccountBadRequest {
	return &RetrieveServiceAccountBadRequest{}
}

/*
RetrieveServiceAccountBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RetrieveServiceAccountBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *RetrieveServiceAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}][%d] retrieveServiceAccountBadRequest  %+v", 400, o.Payload)
}
func (o *RetrieveServiceAccountBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RetrieveServiceAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveServiceAccountUnauthorized creates a RetrieveServiceAccountUnauthorized with default headers values
func NewRetrieveServiceAccountUnauthorized() *RetrieveServiceAccountUnauthorized {
	return &RetrieveServiceAccountUnauthorized{}
}

/*
RetrieveServiceAccountUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RetrieveServiceAccountUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *RetrieveServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}][%d] retrieveServiceAccountUnauthorized  %+v", 401, o.Payload)
}
func (o *RetrieveServiceAccountUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RetrieveServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveServiceAccountForbidden creates a RetrieveServiceAccountForbidden with default headers values
func NewRetrieveServiceAccountForbidden() *RetrieveServiceAccountForbidden {
	return &RetrieveServiceAccountForbidden{}
}

/*
RetrieveServiceAccountForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type RetrieveServiceAccountForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *RetrieveServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}][%d] retrieveServiceAccountForbidden  %+v", 403, o.Payload)
}
func (o *RetrieveServiceAccountForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RetrieveServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveServiceAccountNotFound creates a RetrieveServiceAccountNotFound with default headers values
func NewRetrieveServiceAccountNotFound() *RetrieveServiceAccountNotFound {
	return &RetrieveServiceAccountNotFound{}
}

/*
RetrieveServiceAccountNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type RetrieveServiceAccountNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *RetrieveServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}][%d] retrieveServiceAccountNotFound  %+v", 404, o.Payload)
}
func (o *RetrieveServiceAccountNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RetrieveServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveServiceAccountInternalServerError creates a RetrieveServiceAccountInternalServerError with default headers values
func NewRetrieveServiceAccountInternalServerError() *RetrieveServiceAccountInternalServerError {
	return &RetrieveServiceAccountInternalServerError{}
}

/*
RetrieveServiceAccountInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RetrieveServiceAccountInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RetrieveServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}][%d] retrieveServiceAccountInternalServerError  %+v", 500, o.Payload)
}
func (o *RetrieveServiceAccountInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RetrieveServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
