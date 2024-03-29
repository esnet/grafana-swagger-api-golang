// Code generated by go-swagger; DO NOT EDIT.

package get_current_org

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetCurrentOrgQuotaReader is a Reader for the GetCurrentOrgQuota structure.
type GetCurrentOrgQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentOrgQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentOrgQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentOrgQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCurrentOrgQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCurrentOrgQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentOrgQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentOrgQuotaOK creates a GetCurrentOrgQuotaOK with default headers values
func NewGetCurrentOrgQuotaOK() *GetCurrentOrgQuotaOK {
	return &GetCurrentOrgQuotaOK{}
}

/*
GetCurrentOrgQuotaOK describes a response with status code 200, with default header values.

(empty)
*/
type GetCurrentOrgQuotaOK struct {
	Payload []*models.QuotaDTO
}

func (o *GetCurrentOrgQuotaOK) Error() string {
	return fmt.Sprintf("[GET /org/quotas][%d] getCurrentOrgQuotaOK  %+v", 200, o.Payload)
}
func (o *GetCurrentOrgQuotaOK) GetPayload() []*models.QuotaDTO {
	return o.Payload
}

func (o *GetCurrentOrgQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgQuotaUnauthorized creates a GetCurrentOrgQuotaUnauthorized with default headers values
func NewGetCurrentOrgQuotaUnauthorized() *GetCurrentOrgQuotaUnauthorized {
	return &GetCurrentOrgQuotaUnauthorized{}
}

/*
GetCurrentOrgQuotaUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetCurrentOrgQuotaUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetCurrentOrgQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /org/quotas][%d] getCurrentOrgQuotaUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCurrentOrgQuotaUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgQuotaForbidden creates a GetCurrentOrgQuotaForbidden with default headers values
func NewGetCurrentOrgQuotaForbidden() *GetCurrentOrgQuotaForbidden {
	return &GetCurrentOrgQuotaForbidden{}
}

/*
GetCurrentOrgQuotaForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetCurrentOrgQuotaForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetCurrentOrgQuotaForbidden) Error() string {
	return fmt.Sprintf("[GET /org/quotas][%d] getCurrentOrgQuotaForbidden  %+v", 403, o.Payload)
}
func (o *GetCurrentOrgQuotaForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgQuotaNotFound creates a GetCurrentOrgQuotaNotFound with default headers values
func NewGetCurrentOrgQuotaNotFound() *GetCurrentOrgQuotaNotFound {
	return &GetCurrentOrgQuotaNotFound{}
}

/*
GetCurrentOrgQuotaNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetCurrentOrgQuotaNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetCurrentOrgQuotaNotFound) Error() string {
	return fmt.Sprintf("[GET /org/quotas][%d] getCurrentOrgQuotaNotFound  %+v", 404, o.Payload)
}
func (o *GetCurrentOrgQuotaNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentOrgQuotaInternalServerError creates a GetCurrentOrgQuotaInternalServerError with default headers values
func NewGetCurrentOrgQuotaInternalServerError() *GetCurrentOrgQuotaInternalServerError {
	return &GetCurrentOrgQuotaInternalServerError{}
}

/*
GetCurrentOrgQuotaInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCurrentOrgQuotaInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetCurrentOrgQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /org/quotas][%d] getCurrentOrgQuotaInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCurrentOrgQuotaInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCurrentOrgQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
