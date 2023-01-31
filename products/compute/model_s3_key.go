/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package compute

import (
	"encoding/json"
)

// S3Key struct for S3Key
type S3Key struct {
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The type of the resource.
	Type *Type `json:"type,omitempty"`
	// URL to the object representation (absolute path).
	Href       *string          `json:"href,omitempty"`
	Metadata   *S3KeyMetadata   `json:"metadata,omitempty"`
	Properties *S3KeyProperties `json:"properties"`
}

// NewS3Key instantiates a new S3Key object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3Key(properties S3KeyProperties) *S3Key {
	this := S3Key{}

	this.Properties = &properties

	return &this
}

// NewS3KeyWithDefaults instantiates a new S3Key object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3KeyWithDefaults() *S3Key {
	this := S3Key{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *S3Key) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Key) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *S3Key) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *S3Key) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for Type will be returned
func (o *S3Key) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Key) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *S3Key) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *S3Key) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *S3Key) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Key) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *S3Key) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *S3Key) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for S3KeyMetadata will be returned
func (o *S3Key) GetMetadata() *S3KeyMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Key) GetMetadataOk() (*S3KeyMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *S3Key) SetMetadata(v S3KeyMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *S3Key) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for S3KeyProperties will be returned
func (o *S3Key) GetProperties() *S3KeyProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3Key) GetPropertiesOk() (*S3KeyProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *S3Key) SetProperties(v S3KeyProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *S3Key) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o S3Key) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableS3Key struct {
	value *S3Key
	isSet bool
}

func (v NullableS3Key) Get() *S3Key {
	return v.value
}

func (v *NullableS3Key) Set(val *S3Key) {
	v.value = val
	v.isSet = true
}

func (v NullableS3Key) IsSet() bool {
	return v.isSet
}

func (v *NullableS3Key) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3Key(val *S3Key) *NullableS3Key {
	return &NullableS3Key{value: val, isSet: true}
}

func (v NullableS3Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3Key) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
