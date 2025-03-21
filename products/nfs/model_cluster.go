/*
 * IONOS Cloud - Network File Storage API
 *
 * The RESTful API for managing Network File Storage.
 *
 * API version: 0.1.3
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nfs

import (
	"encoding/json"
)

// checks if the Cluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cluster{}

// Cluster Network File Storage cluster
type Cluster struct {
	Name        string               `json:"name"`
	Connections []ClusterConnections `json:"connections"`
	Nfs         *ClusterNfs          `json:"nfs,omitempty"`
	// The size of the Network File Storage cluster in TiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees.
	Size *int32 `json:"size,omitempty"`
}

// NewCluster instantiates a new Cluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCluster(name string, connections []ClusterConnections) *Cluster {
	this := Cluster{}

	this.Name = name
	this.Connections = connections
	var size int32 = 2
	this.Size = &size

	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	var size int32 = 2
	this.Size = &size
	return &this
}

// GetName returns the Name field value
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetConnections returns the Connections field value
func (o *Cluster) GetConnections() []ClusterConnections {
	if o == nil {
		var ret []ClusterConnections
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetConnectionsOk() ([]ClusterConnections, bool) {
	if o == nil {
		return nil, false
	}
	return o.Connections, true
}

// SetConnections sets field value
func (o *Cluster) SetConnections(v []ClusterConnections) {
	o.Connections = v
}

// GetNfs returns the Nfs field value if set, zero value otherwise.
func (o *Cluster) GetNfs() ClusterNfs {
	if o == nil || IsNil(o.Nfs) {
		var ret ClusterNfs
		return ret
	}
	return *o.Nfs
}

// GetNfsOk returns a tuple with the Nfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetNfsOk() (*ClusterNfs, bool) {
	if o == nil || IsNil(o.Nfs) {
		return nil, false
	}
	return o.Nfs, true
}

// HasNfs returns a boolean if a field has been set.
func (o *Cluster) HasNfs() bool {
	if o != nil && !IsNil(o.Nfs) {
		return true
	}

	return false
}

// SetNfs gets a reference to the given ClusterNfs and assigns it to the Nfs field.
func (o *Cluster) SetNfs(v ClusterNfs) {
	o.Nfs = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Cluster) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Cluster) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *Cluster) SetSize(v int32) {
	o.Size = &v
}

func (o Cluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["connections"] = o.Connections
	if !IsNil(o.Nfs) {
		toSerialize["nfs"] = o.Nfs
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
