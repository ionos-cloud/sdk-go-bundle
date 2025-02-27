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

// checks if the S3Bucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &S3Bucket{}

// S3Bucket struct for S3Bucket
type S3Bucket struct {
	// The name of the Object storage bucket.
	Name string `json:"name"`
}

// NewS3Bucket instantiates a new S3Bucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3Bucket(name string) *S3Bucket {
	this := S3Bucket{}

	this.Name = name

	return &this
}

// NewS3BucketWithDefaults instantiates a new S3Bucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3BucketWithDefaults() *S3Bucket {
	this := S3Bucket{}
	return &this
}

// GetName returns the Name field value
func (o *S3Bucket) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *S3Bucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *S3Bucket) SetName(v string) {
	o.Name = v
}

func (o S3Bucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableS3Bucket struct {
	value *S3Bucket
	isSet bool
}

func (v NullableS3Bucket) Get() *S3Bucket {
	return v.value
}

func (v *NullableS3Bucket) Set(val *S3Bucket) {
	v.value = val
	v.isSet = true
}

func (v NullableS3Bucket) IsSet() bool {
	return v.isSet
}

func (v *NullableS3Bucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3Bucket(val *S3Bucket) *NullableS3Bucket {
	return &NullableS3Bucket{value: val, isSet: true}
}

func (v NullableS3Bucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3Bucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
