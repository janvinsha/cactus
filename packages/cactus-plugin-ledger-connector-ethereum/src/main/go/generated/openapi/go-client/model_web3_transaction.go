/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the Web3Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3Transaction{}

// Web3Transaction struct for Web3Transaction
type Web3Transaction struct {
	Hash string `json:"hash"`
	Nonce string `json:"nonce"`
	BlockHash string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From string `json:"from"`
	To string `json:"to"`
	Value string `json:"value"`
	GasPrice string `json:"gasPrice"`
	Gas string `json:"gas"`
	Input string `json:"input"`
	Type string `json:"type"`
	ChainId string `json:"chainId"`
	V *string `json:"v,omitempty"`
	R *string `json:"r,omitempty"`
	S *string `json:"s,omitempty"`
}

// NewWeb3Transaction instantiates a new Web3Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3Transaction(hash string, nonce string, blockHash string, blockNumber string, transactionIndex string, from string, to string, value string, gasPrice string, gas string, input string, type_ string, chainId string) *Web3Transaction {
	this := Web3Transaction{}
	this.Hash = hash
	this.Nonce = nonce
	this.BlockHash = blockHash
	this.BlockNumber = blockNumber
	this.TransactionIndex = transactionIndex
	this.From = from
	this.To = to
	this.Value = value
	this.GasPrice = gasPrice
	this.Gas = gas
	this.Input = input
	this.Type = type_
	this.ChainId = chainId
	return &this
}

// NewWeb3TransactionWithDefaults instantiates a new Web3Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3TransactionWithDefaults() *Web3Transaction {
	this := Web3Transaction{}
	return &this
}

// GetHash returns the Hash field value
func (o *Web3Transaction) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Web3Transaction) SetHash(v string) {
	o.Hash = v
}

// GetNonce returns the Nonce field value
func (o *Web3Transaction) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *Web3Transaction) SetNonce(v string) {
	o.Nonce = v
}

// GetBlockHash returns the BlockHash field value
func (o *Web3Transaction) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *Web3Transaction) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockNumber returns the BlockNumber field value
func (o *Web3Transaction) GetBlockNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetBlockNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockNumber, true
}

// SetBlockNumber sets field value
func (o *Web3Transaction) SetBlockNumber(v string) {
	o.BlockNumber = v
}

// GetTransactionIndex returns the TransactionIndex field value
func (o *Web3Transaction) GetTransactionIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetTransactionIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIndex, true
}

// SetTransactionIndex sets field value
func (o *Web3Transaction) SetTransactionIndex(v string) {
	o.TransactionIndex = v
}

// GetFrom returns the From field value
func (o *Web3Transaction) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *Web3Transaction) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *Web3Transaction) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *Web3Transaction) SetTo(v string) {
	o.To = v
}

// GetValue returns the Value field value
func (o *Web3Transaction) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Web3Transaction) SetValue(v string) {
	o.Value = v
}

// GetGasPrice returns the GasPrice field value
func (o *Web3Transaction) GetGasPrice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetGasPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GasPrice, true
}

// SetGasPrice sets field value
func (o *Web3Transaction) SetGasPrice(v string) {
	o.GasPrice = v
}

// GetGas returns the Gas field value
func (o *Web3Transaction) GetGas() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gas
}

// GetGasOk returns a tuple with the Gas field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetGasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gas, true
}

// SetGas sets field value
func (o *Web3Transaction) SetGas(v string) {
	o.Gas = v
}

// GetInput returns the Input field value
func (o *Web3Transaction) GetInput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Input
}

// GetInputOk returns a tuple with the Input field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetInputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Input, true
}

// SetInput sets field value
func (o *Web3Transaction) SetInput(v string) {
	o.Input = v
}

// GetType returns the Type field value
func (o *Web3Transaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Web3Transaction) SetType(v string) {
	o.Type = v
}

// GetChainId returns the ChainId field value
func (o *Web3Transaction) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *Web3Transaction) SetChainId(v string) {
	o.ChainId = v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *Web3Transaction) GetV() string {
	if o == nil || IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetVOk() (*string, bool) {
	if o == nil || IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *Web3Transaction) HasV() bool {
	if o != nil && !IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *Web3Transaction) SetV(v string) {
	o.V = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *Web3Transaction) GetR() string {
	if o == nil || IsNil(o.R) {
		var ret string
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetROk() (*string, bool) {
	if o == nil || IsNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *Web3Transaction) HasR() bool {
	if o != nil && !IsNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *Web3Transaction) SetR(v string) {
	o.R = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Web3Transaction) GetS() string {
	if o == nil || IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3Transaction) GetSOk() (*string, bool) {
	if o == nil || IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *Web3Transaction) HasS() bool {
	if o != nil && !IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Web3Transaction) SetS(v string) {
	o.S = &v
}

func (o Web3Transaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["nonce"] = o.Nonce
	toSerialize["blockHash"] = o.BlockHash
	toSerialize["blockNumber"] = o.BlockNumber
	toSerialize["transactionIndex"] = o.TransactionIndex
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	toSerialize["value"] = o.Value
	toSerialize["gasPrice"] = o.GasPrice
	toSerialize["gas"] = o.Gas
	toSerialize["input"] = o.Input
	toSerialize["type"] = o.Type
	toSerialize["chainId"] = o.ChainId
	if !IsNil(o.V) {
		toSerialize["v"] = o.V
	}
	if !IsNil(o.R) {
		toSerialize["r"] = o.R
	}
	if !IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	return toSerialize, nil
}

type NullableWeb3Transaction struct {
	value *Web3Transaction
	isSet bool
}

func (v NullableWeb3Transaction) Get() *Web3Transaction {
	return v.value
}

func (v *NullableWeb3Transaction) Set(val *Web3Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3Transaction) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3Transaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3Transaction(val *Web3Transaction) *NullableWeb3Transaction {
	return &NullableWeb3Transaction{value: val, isSet: true}
}

func (v NullableWeb3Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3Transaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


