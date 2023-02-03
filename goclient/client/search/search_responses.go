// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// SearchReader is a Reader for the Search structure.
type SearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSearchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchOK creates a SearchOK with default headers values
func NewSearchOK() *SearchOK {
	return &SearchOK{}
}

/*
SearchOK describes a response with status code 200, with default header values.

(empty)
*/
type SearchOK struct {
	Payload models.HitList
}

func (o *SearchOK) Error() string {
	return fmt.Sprintf("[GET /search][%d] searchOK  %+v", 200, o.Payload)
}
func (o *SearchOK) GetPayload() models.HitList {
	return o.Payload
}

func (o *SearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUnauthorized creates a SearchUnauthorized with default headers values
func NewSearchUnauthorized() *SearchUnauthorized {
	return &SearchUnauthorized{}
}

/*
SearchUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type SearchUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /search][%d] searchUnauthorized  %+v", 401, o.Payload)
}
func (o *SearchUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUnprocessableEntity creates a SearchUnprocessableEntity with default headers values
func NewSearchUnprocessableEntity() *SearchUnprocessableEntity {
	return &SearchUnprocessableEntity{}
}

/*
SearchUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntityError
*/
type SearchUnprocessableEntity struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /search][%d] searchUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *SearchUnprocessableEntity) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchInternalServerError creates a SearchInternalServerError with default headers values
func NewSearchInternalServerError() *SearchInternalServerError {
	return &SearchInternalServerError{}
}

/*
SearchInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SearchInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *SearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /search][%d] searchInternalServerError  %+v", 500, o.Payload)
}
func (o *SearchInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
