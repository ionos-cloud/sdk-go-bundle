/*
 * Certificate Manager Service API
 *
 * Using the Certificate Manager Service, you can conveniently provision and manage SSL certificates  with IONOS services and your internal connected resources.   For the [Application Load Balancer](https://api.ionos.com/docs/cloud/v6/#Application-Load-Balancers-get-datacenters-datacenterId-applicationloadbalancers), you usually need a certificate to encrypt your HTTPS traffic. The service provides the basic functions of uploading and deleting your certificates for this purpose.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cert

import (
	"encoding/json"
)

// checks if the AutoCertificateCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoCertificateCreate{}

// AutoCertificateCreate struct for AutoCertificateCreate
type AutoCertificateCreate struct {
	// Metadata
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
	Properties AutoCertificate        `json:"properties"`
}

// NewAutoCertificateCreate instantiates a new AutoCertificateCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoCertificateCreate(properties AutoCertificate) *AutoCertificateCreate {
	this := AutoCertificateCreate{}

	this.Properties = properties

	return &this
}

// NewAutoCertificateCreateWithDefaults instantiates a new AutoCertificateCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoCertificateCreateWithDefaults() *AutoCertificateCreate {
	this := AutoCertificateCreate{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AutoCertificateCreate) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoCertificateCreate) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AutoCertificateCreate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AutoCertificateCreate) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *AutoCertificateCreate) GetProperties() AutoCertificate {
	if o == nil {
		var ret AutoCertificate
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *AutoCertificateCreate) GetPropertiesOk() (*AutoCertificate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *AutoCertificateCreate) SetProperties(v AutoCertificate) {
	o.Properties = v
}

func (o AutoCertificateCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoCertificateCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

type NullableAutoCertificateCreate struct {
	value *AutoCertificateCreate
	isSet bool
}

func (v NullableAutoCertificateCreate) Get() *AutoCertificateCreate {
	return v.value
}

func (v *NullableAutoCertificateCreate) Set(val *AutoCertificateCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoCertificateCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoCertificateCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoCertificateCreate(val *AutoCertificateCreate) *NullableAutoCertificateCreate {
	return &NullableAutoCertificateCreate{value: val, isSet: true}
}

func (v NullableAutoCertificateCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoCertificateCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
