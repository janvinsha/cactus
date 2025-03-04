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

// checks if the GetBesuRecordV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBesuRecordV1Request{}

// GetBesuRecordV1Request struct for GetBesuRecordV1Request
type GetBesuRecordV1Request struct {
	InvokeCall *InvokeContractV1Request `json:"invokeCall,omitempty"`
	TransactionHash *string `json:"transactionHash,omitempty"`
}

// NewGetBesuRecordV1Request instantiates a new GetBesuRecordV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBesuRecordV1Request() *GetBesuRecordV1Request {
	this := GetBesuRecordV1Request{}
	return &this
}

// NewGetBesuRecordV1RequestWithDefaults instantiates a new GetBesuRecordV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBesuRecordV1RequestWithDefaults() *GetBesuRecordV1Request {
	this := GetBesuRecordV1Request{}
	return &this
}

// GetInvokeCall returns the InvokeCall field value if set, zero value otherwise.
func (o *GetBesuRecordV1Request) GetInvokeCall() InvokeContractV1Request {
	if o == nil || IsNil(o.InvokeCall) {
		var ret InvokeContractV1Request
		return ret
	}
	return *o.InvokeCall
}

// GetInvokeCallOk returns a tuple with the InvokeCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBesuRecordV1Request) GetInvokeCallOk() (*InvokeContractV1Request, bool) {
	if o == nil || IsNil(o.InvokeCall) {
		return nil, false
	}
	return o.InvokeCall, true
}

// HasInvokeCall returns a boolean if a field has been set.
func (o *GetBesuRecordV1Request) HasInvokeCall() bool {
	if o != nil && !IsNil(o.InvokeCall) {
		return true
	}

	return false
}

// SetInvokeCall gets a reference to the given InvokeContractV1Request and assigns it to the InvokeCall field.
func (o *GetBesuRecordV1Request) SetInvokeCall(v InvokeContractV1Request) {
	o.InvokeCall = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *GetBesuRecordV1Request) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBesuRecordV1Request) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *GetBesuRecordV1Request) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *GetBesuRecordV1Request) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

func (o GetBesuRecordV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBesuRecordV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvokeCall) {
		toSerialize["invokeCall"] = o.InvokeCall
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transactionHash"] = o.TransactionHash
	}
	return toSerialize, nil
}

type NullableGetBesuRecordV1Request struct {
	value *GetBesuRecordV1Request
	isSet bool
}

func (v NullableGetBesuRecordV1Request) Get() *GetBesuRecordV1Request {
	return v.value
}

func (v *NullableGetBesuRecordV1Request) Set(val *GetBesuRecordV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBesuRecordV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBesuRecordV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBesuRecordV1Request(val *GetBesuRecordV1Request) *NullableGetBesuRecordV1Request {
	return &NullableGetBesuRecordV1Request{value: val, isSet: true}
}

func (v NullableGetBesuRecordV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBesuRecordV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


