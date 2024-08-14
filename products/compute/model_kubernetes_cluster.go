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

// checks if the KubernetesCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesCluster{}

// KubernetesCluster struct for KubernetesCluster
type KubernetesCluster struct {
	// The resource unique identifier.
	Id *string `json:"id,omitempty"`
	// The object type.
	Type *string `json:"type,omitempty"`
	// The URL to the object representation (absolute path).
	Href       *string                     `json:"href,omitempty"`
	Metadata   *DatacenterElementMetadata  `json:"metadata,omitempty"`
	Properties KubernetesClusterProperties `json:"properties"`
	Entities   *KubernetesClusterEntities  `json:"entities,omitempty"`
}

// NewKubernetesCluster instantiates a new KubernetesCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesCluster(properties KubernetesClusterProperties) *KubernetesCluster {
	this := KubernetesCluster{}

	this.Properties = properties

	return &this
}

// NewKubernetesClusterWithDefaults instantiates a new KubernetesCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterWithDefaults() *KubernetesCluster {
	this := KubernetesCluster{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KubernetesCluster) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KubernetesCluster) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KubernetesCluster) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KubernetesCluster) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KubernetesCluster) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KubernetesCluster) SetType(v string) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *KubernetesCluster) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *KubernetesCluster) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *KubernetesCluster) SetHref(v string) {
	o.Href = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *KubernetesCluster) GetMetadata() DatacenterElementMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret DatacenterElementMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetMetadataOk() (*DatacenterElementMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KubernetesCluster) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given DatacenterElementMetadata and assigns it to the Metadata field.
func (o *KubernetesCluster) SetMetadata(v DatacenterElementMetadata) {
	o.Metadata = &v
}

// GetProperties returns the Properties field value
func (o *KubernetesCluster) GetProperties() KubernetesClusterProperties {
	if o == nil {
		var ret KubernetesClusterProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetPropertiesOk() (*KubernetesClusterProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *KubernetesCluster) SetProperties(v KubernetesClusterProperties) {
	o.Properties = v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *KubernetesCluster) GetEntities() KubernetesClusterEntities {
	if o == nil || IsNil(o.Entities) {
		var ret KubernetesClusterEntities
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesCluster) GetEntitiesOk() (*KubernetesClusterEntities, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *KubernetesCluster) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given KubernetesClusterEntities and assigns it to the Entities field.
func (o *KubernetesCluster) SetEntities(v KubernetesClusterEntities) {
	o.Entities = &v
}

func (o KubernetesCluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesCluster) ToMap() (map[string]interface{}, error) {
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
	if !IsZero(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	return toSerialize, nil
}

type NullableKubernetesCluster struct {
	value *KubernetesCluster
	isSet bool
}

func (v NullableKubernetesCluster) Get() *KubernetesCluster {
	return v.value
}

func (v *NullableKubernetesCluster) Set(val *KubernetesCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesCluster(val *KubernetesCluster) *NullableKubernetesCluster {
	return &NullableKubernetesCluster{value: val, isSet: true}
}

func (v NullableKubernetesCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
