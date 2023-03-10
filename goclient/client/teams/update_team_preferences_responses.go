// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// UpdateTeamPreferencesReader is a Reader for the UpdateTeamPreferences structure.
type UpdateTeamPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTeamPreferencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTeamPreferencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTeamPreferencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTeamPreferencesOK creates a UpdateTeamPreferencesOK with default headers values
func NewUpdateTeamPreferencesOK() *UpdateTeamPreferencesOK {
	return &UpdateTeamPreferencesOK{}
}

/*
UpdateTeamPreferencesOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type UpdateTeamPreferencesOK struct {
	Payload *models.SuccessResponseBody
}

func (o *UpdateTeamPreferencesOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/preferences][%d] updateTeamPreferencesOK  %+v", 200, o.Payload)
}
func (o *UpdateTeamPreferencesOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *UpdateTeamPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamPreferencesBadRequest creates a UpdateTeamPreferencesBadRequest with default headers values
func NewUpdateTeamPreferencesBadRequest() *UpdateTeamPreferencesBadRequest {
	return &UpdateTeamPreferencesBadRequest{}
}

/*
UpdateTeamPreferencesBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateTeamPreferencesBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateTeamPreferencesBadRequest) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/preferences][%d] updateTeamPreferencesBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateTeamPreferencesBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamPreferencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamPreferencesUnauthorized creates a UpdateTeamPreferencesUnauthorized with default headers values
func NewUpdateTeamPreferencesUnauthorized() *UpdateTeamPreferencesUnauthorized {
	return &UpdateTeamPreferencesUnauthorized{}
}

/*
UpdateTeamPreferencesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateTeamPreferencesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateTeamPreferencesUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/preferences][%d] updateTeamPreferencesUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateTeamPreferencesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamPreferencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamPreferencesInternalServerError creates a UpdateTeamPreferencesInternalServerError with default headers values
func NewUpdateTeamPreferencesInternalServerError() *UpdateTeamPreferencesInternalServerError {
	return &UpdateTeamPreferencesInternalServerError{}
}

/*
UpdateTeamPreferencesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateTeamPreferencesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateTeamPreferencesInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}/preferences][%d] updateTeamPreferencesInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateTeamPreferencesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateTeamPreferencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
