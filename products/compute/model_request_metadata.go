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

	"time"
)

// RequestMetadata struct for RequestMetadata
type RequestMetadata struct {
	// The last time the resource was created.
	CreatedDate *IonosTime `json:"createdDate,omitempty"`
	// The user who created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Resource's Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11  Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter.
	Etag          *string        `json:"etag,omitempty"`
	RequestStatus *RequestStatus `json:"requestStatus,omitempty"`
}

// NewRequestMetadata instantiates a new RequestMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMetadata() *RequestMetadata {
	this := RequestMetadata{}

	return &this
}

// NewRequestMetadataWithDefaults instantiates a new RequestMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMetadataWithDefaults() *RequestMetadata {
	this := RequestMetadata{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *RequestMetadata) GetCreatedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.CreatedDate == nil {
		return nil
	}
	return &o.CreatedDate.Time

}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMetadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.CreatedDate == nil {
		return nil, false
	}
	return &o.CreatedDate.Time, true

}

// SetCreatedDate sets field value
func (o *RequestMetadata) SetCreatedDate(v time.Time) {

	o.CreatedDate = &IonosTime{v}

}

// HasCreatedDate returns a boolean if a field has been set.
func (o *RequestMetadata) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// GetCreatedBy returns the CreatedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RequestMetadata) GetCreatedBy() *string {
	if o == nil {
		return nil
	}

	return o.CreatedBy

}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMetadata) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestMetadata) SetCreatedBy(v string) {

	o.CreatedBy = &v

}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RequestMetadata) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// GetEtag returns the Etag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RequestMetadata) GetEtag() *string {
	if o == nil {
		return nil
	}

	return o.Etag

}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMetadata) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Etag, true
}

// SetEtag sets field value
func (o *RequestMetadata) SetEtag(v string) {

	o.Etag = &v

}

// HasEtag returns a boolean if a field has been set.
func (o *RequestMetadata) HasEtag() bool {
	if o != nil && o.Etag != nil {
		return true
	}

	return false
}

// GetRequestStatus returns the RequestStatus field value
// If the value is explicit nil, the zero value for RequestStatus will be returned
func (o *RequestMetadata) GetRequestStatus() *RequestStatus {
	if o == nil {
		return nil
	}

	return o.RequestStatus

}

// GetRequestStatusOk returns a tuple with the RequestStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMetadata) GetRequestStatusOk() (*RequestStatus, bool) {
	if o == nil {
		return nil, false
	}

	return o.RequestStatus, true
}

// SetRequestStatus sets field value
func (o *RequestMetadata) SetRequestStatus(v RequestStatus) {

	o.RequestStatus = &v

}

// HasRequestStatus returns a boolean if a field has been set.
func (o *RequestMetadata) HasRequestStatus() bool {
	if o != nil && o.RequestStatus != nil {
		return true
	}

	return false
}

func (o RequestMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}

	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}

	if o.Etag != nil {
		toSerialize["etag"] = o.Etag
	}

	if o.RequestStatus != nil {
		toSerialize["requestStatus"] = o.RequestStatus
	}

	return json.Marshal(toSerialize)
}

type NullableRequestMetadata struct {
	value *RequestMetadata
	isSet bool
}

func (v NullableRequestMetadata) Get() *RequestMetadata {
	return v.value
}

func (v *NullableRequestMetadata) Set(val *RequestMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMetadata(val *RequestMetadata) *NullableRequestMetadata {
	return &NullableRequestMetadata{value: val, isSet: true}
}

func (v NullableRequestMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
