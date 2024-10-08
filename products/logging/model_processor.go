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

// checks if the Processor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Processor{}

// Processor struct for Processor
type Processor struct {
	// The source parser to be used
	Source *string `json:"source,omitempty"`
	// Tag is to distinguish different pipelines. must be unique amongst the pipeline's array items.
	Tag *string `json:"tag,omitempty"`
	// Protocol to use as intake
	Protocol *string `json:"protocol,omitempty"`
	// Optional custom labels to filter and report logs
	Labels []string `json:"labels,omitempty"`
	// The configuration of the logs datastore
	Destinations []Destination `json:"destinations,omitempty"`
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

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Processor) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Processor) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Processor) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *Processor) SetSource(v string) {
	o.Source = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Processor) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Processor) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Processor) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *Processor) SetTag(v string) {
	o.Tag = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Processor) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Processor) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Processor) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Processor) SetProtocol(v string) {
	o.Protocol = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Processor) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Processor) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Processor) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *Processor) SetLabels(v []string) {
	o.Labels = v
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *Processor) GetDestinations() []Destination {
	if o == nil || IsNil(o.Destinations) {
		var ret []Destination
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Processor) GetDestinationsOk() ([]Destination, bool) {
	if o == nil || IsNil(o.Destinations) {
		return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *Processor) HasDestinations() bool {
	if o != nil && !IsNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []Destination and assigns it to the Destinations field.
func (o *Processor) SetDestinations(v []Destination) {
	o.Destinations = v
}

func (o Processor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return toSerialize, nil
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
