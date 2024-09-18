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

// RackRequestFormFactor * `2-post-frame` - 2-post frame * `4-post-frame` - 4-post frame * `4-post-cabinet` - 4-post cabinet * `wall-frame` - Wall-mounted frame * `wall-frame-vertical` - Wall-mounted frame (vertical) * `wall-cabinet` - Wall-mounted cabinet * `wall-cabinet-vertical` - Wall-mounted cabinet (vertical)
type RackRequestFormFactor string

// List of RackRequest_form_factor
const (
	RACKREQUESTFORMFACTOR__2_POST_FRAME         RackRequestFormFactor = "2-post-frame"
	RACKREQUESTFORMFACTOR__4_POST_FRAME         RackRequestFormFactor = "4-post-frame"
	RACKREQUESTFORMFACTOR__4_POST_CABINET       RackRequestFormFactor = "4-post-cabinet"
	RACKREQUESTFORMFACTOR_WALL_FRAME            RackRequestFormFactor = "wall-frame"
	RACKREQUESTFORMFACTOR_WALL_FRAME_VERTICAL   RackRequestFormFactor = "wall-frame-vertical"
	RACKREQUESTFORMFACTOR_WALL_CABINET          RackRequestFormFactor = "wall-cabinet"
	RACKREQUESTFORMFACTOR_WALL_CABINET_VERTICAL RackRequestFormFactor = "wall-cabinet-vertical"
	RACKREQUESTFORMFACTOR_EMPTY                 RackRequestFormFactor = ""
)

// All allowed values of RackRequestFormFactor enum
var AllowedRackRequestFormFactorEnumValues = []RackRequestFormFactor{
	"2-post-frame",
	"4-post-frame",
	"4-post-cabinet",
	"wall-frame",
	"wall-frame-vertical",
	"wall-cabinet",
	"wall-cabinet-vertical",
	"",
}

func (v *RackRequestFormFactor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackRequestFormFactor(value)
	for _, existing := range AllowedRackRequestFormFactorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackRequestFormFactor", value)
}

// NewRackRequestFormFactorFromValue returns a pointer to a valid RackRequestFormFactor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackRequestFormFactorFromValue(v string) (*RackRequestFormFactor, error) {
	ev := RackRequestFormFactor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackRequestFormFactor: valid values are %v", v, AllowedRackRequestFormFactorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackRequestFormFactor) IsValid() bool {
	for _, existing := range AllowedRackRequestFormFactorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RackRequest_form_factor value
func (v RackRequestFormFactor) Ptr() *RackRequestFormFactor {
	return &v
}

type NullableRackRequestFormFactor struct {
	value *RackRequestFormFactor
	isSet bool
}

func (v NullableRackRequestFormFactor) Get() *RackRequestFormFactor {
	return v.value
}

func (v *NullableRackRequestFormFactor) Set(val *RackRequestFormFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableRackRequestFormFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableRackRequestFormFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackRequestFormFactor(val *RackRequestFormFactor) *NullableRackRequestFormFactor {
	return &NullableRackRequestFormFactor{value: val, isSet: true}
}

func (v NullableRackRequestFormFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackRequestFormFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
