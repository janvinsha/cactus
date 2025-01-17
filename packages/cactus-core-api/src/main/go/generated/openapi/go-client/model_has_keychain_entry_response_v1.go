/*
Hyperledger Core API

Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-core-api

import (
	"encoding/json"
)

// checks if the HasKeychainEntryResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasKeychainEntryResponseV1{}

// HasKeychainEntryResponseV1 struct for HasKeychainEntryResponseV1
type HasKeychainEntryResponseV1 struct {
	// The key that was used to check the presence of the value in the entry store.
	Key string `json:"key"`
	// Date and time encoded as JSON when the presence check was performed by the plugin backend.
	CheckedAt string `json:"checkedAt"`
	// The boolean true or false indicating the presence or absence of an entry under 'key'.
	IsPresent bool `json:"isPresent"`
}

// NewHasKeychainEntryResponseV1 instantiates a new HasKeychainEntryResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasKeychainEntryResponseV1(key string, checkedAt string, isPresent bool) *HasKeychainEntryResponseV1 {
	this := HasKeychainEntryResponseV1{}
	this.Key = key
	this.CheckedAt = checkedAt
	this.IsPresent = isPresent
	return &this
}

// NewHasKeychainEntryResponseV1WithDefaults instantiates a new HasKeychainEntryResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasKeychainEntryResponseV1WithDefaults() *HasKeychainEntryResponseV1 {
	this := HasKeychainEntryResponseV1{}
	return &this
}

// GetKey returns the Key field value
func (o *HasKeychainEntryResponseV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *HasKeychainEntryResponseV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *HasKeychainEntryResponseV1) SetKey(v string) {
	o.Key = v
}

// GetCheckedAt returns the CheckedAt field value
func (o *HasKeychainEntryResponseV1) GetCheckedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckedAt
}

// GetCheckedAtOk returns a tuple with the CheckedAt field value
// and a boolean to check if the value has been set.
func (o *HasKeychainEntryResponseV1) GetCheckedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckedAt, true
}

// SetCheckedAt sets field value
func (o *HasKeychainEntryResponseV1) SetCheckedAt(v string) {
	o.CheckedAt = v
}

// GetIsPresent returns the IsPresent field value
func (o *HasKeychainEntryResponseV1) GetIsPresent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPresent
}

// GetIsPresentOk returns a tuple with the IsPresent field value
// and a boolean to check if the value has been set.
func (o *HasKeychainEntryResponseV1) GetIsPresentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPresent, true
}

// SetIsPresent sets field value
func (o *HasKeychainEntryResponseV1) SetIsPresent(v bool) {
	o.IsPresent = v
}

func (o HasKeychainEntryResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasKeychainEntryResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["checkedAt"] = o.CheckedAt
	toSerialize["isPresent"] = o.IsPresent
	return toSerialize, nil
}

type NullableHasKeychainEntryResponseV1 struct {
	value *HasKeychainEntryResponseV1
	isSet bool
}

func (v NullableHasKeychainEntryResponseV1) Get() *HasKeychainEntryResponseV1 {
	return v.value
}

func (v *NullableHasKeychainEntryResponseV1) Set(val *HasKeychainEntryResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableHasKeychainEntryResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableHasKeychainEntryResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasKeychainEntryResponseV1(val *HasKeychainEntryResponseV1) *NullableHasKeychainEntryResponseV1 {
	return &NullableHasKeychainEntryResponseV1{value: val, isSet: true}
}

func (v NullableHasKeychainEntryResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasKeychainEntryResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


