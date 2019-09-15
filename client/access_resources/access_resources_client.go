// Code generated by go-swagger; DO NOT EDIT.

package access_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new access resources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access resources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateResource creates resource
*/
func (a *Client) CreateResource(params *CreateResourceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateResourceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createResource",
		Method:             "POST",
		PathPattern:        "/access_resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateResourceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteResource deletes resource
*/
func (a *Client) DeleteResource(params *DeleteResourceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteResourceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteResource",
		Method:             "DELETE",
		PathPattern:        "/access_resources/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteResourceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
EditResource edits resource
*/
func (a *Client) EditResource(params *EditResourceParams, authInfo runtime.ClientAuthInfoWriter) (*EditResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editResource",
		Method:             "PATCH",
		PathPattern:        "/access_resources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EditResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for editResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetResource retrieves information about a resource
*/
func (a *Client) GetResource(params *GetResourceParams, authInfo runtime.ClientAuthInfoWriter) (*GetResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResource",
		Method:             "GET",
		PathPattern:        "/access_resources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListResources lists resources
*/
func (a *Client) ListResources(params *ListResourcesParams, authInfo runtime.ClientAuthInfoWriter) (*ListResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listResources",
		Method:             "GET",
		PathPattern:        "/access_resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListResourcesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
