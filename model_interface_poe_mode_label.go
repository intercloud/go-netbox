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

// InterfacePoeModeLabel the model 'InterfacePoeModeLabel'
type InterfacePoeModeLabel string

// List of Interface_poe_mode_label
const (
	INTERFACEPOEMODELABEL_PD  InterfacePoeModeLabel = "PD"
	INTERFACEPOEMODELABEL_PSE InterfacePoeModeLabel = "PSE"
)

// All allowed values of InterfacePoeModeLabel enum
var AllowedInterfacePoeModeLabelEnumValues = []InterfacePoeModeLabel{
	"PD",
	"PSE",
}

func (v *InterfacePoeModeLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfacePoeModeLabel(value)
	for _, existing := range AllowedInterfacePoeModeLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfacePoeModeLabel", value)
}

// NewInterfacePoeModeLabelFromValue returns a pointer to a valid InterfacePoeModeLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfacePoeModeLabelFromValue(v string) (*InterfacePoeModeLabel, error) {
	ev := InterfacePoeModeLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfacePoeModeLabel: valid values are %v", v, AllowedInterfacePoeModeLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfacePoeModeLabel) IsValid() bool {
	for _, existing := range AllowedInterfacePoeModeLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_poe_mode_label value
func (v InterfacePoeModeLabel) Ptr() *InterfacePoeModeLabel {
	return &v
}

type NullableInterfacePoeModeLabel struct {
	value *InterfacePoeModeLabel
	isSet bool
}

func (v NullableInterfacePoeModeLabel) Get() *InterfacePoeModeLabel {
	return v.value
}

func (v *NullableInterfacePoeModeLabel) Set(val *InterfacePoeModeLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfacePoeModeLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfacePoeModeLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfacePoeModeLabel(val *InterfacePoeModeLabel) *NullableInterfacePoeModeLabel {
	return &NullableInterfacePoeModeLabel{value: val, isSet: true}
}

func (v NullableInterfacePoeModeLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfacePoeModeLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
