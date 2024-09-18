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

// IKEProposalEncryptionAlgorithmValue * `aes-128-cbc` - 128-bit AES (CBC) * `aes-128-gcm` - 128-bit AES (GCM) * `aes-192-cbc` - 192-bit AES (CBC) * `aes-192-gcm` - 192-bit AES (GCM) * `aes-256-cbc` - 256-bit AES (CBC) * `aes-256-gcm` - 256-bit AES (GCM) * `3des-cbc` - 3DES * `des-cbc` - DES
type IKEProposalEncryptionAlgorithmValue string

// List of IKEProposal_encryption_algorithm_value
const (
	IKEPROPOSALENCRYPTIONALGORITHMVALUE_AES_128_CBC IKEProposalEncryptionAlgorithmValue = "aes-128-cbc"
	IKEPROPOSALENCRYPTIONALGORITHMVALUE_AES_128_GCM IKEProposalEncryptionAlgorithmValue = "aes-128-gcm"
	IKEPROPOSALENCRYPTIONALGORITHMVALUE_AES_192_CBC IKEProposalEncryptionAlgorithmValue = "aes-192-cbc"
	IKEPROPOSALENCRYPTIONALGORITHMVALUE_AES_192_GCM IKEProposalEncryptionAlgorithmValue = "aes-192-gcm"
	IKEPROPOSALENCRYPTIONALGORITHMVALUE_AES_256_CBC IKEProposalEncryptionAlgorithmValue = "aes-256-cbc"
	IKEPROPOSALENCRYPTIONALGORITHMVALUE_AES_256_GCM IKEProposalEncryptionAlgorithmValue = "aes-256-gcm"
	IKEPROPOSALENCRYPTIONALGORITHMVALUE__3DES_CBC   IKEProposalEncryptionAlgorithmValue = "3des-cbc"
	IKEPROPOSALENCRYPTIONALGORITHMVALUE_DES_CBC     IKEProposalEncryptionAlgorithmValue = "des-cbc"
)

// All allowed values of IKEProposalEncryptionAlgorithmValue enum
var AllowedIKEProposalEncryptionAlgorithmValueEnumValues = []IKEProposalEncryptionAlgorithmValue{
	"aes-128-cbc",
	"aes-128-gcm",
	"aes-192-cbc",
	"aes-192-gcm",
	"aes-256-cbc",
	"aes-256-gcm",
	"3des-cbc",
	"des-cbc",
}

func (v *IKEProposalEncryptionAlgorithmValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IKEProposalEncryptionAlgorithmValue(value)
	for _, existing := range AllowedIKEProposalEncryptionAlgorithmValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IKEProposalEncryptionAlgorithmValue", value)
}

// NewIKEProposalEncryptionAlgorithmValueFromValue returns a pointer to a valid IKEProposalEncryptionAlgorithmValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIKEProposalEncryptionAlgorithmValueFromValue(v string) (*IKEProposalEncryptionAlgorithmValue, error) {
	ev := IKEProposalEncryptionAlgorithmValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IKEProposalEncryptionAlgorithmValue: valid values are %v", v, AllowedIKEProposalEncryptionAlgorithmValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IKEProposalEncryptionAlgorithmValue) IsValid() bool {
	for _, existing := range AllowedIKEProposalEncryptionAlgorithmValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IKEProposal_encryption_algorithm_value value
func (v IKEProposalEncryptionAlgorithmValue) Ptr() *IKEProposalEncryptionAlgorithmValue {
	return &v
}

type NullableIKEProposalEncryptionAlgorithmValue struct {
	value *IKEProposalEncryptionAlgorithmValue
	isSet bool
}

func (v NullableIKEProposalEncryptionAlgorithmValue) Get() *IKEProposalEncryptionAlgorithmValue {
	return v.value
}

func (v *NullableIKEProposalEncryptionAlgorithmValue) Set(val *IKEProposalEncryptionAlgorithmValue) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposalEncryptionAlgorithmValue) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposalEncryptionAlgorithmValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposalEncryptionAlgorithmValue(val *IKEProposalEncryptionAlgorithmValue) *NullableIKEProposalEncryptionAlgorithmValue {
	return &NullableIKEProposalEncryptionAlgorithmValue{value: val, isSet: true}
}

func (v NullableIKEProposalEncryptionAlgorithmValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposalEncryptionAlgorithmValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
