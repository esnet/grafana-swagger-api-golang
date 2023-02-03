// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetReportReader is a Reader for the GetReport structure.
type GetReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReportOK creates a GetReportOK with default headers values
func NewGetReportOK() *GetReportOK {
	return &GetReportOK{}
}

/*
GetReportOK describes a response with status code 200, with default header values.

(empty)
*/
type GetReportOK struct {
	Payload *models.ConfigDTO
}

func (o *GetReportOK) Error() string {
	return fmt.Sprintf("[GET /reports/{id}][%d] getReportOK  %+v", 200, o.Payload)
}
func (o *GetReportOK) GetPayload() *models.ConfigDTO {
	return o.Payload
}

func (o *GetReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportBadRequest creates a GetReportBadRequest with default headers values
func NewGetReportBadRequest() *GetReportBadRequest {
	return &GetReportBadRequest{}
}

/*
GetReportBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type GetReportBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *GetReportBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/{id}][%d] getReportBadRequest  %+v", 400, o.Payload)
}
func (o *GetReportBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportUnauthorized creates a GetReportUnauthorized with default headers values
func NewGetReportUnauthorized() *GetReportUnauthorized {
	return &GetReportUnauthorized{}
}

/*
GetReportUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetReportUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetReportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/{id}][%d] getReportUnauthorized  %+v", 401, o.Payload)
}
func (o *GetReportUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportForbidden creates a GetReportForbidden with default headers values
func NewGetReportForbidden() *GetReportForbidden {
	return &GetReportForbidden{}
}

/*
GetReportForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetReportForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetReportForbidden) Error() string {
	return fmt.Sprintf("[GET /reports/{id}][%d] getReportForbidden  %+v", 403, o.Payload)
}
func (o *GetReportForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportNotFound creates a GetReportNotFound with default headers values
func NewGetReportNotFound() *GetReportNotFound {
	return &GetReportNotFound{}
}

/*
GetReportNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetReportNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetReportNotFound) Error() string {
	return fmt.Sprintf("[GET /reports/{id}][%d] getReportNotFound  %+v", 404, o.Payload)
}
func (o *GetReportNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportInternalServerError creates a GetReportInternalServerError with default headers values
func NewGetReportInternalServerError() *GetReportInternalServerError {
	return &GetReportInternalServerError{}
}

/*
GetReportInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetReportInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetReportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/{id}][%d] getReportInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReportInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}