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

// checks if the ZoneCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneCreate{}

// ZoneCreate struct for ZoneCreate
type ZoneCreate struct {
	Properties Zone `json:"properties"`
}

// NewZoneCreate instantiates a new ZoneCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneCreate(properties Zone) *ZoneCreate {
	this := ZoneCreate{}

	this.Properties = properties

	return &this
}

// NewZoneCreateWithDefaults instantiates a new ZoneCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneCreateWithDefaults() *ZoneCreate {
	this := ZoneCreate{}
	return &this
}

// GetProperties returns the Properties field value
func (o *ZoneCreate) GetProperties() Zone {
	if o == nil {
		var ret Zone
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ZoneCreate) GetPropertiesOk() (*Zone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ZoneCreate) SetProperties(v Zone) {
	o.Properties = v
}

func (o ZoneCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableZoneCreate struct {
	value *ZoneCreate
	isSet bool
}

func (v NullableZoneCreate) Get() *ZoneCreate {
	return v.value
}

func (v *NullableZoneCreate) Set(val *ZoneCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneCreate(val *ZoneCreate) *NullableZoneCreate {
	return &NullableZoneCreate{value: val, isSet: true}
}

func (v NullableZoneCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
