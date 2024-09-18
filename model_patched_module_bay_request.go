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

// checks if the PatchedModuleBayRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedModuleBayRequest{}

// PatchedModuleBayRequest Adds support for custom fields and tags.
type PatchedModuleBayRequest struct {
	Device          *BriefDeviceRequest        `json:"device,omitempty"`
	Module          NullableBriefModuleRequest `json:"module,omitempty"`
	Name            *string                    `json:"name,omitempty"`
	InstalledModule NullableBriefModuleRequest `json:"installed_module,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	// Identifier to reference when renaming installed components
	Position             *string                `json:"position,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedModuleBayRequest PatchedModuleBayRequest

// NewPatchedModuleBayRequest instantiates a new PatchedModuleBayRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedModuleBayRequest() *PatchedModuleBayRequest {
	this := PatchedModuleBayRequest{}
	return &this
}

// NewPatchedModuleBayRequestWithDefaults instantiates a new PatchedModuleBayRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedModuleBayRequestWithDefaults() *PatchedModuleBayRequest {
	this := PatchedModuleBayRequest{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PatchedModuleBayRequest) GetDevice() BriefDeviceRequest {
	if o == nil || IsNil(o.Device) {
		var ret BriefDeviceRequest
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedModuleBayRequest) GetDeviceOk() (*BriefDeviceRequest, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given BriefDeviceRequest and assigns it to the Device field.
func (o *PatchedModuleBayRequest) SetDevice(v BriefDeviceRequest) {
	o.Device = &v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedModuleBayRequest) GetModule() BriefModuleRequest {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BriefModuleRequest
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedModuleBayRequest) GetModuleOk() (*BriefModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBriefModuleRequest and assigns it to the Module field.
func (o *PatchedModuleBayRequest) SetModule(v BriefModuleRequest) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *PatchedModuleBayRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *PatchedModuleBayRequest) UnsetModule() {
	o.Module.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedModuleBayRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedModuleBayRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedModuleBayRequest) SetName(v string) {
	o.Name = &v
}

// GetInstalledModule returns the InstalledModule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedModuleBayRequest) GetInstalledModule() BriefModuleRequest {
	if o == nil || IsNil(o.InstalledModule.Get()) {
		var ret BriefModuleRequest
		return ret
	}
	return *o.InstalledModule.Get()
}

// GetInstalledModuleOk returns a tuple with the InstalledModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedModuleBayRequest) GetInstalledModuleOk() (*BriefModuleRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstalledModule.Get(), o.InstalledModule.IsSet()
}

// HasInstalledModule returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasInstalledModule() bool {
	if o != nil && o.InstalledModule.IsSet() {
		return true
	}

	return false
}

// SetInstalledModule gets a reference to the given NullableBriefModuleRequest and assigns it to the InstalledModule field.
func (o *PatchedModuleBayRequest) SetInstalledModule(v BriefModuleRequest) {
	o.InstalledModule.Set(&v)
}

// SetInstalledModuleNil sets the value for InstalledModule to be an explicit nil
func (o *PatchedModuleBayRequest) SetInstalledModuleNil() {
	o.InstalledModule.Set(nil)
}

// UnsetInstalledModule ensures that no value is present for InstalledModule, not even an explicit nil
func (o *PatchedModuleBayRequest) UnsetInstalledModule() {
	o.InstalledModule.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedModuleBayRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedModuleBayRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedModuleBayRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PatchedModuleBayRequest) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedModuleBayRequest) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *PatchedModuleBayRequest) SetPosition(v string) {
	o.Position = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedModuleBayRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedModuleBayRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedModuleBayRequest) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedModuleBayRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedModuleBayRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedModuleBayRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedModuleBayRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedModuleBayRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedModuleBayRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedModuleBayRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedModuleBayRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedModuleBayRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.InstalledModule.IsSet() {
		toSerialize["installed_module"] = o.InstalledModule.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *PatchedModuleBayRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedModuleBayRequest := _PatchedModuleBayRequest{}

	err = json.Unmarshal(data, &varPatchedModuleBayRequest)

	if err != nil {
		return err
	}

	*o = PatchedModuleBayRequest(varPatchedModuleBayRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "installed_module")
		delete(additionalProperties, "label")
		delete(additionalProperties, "position")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedModuleBayRequest struct {
	value *PatchedModuleBayRequest
	isSet bool
}

func (v NullablePatchedModuleBayRequest) Get() *PatchedModuleBayRequest {
	return v.value
}

func (v *NullablePatchedModuleBayRequest) Set(val *PatchedModuleBayRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedModuleBayRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedModuleBayRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedModuleBayRequest(val *PatchedModuleBayRequest) *NullablePatchedModuleBayRequest {
	return &NullablePatchedModuleBayRequest{value: val, isSet: true}
}

func (v NullablePatchedModuleBayRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedModuleBayRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
