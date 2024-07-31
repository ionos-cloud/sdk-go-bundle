/*
 * IONOS Logging REST API
 *
 * The logging service offers a centralized platform to collect and store logs from various systems and applications. It includes tools to search, filter, visualize, and create alerts based on your log data.  This API provides programmatic control over logging pipelines, enabling you to create new pipelines or modify existing ones. It mirrors the functionality of the DCD visual tool, ensuring a consistent experience regardless of your chosen interface.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logging

import (
	"encoding/json"
)

// checks if the PipelineListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineListResponse{}

// PipelineListResponse List of pipelines
type PipelineListResponse struct {
	Id    *string    `json:"id,omitempty"`
	Type  *string    `json:"type,omitempty"`
	Items []Pipeline `json:"items,omitempty"`
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *PipelineListResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineListResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PipelineListResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PipelineListResponse) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PipelineListResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineListResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PipelineListResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PipelineListResponse) SetType(v string) {
	o.Type = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PipelineListResponse) GetItems() []Pipeline {
	if o == nil || IsNil(o.Items) {
		var ret []Pipeline
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineListResponse) GetItemsOk() ([]Pipeline, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PipelineListResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Pipeline and assigns it to the Items field.
func (o *PipelineListResponse) SetItems(v []Pipeline) {
	o.Items = v
}

func (o PipelineListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
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
