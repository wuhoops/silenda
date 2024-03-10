// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/wuhoops/silenda/cli/models"
)

// UpdateSecretReader is a Reader for the UpdateSecret structure.
type UpdateSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /secret.updateSecret] updateSecret", response, response.Code())
	}
}

// NewUpdateSecretOK creates a UpdateSecretOK with default headers values
func NewUpdateSecretOK() *UpdateSecretOK {
	return &UpdateSecretOK{}
}

/*
UpdateSecretOK describes a response with status code 200, with default header values.

OK
*/
type UpdateSecretOK struct {
	Payload *models.HandlersResponseString
}

// IsSuccess returns true when this update secret o k response has a 2xx status code
func (o *UpdateSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update secret o k response has a 3xx status code
func (o *UpdateSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update secret o k response has a 4xx status code
func (o *UpdateSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update secret o k response has a 5xx status code
func (o *UpdateSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update secret o k response a status code equal to that given
func (o *UpdateSecretOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update secret o k response
func (o *UpdateSecretOK) Code() int {
	return 200
}

func (o *UpdateSecretOK) Error() string {
	return fmt.Sprintf("[POST /secret.updateSecret][%d] updateSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateSecretOK) String() string {
	return fmt.Sprintf("[POST /secret.updateSecret][%d] updateSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateSecretOK) GetPayload() *models.HandlersResponseString {
	return o.Payload
}

func (o *UpdateSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersResponseString)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecretBadRequest creates a UpdateSecretBadRequest with default headers values
func NewUpdateSecretBadRequest() *UpdateSecretBadRequest {
	return &UpdateSecretBadRequest{}
}

/*
UpdateSecretBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateSecretBadRequest struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this update secret bad request response has a 2xx status code
func (o *UpdateSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update secret bad request response has a 3xx status code
func (o *UpdateSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update secret bad request response has a 4xx status code
func (o *UpdateSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update secret bad request response has a 5xx status code
func (o *UpdateSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update secret bad request response a status code equal to that given
func (o *UpdateSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update secret bad request response
func (o *UpdateSecretBadRequest) Code() int {
	return 400
}

func (o *UpdateSecretBadRequest) Error() string {
	return fmt.Sprintf("[POST /secret.updateSecret][%d] updateSecretBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSecretBadRequest) String() string {
	return fmt.Sprintf("[POST /secret.updateSecret][%d] updateSecretBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSecretBadRequest) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *UpdateSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecretInternalServerError creates a UpdateSecretInternalServerError with default headers values
func NewUpdateSecretInternalServerError() *UpdateSecretInternalServerError {
	return &UpdateSecretInternalServerError{}
}

/*
UpdateSecretInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateSecretInternalServerError struct {
	Payload *models.HandlersErrResponse
}

// IsSuccess returns true when this update secret internal server error response has a 2xx status code
func (o *UpdateSecretInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update secret internal server error response has a 3xx status code
func (o *UpdateSecretInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update secret internal server error response has a 4xx status code
func (o *UpdateSecretInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update secret internal server error response has a 5xx status code
func (o *UpdateSecretInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update secret internal server error response a status code equal to that given
func (o *UpdateSecretInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update secret internal server error response
func (o *UpdateSecretInternalServerError) Code() int {
	return 500
}

func (o *UpdateSecretInternalServerError) Error() string {
	return fmt.Sprintf("[POST /secret.updateSecret][%d] updateSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSecretInternalServerError) String() string {
	return fmt.Sprintf("[POST /secret.updateSecret][%d] updateSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSecretInternalServerError) GetPayload() *models.HandlersErrResponse {
	return o.Payload
}

func (o *UpdateSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HandlersErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
