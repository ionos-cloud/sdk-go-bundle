/*
 * IONOS Cloud - Object Storage Management API
 *
 * Object Storage Management API is a RESTful API that manages the object storage service configuration for IONOS Cloud.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstoragemanagement

import (
	"encoding/json"
)

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error{}

// Error The Error object is used to represent an error response from the API.
type Error struct {
	// The HTTP status code of the operation.
	HttpStatus *int32 `json:"httpStatus,omitempty"`
	// A list of error messages.
	Messages []ErrorMessages `json:"messages,omitempty"`
}

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError() *Error {
	this := Error{}

	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value if set, zero value otherwise.
func (o *Error) GetHttpStatus() int32 {
	if o == nil || IsNil(o.HttpStatus) {
		var ret int32
		return ret
	}
	return *o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetHttpStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.HttpStatus) {
		return nil, false
	}
	return o.HttpStatus, true
}

// HasHttpStatus returns a boolean if a field has been set.
func (o *Error) HasHttpStatus() bool {
	if o != nil && !IsNil(o.HttpStatus) {
		return true
	}

	return false
}

// SetHttpStatus gets a reference to the given int32 and assigns it to the HttpStatus field.
func (o *Error) SetHttpStatus(v int32) {
	o.HttpStatus = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Error) GetMessages() []ErrorMessages {
	if o == nil || IsNil(o.Messages) {
		var ret []ErrorMessages
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetMessagesOk() ([]ErrorMessages, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Error) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ErrorMessages and assigns it to the Messages field.
func (o *Error) SetMessages(v []ErrorMessages) {
	o.Messages = v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpStatus) {
		toSerialize["httpStatus"] = o.HttpStatus
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	return toSerialize, nil
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
