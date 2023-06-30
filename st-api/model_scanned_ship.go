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

// checks if the ScannedShip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannedShip{}

// ScannedShip The ship that was scanned. Details include information about the ship that could be detected by the scanner.
type ScannedShip struct {
	// The globally unique identifier of the ship.
	Symbol string `json:"symbol"`
	Registration ShipRegistration `json:"registration"`
	Nav ShipNav `json:"nav"`
	Frame *ScannedShipFrame `json:"frame,omitempty"`
	Reactor *ScannedShipReactor `json:"reactor,omitempty"`
	Engine ScannedShipEngine `json:"engine"`
	// List of mounts installed in the ship.
	Mounts []ScannedShipMountsInner `json:"mounts,omitempty"`
}

// NewScannedShip instantiates a new ScannedShip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannedShip(symbol string, registration ShipRegistration, nav ShipNav, engine ScannedShipEngine) *ScannedShip {
	this := ScannedShip{}
	this.Symbol = symbol
	this.Registration = registration
	this.Nav = nav
	this.Engine = engine
	return &this
}

// NewScannedShipWithDefaults instantiates a new ScannedShip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannedShipWithDefaults() *ScannedShip {
	this := ScannedShip{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ScannedShip) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ScannedShip) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ScannedShip) SetSymbol(v string) {
	o.Symbol = v
}

// GetRegistration returns the Registration field value
func (o *ScannedShip) GetRegistration() ShipRegistration {
	if o == nil {
		var ret ShipRegistration
		return ret
	}

	return o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value
// and a boolean to check if the value has been set.
func (o *ScannedShip) GetRegistrationOk() (*ShipRegistration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Registration, true
}

// SetRegistration sets field value
func (o *ScannedShip) SetRegistration(v ShipRegistration) {
	o.Registration = v
}

// GetNav returns the Nav field value
func (o *ScannedShip) GetNav() ShipNav {
	if o == nil {
		var ret ShipNav
		return ret
	}

	return o.Nav
}

// GetNavOk returns a tuple with the Nav field value
// and a boolean to check if the value has been set.
func (o *ScannedShip) GetNavOk() (*ShipNav, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nav, true
}

// SetNav sets field value
func (o *ScannedShip) SetNav(v ShipNav) {
	o.Nav = v
}

// GetFrame returns the Frame field value if set, zero value otherwise.
func (o *ScannedShip) GetFrame() ScannedShipFrame {
	if o == nil || IsNil(o.Frame) {
		var ret ScannedShipFrame
		return ret
	}
	return *o.Frame
}

// GetFrameOk returns a tuple with the Frame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedShip) GetFrameOk() (*ScannedShipFrame, bool) {
	if o == nil || IsNil(o.Frame) {
		return nil, false
	}
	return o.Frame, true
}

// HasFrame returns a boolean if a field has been set.
func (o *ScannedShip) HasFrame() bool {
	if o != nil && !IsNil(o.Frame) {
		return true
	}

	return false
}

// SetFrame gets a reference to the given ScannedShipFrame and assigns it to the Frame field.
func (o *ScannedShip) SetFrame(v ScannedShipFrame) {
	o.Frame = &v
}

// GetReactor returns the Reactor field value if set, zero value otherwise.
func (o *ScannedShip) GetReactor() ScannedShipReactor {
	if o == nil || IsNil(o.Reactor) {
		var ret ScannedShipReactor
		return ret
	}
	return *o.Reactor
}

// GetReactorOk returns a tuple with the Reactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedShip) GetReactorOk() (*ScannedShipReactor, bool) {
	if o == nil || IsNil(o.Reactor) {
		return nil, false
	}
	return o.Reactor, true
}

// HasReactor returns a boolean if a field has been set.
func (o *ScannedShip) HasReactor() bool {
	if o != nil && !IsNil(o.Reactor) {
		return true
	}

	return false
}

// SetReactor gets a reference to the given ScannedShipReactor and assigns it to the Reactor field.
func (o *ScannedShip) SetReactor(v ScannedShipReactor) {
	o.Reactor = &v
}

// GetEngine returns the Engine field value
func (o *ScannedShip) GetEngine() ScannedShipEngine {
	if o == nil {
		var ret ScannedShipEngine
		return ret
	}

	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *ScannedShip) GetEngineOk() (*ScannedShipEngine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value
func (o *ScannedShip) SetEngine(v ScannedShipEngine) {
	o.Engine = v
}

// GetMounts returns the Mounts field value if set, zero value otherwise.
func (o *ScannedShip) GetMounts() []ScannedShipMountsInner {
	if o == nil || IsNil(o.Mounts) {
		var ret []ScannedShipMountsInner
		return ret
	}
	return o.Mounts
}

// GetMountsOk returns a tuple with the Mounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedShip) GetMountsOk() ([]ScannedShipMountsInner, bool) {
	if o == nil || IsNil(o.Mounts) {
		return nil, false
	}
	return o.Mounts, true
}

// HasMounts returns a boolean if a field has been set.
func (o *ScannedShip) HasMounts() bool {
	if o != nil && !IsNil(o.Mounts) {
		return true
	}

	return false
}

// SetMounts gets a reference to the given []ScannedShipMountsInner and assigns it to the Mounts field.
func (o *ScannedShip) SetMounts(v []ScannedShipMountsInner) {
	o.Mounts = v
}

func (o ScannedShip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannedShip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["registration"] = o.Registration
	toSerialize["nav"] = o.Nav
	if !IsNil(o.Frame) {
		toSerialize["frame"] = o.Frame
	}
	if !IsNil(o.Reactor) {
		toSerialize["reactor"] = o.Reactor
	}
	toSerialize["engine"] = o.Engine
	if !IsNil(o.Mounts) {
		toSerialize["mounts"] = o.Mounts
	}
	return toSerialize, nil
}

type NullableScannedShip struct {
	value *ScannedShip
	isSet bool
}

func (v NullableScannedShip) Get() *ScannedShip {
	return v.value
}

func (v *NullableScannedShip) Set(val *ScannedShip) {
	v.value = val
	v.isSet = true
}

func (v NullableScannedShip) IsSet() bool {
	return v.isSet
}

func (v *NullableScannedShip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannedShip(val *ScannedShip) *NullableScannedShip {
	return &NullableScannedShip{value: val, isSet: true}
}

func (v NullableScannedShip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannedShip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


