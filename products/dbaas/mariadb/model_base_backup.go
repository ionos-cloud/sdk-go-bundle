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

	"time"
)

// checks if the BaseBackup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseBackup{}

// BaseBackup A single base backup.
type BaseBackup struct {
	// The ISO 8601 creation timestamp.
	Created *IonosTime `json:"created,omitempty"`
	// The size of the backup in MiB. This is the size of the binary backup file that was stored.
	Size *int32 `json:"size,omitempty"`
}

// NewBaseBackup instantiates a new BaseBackup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseBackup() *BaseBackup {
	this := BaseBackup{}

	return &this
}

// NewBaseBackupWithDefaults instantiates a new BaseBackup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseBackupWithDefaults() *BaseBackup {
	this := BaseBackup{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *BaseBackup) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return o.Created.Time
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseBackup) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return &o.Created.Time, true
}

// HasCreated returns a boolean if a field has been set.
func (o *BaseBackup) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *BaseBackup) SetCreated(v time.Time) {
	o.Created = &IonosTime{v}
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BaseBackup) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseBackup) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BaseBackup) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *BaseBackup) SetSize(v int32) {
	o.Size = &v
}

func (o BaseBackup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableBaseBackup struct {
	value *BaseBackup
	isSet bool
}

func (v NullableBaseBackup) Get() *BaseBackup {
	return v.value
}

func (v *NullableBaseBackup) Set(val *BaseBackup) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseBackup(val *BaseBackup) *NullableBaseBackup {
	return &NullableBaseBackup{value: val, isSet: true}
}

func (v NullableBaseBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
