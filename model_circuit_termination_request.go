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

// checks if the CircuitTerminationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CircuitTerminationRequest{}

// CircuitTerminationRequest Adds support for custom fields and tags.
type CircuitTerminationRequest struct {
	Circuit         BriefCircuitRequest                 `json:"circuit"`
	TermSide        Termination1                        `json:"term_side"`
	Site            NullableBriefSiteRequest            `json:"site,omitempty"`
	ProviderNetwork NullableBriefProviderNetworkRequest `json:"provider_network,omitempty"`
	// Physical circuit speed
	PortSpeed NullableInt32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed NullableInt32 `json:"upstream_speed,omitempty"`
	// ID of the local cross-connect
	XconnectId *string `json:"xconnect_id,omitempty"`
	// Patch panel ID and port number(s)
	PpInfo      *string `json:"pp_info,omitempty"`
	Description *string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected        *bool                  `json:"mark_connected,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CircuitTerminationRequest CircuitTerminationRequest

// NewCircuitTerminationRequest instantiates a new CircuitTerminationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCircuitTerminationRequest(circuit BriefCircuitRequest, termSide Termination1) *CircuitTerminationRequest {
	this := CircuitTerminationRequest{}
	this.Circuit = circuit
	this.TermSide = termSide
	return &this
}

// NewCircuitTerminationRequestWithDefaults instantiates a new CircuitTerminationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCircuitTerminationRequestWithDefaults() *CircuitTerminationRequest {
	this := CircuitTerminationRequest{}
	return &this
}

// GetCircuit returns the Circuit field value
func (o *CircuitTerminationRequest) GetCircuit() BriefCircuitRequest {
	if o == nil {
		var ret BriefCircuitRequest
		return ret
	}

	return o.Circuit
}

// GetCircuitOk returns a tuple with the Circuit field value
// and a boolean to check if the value has been set.
func (o *CircuitTerminationRequest) GetCircuitOk() (*BriefCircuitRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Circuit, true
}

// SetCircuit sets field value
func (o *CircuitTerminationRequest) SetCircuit(v BriefCircuitRequest) {
	o.Circuit = v
}

// GetTermSide returns the TermSide field value
func (o *CircuitTerminationRequest) GetTermSide() Termination1 {
	if o == nil {
		var ret Termination1
		return ret
	}

	return o.TermSide
}

// GetTermSideOk returns a tuple with the TermSide field value
// and a boolean to check if the value has been set.
func (o *CircuitTerminationRequest) GetTermSideOk() (*Termination1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TermSide, true
}

// SetTermSide sets field value
func (o *CircuitTerminationRequest) SetTermSide(v Termination1) {
	o.TermSide = v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitTerminationRequest) GetSite() BriefSiteRequest {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSiteRequest
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitTerminationRequest) GetSiteOk() (*BriefSiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSiteRequest and assigns it to the Site field.
func (o *CircuitTerminationRequest) SetSite(v BriefSiteRequest) {
	o.Site.Set(&v)
}

// SetSiteNil sets the value for Site to be an explicit nil
func (o *CircuitTerminationRequest) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *CircuitTerminationRequest) UnsetSite() {
	o.Site.Unset()
}

// GetProviderNetwork returns the ProviderNetwork field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitTerminationRequest) GetProviderNetwork() BriefProviderNetworkRequest {
	if o == nil || IsNil(o.ProviderNetwork.Get()) {
		var ret BriefProviderNetworkRequest
		return ret
	}
	return *o.ProviderNetwork.Get()
}

// GetProviderNetworkOk returns a tuple with the ProviderNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitTerminationRequest) GetProviderNetworkOk() (*BriefProviderNetworkRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderNetwork.Get(), o.ProviderNetwork.IsSet()
}

// HasProviderNetwork returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasProviderNetwork() bool {
	if o != nil && o.ProviderNetwork.IsSet() {
		return true
	}

	return false
}

// SetProviderNetwork gets a reference to the given NullableBriefProviderNetworkRequest and assigns it to the ProviderNetwork field.
func (o *CircuitTerminationRequest) SetProviderNetwork(v BriefProviderNetworkRequest) {
	o.ProviderNetwork.Set(&v)
}

// SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil
func (o *CircuitTerminationRequest) SetProviderNetworkNil() {
	o.ProviderNetwork.Set(nil)
}

// UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
func (o *CircuitTerminationRequest) UnsetProviderNetwork() {
	o.ProviderNetwork.Unset()
}

// GetPortSpeed returns the PortSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitTerminationRequest) GetPortSpeed() int32 {
	if o == nil || IsNil(o.PortSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.PortSpeed.Get()
}

// GetPortSpeedOk returns a tuple with the PortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitTerminationRequest) GetPortSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortSpeed.Get(), o.PortSpeed.IsSet()
}

// HasPortSpeed returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasPortSpeed() bool {
	if o != nil && o.PortSpeed.IsSet() {
		return true
	}

	return false
}

// SetPortSpeed gets a reference to the given NullableInt32 and assigns it to the PortSpeed field.
func (o *CircuitTerminationRequest) SetPortSpeed(v int32) {
	o.PortSpeed.Set(&v)
}

// SetPortSpeedNil sets the value for PortSpeed to be an explicit nil
func (o *CircuitTerminationRequest) SetPortSpeedNil() {
	o.PortSpeed.Set(nil)
}

// UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
func (o *CircuitTerminationRequest) UnsetPortSpeed() {
	o.PortSpeed.Unset()
}

// GetUpstreamSpeed returns the UpstreamSpeed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CircuitTerminationRequest) GetUpstreamSpeed() int32 {
	if o == nil || IsNil(o.UpstreamSpeed.Get()) {
		var ret int32
		return ret
	}
	return *o.UpstreamSpeed.Get()
}

// GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpstreamSpeed.Get(), o.UpstreamSpeed.IsSet()
}

// HasUpstreamSpeed returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasUpstreamSpeed() bool {
	if o != nil && o.UpstreamSpeed.IsSet() {
		return true
	}

	return false
}

// SetUpstreamSpeed gets a reference to the given NullableInt32 and assigns it to the UpstreamSpeed field.
func (o *CircuitTerminationRequest) SetUpstreamSpeed(v int32) {
	o.UpstreamSpeed.Set(&v)
}

// SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil
func (o *CircuitTerminationRequest) SetUpstreamSpeedNil() {
	o.UpstreamSpeed.Set(nil)
}

// UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
func (o *CircuitTerminationRequest) UnsetUpstreamSpeed() {
	o.UpstreamSpeed.Unset()
}

// GetXconnectId returns the XconnectId field value if set, zero value otherwise.
func (o *CircuitTerminationRequest) GetXconnectId() string {
	if o == nil || IsNil(o.XconnectId) {
		var ret string
		return ret
	}
	return *o.XconnectId
}

// GetXconnectIdOk returns a tuple with the XconnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTerminationRequest) GetXconnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.XconnectId) {
		return nil, false
	}
	return o.XconnectId, true
}

// HasXconnectId returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasXconnectId() bool {
	if o != nil && !IsNil(o.XconnectId) {
		return true
	}

	return false
}

// SetXconnectId gets a reference to the given string and assigns it to the XconnectId field.
func (o *CircuitTerminationRequest) SetXconnectId(v string) {
	o.XconnectId = &v
}

// GetPpInfo returns the PpInfo field value if set, zero value otherwise.
func (o *CircuitTerminationRequest) GetPpInfo() string {
	if o == nil || IsNil(o.PpInfo) {
		var ret string
		return ret
	}
	return *o.PpInfo
}

// GetPpInfoOk returns a tuple with the PpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTerminationRequest) GetPpInfoOk() (*string, bool) {
	if o == nil || IsNil(o.PpInfo) {
		return nil, false
	}
	return o.PpInfo, true
}

// HasPpInfo returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasPpInfo() bool {
	if o != nil && !IsNil(o.PpInfo) {
		return true
	}

	return false
}

// SetPpInfo gets a reference to the given string and assigns it to the PpInfo field.
func (o *CircuitTerminationRequest) SetPpInfo(v string) {
	o.PpInfo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CircuitTerminationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTerminationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CircuitTerminationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMarkConnected returns the MarkConnected field value if set, zero value otherwise.
func (o *CircuitTerminationRequest) GetMarkConnected() bool {
	if o == nil || IsNil(o.MarkConnected) {
		var ret bool
		return ret
	}
	return *o.MarkConnected
}

// GetMarkConnectedOk returns a tuple with the MarkConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTerminationRequest) GetMarkConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.MarkConnected) {
		return nil, false
	}
	return o.MarkConnected, true
}

// HasMarkConnected returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasMarkConnected() bool {
	if o != nil && !IsNil(o.MarkConnected) {
		return true
	}

	return false
}

// SetMarkConnected gets a reference to the given bool and assigns it to the MarkConnected field.
func (o *CircuitTerminationRequest) SetMarkConnected(v bool) {
	o.MarkConnected = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CircuitTerminationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTerminationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *CircuitTerminationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CircuitTerminationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CircuitTerminationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CircuitTerminationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CircuitTerminationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o CircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CircuitTerminationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["circuit"] = o.Circuit
	toSerialize["term_side"] = o.TermSide
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.ProviderNetwork.IsSet() {
		toSerialize["provider_network"] = o.ProviderNetwork.Get()
	}
	if o.PortSpeed.IsSet() {
		toSerialize["port_speed"] = o.PortSpeed.Get()
	}
	if o.UpstreamSpeed.IsSet() {
		toSerialize["upstream_speed"] = o.UpstreamSpeed.Get()
	}
	if !IsNil(o.XconnectId) {
		toSerialize["xconnect_id"] = o.XconnectId
	}
	if !IsNil(o.PpInfo) {
		toSerialize["pp_info"] = o.PpInfo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MarkConnected) {
		toSerialize["mark_connected"] = o.MarkConnected
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

func (o *CircuitTerminationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"circuit",
		"term_side",
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

	varCircuitTerminationRequest := _CircuitTerminationRequest{}

	err = json.Unmarshal(data, &varCircuitTerminationRequest)

	if err != nil {
		return err
	}

	*o = CircuitTerminationRequest(varCircuitTerminationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "circuit")
		delete(additionalProperties, "term_side")
		delete(additionalProperties, "site")
		delete(additionalProperties, "provider_network")
		delete(additionalProperties, "port_speed")
		delete(additionalProperties, "upstream_speed")
		delete(additionalProperties, "xconnect_id")
		delete(additionalProperties, "pp_info")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mark_connected")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCircuitTerminationRequest struct {
	value *CircuitTerminationRequest
	isSet bool
}

func (v NullableCircuitTerminationRequest) Get() *CircuitTerminationRequest {
	return v.value
}

func (v *NullableCircuitTerminationRequest) Set(val *CircuitTerminationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCircuitTerminationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCircuitTerminationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCircuitTerminationRequest(val *CircuitTerminationRequest) *NullableCircuitTerminationRequest {
	return &NullableCircuitTerminationRequest{value: val, isSet: true}
}

func (v NullableCircuitTerminationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCircuitTerminationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
