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
)

// checks if the ShipFuel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipFuel{}

// ShipFuel Details of the ship's fuel tanks including how much fuel was consumed during the last transit or action.
type ShipFuel struct {
	// The current amount of fuel in the ship's tanks.
	Current int32 `json:"current"`
	// The maximum amount of fuel the ship's tanks can hold.
	Capacity int32 `json:"capacity"`
	Consumed *ShipFuelConsumed `json:"consumed,omitempty"`
}

// NewShipFuel instantiates a new ShipFuel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipFuel(current int32, capacity int32) *ShipFuel {
	this := ShipFuel{}
	this.Current = current
	this.Capacity = capacity
	return &this
}

// NewShipFuelWithDefaults instantiates a new ShipFuel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipFuelWithDefaults() *ShipFuel {
	this := ShipFuel{}
	return &this
}

// GetCurrent returns the Current field value
func (o *ShipFuel) GetCurrent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *ShipFuel) GetCurrentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *ShipFuel) SetCurrent(v int32) {
	o.Current = v
}

// GetCapacity returns the Capacity field value
func (o *ShipFuel) GetCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *ShipFuel) GetCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *ShipFuel) SetCapacity(v int32) {
	o.Capacity = v
}

// GetConsumed returns the Consumed field value if set, zero value otherwise.
func (o *ShipFuel) GetConsumed() ShipFuelConsumed {
	if o == nil || IsNil(o.Consumed) {
		var ret ShipFuelConsumed
		return ret
	}
	return *o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipFuel) GetConsumedOk() (*ShipFuelConsumed, bool) {
	if o == nil || IsNil(o.Consumed) {
		return nil, false
	}
	return o.Consumed, true
}

// HasConsumed returns a boolean if a field has been set.
func (o *ShipFuel) HasConsumed() bool {
	if o != nil && !IsNil(o.Consumed) {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given ShipFuelConsumed and assigns it to the Consumed field.
func (o *ShipFuel) SetConsumed(v ShipFuelConsumed) {
	o.Consumed = &v
}

func (o ShipFuel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipFuel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current"] = o.Current
	toSerialize["capacity"] = o.Capacity
	if !IsNil(o.Consumed) {
		toSerialize["consumed"] = o.Consumed
	}
	return toSerialize, nil
}

type NullableShipFuel struct {
	value *ShipFuel
	isSet bool
}

func (v NullableShipFuel) Get() *ShipFuel {
	return v.value
}

func (v *NullableShipFuel) Set(val *ShipFuel) {
	v.value = val
	v.isSet = true
}

func (v NullableShipFuel) IsSet() bool {
	return v.isSet
}

func (v *NullableShipFuel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipFuel(val *ShipFuel) *NullableShipFuel {
	return &NullableShipFuel{value: val, isSet: true}
}

func (v NullableShipFuel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipFuel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


