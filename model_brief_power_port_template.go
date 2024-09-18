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

// checks if the BriefPowerPortTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefPowerPortTemplate{}

// BriefPowerPortTemplate Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BriefPowerPortTemplate struct {
	Id      int32  `json:"id"`
	Url     string `json:"url"`
	Display string `json:"display"`
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BriefPowerPortTemplate BriefPowerPortTemplate

// NewBriefPowerPortTemplate instantiates a new BriefPowerPortTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefPowerPortTemplate(id int32, url string, display string, name string) *BriefPowerPortTemplate {
	this := BriefPowerPortTemplate{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	return &this
}

// NewBriefPowerPortTemplateWithDefaults instantiates a new BriefPowerPortTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefPowerPortTemplateWithDefaults() *BriefPowerPortTemplate {
	this := BriefPowerPortTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *BriefPowerPortTemplate) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPortTemplate) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BriefPowerPortTemplate) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *BriefPowerPortTemplate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPortTemplate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BriefPowerPortTemplate) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *BriefPowerPortTemplate) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPortTemplate) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *BriefPowerPortTemplate) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *BriefPowerPortTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BriefPowerPortTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BriefPowerPortTemplate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BriefPowerPortTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BriefPowerPortTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BriefPowerPortTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BriefPowerPortTemplate) SetDescription(v string) {
	o.Description = &v
}

func (o BriefPowerPortTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefPowerPortTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefPowerPortTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
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

	varBriefPowerPortTemplate := _BriefPowerPortTemplate{}

	err = json.Unmarshal(data, &varBriefPowerPortTemplate)

	if err != nil {
		return err
	}

	*o = BriefPowerPortTemplate(varBriefPowerPortTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefPowerPortTemplate struct {
	value *BriefPowerPortTemplate
	isSet bool
}

func (v NullableBriefPowerPortTemplate) Get() *BriefPowerPortTemplate {
	return v.value
}

func (v *NullableBriefPowerPortTemplate) Set(val *BriefPowerPortTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefPowerPortTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefPowerPortTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefPowerPortTemplate(val *BriefPowerPortTemplate) *NullableBriefPowerPortTemplate {
	return &NullableBriefPowerPortTemplate{value: val, isSet: true}
}

func (v NullableBriefPowerPortTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefPowerPortTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
