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

// checks if the GatewayCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayCreate{}

// GatewayCreate struct for GatewayCreate
type GatewayCreate struct {
	// Metadata
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Properties Gateway                `json:"properties"`
}

// NewGatewayCreate instantiates a new GatewayCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayCreate(properties Gateway) *GatewayCreate {
	this := GatewayCreate{}

	this.Properties = properties

	return &this
}

// NewGatewayCreateWithDefaults instantiates a new GatewayCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayCreateWithDefaults() *GatewayCreate {
	this := GatewayCreate{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GatewayCreate) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayCreate) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GatewayCreate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GatewayCreate) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *GatewayCreate) GetProperties() Gateway {
	if o == nil {
		var ret Gateway
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *GatewayCreate) GetPropertiesOk() (*Gateway, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *GatewayCreate) SetProperties(v Gateway) {
	o.Properties = v
}

func (o GatewayCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsZero(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableGatewayCreate struct {
	value *GatewayCreate
	isSet bool
}

func (v NullableGatewayCreate) Get() *GatewayCreate {
	return v.value
}

func (v *NullableGatewayCreate) Set(val *GatewayCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayCreate(val *GatewayCreate) *NullableGatewayCreate {
	return &NullableGatewayCreate{value: val, isSet: true}
}

func (v NullableGatewayCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
