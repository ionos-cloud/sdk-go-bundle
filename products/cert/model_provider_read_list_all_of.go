/*
 * Certificate Manager Service API
 *
 * Using the Certificate Manager Service, you can conveniently provision and manage SSL certificates  with IONOS services and your internal connected resources.   For the [Application Load Balancer](https://api.ionos.com/docs/cloud/v6/#Application-Load-Balancers-get-datacenters-datacenterId-applicationloadbalancers), you usually need a certificate to encrypt your HTTPS traffic. The service provides the basic functions of uploading and deleting your certificates for this purpose.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cert

import (
	"encoding/json"
)

// checks if the ProviderReadListAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderReadListAllOf{}

// ProviderReadListAllOf struct for ProviderReadListAllOf
type ProviderReadListAllOf struct {
	// ID of the list of Provider resources.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the list of Provider resources.
	Href string `json:"href"`
	// The list of Provider resources.
	Items []ProviderRead `json:"items,omitempty"`
}

// NewProviderReadListAllOf instantiates a new ProviderReadListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderReadListAllOf(id string, type_ string, href string) *ProviderReadListAllOf {
	this := ProviderReadListAllOf{}

	this.Id = id
	this.Type = type_
	this.Href = href

	return &this
}

// NewProviderReadListAllOfWithDefaults instantiates a new ProviderReadListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderReadListAllOfWithDefaults() *ProviderReadListAllOf {
	this := ProviderReadListAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *ProviderReadListAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProviderReadListAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProviderReadListAllOf) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ProviderReadListAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProviderReadListAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProviderReadListAllOf) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *ProviderReadListAllOf) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ProviderReadListAllOf) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ProviderReadListAllOf) SetHref(v string) {
	o.Href = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ProviderReadListAllOf) GetItems() []ProviderRead {
	if o == nil || IsNil(o.Items) {
		var ret []ProviderRead
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderReadListAllOf) GetItemsOk() ([]ProviderRead, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ProviderReadListAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ProviderRead and assigns it to the Items field.
func (o *ProviderReadListAllOf) SetItems(v []ProviderRead) {
	o.Items = v
}

func (o ProviderReadListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderReadListAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableProviderReadListAllOf struct {
	value *ProviderReadListAllOf
	isSet bool
}

func (v NullableProviderReadListAllOf) Get() *ProviderReadListAllOf {
	return v.value
}

func (v *NullableProviderReadListAllOf) Set(val *ProviderReadListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderReadListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderReadListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderReadListAllOf(val *ProviderReadListAllOf) *NullableProviderReadListAllOf {
	return &NullableProviderReadListAllOf{value: val, isSet: true}
}

func (v NullableProviderReadListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderReadListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
