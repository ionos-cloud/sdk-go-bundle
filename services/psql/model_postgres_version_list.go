/*
 * IONOS DBaaS REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package psql

import (
	"encoding/json"
)

// PostgresVersionList List of PostgreSQL versions.
type PostgresVersionList struct {
	Data *[]PostgresVersionListData `json:"data,omitempty"`
}

// NewPostgresVersionList instantiates a new PostgresVersionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresVersionList() *PostgresVersionList {
	this := PostgresVersionList{}

	return &this
}

// NewPostgresVersionListWithDefaults instantiates a new PostgresVersionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresVersionListWithDefaults() *PostgresVersionList {
	this := PostgresVersionList{}
	return &this
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for []PostgresVersionListData will be returned
func (o *PostgresVersionList) GetData() *[]PostgresVersionListData {
	if o == nil {
		return nil
	}

	return o.Data

}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostgresVersionList) GetDataOk() (*[]PostgresVersionListData, bool) {
	if o == nil {
		return nil, false
	}

	return o.Data, true
}

// SetData sets field value
func (o *PostgresVersionList) SetData(v []PostgresVersionListData) {

	o.Data = &v

}

// HasData returns a boolean if a field has been set.
func (o *PostgresVersionList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

func (o PostgresVersionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	return json.Marshal(toSerialize)
}

type NullablePostgresVersionList struct {
	value *PostgresVersionList
	isSet bool
}

func (v NullablePostgresVersionList) Get() *PostgresVersionList {
	return v.value
}

func (v *NullablePostgresVersionList) Set(val *PostgresVersionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresVersionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresVersionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresVersionList(val *PostgresVersionList) *NullablePostgresVersionList {
	return &NullablePostgresVersionList{value: val, isSet: true}
}

func (v NullablePostgresVersionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresVersionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
