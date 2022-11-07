/*
 * Container Registry service
 *
 * Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls.
 *
 * API version: 1.0
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"
)

// PutTokenInput struct for PutTokenInput
type PutTokenInput struct {
	Properties *PostTokenProperties `json:"properties"`
}

// NewPutTokenInput instantiates a new PutTokenInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutTokenInput(properties PostTokenProperties) *PutTokenInput {
	this := PutTokenInput{}

	this.Properties = &properties

	return &this
}

// NewPutTokenInputWithDefaults instantiates a new PutTokenInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutTokenInputWithDefaults() *PutTokenInput {
	this := PutTokenInput{}
	return &this
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for PostTokenProperties will be returned
func (o *PutTokenInput) GetProperties() *PostTokenProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutTokenInput) GetPropertiesOk() (*PostTokenProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *PutTokenInput) SetProperties(v PostTokenProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *PutTokenInput) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o PutTokenInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullablePutTokenInput struct {
	value *PutTokenInput
	isSet bool
}

func (v NullablePutTokenInput) Get() *PutTokenInput {
	return v.value
}

func (v *NullablePutTokenInput) Set(val *PutTokenInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePutTokenInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePutTokenInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutTokenInput(val *PutTokenInput) *NullablePutTokenInput {
	return &NullablePutTokenInput{value: val, isSet: true}
}

func (v NullablePutTokenInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutTokenInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
