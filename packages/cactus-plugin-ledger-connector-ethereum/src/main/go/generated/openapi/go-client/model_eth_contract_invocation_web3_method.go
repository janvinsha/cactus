/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
	"fmt"
)

// EthContractInvocationWeb3Method the model 'EthContractInvocationWeb3Method'
type EthContractInvocationWeb3Method string

// List of EthContractInvocationWeb3Method
const (
	SEND EthContractInvocationWeb3Method = "send"
	CALL EthContractInvocationWeb3Method = "call"
	ENCODE_ABI EthContractInvocationWeb3Method = "encodeABI"
	ESTIMATE_GAS EthContractInvocationWeb3Method = "estimateGas"
)

// All allowed values of EthContractInvocationWeb3Method enum
var AllowedEthContractInvocationWeb3MethodEnumValues = []EthContractInvocationWeb3Method{
	"send",
	"call",
	"encodeABI",
	"estimateGas",
}

func (v *EthContractInvocationWeb3Method) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EthContractInvocationWeb3Method(value)
	for _, existing := range AllowedEthContractInvocationWeb3MethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EthContractInvocationWeb3Method", value)
}

// NewEthContractInvocationWeb3MethodFromValue returns a pointer to a valid EthContractInvocationWeb3Method
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEthContractInvocationWeb3MethodFromValue(v string) (*EthContractInvocationWeb3Method, error) {
	ev := EthContractInvocationWeb3Method(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EthContractInvocationWeb3Method: valid values are %v", v, AllowedEthContractInvocationWeb3MethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EthContractInvocationWeb3Method) IsValid() bool {
	for _, existing := range AllowedEthContractInvocationWeb3MethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EthContractInvocationWeb3Method value
func (v EthContractInvocationWeb3Method) Ptr() *EthContractInvocationWeb3Method {
	return &v
}

type NullableEthContractInvocationWeb3Method struct {
	value *EthContractInvocationWeb3Method
	isSet bool
}

func (v NullableEthContractInvocationWeb3Method) Get() *EthContractInvocationWeb3Method {
	return v.value
}

func (v *NullableEthContractInvocationWeb3Method) Set(val *EthContractInvocationWeb3Method) {
	v.value = val
	v.isSet = true
}

func (v NullableEthContractInvocationWeb3Method) IsSet() bool {
	return v.isSet
}

func (v *NullableEthContractInvocationWeb3Method) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthContractInvocationWeb3Method(val *EthContractInvocationWeb3Method) *NullableEthContractInvocationWeb3Method {
	return &NullableEthContractInvocationWeb3Method{value: val, isSet: true}
}

func (v NullableEthContractInvocationWeb3Method) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthContractInvocationWeb3Method) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

