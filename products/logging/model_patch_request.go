/*
 * IONOS Logging REST API
 *
 * Logging Service is a service that provides a centralized logging system where users are able to push and aggregate their system or application logs. This service also provides a visualization platform where users are able to observe, search and filter the logs and also create dashboards and alerts for their data points. This service can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an API. The API allows you to create logging pipelines or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logging

import (
	"encoding/json"
)

// PatchRequest Request payload with any data that is possible to patch a logging pipeline
type PatchRequest struct {
	Properties *PatchRequestProperties `json:"properties"`
}

// NewPatchRequest instantiates a new PatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchRequest(properties PatchRequestProperties) *PatchRequest {
	this := PatchRequest{}

	this.Properties = &properties

	return &this
}

// NewPatchRequestWithDefaults instantiates a new PatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchRequestWithDefaults() *PatchRequest {
	this := PatchRequest{}
	return &this
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for PatchRequestProperties will be returned
func (o *PatchRequest) GetProperties() *PatchRequestProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchRequest) GetPropertiesOk() (*PatchRequestProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *PatchRequest) SetProperties(v PatchRequestProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *PatchRequest) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o PatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullablePatchRequest struct {
	value *PatchRequest
	isSet bool
}

func (v NullablePatchRequest) Get() *PatchRequest {
	return v.value
}

func (v *NullablePatchRequest) Set(val *PatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchRequest(val *PatchRequest) *NullablePatchRequest {
	return &NullablePatchRequest{value: val, isSet: true}
}

func (v NullablePatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
