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

// ListTokensReader is a Reader for the ListTokens structure.
type ListTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListTokensBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListTokensUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTokensForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTokensInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTokensOK creates a ListTokensOK with default headers values
func NewListTokensOK() *ListTokensOK {
	return &ListTokensOK{}
}

/*
ListTokensOK describes a response with status code 200, with default header values.

(empty)
*/
type ListTokensOK struct {
	Payload *models.TokenDTO
}

func (o *ListTokensOK) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}/tokens][%d] listTokensOK  %+v", 200, o.Payload)
}
func (o *ListTokensOK) GetPayload() *models.TokenDTO {
	return o.Payload
}

func (o *ListTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTokensBadRequest creates a ListTokensBadRequest with default headers values
func NewListTokensBadRequest() *ListTokensBadRequest {
	return &ListTokensBadRequest{}
}

/*
ListTokensBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type ListTokensBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *ListTokensBadRequest) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}/tokens][%d] listTokensBadRequest  %+v", 400, o.Payload)
}
func (o *ListTokensBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListTokensBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTokensUnauthorized creates a ListTokensUnauthorized with default headers values
func NewListTokensUnauthorized() *ListTokensUnauthorized {
	return &ListTokensUnauthorized{}
}

/*
ListTokensUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type ListTokensUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *ListTokensUnauthorized) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}/tokens][%d] listTokensUnauthorized  %+v", 401, o.Payload)
}
func (o *ListTokensUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListTokensUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTokensForbidden creates a ListTokensForbidden with default headers values
func NewListTokensForbidden() *ListTokensForbidden {
	return &ListTokensForbidden{}
}

/*
ListTokensForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ListTokensForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *ListTokensForbidden) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}/tokens][%d] listTokensForbidden  %+v", 403, o.Payload)
}
func (o *ListTokensForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListTokensForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTokensInternalServerError creates a ListTokensInternalServerError with default headers values
func NewListTokensInternalServerError() *ListTokensInternalServerError {
	return &ListTokensInternalServerError{}
}

/*
ListTokensInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ListTokensInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *ListTokensInternalServerError) Error() string {
	return fmt.Sprintf("[GET /serviceaccounts/{serviceAccountId}/tokens][%d] listTokensInternalServerError  %+v", 500, o.Payload)
}
func (o *ListTokensInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListTokensInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
