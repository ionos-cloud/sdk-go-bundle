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

// checks if the AccessKeyReadListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessKeyReadListAllOf{}

// AccessKeyReadListAllOf struct for AccessKeyReadListAllOf
type AccessKeyReadListAllOf struct {
	// ID of the list of AccessKey resources.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the list of AccessKey resources.
	Href string `json:"href"`
	// The list of AccessKey resources.
	Items []AccessKeyRead `json:"items,omitempty"`
}

// NewAccessKeyReadListAllOf instantiates a new AccessKeyReadListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessKeyReadListAllOf(id string, type_ string, href string) *AccessKeyReadListAllOf {
	this := AccessKeyReadListAllOf{}

	this.Id = id
	this.Type = type_
	this.Href = href

	return &this
}

// NewAccessKeyReadListAllOfWithDefaults instantiates a new AccessKeyReadListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessKeyReadListAllOfWithDefaults() *AccessKeyReadListAllOf {
	this := AccessKeyReadListAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *AccessKeyReadListAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccessKeyReadListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccessKeyReadListAllOf) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *AccessKeyReadListAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessKeyReadListAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessKeyReadListAllOf) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *AccessKeyReadListAllOf) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *AccessKeyReadListAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *AccessKeyReadListAllOf) SetHref(v string) {
	o.Href = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AccessKeyReadListAllOf) GetItems() []AccessKeyRead {
	if o == nil || IsNil(o.Items) {
		var ret []AccessKeyRead
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessKeyReadListAllOf) GetItemsOk() ([]AccessKeyRead, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AccessKeyReadListAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AccessKeyRead and assigns it to the Items field.
func (o *AccessKeyReadListAllOf) SetItems(v []AccessKeyRead) {
	o.Items = v
}

func (o AccessKeyReadListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessKeyReadListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableAccessKeyReadListAllOf struct {
	value *AccessKeyReadListAllOf
	isSet bool
}

func (v NullableAccessKeyReadListAllOf) Get() *AccessKeyReadListAllOf {
	return v.value
}

func (v *NullableAccessKeyReadListAllOf) Set(val *AccessKeyReadListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessKeyReadListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessKeyReadListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessKeyReadListAllOf(val *AccessKeyReadListAllOf) *NullableAccessKeyReadListAllOf {
	return &NullableAccessKeyReadListAllOf{value: val, isSet: true}
}

func (v NullableAccessKeyReadListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessKeyReadListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
