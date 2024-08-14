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

// checks if the S3ObjectStorageSSO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3ObjectStorageSSO{}

// S3ObjectStorageSSO struct for S3ObjectStorageSSO
type S3ObjectStorageSSO struct {
	// The S3 object storage single sign on url
	SsoUrl *string `json:"ssoUrl,omitempty"`
}

// NewS3ObjectStorageSSO instantiates a new S3ObjectStorageSSO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ObjectStorageSSO() *S3ObjectStorageSSO {
	this := S3ObjectStorageSSO{}

	return &this
}

// NewS3ObjectStorageSSOWithDefaults instantiates a new S3ObjectStorageSSO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ObjectStorageSSOWithDefaults() *S3ObjectStorageSSO {
	this := S3ObjectStorageSSO{}
	return &this
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise.
func (o *S3ObjectStorageSSO) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl) {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ObjectStorageSSO) GetSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUrl) {
		return nil, false
	}
	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *S3ObjectStorageSSO) HasSsoUrl() bool {
	if o != nil && !IsNil(o.SsoUrl) {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *S3ObjectStorageSSO) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

func (o S3ObjectStorageSSO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3ObjectStorageSSO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsoUrl) {
		toSerialize["ssoUrl"] = o.SsoUrl
	}
	return toSerialize, nil
}

type NullableS3ObjectStorageSSO struct {
	value *S3ObjectStorageSSO
	isSet bool
}

func (v NullableS3ObjectStorageSSO) Get() *S3ObjectStorageSSO {
	return v.value
}

func (v *NullableS3ObjectStorageSSO) Set(val *S3ObjectStorageSSO) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ObjectStorageSSO) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ObjectStorageSSO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ObjectStorageSSO(val *S3ObjectStorageSSO) *NullableS3ObjectStorageSSO {
	return &NullableS3ObjectStorageSSO{value: val, isSet: true}
}

func (v NullableS3ObjectStorageSSO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ObjectStorageSSO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
