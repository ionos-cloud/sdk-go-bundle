/*
 * IONOS Cloud - API Gateway
 *
 * API Gateway is an application that acts as a \"front door\" for backend services and APIs, handling client requests and routing them to the appropriate backend.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apigateway

import (
	"encoding/json"
)

// checks if the RouteReadList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteReadList{}

// RouteReadList struct for RouteReadList
type RouteReadList struct {
	// ID of the Route.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the Route.
	Href string `json:"href"`
	// The list of Route resources.
	Items []RouteRead `json:"items,omitempty"`
	// The offset specified in the request (if none was specified, the default offset is 0).
	Offset int32 `json:"offset"`
	// The limit specified in the request (if none was specified, use the endpoint's default pagination limit).
	Limit int32 `json:"limit"`
	Links Links `json:"_links"`
}

// NewRouteReadList instantiates a new RouteReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteReadList(id string, type_ string, href string, offset int32, limit int32, links Links) *RouteReadList {
	this := RouteReadList{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Offset = offset
	this.Limit = limit
	this.Links = links

	return &this
}

// NewRouteReadListWithDefaults instantiates a new RouteReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteReadListWithDefaults() *RouteReadList {
	this := RouteReadList{}
	return &this
}

// GetId returns the Id field value
func (o *RouteReadList) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RouteReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RouteReadList) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RouteReadList) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RouteReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RouteReadList) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *RouteReadList) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *RouteReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *RouteReadList) SetHref(v string) {
	o.Href = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RouteReadList) GetItems() []RouteRead {
	if o == nil || IsNil(o.Items) {
		var ret []RouteRead
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteReadList) GetItemsOk() ([]RouteRead, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RouteReadList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RouteRead and assigns it to the Items field.
func (o *RouteReadList) SetItems(v []RouteRead) {
	o.Items = v
}

// GetOffset returns the Offset field value
func (o *RouteReadList) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *RouteReadList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *RouteReadList) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *RouteReadList) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *RouteReadList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *RouteReadList) SetLimit(v int32) {
	o.Limit = v
}

// GetLinks returns the Links field value
func (o *RouteReadList) GetLinks() Links {
	if o == nil {
		var ret Links
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RouteReadList) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RouteReadList) SetLinks(v Links) {
	o.Links = v
}

func (o RouteReadList) ToMap() (map[string]interface{}, error) {
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

type NullableRouteReadList struct {
	value *RouteReadList
	isSet bool
}

func (v NullableRouteReadList) Get() *RouteReadList {
	return v.value
}

func (v *NullableRouteReadList) Set(val *RouteReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteReadList(val *RouteReadList) *NullableRouteReadList {
	return &NullableRouteReadList{value: val, isSet: true}
}

func (v NullableRouteReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
