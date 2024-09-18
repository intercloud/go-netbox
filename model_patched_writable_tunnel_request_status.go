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

// PatchedWritableTunnelRequestStatus * `planned` - Planned * `active` - Active * `disabled` - Disabled
type PatchedWritableTunnelRequestStatus string

// List of PatchedWritableTunnelRequest_status
const (
	PATCHEDWRITABLETUNNELREQUESTSTATUS_PLANNED  PatchedWritableTunnelRequestStatus = "planned"
	PATCHEDWRITABLETUNNELREQUESTSTATUS_ACTIVE   PatchedWritableTunnelRequestStatus = "active"
	PATCHEDWRITABLETUNNELREQUESTSTATUS_DISABLED PatchedWritableTunnelRequestStatus = "disabled"
)

// All allowed values of PatchedWritableTunnelRequestStatus enum
var AllowedPatchedWritableTunnelRequestStatusEnumValues = []PatchedWritableTunnelRequestStatus{
	"planned",
	"active",
	"disabled",
}

func (v *PatchedWritableTunnelRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritableTunnelRequestStatus(value)
	for _, existing := range AllowedPatchedWritableTunnelRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritableTunnelRequestStatus", value)
}

// NewPatchedWritableTunnelRequestStatusFromValue returns a pointer to a valid PatchedWritableTunnelRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritableTunnelRequestStatusFromValue(v string) (*PatchedWritableTunnelRequestStatus, error) {
	ev := PatchedWritableTunnelRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritableTunnelRequestStatus: valid values are %v", v, AllowedPatchedWritableTunnelRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritableTunnelRequestStatus) IsValid() bool {
	for _, existing := range AllowedPatchedWritableTunnelRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritableTunnelRequest_status value
func (v PatchedWritableTunnelRequestStatus) Ptr() *PatchedWritableTunnelRequestStatus {
	return &v
}

type NullablePatchedWritableTunnelRequestStatus struct {
	value *PatchedWritableTunnelRequestStatus
	isSet bool
}

func (v NullablePatchedWritableTunnelRequestStatus) Get() *PatchedWritableTunnelRequestStatus {
	return v.value
}

func (v *NullablePatchedWritableTunnelRequestStatus) Set(val *PatchedWritableTunnelRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableTunnelRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableTunnelRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableTunnelRequestStatus(val *PatchedWritableTunnelRequestStatus) *NullablePatchedWritableTunnelRequestStatus {
	return &NullablePatchedWritableTunnelRequestStatus{value: val, isSet: true}
}

func (v NullablePatchedWritableTunnelRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableTunnelRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
