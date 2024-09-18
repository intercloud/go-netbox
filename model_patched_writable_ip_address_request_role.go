/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.1 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// PatchedWritableIPAddressRequestRole The functional role of this IP  * `loopback` - Loopback * `secondary` - Secondary * `anycast` - Anycast * `vip` - VIP * `vrrp` - VRRP * `hsrp` - HSRP * `glbp` - GLBP * `carp` - CARP
type PatchedWritableIPAddressRequestRole string

// List of PatchedWritableIPAddressRequest_role
const (
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_LOOPBACK  PatchedWritableIPAddressRequestRole = "loopback"
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_SECONDARY PatchedWritableIPAddressRequestRole = "secondary"
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_ANYCAST   PatchedWritableIPAddressRequestRole = "anycast"
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_VIP       PatchedWritableIPAddressRequestRole = "vip"
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_VRRP      PatchedWritableIPAddressRequestRole = "vrrp"
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_HSRP      PatchedWritableIPAddressRequestRole = "hsrp"
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_GLBP      PatchedWritableIPAddressRequestRole = "glbp"
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_CARP      PatchedWritableIPAddressRequestRole = "carp"
	PATCHEDWRITABLEIPADDRESSREQUESTROLE_EMPTY     PatchedWritableIPAddressRequestRole = ""
)

// All allowed values of PatchedWritableIPAddressRequestRole enum
var AllowedPatchedWritableIPAddressRequestRoleEnumValues = []PatchedWritableIPAddressRequestRole{
	"loopback",
	"secondary",
	"anycast",
	"vip",
	"vrrp",
	"hsrp",
	"glbp",
	"carp",
	"",
}

func (v *PatchedWritableIPAddressRequestRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableIPAddressRequestRole(value)
	for _, existing := range AllowedPatchedWritableIPAddressRequestRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableIPAddressRequestRole", value)
}

// NewPatchedWritableIPAddressRequestRoleFromValue returns a pointer to a valid PatchedWritableIPAddressRequestRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableIPAddressRequestRoleFromValue(v string) (*PatchedWritableIPAddressRequestRole, error) {
	ev := PatchedWritableIPAddressRequestRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableIPAddressRequestRole: valid values are %v", v, AllowedPatchedWritableIPAddressRequestRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableIPAddressRequestRole) IsValid() bool {
	for _, existing := range AllowedPatchedWritableIPAddressRequestRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableIPAddressRequest_role value
func (v PatchedWritableIPAddressRequestRole) Ptr() *PatchedWritableIPAddressRequestRole {
	return &v
}

type NullablePatchedWritableIPAddressRequestRole struct {
	value *PatchedWritableIPAddressRequestRole
	isSet bool
}

func (v NullablePatchedWritableIPAddressRequestRole) Get() *PatchedWritableIPAddressRequestRole {
	return v.value
}

func (v *NullablePatchedWritableIPAddressRequestRole) Set(val *PatchedWritableIPAddressRequestRole) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIPAddressRequestRole) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIPAddressRequestRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIPAddressRequestRole(val *PatchedWritableIPAddressRequestRole) *NullablePatchedWritableIPAddressRequestRole {
	return &NullablePatchedWritableIPAddressRequestRole{value: val, isSet: true}
}

func (v NullablePatchedWritableIPAddressRequestRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIPAddressRequestRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
