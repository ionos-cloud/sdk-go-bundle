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

// checks if the RequestStatusMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestStatusMetadata{}

// RequestStatusMetadata struct for RequestStatusMetadata
type RequestStatusMetadata struct {
	Status  *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	// Resource's Entity Tag as defined in http://www.w3.org/Protocols/rfc2616/rfc2616-sec3.html#sec3.11  Entity Tag is also added as an 'ETag response header to requests which don't use 'depth' parameter.
	Etag    *string         `json:"etag,omitempty"`
	Targets []RequestTarget `json:"targets,omitempty"`
}

// NewRequestStatusMetadata instantiates a new RequestStatusMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestStatusMetadata() *RequestStatusMetadata {
	this := RequestStatusMetadata{}

	return &this
}

// NewRequestStatusMetadataWithDefaults instantiates a new RequestStatusMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestStatusMetadataWithDefaults() *RequestStatusMetadata {
	this := RequestStatusMetadata{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestStatusMetadata) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestStatusMetadata) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestStatusMetadata) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RequestStatusMetadata) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RequestStatusMetadata) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestStatusMetadata) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RequestStatusMetadata) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RequestStatusMetadata) SetMessage(v string) {
	o.Message = &v
}

// GetEtag returns the Etag field value if set, zero value otherwise.
func (o *RequestStatusMetadata) GetEtag() string {
	if o == nil || IsNil(o.Etag) {
		var ret string
		return ret
	}
	return *o.Etag
}

// GetEtagOk returns a tuple with the Etag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestStatusMetadata) GetEtagOk() (*string, bool) {
	if o == nil || IsNil(o.Etag) {
		return nil, false
	}
	return o.Etag, true
}

// HasEtag returns a boolean if a field has been set.
func (o *RequestStatusMetadata) HasEtag() bool {
	if o != nil && !IsNil(o.Etag) {
		return true
	}

	return false
}

// SetEtag gets a reference to the given string and assigns it to the Etag field.
func (o *RequestStatusMetadata) SetEtag(v string) {
	o.Etag = &v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *RequestStatusMetadata) GetTargets() []RequestTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []RequestTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestStatusMetadata) GetTargetsOk() ([]RequestTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *RequestStatusMetadata) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []RequestTarget and assigns it to the Targets field.
func (o *RequestStatusMetadata) SetTargets(v []RequestTarget) {
	o.Targets = v
}

func (o RequestStatusMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Etag) {
		toSerialize["etag"] = o.Etag
	}
	if !IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	return toSerialize, nil
}

type NullableRequestStatusMetadata struct {
	value *RequestStatusMetadata
	isSet bool
}

func (v NullableRequestStatusMetadata) Get() *RequestStatusMetadata {
	return v.value
}

func (v *NullableRequestStatusMetadata) Set(val *RequestStatusMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestStatusMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestStatusMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestStatusMetadata(val *RequestStatusMetadata) *NullableRequestStatusMetadata {
	return &NullableRequestStatusMetadata{value: val, isSet: true}
}

func (v NullableRequestStatusMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestStatusMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
