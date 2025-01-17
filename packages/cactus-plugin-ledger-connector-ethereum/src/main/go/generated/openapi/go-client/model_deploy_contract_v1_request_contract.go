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

// DeployContractV1RequestContract - struct for DeployContractV1RequestContract
type DeployContractV1RequestContract struct {
	ContractJsonDefinition *ContractJsonDefinition
	ContractKeychainDefinition *ContractKeychainDefinition
}

// ContractJsonDefinitionAsDeployContractV1RequestContract is a convenience function that returns ContractJsonDefinition wrapped in DeployContractV1RequestContract
func ContractJsonDefinitionAsDeployContractV1RequestContract(v *ContractJsonDefinition) DeployContractV1RequestContract {
	return DeployContractV1RequestContract{
		ContractJsonDefinition: v,
	}
}

// ContractKeychainDefinitionAsDeployContractV1RequestContract is a convenience function that returns ContractKeychainDefinition wrapped in DeployContractV1RequestContract
func ContractKeychainDefinitionAsDeployContractV1RequestContract(v *ContractKeychainDefinition) DeployContractV1RequestContract {
	return DeployContractV1RequestContract{
		ContractKeychainDefinition: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeployContractV1RequestContract) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ContractJsonDefinition
	err = newStrictDecoder(data).Decode(&dst.ContractJsonDefinition)
	if err == nil {
		jsonContractJsonDefinition, _ := json.Marshal(dst.ContractJsonDefinition)
		if string(jsonContractJsonDefinition) == "{}" { // empty struct
			dst.ContractJsonDefinition = nil
		} else {
			match++
		}
	} else {
		dst.ContractJsonDefinition = nil
	}

	// try to unmarshal data into ContractKeychainDefinition
	err = newStrictDecoder(data).Decode(&dst.ContractKeychainDefinition)
	if err == nil {
		jsonContractKeychainDefinition, _ := json.Marshal(dst.ContractKeychainDefinition)
		if string(jsonContractKeychainDefinition) == "{}" { // empty struct
			dst.ContractKeychainDefinition = nil
		} else {
			match++
		}
	} else {
		dst.ContractKeychainDefinition = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ContractJsonDefinition = nil
		dst.ContractKeychainDefinition = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DeployContractV1RequestContract)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DeployContractV1RequestContract)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeployContractV1RequestContract) MarshalJSON() ([]byte, error) {
	if src.ContractJsonDefinition != nil {
		return json.Marshal(&src.ContractJsonDefinition)
	}

	if src.ContractKeychainDefinition != nil {
		return json.Marshal(&src.ContractKeychainDefinition)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeployContractV1RequestContract) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ContractJsonDefinition != nil {
		return obj.ContractJsonDefinition
	}

	if obj.ContractKeychainDefinition != nil {
		return obj.ContractKeychainDefinition
	}

	// all schemas are nil
	return nil
}

type NullableDeployContractV1RequestContract struct {
	value *DeployContractV1RequestContract
	isSet bool
}

func (v NullableDeployContractV1RequestContract) Get() *DeployContractV1RequestContract {
	return v.value
}

func (v *NullableDeployContractV1RequestContract) Set(val *DeployContractV1RequestContract) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployContractV1RequestContract) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployContractV1RequestContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployContractV1RequestContract(val *DeployContractV1RequestContract) *NullableDeployContractV1RequestContract {
	return &NullableDeployContractV1RequestContract{value: val, isSet: true}
}

func (v NullableDeployContractV1RequestContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployContractV1RequestContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


