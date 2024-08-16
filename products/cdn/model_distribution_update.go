/*
 * IONOS Cloud - CDN Distribution API
 *
 * This API manages CDN distributions.
 *
 * API version: 0.1.7
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the DistributionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributionUpdate{}

// DistributionUpdate struct for DistributionUpdate
type DistributionUpdate struct {
	// The ID (UUID) of the Distribution.
	Id string `json:"id"`
	// Metadata
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Properties DistributionProperties `json:"properties"`
}

// NewDistributionUpdate instantiates a new DistributionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionUpdate(id string, properties DistributionProperties) *DistributionUpdate {
	this := DistributionUpdate{}

	this.Id = id
	this.Properties = properties

	return &this
}

// NewDistributionUpdateWithDefaults instantiates a new DistributionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionUpdateWithDefaults() *DistributionUpdate {
	this := DistributionUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *DistributionUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DistributionUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DistributionUpdate) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DistributionUpdate) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionUpdate) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DistributionUpdate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *DistributionUpdate) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *DistributionUpdate) GetProperties() DistributionProperties {
	if o == nil {
		var ret DistributionProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *DistributionUpdate) GetPropertiesOk() (*DistributionProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *DistributionUpdate) SetProperties(v DistributionProperties) {
	o.Properties = v
}

func (o DistributionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistributionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsZero(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsZero(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableDistributionUpdate struct {
	value *DistributionUpdate
	isSet bool
}

func (v NullableDistributionUpdate) Get() *DistributionUpdate {
	return v.value
}

func (v *NullableDistributionUpdate) Set(val *DistributionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionUpdate(val *DistributionUpdate) *NullableDistributionUpdate {
	return &NullableDistributionUpdate{value: val, isSet: true}
}

func (v NullableDistributionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
