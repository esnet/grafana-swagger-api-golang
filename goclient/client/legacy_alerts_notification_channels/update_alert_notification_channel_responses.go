// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts_notification_channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// UpdateAlertNotificationChannelReader is a Reader for the UpdateAlertNotificationChannel structure.
type UpdateAlertNotificationChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAlertNotificationChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAlertNotificationChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAlertNotificationChannelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAlertNotificationChannelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAlertNotificationChannelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAlertNotificationChannelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAlertNotificationChannelOK creates a UpdateAlertNotificationChannelOK with default headers values
func NewUpdateAlertNotificationChannelOK() *UpdateAlertNotificationChannelOK {
	return &UpdateAlertNotificationChannelOK{}
}

/*
UpdateAlertNotificationChannelOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateAlertNotificationChannelOK struct {
	Payload *models.AlertNotification
}

func (o *UpdateAlertNotificationChannelOK) Error() string {
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelOK  %+v", 200, o.Payload)
}
func (o *UpdateAlertNotificationChannelOK) GetPayload() *models.AlertNotification {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertNotificationChannelUnauthorized creates a UpdateAlertNotificationChannelUnauthorized with default headers values
func NewUpdateAlertNotificationChannelUnauthorized() *UpdateAlertNotificationChannelUnauthorized {
	return &UpdateAlertNotificationChannelUnauthorized{}
}

/*
UpdateAlertNotificationChannelUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateAlertNotificationChannelUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateAlertNotificationChannelUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateAlertNotificationChannelUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertNotificationChannelForbidden creates a UpdateAlertNotificationChannelForbidden with default headers values
func NewUpdateAlertNotificationChannelForbidden() *UpdateAlertNotificationChannelForbidden {
	return &UpdateAlertNotificationChannelForbidden{}
}

/*
UpdateAlertNotificationChannelForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateAlertNotificationChannelForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateAlertNotificationChannelForbidden) Error() string {
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelForbidden  %+v", 403, o.Payload)
}
func (o *UpdateAlertNotificationChannelForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertNotificationChannelNotFound creates a UpdateAlertNotificationChannelNotFound with default headers values
func NewUpdateAlertNotificationChannelNotFound() *UpdateAlertNotificationChannelNotFound {
	return &UpdateAlertNotificationChannelNotFound{}
}

/*
UpdateAlertNotificationChannelNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateAlertNotificationChannelNotFound struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateAlertNotificationChannelNotFound) Error() string {
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelNotFound  %+v", 404, o.Payload)
}
func (o *UpdateAlertNotificationChannelNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertNotificationChannelInternalServerError creates a UpdateAlertNotificationChannelInternalServerError with default headers values
func NewUpdateAlertNotificationChannelInternalServerError() *UpdateAlertNotificationChannelInternalServerError {
	return &UpdateAlertNotificationChannelInternalServerError{}
}

/*
UpdateAlertNotificationChannelInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateAlertNotificationChannelInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *UpdateAlertNotificationChannelInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateAlertNotificationChannelInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
