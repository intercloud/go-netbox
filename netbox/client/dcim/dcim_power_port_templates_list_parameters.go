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

package dcim

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

// NewDcimPowerPortTemplatesListParams creates a new DcimPowerPortTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortTemplatesListParams() *DcimPowerPortTemplatesListParams {
	return &DcimPowerPortTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesListParamsWithTimeout creates a new DcimPowerPortTemplatesListParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesListParams {
	return &DcimPowerPortTemplatesListParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesListParamsWithContext creates a new DcimPowerPortTemplatesListParams object
// with the ability to set a context for a request.
func NewDcimPowerPortTemplatesListParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesListParams {
	return &DcimPowerPortTemplatesListParams{
		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesListParamsWithHTTPClient creates a new DcimPowerPortTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesListParams {
	return &DcimPowerPortTemplatesListParams{
		HTTPClient: client,
	}
}

/*
DcimPowerPortTemplatesListParams contains all the parameters to send to the API endpoint

	for the dcim power port templates list operation.

	Typically these are written to a http.Request.
*/
type DcimPowerPortTemplatesListParams struct {

	// AllocatedDraw.
	AllocatedDraw *string

	// AllocatedDrawGt.
	AllocatedDrawGt *string

	// AllocatedDrawGte.
	AllocatedDrawGte *string

	// AllocatedDrawLt.
	AllocatedDrawLt *string

	// AllocatedDrawLte.
	AllocatedDrawLte *string

	// AllocatedDrawn.
	AllocatedDrawn *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// DevicetypeID.
	DevicetypeID *string

	// DevicetypeIDn.
	DevicetypeIDn *string

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

	// MaximumDraw.
	MaximumDraw *string

	// MaximumDrawGt.
	MaximumDrawGt *string

	// MaximumDrawGte.
	MaximumDrawGte *string

	// MaximumDrawLt.
	MaximumDrawLt *string

	// MaximumDrawLte.
	MaximumDrawLte *string

	// MaximumDrawn.
	MaximumDrawn *string

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

	// Type.
	Type *string

	// Typen.
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesListParams) WithDefaults() *DcimPowerPortTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatedDraw adds the allocatedDraw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDraw(allocatedDraw *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDraw(allocatedDraw)
	return o
}

// SetAllocatedDraw adds the allocatedDraw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDraw(allocatedDraw *string) {
	o.AllocatedDraw = allocatedDraw
}

// WithAllocatedDrawGt adds the allocatedDrawGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawGt(allocatedDrawGt *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawGt(allocatedDrawGt)
	return o
}

// SetAllocatedDrawGt adds the allocatedDrawGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawGt(allocatedDrawGt *string) {
	o.AllocatedDrawGt = allocatedDrawGt
}

// WithAllocatedDrawGte adds the allocatedDrawGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawGte(allocatedDrawGte *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawGte(allocatedDrawGte)
	return o
}

// SetAllocatedDrawGte adds the allocatedDrawGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawGte(allocatedDrawGte *string) {
	o.AllocatedDrawGte = allocatedDrawGte
}

// WithAllocatedDrawLt adds the allocatedDrawLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawLt(allocatedDrawLt *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawLt(allocatedDrawLt)
	return o
}

// SetAllocatedDrawLt adds the allocatedDrawLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawLt(allocatedDrawLt *string) {
	o.AllocatedDrawLt = allocatedDrawLt
}

// WithAllocatedDrawLte adds the allocatedDrawLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawLte(allocatedDrawLte *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawLte(allocatedDrawLte)
	return o
}

// SetAllocatedDrawLte adds the allocatedDrawLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawLte(allocatedDrawLte *string) {
	o.AllocatedDrawLte = allocatedDrawLte
}

// WithAllocatedDrawn adds the allocatedDrawn to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithAllocatedDrawn(allocatedDrawn *string) *DcimPowerPortTemplatesListParams {
	o.SetAllocatedDrawn(allocatedDrawn)
	return o
}

// SetAllocatedDrawn adds the allocatedDrawN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetAllocatedDrawn(allocatedDrawn *string) {
	o.AllocatedDrawn = allocatedDrawn
}

// WithCreated adds the created to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithCreated(created *string) *DcimPowerPortTemplatesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithCreatedGte(createdGte *string) *DcimPowerPortTemplatesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithCreatedLte(createdLte *string) *DcimPowerPortTemplatesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDevicetypeID adds the devicetypeID to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimPowerPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimPowerPortTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithID(id *string) *DcimPowerPortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDGt(iDGt *string) *DcimPowerPortTemplatesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDGte(iDGte *string) *DcimPowerPortTemplatesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDLt(iDLt *string) *DcimPowerPortTemplatesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDLte(iDLte *string) *DcimPowerPortTemplatesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithIDn(iDn *string) *DcimPowerPortTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithLastUpdated(lastUpdated *string) *DcimPowerPortTemplatesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimPowerPortTemplatesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimPowerPortTemplatesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithLimit(limit *int64) *DcimPowerPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaximumDraw adds the maximumDraw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDraw(maximumDraw *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDraw(maximumDraw)
	return o
}

// SetMaximumDraw adds the maximumDraw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDraw(maximumDraw *string) {
	o.MaximumDraw = maximumDraw
}

// WithMaximumDrawGt adds the maximumDrawGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawGt(maximumDrawGt *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawGt(maximumDrawGt)
	return o
}

// SetMaximumDrawGt adds the maximumDrawGt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawGt(maximumDrawGt *string) {
	o.MaximumDrawGt = maximumDrawGt
}

// WithMaximumDrawGte adds the maximumDrawGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawGte(maximumDrawGte *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawGte(maximumDrawGte)
	return o
}

// SetMaximumDrawGte adds the maximumDrawGte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawGte(maximumDrawGte *string) {
	o.MaximumDrawGte = maximumDrawGte
}

// WithMaximumDrawLt adds the maximumDrawLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawLt(maximumDrawLt *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawLt(maximumDrawLt)
	return o
}

// SetMaximumDrawLt adds the maximumDrawLt to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawLt(maximumDrawLt *string) {
	o.MaximumDrawLt = maximumDrawLt
}

// WithMaximumDrawLte adds the maximumDrawLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawLte(maximumDrawLte *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawLte(maximumDrawLte)
	return o
}

// SetMaximumDrawLte adds the maximumDrawLte to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawLte(maximumDrawLte *string) {
	o.MaximumDrawLte = maximumDrawLte
}

// WithMaximumDrawn adds the maximumDrawn to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithMaximumDrawn(maximumDrawn *string) *DcimPowerPortTemplatesListParams {
	o.SetMaximumDrawn(maximumDrawn)
	return o
}

// SetMaximumDrawn adds the maximumDrawN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetMaximumDrawn(maximumDrawn *string) {
	o.MaximumDrawn = maximumDrawn
}

// WithName adds the name to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithName(name *string) *DcimPowerPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameEmpty(nameEmpty *string) *DcimPowerPortTemplatesListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameIc(nameIc *string) *DcimPowerPortTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameIe(nameIe *string) *DcimPowerPortTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameIew(nameIew *string) *DcimPowerPortTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameIsw(nameIsw *string) *DcimPowerPortTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNamen(namen *string) *DcimPowerPortTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameNic(nameNic *string) *DcimPowerPortTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameNie(nameNie *string) *DcimPowerPortTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameNiew(nameNiew *string) *DcimPowerPortTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithNameNisw(nameNisw *string) *DcimPowerPortTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithOffset(offset *int64) *DcimPowerPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithQ(q *string) *DcimPowerPortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithType adds the typeVar to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithType(typeVar *string) *DcimPowerPortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) WithTypen(typen *string) *DcimPowerPortTemplatesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim power port templates list params
func (o *DcimPowerPortTemplatesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllocatedDraw != nil {

		// query param allocated_draw
		var qrAllocatedDraw string

		if o.AllocatedDraw != nil {
			qrAllocatedDraw = *o.AllocatedDraw
		}
		qAllocatedDraw := qrAllocatedDraw
		if qAllocatedDraw != "" {

			if err := r.SetQueryParam("allocated_draw", qAllocatedDraw); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawGt != nil {

		// query param allocated_draw__gt
		var qrAllocatedDrawGt string

		if o.AllocatedDrawGt != nil {
			qrAllocatedDrawGt = *o.AllocatedDrawGt
		}
		qAllocatedDrawGt := qrAllocatedDrawGt
		if qAllocatedDrawGt != "" {

			if err := r.SetQueryParam("allocated_draw__gt", qAllocatedDrawGt); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawGte != nil {

		// query param allocated_draw__gte
		var qrAllocatedDrawGte string

		if o.AllocatedDrawGte != nil {
			qrAllocatedDrawGte = *o.AllocatedDrawGte
		}
		qAllocatedDrawGte := qrAllocatedDrawGte
		if qAllocatedDrawGte != "" {

			if err := r.SetQueryParam("allocated_draw__gte", qAllocatedDrawGte); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawLt != nil {

		// query param allocated_draw__lt
		var qrAllocatedDrawLt string

		if o.AllocatedDrawLt != nil {
			qrAllocatedDrawLt = *o.AllocatedDrawLt
		}
		qAllocatedDrawLt := qrAllocatedDrawLt
		if qAllocatedDrawLt != "" {

			if err := r.SetQueryParam("allocated_draw__lt", qAllocatedDrawLt); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawLte != nil {

		// query param allocated_draw__lte
		var qrAllocatedDrawLte string

		if o.AllocatedDrawLte != nil {
			qrAllocatedDrawLte = *o.AllocatedDrawLte
		}
		qAllocatedDrawLte := qrAllocatedDrawLte
		if qAllocatedDrawLte != "" {

			if err := r.SetQueryParam("allocated_draw__lte", qAllocatedDrawLte); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawn != nil {

		// query param allocated_draw__n
		var qrAllocatedDrawn string

		if o.AllocatedDrawn != nil {
			qrAllocatedDrawn = *o.AllocatedDrawn
		}
		qAllocatedDrawn := qrAllocatedDrawn
		if qAllocatedDrawn != "" {

			if err := r.SetQueryParam("allocated_draw__n", qAllocatedDrawn); err != nil {
				return err
			}
		}
	}

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

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string

		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {

			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}
	}

	if o.DevicetypeIDn != nil {

		// query param devicetype_id__n
		var qrDevicetypeIDn string

		if o.DevicetypeIDn != nil {
			qrDevicetypeIDn = *o.DevicetypeIDn
		}
		qDevicetypeIDn := qrDevicetypeIDn
		if qDevicetypeIDn != "" {

			if err := r.SetQueryParam("devicetype_id__n", qDevicetypeIDn); err != nil {
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

	if o.MaximumDraw != nil {

		// query param maximum_draw
		var qrMaximumDraw string

		if o.MaximumDraw != nil {
			qrMaximumDraw = *o.MaximumDraw
		}
		qMaximumDraw := qrMaximumDraw
		if qMaximumDraw != "" {

			if err := r.SetQueryParam("maximum_draw", qMaximumDraw); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawGt != nil {

		// query param maximum_draw__gt
		var qrMaximumDrawGt string

		if o.MaximumDrawGt != nil {
			qrMaximumDrawGt = *o.MaximumDrawGt
		}
		qMaximumDrawGt := qrMaximumDrawGt
		if qMaximumDrawGt != "" {

			if err := r.SetQueryParam("maximum_draw__gt", qMaximumDrawGt); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawGte != nil {

		// query param maximum_draw__gte
		var qrMaximumDrawGte string

		if o.MaximumDrawGte != nil {
			qrMaximumDrawGte = *o.MaximumDrawGte
		}
		qMaximumDrawGte := qrMaximumDrawGte
		if qMaximumDrawGte != "" {

			if err := r.SetQueryParam("maximum_draw__gte", qMaximumDrawGte); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawLt != nil {

		// query param maximum_draw__lt
		var qrMaximumDrawLt string

		if o.MaximumDrawLt != nil {
			qrMaximumDrawLt = *o.MaximumDrawLt
		}
		qMaximumDrawLt := qrMaximumDrawLt
		if qMaximumDrawLt != "" {

			if err := r.SetQueryParam("maximum_draw__lt", qMaximumDrawLt); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawLte != nil {

		// query param maximum_draw__lte
		var qrMaximumDrawLte string

		if o.MaximumDrawLte != nil {
			qrMaximumDrawLte = *o.MaximumDrawLte
		}
		qMaximumDrawLte := qrMaximumDrawLte
		if qMaximumDrawLte != "" {

			if err := r.SetQueryParam("maximum_draw__lte", qMaximumDrawLte); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawn != nil {

		// query param maximum_draw__n
		var qrMaximumDrawn string

		if o.MaximumDrawn != nil {
			qrMaximumDrawn = *o.MaximumDrawn
		}
		qMaximumDrawn := qrMaximumDrawn
		if qMaximumDrawn != "" {

			if err := r.SetQueryParam("maximum_draw__n", qMaximumDrawn); err != nil {
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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Typen != nil {

		// query param type__n
		var qrTypen string

		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {

			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
