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

// GetFoldersReader is a Reader for the GetFolders structure.
type GetFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFoldersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetFoldersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFoldersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFoldersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFoldersOK creates a GetFoldersOK with default headers values
func NewGetFoldersOK() *GetFoldersOK {
	return &GetFoldersOK{}
}

/*
GetFoldersOK describes a response with status code 200, with default header values.

(empty)
*/
type GetFoldersOK struct {
	Payload []*models.FolderSearchHit
}

func (o *GetFoldersOK) Error() string {
	return fmt.Sprintf("[GET /folders][%d] getFoldersOK  %+v", 200, o.Payload)
}
func (o *GetFoldersOK) GetPayload() []*models.FolderSearchHit {
	return o.Payload
}

func (o *GetFoldersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersUnauthorized creates a GetFoldersUnauthorized with default headers values
func NewGetFoldersUnauthorized() *GetFoldersUnauthorized {
	return &GetFoldersUnauthorized{}
}

/*
GetFoldersUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetFoldersUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetFoldersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /folders][%d] getFoldersUnauthorized  %+v", 401, o.Payload)
}
func (o *GetFoldersUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFoldersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersForbidden creates a GetFoldersForbidden with default headers values
func NewGetFoldersForbidden() *GetFoldersForbidden {
	return &GetFoldersForbidden{}
}

/*
GetFoldersForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetFoldersForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetFoldersForbidden) Error() string {
	return fmt.Sprintf("[GET /folders][%d] getFoldersForbidden  %+v", 403, o.Payload)
}
func (o *GetFoldersForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFoldersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFoldersInternalServerError creates a GetFoldersInternalServerError with default headers values
func NewGetFoldersInternalServerError() *GetFoldersInternalServerError {
	return &GetFoldersInternalServerError{}
}

/*
GetFoldersInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetFoldersInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetFoldersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /folders][%d] getFoldersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFoldersInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetFoldersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
