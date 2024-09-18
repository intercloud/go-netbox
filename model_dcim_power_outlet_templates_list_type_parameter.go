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

// DcimPowerOutletTemplatesListTypeParameter the model 'DcimPowerOutletTemplatesListTypeParameter'
type DcimPowerOutletTemplatesListTypeParameter string

// List of dcim_power_outlet_templates_list_type_parameter
const (
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_CALIFORNIA_STYLE  DcimPowerOutletTemplatesListTypeParameter = "California Style"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_DC                DcimPowerOutletTemplatesListTypeParameter = "DC"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_IEC_60309         DcimPowerOutletTemplatesListTypeParameter = "IEC 60309"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_IEC_60320         DcimPowerOutletTemplatesListTypeParameter = "IEC 60320"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_IEC_60906_1       DcimPowerOutletTemplatesListTypeParameter = "IEC 60906-1"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_ITA_INTERNATIONAL DcimPowerOutletTemplatesListTypeParameter = "ITA/International"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_MOLEX             DcimPowerOutletTemplatesListTypeParameter = "Molex"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_NEMA__LOCKING     DcimPowerOutletTemplatesListTypeParameter = "NEMA (Locking)"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_NEMA__NON_LOCKING DcimPowerOutletTemplatesListTypeParameter = "NEMA (Non-locking)"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_OTHER             DcimPowerOutletTemplatesListTypeParameter = "Other"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_PROPRIETARY       DcimPowerOutletTemplatesListTypeParameter = "Proprietary"
	DCIMPOWEROUTLETTEMPLATESLISTTYPEPARAMETER_USB               DcimPowerOutletTemplatesListTypeParameter = "USB"
)

// All allowed values of DcimPowerOutletTemplatesListTypeParameter enum
var AllowedDcimPowerOutletTemplatesListTypeParameterEnumValues = []DcimPowerOutletTemplatesListTypeParameter{
	"California Style",
	"DC",
	"IEC 60309",
	"IEC 60320",
	"IEC 60906-1",
	"ITA/International",
	"Molex",
	"NEMA (Locking)",
	"NEMA (Non-locking)",
	"Other",
	"Proprietary",
	"USB",
}

func (v *DcimPowerOutletTemplatesListTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DcimPowerOutletTemplatesListTypeParameter(value)
	for _, existing := range AllowedDcimPowerOutletTemplatesListTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DcimPowerOutletTemplatesListTypeParameter", value)
}

// NewDcimPowerOutletTemplatesListTypeParameterFromValue returns a pointer to a valid DcimPowerOutletTemplatesListTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDcimPowerOutletTemplatesListTypeParameterFromValue(v string) (*DcimPowerOutletTemplatesListTypeParameter, error) {
	ev := DcimPowerOutletTemplatesListTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DcimPowerOutletTemplatesListTypeParameter: valid values are %v", v, AllowedDcimPowerOutletTemplatesListTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DcimPowerOutletTemplatesListTypeParameter) IsValid() bool {
	for _, existing := range AllowedDcimPowerOutletTemplatesListTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dcim_power_outlet_templates_list_type_parameter value
func (v DcimPowerOutletTemplatesListTypeParameter) Ptr() *DcimPowerOutletTemplatesListTypeParameter {
	return &v
}

type NullableDcimPowerOutletTemplatesListTypeParameter struct {
	value *DcimPowerOutletTemplatesListTypeParameter
	isSet bool
}

func (v NullableDcimPowerOutletTemplatesListTypeParameter) Get() *DcimPowerOutletTemplatesListTypeParameter {
	return v.value
}

func (v *NullableDcimPowerOutletTemplatesListTypeParameter) Set(val *DcimPowerOutletTemplatesListTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableDcimPowerOutletTemplatesListTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableDcimPowerOutletTemplatesListTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDcimPowerOutletTemplatesListTypeParameter(val *DcimPowerOutletTemplatesListTypeParameter) *NullableDcimPowerOutletTemplatesListTypeParameter {
	return &NullableDcimPowerOutletTemplatesListTypeParameter{value: val, isSet: true}
}

func (v NullableDcimPowerOutletTemplatesListTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDcimPowerOutletTemplatesListTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
