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

// DeviceTypeRequestSubdeviceRole * `parent` - Parent * `child` - Child
type DeviceTypeRequestSubdeviceRole string

// List of DeviceTypeRequest_subdevice_role
const (
	DEVICETYPEREQUESTSUBDEVICEROLE_PARENT DeviceTypeRequestSubdeviceRole = "parent"
	DEVICETYPEREQUESTSUBDEVICEROLE_CHILD  DeviceTypeRequestSubdeviceRole = "child"
	DEVICETYPEREQUESTSUBDEVICEROLE_EMPTY  DeviceTypeRequestSubdeviceRole = ""
)

// All allowed values of DeviceTypeRequestSubdeviceRole enum
var AllowedDeviceTypeRequestSubdeviceRoleEnumValues = []DeviceTypeRequestSubdeviceRole{
	"parent",
	"child",
	"",
}

func (v *DeviceTypeRequestSubdeviceRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeviceTypeRequestSubdeviceRole(value)
	for _, existing := range AllowedDeviceTypeRequestSubdeviceRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeviceTypeRequestSubdeviceRole", value)
}

// NewDeviceTypeRequestSubdeviceRoleFromValue returns a pointer to a valid DeviceTypeRequestSubdeviceRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeviceTypeRequestSubdeviceRoleFromValue(v string) (*DeviceTypeRequestSubdeviceRole, error) {
	ev := DeviceTypeRequestSubdeviceRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeviceTypeRequestSubdeviceRole: valid values are %v", v, AllowedDeviceTypeRequestSubdeviceRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeviceTypeRequestSubdeviceRole) IsValid() bool {
	for _, existing := range AllowedDeviceTypeRequestSubdeviceRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeviceTypeRequest_subdevice_role value
func (v DeviceTypeRequestSubdeviceRole) Ptr() *DeviceTypeRequestSubdeviceRole {
	return &v
}

type NullableDeviceTypeRequestSubdeviceRole struct {
	value *DeviceTypeRequestSubdeviceRole
	isSet bool
}

func (v NullableDeviceTypeRequestSubdeviceRole) Get() *DeviceTypeRequestSubdeviceRole {
	return v.value
}

func (v *NullableDeviceTypeRequestSubdeviceRole) Set(val *DeviceTypeRequestSubdeviceRole) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceTypeRequestSubdeviceRole) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceTypeRequestSubdeviceRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceTypeRequestSubdeviceRole(val *DeviceTypeRequestSubdeviceRole) *NullableDeviceTypeRequestSubdeviceRole {
	return &NullableDeviceTypeRequestSubdeviceRole{value: val, isSet: true}
}

func (v NullableDeviceTypeRequestSubdeviceRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceTypeRequestSubdeviceRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
