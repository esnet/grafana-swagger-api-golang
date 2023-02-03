// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// CallDatasourceResourceWithUIDReader is a Reader for the CallDatasourceResourceWithUID structure.
type CallDatasourceResourceWithUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CallDatasourceResourceWithUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCallDatasourceResourceWithUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCallDatasourceResourceWithUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCallDatasourceResourceWithUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCallDatasourceResourceWithUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCallDatasourceResourceWithUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCallDatasourceResourceWithUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCallDatasourceResourceWithUIDOK creates a CallDatasourceResourceWithUIDOK with default headers values
func NewCallDatasourceResourceWithUIDOK() *CallDatasourceResourceWithUIDOK {
	return &CallDatasourceResourceWithUIDOK{}
}

/*
CallDatasourceResourceWithUIDOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type CallDatasourceResourceWithUIDOK struct {
	Payload *models.SuccessResponseBody
}

func (o *CallDatasourceResourceWithUIDOK) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/resources/{datasource_proxy_route}][%d] callDatasourceResourceWithUidOK  %+v", 200, o.Payload)
}
func (o *CallDatasourceResourceWithUIDOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *CallDatasourceResourceWithUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCallDatasourceResourceWithUIDBadRequest creates a CallDatasourceResourceWithUIDBadRequest with default headers values
func NewCallDatasourceResourceWithUIDBadRequest() *CallDatasourceResourceWithUIDBadRequest {
	return &CallDatasourceResourceWithUIDBadRequest{}
}

/*
CallDatasourceResourceWithUIDBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CallDatasourceResourceWithUIDBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *CallDatasourceResourceWithUIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/resources/{datasource_proxy_route}][%d] callDatasourceResourceWithUidBadRequest  %+v", 400, o.Payload)
}
func (o *CallDatasourceResourceWithUIDBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CallDatasourceResourceWithUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCallDatasourceResourceWithUIDUnauthorized creates a CallDatasourceResourceWithUIDUnauthorized with default headers values
func NewCallDatasourceResourceWithUIDUnauthorized() *CallDatasourceResourceWithUIDUnauthorized {
	return &CallDatasourceResourceWithUIDUnauthorized{}
}

/*
CallDatasourceResourceWithUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CallDatasourceResourceWithUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *CallDatasourceResourceWithUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/resources/{datasource_proxy_route}][%d] callDatasourceResourceWithUidUnauthorized  %+v", 401, o.Payload)
}
func (o *CallDatasourceResourceWithUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CallDatasourceResourceWithUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCallDatasourceResourceWithUIDForbidden creates a CallDatasourceResourceWithUIDForbidden with default headers values
func NewCallDatasourceResourceWithUIDForbidden() *CallDatasourceResourceWithUIDForbidden {
	return &CallDatasourceResourceWithUIDForbidden{}
}

/*
CallDatasourceResourceWithUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CallDatasourceResourceWithUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *CallDatasourceResourceWithUIDForbidden) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/resources/{datasource_proxy_route}][%d] callDatasourceResourceWithUidForbidden  %+v", 403, o.Payload)
}
func (o *CallDatasourceResourceWithUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CallDatasourceResourceWithUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCallDatasourceResourceWithUIDNotFound creates a CallDatasourceResourceWithUIDNotFound with default headers values
func NewCallDatasourceResourceWithUIDNotFound() *CallDatasourceResourceWithUIDNotFound {
	return &CallDatasourceResourceWithUIDNotFound{}
}

/*
CallDatasourceResourceWithUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type CallDatasourceResourceWithUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *CallDatasourceResourceWithUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/resources/{datasource_proxy_route}][%d] callDatasourceResourceWithUidNotFound  %+v", 404, o.Payload)
}
func (o *CallDatasourceResourceWithUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CallDatasourceResourceWithUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCallDatasourceResourceWithUIDInternalServerError creates a CallDatasourceResourceWithUIDInternalServerError with default headers values
func NewCallDatasourceResourceWithUIDInternalServerError() *CallDatasourceResourceWithUIDInternalServerError {
	return &CallDatasourceResourceWithUIDInternalServerError{}
}

/*
CallDatasourceResourceWithUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CallDatasourceResourceWithUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *CallDatasourceResourceWithUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /datasources/uid/{uid}/resources/{datasource_proxy_route}][%d] callDatasourceResourceWithUidInternalServerError  %+v", 500, o.Payload)
}
func (o *CallDatasourceResourceWithUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CallDatasourceResourceWithUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
