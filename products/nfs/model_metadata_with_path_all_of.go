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

// checks if the MetadataWithPathAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataWithPathAllOf{}

// MetadataWithPathAllOf struct for MetadataWithPathAllOf
type MetadataWithPathAllOf struct {
	// The path of the NFS export.
	NfsPath string `json:"nfsPath"`
}

// NewMetadataWithPathAllOf instantiates a new MetadataWithPathAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataWithPathAllOf(nfsPath string) *MetadataWithPathAllOf {
	this := MetadataWithPathAllOf{}

	this.NfsPath = nfsPath

	return &this
}

// NewMetadataWithPathAllOfWithDefaults instantiates a new MetadataWithPathAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithPathAllOfWithDefaults() *MetadataWithPathAllOf {
	this := MetadataWithPathAllOf{}
	return &this
}

// GetNfsPath returns the NfsPath field value
func (o *MetadataWithPathAllOf) GetNfsPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfsPath
}

// GetNfsPathOk returns a tuple with the NfsPath field value
// and a boolean to check if the value has been set.
func (o *MetadataWithPathAllOf) GetNfsPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfsPath, true
}

// SetNfsPath sets field value
func (o *MetadataWithPathAllOf) SetNfsPath(v string) {
	o.NfsPath = v
}

func (o MetadataWithPathAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nfsPath"] = o.NfsPath
	return toSerialize, nil
}

type NullableMetadataWithPathAllOf struct {
	value *MetadataWithPathAllOf
	isSet bool
}

func (v NullableMetadataWithPathAllOf) Get() *MetadataWithPathAllOf {
	return v.value
}

func (v *NullableMetadataWithPathAllOf) Set(val *MetadataWithPathAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataWithPathAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataWithPathAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataWithPathAllOf(val *MetadataWithPathAllOf) *NullableMetadataWithPathAllOf {
	return &NullableMetadataWithPathAllOf{value: val, isSet: true}
}

func (v NullableMetadataWithPathAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataWithPathAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
