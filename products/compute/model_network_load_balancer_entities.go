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

// checks if the NetworkLoadBalancerEntities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkLoadBalancerEntities{}

// NetworkLoadBalancerEntities struct for NetworkLoadBalancerEntities
type NetworkLoadBalancerEntities struct {
	Flowlogs        *FlowLogs                           `json:"flowlogs,omitempty"`
	Forwardingrules *NetworkLoadBalancerForwardingRules `json:"forwardingrules,omitempty"`
}

// NewNetworkLoadBalancerEntities instantiates a new NetworkLoadBalancerEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancerEntities() *NetworkLoadBalancerEntities {
	this := NetworkLoadBalancerEntities{}

	return &this
}

// NewNetworkLoadBalancerEntitiesWithDefaults instantiates a new NetworkLoadBalancerEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerEntitiesWithDefaults() *NetworkLoadBalancerEntities {
	this := NetworkLoadBalancerEntities{}
	return &this
}

// GetFlowlogs returns the Flowlogs field value if set, zero value otherwise.
func (o *NetworkLoadBalancerEntities) GetFlowlogs() FlowLogs {
	if o == nil || IsNil(o.Flowlogs) {
		var ret FlowLogs
		return ret
	}
	return *o.Flowlogs
}

// GetFlowlogsOk returns a tuple with the Flowlogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerEntities) GetFlowlogsOk() (*FlowLogs, bool) {
	if o == nil || IsNil(o.Flowlogs) {
		return nil, false
	}
	return o.Flowlogs, true
}

// HasFlowlogs returns a boolean if a field has been set.
func (o *NetworkLoadBalancerEntities) HasFlowlogs() bool {
	if o != nil && !IsNil(o.Flowlogs) {
		return true
	}

	return false
}

// SetFlowlogs gets a reference to the given FlowLogs and assigns it to the Flowlogs field.
func (o *NetworkLoadBalancerEntities) SetFlowlogs(v FlowLogs) {
	o.Flowlogs = &v
}

// GetForwardingrules returns the Forwardingrules field value if set, zero value otherwise.
func (o *NetworkLoadBalancerEntities) GetForwardingrules() NetworkLoadBalancerForwardingRules {
	if o == nil || IsNil(o.Forwardingrules) {
		var ret NetworkLoadBalancerForwardingRules
		return ret
	}
	return *o.Forwardingrules
}

// GetForwardingrulesOk returns a tuple with the Forwardingrules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerEntities) GetForwardingrulesOk() (*NetworkLoadBalancerForwardingRules, bool) {
	if o == nil || IsNil(o.Forwardingrules) {
		return nil, false
	}
	return o.Forwardingrules, true
}

// HasForwardingrules returns a boolean if a field has been set.
func (o *NetworkLoadBalancerEntities) HasForwardingrules() bool {
	if o != nil && !IsNil(o.Forwardingrules) {
		return true
	}

	return false
}

// SetForwardingrules gets a reference to the given NetworkLoadBalancerForwardingRules and assigns it to the Forwardingrules field.
func (o *NetworkLoadBalancerEntities) SetForwardingrules(v NetworkLoadBalancerForwardingRules) {
	o.Forwardingrules = &v
}

func (o NetworkLoadBalancerEntities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flowlogs) {
		toSerialize["flowlogs"] = o.Flowlogs
	}
	if !IsNil(o.Forwardingrules) {
		toSerialize["forwardingrules"] = o.Forwardingrules
	}
	return toSerialize, nil
}

type NullableNetworkLoadBalancerEntities struct {
	value *NetworkLoadBalancerEntities
	isSet bool
}

func (v NullableNetworkLoadBalancerEntities) Get() *NetworkLoadBalancerEntities {
	return v.value
}

func (v *NullableNetworkLoadBalancerEntities) Set(val *NetworkLoadBalancerEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerEntities(val *NetworkLoadBalancerEntities) *NullableNetworkLoadBalancerEntities {
	return &NullableNetworkLoadBalancerEntities{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
