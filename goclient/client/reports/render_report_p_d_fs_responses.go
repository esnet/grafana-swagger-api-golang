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

// RenderReportPDFsReader is a Reader for the RenderReportPDFs structure.
type RenderReportPDFsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenderReportPDFsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenderReportPDFsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenderReportPDFsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRenderReportPDFsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRenderReportPDFsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRenderReportPDFsOK creates a RenderReportPDFsOK with default headers values
func NewRenderReportPDFsOK() *RenderReportPDFsOK {
	return &RenderReportPDFsOK{}
}

/*
RenderReportPDFsOK describes a response with status code 200, with default header values.

(empty)
*/
type RenderReportPDFsOK struct {
	Payload []uint8
}

func (o *RenderReportPDFsOK) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdfs][%d] renderReportPDFsOK  %+v", 200, o.Payload)
}
func (o *RenderReportPDFsOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *RenderReportPDFsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFsBadRequest creates a RenderReportPDFsBadRequest with default headers values
func NewRenderReportPDFsBadRequest() *RenderReportPDFsBadRequest {
	return &RenderReportPDFsBadRequest{}
}

/*
RenderReportPDFsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RenderReportPDFsBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *RenderReportPDFsBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdfs][%d] renderReportPDFsBadRequest  %+v", 400, o.Payload)
}
func (o *RenderReportPDFsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFsUnauthorized creates a RenderReportPDFsUnauthorized with default headers values
func NewRenderReportPDFsUnauthorized() *RenderReportPDFsUnauthorized {
	return &RenderReportPDFsUnauthorized{}
}

/*
RenderReportPDFsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RenderReportPDFsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *RenderReportPDFsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdfs][%d] renderReportPDFsUnauthorized  %+v", 401, o.Payload)
}
func (o *RenderReportPDFsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFsInternalServerError creates a RenderReportPDFsInternalServerError with default headers values
func NewRenderReportPDFsInternalServerError() *RenderReportPDFsInternalServerError {
	return &RenderReportPDFsInternalServerError{}
}

/*
RenderReportPDFsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RenderReportPDFsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RenderReportPDFsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdfs][%d] renderReportPDFsInternalServerError  %+v", 500, o.Payload)
}
func (o *RenderReportPDFsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
