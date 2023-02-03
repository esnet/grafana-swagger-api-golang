// Code generated by go-swagger; DO NOT EDIT.

package library_elements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetLibraryElementByNameReader is a Reader for the GetLibraryElementByName structure.
type GetLibraryElementByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLibraryElementByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLibraryElementByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLibraryElementByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLibraryElementByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLibraryElementByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLibraryElementByNameOK creates a GetLibraryElementByNameOK with default headers values
func NewGetLibraryElementByNameOK() *GetLibraryElementByNameOK {
	return &GetLibraryElementByNameOK{}
}

/*
GetLibraryElementByNameOK describes a response with status code 200, with default header values.

(empty)
*/
type GetLibraryElementByNameOK struct {
	Payload *models.LibraryElementResponse
}

func (o *GetLibraryElementByNameOK) Error() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameOK  %+v", 200, o.Payload)
}
func (o *GetLibraryElementByNameOK) GetPayload() *models.LibraryElementResponse {
	return o.Payload
}

func (o *GetLibraryElementByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LibraryElementResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementByNameUnauthorized creates a GetLibraryElementByNameUnauthorized with default headers values
func NewGetLibraryElementByNameUnauthorized() *GetLibraryElementByNameUnauthorized {
	return &GetLibraryElementByNameUnauthorized{}
}

/*
GetLibraryElementByNameUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetLibraryElementByNameUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetLibraryElementByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *GetLibraryElementByNameUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementByNameNotFound creates a GetLibraryElementByNameNotFound with default headers values
func NewGetLibraryElementByNameNotFound() *GetLibraryElementByNameNotFound {
	return &GetLibraryElementByNameNotFound{}
}

/*
GetLibraryElementByNameNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetLibraryElementByNameNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *GetLibraryElementByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameNotFound  %+v", 404, o.Payload)
}
func (o *GetLibraryElementByNameNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLibraryElementByNameInternalServerError creates a GetLibraryElementByNameInternalServerError with default headers values
func NewGetLibraryElementByNameInternalServerError() *GetLibraryElementByNameInternalServerError {
	return &GetLibraryElementByNameInternalServerError{}
}

/*
GetLibraryElementByNameInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetLibraryElementByNameInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetLibraryElementByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /library-elements/name/{library_element_name}][%d] getLibraryElementByNameInternalServerError  %+v", 500, o.Payload)
}
func (o *GetLibraryElementByNameInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetLibraryElementByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}