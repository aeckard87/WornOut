// Code generated by go-swagger; DO NOT EDIT.

package descriptors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new descriptors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for descriptors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDescriptor creates a descriptor
*/
func (a *Client) CreateDescriptor(params *CreateDescriptorParams) (*CreateDescriptorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDescriptorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDescriptor",
		Method:             "POST",
		PathPattern:        "/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDescriptorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDescriptorCreated), nil

}

/*
CreateDescriptorByDetail creates a descriptor by detail
*/
func (a *Client) CreateDescriptorByDetail(params *CreateDescriptorByDetailParams) (*CreateDescriptorByDetailCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDescriptorByDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDescriptorByDetail",
		Method:             "POST",
		PathPattern:        "/details/{id}/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDescriptorByDetailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDescriptorByDetailCreated), nil

}

/*
DeleteDescriptor deletes descriptor
*/
func (a *Client) DeleteDescriptor(params *DeleteDescriptorParams) (*DeleteDescriptorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDescriptorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDescriptor",
		Method:             "DELETE",
		PathPattern:        "/Descriptors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDescriptorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDescriptorOK), nil

}

/*
DeleteDescriptors deletes all descriptors
*/
func (a *Client) DeleteDescriptors(params *DeleteDescriptorsParams) (*DeleteDescriptorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDescriptorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDescriptors",
		Method:             "DELETE",
		PathPattern:        "/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDescriptorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDescriptorsOK), nil

}

/*
DeleteDescriptorsByDetail deletes all descriptors by detail
*/
func (a *Client) DeleteDescriptorsByDetail(params *DeleteDescriptorsByDetailParams) (*DeleteDescriptorsByDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDescriptorsByDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDescriptorsByDetail",
		Method:             "DELETE",
		PathPattern:        "/details/{id}/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDescriptorsByDetailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDescriptorsByDetailOK), nil

}

/*
GetDescriptor gets a specific descriptor
*/
func (a *Client) GetDescriptor(params *GetDescriptorParams) (*GetDescriptorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDescriptorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDescriptor",
		Method:             "GET",
		PathPattern:        "/Descriptors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDescriptorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDescriptorOK), nil

}

/*
GetDescriptors gets all descriptors
*/
func (a *Client) GetDescriptors(params *GetDescriptorsParams) (*GetDescriptorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDescriptorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDescriptors",
		Method:             "GET",
		PathPattern:        "/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDescriptorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDescriptorsOK), nil

}

/*
GetDescriptorsByDetail gets all descriptors by detail
*/
func (a *Client) GetDescriptorsByDetail(params *GetDescriptorsByDetailParams) (*GetDescriptorsByDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDescriptorsByDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDescriptorsByDetail",
		Method:             "GET",
		PathPattern:        "/details/{id}/descriptors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDescriptorsByDetailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDescriptorsByDetailOK), nil

}

/*
UpdateDescriptor updates a specific descriptor
*/
func (a *Client) UpdateDescriptor(params *UpdateDescriptorParams) (*UpdateDescriptorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDescriptorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateDescriptor",
		Method:             "PUT",
		PathPattern:        "/Descriptors/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDescriptorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDescriptorCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
