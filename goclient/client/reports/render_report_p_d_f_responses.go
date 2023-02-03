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

// RenderReportPDFReader is a Reader for the RenderReportPDF structure.
type RenderReportPDFReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenderReportPDFReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenderReportPDFOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenderReportPDFBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRenderReportPDFUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRenderReportPDFInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRenderReportPDFOK creates a RenderReportPDFOK with default headers values
func NewRenderReportPDFOK() *RenderReportPDFOK {
	return &RenderReportPDFOK{}
}

/*
RenderReportPDFOK describes a response with status code 200, with default header values.

(empty)
*/
type RenderReportPDFOK struct {
	Payload []uint8
}

func (o *RenderReportPDFOK) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFOK  %+v", 200, o.Payload)
}
func (o *RenderReportPDFOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *RenderReportPDFOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFBadRequest creates a RenderReportPDFBadRequest with default headers values
func NewRenderReportPDFBadRequest() *RenderReportPDFBadRequest {
	return &RenderReportPDFBadRequest{}
}

/*
RenderReportPDFBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type RenderReportPDFBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *RenderReportPDFBadRequest) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFBadRequest  %+v", 400, o.Payload)
}
func (o *RenderReportPDFBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFUnauthorized creates a RenderReportPDFUnauthorized with default headers values
func NewRenderReportPDFUnauthorized() *RenderReportPDFUnauthorized {
	return &RenderReportPDFUnauthorized{}
}

/*
RenderReportPDFUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type RenderReportPDFUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *RenderReportPDFUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFUnauthorized  %+v", 401, o.Payload)
}
func (o *RenderReportPDFUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderReportPDFInternalServerError creates a RenderReportPDFInternalServerError with default headers values
func NewRenderReportPDFInternalServerError() *RenderReportPDFInternalServerError {
	return &RenderReportPDFInternalServerError{}
}

/*
RenderReportPDFInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type RenderReportPDFInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *RenderReportPDFInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reports/render/pdf/{dashboardID}][%d] renderReportPDFInternalServerError  %+v", 500, o.Payload)
}
func (o *RenderReportPDFInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *RenderReportPDFInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}