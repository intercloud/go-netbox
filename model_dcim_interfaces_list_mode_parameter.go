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

// DcimInterfacesListModeParameter the model 'DcimInterfacesListModeParameter'
type DcimInterfacesListModeParameter string

// List of dcim_interfaces_list_mode_parameter
const (
	DCIMINTERFACESLISTMODEPARAMETER_ACCESS     DcimInterfacesListModeParameter = "access"
	DCIMINTERFACESLISTMODEPARAMETER_TAGGED     DcimInterfacesListModeParameter = "tagged"
	DCIMINTERFACESLISTMODEPARAMETER_TAGGED_ALL DcimInterfacesListModeParameter = "tagged-all"
)

// All allowed values of DcimInterfacesListModeParameter enum
var AllowedDcimInterfacesListModeParameterEnumValues = []DcimInterfacesListModeParameter{
	"access",
	"tagged",
	"tagged-all",
}

func (v *DcimInterfacesListModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimInterfacesListModeParameter(value)
	for _, existing := range AllowedDcimInterfacesListModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimInterfacesListModeParameter", value)
}

// NewDcimInterfacesListModeParameterFromValue returns a pointer to a valid DcimInterfacesListModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimInterfacesListModeParameterFromValue(v string) (*DcimInterfacesListModeParameter, error) {
	ev := DcimInterfacesListModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimInterfacesListModeParameter: valid values are %v", v, AllowedDcimInterfacesListModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimInterfacesListModeParameter) IsValid() bool {
	for _, existing := range AllowedDcimInterfacesListModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_interfaces_list_mode_parameter value
func (v DcimInterfacesListModeParameter) Ptr() *DcimInterfacesListModeParameter {
	return &v
}

type NullableDcimInterfacesListModeParameter struct {
	value *DcimInterfacesListModeParameter
	isSet bool
}

func (v NullableDcimInterfacesListModeParameter) Get() *DcimInterfacesListModeParameter {
	return v.value
}

func (v *NullableDcimInterfacesListModeParameter) Set(val *DcimInterfacesListModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimInterfacesListModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimInterfacesListModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimInterfacesListModeParameter(val *DcimInterfacesListModeParameter) *NullableDcimInterfacesListModeParameter {
	return &NullableDcimInterfacesListModeParameter{value: val, isSet: true}
}

func (v NullableDcimInterfacesListModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimInterfacesListModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
