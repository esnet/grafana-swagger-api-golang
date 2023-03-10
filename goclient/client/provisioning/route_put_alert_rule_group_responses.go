// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
)

// RoutePutAlertRuleGroupReader is a Reader for the RoutePutAlertRuleGroup structure.
type RoutePutAlertRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RoutePutAlertRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRoutePutAlertRuleGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRoutePutAlertRuleGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRoutePutAlertRuleGroupOK creates a RoutePutAlertRuleGroupOK with default headers values
func NewRoutePutAlertRuleGroupOK() *RoutePutAlertRuleGroupOK {
	return &RoutePutAlertRuleGroupOK{}
}

/*
RoutePutAlertRuleGroupOK describes a response with status code 200, with default header values.

AlertRuleGroup
*/
type RoutePutAlertRuleGroupOK struct {
	Payload *models.AlertRuleGroup
}

func (o *RoutePutAlertRuleGroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] routePutAlertRuleGroupOK  %+v", 200, o.Payload)
}
func (o *RoutePutAlertRuleGroupOK) GetPayload() *models.AlertRuleGroup {
	return o.Payload
}

func (o *RoutePutAlertRuleGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRuleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRoutePutAlertRuleGroupBadRequest creates a RoutePutAlertRuleGroupBadRequest with default headers values
func NewRoutePutAlertRuleGroupBadRequest() *RoutePutAlertRuleGroupBadRequest {
	return &RoutePutAlertRuleGroupBadRequest{}
}

/*
RoutePutAlertRuleGroupBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type RoutePutAlertRuleGroupBadRequest struct {
	Payload *models.ValidationError
}

func (o *RoutePutAlertRuleGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] routePutAlertRuleGroupBadRequest  %+v", 400, o.Payload)
}
func (o *RoutePutAlertRuleGroupBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RoutePutAlertRuleGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
