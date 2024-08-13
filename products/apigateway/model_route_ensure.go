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

// checks if the RouteEnsure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteEnsure{}

// RouteEnsure struct for RouteEnsure
type RouteEnsure struct {
	// The ID (UUID) of the Route.
	Id string `json:"id"`
	// Metadata
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Properties Route                  `json:"properties"`
}

// NewRouteEnsure instantiates a new RouteEnsure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteEnsure(id string, properties Route) *RouteEnsure {
	this := RouteEnsure{}

	this.Id = id
	this.Properties = properties

	return &this
}

// NewRouteEnsureWithDefaults instantiates a new RouteEnsure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteEnsureWithDefaults() *RouteEnsure {
	this := RouteEnsure{}
	return &this
}

// GetId returns the Id field value
func (o *RouteEnsure) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RouteEnsure) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RouteEnsure) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RouteEnsure) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteEnsure) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RouteEnsure) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *RouteEnsure) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *RouteEnsure) GetProperties() Route {
	if o == nil {
		var ret Route
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *RouteEnsure) GetPropertiesOk() (*Route, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *RouteEnsure) SetProperties(v Route) {
	o.Properties = v
}

func (o RouteEnsure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteEnsure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsZero(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsZero(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableRouteEnsure struct {
	value *RouteEnsure
	isSet bool
}

func (v NullableRouteEnsure) Get() *RouteEnsure {
	return v.value
}

func (v *NullableRouteEnsure) Set(val *RouteEnsure) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteEnsure) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteEnsure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteEnsure(val *RouteEnsure) *NullableRouteEnsure {
	return &NullableRouteEnsure{value: val, isSet: true}
}

func (v NullableRouteEnsure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteEnsure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
