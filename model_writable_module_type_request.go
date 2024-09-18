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

// checks if the WritableModuleTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WritableModuleTypeRequest{}

// WritableModuleTypeRequest Adds support for custom fields and tags.
type WritableModuleTypeRequest struct {
	Manufacturer BriefManufacturerRequest `json:"manufacturer"`
	Model        string                   `json:"model"`
	// Discrete part number (optional)
	PartNumber           *string                    `json:"part_number,omitempty"`
	Airflow              *ModuleTypeAirflowValue    `json:"airflow,omitempty"`
	Weight               NullableFloat64            `json:"weight,omitempty"`
	WeightUnit           *DeviceTypeWeightUnitValue `json:"weight_unit,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	Comments             *string                    `json:"comments,omitempty"`
	Tags                 []NestedTagRequest         `json:"tags,omitempty"`
	CustomFields         map[string]interface{}     `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WritableModuleTypeRequest WritableModuleTypeRequest

// NewWritableModuleTypeRequest instantiates a new WritableModuleTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWritableModuleTypeRequest(manufacturer BriefManufacturerRequest, model string) *WritableModuleTypeRequest {
	this := WritableModuleTypeRequest{}
	this.Manufacturer = manufacturer
	this.Model = model
	return &this
}

// NewWritableModuleTypeRequestWithDefaults instantiates a new WritableModuleTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWritableModuleTypeRequestWithDefaults() *WritableModuleTypeRequest {
	this := WritableModuleTypeRequest{}
	return &this
}

// GetManufacturer returns the Manufacturer field value
func (o *WritableModuleTypeRequest) GetManufacturer() BriefManufacturerRequest {
	if o == nil {
		var ret BriefManufacturerRequest
		return ret
	}

	return o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetManufacturerOk() (*BriefManufacturerRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Manufacturer, true
}

// SetManufacturer sets field value
func (o *WritableModuleTypeRequest) SetManufacturer(v BriefManufacturerRequest) {
	o.Manufacturer = v
}

// GetModel returns the Model field value
func (o *WritableModuleTypeRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *WritableModuleTypeRequest) SetModel(v string) {
	o.Model = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *WritableModuleTypeRequest) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *WritableModuleTypeRequest) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *WritableModuleTypeRequest) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetAirflow returns the Airflow field value if set, zero value otherwise.
func (o *WritableModuleTypeRequest) GetAirflow() ModuleTypeAirflowValue {
	if o == nil || IsNil(o.Airflow) {
		var ret ModuleTypeAirflowValue
		return ret
	}
	return *o.Airflow
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetAirflowOk() (*ModuleTypeAirflowValue, bool) {
	if o == nil || IsNil(o.Airflow) {
		return nil, false
	}
	return o.Airflow, true
}

// HasAirflow returns a boolean if a field has been set.
func (o *WritableModuleTypeRequest) HasAirflow() bool {
	if o != nil && !IsNil(o.Airflow) {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given ModuleTypeAirflowValue and assigns it to the Airflow field.
func (o *WritableModuleTypeRequest) SetAirflow(v ModuleTypeAirflowValue) {
	o.Airflow = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WritableModuleTypeRequest) GetWeight() float64 {
	if o == nil || IsNil(o.Weight.Get()) {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WritableModuleTypeRequest) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *WritableModuleTypeRequest) HasWeight() bool {
	if o != nil && o.Weight.IsSet() {
		return true
	}

	return false
}

// SetWeight gets a reference to the given NullableFloat64 and assigns it to the Weight field.
func (o *WritableModuleTypeRequest) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil
func (o *WritableModuleTypeRequest) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil
func (o *WritableModuleTypeRequest) UnsetWeight() {
	o.Weight.Unset()
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *WritableModuleTypeRequest) GetWeightUnit() DeviceTypeWeightUnitValue {
	if o == nil || IsNil(o.WeightUnit) {
		var ret DeviceTypeWeightUnitValue
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetWeightUnitOk() (*DeviceTypeWeightUnitValue, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *WritableModuleTypeRequest) HasWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given DeviceTypeWeightUnitValue and assigns it to the WeightUnit field.
func (o *WritableModuleTypeRequest) SetWeightUnit(v DeviceTypeWeightUnitValue) {
	o.WeightUnit = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WritableModuleTypeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WritableModuleTypeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WritableModuleTypeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WritableModuleTypeRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WritableModuleTypeRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WritableModuleTypeRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WritableModuleTypeRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WritableModuleTypeRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WritableModuleTypeRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WritableModuleTypeRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WritableModuleTypeRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WritableModuleTypeRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WritableModuleTypeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o WritableModuleTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WritableModuleTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["manufacturer"] = o.Manufacturer
	toSerialize["model"] = o.Model
	if !IsNil(o.PartNumber) {
		toSerialize["part_number"] = o.PartNumber
	}
	if !IsNil(o.Airflow) {
		toSerialize["airflow"] = o.Airflow
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}
	if !IsNil(o.WeightUnit) {
		toSerialize["weight_unit"] = o.WeightUnit
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WritableModuleTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"manufacturer",
		"model",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWritableModuleTypeRequest := _WritableModuleTypeRequest{}

	err = json.Unmarshal(data, &varWritableModuleTypeRequest)

	if err != nil {
		return err
	}

	*o = WritableModuleTypeRequest(varWritableModuleTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "manufacturer")
		delete(additionalProperties, "model")
		delete(additionalProperties, "part_number")
		delete(additionalProperties, "airflow")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "weight_unit")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWritableModuleTypeRequest struct {
	value *WritableModuleTypeRequest
	isSet bool
}

func (v NullableWritableModuleTypeRequest) Get() *WritableModuleTypeRequest {
	return v.value
}

func (v *NullableWritableModuleTypeRequest) Set(val *WritableModuleTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWritableModuleTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWritableModuleTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWritableModuleTypeRequest(val *WritableModuleTypeRequest) *NullableWritableModuleTypeRequest {
	return &NullableWritableModuleTypeRequest{value: val, isSet: true}
}

func (v NullableWritableModuleTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWritableModuleTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
