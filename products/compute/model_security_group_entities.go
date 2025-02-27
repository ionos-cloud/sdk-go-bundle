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

// checks if the SecurityGroupEntities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityGroupEntities{}

// SecurityGroupEntities struct for SecurityGroupEntities
type SecurityGroupEntities struct {
	Rules   *FirewallRules `json:"rules,omitempty"`
	Nics    *Nics          `json:"nics,omitempty"`
	Servers *Servers       `json:"servers,omitempty"`
}

// NewSecurityGroupEntities instantiates a new SecurityGroupEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityGroupEntities() *SecurityGroupEntities {
	this := SecurityGroupEntities{}

	return &this
}

// NewSecurityGroupEntitiesWithDefaults instantiates a new SecurityGroupEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityGroupEntitiesWithDefaults() *SecurityGroupEntities {
	this := SecurityGroupEntities{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *SecurityGroupEntities) GetRules() FirewallRules {
	if o == nil || IsNil(o.Rules) {
		var ret FirewallRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupEntities) GetRulesOk() (*FirewallRules, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SecurityGroupEntities) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given FirewallRules and assigns it to the Rules field.
func (o *SecurityGroupEntities) SetRules(v FirewallRules) {
	o.Rules = &v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *SecurityGroupEntities) GetNics() Nics {
	if o == nil || IsNil(o.Nics) {
		var ret Nics
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupEntities) GetNicsOk() (*Nics, bool) {
	if o == nil || IsNil(o.Nics) {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *SecurityGroupEntities) HasNics() bool {
	if o != nil && !IsNil(o.Nics) {
		return true
	}

	return false
}

// SetNics gets a reference to the given Nics and assigns it to the Nics field.
func (o *SecurityGroupEntities) SetNics(v Nics) {
	o.Nics = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *SecurityGroupEntities) GetServers() Servers {
	if o == nil || IsNil(o.Servers) {
		var ret Servers
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityGroupEntities) GetServersOk() (*Servers, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *SecurityGroupEntities) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given Servers and assigns it to the Servers field.
func (o *SecurityGroupEntities) SetServers(v Servers) {
	o.Servers = &v
}

func (o SecurityGroupEntities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Nics) {
		toSerialize["nics"] = o.Nics
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

type NullableSecurityGroupEntities struct {
	value *SecurityGroupEntities
	isSet bool
}

func (v NullableSecurityGroupEntities) Get() *SecurityGroupEntities {
	return v.value
}

func (v *NullableSecurityGroupEntities) Set(val *SecurityGroupEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityGroupEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityGroupEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityGroupEntities(val *SecurityGroupEntities) *NullableSecurityGroupEntities {
	return &NullableSecurityGroupEntities{value: val, isSet: true}
}

func (v NullableSecurityGroupEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityGroupEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
