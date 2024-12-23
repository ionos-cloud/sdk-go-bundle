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

// checks if the RouteRead type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteRead{}

// RouteRead struct for RouteRead
type RouteRead struct {
	// The ID (UUID) of the Route.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The URL of the Route.
	Href       string               `json:"href"`
	Metadata   MetadataWithEndpoint `json:"metadata"`
	Properties Route                `json:"properties"`
}

// NewRouteRead instantiates a new RouteRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteRead(id string, type_ string, href string, metadata MetadataWithEndpoint, properties Route) *RouteRead {
	this := RouteRead{}

	this.Id = id
	this.Type = type_
	this.Href = href
	this.Metadata = metadata
	this.Properties = properties

	return &this
}

// NewRouteReadWithDefaults instantiates a new RouteRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteReadWithDefaults() *RouteRead {
	this := RouteRead{}
	return &this
}

// GetId returns the Id field value
func (o *RouteRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RouteRead) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RouteRead) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *RouteRead) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RouteRead) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RouteRead) SetType(v string) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *RouteRead) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *RouteRead) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *RouteRead) SetHref(v string) {
	o.Href = v
}

// GetMetadata returns the Metadata field value
func (o *RouteRead) GetMetadata() MetadataWithEndpoint {
	if o == nil {
		var ret MetadataWithEndpoint
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *RouteRead) GetMetadataOk() (*MetadataWithEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *RouteRead) SetMetadata(v MetadataWithEndpoint) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *RouteRead) GetProperties() Route {
	if o == nil {
		var ret Route
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *RouteRead) GetPropertiesOk() (*Route, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *RouteRead) SetProperties(v Route) {
	o.Properties = v
}

func (o RouteRead) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["metadata"] = o.Metadata
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableRouteRead struct {
	value *RouteRead
	isSet bool
}

func (v NullableRouteRead) Get() *RouteRead {
	return v.value
}

func (v *NullableRouteRead) Set(val *RouteRead) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteRead) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteRead(val *RouteRead) *NullableRouteRead {
	return &NullableRouteRead{value: val, isSet: true}
}

func (v NullableRouteRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
