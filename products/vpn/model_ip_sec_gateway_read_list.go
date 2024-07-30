/*
 * VPN Gateways
 *
 * POC Docs for VPN gateway as service
 *
 * API version: 0.0.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vpn

import (
	"encoding/json"
)

// checks if the IPSecGatewayReadList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPSecGatewayReadList{}

// IPSecGatewayReadList struct for IPSecGatewayReadList
type IPSecGatewayReadList struct {
	// ID of the list of IPSecGateway resources.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the list of IPSecGateway resources.
	Href string `json:"href"`
	// The list of IPSecGateway resources.
	Items []IPSecGatewayRead `json:"items,omitempty"`
	// The offset specified in the request (if none was specified, the default offset is 0).
	Offset int32 `json:"offset"`
	// The limit specified in the request (if none was specified, use the endpoint's default pagination limit).
	Limit int32 `json:"limit"`
	Links Links `json:"_links"`
}

// NewIPSecGatewayReadList instantiates a new IPSecGatewayReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecGatewayReadList(id string, type_ string, href string, offset int32, limit int32, links Links) *IPSecGatewayReadList {
	this := IPSecGatewayReadList{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Offset = offset
	this.Limit = limit
	this.Links = links

	return &this
}

// NewIPSecGatewayReadListWithDefaults instantiates a new IPSecGatewayReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecGatewayReadListWithDefaults() *IPSecGatewayReadList {
	this := IPSecGatewayReadList{}
	return &this
}

// GetId returns the Id field value
func (o *IPSecGatewayReadList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IPSecGatewayReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IPSecGatewayReadList) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *IPSecGatewayReadList) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IPSecGatewayReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IPSecGatewayReadList) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *IPSecGatewayReadList) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *IPSecGatewayReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *IPSecGatewayReadList) SetHref(v string) {
	o.Href = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *IPSecGatewayReadList) GetItems() []IPSecGatewayRead {
	if o == nil || IsNil(o.Items) {
		var ret []IPSecGatewayRead
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecGatewayReadList) GetItemsOk() ([]IPSecGatewayRead, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *IPSecGatewayReadList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []IPSecGatewayRead and assigns it to the Items field.
func (o *IPSecGatewayReadList) SetItems(v []IPSecGatewayRead) {
	o.Items = v
}

// GetOffset returns the Offset field value
func (o *IPSecGatewayReadList) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *IPSecGatewayReadList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *IPSecGatewayReadList) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *IPSecGatewayReadList) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *IPSecGatewayReadList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *IPSecGatewayReadList) SetLimit(v int32) {
	o.Limit = v
}

// GetLinks returns the Links field value
func (o *IPSecGatewayReadList) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *IPSecGatewayReadList) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *IPSecGatewayReadList) SetLinks(v Links) {
	o.Links = v
}

func (o IPSecGatewayReadList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPSecGatewayReadList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsZero(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsZero(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsZero(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsZero(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsZero(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsZero(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableIPSecGatewayReadList struct {
	value *IPSecGatewayReadList
	isSet bool
}

func (v NullableIPSecGatewayReadList) Get() *IPSecGatewayReadList {
	return v.value
}

func (v *NullableIPSecGatewayReadList) Set(val *IPSecGatewayReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecGatewayReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecGatewayReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecGatewayReadList(val *IPSecGatewayReadList) *NullableIPSecGatewayReadList {
	return &NullableIPSecGatewayReadList{value: val, isSet: true}
}

func (v NullableIPSecGatewayReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecGatewayReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
