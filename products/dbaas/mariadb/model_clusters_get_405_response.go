/*
 * IONOS DBaaS MariaDB REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional MariaDB database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mariadb

import (
	"encoding/json"
)

// checks if the ClustersGet405Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClustersGet405Response{}

// ClustersGet405Response struct for ClustersGet405Response
type ClustersGet405Response struct {
	// The HTTP status code of the operation.
	HttpStatus int32          `json:"httpStatus"`
	Messages   []ErrorMessage `json:"messages"`
}

// NewClustersGet405Response instantiates a new ClustersGet405Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClustersGet405Response(httpStatus int32, messages []ErrorMessage) *ClustersGet405Response {
	this := ClustersGet405Response{}

	this.HttpStatus = httpStatus
	this.Messages = messages

	return &this
}

// NewClustersGet405ResponseWithDefaults instantiates a new ClustersGet405Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClustersGet405ResponseWithDefaults() *ClustersGet405Response {
	this := ClustersGet405Response{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value
func (o *ClustersGet405Response) GetHttpStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
func (o *ClustersGet405Response) GetHttpStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ClustersGet405Response) SetHttpStatus(v int32) {
	o.HttpStatus = v
}

// GetMessages returns the Messages field value
func (o *ClustersGet405Response) GetMessages() []ErrorMessage {
	if o == nil {
		var ret []ErrorMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ClustersGet405Response) GetMessagesOk() ([]ErrorMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *ClustersGet405Response) SetMessages(v []ErrorMessage) {
	o.Messages = v
}

func (o ClustersGet405Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["httpStatus"] = o.HttpStatus
	toSerialize["messages"] = o.Messages
	return toSerialize, nil
}

type NullableClustersGet405Response struct {
	value *ClustersGet405Response
	isSet bool
}

func (v NullableClustersGet405Response) Get() *ClustersGet405Response {
	return v.value
}

func (v *NullableClustersGet405Response) Set(val *ClustersGet405Response) {
	v.value = val
	v.isSet = true
}

func (v NullableClustersGet405Response) IsSet() bool {
	return v.isSet
}

func (v *NullableClustersGet405Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClustersGet405Response(val *ClustersGet405Response) *NullableClustersGet405Response {
	return &NullableClustersGet405Response{value: val, isSet: true}
}

func (v NullableClustersGet405Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClustersGet405Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
