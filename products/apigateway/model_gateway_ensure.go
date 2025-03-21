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

// checks if the GatewayEnsure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayEnsure{}

// GatewayEnsure struct for GatewayEnsure
type GatewayEnsure struct {
	// The ID (UUID) of the Gateway.
	Id string `json:"id"`
	// Metadata
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Properties Gateway                `json:"properties"`
}

// NewGatewayEnsure instantiates a new GatewayEnsure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayEnsure(id string, properties Gateway) *GatewayEnsure {
	this := GatewayEnsure{}

	this.Id = id
	this.Properties = properties

	return &this
}

// NewGatewayEnsureWithDefaults instantiates a new GatewayEnsure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayEnsureWithDefaults() *GatewayEnsure {
	this := GatewayEnsure{}
	return &this
}

// GetId returns the Id field value
func (o *GatewayEnsure) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GatewayEnsure) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GatewayEnsure) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GatewayEnsure) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayEnsure) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GatewayEnsure) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GatewayEnsure) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *GatewayEnsure) GetProperties() Gateway {
	if o == nil {
		var ret Gateway
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *GatewayEnsure) GetPropertiesOk() (*Gateway, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *GatewayEnsure) SetProperties(v Gateway) {
	o.Properties = v
}

func (o GatewayEnsure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableGatewayEnsure struct {
	value *GatewayEnsure
	isSet bool
}

func (v NullableGatewayEnsure) Get() *GatewayEnsure {
	return v.value
}

func (v *NullableGatewayEnsure) Set(val *GatewayEnsure) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayEnsure) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayEnsure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayEnsure(val *GatewayEnsure) *NullableGatewayEnsure {
	return &NullableGatewayEnsure{value: val, isSet: true}
}

func (v NullableGatewayEnsure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayEnsure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
