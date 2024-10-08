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

// checks if the ClusterLogs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterLogs{}

// ClusterLogs The logs of the PostgreSQL cluster.
type ClusterLogs struct {
	Instances []ClusterLogsInstances `json:"instances,omitempty"`
}

// NewClusterLogs instantiates a new ClusterLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLogs() *ClusterLogs {
	this := ClusterLogs{}

	return &this
}

// NewClusterLogsWithDefaults instantiates a new ClusterLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLogsWithDefaults() *ClusterLogs {
	this := ClusterLogs{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *ClusterLogs) GetInstances() []ClusterLogsInstances {
	if o == nil || IsNil(o.Instances) {
		var ret []ClusterLogsInstances
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogs) GetInstancesOk() ([]ClusterLogsInstances, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *ClusterLogs) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []ClusterLogsInstances and assigns it to the Instances field.
func (o *ClusterLogs) SetInstances(v []ClusterLogsInstances) {
	o.Instances = v
}

func (o ClusterLogs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	return toSerialize, nil
}

type NullableClusterLogs struct {
	value *ClusterLogs
	isSet bool
}

func (v NullableClusterLogs) Get() *ClusterLogs {
	return v.value
}

func (v *NullableClusterLogs) Set(val *ClusterLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLogs(val *ClusterLogs) *NullableClusterLogs {
	return &NullableClusterLogs{value: val, isSet: true}
}

func (v NullableClusterLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
