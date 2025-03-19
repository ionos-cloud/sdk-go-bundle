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

// checks if the Region type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Region{}

// Region IONOS Cloud object storage regions they define the location of the bucket, can also be used as `LocationConstraint` for bucket creation.
type Region struct {
	// The version of the region properties
	Version int32 `json:"version"`
	// The endpoint URL for the region
	Endpoint string `json:"endpoint"`
	// The website URL for the region
	Website    string           `json:"website"`
	Capability RegionCapability `json:"capability"`
	// The available classes in the region
	StorageClasses []string `json:"storageClasses,omitempty"`
	// The data center location of the region as per [Get Location](/docs/cloud/v6/#tag/Locations/operation/locationsGet). *Can't be used as `LocationConstraint` on bucket creation.*
	Location string `json:"location"`
}

// NewRegion instantiates a new Region object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegion(version int32, endpoint string, website string, capability RegionCapability, location string) *Region {
	this := Region{}

	this.Version = version
	this.Endpoint = endpoint
	this.Website = website
	this.Capability = capability
	this.Location = location

	return &this
}

// NewRegionWithDefaults instantiates a new Region object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionWithDefaults() *Region {
	this := Region{}
	return &this
}

// GetVersion returns the Version field value
func (o *Region) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Region) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Region) SetVersion(v int32) {
	o.Version = v
}

// GetEndpoint returns the Endpoint field value
func (o *Region) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *Region) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *Region) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetWebsite returns the Website field value
func (o *Region) GetWebsite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Website
}

// GetWebsiteOk returns a tuple with the Website field value
// and a boolean to check if the value has been set.
func (o *Region) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Website, true
}

// SetWebsite sets field value
func (o *Region) SetWebsite(v string) {
	o.Website = v
}

// GetCapability returns the Capability field value
func (o *Region) GetCapability() RegionCapability {
	if o == nil {
		var ret RegionCapability
		return ret
	}

	return o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value
// and a boolean to check if the value has been set.
func (o *Region) GetCapabilityOk() (*RegionCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capability, true
}

// SetCapability sets field value
func (o *Region) SetCapability(v RegionCapability) {
	o.Capability = v
}

// GetStorageClasses returns the StorageClasses field value if set, zero value otherwise.
func (o *Region) GetStorageClasses() []string {
	if o == nil || IsNil(o.StorageClasses) {
		var ret []string
		return ret
	}
	return o.StorageClasses
}

// GetStorageClassesOk returns a tuple with the StorageClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Region) GetStorageClassesOk() ([]string, bool) {
	if o == nil || IsNil(o.StorageClasses) {
		return nil, false
	}
	return o.StorageClasses, true
}

// HasStorageClasses returns a boolean if a field has been set.
func (o *Region) HasStorageClasses() bool {
	if o != nil && !IsNil(o.StorageClasses) {
		return true
	}

	return false
}

// SetStorageClasses gets a reference to the given []string and assigns it to the StorageClasses field.
func (o *Region) SetStorageClasses(v []string) {
	o.StorageClasses = v
}

// GetLocation returns the Location field value
func (o *Region) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Region) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Region) SetLocation(v string) {
	o.Location = v
}

func (o Region) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Region) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["website"] = o.Website
	toSerialize["capability"] = o.Capability
	if !IsNil(o.StorageClasses) {
		toSerialize["storageClasses"] = o.StorageClasses
	}
	toSerialize["location"] = o.Location
	return toSerialize, nil
}

type NullableRegion struct {
	value *Region
	isSet bool
}

func (v NullableRegion) Get() *Region {
	return v.value
}

func (v *NullableRegion) Set(val *Region) {
	v.value = val
	v.isSet = true
}

func (v NullableRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegion(val *Region) *NullableRegion {
	return &NullableRegion{value: val, isSet: true}
}

func (v NullableRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
