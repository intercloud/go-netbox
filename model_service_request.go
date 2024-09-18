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

// checks if the ServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceRequest{}

// ServiceRequest Adds support for custom fields and tags.
type ServiceRequest struct {
	Device               NullableBriefDeviceRequest             `json:"device,omitempty"`
	VirtualMachine       NullableBriefVirtualMachineRequest     `json:"virtual_machine,omitempty"`
	Name                 string                                 `json:"name"`
	Protocol             *PatchedWritableServiceRequestProtocol `json:"protocol,omitempty"`
	Ports                []int32                                `json:"ports"`
	Ipaddresses          []int32                                `json:"ipaddresses,omitempty"`
	Description          *string                                `json:"description,omitempty"`
	Comments             *string                                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest                     `json:"tags,omitempty"`
	CustomFields         map[string]interface{}                 `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceRequest ServiceRequest

// NewServiceRequest instantiates a new ServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRequest(name string, ports []int32) *ServiceRequest {
	this := ServiceRequest{}
	this.Name = name
	this.Ports = ports
	return &this
}

// NewServiceRequestWithDefaults instantiates a new ServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRequestWithDefaults() *ServiceRequest {
	this := ServiceRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceRequest) GetDevice() BriefDeviceRequest {
	if o == nil || IsNil(o.Device.Get()) {
		var ret BriefDeviceRequest
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *ServiceRequest) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableBriefDeviceRequest and assigns it to the Device field.
func (o *ServiceRequest) SetDevice(v BriefDeviceRequest) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *ServiceRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *ServiceRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceRequest) GetVirtualMachine() BriefVirtualMachineRequest {
	if o == nil || IsNil(o.VirtualMachine.Get()) {
		var ret BriefVirtualMachineRequest
		return ret
	}
	return *o.VirtualMachine.Get()
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceRequest) GetVirtualMachineOk() (*BriefVirtualMachineRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualMachine.Get(), o.VirtualMachine.IsSet()
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *ServiceRequest) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine.IsSet() {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given NullableBriefVirtualMachineRequest and assigns it to the VirtualMachine field.
func (o *ServiceRequest) SetVirtualMachine(v BriefVirtualMachineRequest) {
	o.VirtualMachine.Set(&v)
}

// SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil
func (o *ServiceRequest) SetVirtualMachineNil() {
	o.VirtualMachine.Set(nil)
}

// UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil
func (o *ServiceRequest) UnsetVirtualMachine() {
	o.VirtualMachine.Unset()
}

// GetName returns the Name field value
func (o *ServiceRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceRequest) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ServiceRequest) GetProtocol() PatchedWritableServiceRequestProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret PatchedWritableServiceRequestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetProtocolOk() (*PatchedWritableServiceRequestProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ServiceRequest) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given PatchedWritableServiceRequestProtocol and assigns it to the Protocol field.
func (o *ServiceRequest) SetProtocol(v PatchedWritableServiceRequestProtocol) {
	o.Protocol = &v
}

// GetPorts returns the Ports field value
func (o *ServiceRequest) GetPorts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetPortsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ports, true
}

// SetPorts sets field value
func (o *ServiceRequest) SetPorts(v []int32) {
	o.Ports = v
}

// GetIpaddresses returns the Ipaddresses field value if set, zero value otherwise.
func (o *ServiceRequest) GetIpaddresses() []int32 {
	if o == nil || IsNil(o.Ipaddresses) {
		var ret []int32
		return ret
	}
	return o.Ipaddresses
}

// GetIpaddressesOk returns a tuple with the Ipaddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetIpaddressesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Ipaddresses) {
		return nil, false
	}
	return o.Ipaddresses, true
}

// HasIpaddresses returns a boolean if a field has been set.
func (o *ServiceRequest) HasIpaddresses() bool {
	if o != nil && !IsNil(o.Ipaddresses) {
		return true
	}

	return false
}

// SetIpaddresses gets a reference to the given []int32 and assigns it to the Ipaddresses field.
func (o *ServiceRequest) SetIpaddresses(v []int32) {
	o.Ipaddresses = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ServiceRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ServiceRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ServiceRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *ServiceRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ServiceRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ServiceRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ServiceRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o ServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.VirtualMachine.IsSet() {
		toSerialize["virtual_machine"] = o.VirtualMachine.Get()
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	toSerialize["ports"] = o.Ports
	if !IsNil(o.Ipaddresses) {
		toSerialize["ipaddresses"] = o.Ipaddresses
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

func (o *ServiceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"ports",
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

	varServiceRequest := _ServiceRequest{}

	err = json.Unmarshal(data, &varServiceRequest)

	if err != nil {
		return err
	}

	*o = ServiceRequest(varServiceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "virtual_machine")
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "ports")
		delete(additionalProperties, "ipaddresses")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceRequest struct {
	value *ServiceRequest
	isSet bool
}

func (v NullableServiceRequest) Get() *ServiceRequest {
	return v.value
}

func (v *NullableServiceRequest) Set(val *ServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRequest(val *ServiceRequest) *NullableServiceRequest {
	return &NullableServiceRequest{value: val, isSet: true}
}

func (v NullableServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
