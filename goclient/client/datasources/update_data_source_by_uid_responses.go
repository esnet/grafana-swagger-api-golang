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

// UpdateDataSourceByUIDReader is a Reader for the UpdateDataSourceByUID structure.
type UpdateDataSourceByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDataSourceByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDataSourceByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDataSourceByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDataSourceByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDataSourceByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDataSourceByUIDOK creates a UpdateDataSourceByUIDOK with default headers values
func NewUpdateDataSourceByUIDOK() *UpdateDataSourceByUIDOK {
	return &UpdateDataSourceByUIDOK{}
}

/*
UpdateDataSourceByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateDataSourceByUIDOK struct {
	Payload *models.UpdateDataSourceByUIDOKBody
}

func (o *UpdateDataSourceByUIDOK) Error() string {
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidOK  %+v", 200, o.Payload)
}
func (o *UpdateDataSourceByUIDOK) GetPayload() *models.UpdateDataSourceByUIDOKBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDataSourceByUIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDataSourceByUIDUnauthorized creates a UpdateDataSourceByUIDUnauthorized with default headers values
func NewUpdateDataSourceByUIDUnauthorized() *UpdateDataSourceByUIDUnauthorized {
	return &UpdateDataSourceByUIDUnauthorized{}
}

/*
UpdateDataSourceByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateDataSourceByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateDataSourceByUIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateDataSourceByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDataSourceByUIDForbidden creates a UpdateDataSourceByUIDForbidden with default headers values
func NewUpdateDataSourceByUIDForbidden() *UpdateDataSourceByUIDForbidden {
	return &UpdateDataSourceByUIDForbidden{}
}

/*
UpdateDataSourceByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateDataSourceByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateDataSourceByUIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidForbidden  %+v", 403, o.Payload)
}
func (o *UpdateDataSourceByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDataSourceByUIDInternalServerError creates a UpdateDataSourceByUIDInternalServerError with default headers values
func NewUpdateDataSourceByUIDInternalServerError() *UpdateDataSourceByUIDInternalServerError {
	return &UpdateDataSourceByUIDInternalServerError{}
}

/*
UpdateDataSourceByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateDataSourceByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateDataSourceByUIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /datasources/uid/{uid}][%d] updateDataSourceByUidInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateDataSourceByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateDataSourceByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}