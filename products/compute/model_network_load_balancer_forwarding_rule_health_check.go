/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// checks if the NetworkLoadBalancerForwardingRuleHealthCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkLoadBalancerForwardingRuleHealthCheck{}

// NetworkLoadBalancerForwardingRuleHealthCheck struct for NetworkLoadBalancerForwardingRuleHealthCheck
type NetworkLoadBalancerForwardingRuleHealthCheck struct {
	// The maximum time in milliseconds to wait for the client to acknowledge or send data; default is 50,000 (50 seconds).
	ClientTimeout *int32 `json:"clientTimeout,omitempty"`
	// The maximum time in milliseconds to wait for a connection attempt to a target to succeed; default is 5000 (five seconds).
	ConnectTimeout *int32 `json:"connectTimeout,omitempty"`
	// The maximum time in milliseconds that a target can remain inactive; default is 50,000 (50 seconds).
	TargetTimeout *int32 `json:"targetTimeout,omitempty"`
	// The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535 and default is three reconnection attempts.
	Retries *int32 `json:"retries,omitempty"`
}

// NewNetworkLoadBalancerForwardingRuleHealthCheck instantiates a new NetworkLoadBalancerForwardingRuleHealthCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancerForwardingRuleHealthCheck() *NetworkLoadBalancerForwardingRuleHealthCheck {
	this := NetworkLoadBalancerForwardingRuleHealthCheck{}

	return &this
}

// NewNetworkLoadBalancerForwardingRuleHealthCheckWithDefaults instantiates a new NetworkLoadBalancerForwardingRuleHealthCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerForwardingRuleHealthCheckWithDefaults() *NetworkLoadBalancerForwardingRuleHealthCheck {
	this := NetworkLoadBalancerForwardingRuleHealthCheck{}
	return &this
}

// GetClientTimeout returns the ClientTimeout field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetClientTimeout() int32 {
	if o == nil || IsNil(o.ClientTimeout) {
		var ret int32
		return ret
	}
	return *o.ClientTimeout
}

// GetClientTimeoutOk returns a tuple with the ClientTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetClientTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.ClientTimeout) {
		return nil, false
	}
	return o.ClientTimeout, true
}

// HasClientTimeout returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasClientTimeout() bool {
	if o != nil && !IsNil(o.ClientTimeout) {
		return true
	}

	return false
}

// SetClientTimeout gets a reference to the given int32 and assigns it to the ClientTimeout field.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetClientTimeout(v int32) {
	o.ClientTimeout = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetConnectTimeout() int32 {
	if o == nil || IsNil(o.ConnectTimeout) {
		var ret int32
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetConnectTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasConnectTimeout() bool {
	if o != nil && !IsNil(o.ConnectTimeout) {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given int32 and assigns it to the ConnectTimeout field.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetConnectTimeout(v int32) {
	o.ConnectTimeout = &v
}

// GetTargetTimeout returns the TargetTimeout field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetTargetTimeout() int32 {
	if o == nil || IsNil(o.TargetTimeout) {
		var ret int32
		return ret
	}
	return *o.TargetTimeout
}

// GetTargetTimeoutOk returns a tuple with the TargetTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetTargetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetTimeout) {
		return nil, false
	}
	return o.TargetTimeout, true
}

// HasTargetTimeout returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasTargetTimeout() bool {
	if o != nil && !IsNil(o.TargetTimeout) {
		return true
	}

	return false
}

// SetTargetTimeout gets a reference to the given int32 and assigns it to the TargetTimeout field.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetTargetTimeout(v int32) {
	o.TargetTimeout = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetRetries() int32 {
	if o == nil || IsNil(o.Retries) {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) GetRetriesOk() (*int32, bool) {
	if o == nil || IsNil(o.Retries) {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) HasRetries() bool {
	if o != nil && !IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *NetworkLoadBalancerForwardingRuleHealthCheck) SetRetries(v int32) {
	o.Retries = &v
}

func (o NetworkLoadBalancerForwardingRuleHealthCheck) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkLoadBalancerForwardingRuleHealthCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientTimeout) {
		toSerialize["clientTimeout"] = o.ClientTimeout
	}
	if !IsNil(o.ConnectTimeout) {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if !IsNil(o.TargetTimeout) {
		toSerialize["targetTimeout"] = o.TargetTimeout
	}
	if !IsNil(o.Retries) {
		toSerialize["retries"] = o.Retries
	}
	return toSerialize, nil
}

type NullableNetworkLoadBalancerForwardingRuleHealthCheck struct {
	value *NetworkLoadBalancerForwardingRuleHealthCheck
	isSet bool
}

func (v NullableNetworkLoadBalancerForwardingRuleHealthCheck) Get() *NetworkLoadBalancerForwardingRuleHealthCheck {
	return v.value
}

func (v *NullableNetworkLoadBalancerForwardingRuleHealthCheck) Set(val *NetworkLoadBalancerForwardingRuleHealthCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerForwardingRuleHealthCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerForwardingRuleHealthCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerForwardingRuleHealthCheck(val *NetworkLoadBalancerForwardingRuleHealthCheck) *NullableNetworkLoadBalancerForwardingRuleHealthCheck {
	return &NullableNetworkLoadBalancerForwardingRuleHealthCheck{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerForwardingRuleHealthCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerForwardingRuleHealthCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
