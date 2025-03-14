/*
 * IONOS Cloud - API Gateway
 *
 * API Gateway is an application that acts as a \"front door\" for backend services and APIs, handling client requests and routing them to the appropriate backend.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apigateway

import (
	"encoding/json"
)

// checks if the Gateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gateway{}

// Gateway An API gateway consists of the generic rules and configurations of an API Gateway.
type Gateway struct {
	Name string `json:"name"`
	// This field enables or disables the collection and reporting of logs for observability of this instance.
	Logs *bool `json:"logs,omitempty"`
	// This field enables or disables the collection and reporting of metrics for observability of this instance.
	Metrics       *bool                  `json:"metrics,omitempty"`
	CustomDomains []GatewayCustomDomains `json:"customDomains,omitempty"`
}

// NewGateway instantiates a new Gateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateway(name string) *Gateway {
	this := Gateway{}

	this.Name = name
	var logs bool = false
	this.Logs = &logs
	var metrics bool = false
	this.Metrics = &metrics

	return &this
}

// NewGatewayWithDefaults instantiates a new Gateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayWithDefaults() *Gateway {
	this := Gateway{}
	var logs bool = false
	this.Logs = &logs
	var metrics bool = false
	this.Metrics = &metrics
	return &this
}

// GetName returns the Name field value
func (o *Gateway) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Gateway) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Gateway) SetName(v string) {
	o.Name = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *Gateway) GetLogs() bool {
	if o == nil || IsNil(o.Logs) {
		var ret bool
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetLogsOk() (*bool, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *Gateway) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given bool and assigns it to the Logs field.
func (o *Gateway) SetLogs(v bool) {
	o.Logs = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *Gateway) GetMetrics() bool {
	if o == nil || IsNil(o.Metrics) {
		var ret bool
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetMetricsOk() (*bool, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *Gateway) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given bool and assigns it to the Metrics field.
func (o *Gateway) SetMetrics(v bool) {
	o.Metrics = &v
}

// GetCustomDomains returns the CustomDomains field value if set, zero value otherwise.
func (o *Gateway) GetCustomDomains() []GatewayCustomDomains {
	if o == nil || IsNil(o.CustomDomains) {
		var ret []GatewayCustomDomains
		return ret
	}
	return o.CustomDomains
}

// GetCustomDomainsOk returns a tuple with the CustomDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetCustomDomainsOk() ([]GatewayCustomDomains, bool) {
	if o == nil || IsNil(o.CustomDomains) {
		return nil, false
	}
	return o.CustomDomains, true
}

// HasCustomDomains returns a boolean if a field has been set.
func (o *Gateway) HasCustomDomains() bool {
	if o != nil && !IsNil(o.CustomDomains) {
		return true
	}

	return false
}

// SetCustomDomains gets a reference to the given []GatewayCustomDomains and assigns it to the CustomDomains field.
func (o *Gateway) SetCustomDomains(v []GatewayCustomDomains) {
	o.CustomDomains = v
}

func (o Gateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.CustomDomains) {
		toSerialize["customDomains"] = o.CustomDomains
	}
	return toSerialize, nil
}

type NullableGateway struct {
	value *Gateway
	isSet bool
}

func (v NullableGateway) Get() *Gateway {
	return v.value
}

func (v *NullableGateway) Set(val *Gateway) {
	v.value = val
	v.isSet = true
}

func (v NullableGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateway(val *Gateway) *NullableGateway {
	return &NullableGateway{value: val, isSet: true}
}

func (v NullableGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
