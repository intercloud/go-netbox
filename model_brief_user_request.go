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

// checks if the BriefUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BriefUserRequest{}

// BriefUserRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BriefUserRequest struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username             string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _BriefUserRequest BriefUserRequest

// NewBriefUserRequest instantiates a new BriefUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBriefUserRequest(username string) *BriefUserRequest {
	this := BriefUserRequest{}
	this.Username = username
	return &this
}

// NewBriefUserRequestWithDefaults instantiates a new BriefUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBriefUserRequestWithDefaults() *BriefUserRequest {
	this := BriefUserRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *BriefUserRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *BriefUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *BriefUserRequest) SetUsername(v string) {
	o.Username = v
}

func (o BriefUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BriefUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BriefUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varBriefUserRequest := _BriefUserRequest{}

	err = json.Unmarshal(data, &varBriefUserRequest)

	if err != nil {
		return err
	}

	*o = BriefUserRequest(varBriefUserRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBriefUserRequest struct {
	value *BriefUserRequest
	isSet bool
}

func (v NullableBriefUserRequest) Get() *BriefUserRequest {
	return v.value
}

func (v *NullableBriefUserRequest) Set(val *BriefUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBriefUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBriefUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBriefUserRequest(val *BriefUserRequest) *NullableBriefUserRequest {
	return &NullableBriefUserRequest{value: val, isSet: true}
}

func (v NullableBriefUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBriefUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
