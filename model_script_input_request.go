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
	"time"
)

// checks if the ScriptInputRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScriptInputRequest{}

// ScriptInputRequest struct for ScriptInputRequest
type ScriptInputRequest struct {
	Data                 interface{}   `json:"data"`
	Commit               bool          `json:"commit"`
	ScheduleAt           NullableTime  `json:"schedule_at,omitempty"`
	Interval             NullableInt32 `json:"interval,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScriptInputRequest ScriptInputRequest

// NewScriptInputRequest instantiates a new ScriptInputRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScriptInputRequest(data interface{}, commit bool) *ScriptInputRequest {
	this := ScriptInputRequest{}
	this.Data = data
	this.Commit = commit
	return &this
}

// NewScriptInputRequestWithDefaults instantiates a new ScriptInputRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScriptInputRequestWithDefaults() *ScriptInputRequest {
	this := ScriptInputRequest{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ScriptInputRequest) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScriptInputRequest) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ScriptInputRequest) SetData(v interface{}) {
	o.Data = v
}

// GetCommit returns the Commit field value
func (o *ScriptInputRequest) GetCommit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *ScriptInputRequest) GetCommitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *ScriptInputRequest) SetCommit(v bool) {
	o.Commit = v
}

// GetScheduleAt returns the ScheduleAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScriptInputRequest) GetScheduleAt() time.Time {
	if o == nil || IsNil(o.ScheduleAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ScheduleAt.Get()
}

// GetScheduleAtOk returns a tuple with the ScheduleAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScriptInputRequest) GetScheduleAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleAt.Get(), o.ScheduleAt.IsSet()
}

// HasScheduleAt returns a boolean if a field has been set.
func (o *ScriptInputRequest) HasScheduleAt() bool {
	if o != nil && o.ScheduleAt.IsSet() {
		return true
	}

	return false
}

// SetScheduleAt gets a reference to the given NullableTime and assigns it to the ScheduleAt field.
func (o *ScriptInputRequest) SetScheduleAt(v time.Time) {
	o.ScheduleAt.Set(&v)
}

// SetScheduleAtNil sets the value for ScheduleAt to be an explicit nil
func (o *ScriptInputRequest) SetScheduleAtNil() {
	o.ScheduleAt.Set(nil)
}

// UnsetScheduleAt ensures that no value is present for ScheduleAt, not even an explicit nil
func (o *ScriptInputRequest) UnsetScheduleAt() {
	o.ScheduleAt.Unset()
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScriptInputRequest) GetInterval() int32 {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret int32
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScriptInputRequest) GetIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *ScriptInputRequest) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableInt32 and assigns it to the Interval field.
func (o *ScriptInputRequest) SetInterval(v int32) {
	o.Interval.Set(&v)
}

// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *ScriptInputRequest) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *ScriptInputRequest) UnsetInterval() {
	o.Interval.Unset()
}

func (o ScriptInputRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScriptInputRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["commit"] = o.Commit
	if o.ScheduleAt.IsSet() {
		toSerialize["schedule_at"] = o.ScheduleAt.Get()
	}
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScriptInputRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"commit",
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

	varScriptInputRequest := _ScriptInputRequest{}

	err = json.Unmarshal(data, &varScriptInputRequest)

	if err != nil {
		return err
	}

	*o = ScriptInputRequest(varScriptInputRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "commit")
		delete(additionalProperties, "schedule_at")
		delete(additionalProperties, "interval")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScriptInputRequest struct {
	value *ScriptInputRequest
	isSet bool
}

func (v NullableScriptInputRequest) Get() *ScriptInputRequest {
	return v.value
}

func (v *NullableScriptInputRequest) Set(val *ScriptInputRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScriptInputRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScriptInputRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScriptInputRequest(val *ScriptInputRequest) *NullableScriptInputRequest {
	return &NullableScriptInputRequest{value: val, isSet: true}
}

func (v NullableScriptInputRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScriptInputRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
