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

package circuits

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
	"github.com/go-openapi/swag"

	"github.com/intercloud/go-netbox/netbox/models"
)

// NewCircuitsCircuitTypesUpdateParams creates a new CircuitsCircuitTypesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTypesUpdateParams() *CircuitsCircuitTypesUpdateParams {
	return &CircuitsCircuitTypesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesUpdateParamsWithTimeout creates a new CircuitsCircuitTypesUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTypesUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesUpdateParams {
	return &CircuitsCircuitTypesUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesUpdateParamsWithContext creates a new CircuitsCircuitTypesUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTypesUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTypesUpdateParams {
	return &CircuitsCircuitTypesUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTypesUpdateParamsWithHTTPClient creates a new CircuitsCircuitTypesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTypesUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesUpdateParams {
	return &CircuitsCircuitTypesUpdateParams{
		HTTPClient: client,
	}
}

/*
CircuitsCircuitTypesUpdateParams contains all the parameters to send to the API endpoint

	for the circuits circuit types update operation.

	Typically these are written to a http.Request.
*/
type CircuitsCircuitTypesUpdateParams struct {

	// Data.
	Data *models.CircuitType

	/* ID.

	   A unique integer value identifying this circuit type.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesUpdateParams) WithDefaults() *CircuitsCircuitTypesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit types update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithData(data *models.CircuitType) *CircuitsCircuitTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetData(data *models.CircuitType) {
	o.Data = data
}

// WithID adds the id to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithID(id int64) *CircuitsCircuitTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
