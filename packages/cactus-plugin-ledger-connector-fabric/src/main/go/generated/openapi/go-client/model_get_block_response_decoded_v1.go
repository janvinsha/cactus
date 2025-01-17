/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the GetBlockResponseDecodedV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBlockResponseDecodedV1{}

// GetBlockResponseDecodedV1 When skipDecode is false (default) then decoded block object is returned.
type GetBlockResponseDecodedV1 struct {
	// Full hyperledger fabric block data.
	DecodedBlock interface{} `json:"decodedBlock"`
}

// NewGetBlockResponseDecodedV1 instantiates a new GetBlockResponseDecodedV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockResponseDecodedV1(decodedBlock interface{}) *GetBlockResponseDecodedV1 {
	this := GetBlockResponseDecodedV1{}
	this.DecodedBlock = decodedBlock
	return &this
}

// NewGetBlockResponseDecodedV1WithDefaults instantiates a new GetBlockResponseDecodedV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockResponseDecodedV1WithDefaults() *GetBlockResponseDecodedV1 {
	this := GetBlockResponseDecodedV1{}
	return &this
}

// GetDecodedBlock returns the DecodedBlock field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GetBlockResponseDecodedV1) GetDecodedBlock() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.DecodedBlock
}

// GetDecodedBlockOk returns a tuple with the DecodedBlock field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBlockResponseDecodedV1) GetDecodedBlockOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DecodedBlock) {
		return nil, false
	}
	return &o.DecodedBlock, true
}

// SetDecodedBlock sets field value
func (o *GetBlockResponseDecodedV1) SetDecodedBlock(v interface{}) {
	o.DecodedBlock = v
}

func (o GetBlockResponseDecodedV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBlockResponseDecodedV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DecodedBlock != nil {
		toSerialize["decodedBlock"] = o.DecodedBlock
	}
	return toSerialize, nil
}

type NullableGetBlockResponseDecodedV1 struct {
	value *GetBlockResponseDecodedV1
	isSet bool
}

func (v NullableGetBlockResponseDecodedV1) Get() *GetBlockResponseDecodedV1 {
	return v.value
}

func (v *NullableGetBlockResponseDecodedV1) Set(val *GetBlockResponseDecodedV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockResponseDecodedV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockResponseDecodedV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockResponseDecodedV1(val *GetBlockResponseDecodedV1) *NullableGetBlockResponseDecodedV1 {
	return &NullableGetBlockResponseDecodedV1{value: val, isSet: true}
}

func (v NullableGetBlockResponseDecodedV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockResponseDecodedV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


