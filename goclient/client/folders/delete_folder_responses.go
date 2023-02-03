// Code generated by go-swagger; DO NOT EDIT.

package folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// DeleteFolderReader is a Reader for the DeleteFolder structure.
type DeleteFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFolderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFolderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFolderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFolderOK creates a DeleteFolderOK with default headers values
func NewDeleteFolderOK() *DeleteFolderOK {
	return &DeleteFolderOK{}
}

/*
DeleteFolderOK describes a response with status code 200, with default header values.

(empty)
*/
type DeleteFolderOK struct {
	Payload *models.DeleteFolderOKBody
}

func (o *DeleteFolderOK) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_uid}][%d] deleteFolderOK  %+v", 200, o.Payload)
}
func (o *DeleteFolderOK) GetPayload() *models.DeleteFolderOKBody {
	return o.Payload
}

func (o *DeleteFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteFolderOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFolderBadRequest creates a DeleteFolderBadRequest with default headers values
func NewDeleteFolderBadRequest() *DeleteFolderBadRequest {
	return &DeleteFolderBadRequest{}
}

/*
DeleteFolderBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DeleteFolderBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteFolderBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_uid}][%d] deleteFolderBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteFolderBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFolderUnauthorized creates a DeleteFolderUnauthorized with default headers values
func NewDeleteFolderUnauthorized() *DeleteFolderUnauthorized {
	return &DeleteFolderUnauthorized{}
}

/*
DeleteFolderUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteFolderUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteFolderUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_uid}][%d] deleteFolderUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteFolderUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteFolderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFolderForbidden creates a DeleteFolderForbidden with default headers values
func NewDeleteFolderForbidden() *DeleteFolderForbidden {
	return &DeleteFolderForbidden{}
}

/*
DeleteFolderForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteFolderForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteFolderForbidden) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_uid}][%d] deleteFolderForbidden  %+v", 403, o.Payload)
}
func (o *DeleteFolderForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteFolderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFolderNotFound creates a DeleteFolderNotFound with default headers values
func NewDeleteFolderNotFound() *DeleteFolderNotFound {
	return &DeleteFolderNotFound{}
}

/*
DeleteFolderNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DeleteFolderNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteFolderNotFound) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_uid}][%d] deleteFolderNotFound  %+v", 404, o.Payload)
}
func (o *DeleteFolderNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFolderInternalServerError creates a DeleteFolderInternalServerError with default headers values
func NewDeleteFolderInternalServerError() *DeleteFolderInternalServerError {
	return &DeleteFolderInternalServerError{}
}

/*
DeleteFolderInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteFolderInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *DeleteFolderInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /folders/{folder_uid}][%d] deleteFolderInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteFolderInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteFolderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
