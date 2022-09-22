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

package users

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

// NewUsersPermissionsPartialUpdateParams creates a new UsersPermissionsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersPermissionsPartialUpdateParams() *UsersPermissionsPartialUpdateParams {
	return &UsersPermissionsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPermissionsPartialUpdateParamsWithTimeout creates a new UsersPermissionsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewUsersPermissionsPartialUpdateParamsWithTimeout(timeout time.Duration) *UsersPermissionsPartialUpdateParams {
	return &UsersPermissionsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewUsersPermissionsPartialUpdateParamsWithContext creates a new UsersPermissionsPartialUpdateParams object
// with the ability to set a context for a request.
func NewUsersPermissionsPartialUpdateParamsWithContext(ctx context.Context) *UsersPermissionsPartialUpdateParams {
	return &UsersPermissionsPartialUpdateParams{
		Context: ctx,
	}
}

// NewUsersPermissionsPartialUpdateParamsWithHTTPClient creates a new UsersPermissionsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersPermissionsPartialUpdateParamsWithHTTPClient(client *http.Client) *UsersPermissionsPartialUpdateParams {
	return &UsersPermissionsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
UsersPermissionsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the users permissions partial update operation.

	Typically these are written to a http.Request.
*/
type UsersPermissionsPartialUpdateParams struct {

	// Data.
	Data *models.WritableObjectPermission

	/* ID.

	   A unique integer value identifying this permission.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users permissions partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsPartialUpdateParams) WithDefaults() *UsersPermissionsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users permissions partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) WithTimeout(timeout time.Duration) *UsersPermissionsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) WithContext(ctx context.Context) *UsersPermissionsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) WithHTTPClient(client *http.Client) *UsersPermissionsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) WithData(data *models.WritableObjectPermission) *UsersPermissionsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) SetData(data *models.WritableObjectPermission) {
	o.Data = data
}

// WithID adds the id to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) WithID(id int64) *UsersPermissionsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users permissions partial update params
func (o *UsersPermissionsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPermissionsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
