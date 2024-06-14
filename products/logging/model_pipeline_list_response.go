/*
 * IONOS Logging REST API
 *
 * Logging as a Service (LaaS) is a service that provides a centralized logging system where users are able to push and aggregate their system or application logs. This service also provides a visualization platform where users are able to observe, search and filter the logs and also create dashboards and alerts for their data points. This service can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an API. The API allows you to create logging pipelines or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logging

import (
	"encoding/json"
)

// PipelineListResponse List of pipelines
type PipelineListResponse struct {
	Id    *string     `json:"id,omitempty"`
	Type  *string     `json:"type,omitempty"`
	Items *[]Pipeline `json:"items,omitempty"`
}

// NewPipelineListResponse instantiates a new PipelineListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineListResponse() *PipelineListResponse {
	this := PipelineListResponse{}

	return &this
}

// NewPipelineListResponseWithDefaults instantiates a new PipelineListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineListResponseWithDefaults() *PipelineListResponse {
	this := PipelineListResponse{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PipelineListResponse) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineListResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *PipelineListResponse) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *PipelineListResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PipelineListResponse) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineListResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *PipelineListResponse) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *PipelineListResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []Pipeline will be returned
func (o *PipelineListResponse) GetItems() *[]Pipeline {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineListResponse) GetItemsOk() (*[]Pipeline, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *PipelineListResponse) SetItems(v []Pipeline) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *PipelineListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

func (o PipelineListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	return json.Marshal(toSerialize)
}

type NullablePipelineListResponse struct {
	value *PipelineListResponse
	isSet bool
}

func (v NullablePipelineListResponse) Get() *PipelineListResponse {
	return v.value
}

func (v *NullablePipelineListResponse) Set(val *PipelineListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineListResponse(val *PipelineListResponse) *NullablePipelineListResponse {
	return &NullablePipelineListResponse{value: val, isSet: true}
}

func (v NullablePipelineListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
