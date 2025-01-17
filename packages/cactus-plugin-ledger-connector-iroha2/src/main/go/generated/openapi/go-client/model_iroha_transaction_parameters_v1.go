/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
)

// checks if the IrohaTransactionParametersV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IrohaTransactionParametersV1{}

// IrohaTransactionParametersV1 Iroha V2 transaction payload parameters
type IrohaTransactionParametersV1 struct {
	// BigInt time to live.
	Ttl *string `json:"ttl,omitempty"`
	// BigInt creation time
	CreationTime *string `json:"creationTime,omitempty"`
	// Transaction nonce
	Nonce *float32 `json:"nonce,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IrohaTransactionParametersV1 IrohaTransactionParametersV1

// NewIrohaTransactionParametersV1 instantiates a new IrohaTransactionParametersV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIrohaTransactionParametersV1() *IrohaTransactionParametersV1 {
	this := IrohaTransactionParametersV1{}
	return &this
}

// NewIrohaTransactionParametersV1WithDefaults instantiates a new IrohaTransactionParametersV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIrohaTransactionParametersV1WithDefaults() *IrohaTransactionParametersV1 {
	this := IrohaTransactionParametersV1{}
	return &this
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *IrohaTransactionParametersV1) GetTtl() string {
	if o == nil || IsNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IrohaTransactionParametersV1) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *IrohaTransactionParametersV1) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *IrohaTransactionParametersV1) SetTtl(v string) {
	o.Ttl = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *IrohaTransactionParametersV1) GetCreationTime() string {
	if o == nil || IsNil(o.CreationTime) {
		var ret string
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IrohaTransactionParametersV1) GetCreationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *IrohaTransactionParametersV1) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given string and assigns it to the CreationTime field.
func (o *IrohaTransactionParametersV1) SetCreationTime(v string) {
	o.CreationTime = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *IrohaTransactionParametersV1) GetNonce() float32 {
	if o == nil || IsNil(o.Nonce) {
		var ret float32
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IrohaTransactionParametersV1) GetNonceOk() (*float32, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *IrohaTransactionParametersV1) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given float32 and assigns it to the Nonce field.
func (o *IrohaTransactionParametersV1) SetNonce(v float32) {
	o.Nonce = &v
}

func (o IrohaTransactionParametersV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IrohaTransactionParametersV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IrohaTransactionParametersV1) UnmarshalJSON(bytes []byte) (err error) {
	varIrohaTransactionParametersV1 := _IrohaTransactionParametersV1{}

	if err = json.Unmarshal(bytes, &varIrohaTransactionParametersV1); err == nil {
		*o = IrohaTransactionParametersV1(varIrohaTransactionParametersV1)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ttl")
		delete(additionalProperties, "creationTime")
		delete(additionalProperties, "nonce")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIrohaTransactionParametersV1 struct {
	value *IrohaTransactionParametersV1
	isSet bool
}

func (v NullableIrohaTransactionParametersV1) Get() *IrohaTransactionParametersV1 {
	return v.value
}

func (v *NullableIrohaTransactionParametersV1) Set(val *IrohaTransactionParametersV1) {
	v.value = val
	v.isSet = true
}

func (v NullableIrohaTransactionParametersV1) IsSet() bool {
	return v.isSet
}

func (v *NullableIrohaTransactionParametersV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIrohaTransactionParametersV1(val *IrohaTransactionParametersV1) *NullableIrohaTransactionParametersV1 {
	return &NullableIrohaTransactionParametersV1{value: val, isSet: true}
}

func (v NullableIrohaTransactionParametersV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIrohaTransactionParametersV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


