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

// checks if the NetworkLoadBalancerForwardingRuleTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkLoadBalancerForwardingRuleTarget{}

// NetworkLoadBalancerForwardingRuleTarget struct for NetworkLoadBalancerForwardingRuleTarget
type NetworkLoadBalancerForwardingRuleTarget struct {
	// The IP of the balanced target VM.
	Ip string `json:"ip"`
	// The port of the balanced target service; valid range is 1 to 65535.
	Port int32 `json:"port"`
	// Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1. Targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best to assign weights in the middle of the range to leave room for later adjustments.
	Weight int32 `json:"weight"`
	// Proxy protocol version.
	ProxyProtocol *string                                             `json:"proxyProtocol,omitempty"`
	HealthCheck   *NetworkLoadBalancerForwardingRuleTargetHealthCheck `json:"healthCheck,omitempty"`
}

// NewNetworkLoadBalancerForwardingRuleTarget instantiates a new NetworkLoadBalancerForwardingRuleTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancerForwardingRuleTarget(ip string, port int32, weight int32) *NetworkLoadBalancerForwardingRuleTarget {
	this := NetworkLoadBalancerForwardingRuleTarget{}

	this.Ip = ip
	this.Port = port
	this.Weight = weight
	var proxyProtocol string = "none"
	this.ProxyProtocol = &proxyProtocol

	return &this
}

// NewNetworkLoadBalancerForwardingRuleTargetWithDefaults instantiates a new NetworkLoadBalancerForwardingRuleTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerForwardingRuleTargetWithDefaults() *NetworkLoadBalancerForwardingRuleTarget {
	this := NetworkLoadBalancerForwardingRuleTarget{}
	var proxyProtocol string = "none"
	this.ProxyProtocol = &proxyProtocol
	return &this
}

// GetIp returns the Ip field value
func (o *NetworkLoadBalancerForwardingRuleTarget) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *NetworkLoadBalancerForwardingRuleTarget) SetIp(v string) {
	o.Ip = v
}

// GetPort returns the Port field value
func (o *NetworkLoadBalancerForwardingRuleTarget) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NetworkLoadBalancerForwardingRuleTarget) SetPort(v int32) {
	o.Port = v
}

// GetWeight returns the Weight field value
func (o *NetworkLoadBalancerForwardingRuleTarget) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *NetworkLoadBalancerForwardingRuleTarget) SetWeight(v int32) {
	o.Weight = v
}

// GetProxyProtocol returns the ProxyProtocol field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRuleTarget) GetProxyProtocol() string {
	if o == nil || IsNil(o.ProxyProtocol) {
		var ret string
		return ret
	}
	return *o.ProxyProtocol
}

// GetProxyProtocolOk returns a tuple with the ProxyProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) GetProxyProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyProtocol) {
		return nil, false
	}
	return o.ProxyProtocol, true
}

// HasProxyProtocol returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) HasProxyProtocol() bool {
	if o != nil && !IsNil(o.ProxyProtocol) {
		return true
	}

	return false
}

// SetProxyProtocol gets a reference to the given string and assigns it to the ProxyProtocol field.
func (o *NetworkLoadBalancerForwardingRuleTarget) SetProxyProtocol(v string) {
	o.ProxyProtocol = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRuleTarget) GetHealthCheck() NetworkLoadBalancerForwardingRuleTargetHealthCheck {
	if o == nil || IsNil(o.HealthCheck) {
		var ret NetworkLoadBalancerForwardingRuleTargetHealthCheck
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) GetHealthCheckOk() (*NetworkLoadBalancerForwardingRuleTargetHealthCheck, bool) {
	if o == nil || IsNil(o.HealthCheck) {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleTarget) HasHealthCheck() bool {
	if o != nil && !IsNil(o.HealthCheck) {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given NetworkLoadBalancerForwardingRuleTargetHealthCheck and assigns it to the HealthCheck field.
func (o *NetworkLoadBalancerForwardingRuleTarget) SetHealthCheck(v NetworkLoadBalancerForwardingRuleTargetHealthCheck) {
	o.HealthCheck = &v
}

func (o NetworkLoadBalancerForwardingRuleTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkLoadBalancerForwardingRuleTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsZero(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsZero(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsZero(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.ProxyProtocol) {
		toSerialize["proxyProtocol"] = o.ProxyProtocol
	}
	if !IsNil(o.HealthCheck) {
		toSerialize["healthCheck"] = o.HealthCheck
	}
	return toSerialize, nil
}

type NullableNetworkLoadBalancerForwardingRuleTarget struct {
	value *NetworkLoadBalancerForwardingRuleTarget
	isSet bool
}

func (v NullableNetworkLoadBalancerForwardingRuleTarget) Get() *NetworkLoadBalancerForwardingRuleTarget {
	return v.value
}

func (v *NullableNetworkLoadBalancerForwardingRuleTarget) Set(val *NetworkLoadBalancerForwardingRuleTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerForwardingRuleTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerForwardingRuleTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerForwardingRuleTarget(val *NetworkLoadBalancerForwardingRuleTarget) *NullableNetworkLoadBalancerForwardingRuleTarget {
	return &NullableNetworkLoadBalancerForwardingRuleTarget{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerForwardingRuleTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerForwardingRuleTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
