// Code generated by go-swagger; DO NOT EDIT.

package access_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/oNaiPs/fyde-cli/models"
)

// EditResourceReader is a Reader for the EditResource structure.
type EditResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEditResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEditResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditResourceOK creates a EditResourceOK with default headers values
func NewEditResourceOK() *EditResourceOK {
	return &EditResourceOK{}
}

/*EditResourceOK handles this case with default header values.

Resource edited
*/
type EditResourceOK struct {
	Payload *EditResourceOKBody
}

func (o *EditResourceOK) Error() string {
	return fmt.Sprintf("[PATCH /access_resources/{id}][%d] editResourceOK  %+v", 200, o.Payload)
}

func (o *EditResourceOK) GetPayload() *EditResourceOKBody {
	return o.Payload
}

func (o *EditResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(EditResourceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditResourceUnauthorized creates a EditResourceUnauthorized with default headers values
func NewEditResourceUnauthorized() *EditResourceUnauthorized {
	return &EditResourceUnauthorized{}
}

/*EditResourceUnauthorized handles this case with default header values.

unauthorized: invalid credentials or missing authentication headers
*/
type EditResourceUnauthorized struct {
	Payload *models.GenericUnauthorizedResponse
}

func (o *EditResourceUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /access_resources/{id}][%d] editResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *EditResourceUnauthorized) GetPayload() *models.GenericUnauthorizedResponse {
	return o.Payload
}

func (o *EditResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericUnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditResourceNotFound creates a EditResourceNotFound with default headers values
func NewEditResourceNotFound() *EditResourceNotFound {
	return &EditResourceNotFound{}
}

/*EditResourceNotFound handles this case with default header values.

resource not found
*/
type EditResourceNotFound struct {
}

func (o *EditResourceNotFound) Error() string {
	return fmt.Sprintf("[PATCH /access_resources/{id}][%d] editResourceNotFound ", 404)
}

func (o *EditResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*AccessResourceAO1AccessCount access resource a o1 access count
swagger:model AccessResourceAO1AccessCount
*/
type AccessResourceAO1AccessCount struct {

	// denied
	Denied int64 `json:"denied,omitempty"`

	// granted
	Granted int64 `json:"granted,omitempty"`
}

// Validate validates this access resource a o1 access count
func (o *AccessResourceAO1AccessCount) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AccessResourceAO1AccessCount) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AccessResourceAO1AccessCount) UnmarshalBinary(b []byte) error {
	var res AccessResourceAO1AccessCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*EditResourceBody edit resource body
swagger:model EditResourceBody
*/
type EditResourceBody struct {

	// access resource
	AccessResource struct {
		models.AccessResource

		// access count
		AccessCount *AccessResourceAO1AccessCount `json:"access_count,omitempty"`

		// access proxy name
		AccessProxyName string `json:"access_proxy_name,omitempty"`

		// last access at
		// Format: date-time
		LastAccessAt *strfmt.DateTime `json:"last_access_at,omitempty"`
	} `json:"access_resource,omitempty"`
}

// Validate validates this edit resource body
func (o *EditResourceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EditResourceBody) validateAccessResource(formats strfmt.Registry) error {

	if swag.IsZero(o.AccessResource) { // not required
		return nil
	}

	if o.AccessResource.AccessCount != nil {
		if err := o.AccessResource.AccessCount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource" + "." + "access_resource" + "." + "access_count")
			}
			return err
		}
	}

	if err := validate.FormatOf("resource"+"."+"access_resource"+"."+"last_access_at", "body", "date-time", o.AccessResource.LastAccessAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EditResourceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditResourceBody) UnmarshalBinary(b []byte) error {
	var res EditResourceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*EditResourceOKBody edit resource o k body
swagger:model EditResourceOKBody
*/
type EditResourceOKBody struct {
	models.AccessResource

	// access proxy name
	AccessProxyName string `json:"access_proxy_name,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *EditResourceOKBody) UnmarshalJSON(raw []byte) error {
	// EditResourceOKBodyAO0
	var editResourceOKBodyAO0 models.AccessResource
	if err := swag.ReadJSON(raw, &editResourceOKBodyAO0); err != nil {
		return err
	}
	o.AccessResource = editResourceOKBodyAO0

	// EditResourceOKBodyAO1
	var dataEditResourceOKBodyAO1 struct {
		AccessProxyName string `json:"access_proxy_name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataEditResourceOKBodyAO1); err != nil {
		return err
	}

	o.AccessProxyName = dataEditResourceOKBodyAO1.AccessProxyName

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o EditResourceOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	editResourceOKBodyAO0, err := swag.WriteJSON(o.AccessResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, editResourceOKBodyAO0)

	var dataEditResourceOKBodyAO1 struct {
		AccessProxyName string `json:"access_proxy_name,omitempty"`
	}

	dataEditResourceOKBodyAO1.AccessProxyName = o.AccessProxyName

	jsonDataEditResourceOKBodyAO1, errEditResourceOKBodyAO1 := swag.WriteJSON(dataEditResourceOKBodyAO1)
	if errEditResourceOKBodyAO1 != nil {
		return nil, errEditResourceOKBodyAO1
	}
	_parts = append(_parts, jsonDataEditResourceOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this edit resource o k body
func (o *EditResourceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.AccessResource
	if err := o.AccessResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *EditResourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EditResourceOKBody) UnmarshalBinary(b []byte) error {
	var res EditResourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
