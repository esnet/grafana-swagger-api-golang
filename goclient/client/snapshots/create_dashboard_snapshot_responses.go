// Code generated by go-swagger; DO NOT EDIT.

package snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// CreateDashboardSnapshotReader is a Reader for the CreateDashboardSnapshot structure.
type CreateDashboardSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDashboardSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateDashboardSnapshotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDashboardSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDashboardSnapshotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDashboardSnapshotOK creates a CreateDashboardSnapshotOK with default headers values
func NewCreateDashboardSnapshotOK() *CreateDashboardSnapshotOK {
	return &CreateDashboardSnapshotOK{}
}

/*
CreateDashboardSnapshotOK describes a response with status code 200, with default header values.

(empty)
*/
type CreateDashboardSnapshotOK struct {
	Payload *models.CreateDashboardSnapshotOKBody
}

func (o *CreateDashboardSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /snapshots][%d] createDashboardSnapshotOK  %+v", 200, o.Payload)
}
func (o *CreateDashboardSnapshotOK) GetPayload() *models.CreateDashboardSnapshotOKBody {
	return o.Payload
}

func (o *CreateDashboardSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDashboardSnapshotOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardSnapshotUnauthorized creates a CreateDashboardSnapshotUnauthorized with default headers values
func NewCreateDashboardSnapshotUnauthorized() *CreateDashboardSnapshotUnauthorized {
	return &CreateDashboardSnapshotUnauthorized{}
}

/*
CreateDashboardSnapshotUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreateDashboardSnapshotUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateDashboardSnapshotUnauthorized) Error() string {
	return fmt.Sprintf("[POST /snapshots][%d] createDashboardSnapshotUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateDashboardSnapshotUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateDashboardSnapshotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardSnapshotForbidden creates a CreateDashboardSnapshotForbidden with default headers values
func NewCreateDashboardSnapshotForbidden() *CreateDashboardSnapshotForbidden {
	return &CreateDashboardSnapshotForbidden{}
}

/*
CreateDashboardSnapshotForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateDashboardSnapshotForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateDashboardSnapshotForbidden) Error() string {
	return fmt.Sprintf("[POST /snapshots][%d] createDashboardSnapshotForbidden  %+v", 403, o.Payload)
}
func (o *CreateDashboardSnapshotForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateDashboardSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardSnapshotInternalServerError creates a CreateDashboardSnapshotInternalServerError with default headers values
func NewCreateDashboardSnapshotInternalServerError() *CreateDashboardSnapshotInternalServerError {
	return &CreateDashboardSnapshotInternalServerError{}
}

/*
CreateDashboardSnapshotInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateDashboardSnapshotInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *CreateDashboardSnapshotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /snapshots][%d] createDashboardSnapshotInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateDashboardSnapshotInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateDashboardSnapshotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
