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

// checks if the PostgresVersionListData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostgresVersionListData{}

// PostgresVersionListData struct for PostgresVersionListData
type PostgresVersionListData struct {
	Name *string `json:"name,omitempty"`
}

// NewPostgresVersionListData instantiates a new PostgresVersionListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresVersionListData() *PostgresVersionListData {
	this := PostgresVersionListData{}

	return &this
}

// NewPostgresVersionListDataWithDefaults instantiates a new PostgresVersionListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresVersionListDataWithDefaults() *PostgresVersionListData {
	this := PostgresVersionListData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostgresVersionListData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresVersionListData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostgresVersionListData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostgresVersionListData) SetName(v string) {
	o.Name = &v
}

func (o PostgresVersionListData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePostgresVersionListData struct {
	value *PostgresVersionListData
	isSet bool
}

func (v NullablePostgresVersionListData) Get() *PostgresVersionListData {
	return v.value
}

func (v *NullablePostgresVersionListData) Set(val *PostgresVersionListData) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresVersionListData) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresVersionListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresVersionListData(val *PostgresVersionListData) *NullablePostgresVersionListData {
	return &NullablePostgresVersionListData{value: val, isSet: true}
}

func (v NullablePostgresVersionListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresVersionListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
