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

// checks if the NetworkLoadBalancerForwardingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkLoadBalancerForwardingRule{}

// NetworkLoadBalancerForwardingRule struct for NetworkLoadBalancerForwardingRule
type NetworkLoadBalancerForwardingRule struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href       *string                                     `json:"href,omitempty"`
	Metadata   *DatacenterElementMetadata                  `json:"metadata,omitempty"`
	Properties NetworkLoadBalancerForwardingRuleProperties `json:"properties"`
}

// NewNetworkLoadBalancerForwardingRule instantiates a new NetworkLoadBalancerForwardingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkLoadBalancerForwardingRule(properties NetworkLoadBalancerForwardingRuleProperties) *NetworkLoadBalancerForwardingRule {
	this := NetworkLoadBalancerForwardingRule{}

	this.Properties = properties

	return &this
}

// NewNetworkLoadBalancerForwardingRuleWithDefaults instantiates a new NetworkLoadBalancerForwardingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkLoadBalancerForwardingRuleWithDefaults() *NetworkLoadBalancerForwardingRule {
	this := NetworkLoadBalancerForwardingRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkLoadBalancerForwardingRule) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRule) GetType() Type {
	if o == nil || IsNil(o.Type) {
		var ret Type
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRule) GetTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given Type and assigns it to the Type field.
func (o *NetworkLoadBalancerForwardingRule) SetType(v Type) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRule) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRule) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRule) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *NetworkLoadBalancerForwardingRule) SetHref(v string) {
	o.Href = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NetworkLoadBalancerForwardingRule) GetMetadata() DatacenterElementMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret DatacenterElementMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRule) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NetworkLoadBalancerForwardingRule) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DatacenterElementMetadata and assigns it to the Metadata field.
func (o *NetworkLoadBalancerForwardingRule) SetMetadata(v DatacenterElementMetadata) {
	o.Metadata = &v
}

// GetProperties returns the Properties field value
func (o *NetworkLoadBalancerForwardingRule) GetProperties() NetworkLoadBalancerForwardingRuleProperties {
	if o == nil {
		var ret NetworkLoadBalancerForwardingRuleProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *NetworkLoadBalancerForwardingRule) GetPropertiesOk() (*NetworkLoadBalancerForwardingRuleProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *NetworkLoadBalancerForwardingRule) SetProperties(v NetworkLoadBalancerForwardingRuleProperties) {
	o.Properties = v
}

func (o NetworkLoadBalancerForwardingRule) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableNetworkLoadBalancerForwardingRule struct {
	value *NetworkLoadBalancerForwardingRule
	isSet bool
}

func (v NullableNetworkLoadBalancerForwardingRule) Get() *NetworkLoadBalancerForwardingRule {
	return v.value
}

func (v *NullableNetworkLoadBalancerForwardingRule) Set(val *NetworkLoadBalancerForwardingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkLoadBalancerForwardingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkLoadBalancerForwardingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkLoadBalancerForwardingRule(val *NetworkLoadBalancerForwardingRule) *NullableNetworkLoadBalancerForwardingRule {
	return &NullableNetworkLoadBalancerForwardingRule{value: val, isSet: true}
}

func (v NullableNetworkLoadBalancerForwardingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkLoadBalancerForwardingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
