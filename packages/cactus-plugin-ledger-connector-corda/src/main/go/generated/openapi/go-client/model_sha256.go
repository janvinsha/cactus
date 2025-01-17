/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the SHA256 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SHA256{}

// SHA256 SHA-256 is part of the SHA-2 hash function family. Generated hash is fixed size, 256-bits (32-bytes).
type SHA256 struct {
	Bytes string `json:"bytes"`
	Offset int32 `json:"offset"`
	Size int32 `json:"size"`
}

// NewSHA256 instantiates a new SHA256 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSHA256(bytes string, offset int32, size int32) *SHA256 {
	this := SHA256{}
	this.Bytes = bytes
	this.Offset = offset
	this.Size = size
	return &this
}

// NewSHA256WithDefaults instantiates a new SHA256 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSHA256WithDefaults() *SHA256 {
	this := SHA256{}
	return &this
}

// GetBytes returns the Bytes field value
func (o *SHA256) GetBytes() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value
// and a boolean to check if the value has been set.
func (o *SHA256) GetBytesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bytes, true
}

// SetBytes sets field value
func (o *SHA256) SetBytes(v string) {
	o.Bytes = v
}

// GetOffset returns the Offset field value
func (o *SHA256) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *SHA256) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *SHA256) SetOffset(v int32) {
	o.Offset = v
}

// GetSize returns the Size field value
func (o *SHA256) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *SHA256) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *SHA256) SetSize(v int32) {
	o.Size = v
}

func (o SHA256) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SHA256) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bytes"] = o.Bytes
	toSerialize["offset"] = o.Offset
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

type NullableSHA256 struct {
	value *SHA256
	isSet bool
}

func (v NullableSHA256) Get() *SHA256 {
	return v.value
}

func (v *NullableSHA256) Set(val *SHA256) {
	v.value = val
	v.isSet = true
}

func (v NullableSHA256) IsSet() bool {
	return v.isSet
}

func (v *NullableSHA256) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSHA256(val *SHA256) *NullableSHA256 {
	return &NullableSHA256{value: val, isSet: true}
}

func (v NullableSHA256) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSHA256) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


