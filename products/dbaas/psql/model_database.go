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

// checks if the Database type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Database{}

// Database struct for Database
type Database struct {
	Properties DatabaseProperties `json:"properties"`
}

// NewDatabase instantiates a new Database object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabase(properties DatabaseProperties) *Database {
	this := Database{}

	this.Properties = properties

	return &this
}

// NewDatabaseWithDefaults instantiates a new Database object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseWithDefaults() *Database {
	this := Database{}
	return &this
}

// GetProperties returns the Properties field value
func (o *Database) GetProperties() DatabaseProperties {
	if o == nil {
		var ret DatabaseProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *Database) GetPropertiesOk() (*DatabaseProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *Database) SetProperties(v DatabaseProperties) {
	o.Properties = v
}

func (o Database) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Database) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsZero(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableDatabase struct {
	value *Database
	isSet bool
}

func (v NullableDatabase) Get() *Database {
	return v.value
}

func (v *NullableDatabase) Set(val *Database) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabase(val *Database) *NullableDatabase {
	return &NullableDatabase{value: val, isSet: true}
}

func (v NullableDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
