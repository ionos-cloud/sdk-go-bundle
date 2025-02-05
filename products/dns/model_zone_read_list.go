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

// checks if the ZoneReadList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZoneReadList{}

// ZoneReadList struct for ZoneReadList
type ZoneReadList struct {
	// ID (UUID) created to identify this action.
	Id   string `json:"id"`
	Type string `json:"type"`
	Href string `json:"href"`
	// Pagination offset.
	Offset float32 `json:"offset"`
	// Pagination limit.
	Limit float32    `json:"limit"`
	Links Links      `json:"_links"`
	Items []ZoneRead `json:"items"`
}

// NewZoneReadList instantiates a new ZoneReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneReadList(id string, type_ string, href string, offset float32, limit float32, links Links, items []ZoneRead) *ZoneReadList {
	this := ZoneReadList{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Offset = offset
	this.Limit = limit
	this.Links = links
	this.Items = items

	return &this
}

// NewZoneReadListWithDefaults instantiates a new ZoneReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneReadListWithDefaults() *ZoneReadList {
	this := ZoneReadList{}
	return &this
}

// GetId returns the Id field value
func (o *ZoneReadList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ZoneReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ZoneReadList) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ZoneReadList) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ZoneReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ZoneReadList) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *ZoneReadList) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ZoneReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ZoneReadList) SetHref(v string) {
	o.Href = v
}

// GetOffset returns the Offset field value
func (o *ZoneReadList) GetOffset() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ZoneReadList) GetOffsetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ZoneReadList) SetOffset(v float32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *ZoneReadList) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ZoneReadList) GetLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ZoneReadList) SetLimit(v float32) {
	o.Limit = v
}

// GetLinks returns the Links field value
func (o *ZoneReadList) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ZoneReadList) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ZoneReadList) SetLinks(v Links) {
	o.Links = v
}

// GetItems returns the Items field value
func (o *ZoneReadList) GetItems() []ZoneRead {
	if o == nil {
		var ret []ZoneRead
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ZoneReadList) GetItemsOk() ([]ZoneRead, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ZoneReadList) SetItems(v []ZoneRead) {
	o.Items = v
}

func (o ZoneReadList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["offset"] = o.Offset
	toSerialize["limit"] = o.Limit
	toSerialize["_links"] = o.Links
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableZoneReadList struct {
	value *ZoneReadList
	isSet bool
}

func (v NullableZoneReadList) Get() *ZoneReadList {
	return v.value
}

func (v *NullableZoneReadList) Set(val *ZoneReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneReadList(val *ZoneReadList) *NullableZoneReadList {
	return &NullableZoneReadList{value: val, isSet: true}
}

func (v NullableZoneReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
