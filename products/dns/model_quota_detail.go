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

// checks if the QuotaDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaDetail{}

// QuotaDetail Count of zones and records. This schema is used to show both usage and defined limits (quota)
type QuotaDetail struct {
	// count of the number of zones
	Zones int32 `json:"zones"`
	// count of the number of secondary zones
	SecondaryZones int32 `json:"secondaryZones"`
	// count of the number of records
	Records int32 `json:"records"`
	// count of the number of reverse DNS records
	ReverseRecords int32 `json:"reverseRecords"`
}

// NewQuotaDetail instantiates a new QuotaDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaDetail(zones int32, secondaryZones int32, records int32, reverseRecords int32) *QuotaDetail {
	this := QuotaDetail{}

	this.Zones = zones
	this.SecondaryZones = secondaryZones
	this.Records = records
	this.ReverseRecords = reverseRecords

	return &this
}

// NewQuotaDetailWithDefaults instantiates a new QuotaDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaDetailWithDefaults() *QuotaDetail {
	this := QuotaDetail{}
	var zones int32 = 0
	this.Zones = zones
	var secondaryZones int32 = 0
	this.SecondaryZones = secondaryZones
	var records int32 = 0
	this.Records = records
	var reverseRecords int32 = 0
	this.ReverseRecords = reverseRecords
	return &this
}

// GetZones returns the Zones field value
func (o *QuotaDetail) GetZones() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value
// and a boolean to check if the value has been set.
func (o *QuotaDetail) GetZonesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zones, true
}

// SetZones sets field value
func (o *QuotaDetail) SetZones(v int32) {
	o.Zones = v
}

// GetSecondaryZones returns the SecondaryZones field value
func (o *QuotaDetail) GetSecondaryZones() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecondaryZones
}

// GetSecondaryZonesOk returns a tuple with the SecondaryZones field value
// and a boolean to check if the value has been set.
func (o *QuotaDetail) GetSecondaryZonesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryZones, true
}

// SetSecondaryZones sets field value
func (o *QuotaDetail) SetSecondaryZones(v int32) {
	o.SecondaryZones = v
}

// GetRecords returns the Records field value
func (o *QuotaDetail) GetRecords() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *QuotaDetail) GetRecordsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Records, true
}

// SetRecords sets field value
func (o *QuotaDetail) SetRecords(v int32) {
	o.Records = v
}

// GetReverseRecords returns the ReverseRecords field value
func (o *QuotaDetail) GetReverseRecords() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReverseRecords
}

// GetReverseRecordsOk returns a tuple with the ReverseRecords field value
// and a boolean to check if the value has been set.
func (o *QuotaDetail) GetReverseRecordsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReverseRecords, true
}

// SetReverseRecords sets field value
func (o *QuotaDetail) SetReverseRecords(v int32) {
	o.ReverseRecords = v
}

func (o QuotaDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotaDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["zones"] = o.Zones
	toSerialize["secondaryZones"] = o.SecondaryZones
	toSerialize["records"] = o.Records
	toSerialize["reverseRecords"] = o.ReverseRecords
	return toSerialize, nil
}

type NullableQuotaDetail struct {
	value *QuotaDetail
	isSet bool
}

func (v NullableQuotaDetail) Get() *QuotaDetail {
	return v.value
}

func (v *NullableQuotaDetail) Set(val *QuotaDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaDetail(val *QuotaDetail) *NullableQuotaDetail {
	return &NullableQuotaDetail{value: val, isSet: true}
}

func (v NullableQuotaDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
