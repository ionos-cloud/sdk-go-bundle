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

// checks if the TargetGroupPut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetGroupPut{}

// TargetGroupPut struct for TargetGroupPut
type TargetGroupPut struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
	// The URL to the object representation (absolute path).
	Href       *string               `json:"href,omitempty"`
	Properties TargetGroupProperties `json:"properties"`
}

// NewTargetGroupPut instantiates a new TargetGroupPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetGroupPut(properties TargetGroupProperties) *TargetGroupPut {
	this := TargetGroupPut{}

	this.Properties = properties

	return &this
}

// NewTargetGroupPutWithDefaults instantiates a new TargetGroupPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetGroupPutWithDefaults() *TargetGroupPut {
	this := TargetGroupPut{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TargetGroupPut) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetGroupPut) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TargetGroupPut) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TargetGroupPut) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TargetGroupPut) GetType() Type {
	if o == nil || IsNil(o.Type) {
		var ret Type
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetGroupPut) GetTypeOk() (*Type, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TargetGroupPut) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given Type and assigns it to the Type field.
func (o *TargetGroupPut) SetType(v Type) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *TargetGroupPut) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetGroupPut) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *TargetGroupPut) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *TargetGroupPut) SetHref(v string) {
	o.Href = &v
}

// GetProperties returns the Properties field value
func (o *TargetGroupPut) GetProperties() TargetGroupProperties {
	if o == nil {
		var ret TargetGroupProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *TargetGroupPut) GetPropertiesOk() (*TargetGroupProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *TargetGroupPut) SetProperties(v TargetGroupProperties) {
	o.Properties = v
}

func (o TargetGroupPut) ToMap() (map[string]interface{}, error) {
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

type NullableTargetGroupPut struct {
	value *TargetGroupPut
	isSet bool
}

func (v NullableTargetGroupPut) Get() *TargetGroupPut {
	return v.value
}

func (v *NullableTargetGroupPut) Set(val *TargetGroupPut) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetGroupPut) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetGroupPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetGroupPut(val *TargetGroupPut) *NullableTargetGroupPut {
	return &NullableTargetGroupPut{value: val, isSet: true}
}

func (v NullableTargetGroupPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetGroupPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
