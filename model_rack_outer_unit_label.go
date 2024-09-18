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

// RackOuterUnitLabel the model 'RackOuterUnitLabel'
type RackOuterUnitLabel string

// List of Rack_outer_unit_label
const (
	RACKOUTERUNITLABEL_MILLIMETERS RackOuterUnitLabel = "Millimeters"
	RACKOUTERUNITLABEL_INCHES      RackOuterUnitLabel = "Inches"
)

// All allowed values of RackOuterUnitLabel enum
var AllowedRackOuterUnitLabelEnumValues = []RackOuterUnitLabel{
	"Millimeters",
	"Inches",
}

func (v *RackOuterUnitLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackOuterUnitLabel(value)
	for _, existing := range AllowedRackOuterUnitLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackOuterUnitLabel", value)
}

// NewRackOuterUnitLabelFromValue returns a pointer to a valid RackOuterUnitLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackOuterUnitLabelFromValue(v string) (*RackOuterUnitLabel, error) {
	ev := RackOuterUnitLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackOuterUnitLabel: valid values are %v", v, AllowedRackOuterUnitLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackOuterUnitLabel) IsValid() bool {
	for _, existing := range AllowedRackOuterUnitLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_outer_unit_label value
func (v RackOuterUnitLabel) Ptr() *RackOuterUnitLabel {
	return &v
}

type NullableRackOuterUnitLabel struct {
	value *RackOuterUnitLabel
	isSet bool
}

func (v NullableRackOuterUnitLabel) Get() *RackOuterUnitLabel {
	return v.value
}

func (v *NullableRackOuterUnitLabel) Set(val *RackOuterUnitLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableRackOuterUnitLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableRackOuterUnitLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackOuterUnitLabel(val *RackOuterUnitLabel) *NullableRackOuterUnitLabel {
	return &NullableRackOuterUnitLabel{value: val, isSet: true}
}

func (v NullableRackOuterUnitLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackOuterUnitLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
