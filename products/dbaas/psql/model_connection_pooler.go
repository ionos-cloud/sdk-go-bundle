/*
 * IONOS DBaaS PostgreSQL REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional PostgreSQL database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psql

import (
	"encoding/json"
)

// checks if the ConnectionPooler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionPooler{}

// ConnectionPooler Configuration options for the connection pooler
type ConnectionPooler struct {
	Enabled  *bool     `json:"enabled,omitempty"`
	PoolMode *PoolMode `json:"poolMode,omitempty"`
}

// NewConnectionPooler instantiates a new ConnectionPooler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionPooler() *ConnectionPooler {
	this := ConnectionPooler{}

	return &this
}

// NewConnectionPoolerWithDefaults instantiates a new ConnectionPooler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionPoolerWithDefaults() *ConnectionPooler {
	this := ConnectionPooler{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ConnectionPooler) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionPooler) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ConnectionPooler) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ConnectionPooler) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPoolMode returns the PoolMode field value if set, zero value otherwise.
func (o *ConnectionPooler) GetPoolMode() PoolMode {
	if o == nil || IsNil(o.PoolMode) {
		var ret PoolMode
		return ret
	}
	return *o.PoolMode
}

// GetPoolModeOk returns a tuple with the PoolMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionPooler) GetPoolModeOk() (*PoolMode, bool) {
	if o == nil || IsNil(o.PoolMode) {
		return nil, false
	}
	return o.PoolMode, true
}

// HasPoolMode returns a boolean if a field has been set.
func (o *ConnectionPooler) HasPoolMode() bool {
	if o != nil && !IsNil(o.PoolMode) {
		return true
	}

	return false
}

// SetPoolMode gets a reference to the given PoolMode and assigns it to the PoolMode field.
func (o *ConnectionPooler) SetPoolMode(v PoolMode) {
	o.PoolMode = &v
}

func (o ConnectionPooler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.PoolMode) {
		toSerialize["poolMode"] = o.PoolMode
	}
	return toSerialize, nil
}

type NullableConnectionPooler struct {
	value *ConnectionPooler
	isSet bool
}

func (v NullableConnectionPooler) Get() *ConnectionPooler {
	return v.value
}

func (v *NullableConnectionPooler) Set(val *ConnectionPooler) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionPooler) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionPooler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionPooler(val *ConnectionPooler) *NullableConnectionPooler {
	return &NullableConnectionPooler{value: val, isSet: true}
}

func (v NullableConnectionPooler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionPooler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}