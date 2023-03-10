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

// UserSetUsingOrgReader is a Reader for the UserSetUsingOrg structure.
type UserSetUsingOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSetUsingOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserSetUsingOrgOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserSetUsingOrgBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserSetUsingOrgUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserSetUsingOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserSetUsingOrgInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserSetUsingOrgOK creates a UserSetUsingOrgOK with default headers values
func NewUserSetUsingOrgOK() *UserSetUsingOrgOK {
	return &UserSetUsingOrgOK{}
}

/*
UserSetUsingOrgOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UserSetUsingOrgOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UserSetUsingOrgOK) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgOK  %+v", 200, o.Payload)
}
func (o *UserSetUsingOrgOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSetUsingOrgBadRequest creates a UserSetUsingOrgBadRequest with default headers values
func NewUserSetUsingOrgBadRequest() *UserSetUsingOrgBadRequest {
	return &UserSetUsingOrgBadRequest{}
}

/*
UserSetUsingOrgBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UserSetUsingOrgBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UserSetUsingOrgBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgBadRequest  %+v", 400, o.Payload)
}
func (o *UserSetUsingOrgBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSetUsingOrgUnauthorized creates a UserSetUsingOrgUnauthorized with default headers values
func NewUserSetUsingOrgUnauthorized() *UserSetUsingOrgUnauthorized {
	return &UserSetUsingOrgUnauthorized{}
}

/*
UserSetUsingOrgUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UserSetUsingOrgUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UserSetUsingOrgUnauthorized) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgUnauthorized  %+v", 401, o.Payload)
}
func (o *UserSetUsingOrgUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSetUsingOrgForbidden creates a UserSetUsingOrgForbidden with default headers values
func NewUserSetUsingOrgForbidden() *UserSetUsingOrgForbidden {
	return &UserSetUsingOrgForbidden{}
}

/*
UserSetUsingOrgForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UserSetUsingOrgForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UserSetUsingOrgForbidden) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgForbidden  %+v", 403, o.Payload)
}
func (o *UserSetUsingOrgForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserSetUsingOrgInternalServerError creates a UserSetUsingOrgInternalServerError with default headers values
func NewUserSetUsingOrgInternalServerError() *UserSetUsingOrgInternalServerError {
	return &UserSetUsingOrgInternalServerError{}
}

/*
UserSetUsingOrgInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UserSetUsingOrgInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UserSetUsingOrgInternalServerError) Error() string {
	return fmt.Sprintf("[POST /user/using/{org_id}][%d] userSetUsingOrgInternalServerError  %+v", 500, o.Payload)
}
func (o *UserSetUsingOrgInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UserSetUsingOrgInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
