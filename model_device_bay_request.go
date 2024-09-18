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

// checks if the DeviceBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceBayRequest{}

// DeviceBayRequest Adds support for custom fields and tags.
type DeviceBayRequest struct {
	Device BriefDeviceRequest `json:"device"`
	Name   string             `json:"name"`
	// Physical label
	Label                *string                    `json:"label,omitempty"`
	Description          *string                    `json:"description,omitempty"`
	InstalledDevice      NullableBriefDeviceRequest `json:"installed_device,omitempty"`
	Tags                 []NestedTagRequest         `json:"tags,omitempty"`
	CustomFields         map[string]interface{}     `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceBayRequest DeviceBayRequest

// NewDeviceBayRequest instantiates a new DeviceBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceBayRequest(device BriefDeviceRequest, name string) *DeviceBayRequest {
	this := DeviceBayRequest{}
	this.Device = device
	this.Name = name
	return &this
}

// NewDeviceBayRequestWithDefaults instantiates a new DeviceBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceBayRequestWithDefaults() *DeviceBayRequest {
	this := DeviceBayRequest{}
	return &this
}

// GetDevice returns the Device field value
func (o *DeviceBayRequest) GetDevice() BriefDeviceRequest {
	if o == nil {
		var ret BriefDeviceRequest
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *DeviceBayRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *DeviceBayRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = v
}

// GetName returns the Name field value
func (o *DeviceBayRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceBayRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceBayRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DeviceBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DeviceBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DeviceBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetInstalledDevice returns the InstalledDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceBayRequest) GetInstalledDevice() BriefDeviceRequest {
	if o == nil || IsNil(o.InstalledDevice.Get()) {
		var ret BriefDeviceRequest
		return ret
	}
	return *o.InstalledDevice.Get()
}

// GetInstalledDeviceOk returns a tuple with the InstalledDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceBayRequest) GetInstalledDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledDevice.Get(), o.InstalledDevice.IsSet()
}

// HasInstalledDevice returns a boolean if a field has been set.
func (o *DeviceBayRequest) HasInstalledDevice() bool {
	if o != nil && o.InstalledDevice.IsSet() {
		return true
	}

	return false
}

// SetInstalledDevice gets a reference to the given NullableBriefDeviceRequest and assigns it to the InstalledDevice field.
func (o *DeviceBayRequest) SetInstalledDevice(v BriefDeviceRequest) {
	o.InstalledDevice.Set(&v)
}

// SetInstalledDeviceNil sets the value for InstalledDevice to be an explicit nil
func (o *DeviceBayRequest) SetInstalledDeviceNil() {
	o.InstalledDevice.Set(nil)
}

// UnsetInstalledDevice ensures that no value is present for InstalledDevice, not even an explicit nil
func (o *DeviceBayRequest) UnsetInstalledDevice() {
	o.InstalledDevice.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceBayRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBayRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *DeviceBayRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *DeviceBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *DeviceBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *DeviceBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o DeviceBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["device"] = o.Device
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.InstalledDevice.IsSet() {
		toSerialize["installed_device"] = o.InstalledDevice.Get()
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

func (o *DeviceBayRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"device",
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

	varDeviceBayRequest := _DeviceBayRequest{}

	err = json.Unmarshal(data, &varDeviceBayRequest)

	if err != nil {
		return err
	}

	*o = DeviceBayRequest(varDeviceBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "installed_device")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceBayRequest struct {
	value *DeviceBayRequest
	isSet bool
}

func (v NullableDeviceBayRequest) Get() *DeviceBayRequest {
	return v.value
}

func (v *NullableDeviceBayRequest) Set(val *DeviceBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceBayRequest(val *DeviceBayRequest) *NullableDeviceBayRequest {
	return &NullableDeviceBayRequest{value: val, isSet: true}
}

func (v NullableDeviceBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
