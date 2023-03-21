// Code generated by go-swagger; DO NOT EDIT.

package orgs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetOrgQuotaReader is a Reader for the GetOrgQuota structure.
type GetOrgQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrgQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgQuotaOK creates a GetOrgQuotaOK with default headers values
func NewGetOrgQuotaOK() *GetOrgQuotaOK {
	return &GetOrgQuotaOK{}
}

/*
GetOrgQuotaOK describes a response with status code 200, with default header values.

(empty)
*/
type GetOrgQuotaOK struct {
	Payload []*models.QuotaDTO
}

func (o *GetOrgQuotaOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaOK  %+v", 200, o.Payload)
}
func (o *GetOrgQuotaOK) GetPayload() []*models.QuotaDTO {
	return o.Payload
}

func (o *GetOrgQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgQuotaUnauthorized creates a GetOrgQuotaUnauthorized with default headers values
func NewGetOrgQuotaUnauthorized() *GetOrgQuotaUnauthorized {
	return &GetOrgQuotaUnauthorized{}
}

/*
GetOrgQuotaUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetOrgQuotaUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetOrgQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaUnauthorized  %+v", 401, o.Payload)
}
func (o *GetOrgQuotaUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgQuotaForbidden creates a GetOrgQuotaForbidden with default headers values
func NewGetOrgQuotaForbidden() *GetOrgQuotaForbidden {
	return &GetOrgQuotaForbidden{}
}

/*
GetOrgQuotaForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetOrgQuotaForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetOrgQuotaForbidden) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaForbidden  %+v", 403, o.Payload)
}
func (o *GetOrgQuotaForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgQuotaNotFound creates a GetOrgQuotaNotFound with default headers values
func NewGetOrgQuotaNotFound() *GetOrgQuotaNotFound {
	return &GetOrgQuotaNotFound{}
}

/*
GetOrgQuotaNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetOrgQuotaNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetOrgQuotaNotFound) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaNotFound  %+v", 404, o.Payload)
}
func (o *GetOrgQuotaNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgQuotaInternalServerError creates a GetOrgQuotaInternalServerError with default headers values
func NewGetOrgQuotaInternalServerError() *GetOrgQuotaInternalServerError {
	return &GetOrgQuotaInternalServerError{}
}

/*
GetOrgQuotaInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetOrgQuotaInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetOrgQuotaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaInternalServerError  %+v", 500, o.Payload)
}
func (o *GetOrgQuotaInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
