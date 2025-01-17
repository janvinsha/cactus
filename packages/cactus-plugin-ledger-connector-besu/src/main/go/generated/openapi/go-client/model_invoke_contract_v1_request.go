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

// checks if the InvokeContractV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvokeContractV1Request{}

// InvokeContractV1Request struct for InvokeContractV1Request
type InvokeContractV1Request struct {
	ContractName string `json:"contractName"`
	SigningCredential Web3SigningCredential `json:"signingCredential"`
	InvocationType EthContractInvocationType `json:"invocationType"`
	// The name of the contract method to invoke.
	MethodName string `json:"methodName"`
	// The list of arguments to pass in to the contract method being invoked.
	Params []interface{} `json:"params"`
	// The application binary interface of the solidity contract, optional parameter
	ContractAbi []interface{} `json:"contractAbi,omitempty"`
	// Address of the solidity contract, optional parameter
	ContractAddress *string `json:"contractAddress,omitempty"`
	Value *Web3BlockHeaderTimestamp `json:"value,omitempty"`
	Gas *Web3BlockHeaderTimestamp `json:"gas,omitempty"`
	GasPrice *Web3BlockHeaderTimestamp `json:"gasPrice,omitempty"`
	Nonce *float32 `json:"nonce,omitempty"`
	// The amount of milliseconds to wait for a transaction receipt beforegiving up and crashing. Only has any effect if the invocation type is SEND
	TimeoutMs *float32 `json:"timeoutMs,omitempty"`
	// The keychainId for retrieve the contracts json.
	KeychainId *string `json:"keychainId,omitempty"`
	PrivateTransactionConfig *BesuPrivateTransactionConfig `json:"privateTransactionConfig,omitempty"`
}

// NewInvokeContractV1Request instantiates a new InvokeContractV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvokeContractV1Request(contractName string, signingCredential Web3SigningCredential, invocationType EthContractInvocationType, methodName string, params []interface{}) *InvokeContractV1Request {
	this := InvokeContractV1Request{}
	this.ContractName = contractName
	this.SigningCredential = signingCredential
	this.InvocationType = invocationType
	this.MethodName = methodName
	this.Params = params
	var timeoutMs float32 = 60000
	this.TimeoutMs = &timeoutMs
	return &this
}

// NewInvokeContractV1RequestWithDefaults instantiates a new InvokeContractV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvokeContractV1RequestWithDefaults() *InvokeContractV1Request {
	this := InvokeContractV1Request{}
	var timeoutMs float32 = 60000
	this.TimeoutMs = &timeoutMs
	return &this
}

// GetContractName returns the ContractName field value
func (o *InvokeContractV1Request) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *InvokeContractV1Request) SetContractName(v string) {
	o.ContractName = v
}

// GetSigningCredential returns the SigningCredential field value
func (o *InvokeContractV1Request) GetSigningCredential() Web3SigningCredential {
	if o == nil {
		var ret Web3SigningCredential
		return ret
	}

	return o.SigningCredential
}

// GetSigningCredentialOk returns a tuple with the SigningCredential field value
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetSigningCredentialOk() (*Web3SigningCredential, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningCredential, true
}

// SetSigningCredential sets field value
func (o *InvokeContractV1Request) SetSigningCredential(v Web3SigningCredential) {
	o.SigningCredential = v
}

// GetInvocationType returns the InvocationType field value
func (o *InvokeContractV1Request) GetInvocationType() EthContractInvocationType {
	if o == nil {
		var ret EthContractInvocationType
		return ret
	}

	return o.InvocationType
}

// GetInvocationTypeOk returns a tuple with the InvocationType field value
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetInvocationTypeOk() (*EthContractInvocationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvocationType, true
}

// SetInvocationType sets field value
func (o *InvokeContractV1Request) SetInvocationType(v EthContractInvocationType) {
	o.InvocationType = v
}

// GetMethodName returns the MethodName field value
func (o *InvokeContractV1Request) GetMethodName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MethodName
}

// GetMethodNameOk returns a tuple with the MethodName field value
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetMethodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MethodName, true
}

// SetMethodName sets field value
func (o *InvokeContractV1Request) SetMethodName(v string) {
	o.MethodName = v
}

// GetParams returns the Params field value
func (o *InvokeContractV1Request) GetParams() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetParamsOk() ([]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Params, true
}

// SetParams sets field value
func (o *InvokeContractV1Request) SetParams(v []interface{}) {
	o.Params = v
}

// GetContractAbi returns the ContractAbi field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetContractAbi() []interface{} {
	if o == nil || IsNil(o.ContractAbi) {
		var ret []interface{}
		return ret
	}
	return o.ContractAbi
}

// GetContractAbiOk returns a tuple with the ContractAbi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetContractAbiOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ContractAbi) {
		return nil, false
	}
	return o.ContractAbi, true
}

// HasContractAbi returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasContractAbi() bool {
	if o != nil && !IsNil(o.ContractAbi) {
		return true
	}

	return false
}

// SetContractAbi gets a reference to the given []interface{} and assigns it to the ContractAbi field.
func (o *InvokeContractV1Request) SetContractAbi(v []interface{}) {
	o.ContractAbi = v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetContractAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasContractAddress() bool {
	if o != nil && !IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *InvokeContractV1Request) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetValue() Web3BlockHeaderTimestamp {
	if o == nil || IsNil(o.Value) {
		var ret Web3BlockHeaderTimestamp
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetValueOk() (*Web3BlockHeaderTimestamp, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Web3BlockHeaderTimestamp and assigns it to the Value field.
func (o *InvokeContractV1Request) SetValue(v Web3BlockHeaderTimestamp) {
	o.Value = &v
}

// GetGas returns the Gas field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetGas() Web3BlockHeaderTimestamp {
	if o == nil || IsNil(o.Gas) {
		var ret Web3BlockHeaderTimestamp
		return ret
	}
	return *o.Gas
}

// GetGasOk returns a tuple with the Gas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetGasOk() (*Web3BlockHeaderTimestamp, bool) {
	if o == nil || IsNil(o.Gas) {
		return nil, false
	}
	return o.Gas, true
}

// HasGas returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasGas() bool {
	if o != nil && !IsNil(o.Gas) {
		return true
	}

	return false
}

// SetGas gets a reference to the given Web3BlockHeaderTimestamp and assigns it to the Gas field.
func (o *InvokeContractV1Request) SetGas(v Web3BlockHeaderTimestamp) {
	o.Gas = &v
}

// GetGasPrice returns the GasPrice field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetGasPrice() Web3BlockHeaderTimestamp {
	if o == nil || IsNil(o.GasPrice) {
		var ret Web3BlockHeaderTimestamp
		return ret
	}
	return *o.GasPrice
}

// GetGasPriceOk returns a tuple with the GasPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetGasPriceOk() (*Web3BlockHeaderTimestamp, bool) {
	if o == nil || IsNil(o.GasPrice) {
		return nil, false
	}
	return o.GasPrice, true
}

// HasGasPrice returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasGasPrice() bool {
	if o != nil && !IsNil(o.GasPrice) {
		return true
	}

	return false
}

// SetGasPrice gets a reference to the given Web3BlockHeaderTimestamp and assigns it to the GasPrice field.
func (o *InvokeContractV1Request) SetGasPrice(v Web3BlockHeaderTimestamp) {
	o.GasPrice = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetNonce() float32 {
	if o == nil || IsNil(o.Nonce) {
		var ret float32
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetNonceOk() (*float32, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given float32 and assigns it to the Nonce field.
func (o *InvokeContractV1Request) SetNonce(v float32) {
	o.Nonce = &v
}

// GetTimeoutMs returns the TimeoutMs field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetTimeoutMs() float32 {
	if o == nil || IsNil(o.TimeoutMs) {
		var ret float32
		return ret
	}
	return *o.TimeoutMs
}

// GetTimeoutMsOk returns a tuple with the TimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetTimeoutMsOk() (*float32, bool) {
	if o == nil || IsNil(o.TimeoutMs) {
		return nil, false
	}
	return o.TimeoutMs, true
}

// HasTimeoutMs returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasTimeoutMs() bool {
	if o != nil && !IsNil(o.TimeoutMs) {
		return true
	}

	return false
}

// SetTimeoutMs gets a reference to the given float32 and assigns it to the TimeoutMs field.
func (o *InvokeContractV1Request) SetTimeoutMs(v float32) {
	o.TimeoutMs = &v
}

// GetKeychainId returns the KeychainId field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetKeychainId() string {
	if o == nil || IsNil(o.KeychainId) {
		var ret string
		return ret
	}
	return *o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetKeychainIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeychainId) {
		return nil, false
	}
	return o.KeychainId, true
}

// HasKeychainId returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasKeychainId() bool {
	if o != nil && !IsNil(o.KeychainId) {
		return true
	}

	return false
}

// SetKeychainId gets a reference to the given string and assigns it to the KeychainId field.
func (o *InvokeContractV1Request) SetKeychainId(v string) {
	o.KeychainId = &v
}

// GetPrivateTransactionConfig returns the PrivateTransactionConfig field value if set, zero value otherwise.
func (o *InvokeContractV1Request) GetPrivateTransactionConfig() BesuPrivateTransactionConfig {
	if o == nil || IsNil(o.PrivateTransactionConfig) {
		var ret BesuPrivateTransactionConfig
		return ret
	}
	return *o.PrivateTransactionConfig
}

// GetPrivateTransactionConfigOk returns a tuple with the PrivateTransactionConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvokeContractV1Request) GetPrivateTransactionConfigOk() (*BesuPrivateTransactionConfig, bool) {
	if o == nil || IsNil(o.PrivateTransactionConfig) {
		return nil, false
	}
	return o.PrivateTransactionConfig, true
}

// HasPrivateTransactionConfig returns a boolean if a field has been set.
func (o *InvokeContractV1Request) HasPrivateTransactionConfig() bool {
	if o != nil && !IsNil(o.PrivateTransactionConfig) {
		return true
	}

	return false
}

// SetPrivateTransactionConfig gets a reference to the given BesuPrivateTransactionConfig and assigns it to the PrivateTransactionConfig field.
func (o *InvokeContractV1Request) SetPrivateTransactionConfig(v BesuPrivateTransactionConfig) {
	o.PrivateTransactionConfig = &v
}

func (o InvokeContractV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvokeContractV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contractName"] = o.ContractName
	toSerialize["signingCredential"] = o.SigningCredential
	toSerialize["invocationType"] = o.InvocationType
	toSerialize["methodName"] = o.MethodName
	toSerialize["params"] = o.Params
	if !IsNil(o.ContractAbi) {
		toSerialize["contractAbi"] = o.ContractAbi
	}
	if !IsNil(o.ContractAddress) {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Gas) {
		toSerialize["gas"] = o.Gas
	}
	if !IsNil(o.GasPrice) {
		toSerialize["gasPrice"] = o.GasPrice
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.TimeoutMs) {
		toSerialize["timeoutMs"] = o.TimeoutMs
	}
	if !IsNil(o.KeychainId) {
		toSerialize["keychainId"] = o.KeychainId
	}
	if !IsNil(o.PrivateTransactionConfig) {
		toSerialize["privateTransactionConfig"] = o.PrivateTransactionConfig
	}
	return toSerialize, nil
}

type NullableInvokeContractV1Request struct {
	value *InvokeContractV1Request
	isSet bool
}

func (v NullableInvokeContractV1Request) Get() *InvokeContractV1Request {
	return v.value
}

func (v *NullableInvokeContractV1Request) Set(val *InvokeContractV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableInvokeContractV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableInvokeContractV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvokeContractV1Request(val *InvokeContractV1Request) *NullableInvokeContractV1Request {
	return &NullableInvokeContractV1Request{value: val, isSet: true}
}

func (v NullableInvokeContractV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvokeContractV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


