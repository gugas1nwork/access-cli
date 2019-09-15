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

// GetResourceReader is a Reader for the GetResource structure.
type GetResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResourceOK creates a GetResourceOK with default headers values
func NewGetResourceOK() *GetResourceOK {
	return &GetResourceOK{}
}

/*GetResourceOK handles this case with default header values.

Resource info
*/
type GetResourceOK struct {
	Payload *GetResourceOKBody
}

func (o *GetResourceOK) Error() string {
	return fmt.Sprintf("[GET /access_resources/{id}][%d] getResourceOK  %+v", 200, o.Payload)
}

func (o *GetResourceOK) GetPayload() *GetResourceOKBody {
	return o.Payload
}

func (o *GetResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetResourceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceUnauthorized creates a GetResourceUnauthorized with default headers values
func NewGetResourceUnauthorized() *GetResourceUnauthorized {
	return &GetResourceUnauthorized{}
}

/*GetResourceUnauthorized handles this case with default header values.

unauthorized: invalid credentials or missing authentication headers
*/
type GetResourceUnauthorized struct {
	Payload *models.GenericUnauthorizedResponse
}

func (o *GetResourceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /access_resources/{id}][%d] getResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResourceUnauthorized) GetPayload() *models.GenericUnauthorizedResponse {
	return o.Payload
}

func (o *GetResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericUnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceNotFound creates a GetResourceNotFound with default headers values
func NewGetResourceNotFound() *GetResourceNotFound {
	return &GetResourceNotFound{}
}

/*GetResourceNotFound handles this case with default header values.

resource not found
*/
type GetResourceNotFound struct {
}

func (o *GetResourceNotFound) Error() string {
	return fmt.Sprintf("[GET /access_resources/{id}][%d] getResourceNotFound ", 404)
}

func (o *GetResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetResourceOKBody get resource o k body
swagger:model GetResourceOKBody
*/
type GetResourceOKBody struct {
	models.AccessResource

	// access count
	AccessCount *GetResourceOKBodyAO1AccessCount `json:"access_count,omitempty"`

	// access proxy name
	AccessProxyName string `json:"access_proxy_name,omitempty"`

	// last access at
	// Format: date-time
	LastAccessAt *strfmt.DateTime `json:"last_access_at,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetResourceOKBody) UnmarshalJSON(raw []byte) error {
	// GetResourceOKBodyAO0
	var getResourceOKBodyAO0 models.AccessResource
	if err := swag.ReadJSON(raw, &getResourceOKBodyAO0); err != nil {
		return err
	}
	o.AccessResource = getResourceOKBodyAO0

	// GetResourceOKBodyAO1
	var dataGetResourceOKBodyAO1 struct {
		AccessCount *GetResourceOKBodyAO1AccessCount `json:"access_count,omitempty"`

		AccessProxyName string `json:"access_proxy_name,omitempty"`

		LastAccessAt *strfmt.DateTime `json:"last_access_at,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetResourceOKBodyAO1); err != nil {
		return err
	}

	o.AccessCount = dataGetResourceOKBodyAO1.AccessCount

	o.AccessProxyName = dataGetResourceOKBodyAO1.AccessProxyName

	o.LastAccessAt = dataGetResourceOKBodyAO1.LastAccessAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetResourceOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getResourceOKBodyAO0, err := swag.WriteJSON(o.AccessResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getResourceOKBodyAO0)

	var dataGetResourceOKBodyAO1 struct {
		AccessCount *GetResourceOKBodyAO1AccessCount `json:"access_count,omitempty"`

		AccessProxyName string `json:"access_proxy_name,omitempty"`

		LastAccessAt *strfmt.DateTime `json:"last_access_at,omitempty"`
	}

	dataGetResourceOKBodyAO1.AccessCount = o.AccessCount

	dataGetResourceOKBodyAO1.AccessProxyName = o.AccessProxyName

	dataGetResourceOKBodyAO1.LastAccessAt = o.LastAccessAt

	jsonDataGetResourceOKBodyAO1, errGetResourceOKBodyAO1 := swag.WriteJSON(dataGetResourceOKBodyAO1)
	if errGetResourceOKBodyAO1 != nil {
		return nil, errGetResourceOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetResourceOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get resource o k body
func (o *GetResourceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.AccessResource
	if err := o.AccessResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAccessCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastAccessAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourceOKBody) validateAccessCount(formats strfmt.Registry) error {

	if swag.IsZero(o.AccessCount) { // not required
		return nil
	}

	if o.AccessCount != nil {
		if err := o.AccessCount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getResourceOK" + "." + "access_count")
			}
			return err
		}
	}

	return nil
}

func (o *GetResourceOKBody) validateLastAccessAt(formats strfmt.Registry) error {

	if swag.IsZero(o.LastAccessAt) { // not required
		return nil
	}

	if err := validate.FormatOf("getResourceOK"+"."+"last_access_at", "body", "date-time", o.LastAccessAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceOKBody) UnmarshalBinary(b []byte) error {
	var res GetResourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetResourceOKBodyAO1AccessCount get resource o k body a o1 access count
swagger:model GetResourceOKBodyAO1AccessCount
*/
type GetResourceOKBodyAO1AccessCount struct {

	// denied
	Denied int64 `json:"denied,omitempty"`

	// granted
	Granted int64 `json:"granted,omitempty"`
}

// Validate validates this get resource o k body a o1 access count
func (o *GetResourceOKBodyAO1AccessCount) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetResourceOKBodyAO1AccessCount) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetResourceOKBodyAO1AccessCount) UnmarshalBinary(b []byte) error {
	var res GetResourceOKBodyAO1AccessCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
