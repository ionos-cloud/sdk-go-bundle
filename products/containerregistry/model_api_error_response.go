/*
 * Container Registry service
 *
 * ## Overview Container Registry service enables IONOS clients to manage docker and OCI compliant registries for use by their managed Kubernetes clusters. Use a Container Registry to ensure you have a privately accessed registry to efficiently support image pulls. ## Changelog ### 1.1.0  - Added new endpoints for Repositories  - Added new endpoints for Artifacts  - Added new endpoints for Vulnerabilities  - Added registry vulnerabilityScanning feature ### 1.2.0  - Added registry `apiSubnetAllowList` ### 1.2.1  - Amended `apiSubnetAllowList` Regex
 *
 * API version: 1.2.1
 * Contact: support@cloud.ionos.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package containerregistry

import (
	"encoding/json"
)

// checks if the ApiErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorResponse{}

// ApiErrorResponse struct for ApiErrorResponse
type ApiErrorResponse struct {
	HttpStatus int32             `json:"httpStatus"`
	Messages   []ApiErrorMessage `json:"messages"`
}

// NewApiErrorResponse instantiates a new ApiErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorResponse(httpStatus int32, messages []ApiErrorMessage) *ApiErrorResponse {
	this := ApiErrorResponse{}

	this.HttpStatus = httpStatus
	this.Messages = messages

	return &this
}

// NewApiErrorResponseWithDefaults instantiates a new ApiErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorResponseWithDefaults() *ApiErrorResponse {
	this := ApiErrorResponse{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value
func (o *ApiErrorResponse) GetHttpStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetHttpStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ApiErrorResponse) SetHttpStatus(v int32) {
	o.HttpStatus = v
}

// GetMessages returns the Messages field value
func (o *ApiErrorResponse) GetMessages() []ApiErrorMessage {
	if o == nil {
		var ret []ApiErrorMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetMessagesOk() ([]ApiErrorMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *ApiErrorResponse) SetMessages(v []ApiErrorMessage) {
	o.Messages = v
}

func (o ApiErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["httpStatus"] = o.HttpStatus
	toSerialize["messages"] = o.Messages
	return toSerialize, nil
}

type NullableApiErrorResponse struct {
	value *ApiErrorResponse
	isSet bool
}

func (v NullableApiErrorResponse) Get() *ApiErrorResponse {
	return v.value
}

func (v *NullableApiErrorResponse) Set(val *ApiErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorResponse(val *ApiErrorResponse) *NullableApiErrorResponse {
	return &NullableApiErrorResponse{value: val, isSet: true}
}

func (v NullableApiErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
