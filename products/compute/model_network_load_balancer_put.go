/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// checks if the NetworkLoadBalancerPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkLoadBalancerPut{}

// NetworkLoadBalancerPut struct for NetworkLoadBalancerPut
type NetworkLoadBalancerPut struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href       *string                       `json:"href,omitempty"`
	Properties NetworkLoadBalancerProperties `json:"properties"`
}

// NewNetworkLoadBalancerPut instantiates a new NetworkLoadBalancerPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancerPut(properties NetworkLoadBalancerProperties) *NetworkLoadBalancerPut {
	this := NetworkLoadBalancerPut{}

	this.Properties = properties

	return &this
}

// NewNetworkLoadBalancerPutWithDefaults instantiates a new NetworkLoadBalancerPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerPutWithDefaults() *NetworkLoadBalancerPut {
	this := NetworkLoadBalancerPut{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkLoadBalancerPut) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerPut) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkLoadBalancerPut) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkLoadBalancerPut) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkLoadBalancerPut) GetType() Type {
	if o == nil || IsNil(o.Type) {
		var ret Type
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerPut) GetTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkLoadBalancerPut) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given Type and assigns it to the Type field.
func (o *NetworkLoadBalancerPut) SetType(v Type) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *NetworkLoadBalancerPut) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerPut) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *NetworkLoadBalancerPut) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *NetworkLoadBalancerPut) SetHref(v string) {
	o.Href = &v
}

// GetProperties returns the Properties field value
func (o *NetworkLoadBalancerPut) GetProperties() NetworkLoadBalancerProperties {
	if o == nil {
		var ret NetworkLoadBalancerProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerPut) GetPropertiesOk() (*NetworkLoadBalancerProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *NetworkLoadBalancerPut) SetProperties(v NetworkLoadBalancerProperties) {
	o.Properties = v
}

func (o NetworkLoadBalancerPut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableNetworkLoadBalancerPut struct {
	value *NetworkLoadBalancerPut
	isSet bool
}

func (v NullableNetworkLoadBalancerPut) Get() *NetworkLoadBalancerPut {
	return v.value
}

func (v *NullableNetworkLoadBalancerPut) Set(val *NetworkLoadBalancerPut) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerPut) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerPut(val *NetworkLoadBalancerPut) *NullableNetworkLoadBalancerPut {
	return &NullableNetworkLoadBalancerPut{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
