/*
 * IONOS Cloud - DNS API
 *
 * Cloud DNS service helps IONOS Cloud customers to automate DNS Zone and Record management.
 *
 * API version: 1.16.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the SecondaryZoneRecordReadListMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecondaryZoneRecordReadListMetadata{}

// SecondaryZoneRecordReadListMetadata Shows the specific properties for secondary zones
type SecondaryZoneRecordReadListMetadata struct {
	// Indicates IP addresses of primary nameservers for a secondary zone. Accepts IPv4 and IPv6 addresses
	PrimaryIps []string `json:"primaryIps"`
}

// NewSecondaryZoneRecordReadListMetadata instantiates a new SecondaryZoneRecordReadListMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecondaryZoneRecordReadListMetadata(primaryIps []string) *SecondaryZoneRecordReadListMetadata {
	this := SecondaryZoneRecordReadListMetadata{}

	this.PrimaryIps = primaryIps

	return &this
}

// NewSecondaryZoneRecordReadListMetadataWithDefaults instantiates a new SecondaryZoneRecordReadListMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecondaryZoneRecordReadListMetadataWithDefaults() *SecondaryZoneRecordReadListMetadata {
	this := SecondaryZoneRecordReadListMetadata{}
	return &this
}

// GetPrimaryIps returns the PrimaryIps field value
func (o *SecondaryZoneRecordReadListMetadata) GetPrimaryIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PrimaryIps
}

// GetPrimaryIpsOk returns a tuple with the PrimaryIps field value
// and a boolean to check if the value has been set.
func (o *SecondaryZoneRecordReadListMetadata) GetPrimaryIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIps, true
}

// SetPrimaryIps sets field value
func (o *SecondaryZoneRecordReadListMetadata) SetPrimaryIps(v []string) {
	o.PrimaryIps = v
}

func (o SecondaryZoneRecordReadListMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecondaryZoneRecordReadListMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsZero(o.PrimaryIps) {
		toSerialize["primaryIps"] = o.PrimaryIps
	}
	return toSerialize, nil
}

type NullableSecondaryZoneRecordReadListMetadata struct {
	value *SecondaryZoneRecordReadListMetadata
	isSet bool
}

func (v NullableSecondaryZoneRecordReadListMetadata) Get() *SecondaryZoneRecordReadListMetadata {
	return v.value
}

func (v *NullableSecondaryZoneRecordReadListMetadata) Set(val *SecondaryZoneRecordReadListMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSecondaryZoneRecordReadListMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSecondaryZoneRecordReadListMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecondaryZoneRecordReadListMetadata(val *SecondaryZoneRecordReadListMetadata) *NullableSecondaryZoneRecordReadListMetadata {
	return &NullableSecondaryZoneRecordReadListMetadata{value: val, isSet: true}
}

func (v NullableSecondaryZoneRecordReadListMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecondaryZoneRecordReadListMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}