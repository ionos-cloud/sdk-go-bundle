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

// Processor struct for Processor
type Processor struct {
	// The source parser to be used
	Source *string `json:"source,omitempty"`
	// Tag is to distinguish different pipelines. must be unique amongst the pipeline's array items.
	Tag *string `json:"tag,omitempty"`
	// Protocol to use as intake
	Protocol *string `json:"protocol,omitempty"`
	// Optional custom labels to filter and report logs
	Labels *[]string `json:"labels,omitempty"`
	// The configuration of the logs datastore
	Destinations *[]Destination `json:"destinations,omitempty"`
}

// NewProcessor instantiates a new Processor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessor() *Processor {
	this := Processor{}

	return &this
}

// NewProcessorWithDefaults instantiates a new Processor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorWithDefaults() *Processor {
	this := Processor{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Processor) GetSource() *string {
	if o == nil {
		return nil
	}

	return o.Source

}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Processor) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Source, true
}

// SetSource sets field value
func (o *Processor) SetSource(v string) {

	o.Source = &v

}

// HasSource returns a boolean if a field has been set.
func (o *Processor) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// GetTag returns the Tag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Processor) GetTag() *string {
	if o == nil {
		return nil
	}

	return o.Tag

}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Processor) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Tag, true
}

// SetTag sets field value
func (o *Processor) SetTag(v string) {

	o.Tag = &v

}

// HasTag returns a boolean if a field has been set.
func (o *Processor) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// GetProtocol returns the Protocol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Processor) GetProtocol() *string {
	if o == nil {
		return nil
	}

	return o.Protocol

}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Processor) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Protocol, true
}

// SetProtocol sets field value
func (o *Processor) SetProtocol(v string) {

	o.Protocol = &v

}

// HasProtocol returns a boolean if a field has been set.
func (o *Processor) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// GetLabels returns the Labels field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *Processor) GetLabels() *[]string {
	if o == nil {
		return nil
	}

	return o.Labels

}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Processor) GetLabelsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Labels, true
}

// SetLabels sets field value
func (o *Processor) SetLabels(v []string) {

	o.Labels = &v

}

// HasLabels returns a boolean if a field has been set.
func (o *Processor) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// GetDestinations returns the Destinations field value
// If the value is explicit nil, the zero value for []Destination will be returned
func (o *Processor) GetDestinations() *[]Destination {
	if o == nil {
		return nil
	}

	return o.Destinations

}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Processor) GetDestinationsOk() (*[]Destination, bool) {
	if o == nil {
		return nil, false
	}

	return o.Destinations, true
}

// SetDestinations sets field value
func (o *Processor) SetDestinations(v []Destination) {

	o.Destinations = &v

}

// HasDestinations returns a boolean if a field has been set.
func (o *Processor) HasDestinations() bool {
	if o != nil && o.Destinations != nil {
		return true
	}

	return false
}

func (o Processor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}

	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}

	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}

	if o.Destinations != nil {
		toSerialize["destinations"] = o.Destinations
	}

	return json.Marshal(toSerialize)
}

type NullableProcessor struct {
	value *Processor
	isSet bool
}

func (v NullableProcessor) Get() *Processor {
	return v.value
}

func (v *NullableProcessor) Set(val *Processor) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessor) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessor(val *Processor) *NullableProcessor {
	return &NullableProcessor{value: val, isSet: true}
}

func (v NullableProcessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
