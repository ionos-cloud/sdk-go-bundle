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

// checks if the UpstreamGeoRestrictions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpstreamGeoRestrictions{}

// UpstreamGeoRestrictions This field manages the list of countries that are allowed or blocked from accessing the resource from the upstream host based on their ISO 3166-1 alpha-2 codes.
type UpstreamGeoRestrictions struct {
	// Country codes, the format should be based on ISO 3166-1 alpha-2 codes standard. Those codes are used to either blacklist or whitelist countries in geoIPBlock.
	BlockList []string `json:"blockList,omitempty"`
	// Country codes, the format should be based on ISO 3166-1 alpha-2 codes standard. Those codes are used to either blacklist or whitelist countries in geoIPBlock.
	AllowList []string `json:"allowList,omitempty"`
}

// NewUpstreamGeoRestrictions instantiates a new UpstreamGeoRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpstreamGeoRestrictions() *UpstreamGeoRestrictions {
	this := UpstreamGeoRestrictions{}

	return &this
}

// NewUpstreamGeoRestrictionsWithDefaults instantiates a new UpstreamGeoRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpstreamGeoRestrictionsWithDefaults() *UpstreamGeoRestrictions {
	this := UpstreamGeoRestrictions{}
	return &this
}

// GetBlockList returns the BlockList field value if set, zero value otherwise.
func (o *UpstreamGeoRestrictions) GetBlockList() []string {
	if o == nil || IsNil(o.BlockList) {
		var ret []string
		return ret
	}
	return o.BlockList
}

// GetBlockListOk returns a tuple with the BlockList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamGeoRestrictions) GetBlockListOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockList) {
		return nil, false
	}
	return o.BlockList, true
}

// HasBlockList returns a boolean if a field has been set.
func (o *UpstreamGeoRestrictions) HasBlockList() bool {
	if o != nil && !IsNil(o.BlockList) {
		return true
	}

	return false
}

// SetBlockList gets a reference to the given []string and assigns it to the BlockList field.
func (o *UpstreamGeoRestrictions) SetBlockList(v []string) {
	o.BlockList = v
}

// GetAllowList returns the AllowList field value if set, zero value otherwise.
func (o *UpstreamGeoRestrictions) GetAllowList() []string {
	if o == nil || IsNil(o.AllowList) {
		var ret []string
		return ret
	}
	return o.AllowList
}

// GetAllowListOk returns a tuple with the AllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamGeoRestrictions) GetAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowList) {
		return nil, false
	}
	return o.AllowList, true
}

// HasAllowList returns a boolean if a field has been set.
func (o *UpstreamGeoRestrictions) HasAllowList() bool {
	if o != nil && !IsNil(o.AllowList) {
		return true
	}

	return false
}

// SetAllowList gets a reference to the given []string and assigns it to the AllowList field.
func (o *UpstreamGeoRestrictions) SetAllowList(v []string) {
	o.AllowList = v
}

func (o UpstreamGeoRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlockList) {
		toSerialize["blockList"] = o.BlockList
	}
	if !IsNil(o.AllowList) {
		toSerialize["allowList"] = o.AllowList
	}
	return toSerialize, nil
}

type NullableUpstreamGeoRestrictions struct {
	value *UpstreamGeoRestrictions
	isSet bool
}

func (v NullableUpstreamGeoRestrictions) Get() *UpstreamGeoRestrictions {
	return v.value
}

func (v *NullableUpstreamGeoRestrictions) Set(val *UpstreamGeoRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableUpstreamGeoRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableUpstreamGeoRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpstreamGeoRestrictions(val *UpstreamGeoRestrictions) *NullableUpstreamGeoRestrictions {
	return &NullableUpstreamGeoRestrictions{value: val, isSet: true}
}

func (v NullableUpstreamGeoRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpstreamGeoRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
