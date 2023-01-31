/*
 * IONOS DBaaS MongoDB REST API
 *
 * With IONOS Cloud Database as a Service, you have the ability to quickly set up and manage a MongoDB database. You can also delete clusters, manage backups and users via the API.   MongoDB is an open source, cross-platform, document-oriented database program. Classified as a NoSQL database program, it uses JSON-like documents with optional schemas.  The MongoDB API allows you to create additional database clusters or modify existing ones. Both tools, the Data Center Designer (DCD) and the API use the same concepts consistently and are well suited for smooth and intuitive use.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongo

import (
	"encoding/json"
)

// ErrorResponse struct for ErrorResponse
type ErrorResponse struct {
	// The HTTP status code of the operation.
	HttpStatus *int32          `json:"httpStatus,omitempty"`
	Messages   *[]ErrorMessage `json:"messages,omitempty"`
}

// NewErrorResponse instantiates a new ErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponse() *ErrorResponse {
	this := ErrorResponse{}

	return &this
}

// NewErrorResponseWithDefaults instantiates a new ErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseWithDefaults() *ErrorResponse {
	this := ErrorResponse{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ErrorResponse) GetHttpStatus() *int32 {
	if o == nil {
		return nil
	}

	return o.HttpStatus

}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorResponse) GetHttpStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ErrorResponse) SetHttpStatus(v int32) {

	o.HttpStatus = &v

}

// HasHttpStatus returns a boolean if a field has been set.
func (o *ErrorResponse) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// GetMessages returns the Messages field value
// If the value is explicit nil, the zero value for []ErrorMessage will be returned
func (o *ErrorResponse) GetMessages() *[]ErrorMessage {
	if o == nil {
		return nil
	}

	return o.Messages

}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorResponse) GetMessagesOk() (*[]ErrorMessage, bool) {
	if o == nil {
		return nil, false
	}

	return o.Messages, true
}

// SetMessages sets field value
func (o *ErrorResponse) SetMessages(v []ErrorMessage) {

	o.Messages = &v

}

// HasMessages returns a boolean if a field has been set.
func (o *ErrorResponse) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

func (o ErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpStatus != nil {
		toSerialize["httpStatus"] = o.HttpStatus
	}

	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}

	return json.Marshal(toSerialize)
}

type NullableErrorResponse struct {
	value *ErrorResponse
	isSet bool
}

func (v NullableErrorResponse) Get() *ErrorResponse {
	return v.value
}

func (v *NullableErrorResponse) Set(val *ErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponse(val *ErrorResponse) *NullableErrorResponse {
	return &NullableErrorResponse{value: val, isSet: true}
}

func (v NullableErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
