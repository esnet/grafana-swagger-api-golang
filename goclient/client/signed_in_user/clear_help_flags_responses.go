// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// ClearHelpFlagsReader is a Reader for the ClearHelpFlags structure.
type ClearHelpFlagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClearHelpFlagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClearHelpFlagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewClearHelpFlagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewClearHelpFlagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewClearHelpFlagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewClearHelpFlagsOK creates a ClearHelpFlagsOK with default headers values
func NewClearHelpFlagsOK() *ClearHelpFlagsOK {
	return &ClearHelpFlagsOK{}
}

/*
ClearHelpFlagsOK describes a response with status code 200, with default header values.

(empty)
*/
type ClearHelpFlagsOK struct {
	Payload *models.ClearHelpFlagsOKBody
}

func (o *ClearHelpFlagsOK) Error() string {
	return fmt.Sprintf("[GET /user/helpflags/clear][%d] clearHelpFlagsOK  %+v", 200, o.Payload)
}
func (o *ClearHelpFlagsOK) GetPayload() *models.ClearHelpFlagsOKBody {
	return o.Payload
}

func (o *ClearHelpFlagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClearHelpFlagsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClearHelpFlagsUnauthorized creates a ClearHelpFlagsUnauthorized with default headers values
func NewClearHelpFlagsUnauthorized() *ClearHelpFlagsUnauthorized {
	return &ClearHelpFlagsUnauthorized{}
}

/*
ClearHelpFlagsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type ClearHelpFlagsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *ClearHelpFlagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/helpflags/clear][%d] clearHelpFlagsUnauthorized  %+v", 401, o.Payload)
}
func (o *ClearHelpFlagsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ClearHelpFlagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClearHelpFlagsForbidden creates a ClearHelpFlagsForbidden with default headers values
func NewClearHelpFlagsForbidden() *ClearHelpFlagsForbidden {
	return &ClearHelpFlagsForbidden{}
}

/*
ClearHelpFlagsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ClearHelpFlagsForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *ClearHelpFlagsForbidden) Error() string {
	return fmt.Sprintf("[GET /user/helpflags/clear][%d] clearHelpFlagsForbidden  %+v", 403, o.Payload)
}
func (o *ClearHelpFlagsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ClearHelpFlagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClearHelpFlagsInternalServerError creates a ClearHelpFlagsInternalServerError with default headers values
func NewClearHelpFlagsInternalServerError() *ClearHelpFlagsInternalServerError {
	return &ClearHelpFlagsInternalServerError{}
}

/*
ClearHelpFlagsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ClearHelpFlagsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *ClearHelpFlagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user/helpflags/clear][%d] clearHelpFlagsInternalServerError  %+v", 500, o.Payload)
}
func (o *ClearHelpFlagsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ClearHelpFlagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}