// Code generated by go-swagger; DO NOT EDIT.

package query_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// MigrateQueriesReader is a Reader for the MigrateQueries structure.
type MigrateQueriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateQueriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateQueriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMigrateQueriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMigrateQueriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMigrateQueriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMigrateQueriesOK creates a MigrateQueriesOK with default headers values
func NewMigrateQueriesOK() *MigrateQueriesOK {
	return &MigrateQueriesOK{}
}

/*
MigrateQueriesOK describes a response with status code 200, with default header values.

(empty)
*/
type MigrateQueriesOK struct {
	Payload *models.QueryHistoryMigrationResponse
}

func (o *MigrateQueriesOK) Error() string {
	return fmt.Sprintf("[POST /query-history/migrate][%d] migrateQueriesOK  %+v", 200, o.Payload)
}
func (o *MigrateQueriesOK) GetPayload() *models.QueryHistoryMigrationResponse {
	return o.Payload
}

func (o *MigrateQueriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryHistoryMigrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateQueriesBadRequest creates a MigrateQueriesBadRequest with default headers values
func NewMigrateQueriesBadRequest() *MigrateQueriesBadRequest {
	return &MigrateQueriesBadRequest{}
}

/*
MigrateQueriesBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type MigrateQueriesBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *MigrateQueriesBadRequest) Error() string {
	return fmt.Sprintf("[POST /query-history/migrate][%d] migrateQueriesBadRequest  %+v", 400, o.Payload)
}
func (o *MigrateQueriesBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MigrateQueriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateQueriesUnauthorized creates a MigrateQueriesUnauthorized with default headers values
func NewMigrateQueriesUnauthorized() *MigrateQueriesUnauthorized {
	return &MigrateQueriesUnauthorized{}
}

/*
MigrateQueriesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type MigrateQueriesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *MigrateQueriesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /query-history/migrate][%d] migrateQueriesUnauthorized  %+v", 401, o.Payload)
}
func (o *MigrateQueriesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MigrateQueriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateQueriesInternalServerError creates a MigrateQueriesInternalServerError with default headers values
func NewMigrateQueriesInternalServerError() *MigrateQueriesInternalServerError {
	return &MigrateQueriesInternalServerError{}
}

/*
MigrateQueriesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type MigrateQueriesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *MigrateQueriesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /query-history/migrate][%d] migrateQueriesInternalServerError  %+v", 500, o.Payload)
}
func (o *MigrateQueriesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MigrateQueriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
