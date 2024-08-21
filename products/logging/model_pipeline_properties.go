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

// checks if the PipelineProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineProperties{}

// PipelineProperties A pipeline properties
type PipelineProperties struct {
	// The friendly name of your pipeline.
	Name *string `json:"name,omitempty"`
	// The information of the log aggregator
	Logs []PipelineResponse `json:"logs,omitempty"`
	// The address to connect fluentBit compatible logging agents to
	TcpAddress *string `json:"tcpAddress,omitempty"`
	// The address to post logs to using JSON with basic auth
	HttpAddress *string `json:"httpAddress,omitempty"`
	// The address of the client's grafana instance
	GrafanaAddress *string `json:"grafanaAddress,omitempty"`
}

// NewPipelineProperties instantiates a new PipelineProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineProperties() *PipelineProperties {
	this := PipelineProperties{}

	return &this
}

// NewPipelinePropertiesWithDefaults instantiates a new PipelineProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelinePropertiesWithDefaults() *PipelineProperties {
	this := PipelineProperties{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineProperties) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineProperties) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineProperties) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineProperties) SetName(v string) {
	o.Name = &v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *PipelineProperties) GetLogs() []PipelineResponse {
	if o == nil || IsNil(o.Logs) {
		var ret []PipelineResponse
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineProperties) GetLogsOk() ([]PipelineResponse, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *PipelineProperties) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []PipelineResponse and assigns it to the Logs field.
func (o *PipelineProperties) SetLogs(v []PipelineResponse) {
	o.Logs = v
}

// GetTcpAddress returns the TcpAddress field value if set, zero value otherwise.
func (o *PipelineProperties) GetTcpAddress() string {
	if o == nil || IsNil(o.TcpAddress) {
		var ret string
		return ret
	}
	return *o.TcpAddress
}

// GetTcpAddressOk returns a tuple with the TcpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineProperties) GetTcpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.TcpAddress) {
		return nil, false
	}
	return o.TcpAddress, true
}

// HasTcpAddress returns a boolean if a field has been set.
func (o *PipelineProperties) HasTcpAddress() bool {
	if o != nil && !IsNil(o.TcpAddress) {
		return true
	}

	return false
}

// SetTcpAddress gets a reference to the given string and assigns it to the TcpAddress field.
func (o *PipelineProperties) SetTcpAddress(v string) {
	o.TcpAddress = &v
}

// GetHttpAddress returns the HttpAddress field value if set, zero value otherwise.
func (o *PipelineProperties) GetHttpAddress() string {
	if o == nil || IsNil(o.HttpAddress) {
		var ret string
		return ret
	}
	return *o.HttpAddress
}

// GetHttpAddressOk returns a tuple with the HttpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineProperties) GetHttpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.HttpAddress) {
		return nil, false
	}
	return o.HttpAddress, true
}

// HasHttpAddress returns a boolean if a field has been set.
func (o *PipelineProperties) HasHttpAddress() bool {
	if o != nil && !IsNil(o.HttpAddress) {
		return true
	}

	return false
}

// SetHttpAddress gets a reference to the given string and assigns it to the HttpAddress field.
func (o *PipelineProperties) SetHttpAddress(v string) {
	o.HttpAddress = &v
}

// GetGrafanaAddress returns the GrafanaAddress field value if set, zero value otherwise.
func (o *PipelineProperties) GetGrafanaAddress() string {
	if o == nil || IsNil(o.GrafanaAddress) {
		var ret string
		return ret
	}
	return *o.GrafanaAddress
}

// GetGrafanaAddressOk returns a tuple with the GrafanaAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineProperties) GetGrafanaAddressOk() (*string, bool) {
	if o == nil || IsNil(o.GrafanaAddress) {
		return nil, false
	}
	return o.GrafanaAddress, true
}

// HasGrafanaAddress returns a boolean if a field has been set.
func (o *PipelineProperties) HasGrafanaAddress() bool {
	if o != nil && !IsNil(o.GrafanaAddress) {
		return true
	}

	return false
}

// SetGrafanaAddress gets a reference to the given string and assigns it to the GrafanaAddress field.
func (o *PipelineProperties) SetGrafanaAddress(v string) {
	o.GrafanaAddress = &v
}

func (o PipelineProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.TcpAddress) {
		toSerialize["tcpAddress"] = o.TcpAddress
	}
	if !IsNil(o.HttpAddress) {
		toSerialize["httpAddress"] = o.HttpAddress
	}
	if !IsNil(o.GrafanaAddress) {
		toSerialize["grafanaAddress"] = o.GrafanaAddress
	}
	return toSerialize, nil
}

type NullablePipelineProperties struct {
	value *PipelineProperties
	isSet bool
}

func (v NullablePipelineProperties) Get() *PipelineProperties {
	return v.value
}

func (v *NullablePipelineProperties) Set(val *PipelineProperties) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineProperties) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineProperties(val *PipelineProperties) *NullablePipelineProperties {
	return &NullablePipelineProperties{value: val, isSet: true}
}

func (v NullablePipelineProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
