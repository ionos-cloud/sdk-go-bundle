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

// checks if the ClusterRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterRead{}

// ClusterRead struct for ClusterRead
type ClusterRead struct {
	// The ID (UUID) of the cluster.
	Id string `json:"id"`
	// The resource type
	Type string `json:"type"`
	// The URL of the cluster.
	Href       string             `json:"href"`
	Metadata   MetadataWithStatus `json:"metadata"`
	Properties Cluster            `json:"properties"`
}

// NewClusterRead instantiates a new ClusterRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRead(id string, type_ string, href string, metadata MetadataWithStatus, properties Cluster) *ClusterRead {
	this := ClusterRead{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Metadata = metadata
	this.Properties = properties

	return &this
}

// NewClusterReadWithDefaults instantiates a new ClusterRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterReadWithDefaults() *ClusterRead {
	this := ClusterRead{}
	return &this
}

// GetId returns the Id field value
func (o *ClusterRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClusterRead) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ClusterRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClusterRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClusterRead) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *ClusterRead) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ClusterRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ClusterRead) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value
func (o *ClusterRead) GetMetadata() MetadataWithStatus {
	if o == nil {
		var ret MetadataWithStatus
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ClusterRead) GetMetadataOk() (*MetadataWithStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ClusterRead) SetMetadata(v MetadataWithStatus) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *ClusterRead) GetProperties() Cluster {
	if o == nil {
		var ret Cluster
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ClusterRead) GetPropertiesOk() (*Cluster, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ClusterRead) SetProperties(v Cluster) {
	o.Properties = v
}

func (o ClusterRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["metadata"] = o.Metadata
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableClusterRead struct {
	value *ClusterRead
	isSet bool
}

func (v NullableClusterRead) Get() *ClusterRead {
	return v.value
}

func (v *NullableClusterRead) Set(val *ClusterRead) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRead) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRead(val *ClusterRead) *NullableClusterRead {
	return &NullableClusterRead{value: val, isSet: true}
}

func (v NullableClusterRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
