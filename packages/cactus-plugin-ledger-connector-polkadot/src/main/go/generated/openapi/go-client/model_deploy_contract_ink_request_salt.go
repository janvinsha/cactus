/*
Hyperledger Cactus Plugin - Connector Polkadot

Can perform basic tasks on a Polkadot parachain

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-polkadot

import (
	"encoding/json"
	"fmt"
)

// DeployContractInkRequestSalt - struct for DeployContractInkRequestSalt
type DeployContractInkRequestSalt struct {
	String *string
}

// stringAsDeployContractInkRequestSalt is a convenience function that returns string wrapped in DeployContractInkRequestSalt
func StringAsDeployContractInkRequestSalt(v *string) DeployContractInkRequestSalt {
	return DeployContractInkRequestSalt{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeployContractInkRequestSalt) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DeployContractInkRequestSalt)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DeployContractInkRequestSalt)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeployContractInkRequestSalt) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeployContractInkRequestSalt) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDeployContractInkRequestSalt struct {
	value *DeployContractInkRequestSalt
	isSet bool
}

func (v NullableDeployContractInkRequestSalt) Get() *DeployContractInkRequestSalt {
	return v.value
}

func (v *NullableDeployContractInkRequestSalt) Set(val *DeployContractInkRequestSalt) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployContractInkRequestSalt) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployContractInkRequestSalt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployContractInkRequestSalt(val *DeployContractInkRequestSalt) *NullableDeployContractInkRequestSalt {
	return &NullableDeployContractInkRequestSalt{value: val, isSet: true}
}

func (v NullableDeployContractInkRequestSalt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployContractInkRequestSalt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


