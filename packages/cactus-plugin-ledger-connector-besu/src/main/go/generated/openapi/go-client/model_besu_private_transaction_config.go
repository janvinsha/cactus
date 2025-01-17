/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the BesuPrivateTransactionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BesuPrivateTransactionConfig{}

// BesuPrivateTransactionConfig struct for BesuPrivateTransactionConfig
type BesuPrivateTransactionConfig struct {
	PrivateFrom string `json:"privateFrom"`
	PrivateFor []interface{} `json:"privateFor"`
}

// NewBesuPrivateTransactionConfig instantiates a new BesuPrivateTransactionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBesuPrivateTransactionConfig(privateFrom string, privateFor []interface{}) *BesuPrivateTransactionConfig {
	this := BesuPrivateTransactionConfig{}
	this.PrivateFrom = privateFrom
	this.PrivateFor = privateFor
	return &this
}

// NewBesuPrivateTransactionConfigWithDefaults instantiates a new BesuPrivateTransactionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBesuPrivateTransactionConfigWithDefaults() *BesuPrivateTransactionConfig {
	this := BesuPrivateTransactionConfig{}
	return &this
}

// GetPrivateFrom returns the PrivateFrom field value
func (o *BesuPrivateTransactionConfig) GetPrivateFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateFrom
}

// GetPrivateFromOk returns a tuple with the PrivateFrom field value
// and a boolean to check if the value has been set.
func (o *BesuPrivateTransactionConfig) GetPrivateFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateFrom, true
}

// SetPrivateFrom sets field value
func (o *BesuPrivateTransactionConfig) SetPrivateFrom(v string) {
	o.PrivateFrom = v
}

// GetPrivateFor returns the PrivateFor field value
func (o *BesuPrivateTransactionConfig) GetPrivateFor() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.PrivateFor
}

// GetPrivateForOk returns a tuple with the PrivateFor field value
// and a boolean to check if the value has been set.
func (o *BesuPrivateTransactionConfig) GetPrivateForOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateFor, true
}

// SetPrivateFor sets field value
func (o *BesuPrivateTransactionConfig) SetPrivateFor(v []interface{}) {
	o.PrivateFor = v
}

func (o BesuPrivateTransactionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BesuPrivateTransactionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privateFrom"] = o.PrivateFrom
	toSerialize["privateFor"] = o.PrivateFor
	return toSerialize, nil
}

type NullableBesuPrivateTransactionConfig struct {
	value *BesuPrivateTransactionConfig
	isSet bool
}

func (v NullableBesuPrivateTransactionConfig) Get() *BesuPrivateTransactionConfig {
	return v.value
}

func (v *NullableBesuPrivateTransactionConfig) Set(val *BesuPrivateTransactionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBesuPrivateTransactionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBesuPrivateTransactionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBesuPrivateTransactionConfig(val *BesuPrivateTransactionConfig) *NullableBesuPrivateTransactionConfig {
	return &NullableBesuPrivateTransactionConfig{value: val, isSet: true}
}

func (v NullableBesuPrivateTransactionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBesuPrivateTransactionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


