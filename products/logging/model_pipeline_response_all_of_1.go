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

// checks if the PipelineResponseAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineResponseAllOf1{}

// PipelineResponseAllOf1 struct for PipelineResponseAllOf1
type PipelineResponseAllOf1 struct {
	Destinations []Destination `json:"destinations,omitempty"`
}

// NewPipelineResponseAllOf1 instantiates a new PipelineResponseAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineResponseAllOf1() *PipelineResponseAllOf1 {
	this := PipelineResponseAllOf1{}

	return &this
}

// NewPipelineResponseAllOf1WithDefaults instantiates a new PipelineResponseAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineResponseAllOf1WithDefaults() *PipelineResponseAllOf1 {
	this := PipelineResponseAllOf1{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *PipelineResponseAllOf1) GetDestinations() []Destination {
	if o == nil || IsNil(o.Destinations) {
		var ret []Destination
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineResponseAllOf1) GetDestinationsOk() ([]Destination, bool) {
	if o == nil || IsNil(o.Destinations) {
		return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *PipelineResponseAllOf1) HasDestinations() bool {
	if o != nil && !IsNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []Destination and assigns it to the Destinations field.
func (o *PipelineResponseAllOf1) SetDestinations(v []Destination) {
	o.Destinations = v
}

func (o PipelineResponseAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineResponseAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return toSerialize, nil
}

type NullablePipelineResponseAllOf1 struct {
	value *PipelineResponseAllOf1
	isSet bool
}

func (v NullablePipelineResponseAllOf1) Get() *PipelineResponseAllOf1 {
	return v.value
}

func (v *NullablePipelineResponseAllOf1) Set(val *PipelineResponseAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineResponseAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineResponseAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineResponseAllOf1(val *PipelineResponseAllOf1) *NullablePipelineResponseAllOf1 {
	return &NullablePipelineResponseAllOf1{value: val, isSet: true}
}

func (v NullablePipelineResponseAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineResponseAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
