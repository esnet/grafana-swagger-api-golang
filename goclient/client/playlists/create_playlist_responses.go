// Code generated by go-swagger; DO NOT EDIT.

package playlists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// CreatePlaylistReader is a Reader for the CreatePlaylist structure.
type CreatePlaylistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePlaylistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePlaylistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreatePlaylistUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePlaylistForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreatePlaylistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePlaylistInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePlaylistOK creates a CreatePlaylistOK with default headers values
func NewCreatePlaylistOK() *CreatePlaylistOK {
	return &CreatePlaylistOK{}
}

/*
CreatePlaylistOK describes a response with status code 200, with default header values.

(empty)
*/
type CreatePlaylistOK struct {
	Payload *models.Playlist
}

func (o *CreatePlaylistOK) Error() string {
	return fmt.Sprintf("[POST /playlists][%d] createPlaylistOK  %+v", 200, o.Payload)
}
func (o *CreatePlaylistOK) GetPayload() *models.Playlist {
	return o.Payload
}

func (o *CreatePlaylistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Playlist)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePlaylistUnauthorized creates a CreatePlaylistUnauthorized with default headers values
func NewCreatePlaylistUnauthorized() *CreatePlaylistUnauthorized {
	return &CreatePlaylistUnauthorized{}
}

/*
CreatePlaylistUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreatePlaylistUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *CreatePlaylistUnauthorized) Error() string {
	return fmt.Sprintf("[POST /playlists][%d] createPlaylistUnauthorized  %+v", 401, o.Payload)
}
func (o *CreatePlaylistUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreatePlaylistUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePlaylistForbidden creates a CreatePlaylistForbidden with default headers values
func NewCreatePlaylistForbidden() *CreatePlaylistForbidden {
	return &CreatePlaylistForbidden{}
}

/*
CreatePlaylistForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreatePlaylistForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *CreatePlaylistForbidden) Error() string {
	return fmt.Sprintf("[POST /playlists][%d] createPlaylistForbidden  %+v", 403, o.Payload)
}
func (o *CreatePlaylistForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreatePlaylistForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePlaylistNotFound creates a CreatePlaylistNotFound with default headers values
func NewCreatePlaylistNotFound() *CreatePlaylistNotFound {
	return &CreatePlaylistNotFound{}
}

/*
CreatePlaylistNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type CreatePlaylistNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *CreatePlaylistNotFound) Error() string {
	return fmt.Sprintf("[POST /playlists][%d] createPlaylistNotFound  %+v", 404, o.Payload)
}
func (o *CreatePlaylistNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreatePlaylistNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePlaylistInternalServerError creates a CreatePlaylistInternalServerError with default headers values
func NewCreatePlaylistInternalServerError() *CreatePlaylistInternalServerError {
	return &CreatePlaylistInternalServerError{}
}

/*
CreatePlaylistInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreatePlaylistInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *CreatePlaylistInternalServerError) Error() string {
	return fmt.Sprintf("[POST /playlists][%d] createPlaylistInternalServerError  %+v", 500, o.Payload)
}
func (o *CreatePlaylistInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreatePlaylistInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}