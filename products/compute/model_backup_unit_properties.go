/*
 * CLOUD API
 *
 *  IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// checks if the BackupUnitProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackupUnitProperties{}

// BackupUnitProperties struct for BackupUnitProperties
type BackupUnitProperties struct {
	// The name of the  resource (alphanumeric characters only).
	Name string `json:"name"`
	// The password associated with that resource.
	Password *string `json:"password,omitempty"`
	// The email associated with the backup unit. Bear in mind that this email does not be the same email as of the user.
	Email *string `json:"email,omitempty"`
}

// NewBackupUnitProperties instantiates a new BackupUnitProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupUnitProperties(name string) *BackupUnitProperties {
	this := BackupUnitProperties{}

	this.Name = name

	return &this
}

// NewBackupUnitPropertiesWithDefaults instantiates a new BackupUnitProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupUnitPropertiesWithDefaults() *BackupUnitProperties {
	this := BackupUnitProperties{}
	return &this
}

// GetName returns the Name field value
func (o *BackupUnitProperties) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackupUnitProperties) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BackupUnitProperties) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *BackupUnitProperties) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupUnitProperties) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *BackupUnitProperties) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *BackupUnitProperties) SetPassword(v string) {
	o.Password = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BackupUnitProperties) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupUnitProperties) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BackupUnitProperties) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BackupUnitProperties) SetEmail(v string) {
	o.Email = &v
}

func (o BackupUnitProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableBackupUnitProperties struct {
	value *BackupUnitProperties
	isSet bool
}

func (v NullableBackupUnitProperties) Get() *BackupUnitProperties {
	return v.value
}

func (v *NullableBackupUnitProperties) Set(val *BackupUnitProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupUnitProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupUnitProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupUnitProperties(val *BackupUnitProperties) *NullableBackupUnitProperties {
	return &NullableBackupUnitProperties{value: val, isSet: true}
}

func (v NullableBackupUnitProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupUnitProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
