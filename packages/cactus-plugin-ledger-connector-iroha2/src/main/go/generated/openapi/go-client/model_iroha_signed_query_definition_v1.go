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

// checks if the IrohaSignedQueryDefinitionV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IrohaSignedQueryDefinitionV1{}

// IrohaSignedQueryDefinitionV1 Iroha V2 signed query definition
type IrohaSignedQueryDefinitionV1 struct {
	// Name of the query to be executed.
	Query IrohaQuery `json:"query"`
	// Signed query transaction binary data received from generate-transaction endpoint.
	Payload string `json:"payload"`
}

// NewIrohaSignedQueryDefinitionV1 instantiates a new IrohaSignedQueryDefinitionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIrohaSignedQueryDefinitionV1(query IrohaQuery, payload string) *IrohaSignedQueryDefinitionV1 {
	this := IrohaSignedQueryDefinitionV1{}
	this.Query = query
	this.Payload = payload
	return &this
}

// NewIrohaSignedQueryDefinitionV1WithDefaults instantiates a new IrohaSignedQueryDefinitionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIrohaSignedQueryDefinitionV1WithDefaults() *IrohaSignedQueryDefinitionV1 {
	this := IrohaSignedQueryDefinitionV1{}
	return &this
}

// GetQuery returns the Query field value
func (o *IrohaSignedQueryDefinitionV1) GetQuery() IrohaQuery {
	if o == nil {
		var ret IrohaQuery
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *IrohaSignedQueryDefinitionV1) GetQueryOk() (*IrohaQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *IrohaSignedQueryDefinitionV1) SetQuery(v IrohaQuery) {
	o.Query = v
}

// GetPayload returns the Payload field value
func (o *IrohaSignedQueryDefinitionV1) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *IrohaSignedQueryDefinitionV1) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *IrohaSignedQueryDefinitionV1) SetPayload(v string) {
	o.Payload = v
}

func (o IrohaSignedQueryDefinitionV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IrohaSignedQueryDefinitionV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query"] = o.Query
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

type NullableIrohaSignedQueryDefinitionV1 struct {
	value *IrohaSignedQueryDefinitionV1
	isSet bool
}

func (v NullableIrohaSignedQueryDefinitionV1) Get() *IrohaSignedQueryDefinitionV1 {
	return v.value
}

func (v *NullableIrohaSignedQueryDefinitionV1) Set(val *IrohaSignedQueryDefinitionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableIrohaSignedQueryDefinitionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableIrohaSignedQueryDefinitionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIrohaSignedQueryDefinitionV1(val *IrohaSignedQueryDefinitionV1) *NullableIrohaSignedQueryDefinitionV1 {
	return &NullableIrohaSignedQueryDefinitionV1{value: val, isSet: true}
}

func (v NullableIrohaSignedQueryDefinitionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIrohaSignedQueryDefinitionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


