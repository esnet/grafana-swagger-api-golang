// Code generated by go-swagger; DO NOT EDIT.

package dashboard_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetDashboardVersionsByUIDReader is a Reader for the GetDashboardVersionsByUID structure.
type GetDashboardVersionsByUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardVersionsByUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardVersionsByUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDashboardVersionsByUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDashboardVersionsByUIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDashboardVersionsByUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDashboardVersionsByUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDashboardVersionsByUIDOK creates a GetDashboardVersionsByUIDOK with default headers values
func NewGetDashboardVersionsByUIDOK() *GetDashboardVersionsByUIDOK {
	return &GetDashboardVersionsByUIDOK{}
}

/*
GetDashboardVersionsByUIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetDashboardVersionsByUIDOK struct {
	Payload []*models.DashboardVersionMeta
}

func (o *GetDashboardVersionsByUIDOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}/versions][%d] getDashboardVersionsByUidOK  %+v", 200, o.Payload)
}
func (o *GetDashboardVersionsByUIDOK) GetPayload() []*models.DashboardVersionMeta {
	return o.Payload
}

func (o *GetDashboardVersionsByUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionsByUIDUnauthorized creates a GetDashboardVersionsByUIDUnauthorized with default headers values
func NewGetDashboardVersionsByUIDUnauthorized() *GetDashboardVersionsByUIDUnauthorized {
	return &GetDashboardVersionsByUIDUnauthorized{}
}

/*
GetDashboardVersionsByUIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetDashboardVersionsByUIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetDashboardVersionsByUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}/versions][%d] getDashboardVersionsByUidUnauthorized  %+v", 401, o.Payload)
}
func (o *GetDashboardVersionsByUIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionsByUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionsByUIDForbidden creates a GetDashboardVersionsByUIDForbidden with default headers values
func NewGetDashboardVersionsByUIDForbidden() *GetDashboardVersionsByUIDForbidden {
	return &GetDashboardVersionsByUIDForbidden{}
}

/*
GetDashboardVersionsByUIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetDashboardVersionsByUIDForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetDashboardVersionsByUIDForbidden) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}/versions][%d] getDashboardVersionsByUidForbidden  %+v", 403, o.Payload)
}
func (o *GetDashboardVersionsByUIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionsByUIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionsByUIDNotFound creates a GetDashboardVersionsByUIDNotFound with default headers values
func NewGetDashboardVersionsByUIDNotFound() *GetDashboardVersionsByUIDNotFound {
	return &GetDashboardVersionsByUIDNotFound{}
}

/*
GetDashboardVersionsByUIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetDashboardVersionsByUIDNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetDashboardVersionsByUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}/versions][%d] getDashboardVersionsByUidNotFound  %+v", 404, o.Payload)
}
func (o *GetDashboardVersionsByUIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionsByUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardVersionsByUIDInternalServerError creates a GetDashboardVersionsByUIDInternalServerError with default headers values
func NewGetDashboardVersionsByUIDInternalServerError() *GetDashboardVersionsByUIDInternalServerError {
	return &GetDashboardVersionsByUIDInternalServerError{}
}

/*
GetDashboardVersionsByUIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetDashboardVersionsByUIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetDashboardVersionsByUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dashboards/uid/{uid}/versions][%d] getDashboardVersionsByUidInternalServerError  %+v", 500, o.Payload)
}
func (o *GetDashboardVersionsByUIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetDashboardVersionsByUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
