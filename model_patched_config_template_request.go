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

// checks if the PatchedConfigTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedConfigTemplateRequest{}

// PatchedConfigTemplateRequest Introduces support for Tag assignment. Adds `tags` serialization, and handles tag assignment on create() and update().
type PatchedConfigTemplateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Any <a href=\"https://jinja.palletsprojects.com/en/3.1.x/api/#jinja2.Environment\">additional parameters</a> to pass when constructing the Jinja2 environment.
	EnvironmentParams interface{} `json:"environment_params,omitempty"`
	// Jinja2 template code.
	TemplateCode         *string                 `json:"template_code,omitempty"`
	DataSource           *BriefDataSourceRequest `json:"data_source,omitempty"`
	Tags                 []NestedTagRequest      `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedConfigTemplateRequest PatchedConfigTemplateRequest

// NewPatchedConfigTemplateRequest instantiates a new PatchedConfigTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedConfigTemplateRequest() *PatchedConfigTemplateRequest {
	this := PatchedConfigTemplateRequest{}
	return &this
}

// NewPatchedConfigTemplateRequestWithDefaults instantiates a new PatchedConfigTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedConfigTemplateRequestWithDefaults() *PatchedConfigTemplateRequest {
	this := PatchedConfigTemplateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedConfigTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedConfigTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedConfigTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedConfigTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedConfigTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedConfigTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironmentParams returns the EnvironmentParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedConfigTemplateRequest) GetEnvironmentParams() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnvironmentParams
}

// GetEnvironmentParamsOk returns a tuple with the EnvironmentParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedConfigTemplateRequest) GetEnvironmentParamsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnvironmentParams) {
		return nil, false
	}
	return &o.EnvironmentParams, true
}

// HasEnvironmentParams returns a boolean if a field has been set.
func (o *PatchedConfigTemplateRequest) HasEnvironmentParams() bool {
	if o != nil && !IsNil(o.EnvironmentParams) {
		return true
	}

	return false
}

// SetEnvironmentParams gets a reference to the given interface{} and assigns it to the EnvironmentParams field.
func (o *PatchedConfigTemplateRequest) SetEnvironmentParams(v interface{}) {
	o.EnvironmentParams = v
}

// GetTemplateCode returns the TemplateCode field value if set, zero value otherwise.
func (o *PatchedConfigTemplateRequest) GetTemplateCode() string {
	if o == nil || IsNil(o.TemplateCode) {
		var ret string
		return ret
	}
	return *o.TemplateCode
}

// GetTemplateCodeOk returns a tuple with the TemplateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigTemplateRequest) GetTemplateCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateCode) {
		return nil, false
	}
	return o.TemplateCode, true
}

// HasTemplateCode returns a boolean if a field has been set.
func (o *PatchedConfigTemplateRequest) HasTemplateCode() bool {
	if o != nil && !IsNil(o.TemplateCode) {
		return true
	}

	return false
}

// SetTemplateCode gets a reference to the given string and assigns it to the TemplateCode field.
func (o *PatchedConfigTemplateRequest) SetTemplateCode(v string) {
	o.TemplateCode = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *PatchedConfigTemplateRequest) GetDataSource() BriefDataSourceRequest {
	if o == nil || IsNil(o.DataSource) {
		var ret BriefDataSourceRequest
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigTemplateRequest) GetDataSourceOk() (*BriefDataSourceRequest, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *PatchedConfigTemplateRequest) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given BriefDataSourceRequest and assigns it to the DataSource field.
func (o *PatchedConfigTemplateRequest) SetDataSource(v BriefDataSourceRequest) {
	o.DataSource = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedConfigTemplateRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedConfigTemplateRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedConfigTemplateRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedConfigTemplateRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o PatchedConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedConfigTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.EnvironmentParams != nil {
		toSerialize["environment_params"] = o.EnvironmentParams
	}
	if !IsNil(o.TemplateCode) {
		toSerialize["template_code"] = o.TemplateCode
	}
	if !IsNil(o.DataSource) {
		toSerialize["data_source"] = o.DataSource
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedConfigTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedConfigTemplateRequest := _PatchedConfigTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedConfigTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedConfigTemplateRequest(varPatchedConfigTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "environment_params")
		delete(additionalProperties, "template_code")
		delete(additionalProperties, "data_source")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedConfigTemplateRequest struct {
	value *PatchedConfigTemplateRequest
	isSet bool
}

func (v NullablePatchedConfigTemplateRequest) Get() *PatchedConfigTemplateRequest {
	return v.value
}

func (v *NullablePatchedConfigTemplateRequest) Set(val *PatchedConfigTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedConfigTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedConfigTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedConfigTemplateRequest(val *PatchedConfigTemplateRequest) *NullablePatchedConfigTemplateRequest {
	return &NullablePatchedConfigTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedConfigTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedConfigTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
