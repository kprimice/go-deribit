// Code generated by go-swagger; DO NOT EDIT.

package supporting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new supporting API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for supporting API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetPublicGetTime retrieves the current time in milliseconds this API endpoint can be used to check the clock skew between your software and deribit s systems
*/
func (a *Client) GetPublicGetTime(params *GetPublicGetTimeParams) (*GetPublicGetTimeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicGetTimeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPublicGetTime",
		Method:             "GET",
		PathPattern:        "/public/get_time",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPublicGetTimeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPublicGetTimeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPublicGetTime: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
