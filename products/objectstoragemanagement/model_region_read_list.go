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

// checks if the RegionReadList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionReadList{}

// RegionReadList struct for RegionReadList
type RegionReadList struct {
	// ID of the list of Region resources.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the list of Region resources.
	Href string `json:"href"`
	// The list of Region resources.
	Items []RegionRead `json:"items,omitempty"`
	// The offset specified in the request (if none was specified, the default offset is 0).
	Offset int32 `json:"offset"`
	// The limit specified in the request (if none was specified, use the endpoint's default pagination limit).
	Limit int32 `json:"limit"`
	Links Links `json:"_links"`
}

// NewRegionReadList instantiates a new RegionReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionReadList(id string, type_ string, href string, offset int32, limit int32, links Links) *RegionReadList {
	this := RegionReadList{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Offset = offset
	this.Limit = limit
	this.Links = links

	return &this
}

// NewRegionReadListWithDefaults instantiates a new RegionReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionReadListWithDefaults() *RegionReadList {
	this := RegionReadList{}
	return &this
}

// GetId returns the Id field value
func (o *RegionReadList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegionReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegionReadList) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RegionReadList) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RegionReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RegionReadList) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *RegionReadList) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *RegionReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *RegionReadList) SetHref(v string) {
	o.Href = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RegionReadList) GetItems() []RegionRead {
	if o == nil || IsNil(o.Items) {
		var ret []RegionRead
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionReadList) GetItemsOk() ([]RegionRead, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RegionReadList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RegionRead and assigns it to the Items field.
func (o *RegionReadList) SetItems(v []RegionRead) {
	o.Items = v
}

// GetOffset returns the Offset field value
func (o *RegionReadList) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *RegionReadList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *RegionReadList) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *RegionReadList) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *RegionReadList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *RegionReadList) SetLimit(v int32) {
	o.Limit = v
}

// GetLinks returns the Links field value
func (o *RegionReadList) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RegionReadList) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RegionReadList) SetLinks(v Links) {
	o.Links = v
}

func (o RegionReadList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionReadList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["offset"] = o.Offset
	toSerialize["limit"] = o.Limit
	toSerialize["_links"] = o.Links
	return toSerialize, nil
}

type NullableRegionReadList struct {
	value *RegionReadList
	isSet bool
}

func (v NullableRegionReadList) Get() *RegionReadList {
	return v.value
}

func (v *NullableRegionReadList) Set(val *RegionReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionReadList(val *RegionReadList) *NullableRegionReadList {
	return &NullableRegionReadList{value: val, isSet: true}
}

func (v NullableRegionReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
