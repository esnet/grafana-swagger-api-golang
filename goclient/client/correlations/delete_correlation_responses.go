// Code generated by go-swagger; DO NOT EDIT.

package correlations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// DeleteCorrelationReader is a Reader for the DeleteCorrelation structure.
type DeleteCorrelationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCorrelationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCorrelationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCorrelationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCorrelationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCorrelationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCorrelationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCorrelationOK creates a DeleteCorrelationOK with default headers values
func NewDeleteCorrelationOK() *DeleteCorrelationOK {
	return &DeleteCorrelationOK{}
}

/*
DeleteCorrelationOK describes a response with status code 200, with default header values.

(empty)
*/
type DeleteCorrelationOK struct {
	Payload *models.DeleteCorrelationResponseBody
}

func (o *DeleteCorrelationOK) Error() string {
	return fmt.Sprintf("[DELETE /datasources/uid/{uid}/correlations/{correlationUID}][%d] deleteCorrelationOK  %+v", 200, o.Payload)
}
func (o *DeleteCorrelationOK) GetPayload() *models.DeleteCorrelationResponseBody {
	return o.Payload
}

func (o *DeleteCorrelationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCorrelationResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCorrelationUnauthorized creates a DeleteCorrelationUnauthorized with default headers values
func NewDeleteCorrelationUnauthorized() *DeleteCorrelationUnauthorized {
	return &DeleteCorrelationUnauthorized{}
}

/*
DeleteCorrelationUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteCorrelationUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteCorrelationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /datasources/uid/{uid}/correlations/{correlationUID}][%d] deleteCorrelationUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteCorrelationUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteCorrelationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCorrelationForbidden creates a DeleteCorrelationForbidden with default headers values
func NewDeleteCorrelationForbidden() *DeleteCorrelationForbidden {
	return &DeleteCorrelationForbidden{}
}

/*
DeleteCorrelationForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteCorrelationForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteCorrelationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /datasources/uid/{uid}/correlations/{correlationUID}][%d] deleteCorrelationForbidden  %+v", 403, o.Payload)
}
func (o *DeleteCorrelationForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteCorrelationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCorrelationNotFound creates a DeleteCorrelationNotFound with default headers values
func NewDeleteCorrelationNotFound() *DeleteCorrelationNotFound {
	return &DeleteCorrelationNotFound{}
}

/*
DeleteCorrelationNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteCorrelationNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteCorrelationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /datasources/uid/{uid}/correlations/{correlationUID}][%d] deleteCorrelationNotFound  %+v", 404, o.Payload)
}
func (o *DeleteCorrelationNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteCorrelationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCorrelationInternalServerError creates a DeleteCorrelationInternalServerError with default headers values
func NewDeleteCorrelationInternalServerError() *DeleteCorrelationInternalServerError {
	return &DeleteCorrelationInternalServerError{}
}

/*
DeleteCorrelationInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteCorrelationInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteCorrelationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /datasources/uid/{uid}/correlations/{correlationUID}][%d] deleteCorrelationInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteCorrelationInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteCorrelationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
