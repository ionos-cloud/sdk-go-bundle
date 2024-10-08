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

// checks if the ProviderPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProviderPatch{}

// ProviderPatch struct for ProviderPatch
type ProviderPatch struct {
	// Metadata
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Properties PatchName              `json:"properties"`
}

// NewProviderPatch instantiates a new ProviderPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderPatch(properties PatchName) *ProviderPatch {
	this := ProviderPatch{}

	this.Properties = properties

	return &this
}

// NewProviderPatchWithDefaults instantiates a new ProviderPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderPatchWithDefaults() *ProviderPatch {
	this := ProviderPatch{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ProviderPatch) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderPatch) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProviderPatch) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ProviderPatch) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *ProviderPatch) GetProperties() PatchName {
	if o == nil {
		var ret PatchName
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ProviderPatch) GetPropertiesOk() (*PatchName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ProviderPatch) SetProperties(v PatchName) {
	o.Properties = v
}

func (o ProviderPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableProviderPatch struct {
	value *ProviderPatch
	isSet bool
}

func (v NullableProviderPatch) Get() *ProviderPatch {
	return v.value
}

func (v *NullableProviderPatch) Set(val *ProviderPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderPatch(val *ProviderPatch) *NullableProviderPatch {
	return &NullableProviderPatch{value: val, isSet: true}
}

func (v NullableProviderPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
