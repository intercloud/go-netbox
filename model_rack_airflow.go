/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.1 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the RackAirflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackAirflow{}

// RackAirflow struct for RackAirflow
type RackAirflow struct {
	Value                *PatchedWritableRackRequestAirflow `json:"value,omitempty"`
	Label                *RackAirflowLabel                  `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackAirflow RackAirflow

// NewRackAirflow instantiates a new RackAirflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackAirflow() *RackAirflow {
	this := RackAirflow{}
	return &this
}

// NewRackAirflowWithDefaults instantiates a new RackAirflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackAirflowWithDefaults() *RackAirflow {
	this := RackAirflow{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RackAirflow) GetValue() PatchedWritableRackRequestAirflow {
	if o == nil || IsNil(o.Value) {
		var ret PatchedWritableRackRequestAirflow
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackAirflow) GetValueOk() (*PatchedWritableRackRequestAirflow, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RackAirflow) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given PatchedWritableRackRequestAirflow and assigns it to the Value field.
func (o *RackAirflow) SetValue(v PatchedWritableRackRequestAirflow) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *RackAirflow) GetLabel() RackAirflowLabel {
	if o == nil || IsNil(o.Label) {
		var ret RackAirflowLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackAirflow) GetLabelOk() (*RackAirflowLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *RackAirflow) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given RackAirflowLabel and assigns it to the Label field.
func (o *RackAirflow) SetLabel(v RackAirflowLabel) {
	o.Label = &v
}

func (o RackAirflow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackAirflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RackAirflow) UnmarshalJSON(data []byte) (err error) {
	varRackAirflow := _RackAirflow{}

	err = json.Unmarshal(data, &varRackAirflow)

	if err != nil {
		return err
	}

	*o = RackAirflow(varRackAirflow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackAirflow struct {
	value *RackAirflow
	isSet bool
}

func (v NullableRackAirflow) Get() *RackAirflow {
	return v.value
}

func (v *NullableRackAirflow) Set(val *RackAirflow) {
	v.value = val
	v.isSet = true
}

func (v NullableRackAirflow) IsSet() bool {
	return v.isSet
}

func (v *NullableRackAirflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackAirflow(val *RackAirflow) *NullableRackAirflow {
	return &NullableRackAirflow{value: val, isSet: true}
}

func (v NullableRackAirflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackAirflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
