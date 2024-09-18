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

// checks if the IPSecPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPSecPolicyRequest{}

// IPSecPolicyRequest Adds support for custom fields and tags.
type IPSecPolicyRequest struct {
	Name                 string                 `json:"name"`
	Description          *string                `json:"description,omitempty"`
	Proposals            []int32                `json:"proposals,omitempty"`
	PfsGroup             *IKEProposalGroupValue `json:"pfs_group,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPSecPolicyRequest IPSecPolicyRequest

// NewIPSecPolicyRequest instantiates a new IPSecPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecPolicyRequest(name string) *IPSecPolicyRequest {
	this := IPSecPolicyRequest{}
	this.Name = name
	return &this
}

// NewIPSecPolicyRequestWithDefaults instantiates a new IPSecPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecPolicyRequestWithDefaults() *IPSecPolicyRequest {
	this := IPSecPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *IPSecPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IPSecPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IPSecPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IPSecPolicyRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IPSecPolicyRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IPSecPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetProposals returns the Proposals field value if set, zero value otherwise.
func (o *IPSecPolicyRequest) GetProposals() []int32 {
	if o == nil || IsNil(o.Proposals) {
		var ret []int32
		return ret
	}
	return o.Proposals
}

// GetProposalsOk returns a tuple with the Proposals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicyRequest) GetProposalsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Proposals) {
		return nil, false
	}
	return o.Proposals, true
}

// HasProposals returns a boolean if a field has been set.
func (o *IPSecPolicyRequest) HasProposals() bool {
	if o != nil && !IsNil(o.Proposals) {
		return true
	}

	return false
}

// SetProposals gets a reference to the given []int32 and assigns it to the Proposals field.
func (o *IPSecPolicyRequest) SetProposals(v []int32) {
	o.Proposals = v
}

// GetPfsGroup returns the PfsGroup field value if set, zero value otherwise.
func (o *IPSecPolicyRequest) GetPfsGroup() IKEProposalGroupValue {
	if o == nil || IsNil(o.PfsGroup) {
		var ret IKEProposalGroupValue
		return ret
	}
	return *o.PfsGroup
}

// GetPfsGroupOk returns a tuple with the PfsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicyRequest) GetPfsGroupOk() (*IKEProposalGroupValue, bool) {
	if o == nil || IsNil(o.PfsGroup) {
		return nil, false
	}
	return o.PfsGroup, true
}

// HasPfsGroup returns a boolean if a field has been set.
func (o *IPSecPolicyRequest) HasPfsGroup() bool {
	if o != nil && !IsNil(o.PfsGroup) {
		return true
	}

	return false
}

// SetPfsGroup gets a reference to the given IKEProposalGroupValue and assigns it to the PfsGroup field.
func (o *IPSecPolicyRequest) SetPfsGroup(v IKEProposalGroupValue) {
	o.PfsGroup = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPSecPolicyRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicyRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPSecPolicyRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPSecPolicyRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPSecPolicyRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicyRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPSecPolicyRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *IPSecPolicyRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IPSecPolicyRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecPolicyRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IPSecPolicyRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IPSecPolicyRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o IPSecPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPSecPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Proposals) {
		toSerialize["proposals"] = o.Proposals
	}
	if !IsNil(o.PfsGroup) {
		toSerialize["pfs_group"] = o.PfsGroup
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

func (o *IPSecPolicyRequest) UnmarshalJSON(data []byte) (err error) {
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

	varIPSecPolicyRequest := _IPSecPolicyRequest{}

	err = json.Unmarshal(data, &varIPSecPolicyRequest)

	if err != nil {
		return err
	}

	*o = IPSecPolicyRequest(varIPSecPolicyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "proposals")
		delete(additionalProperties, "pfs_group")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPSecPolicyRequest struct {
	value *IPSecPolicyRequest
	isSet bool
}

func (v NullableIPSecPolicyRequest) Get() *IPSecPolicyRequest {
	return v.value
}

func (v *NullableIPSecPolicyRequest) Set(val *IPSecPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecPolicyRequest(val *IPSecPolicyRequest) *NullableIPSecPolicyRequest {
	return &NullableIPSecPolicyRequest{value: val, isSet: true}
}

func (v NullableIPSecPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
