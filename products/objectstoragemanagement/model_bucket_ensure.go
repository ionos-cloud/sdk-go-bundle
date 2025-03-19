/*
 * IONOS Cloud - Object Storage Management API
 *
 * Object Storage Management API is a RESTful API that manages the object storage service configuration for IONOS Cloud.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstoragemanagement

import (
	"encoding/json"
)

// checks if the BucketEnsure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketEnsure{}

// BucketEnsure struct for BucketEnsure
type BucketEnsure struct {
	// The Bucket of the Bucket.
	Id string `json:"id"`
	// Metadata
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Properties Bucket                 `json:"properties"`
}

// NewBucketEnsure instantiates a new BucketEnsure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketEnsure(id string, properties Bucket) *BucketEnsure {
	this := BucketEnsure{}

	this.Id = id
	this.Properties = properties

	return &this
}

// NewBucketEnsureWithDefaults instantiates a new BucketEnsure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketEnsureWithDefaults() *BucketEnsure {
	this := BucketEnsure{}
	return &this
}

// GetId returns the Id field value
func (o *BucketEnsure) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BucketEnsure) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BucketEnsure) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BucketEnsure) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketEnsure) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BucketEnsure) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *BucketEnsure) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *BucketEnsure) GetProperties() Bucket {
	if o == nil {
		var ret Bucket
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *BucketEnsure) GetPropertiesOk() (*Bucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *BucketEnsure) SetProperties(v Bucket) {
	o.Properties = v
}

func (o BucketEnsure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketEnsure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableBucketEnsure struct {
	value *BucketEnsure
	isSet bool
}

func (v NullableBucketEnsure) Get() *BucketEnsure {
	return v.value
}

func (v *NullableBucketEnsure) Set(val *BucketEnsure) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketEnsure) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketEnsure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketEnsure(val *BucketEnsure) *NullableBucketEnsure {
	return &NullableBucketEnsure{value: val, isSet: true}
}

func (v NullableBucketEnsure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketEnsure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
