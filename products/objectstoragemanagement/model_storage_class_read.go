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

// checks if the StorageClassRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageClassRead{}

// StorageClassRead struct for StorageClassRead
type StorageClassRead struct {
	// The StorageClass of the StorageClass.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the StorageClass.
	Href       string                 `json:"href"`
	Metadata   map[string]interface{} `json:"metadata"`
	Properties StorageClass           `json:"properties"`
}

// NewStorageClassRead instantiates a new StorageClassRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageClassRead(id string, type_ string, href string, metadata map[string]interface{}, properties StorageClass) *StorageClassRead {
	this := StorageClassRead{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Metadata = metadata
	this.Properties = properties

	return &this
}

// NewStorageClassReadWithDefaults instantiates a new StorageClassRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageClassReadWithDefaults() *StorageClassRead {
	this := StorageClassRead{}
	return &this
}

// GetId returns the Id field value
func (o *StorageClassRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StorageClassRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StorageClassRead) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *StorageClassRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StorageClassRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StorageClassRead) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *StorageClassRead) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *StorageClassRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *StorageClassRead) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value
func (o *StorageClassRead) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *StorageClassRead) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *StorageClassRead) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *StorageClassRead) GetProperties() StorageClass {
	if o == nil {
		var ret StorageClass
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *StorageClassRead) GetPropertiesOk() (*StorageClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *StorageClassRead) SetProperties(v StorageClass) {
	o.Properties = v
}

func (o StorageClassRead) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageClassRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["metadata"] = o.Metadata
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableStorageClassRead struct {
	value *StorageClassRead
	isSet bool
}

func (v NullableStorageClassRead) Get() *StorageClassRead {
	return v.value
}

func (v *NullableStorageClassRead) Set(val *StorageClassRead) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageClassRead) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageClassRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageClassRead(val *StorageClassRead) *NullableStorageClassRead {
	return &NullableStorageClassRead{value: val, isSet: true}
}

func (v NullableStorageClassRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageClassRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
