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

// checks if the NestedWirelessLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedWirelessLinkRequest{}

// NestedWirelessLinkRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedWirelessLinkRequest struct {
	Ssid                 *string `json:"ssid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedWirelessLinkRequest NestedWirelessLinkRequest

// NewNestedWirelessLinkRequest instantiates a new NestedWirelessLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedWirelessLinkRequest() *NestedWirelessLinkRequest {
	this := NestedWirelessLinkRequest{}
	return &this
}

// NewNestedWirelessLinkRequestWithDefaults instantiates a new NestedWirelessLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedWirelessLinkRequestWithDefaults() *NestedWirelessLinkRequest {
	this := NestedWirelessLinkRequest{}
	return &this
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *NestedWirelessLinkRequest) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedWirelessLinkRequest) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *NestedWirelessLinkRequest) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *NestedWirelessLinkRequest) SetSsid(v string) {
	o.Ssid = &v
}

func (o NestedWirelessLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedWirelessLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedWirelessLinkRequest) UnmarshalJSON(data []byte) (err error) {
	varNestedWirelessLinkRequest := _NestedWirelessLinkRequest{}

	err = json.Unmarshal(data, &varNestedWirelessLinkRequest)

	if err != nil {
		return err
	}

	*o = NestedWirelessLinkRequest(varNestedWirelessLinkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ssid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedWirelessLinkRequest struct {
	value *NestedWirelessLinkRequest
	isSet bool
}

func (v NullableNestedWirelessLinkRequest) Get() *NestedWirelessLinkRequest {
	return v.value
}

func (v *NullableNestedWirelessLinkRequest) Set(val *NestedWirelessLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedWirelessLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedWirelessLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedWirelessLinkRequest(val *NestedWirelessLinkRequest) *NullableNestedWirelessLinkRequest {
	return &NullableNestedWirelessLinkRequest{value: val, isSet: true}
}

func (v NullableNestedWirelessLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedWirelessLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
