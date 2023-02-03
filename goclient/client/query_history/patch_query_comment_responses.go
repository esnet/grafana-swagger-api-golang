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

// PatchQueryCommentReader is a Reader for the PatchQueryComment structure.
type PatchQueryCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchQueryCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchQueryCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchQueryCommentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchQueryCommentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchQueryCommentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchQueryCommentOK creates a PatchQueryCommentOK with default headers values
func NewPatchQueryCommentOK() *PatchQueryCommentOK {
	return &PatchQueryCommentOK{}
}

/*
PatchQueryCommentOK describes a response with status code 200, with default header values.

(empty)
*/
type PatchQueryCommentOK struct {
	Payload *models.QueryHistoryResponse
}

func (o *PatchQueryCommentOK) Error() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentOK  %+v", 200, o.Payload)
}
func (o *PatchQueryCommentOK) GetPayload() *models.QueryHistoryResponse {
	return o.Payload
}

func (o *PatchQueryCommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueryCommentBadRequest creates a PatchQueryCommentBadRequest with default headers values
func NewPatchQueryCommentBadRequest() *PatchQueryCommentBadRequest {
	return &PatchQueryCommentBadRequest{}
}

/*
PatchQueryCommentBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type PatchQueryCommentBadRequest struct {
	Payload *models.ErrorResponseBody
}

func (o *PatchQueryCommentBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentBadRequest  %+v", 400, o.Payload)
}
func (o *PatchQueryCommentBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchQueryCommentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueryCommentUnauthorized creates a PatchQueryCommentUnauthorized with default headers values
func NewPatchQueryCommentUnauthorized() *PatchQueryCommentUnauthorized {
	return &PatchQueryCommentUnauthorized{}
}

/*
PatchQueryCommentUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type PatchQueryCommentUnauthorized struct {
	Payload *models.ErrorResponseBody
}

func (o *PatchQueryCommentUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentUnauthorized  %+v", 401, o.Payload)
}
func (o *PatchQueryCommentUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchQueryCommentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueryCommentInternalServerError creates a PatchQueryCommentInternalServerError with default headers values
func NewPatchQueryCommentInternalServerError() *PatchQueryCommentInternalServerError {
	return &PatchQueryCommentInternalServerError{}
}

/*
PatchQueryCommentInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PatchQueryCommentInternalServerError struct {
	Payload *models.ErrorResponseBody
}

func (o *PatchQueryCommentInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /query-history/{query_history_uid}][%d] patchQueryCommentInternalServerError  %+v", 500, o.Payload)
}
func (o *PatchQueryCommentInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PatchQueryCommentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
