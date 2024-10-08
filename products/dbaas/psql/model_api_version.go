/*
 * IONOS DBaaS PostgreSQL REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional PostgreSQL database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psql

import (
	"encoding/json"
)

// checks if the APIVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIVersion{}

// APIVersion struct for APIVersion
type APIVersion struct {
	Name       *string `json:"name,omitempty"`
	SwaggerUrl *string `json:"swaggerUrl,omitempty"`
}

// NewAPIVersion instantiates a new APIVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIVersion() *APIVersion {
	this := APIVersion{}

	return &this
}

// NewAPIVersionWithDefaults instantiates a new APIVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIVersionWithDefaults() *APIVersion {
	this := APIVersion{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *APIVersion) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIVersion) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *APIVersion) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *APIVersion) SetName(v string) {
	o.Name = &v
}

// GetSwaggerUrl returns the SwaggerUrl field value if set, zero value otherwise.
func (o *APIVersion) GetSwaggerUrl() string {
	if o == nil || IsNil(o.SwaggerUrl) {
		var ret string
		return ret
	}
	return *o.SwaggerUrl
}

// GetSwaggerUrlOk returns a tuple with the SwaggerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIVersion) GetSwaggerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SwaggerUrl) {
		return nil, false
	}
	return o.SwaggerUrl, true
}

// HasSwaggerUrl returns a boolean if a field has been set.
func (o *APIVersion) HasSwaggerUrl() bool {
	if o != nil && !IsNil(o.SwaggerUrl) {
		return true
	}

	return false
}

// SetSwaggerUrl gets a reference to the given string and assigns it to the SwaggerUrl field.
func (o *APIVersion) SetSwaggerUrl(v string) {
	o.SwaggerUrl = &v
}

func (o APIVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SwaggerUrl) {
		toSerialize["swaggerUrl"] = o.SwaggerUrl
	}
	return toSerialize, nil
}

type NullableAPIVersion struct {
	value *APIVersion
	isSet bool
}

func (v NullableAPIVersion) Get() *APIVersion {
	return v.value
}

func (v *NullableAPIVersion) Set(val *APIVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIVersion(val *APIVersion) *NullableAPIVersion {
	return &NullableAPIVersion{value: val, isSet: true}
}

func (v NullableAPIVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
