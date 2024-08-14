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

// checks if the KubernetesClusters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesClusters{}

// KubernetesClusters struct for KubernetesClusters
type KubernetesClusters struct {
	// The unique representation of the K8s cluster as a resource collection.
	Id *string `json:"id,omitempty"`
	// The resource type within a collection.
	Type *string `json:"type,omitempty"`
	// The URL to the collection representation (absolute path).
	Href *string `json:"href,omitempty"`
	// Array of K8s clusters in the collection.
	Items []KubernetesCluster `json:"items,omitempty"`
}

// NewKubernetesClusters instantiates a new KubernetesClusters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusters() *KubernetesClusters {
	this := KubernetesClusters{}

	return &this
}

// NewKubernetesClustersWithDefaults instantiates a new KubernetesClusters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClustersWithDefaults() *KubernetesClusters {
	this := KubernetesClusters{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesClusters) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusters) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesClusters) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KubernetesClusters) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KubernetesClusters) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusters) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KubernetesClusters) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KubernetesClusters) SetType(v string) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *KubernetesClusters) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusters) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *KubernetesClusters) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *KubernetesClusters) SetHref(v string) {
	o.Href = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *KubernetesClusters) GetItems() []KubernetesCluster {
	if o == nil || IsNil(o.Items) {
		var ret []KubernetesCluster
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusters) GetItemsOk() ([]KubernetesCluster, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *KubernetesClusters) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []KubernetesCluster and assigns it to the Items field.
func (o *KubernetesClusters) SetItems(v []KubernetesCluster) {
	o.Items = v
}

func (o KubernetesClusters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesClusters) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableKubernetesClusters struct {
	value *KubernetesClusters
	isSet bool
}

func (v NullableKubernetesClusters) Get() *KubernetesClusters {
	return v.value
}

func (v *NullableKubernetesClusters) Set(val *KubernetesClusters) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusters) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusters(val *KubernetesClusters) *NullableKubernetesClusters {
	return &NullableKubernetesClusters{value: val, isSet: true}
}

func (v NullableKubernetesClusters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
