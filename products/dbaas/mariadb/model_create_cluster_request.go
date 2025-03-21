/*
 * IONOS DBaaS MariaDB REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional MariaDB database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mariadb

import (
	"encoding/json"
)

// checks if the CreateClusterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClusterRequest{}

// CreateClusterRequest Request payload with all data needed to create a new MariaDB cluster.
type CreateClusterRequest struct {
	Properties *CreateClusterProperties `json:"properties,omitempty"`
}

// NewCreateClusterRequest instantiates a new CreateClusterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequest() *CreateClusterRequest {
	this := CreateClusterRequest{}

	return &this
}

// NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestWithDefaults() *CreateClusterRequest {
	this := CreateClusterRequest{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CreateClusterRequest) GetProperties() CreateClusterProperties {
	if o == nil || IsNil(o.Properties) {
		var ret CreateClusterProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateClusterRequest) GetPropertiesOk() (*CreateClusterProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CreateClusterRequest) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given CreateClusterProperties and assigns it to the Properties field.
func (o *CreateClusterRequest) SetProperties(v CreateClusterProperties) {
	o.Properties = &v
}

func (o CreateClusterRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClusterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableCreateClusterRequest struct {
	value *CreateClusterRequest
	isSet bool
}

func (v NullableCreateClusterRequest) Get() *CreateClusterRequest {
	return v.value
}

func (v *NullableCreateClusterRequest) Set(val *CreateClusterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequest(val *CreateClusterRequest) *NullableCreateClusterRequest {
	return &NullableCreateClusterRequest{value: val, isSet: true}
}

func (v NullableCreateClusterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
