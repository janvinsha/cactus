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

// checks if the GetBlockV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBlockV1Request{}

// GetBlockV1Request struct for GetBlockV1Request
type GetBlockV1Request struct {
	BlockHashOrBlockNumber interface{} `json:"blockHashOrBlockNumber"`
}

// NewGetBlockV1Request instantiates a new GetBlockV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBlockV1Request(blockHashOrBlockNumber interface{}) *GetBlockV1Request {
	this := GetBlockV1Request{}
	this.BlockHashOrBlockNumber = blockHashOrBlockNumber
	return &this
}

// NewGetBlockV1RequestWithDefaults instantiates a new GetBlockV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBlockV1RequestWithDefaults() *GetBlockV1Request {
	this := GetBlockV1Request{}
	return &this
}

// GetBlockHashOrBlockNumber returns the BlockHashOrBlockNumber field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *GetBlockV1Request) GetBlockHashOrBlockNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.BlockHashOrBlockNumber
}

// GetBlockHashOrBlockNumberOk returns a tuple with the BlockHashOrBlockNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBlockV1Request) GetBlockHashOrBlockNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BlockHashOrBlockNumber) {
		return nil, false
	}
	return &o.BlockHashOrBlockNumber, true
}

// SetBlockHashOrBlockNumber sets field value
func (o *GetBlockV1Request) SetBlockHashOrBlockNumber(v interface{}) {
	o.BlockHashOrBlockNumber = v
}

func (o GetBlockV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBlockV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockHashOrBlockNumber != nil {
		toSerialize["blockHashOrBlockNumber"] = o.BlockHashOrBlockNumber
	}
	return toSerialize, nil
}

type NullableGetBlockV1Request struct {
	value *GetBlockV1Request
	isSet bool
}

func (v NullableGetBlockV1Request) Get() *GetBlockV1Request {
	return v.value
}

func (v *NullableGetBlockV1Request) Set(val *GetBlockV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBlockV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBlockV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBlockV1Request(val *GetBlockV1Request) *NullableGetBlockV1Request {
	return &NullableGetBlockV1Request{value: val, isSet: true}
}

func (v NullableGetBlockV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBlockV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


