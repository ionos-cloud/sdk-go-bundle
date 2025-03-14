/*
 * IONOS Cloud - DNS API
 *
 * Cloud DNS service helps IONOS Cloud customers to automate DNS Zone and Record management.
 *
 * API version: 1.17.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the ZoneReadAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneReadAllOf{}

// ZoneReadAllOf struct for ZoneReadAllOf
type ZoneReadAllOf struct {
	Properties Zone `json:"properties"`
}

// NewZoneReadAllOf instantiates a new ZoneReadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneReadAllOf(properties Zone) *ZoneReadAllOf {
	this := ZoneReadAllOf{}

	this.Properties = properties

	return &this
}

// NewZoneReadAllOfWithDefaults instantiates a new ZoneReadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneReadAllOfWithDefaults() *ZoneReadAllOf {
	this := ZoneReadAllOf{}
	return &this
}

// GetProperties returns the Properties field value
func (o *ZoneReadAllOf) GetProperties() Zone {
	if o == nil {
		var ret Zone
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ZoneReadAllOf) GetPropertiesOk() (*Zone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ZoneReadAllOf) SetProperties(v Zone) {
	o.Properties = v
}

func (o ZoneReadAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneReadAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableZoneReadAllOf struct {
	value *ZoneReadAllOf
	isSet bool
}

func (v NullableZoneReadAllOf) Get() *ZoneReadAllOf {
	return v.value
}

func (v *NullableZoneReadAllOf) Set(val *ZoneReadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneReadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneReadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneReadAllOf(val *ZoneReadAllOf) *NullableZoneReadAllOf {
	return &NullableZoneReadAllOf{value: val, isSet: true}
}

func (v NullableZoneReadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneReadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
