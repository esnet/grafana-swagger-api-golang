// Code generated by go-swagger; DO NOT EDIT.

package signed_in_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// GetSignedInUserTeamListReader is a Reader for the GetSignedInUserTeamList structure.
type GetSignedInUserTeamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSignedInUserTeamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSignedInUserTeamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSignedInUserTeamListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSignedInUserTeamListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSignedInUserTeamListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSignedInUserTeamListOK creates a GetSignedInUserTeamListOK with default headers values
func NewGetSignedInUserTeamListOK() *GetSignedInUserTeamListOK {
	return &GetSignedInUserTeamListOK{}
}

/*
GetSignedInUserTeamListOK describes a response with status code 200, with default header values.

(empty)
*/
type GetSignedInUserTeamListOK struct {
	Payload []*models.TeamDTO
}

func (o *GetSignedInUserTeamListOK) Error() string {
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListOK  %+v", 200, o.Payload)
}
func (o *GetSignedInUserTeamListOK) GetPayload() []*models.TeamDTO {
	return o.Payload
}

func (o *GetSignedInUserTeamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserTeamListUnauthorized creates a GetSignedInUserTeamListUnauthorized with default headers values
func NewGetSignedInUserTeamListUnauthorized() *GetSignedInUserTeamListUnauthorized {
	return &GetSignedInUserTeamListUnauthorized{}
}

/*
GetSignedInUserTeamListUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetSignedInUserTeamListUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSignedInUserTeamListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSignedInUserTeamListUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserTeamListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserTeamListForbidden creates a GetSignedInUserTeamListForbidden with default headers values
func NewGetSignedInUserTeamListForbidden() *GetSignedInUserTeamListForbidden {
	return &GetSignedInUserTeamListForbidden{}
}

/*
GetSignedInUserTeamListForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetSignedInUserTeamListForbidden struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSignedInUserTeamListForbidden) Error() string {
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListForbidden  %+v", 403, o.Payload)
}
func (o *GetSignedInUserTeamListForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserTeamListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSignedInUserTeamListInternalServerError creates a GetSignedInUserTeamListInternalServerError with default headers values
func NewGetSignedInUserTeamListInternalServerError() *GetSignedInUserTeamListInternalServerError {
	return &GetSignedInUserTeamListInternalServerError{}
}

/*
GetSignedInUserTeamListInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetSignedInUserTeamListInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *GetSignedInUserTeamListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user/teams][%d] getSignedInUserTeamListInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSignedInUserTeamListInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetSignedInUserTeamListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
