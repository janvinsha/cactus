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

// checks if the GenerateTransactionRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateTransactionRequestV1{}

// GenerateTransactionRequestV1 Request for generating transaction or query payload that can be signed on the client side.
type GenerateTransactionRequestV1 struct {
	Request GenerateTransactionRequestV1Request `json:"request"`
	BaseConfig *Iroha2BaseConfig `json:"baseConfig,omitempty"`
}

// NewGenerateTransactionRequestV1 instantiates a new GenerateTransactionRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateTransactionRequestV1(request GenerateTransactionRequestV1Request) *GenerateTransactionRequestV1 {
	this := GenerateTransactionRequestV1{}
	this.Request = request
	return &this
}

// NewGenerateTransactionRequestV1WithDefaults instantiates a new GenerateTransactionRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateTransactionRequestV1WithDefaults() *GenerateTransactionRequestV1 {
	this := GenerateTransactionRequestV1{}
	return &this
}

// GetRequest returns the Request field value
func (o *GenerateTransactionRequestV1) GetRequest() GenerateTransactionRequestV1Request {
	if o == nil {
		var ret GenerateTransactionRequestV1Request
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *GenerateTransactionRequestV1) GetRequestOk() (*GenerateTransactionRequestV1Request, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *GenerateTransactionRequestV1) SetRequest(v GenerateTransactionRequestV1Request) {
	o.Request = v
}

// GetBaseConfig returns the BaseConfig field value if set, zero value otherwise.
func (o *GenerateTransactionRequestV1) GetBaseConfig() Iroha2BaseConfig {
	if o == nil || IsNil(o.BaseConfig) {
		var ret Iroha2BaseConfig
		return ret
	}
	return *o.BaseConfig
}

// GetBaseConfigOk returns a tuple with the BaseConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateTransactionRequestV1) GetBaseConfigOk() (*Iroha2BaseConfig, bool) {
	if o == nil || IsNil(o.BaseConfig) {
		return nil, false
	}
	return o.BaseConfig, true
}

// HasBaseConfig returns a boolean if a field has been set.
func (o *GenerateTransactionRequestV1) HasBaseConfig() bool {
	if o != nil && !IsNil(o.BaseConfig) {
		return true
	}

	return false
}

// SetBaseConfig gets a reference to the given Iroha2BaseConfig and assigns it to the BaseConfig field.
func (o *GenerateTransactionRequestV1) SetBaseConfig(v Iroha2BaseConfig) {
	o.BaseConfig = &v
}

func (o GenerateTransactionRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateTransactionRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["request"] = o.Request
	if !IsNil(o.BaseConfig) {
		toSerialize["baseConfig"] = o.BaseConfig
	}
	return toSerialize, nil
}

type NullableGenerateTransactionRequestV1 struct {
	value *GenerateTransactionRequestV1
	isSet bool
}

func (v NullableGenerateTransactionRequestV1) Get() *GenerateTransactionRequestV1 {
	return v.value
}

func (v *NullableGenerateTransactionRequestV1) Set(val *GenerateTransactionRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateTransactionRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateTransactionRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateTransactionRequestV1(val *GenerateTransactionRequestV1) *NullableGenerateTransactionRequestV1 {
	return &NullableGenerateTransactionRequestV1{value: val, isSet: true}
}

func (v NullableGenerateTransactionRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateTransactionRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


