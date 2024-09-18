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

// CableLengthUnitValue * `km` - Kilometers * `m` - Meters * `cm` - Centimeters * `mi` - Miles * `ft` - Feet * `in` - Inches
type CableLengthUnitValue string

// List of Cable_length_unit_value
const (
	CABLELENGTHUNITVALUE_KM    CableLengthUnitValue = "km"
	CABLELENGTHUNITVALUE_M     CableLengthUnitValue = "m"
	CABLELENGTHUNITVALUE_CM    CableLengthUnitValue = "cm"
	CABLELENGTHUNITVALUE_MI    CableLengthUnitValue = "mi"
	CABLELENGTHUNITVALUE_FT    CableLengthUnitValue = "ft"
	CABLELENGTHUNITVALUE_IN    CableLengthUnitValue = "in"
	CABLELENGTHUNITVALUE_EMPTY CableLengthUnitValue = ""
)

// All allowed values of CableLengthUnitValue enum
var AllowedCableLengthUnitValueEnumValues = []CableLengthUnitValue{
	"km",
	"m",
	"cm",
	"mi",
	"ft",
	"in",
	"",
}

func (v *CableLengthUnitValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CableLengthUnitValue(value)
	for _, existing := range AllowedCableLengthUnitValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CableLengthUnitValue", value)
}

// NewCableLengthUnitValueFromValue returns a pointer to a valid CableLengthUnitValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCableLengthUnitValueFromValue(v string) (*CableLengthUnitValue, error) {
	ev := CableLengthUnitValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CableLengthUnitValue: valid values are %v", v, AllowedCableLengthUnitValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CableLengthUnitValue) IsValid() bool {
	for _, existing := range AllowedCableLengthUnitValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cable_length_unit_value value
func (v CableLengthUnitValue) Ptr() *CableLengthUnitValue {
	return &v
}

type NullableCableLengthUnitValue struct {
	value *CableLengthUnitValue
	isSet bool
}

func (v NullableCableLengthUnitValue) Get() *CableLengthUnitValue {
	return v.value
}

func (v *NullableCableLengthUnitValue) Set(val *CableLengthUnitValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCableLengthUnitValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCableLengthUnitValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCableLengthUnitValue(val *CableLengthUnitValue) *NullableCableLengthUnitValue {
	return &NullableCableLengthUnitValue{value: val, isSet: true}
}

func (v NullableCableLengthUnitValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCableLengthUnitValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
