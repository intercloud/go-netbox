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

// checks if the CustomFieldType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldType{}

// CustomFieldType struct for CustomFieldType
type CustomFieldType struct {
	Value                *CustomFieldTypeValue `json:"value,omitempty"`
	Label                *CustomFieldTypeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomFieldType CustomFieldType

// NewCustomFieldType instantiates a new CustomFieldType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldType() *CustomFieldType {
	this := CustomFieldType{}
	return &this
}

// NewCustomFieldTypeWithDefaults instantiates a new CustomFieldType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldTypeWithDefaults() *CustomFieldType {
	this := CustomFieldType{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomFieldType) GetValue() CustomFieldTypeValue {
	if o == nil || IsNil(o.Value) {
		var ret CustomFieldTypeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldType) GetValueOk() (*CustomFieldTypeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomFieldType) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CustomFieldTypeValue and assigns it to the Value field.
func (o *CustomFieldType) SetValue(v CustomFieldTypeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CustomFieldType) GetLabel() CustomFieldTypeLabel {
	if o == nil || IsNil(o.Label) {
		var ret CustomFieldTypeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldType) GetLabelOk() (*CustomFieldTypeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CustomFieldType) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given CustomFieldTypeLabel and assigns it to the Label field.
func (o *CustomFieldType) SetLabel(v CustomFieldTypeLabel) {
	o.Label = &v
}

func (o CustomFieldType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldType) ToMap() (map[string]interface{}, error) {
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

func (o *CustomFieldType) UnmarshalJSON(data []byte) (err error) {
	varCustomFieldType := _CustomFieldType{}

	err = json.Unmarshal(data, &varCustomFieldType)

	if err != nil {
		return err
	}

	*o = CustomFieldType(varCustomFieldType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomFieldType struct {
	value *CustomFieldType
	isSet bool
}

func (v NullableCustomFieldType) Get() *CustomFieldType {
	return v.value
}

func (v *NullableCustomFieldType) Set(val *CustomFieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldType(val *CustomFieldType) *NullableCustomFieldType {
	return &NullableCustomFieldType{value: val, isSet: true}
}

func (v NullableCustomFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
