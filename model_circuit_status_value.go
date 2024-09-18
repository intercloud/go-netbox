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

// CircuitStatusValue * `planned` - Planned * `provisioning` - Provisioning * `active` - Active * `offline` - Offline * `deprovisioning` - Deprovisioning * `decommissioned` - Decommissioned
type CircuitStatusValue string

// List of Circuit_status_value
const (
	CIRCUITSTATUSVALUE_PLANNED        CircuitStatusValue = "planned"
	CIRCUITSTATUSVALUE_PROVISIONING   CircuitStatusValue = "provisioning"
	CIRCUITSTATUSVALUE_ACTIVE         CircuitStatusValue = "active"
	CIRCUITSTATUSVALUE_OFFLINE        CircuitStatusValue = "offline"
	CIRCUITSTATUSVALUE_DEPROVISIONING CircuitStatusValue = "deprovisioning"
	CIRCUITSTATUSVALUE_DECOMMISSIONED CircuitStatusValue = "decommissioned"
)

// All allowed values of CircuitStatusValue enum
var AllowedCircuitStatusValueEnumValues = []CircuitStatusValue{
	"planned",
	"provisioning",
	"active",
	"offline",
	"deprovisioning",
	"decommissioned",
}

func (v *CircuitStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CircuitStatusValue(value)
	for _, existing := range AllowedCircuitStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CircuitStatusValue", value)
}

// NewCircuitStatusValueFromValue returns a pointer to a valid CircuitStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCircuitStatusValueFromValue(v string) (*CircuitStatusValue, error) {
	ev := CircuitStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CircuitStatusValue: valid values are %v", v, AllowedCircuitStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CircuitStatusValue) IsValid() bool {
	for _, existing := range AllowedCircuitStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Circuit_status_value value
func (v CircuitStatusValue) Ptr() *CircuitStatusValue {
	return &v
}

type NullableCircuitStatusValue struct {
	value *CircuitStatusValue
	isSet bool
}

func (v NullableCircuitStatusValue) Get() *CircuitStatusValue {
	return v.value
}

func (v *NullableCircuitStatusValue) Set(val *CircuitStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitStatusValue(val *CircuitStatusValue) *NullableCircuitStatusValue {
	return &NullableCircuitStatusValue{value: val, isSet: true}
}

func (v NullableCircuitStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
