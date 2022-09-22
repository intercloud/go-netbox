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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableIPAddress writable IP address
//
// swagger:model WritableIPAddress
type WritableIPAddress struct {

	// Address
	//
	// IPv4 or IPv6 address (with mask)
	// Required: true
	Address *string `json:"address"`

	// Assigned object
	// Read Only: true
	AssignedObject map[string]*string `json:"assigned_object,omitempty"`

	// Assigned object id
	// Maximum: 2.147483647e+09
	// Minimum: 0
	AssignedObjectID *int64 `json:"assigned_object_id,omitempty"`

	// Assigned object type
	AssignedObjectType *string `json:"assigned_object_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// DNS Name
	//
	// Hostname or FQDN (not case-sensitive)
	// Max Length: 255
	// Pattern: ^[0-9A-Za-z._-]+$
	DNSName string `json:"dns_name,omitempty"`

	// Family
	// Read Only: true
	Family string `json:"family,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// NAT (Inside)
	//
	// The IP for which this address is the "outside" IP
	NatInside *int64 `json:"nat_inside,omitempty"`

	// Nat outside
	// Read Only: true
	NatOutside string `json:"nat_outside,omitempty"`

	// Role
	//
	// The functional role of this IP
	// Enum: [loopback secondary anycast vip vrrp hsrp glbp carp]
	Role string `json:"role,omitempty"`

	// Status
	//
	// The operational status of this IP
	// Enum: [active reserved deprecated dhcp slaac]
	Status string `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// VRF
	Vrf *int64 `json:"vrf,omitempty"`
}

// Validate validates this writable IP address
func (m *WritableIPAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableIPAddress) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateAssignedObjectID(formats strfmt.Registry) error {
	if swag.IsZero(m.AssignedObjectID) { // not required
		return nil
	}

	if err := validate.MinimumInt("assigned_object_id", "body", *m.AssignedObjectID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("assigned_object_id", "body", *m.AssignedObjectID, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateDNSName(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSName) { // not required
		return nil
	}

	if err := validate.MaxLength("dns_name", "body", m.DNSName, 255); err != nil {
		return err
	}

	if err := validate.Pattern("dns_name", "body", m.DNSName, `^[0-9A-Za-z._-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

var writableIpAddressTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["loopback","secondary","anycast","vip","vrrp","hsrp","glbp","carp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableIpAddressTypeRolePropEnum = append(writableIpAddressTypeRolePropEnum, v)
	}
}

const (

	// WritableIPAddressRoleLoopback captures enum value "loopback"
	WritableIPAddressRoleLoopback string = "loopback"

	// WritableIPAddressRoleSecondary captures enum value "secondary"
	WritableIPAddressRoleSecondary string = "secondary"

	// WritableIPAddressRoleAnycast captures enum value "anycast"
	WritableIPAddressRoleAnycast string = "anycast"

	// WritableIPAddressRoleVip captures enum value "vip"
	WritableIPAddressRoleVip string = "vip"

	// WritableIPAddressRoleVrrp captures enum value "vrrp"
	WritableIPAddressRoleVrrp string = "vrrp"

	// WritableIPAddressRoleHsrp captures enum value "hsrp"
	WritableIPAddressRoleHsrp string = "hsrp"

	// WritableIPAddressRoleGlbp captures enum value "glbp"
	WritableIPAddressRoleGlbp string = "glbp"

	// WritableIPAddressRoleCarp captures enum value "carp"
	WritableIPAddressRoleCarp string = "carp"
)

// prop value enum
func (m *WritableIPAddress) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableIpAddressTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableIPAddress) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

var writableIpAddressTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","reserved","deprecated","dhcp","slaac"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableIpAddressTypeStatusPropEnum = append(writableIpAddressTypeStatusPropEnum, v)
	}
}

const (

	// WritableIPAddressStatusActive captures enum value "active"
	WritableIPAddressStatusActive string = "active"

	// WritableIPAddressStatusReserved captures enum value "reserved"
	WritableIPAddressStatusReserved string = "reserved"

	// WritableIPAddressStatusDeprecated captures enum value "deprecated"
	WritableIPAddressStatusDeprecated string = "deprecated"

	// WritableIPAddressStatusDhcp captures enum value "dhcp"
	WritableIPAddressStatusDhcp string = "dhcp"

	// WritableIPAddressStatusSlaac captures enum value "slaac"
	WritableIPAddressStatusSlaac string = "slaac"
)

// prop value enum
func (m *WritableIPAddress) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableIpAddressTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableIPAddress) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableIPAddress) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable IP address based on the context it is used
func (m *WritableIPAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignedObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNatOutside(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableIPAddress) contextValidateAssignedObject(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *WritableIPAddress) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateFamily(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "family", "body", string(m.Family)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateNatOutside(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "nat_outside", "body", string(m.NatOutside)); err != nil {
		return err
	}

	return nil
}

func (m *WritableIPAddress) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableIPAddress) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableIPAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableIPAddress) UnmarshalBinary(b []byte) error {
	var res WritableIPAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
