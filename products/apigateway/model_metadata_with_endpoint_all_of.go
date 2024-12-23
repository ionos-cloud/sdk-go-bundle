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

// checks if the MetadataWithEndpointAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataWithEndpointAllOf{}

// MetadataWithEndpointAllOf struct for MetadataWithEndpointAllOf
type MetadataWithEndpointAllOf struct {
	// The public endpoint of the API Gateway instance.
	PublicEndpoint string `json:"publicEndpoint"`
}

// NewMetadataWithEndpointAllOf instantiates a new MetadataWithEndpointAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataWithEndpointAllOf(publicEndpoint string) *MetadataWithEndpointAllOf {
	this := MetadataWithEndpointAllOf{}

	this.PublicEndpoint = publicEndpoint

	return &this
}

// NewMetadataWithEndpointAllOfWithDefaults instantiates a new MetadataWithEndpointAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithEndpointAllOfWithDefaults() *MetadataWithEndpointAllOf {
	this := MetadataWithEndpointAllOf{}
	return &this
}

// GetPublicEndpoint returns the PublicEndpoint field value
func (o *MetadataWithEndpointAllOf) GetPublicEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicEndpoint
}

// GetPublicEndpointOk returns a tuple with the PublicEndpoint field value
// and a boolean to check if the value has been set.
func (o *MetadataWithEndpointAllOf) GetPublicEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicEndpoint, true
}

// SetPublicEndpoint sets field value
func (o *MetadataWithEndpointAllOf) SetPublicEndpoint(v string) {
	o.PublicEndpoint = v
}

func (o MetadataWithEndpointAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publicEndpoint"] = o.PublicEndpoint
	return toSerialize, nil
}

type NullableMetadataWithEndpointAllOf struct {
	value *MetadataWithEndpointAllOf
	isSet bool
}

func (v NullableMetadataWithEndpointAllOf) Get() *MetadataWithEndpointAllOf {
	return v.value
}

func (v *NullableMetadataWithEndpointAllOf) Set(val *MetadataWithEndpointAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataWithEndpointAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataWithEndpointAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataWithEndpointAllOf(val *MetadataWithEndpointAllOf) *NullableMetadataWithEndpointAllOf {
	return &NullableMetadataWithEndpointAllOf{value: val, isSet: true}
}

func (v NullableMetadataWithEndpointAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataWithEndpointAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
