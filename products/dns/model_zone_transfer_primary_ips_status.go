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

// checks if the ZoneTransferPrimaryIpsStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneTransferPrimaryIpsStatus{}

// ZoneTransferPrimaryIpsStatus Indicates, for secondary zones, the transfer status for each one single IP defined on primaryIps field
type ZoneTransferPrimaryIpsStatus struct {
	Type  string                        `json:"type"`
	Items []ZoneTransferPrimaryIpStatus `json:"items"`
}

// NewZoneTransferPrimaryIpsStatus instantiates a new ZoneTransferPrimaryIpsStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneTransferPrimaryIpsStatus(type_ string, items []ZoneTransferPrimaryIpStatus) *ZoneTransferPrimaryIpsStatus {
	this := ZoneTransferPrimaryIpsStatus{}

	this.Type = type_
	this.Items = items

	return &this
}

// NewZoneTransferPrimaryIpsStatusWithDefaults instantiates a new ZoneTransferPrimaryIpsStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneTransferPrimaryIpsStatusWithDefaults() *ZoneTransferPrimaryIpsStatus {
	this := ZoneTransferPrimaryIpsStatus{}
	return &this
}

// GetType returns the Type field value
func (o *ZoneTransferPrimaryIpsStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ZoneTransferPrimaryIpsStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ZoneTransferPrimaryIpsStatus) SetType(v string) {
	o.Type = v
}

// GetItems returns the Items field value
func (o *ZoneTransferPrimaryIpsStatus) GetItems() []ZoneTransferPrimaryIpStatus {
	if o == nil {
		var ret []ZoneTransferPrimaryIpStatus
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ZoneTransferPrimaryIpsStatus) GetItemsOk() ([]ZoneTransferPrimaryIpStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ZoneTransferPrimaryIpsStatus) SetItems(v []ZoneTransferPrimaryIpStatus) {
	o.Items = v
}

func (o ZoneTransferPrimaryIpsStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZoneTransferPrimaryIpsStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableZoneTransferPrimaryIpsStatus struct {
	value *ZoneTransferPrimaryIpsStatus
	isSet bool
}

func (v NullableZoneTransferPrimaryIpsStatus) Get() *ZoneTransferPrimaryIpsStatus {
	return v.value
}

func (v *NullableZoneTransferPrimaryIpsStatus) Set(val *ZoneTransferPrimaryIpsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneTransferPrimaryIpsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneTransferPrimaryIpsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneTransferPrimaryIpsStatus(val *ZoneTransferPrimaryIpsStatus) *NullableZoneTransferPrimaryIpsStatus {
	return &NullableZoneTransferPrimaryIpsStatus{value: val, isSet: true}
}

func (v NullableZoneTransferPrimaryIpsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneTransferPrimaryIpsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
