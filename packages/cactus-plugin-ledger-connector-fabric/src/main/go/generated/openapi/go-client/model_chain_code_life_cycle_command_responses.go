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

// checks if the ChainCodeLifeCycleCommandResponses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChainCodeLifeCycleCommandResponses{}

// ChainCodeLifeCycleCommandResponses struct for ChainCodeLifeCycleCommandResponses
type ChainCodeLifeCycleCommandResponses struct {
	Packaging *SSHExecCommandResponse `json:"packaging,omitempty"`
	InstallList []SSHExecCommandResponse `json:"installList"`
	QueryInstalledList []SSHExecCommandResponse `json:"queryInstalledList"`
	ApproveForMyOrgList []SSHExecCommandResponse `json:"approveForMyOrgList"`
	Commit *SSHExecCommandResponse `json:"commit,omitempty"`
	QueryCommitted *SSHExecCommandResponse `json:"queryCommitted,omitempty"`
	Init *SSHExecCommandResponse `json:"init,omitempty"`
}

// NewChainCodeLifeCycleCommandResponses instantiates a new ChainCodeLifeCycleCommandResponses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChainCodeLifeCycleCommandResponses(installList []SSHExecCommandResponse, queryInstalledList []SSHExecCommandResponse, approveForMyOrgList []SSHExecCommandResponse) *ChainCodeLifeCycleCommandResponses {
	this := ChainCodeLifeCycleCommandResponses{}
	this.InstallList = installList
	this.QueryInstalledList = queryInstalledList
	this.ApproveForMyOrgList = approveForMyOrgList
	return &this
}

// NewChainCodeLifeCycleCommandResponsesWithDefaults instantiates a new ChainCodeLifeCycleCommandResponses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChainCodeLifeCycleCommandResponsesWithDefaults() *ChainCodeLifeCycleCommandResponses {
	this := ChainCodeLifeCycleCommandResponses{}
	return &this
}

// GetPackaging returns the Packaging field value if set, zero value otherwise.
func (o *ChainCodeLifeCycleCommandResponses) GetPackaging() SSHExecCommandResponse {
	if o == nil || IsNil(o.Packaging) {
		var ret SSHExecCommandResponse
		return ret
	}
	return *o.Packaging
}

// GetPackagingOk returns a tuple with the Packaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChainCodeLifeCycleCommandResponses) GetPackagingOk() (*SSHExecCommandResponse, bool) {
	if o == nil || IsNil(o.Packaging) {
		return nil, false
	}
	return o.Packaging, true
}

// HasPackaging returns a boolean if a field has been set.
func (o *ChainCodeLifeCycleCommandResponses) HasPackaging() bool {
	if o != nil && !IsNil(o.Packaging) {
		return true
	}

	return false
}

// SetPackaging gets a reference to the given SSHExecCommandResponse and assigns it to the Packaging field.
func (o *ChainCodeLifeCycleCommandResponses) SetPackaging(v SSHExecCommandResponse) {
	o.Packaging = &v
}

// GetInstallList returns the InstallList field value
func (o *ChainCodeLifeCycleCommandResponses) GetInstallList() []SSHExecCommandResponse {
	if o == nil {
		var ret []SSHExecCommandResponse
		return ret
	}

	return o.InstallList
}

// GetInstallListOk returns a tuple with the InstallList field value
// and a boolean to check if the value has been set.
func (o *ChainCodeLifeCycleCommandResponses) GetInstallListOk() ([]SSHExecCommandResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallList, true
}

// SetInstallList sets field value
func (o *ChainCodeLifeCycleCommandResponses) SetInstallList(v []SSHExecCommandResponse) {
	o.InstallList = v
}

// GetQueryInstalledList returns the QueryInstalledList field value
func (o *ChainCodeLifeCycleCommandResponses) GetQueryInstalledList() []SSHExecCommandResponse {
	if o == nil {
		var ret []SSHExecCommandResponse
		return ret
	}

	return o.QueryInstalledList
}

// GetQueryInstalledListOk returns a tuple with the QueryInstalledList field value
// and a boolean to check if the value has been set.
func (o *ChainCodeLifeCycleCommandResponses) GetQueryInstalledListOk() ([]SSHExecCommandResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryInstalledList, true
}

// SetQueryInstalledList sets field value
func (o *ChainCodeLifeCycleCommandResponses) SetQueryInstalledList(v []SSHExecCommandResponse) {
	o.QueryInstalledList = v
}

// GetApproveForMyOrgList returns the ApproveForMyOrgList field value
func (o *ChainCodeLifeCycleCommandResponses) GetApproveForMyOrgList() []SSHExecCommandResponse {
	if o == nil {
		var ret []SSHExecCommandResponse
		return ret
	}

	return o.ApproveForMyOrgList
}

// GetApproveForMyOrgListOk returns a tuple with the ApproveForMyOrgList field value
// and a boolean to check if the value has been set.
func (o *ChainCodeLifeCycleCommandResponses) GetApproveForMyOrgListOk() ([]SSHExecCommandResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproveForMyOrgList, true
}

// SetApproveForMyOrgList sets field value
func (o *ChainCodeLifeCycleCommandResponses) SetApproveForMyOrgList(v []SSHExecCommandResponse) {
	o.ApproveForMyOrgList = v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *ChainCodeLifeCycleCommandResponses) GetCommit() SSHExecCommandResponse {
	if o == nil || IsNil(o.Commit) {
		var ret SSHExecCommandResponse
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChainCodeLifeCycleCommandResponses) GetCommitOk() (*SSHExecCommandResponse, bool) {
	if o == nil || IsNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *ChainCodeLifeCycleCommandResponses) HasCommit() bool {
	if o != nil && !IsNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given SSHExecCommandResponse and assigns it to the Commit field.
func (o *ChainCodeLifeCycleCommandResponses) SetCommit(v SSHExecCommandResponse) {
	o.Commit = &v
}

// GetQueryCommitted returns the QueryCommitted field value if set, zero value otherwise.
func (o *ChainCodeLifeCycleCommandResponses) GetQueryCommitted() SSHExecCommandResponse {
	if o == nil || IsNil(o.QueryCommitted) {
		var ret SSHExecCommandResponse
		return ret
	}
	return *o.QueryCommitted
}

// GetQueryCommittedOk returns a tuple with the QueryCommitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChainCodeLifeCycleCommandResponses) GetQueryCommittedOk() (*SSHExecCommandResponse, bool) {
	if o == nil || IsNil(o.QueryCommitted) {
		return nil, false
	}
	return o.QueryCommitted, true
}

// HasQueryCommitted returns a boolean if a field has been set.
func (o *ChainCodeLifeCycleCommandResponses) HasQueryCommitted() bool {
	if o != nil && !IsNil(o.QueryCommitted) {
		return true
	}

	return false
}

// SetQueryCommitted gets a reference to the given SSHExecCommandResponse and assigns it to the QueryCommitted field.
func (o *ChainCodeLifeCycleCommandResponses) SetQueryCommitted(v SSHExecCommandResponse) {
	o.QueryCommitted = &v
}

// GetInit returns the Init field value if set, zero value otherwise.
func (o *ChainCodeLifeCycleCommandResponses) GetInit() SSHExecCommandResponse {
	if o == nil || IsNil(o.Init) {
		var ret SSHExecCommandResponse
		return ret
	}
	return *o.Init
}

// GetInitOk returns a tuple with the Init field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChainCodeLifeCycleCommandResponses) GetInitOk() (*SSHExecCommandResponse, bool) {
	if o == nil || IsNil(o.Init) {
		return nil, false
	}
	return o.Init, true
}

// HasInit returns a boolean if a field has been set.
func (o *ChainCodeLifeCycleCommandResponses) HasInit() bool {
	if o != nil && !IsNil(o.Init) {
		return true
	}

	return false
}

// SetInit gets a reference to the given SSHExecCommandResponse and assigns it to the Init field.
func (o *ChainCodeLifeCycleCommandResponses) SetInit(v SSHExecCommandResponse) {
	o.Init = &v
}

func (o ChainCodeLifeCycleCommandResponses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChainCodeLifeCycleCommandResponses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Packaging) {
		toSerialize["packaging"] = o.Packaging
	}
	toSerialize["installList"] = o.InstallList
	toSerialize["queryInstalledList"] = o.QueryInstalledList
	toSerialize["approveForMyOrgList"] = o.ApproveForMyOrgList
	if !IsNil(o.Commit) {
		toSerialize["commit"] = o.Commit
	}
	if !IsNil(o.QueryCommitted) {
		toSerialize["queryCommitted"] = o.QueryCommitted
	}
	if !IsNil(o.Init) {
		toSerialize["init"] = o.Init
	}
	return toSerialize, nil
}

type NullableChainCodeLifeCycleCommandResponses struct {
	value *ChainCodeLifeCycleCommandResponses
	isSet bool
}

func (v NullableChainCodeLifeCycleCommandResponses) Get() *ChainCodeLifeCycleCommandResponses {
	return v.value
}

func (v *NullableChainCodeLifeCycleCommandResponses) Set(val *ChainCodeLifeCycleCommandResponses) {
	v.value = val
	v.isSet = true
}

func (v NullableChainCodeLifeCycleCommandResponses) IsSet() bool {
	return v.isSet
}

func (v *NullableChainCodeLifeCycleCommandResponses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChainCodeLifeCycleCommandResponses(val *ChainCodeLifeCycleCommandResponses) *NullableChainCodeLifeCycleCommandResponses {
	return &NullableChainCodeLifeCycleCommandResponses{value: val, isSet: true}
}

func (v NullableChainCodeLifeCycleCommandResponses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChainCodeLifeCycleCommandResponses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


