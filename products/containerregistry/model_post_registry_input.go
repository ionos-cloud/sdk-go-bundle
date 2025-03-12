/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature ### 1.2.0  - Added registry `apiSubnetAllowList` ### 1.2.1  - Amended `apiSubnetAllowList` Regex
 *
 * API version: 1.2.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"
)

// checks if the PostRegistryInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostRegistryInput{}

// PostRegistryInput struct for PostRegistryInput
type PostRegistryInput struct {
	Properties PostRegistryProperties `json:"properties"`
}

// NewPostRegistryInput instantiates a new PostRegistryInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRegistryInput(properties PostRegistryProperties) *PostRegistryInput {
	this := PostRegistryInput{}

	this.Properties = properties

	return &this
}

// NewPostRegistryInputWithDefaults instantiates a new PostRegistryInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRegistryInputWithDefaults() *PostRegistryInput {
	this := PostRegistryInput{}
	return &this
}

// GetProperties returns the Properties field value
func (o *PostRegistryInput) GetProperties() PostRegistryProperties {
	if o == nil {
		var ret PostRegistryProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *PostRegistryInput) GetPropertiesOk() (*PostRegistryProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *PostRegistryInput) SetProperties(v PostRegistryProperties) {
	o.Properties = v
}

func (o PostRegistryInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRegistryInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullablePostRegistryInput struct {
	value *PostRegistryInput
	isSet bool
}

func (v NullablePostRegistryInput) Get() *PostRegistryInput {
	return v.value
}

func (v *NullablePostRegistryInput) Set(val *PostRegistryInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRegistryInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRegistryInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRegistryInput(val *PostRegistryInput) *NullablePostRegistryInput {
	return &NullablePostRegistryInput{value: val, isSet: true}
}

func (v NullablePostRegistryInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRegistryInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
