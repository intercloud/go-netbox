// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/intercloud/go-netbox/netbox/models"
)

// NewVirtualizationInterfacesCreateParams creates a new VirtualizationInterfacesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationInterfacesCreateParams() *VirtualizationInterfacesCreateParams {
	return &VirtualizationInterfacesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesCreateParamsWithTimeout creates a new VirtualizationInterfacesCreateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationInterfacesCreateParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesCreateParams {
	return &VirtualizationInterfacesCreateParams{
		timeout: timeout,
	}
}

// NewVirtualizationInterfacesCreateParamsWithContext creates a new VirtualizationInterfacesCreateParams object
// with the ability to set a context for a request.
func NewVirtualizationInterfacesCreateParamsWithContext(ctx context.Context) *VirtualizationInterfacesCreateParams {
	return &VirtualizationInterfacesCreateParams{
		Context: ctx,
	}
}

// NewVirtualizationInterfacesCreateParamsWithHTTPClient creates a new VirtualizationInterfacesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationInterfacesCreateParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesCreateParams {
	return &VirtualizationInterfacesCreateParams{
		HTTPClient: client,
	}
}

/*
VirtualizationInterfacesCreateParams contains all the parameters to send to the API endpoint

	for the virtualization interfaces create operation.

	Typically these are written to a http.Request.
*/
type VirtualizationInterfacesCreateParams struct {

	// Data.
	Data *models.WritableVMInterface

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization interfaces create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesCreateParams) WithDefaults() *VirtualizationInterfacesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization interfaces create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) WithContext(ctx context.Context) *VirtualizationInterfacesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) WithData(data *models.WritableVMInterface) *VirtualizationInterfacesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization interfaces create params
func (o *VirtualizationInterfacesCreateParams) SetData(data *models.WritableVMInterface) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
