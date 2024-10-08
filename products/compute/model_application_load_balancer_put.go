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

// checks if the ApplicationLoadBalancerPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationLoadBalancerPut{}

// ApplicationLoadBalancerPut struct for ApplicationLoadBalancerPut
type ApplicationLoadBalancerPut struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// The URL to the object representation (absolute path).
	Href       *string                           `json:"href,omitempty"`
	Properties ApplicationLoadBalancerProperties `json:"properties"`
}

// NewApplicationLoadBalancerPut instantiates a new ApplicationLoadBalancerPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLoadBalancerPut(properties ApplicationLoadBalancerProperties) *ApplicationLoadBalancerPut {
	this := ApplicationLoadBalancerPut{}

	this.Properties = properties

	return &this
}

// NewApplicationLoadBalancerPutWithDefaults instantiates a new ApplicationLoadBalancerPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLoadBalancerPutWithDefaults() *ApplicationLoadBalancerPut {
	this := ApplicationLoadBalancerPut{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerPut) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerPut) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerPut) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationLoadBalancerPut) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerPut) GetType() Type {
	if o == nil || IsNil(o.Type) {
		var ret Type
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerPut) GetTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerPut) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given Type and assigns it to the Type field.
func (o *ApplicationLoadBalancerPut) SetType(v Type) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerPut) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerPut) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerPut) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ApplicationLoadBalancerPut) SetHref(v string) {
	o.Href = &v
}

// GetProperties returns the Properties field value
func (o *ApplicationLoadBalancerPut) GetProperties() ApplicationLoadBalancerProperties {
	if o == nil {
		var ret ApplicationLoadBalancerProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerPut) GetPropertiesOk() (*ApplicationLoadBalancerProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ApplicationLoadBalancerPut) SetProperties(v ApplicationLoadBalancerProperties) {
	o.Properties = v
}

func (o ApplicationLoadBalancerPut) ToMap() (map[string]interface{}, error) {
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

type NullableApplicationLoadBalancerPut struct {
	value *ApplicationLoadBalancerPut
	isSet bool
}

func (v NullableApplicationLoadBalancerPut) Get() *ApplicationLoadBalancerPut {
	return v.value
}

func (v *NullableApplicationLoadBalancerPut) Set(val *ApplicationLoadBalancerPut) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLoadBalancerPut) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLoadBalancerPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLoadBalancerPut(val *ApplicationLoadBalancerPut) *NullableApplicationLoadBalancerPut {
	return &NullableApplicationLoadBalancerPut{value: val, isSet: true}
}

func (v NullableApplicationLoadBalancerPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLoadBalancerPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
