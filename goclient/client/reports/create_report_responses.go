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

// CreateReportReader is a Reader for the CreateReport structure.
type CreateReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateReportOK creates a CreateReportOK with default headers values
func NewCreateReportOK() *CreateReportOK {
	return &CreateReportOK{}
}

/*
CreateReportOK describes a response with status code 200, with default header values.

(empty)
*/
type CreateReportOK struct {
	Payload *models.CreateReportOKBody
}

func (o *CreateReportOK) Error() string {
	return fmt.Sprintf("[POST /reports][%d] createReportOK  %+v", 200, o.Payload)
}
func (o *CreateReportOK) GetPayload() *models.CreateReportOKBody {
	return o.Payload
}

func (o *CreateReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateReportOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportBadRequest creates a CreateReportBadRequest with default headers values
func NewCreateReportBadRequest() *CreateReportBadRequest {
	return &CreateReportBadRequest{}
}

/*
CreateReportBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CreateReportBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateReportBadRequest) Error() string {
	return fmt.Sprintf("[POST /reports][%d] createReportBadRequest  %+v", 400, o.Payload)
}
func (o *CreateReportBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportUnauthorized creates a CreateReportUnauthorized with default headers values
func NewCreateReportUnauthorized() *CreateReportUnauthorized {
	return &CreateReportUnauthorized{}
}

/*
CreateReportUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreateReportUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateReportUnauthorized) Error() string {
	return fmt.Sprintf("[POST /reports][%d] createReportUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateReportUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportForbidden creates a CreateReportForbidden with default headers values
func NewCreateReportForbidden() *CreateReportForbidden {
	return &CreateReportForbidden{}
}

/*
CreateReportForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateReportForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateReportForbidden) Error() string {
	return fmt.Sprintf("[POST /reports][%d] createReportForbidden  %+v", 403, o.Payload)
}
func (o *CreateReportForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportNotFound creates a CreateReportNotFound with default headers values
func NewCreateReportNotFound() *CreateReportNotFound {
	return &CreateReportNotFound{}
}

/*
CreateReportNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type CreateReportNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateReportNotFound) Error() string {
	return fmt.Sprintf("[POST /reports][%d] createReportNotFound  %+v", 404, o.Payload)
}
func (o *CreateReportNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateReportInternalServerError creates a CreateReportInternalServerError with default headers values
func NewCreateReportInternalServerError() *CreateReportInternalServerError {
	return &CreateReportInternalServerError{}
}

/*
CreateReportInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateReportInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateReportInternalServerError) Error() string {
	return fmt.Sprintf("[POST /reports][%d] createReportInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateReportInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}