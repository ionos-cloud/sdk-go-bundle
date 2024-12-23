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

// checks if the DataCenterEntities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataCenterEntities{}

// DataCenterEntities struct for DataCenterEntities
type DataCenterEntities struct {
	Servers              *Servers              `json:"servers,omitempty"`
	Volumes              *Volumes              `json:"volumes,omitempty"`
	Loadbalancers        *Loadbalancers        `json:"loadbalancers,omitempty"`
	Lans                 *Lans                 `json:"lans,omitempty"`
	Networkloadbalancers *NetworkLoadBalancers `json:"networkloadbalancers,omitempty"`
	Natgateways          *NatGateways          `json:"natgateways,omitempty"`
	Securitygroups       *SecurityGroups       `json:"securitygroups,omitempty"`
}

// NewDataCenterEntities instantiates a new DataCenterEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataCenterEntities() *DataCenterEntities {
	this := DataCenterEntities{}

	return &this
}

// NewDataCenterEntitiesWithDefaults instantiates a new DataCenterEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataCenterEntitiesWithDefaults() *DataCenterEntities {
	this := DataCenterEntities{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *DataCenterEntities) GetServers() Servers {
	if o == nil || IsNil(o.Servers) {
		var ret Servers
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenterEntities) GetServersOk() (*Servers, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *DataCenterEntities) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given Servers and assigns it to the Servers field.
func (o *DataCenterEntities) SetServers(v Servers) {
	o.Servers = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *DataCenterEntities) GetVolumes() Volumes {
	if o == nil || IsNil(o.Volumes) {
		var ret Volumes
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenterEntities) GetVolumesOk() (*Volumes, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *DataCenterEntities) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given Volumes and assigns it to the Volumes field.
func (o *DataCenterEntities) SetVolumes(v Volumes) {
	o.Volumes = &v
}

// GetLoadbalancers returns the Loadbalancers field value if set, zero value otherwise.
func (o *DataCenterEntities) GetLoadbalancers() Loadbalancers {
	if o == nil || IsNil(o.Loadbalancers) {
		var ret Loadbalancers
		return ret
	}
	return *o.Loadbalancers
}

// GetLoadbalancersOk returns a tuple with the Loadbalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenterEntities) GetLoadbalancersOk() (*Loadbalancers, bool) {
	if o == nil || IsNil(o.Loadbalancers) {
		return nil, false
	}
	return o.Loadbalancers, true
}

// HasLoadbalancers returns a boolean if a field has been set.
func (o *DataCenterEntities) HasLoadbalancers() bool {
	if o != nil && !IsNil(o.Loadbalancers) {
		return true
	}

	return false
}

// SetLoadbalancers gets a reference to the given Loadbalancers and assigns it to the Loadbalancers field.
func (o *DataCenterEntities) SetLoadbalancers(v Loadbalancers) {
	o.Loadbalancers = &v
}

// GetLans returns the Lans field value if set, zero value otherwise.
func (o *DataCenterEntities) GetLans() Lans {
	if o == nil || IsNil(o.Lans) {
		var ret Lans
		return ret
	}
	return *o.Lans
}

// GetLansOk returns a tuple with the Lans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenterEntities) GetLansOk() (*Lans, bool) {
	if o == nil || IsNil(o.Lans) {
		return nil, false
	}
	return o.Lans, true
}

// HasLans returns a boolean if a field has been set.
func (o *DataCenterEntities) HasLans() bool {
	if o != nil && !IsNil(o.Lans) {
		return true
	}

	return false
}

// SetLans gets a reference to the given Lans and assigns it to the Lans field.
func (o *DataCenterEntities) SetLans(v Lans) {
	o.Lans = &v
}

// GetNetworkloadbalancers returns the Networkloadbalancers field value if set, zero value otherwise.
func (o *DataCenterEntities) GetNetworkloadbalancers() NetworkLoadBalancers {
	if o == nil || IsNil(o.Networkloadbalancers) {
		var ret NetworkLoadBalancers
		return ret
	}
	return *o.Networkloadbalancers
}

// GetNetworkloadbalancersOk returns a tuple with the Networkloadbalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenterEntities) GetNetworkloadbalancersOk() (*NetworkLoadBalancers, bool) {
	if o == nil || IsNil(o.Networkloadbalancers) {
		return nil, false
	}
	return o.Networkloadbalancers, true
}

// HasNetworkloadbalancers returns a boolean if a field has been set.
func (o *DataCenterEntities) HasNetworkloadbalancers() bool {
	if o != nil && !IsNil(o.Networkloadbalancers) {
		return true
	}

	return false
}

// SetNetworkloadbalancers gets a reference to the given NetworkLoadBalancers and assigns it to the Networkloadbalancers field.
func (o *DataCenterEntities) SetNetworkloadbalancers(v NetworkLoadBalancers) {
	o.Networkloadbalancers = &v
}

// GetNatgateways returns the Natgateways field value if set, zero value otherwise.
func (o *DataCenterEntities) GetNatgateways() NatGateways {
	if o == nil || IsNil(o.Natgateways) {
		var ret NatGateways
		return ret
	}
	return *o.Natgateways
}

// GetNatgatewaysOk returns a tuple with the Natgateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenterEntities) GetNatgatewaysOk() (*NatGateways, bool) {
	if o == nil || IsNil(o.Natgateways) {
		return nil, false
	}
	return o.Natgateways, true
}

// HasNatgateways returns a boolean if a field has been set.
func (o *DataCenterEntities) HasNatgateways() bool {
	if o != nil && !IsNil(o.Natgateways) {
		return true
	}

	return false
}

// SetNatgateways gets a reference to the given NatGateways and assigns it to the Natgateways field.
func (o *DataCenterEntities) SetNatgateways(v NatGateways) {
	o.Natgateways = &v
}

// GetSecuritygroups returns the Securitygroups field value if set, zero value otherwise.
func (o *DataCenterEntities) GetSecuritygroups() SecurityGroups {
	if o == nil || IsNil(o.Securitygroups) {
		var ret SecurityGroups
		return ret
	}
	return *o.Securitygroups
}

// GetSecuritygroupsOk returns a tuple with the Securitygroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCenterEntities) GetSecuritygroupsOk() (*SecurityGroups, bool) {
	if o == nil || IsNil(o.Securitygroups) {
		return nil, false
	}
	return o.Securitygroups, true
}

// HasSecuritygroups returns a boolean if a field has been set.
func (o *DataCenterEntities) HasSecuritygroups() bool {
	if o != nil && !IsNil(o.Securitygroups) {
		return true
	}

	return false
}

// SetSecuritygroups gets a reference to the given SecurityGroups and assigns it to the Securitygroups field.
func (o *DataCenterEntities) SetSecuritygroups(v SecurityGroups) {
	o.Securitygroups = &v
}

func (o DataCenterEntities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	if !IsNil(o.Loadbalancers) {
		toSerialize["loadbalancers"] = o.Loadbalancers
	}
	if !IsNil(o.Lans) {
		toSerialize["lans"] = o.Lans
	}
	if !IsNil(o.Networkloadbalancers) {
		toSerialize["networkloadbalancers"] = o.Networkloadbalancers
	}
	if !IsNil(o.Natgateways) {
		toSerialize["natgateways"] = o.Natgateways
	}
	if !IsNil(o.Securitygroups) {
		toSerialize["securitygroups"] = o.Securitygroups
	}
	return toSerialize, nil
}

type NullableDataCenterEntities struct {
	value *DataCenterEntities
	isSet bool
}

func (v NullableDataCenterEntities) Get() *DataCenterEntities {
	return v.value
}

func (v *NullableDataCenterEntities) Set(val *DataCenterEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableDataCenterEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableDataCenterEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataCenterEntities(val *DataCenterEntities) *NullableDataCenterEntities {
	return &NullableDataCenterEntities{value: val, isSet: true}
}

func (v NullableDataCenterEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataCenterEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
