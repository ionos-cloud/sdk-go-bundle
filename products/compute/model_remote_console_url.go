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

// checks if the RemoteConsoleUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteConsoleUrl{}

// RemoteConsoleUrl struct for RemoteConsoleUrl
type RemoteConsoleUrl struct {
	// The remote console url with the jwToken parameter for access
	Url *string `json:"url,omitempty"`
}

// NewRemoteConsoleUrl instantiates a new RemoteConsoleUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteConsoleUrl() *RemoteConsoleUrl {
	this := RemoteConsoleUrl{}

	return &this
}

// NewRemoteConsoleUrlWithDefaults instantiates a new RemoteConsoleUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteConsoleUrlWithDefaults() *RemoteConsoleUrl {
	this := RemoteConsoleUrl{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RemoteConsoleUrl) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConsoleUrl) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RemoteConsoleUrl) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RemoteConsoleUrl) SetUrl(v string) {
	o.Url = &v
}

func (o RemoteConsoleUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteConsoleUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableRemoteConsoleUrl struct {
	value *RemoteConsoleUrl
	isSet bool
}

func (v NullableRemoteConsoleUrl) Get() *RemoteConsoleUrl {
	return v.value
}

func (v *NullableRemoteConsoleUrl) Set(val *RemoteConsoleUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteConsoleUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteConsoleUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteConsoleUrl(val *RemoteConsoleUrl) *NullableRemoteConsoleUrl {
	return &NullableRemoteConsoleUrl{value: val, isSet: true}
}

func (v NullableRemoteConsoleUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteConsoleUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
