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

// checks if the PatchedProviderNetworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedProviderNetworkRequest{}

// PatchedProviderNetworkRequest Adds support for custom fields and tags.
type PatchedProviderNetworkRequest struct {
	Provider             *BriefProviderRequest  `json:"provider,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	ServiceId            *string                `json:"service_id,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedProviderNetworkRequest PatchedProviderNetworkRequest

// NewPatchedProviderNetworkRequest instantiates a new PatchedProviderNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedProviderNetworkRequest() *PatchedProviderNetworkRequest {
	this := PatchedProviderNetworkRequest{}
	return &this
}

// NewPatchedProviderNetworkRequestWithDefaults instantiates a new PatchedProviderNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedProviderNetworkRequestWithDefaults() *PatchedProviderNetworkRequest {
	this := PatchedProviderNetworkRequest{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PatchedProviderNetworkRequest) GetProvider() BriefProviderRequest {
	if o == nil || IsNil(o.Provider) {
		var ret BriefProviderRequest
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProviderNetworkRequest) GetProviderOk() (*BriefProviderRequest, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PatchedProviderNetworkRequest) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given BriefProviderRequest and assigns it to the Provider field.
func (o *PatchedProviderNetworkRequest) SetProvider(v BriefProviderRequest) {
	o.Provider = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedProviderNetworkRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProviderNetworkRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedProviderNetworkRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedProviderNetworkRequest) SetName(v string) {
	o.Name = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *PatchedProviderNetworkRequest) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProviderNetworkRequest) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *PatchedProviderNetworkRequest) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *PatchedProviderNetworkRequest) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedProviderNetworkRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProviderNetworkRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedProviderNetworkRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedProviderNetworkRequest) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedProviderNetworkRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProviderNetworkRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedProviderNetworkRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedProviderNetworkRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedProviderNetworkRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProviderNetworkRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedProviderNetworkRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedProviderNetworkRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedProviderNetworkRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedProviderNetworkRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedProviderNetworkRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedProviderNetworkRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedProviderNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedProviderNetworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
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

func (o *PatchedProviderNetworkRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedProviderNetworkRequest := _PatchedProviderNetworkRequest{}

	err = json.Unmarshal(data, &varPatchedProviderNetworkRequest)

	if err != nil {
		return err
	}

	*o = PatchedProviderNetworkRequest(varPatchedProviderNetworkRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "provider")
		delete(additionalProperties, "name")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedProviderNetworkRequest struct {
	value *PatchedProviderNetworkRequest
	isSet bool
}

func (v NullablePatchedProviderNetworkRequest) Get() *PatchedProviderNetworkRequest {
	return v.value
}

func (v *NullablePatchedProviderNetworkRequest) Set(val *PatchedProviderNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedProviderNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedProviderNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedProviderNetworkRequest(val *PatchedProviderNetworkRequest) *NullablePatchedProviderNetworkRequest {
	return &NullablePatchedProviderNetworkRequest{value: val, isSet: true}
}

func (v NullablePatchedProviderNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedProviderNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
