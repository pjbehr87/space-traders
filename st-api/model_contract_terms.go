/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stapi

import (
	"encoding/json"
	"time"
)

// checks if the ContractTerms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractTerms{}

// ContractTerms Terms of the contract needed to fulfill it.
type ContractTerms struct {
	// The deadline for the contract.
	Deadline time.Time `json:"deadline"`
	Payment ContractPayment `json:"payment"`
	// The cargo that needs to be delivered to fulfill the contract.
	Deliver []ContractDeliverGood `json:"deliver,omitempty"`
}

// NewContractTerms instantiates a new ContractTerms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractTerms(deadline time.Time, payment ContractPayment) *ContractTerms {
	this := ContractTerms{}
	this.Deadline = deadline
	this.Payment = payment
	return &this
}

// NewContractTermsWithDefaults instantiates a new ContractTerms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractTermsWithDefaults() *ContractTerms {
	this := ContractTerms{}
	return &this
}

// GetDeadline returns the Deadline field value
func (o *ContractTerms) GetDeadline() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value
// and a boolean to check if the value has been set.
func (o *ContractTerms) GetDeadlineOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deadline, true
}

// SetDeadline sets field value
func (o *ContractTerms) SetDeadline(v time.Time) {
	o.Deadline = v
}

// GetPayment returns the Payment field value
func (o *ContractTerms) GetPayment() ContractPayment {
	if o == nil {
		var ret ContractPayment
		return ret
	}

	return o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value
// and a boolean to check if the value has been set.
func (o *ContractTerms) GetPaymentOk() (*ContractPayment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payment, true
}

// SetPayment sets field value
func (o *ContractTerms) SetPayment(v ContractPayment) {
	o.Payment = v
}

// GetDeliver returns the Deliver field value if set, zero value otherwise.
func (o *ContractTerms) GetDeliver() []ContractDeliverGood {
	if o == nil || IsNil(o.Deliver) {
		var ret []ContractDeliverGood
		return ret
	}
	return o.Deliver
}

// GetDeliverOk returns a tuple with the Deliver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractTerms) GetDeliverOk() ([]ContractDeliverGood, bool) {
	if o == nil || IsNil(o.Deliver) {
		return nil, false
	}
	return o.Deliver, true
}

// HasDeliver returns a boolean if a field has been set.
func (o *ContractTerms) HasDeliver() bool {
	if o != nil && !IsNil(o.Deliver) {
		return true
	}

	return false
}

// SetDeliver gets a reference to the given []ContractDeliverGood and assigns it to the Deliver field.
func (o *ContractTerms) SetDeliver(v []ContractDeliverGood) {
	o.Deliver = v
}

func (o ContractTerms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractTerms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deadline"] = o.Deadline
	toSerialize["payment"] = o.Payment
	if !IsNil(o.Deliver) {
		toSerialize["deliver"] = o.Deliver
	}
	return toSerialize, nil
}

type NullableContractTerms struct {
	value *ContractTerms
	isSet bool
}

func (v NullableContractTerms) Get() *ContractTerms {
	return v.value
}

func (v *NullableContractTerms) Set(val *ContractTerms) {
	v.value = val
	v.isSet = true
}

func (v NullableContractTerms) IsSet() bool {
	return v.isSet
}

func (v *NullableContractTerms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractTerms(val *ContractTerms) *NullableContractTerms {
	return &NullableContractTerms{value: val, isSet: true}
}

func (v NullableContractTerms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractTerms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


