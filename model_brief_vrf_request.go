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

// checks if the BriefVRFRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefVRFRequest{}

// BriefVRFRequest Adds support for custom fields and tags.
type BriefVRFRequest struct {
	Name string `json:"name"`
	// Unique route distinguisher (as defined in RFC 4364)
	Rd                   NullableString `json:"rd,omitempty"`
	Description          *string        `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefVRFRequest BriefVRFRequest

// NewBriefVRFRequest instantiates a new BriefVRFRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefVRFRequest(name string) *BriefVRFRequest {
	this := BriefVRFRequest{}
	this.Name = name
	return &this
}

// NewBriefVRFRequestWithDefaults instantiates a new BriefVRFRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefVRFRequestWithDefaults() *BriefVRFRequest {
	this := BriefVRFRequest{}
	return &this
}

// GetName returns the Name field value
func (o *BriefVRFRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefVRFRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefVRFRequest) SetName(v string) {
	o.Name = v
}

// GetRd returns the Rd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BriefVRFRequest) GetRd() string {
	if o == nil || IsNil(o.Rd.Get()) {
		var ret string
		return ret
	}
	return *o.Rd.Get()
}

// GetRdOk returns a tuple with the Rd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BriefVRFRequest) GetRdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rd.Get(), o.Rd.IsSet()
}

// HasRd returns a boolean if a field has been set.
func (o *BriefVRFRequest) HasRd() bool {
	if o != nil && o.Rd.IsSet() {
		return true
	}

	return false
}

// SetRd gets a reference to the given NullableString and assigns it to the Rd field.
func (o *BriefVRFRequest) SetRd(v string) {
	o.Rd.Set(&v)
}

// SetRdNil sets the value for Rd to be an explicit nil
func (o *BriefVRFRequest) SetRdNil() {
	o.Rd.Set(nil)
}

// UnsetRd ensures that no value is present for Rd, not even an explicit nil
func (o *BriefVRFRequest) UnsetRd() {
	o.Rd.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefVRFRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefVRFRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefVRFRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefVRFRequest) SetDescription(v string) {
	o.Description = &v
}

func (o BriefVRFRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefVRFRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Rd.IsSet() {
		toSerialize["rd"] = o.Rd.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefVRFRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varBriefVRFRequest := _BriefVRFRequest{}

	err = json.Unmarshal(data, &varBriefVRFRequest)

	if err != nil {
		return err
	}

	*o = BriefVRFRequest(varBriefVRFRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "rd")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefVRFRequest struct {
	value *BriefVRFRequest
	isSet bool
}

func (v NullableBriefVRFRequest) Get() *BriefVRFRequest {
	return v.value
}

func (v *NullableBriefVRFRequest) Set(val *BriefVRFRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefVRFRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefVRFRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefVRFRequest(val *BriefVRFRequest) *NullableBriefVRFRequest {
	return &NullableBriefVRFRequest{value: val, isSet: true}
}

func (v NullableBriefVRFRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefVRFRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
