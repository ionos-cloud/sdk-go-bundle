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

// checks if the KubernetesNodePoolForPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesNodePoolForPost{}

// KubernetesNodePoolForPost struct for KubernetesNodePoolForPost
type KubernetesNodePoolForPost struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The object type.
	Type *string `json:"type,omitempty"`
	// The URL to the object representation (absolute path).
	Href       *string                             `json:"href,omitempty"`
	Metadata   *DatacenterElementMetadata          `json:"metadata,omitempty"`
	Properties KubernetesNodePoolPropertiesForPost `json:"properties"`
}

// NewKubernetesNodePoolForPost instantiates a new KubernetesNodePoolForPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodePoolForPost(properties KubernetesNodePoolPropertiesForPost) *KubernetesNodePoolForPost {
	this := KubernetesNodePoolForPost{}

	this.Properties = properties

	return &this
}

// NewKubernetesNodePoolForPostWithDefaults instantiates a new KubernetesNodePoolForPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodePoolForPostWithDefaults() *KubernetesNodePoolForPost {
	this := KubernetesNodePoolForPost{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesNodePoolForPost) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodePoolForPost) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesNodePoolForPost) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KubernetesNodePoolForPost) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KubernetesNodePoolForPost) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodePoolForPost) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KubernetesNodePoolForPost) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KubernetesNodePoolForPost) SetType(v string) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *KubernetesNodePoolForPost) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodePoolForPost) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *KubernetesNodePoolForPost) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *KubernetesNodePoolForPost) SetHref(v string) {
	o.Href = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *KubernetesNodePoolForPost) GetMetadata() DatacenterElementMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret DatacenterElementMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodePoolForPost) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KubernetesNodePoolForPost) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DatacenterElementMetadata and assigns it to the Metadata field.
func (o *KubernetesNodePoolForPost) SetMetadata(v DatacenterElementMetadata) {
	o.Metadata = &v
}

// GetProperties returns the Properties field value
func (o *KubernetesNodePoolForPost) GetProperties() KubernetesNodePoolPropertiesForPost {
	if o == nil {
		var ret KubernetesNodePoolPropertiesForPost
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodePoolForPost) GetPropertiesOk() (*KubernetesNodePoolPropertiesForPost, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *KubernetesNodePoolForPost) SetProperties(v KubernetesNodePoolPropertiesForPost) {
	o.Properties = v
}

func (o KubernetesNodePoolForPost) ToMap() (map[string]interface{}, error) {
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

type NullableKubernetesNodePoolForPost struct {
	value *KubernetesNodePoolForPost
	isSet bool
}

func (v NullableKubernetesNodePoolForPost) Get() *KubernetesNodePoolForPost {
	return v.value
}

func (v *NullableKubernetesNodePoolForPost) Set(val *KubernetesNodePoolForPost) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodePoolForPost) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodePoolForPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodePoolForPost(val *KubernetesNodePoolForPost) *NullableKubernetesNodePoolForPost {
	return &NullableKubernetesNodePoolForPost{value: val, isSet: true}
}

func (v NullableKubernetesNodePoolForPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodePoolForPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
