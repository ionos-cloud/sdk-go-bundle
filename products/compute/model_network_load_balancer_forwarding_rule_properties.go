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

// checks if the NetworkLoadBalancerForwardingRuleProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkLoadBalancerForwardingRuleProperties{}

// NetworkLoadBalancerForwardingRuleProperties struct for NetworkLoadBalancerForwardingRuleProperties
type NetworkLoadBalancerForwardingRuleProperties struct {
	// The name of the Network Load Balancer forwarding rule.
	Name string `json:"name"`
	// Balancing algorithm
	Algorithm string `json:"algorithm"`
	// Balancing protocol
	Protocol string `json:"protocol"`
	// Listening (inbound) IP.
	ListenerIp string `json:"listenerIp"`
	// Listening (inbound) port number; valid range is 1 to 65535.
	ListenerPort int32                                         `json:"listenerPort"`
	HealthCheck  *NetworkLoadBalancerForwardingRuleHealthCheck `json:"healthCheck,omitempty"`
	// Array of items in the collection.
	Targets []NetworkLoadBalancerForwardingRuleTarget `json:"targets"`
}

// NewNetworkLoadBalancerForwardingRuleProperties instantiates a new NetworkLoadBalancerForwardingRuleProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancerForwardingRuleProperties(name string, algorithm string, protocol string, listenerIp string, listenerPort int32, targets []NetworkLoadBalancerForwardingRuleTarget) *NetworkLoadBalancerForwardingRuleProperties {
	this := NetworkLoadBalancerForwardingRuleProperties{}

	this.Name = name
	this.Algorithm = algorithm
	this.Protocol = protocol
	this.ListenerIp = listenerIp
	this.ListenerPort = listenerPort
	this.Targets = targets

	return &this
}

// NewNetworkLoadBalancerForwardingRulePropertiesWithDefaults instantiates a new NetworkLoadBalancerForwardingRuleProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerForwardingRulePropertiesWithDefaults() *NetworkLoadBalancerForwardingRuleProperties {
	this := NetworkLoadBalancerForwardingRuleProperties{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkLoadBalancerForwardingRuleProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkLoadBalancerForwardingRuleProperties) SetName(v string) {
	o.Name = v
}

// GetAlgorithm returns the Algorithm field value
func (o *NetworkLoadBalancerForwardingRuleProperties) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleProperties) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *NetworkLoadBalancerForwardingRuleProperties) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetProtocol returns the Protocol field value
func (o *NetworkLoadBalancerForwardingRuleProperties) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleProperties) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *NetworkLoadBalancerForwardingRuleProperties) SetProtocol(v string) {
	o.Protocol = v
}

// GetListenerIp returns the ListenerIp field value
func (o *NetworkLoadBalancerForwardingRuleProperties) GetListenerIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListenerIp
}

// GetListenerIpOk returns a tuple with the ListenerIp field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleProperties) GetListenerIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListenerIp, true
}

// SetListenerIp sets field value
func (o *NetworkLoadBalancerForwardingRuleProperties) SetListenerIp(v string) {
	o.ListenerIp = v
}

// GetListenerPort returns the ListenerPort field value
func (o *NetworkLoadBalancerForwardingRuleProperties) GetListenerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListenerPort
}

// GetListenerPortOk returns a tuple with the ListenerPort field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleProperties) GetListenerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListenerPort, true
}

// SetListenerPort sets field value
func (o *NetworkLoadBalancerForwardingRuleProperties) SetListenerPort(v int32) {
	o.ListenerPort = v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRuleProperties) GetHealthCheck() NetworkLoadBalancerForwardingRuleHealthCheck {
	if o == nil || IsNil(o.HealthCheck) {
		var ret NetworkLoadBalancerForwardingRuleHealthCheck
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleProperties) GetHealthCheckOk() (*NetworkLoadBalancerForwardingRuleHealthCheck, bool) {
	if o == nil || IsNil(o.HealthCheck) {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRuleProperties) HasHealthCheck() bool {
	if o != nil && !IsNil(o.HealthCheck) {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given NetworkLoadBalancerForwardingRuleHealthCheck and assigns it to the HealthCheck field.
func (o *NetworkLoadBalancerForwardingRuleProperties) SetHealthCheck(v NetworkLoadBalancerForwardingRuleHealthCheck) {
	o.HealthCheck = &v
}

// GetTargets returns the Targets field value
func (o *NetworkLoadBalancerForwardingRuleProperties) GetTargets() []NetworkLoadBalancerForwardingRuleTarget {
	if o == nil {
		var ret []NetworkLoadBalancerForwardingRuleTarget
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRuleProperties) GetTargetsOk() ([]NetworkLoadBalancerForwardingRuleTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *NetworkLoadBalancerForwardingRuleProperties) SetTargets(v []NetworkLoadBalancerForwardingRuleTarget) {
	o.Targets = v
}

func (o NetworkLoadBalancerForwardingRuleProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["protocol"] = o.Protocol
	toSerialize["listenerIp"] = o.ListenerIp
	toSerialize["listenerPort"] = o.ListenerPort
	if !IsNil(o.HealthCheck) {
		toSerialize["healthCheck"] = o.HealthCheck
	}
	toSerialize["targets"] = o.Targets
	return toSerialize, nil
}

type NullableNetworkLoadBalancerForwardingRuleProperties struct {
	value *NetworkLoadBalancerForwardingRuleProperties
	isSet bool
}

func (v NullableNetworkLoadBalancerForwardingRuleProperties) Get() *NetworkLoadBalancerForwardingRuleProperties {
	return v.value
}

func (v *NullableNetworkLoadBalancerForwardingRuleProperties) Set(val *NetworkLoadBalancerForwardingRuleProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerForwardingRuleProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerForwardingRuleProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerForwardingRuleProperties(val *NetworkLoadBalancerForwardingRuleProperties) *NullableNetworkLoadBalancerForwardingRuleProperties {
	return &NullableNetworkLoadBalancerForwardingRuleProperties{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerForwardingRuleProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerForwardingRuleProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
