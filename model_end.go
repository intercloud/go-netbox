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

// End the model 'End'
type End string

// List of End
const (
	END_A End = "A"
	END_B End = "B"
)

// All allowed values of End enum
var AllowedEndEnumValues = []End{
	"A",
	"B",
}

func (v *End) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := End(value)
	for _, existing := range AllowedEndEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid End", value)
}

// NewEndFromValue returns a pointer to a valid End
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEndFromValue(v string) (*End, error) {
	ev := End(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for End: valid values are %v", v, AllowedEndEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v End) IsValid() bool {
	for _, existing := range AllowedEndEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to End value
func (v End) Ptr() *End {
	return &v
}

type NullableEnd struct {
	value *End
	isSet bool
}

func (v NullableEnd) Get() *End {
	return v.value
}

func (v *NullableEnd) Set(val *End) {
	v.value = val
	v.isSet = true
}

func (v NullableEnd) IsSet() bool {
	return v.isSet
}

func (v *NullableEnd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnd(val *End) *NullableEnd {
	return &NullableEnd{value: val, isSet: true}
}

func (v NullableEnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
