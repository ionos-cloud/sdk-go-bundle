/*
 * IONOS Cloud - Object Storage Management API
 *
 * Object Storage Management API is a RESTful API that manages the object storage service configuration for IONOS Cloud.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstoragemgmt

import (
	"encoding/json"
)

// checks if the RegionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionCreate{}

// RegionCreate struct for RegionCreate
type RegionCreate struct {
	// Metadata
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Properties Region                 `json:"properties"`
}

// NewRegionCreate instantiates a new RegionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionCreate(properties Region) *RegionCreate {
	this := RegionCreate{}

	this.Properties = properties

	return &this
}

// NewRegionCreateWithDefaults instantiates a new RegionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionCreateWithDefaults() *RegionCreate {
	this := RegionCreate{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RegionCreate) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionCreate) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RegionCreate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RegionCreate) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *RegionCreate) GetProperties() Region {
	if o == nil {
		var ret Region
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *RegionCreate) GetPropertiesOk() (*Region, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *RegionCreate) SetProperties(v Region) {
	o.Properties = v
}

func (o RegionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableRegionCreate struct {
	value *RegionCreate
	isSet bool
}

func (v NullableRegionCreate) Get() *RegionCreate {
	return v.value
}

func (v *NullableRegionCreate) Set(val *RegionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionCreate(val *RegionCreate) *NullableRegionCreate {
	return &NullableRegionCreate{value: val, isSet: true}
}

func (v NullableRegionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
