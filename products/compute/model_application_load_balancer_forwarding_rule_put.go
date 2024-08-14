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

// checks if the ApplicationLoadBalancerForwardingRulePut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationLoadBalancerForwardingRulePut{}

// ApplicationLoadBalancerForwardingRulePut struct for ApplicationLoadBalancerForwardingRulePut
type ApplicationLoadBalancerForwardingRulePut struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// The URL to the object representation (absolute path).
	Href       *string                                         `json:"href,omitempty"`
	Properties ApplicationLoadBalancerForwardingRuleProperties `json:"properties"`
}

// NewApplicationLoadBalancerForwardingRulePut instantiates a new ApplicationLoadBalancerForwardingRulePut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLoadBalancerForwardingRulePut(properties ApplicationLoadBalancerForwardingRuleProperties) *ApplicationLoadBalancerForwardingRulePut {
	this := ApplicationLoadBalancerForwardingRulePut{}

	this.Properties = properties

	return &this
}

// NewApplicationLoadBalancerForwardingRulePutWithDefaults instantiates a new ApplicationLoadBalancerForwardingRulePut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLoadBalancerForwardingRulePutWithDefaults() *ApplicationLoadBalancerForwardingRulePut {
	this := ApplicationLoadBalancerForwardingRulePut{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerForwardingRulePut) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationLoadBalancerForwardingRulePut) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerForwardingRulePut) GetType() Type {
	if o == nil || IsNil(o.Type) {
		var ret Type
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) GetTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given Type and assigns it to the Type field.
func (o *ApplicationLoadBalancerForwardingRulePut) SetType(v Type) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ApplicationLoadBalancerForwardingRulePut) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ApplicationLoadBalancerForwardingRulePut) SetHref(v string) {
	o.Href = &v
}

// GetProperties returns the Properties field value
func (o *ApplicationLoadBalancerForwardingRulePut) GetProperties() ApplicationLoadBalancerForwardingRuleProperties {
	if o == nil {
		var ret ApplicationLoadBalancerForwardingRuleProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ApplicationLoadBalancerForwardingRulePut) GetPropertiesOk() (*ApplicationLoadBalancerForwardingRuleProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ApplicationLoadBalancerForwardingRulePut) SetProperties(v ApplicationLoadBalancerForwardingRuleProperties) {
	o.Properties = v
}

func (o ApplicationLoadBalancerForwardingRulePut) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationLoadBalancerForwardingRulePut) ToMap() (map[string]interface{}, error) {
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
	if !IsZero(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableApplicationLoadBalancerForwardingRulePut struct {
	value *ApplicationLoadBalancerForwardingRulePut
	isSet bool
}

func (v NullableApplicationLoadBalancerForwardingRulePut) Get() *ApplicationLoadBalancerForwardingRulePut {
	return v.value
}

func (v *NullableApplicationLoadBalancerForwardingRulePut) Set(val *ApplicationLoadBalancerForwardingRulePut) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLoadBalancerForwardingRulePut) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLoadBalancerForwardingRulePut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLoadBalancerForwardingRulePut(val *ApplicationLoadBalancerForwardingRulePut) *NullableApplicationLoadBalancerForwardingRulePut {
	return &NullableApplicationLoadBalancerForwardingRulePut{value: val, isSet: true}
}

func (v NullableApplicationLoadBalancerForwardingRulePut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLoadBalancerForwardingRulePut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
