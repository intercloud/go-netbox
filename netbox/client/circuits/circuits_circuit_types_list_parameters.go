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
)

// NewCircuitsCircuitTypesListParams creates a new CircuitsCircuitTypesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTypesListParams() *CircuitsCircuitTypesListParams {
	return &CircuitsCircuitTypesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesListParamsWithTimeout creates a new CircuitsCircuitTypesListParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTypesListParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesListParams {
	return &CircuitsCircuitTypesListParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesListParamsWithContext creates a new CircuitsCircuitTypesListParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTypesListParamsWithContext(ctx context.Context) *CircuitsCircuitTypesListParams {
	return &CircuitsCircuitTypesListParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTypesListParamsWithHTTPClient creates a new CircuitsCircuitTypesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTypesListParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesListParams {
	return &CircuitsCircuitTypesListParams{
		HTTPClient: client,
	}
}

/*
CircuitsCircuitTypesListParams contains all the parameters to send to the API endpoint

	for the circuits circuit types list operation.

	Typically these are written to a http.Request.
*/
type CircuitsCircuitTypesListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Slug.
	Slug *string

	// SlugEmpty.
	SlugEmpty *string

	// SlugIc.
	SlugIc *string

	// SlugIe.
	SlugIe *string

	// SlugIew.
	SlugIew *string

	// SlugIsw.
	SlugIsw *string

	// Slugn.
	Slugn *string

	// SlugNic.
	SlugNic *string

	// SlugNie.
	SlugNie *string

	// SlugNiew.
	SlugNiew *string

	// SlugNisw.
	SlugNisw *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit types list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesListParams) WithDefaults() *CircuitsCircuitTypesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit types list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithContext(ctx context.Context) *CircuitsCircuitTypesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithCreated(created *string) *CircuitsCircuitTypesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithCreatedGte(createdGte *string) *CircuitsCircuitTypesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithCreatedLte(createdLte *string) *CircuitsCircuitTypesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithID adds the id to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithID(id *string) *CircuitsCircuitTypesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithIDGt(iDGt *string) *CircuitsCircuitTypesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithIDGte(iDGte *string) *CircuitsCircuitTypesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithIDLt(iDLt *string) *CircuitsCircuitTypesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithIDLte(iDLte *string) *CircuitsCircuitTypesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithIDn(iDn *string) *CircuitsCircuitTypesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithLastUpdated(lastUpdated *string) *CircuitsCircuitTypesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *CircuitsCircuitTypesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *CircuitsCircuitTypesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithLimit(limit *int64) *CircuitsCircuitTypesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithName(name *string) *CircuitsCircuitTypesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameEmpty(nameEmpty *string) *CircuitsCircuitTypesListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameIc(nameIc *string) *CircuitsCircuitTypesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameIe(nameIe *string) *CircuitsCircuitTypesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameIew(nameIew *string) *CircuitsCircuitTypesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameIsw(nameIsw *string) *CircuitsCircuitTypesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNamen(namen *string) *CircuitsCircuitTypesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameNic(nameNic *string) *CircuitsCircuitTypesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameNie(nameNie *string) *CircuitsCircuitTypesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameNiew(nameNiew *string) *CircuitsCircuitTypesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithNameNisw(nameNisw *string) *CircuitsCircuitTypesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithOffset(offset *int64) *CircuitsCircuitTypesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithQ(q *string) *CircuitsCircuitTypesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlug(slug *string) *CircuitsCircuitTypesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugEmpty adds the slugEmpty to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugEmpty(slugEmpty *string) *CircuitsCircuitTypesListParams {
	o.SetSlugEmpty(slugEmpty)
	return o
}

// SetSlugEmpty adds the slugEmpty to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugEmpty(slugEmpty *string) {
	o.SlugEmpty = slugEmpty
}

// WithSlugIc adds the slugIc to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugIc(slugIc *string) *CircuitsCircuitTypesListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugIe(slugIe *string) *CircuitsCircuitTypesListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugIew(slugIew *string) *CircuitsCircuitTypesListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugIsw(slugIsw *string) *CircuitsCircuitTypesListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugn(slugn *string) *CircuitsCircuitTypesListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugNic(slugNic *string) *CircuitsCircuitTypesListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugNie(slugNie *string) *CircuitsCircuitTypesListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugNiew(slugNiew *string) *CircuitsCircuitTypesListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithSlugNisw(slugNisw *string) *CircuitsCircuitTypesListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WithTag adds the tag to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithTag(tag *string) *CircuitsCircuitTypesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) WithTagn(tagn *string) *CircuitsCircuitTypesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the circuits circuit types list params
func (o *CircuitsCircuitTypesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string

		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {

			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}
	}

	if o.SlugEmpty != nil {

		// query param slug__empty
		var qrSlugEmpty string

		if o.SlugEmpty != nil {
			qrSlugEmpty = *o.SlugEmpty
		}
		qSlugEmpty := qrSlugEmpty
		if qSlugEmpty != "" {

			if err := r.SetQueryParam("slug__empty", qSlugEmpty); err != nil {
				return err
			}
		}
	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string

		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {

			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}
	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string

		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {

			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}
	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string

		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {

			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}
	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string

		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {

			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}
	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string

		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {

			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}
	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string

		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {

			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}
	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string

		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {

			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}
	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string

		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {

			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}
	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string

		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {

			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
