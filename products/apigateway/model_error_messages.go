/*
 * IONOS Cloud - API Gateway
 *
 * API Gateway is an application that acts as a \"front door\" for backend services and APIs, handling client requests and routing them to the appropriate backend.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apigateway

import (
	"encoding/json"
)

// checks if the ErrorMessages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorMessages{}

// ErrorMessages struct for ErrorMessages
type ErrorMessages struct {
	// Application internal error code
	ErrorCode *string `json:"errorCode,omitempty"`
	// A human readable explanation specific to this occurrence of the problem.
	Message *string `json:"message,omitempty"`
}

// NewErrorMessages instantiates a new ErrorMessages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessages() *ErrorMessages {
	this := ErrorMessages{}

	return &this
}

// NewErrorMessagesWithDefaults instantiates a new ErrorMessages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessagesWithDefaults() *ErrorMessages {
	this := ErrorMessages{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ErrorMessages) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessages) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ErrorMessages) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ErrorMessages) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorMessages) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessages) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorMessages) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorMessages) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorMessages) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorMessages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableErrorMessages struct {
	value *ErrorMessages
	isSet bool
}

func (v NullableErrorMessages) Get() *ErrorMessages {
	return v.value
}

func (v *NullableErrorMessages) Set(val *ErrorMessages) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessages) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessages(val *ErrorMessages) *NullableErrorMessages {
	return &NullableErrorMessages{value: val, isSet: true}
}

func (v NullableErrorMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
